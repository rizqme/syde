package cli

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/client"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/feedloop/syde/skill"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Manage design plans",
}

// printIndented prints a multi-line string with each line prefixed by indent.
func printIndented(text, indent string) {
	for _, line := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		fmt.Printf("%s%s\n", indent, line)
	}
}

var (
	planCreateBackground string
	planCreateObjective  string
	planCreateScope      string
	planCreateDesign     string
)

var planCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		plan := &model.PlanEntity{
			BaseEntity: model.BaseEntity{
				Kind: model.KindPlan,
				Name: name,
			},
			PlanStatus: model.PlanDraft,
			Background: planCreateBackground,
			Objective:  planCreateObjective,
			PlanScope:  planCreateScope,
			Design:     planCreateDesign,
			Source:     "manual",
			CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		}

		filePath, err := store.Create(plan, "")
		if err != nil {
			return err
		}

		fmt.Printf("Created plan: %s\n", name)
		fmt.Printf("  ID: %s\n", plan.ID)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

var planListCmd = &cobra.Command{
	Use:   "list",
	Short: "List plans",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		plans, err := store.List(model.KindPlan)
		if err != nil {
			return err
		}

		if len(plans) == 0 {
			fmt.Println("No plans found.")
			return nil
		}

		for _, ewb := range plans {
			p, ok := ewb.Entity.(*model.PlanEntity)
			if !ok {
				continue
			}
			progress := p.Progress()
			completedPhases := 0
			for _, ph := range p.Phases {
				if ph.Status == model.PhaseCompleted {
					completedPhases++
				}
			}
			fmt.Printf("  %-30s %-12s %d/%d phases (%.0f%%)\n",
				p.Name, p.PlanStatus, completedPhases, len(p.Phases), progress)
		}
		return nil
	},
}

var planShowFull bool

var planShowCmd = &cobra.Command{
	Use:   "show <slug>",
	Short: "Show plan details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, slug)
		if err != nil {
			return fmt.Errorf("plan not found: %s", slug)
		}

		p := entity.(*model.PlanEntity)
		fmt.Printf("═══ Plan: %s ═══\n", p.Name)
		fmt.Printf("Status: %s | Progress: %.0f%%\n", p.PlanStatus, p.Progress())
		if p.CreatedAt != "" {
			fmt.Printf("Created: %s\n", p.CreatedAt)
		}
		if p.ApprovedAt != "" {
			fmt.Printf("Approved: %s\n", p.ApprovedAt)
		}
		if p.CompletedAt != "" {
			fmt.Printf("Completed: %s\n", p.CompletedAt)
		}
		fmt.Println()

		if p.Background != "" {
			fmt.Println("BACKGROUND")
			printIndented(p.Background, "  ")
			fmt.Println()
		}
		if p.Objective != "" {
			fmt.Println("OBJECTIVE")
			printIndented(p.Objective, "  ")
			fmt.Println()
		}
		if p.PlanScope != "" {
			fmt.Println("SCOPE")
			printIndented(p.PlanScope, "  ")
			fmt.Println()
		}

		if len(p.Phases) > 0 {
			impl := p.Phases

			if len(impl) > 0 {
				fmt.Println("Implementation Phases:")
				// Build parent→children map
				children := map[string][]model.PlanPhase{}
				var roots []model.PlanPhase
				for _, ph := range impl {
					if ph.ParentPhase != "" {
						children[ph.ParentPhase] = append(children[ph.ParentPhase], ph)
					} else {
						roots = append(roots, ph)
					}
				}
				// Load tasks for status lookup
				taskMap := map[string]*model.TaskEntity{}
				allTasks, _ := store.List(model.KindTask)
				for _, ewb := range allTasks {
					t := ewb.Entity.(*model.TaskEntity)
					taskMap[utils.Slugify(t.Name)] = t
				}

				// Render tree
				var renderPhase func(ph model.PlanPhase, indent string)
				renderPhase = func(ph model.PlanPhase, indent string) {
					icon := "○"
					switch ph.Status {
					case model.PhaseCompleted:
						icon = "✓"
					case model.PhaseInProgress:
						icon = "●"
					case model.PhaseSkipped:
						icon = "–"
					}
					name := ph.Name
					if name == "" {
						name = ph.Description
					}
					// Show aggregated counts for parent phases
					childTasks := p.CollectTasks(ph.ID)
					suffix := ""
					if len(childTasks) > 0 {
						suffix = fmt.Sprintf(" [%d tasks]", len(childTasks))
					}
					fmt.Printf("%s%s %s — %s%s\n", indent, icon, name, ph.Status, suffix)
					detailIndent := indent + "    "
					if ph.Name != "" && ph.Description != "" {
						fmt.Printf("%s%s\n", detailIndent, ph.Description)
					}
					if ph.Objective != "" {
						fmt.Printf("%sObjective: %s\n", detailIndent, ph.Objective)
					}
					if ph.Changes != "" {
						fmt.Printf("%sChanges:   %s\n", detailIndent, ph.Changes)
					}
					if planShowFull {
						if ph.Details != "" {
							fmt.Printf("%sDetails:   %s\n", detailIndent, ph.Details)
						}
						if ph.Notes != "" {
							fmt.Printf("%sNotes:     %s\n", detailIndent, ph.Notes)
						}
					}
					// Show tasks as third level (always, not just in --full)
					for _, taskSlug := range ph.Tasks {
						taskIcon := "○"
						taskStatus := "pending"
						var tt *model.TaskEntity
						if t, ok := taskMap[taskSlug]; ok {
							tt = t
							taskStatus = string(t.TaskStatus)
							switch t.TaskStatus {
							case model.TaskCompleted:
								taskIcon = "✓"
							case model.TaskInProgress:
								taskIcon = "●"
							case model.TaskBlocked:
								taskIcon = "✗"
							case model.TaskCancelled:
								taskIcon = "–"
							}
						}
						fmt.Printf("%s  %s %s — %s\n", indent, taskIcon, taskSlug, taskStatus)
						if tt != nil {
							taskIndent := indent + "      "
							if tt.Objective != "" {
								fmt.Printf("%sObjective:  %s\n", taskIndent, tt.Objective)
							}
							if planShowFull {
								if tt.Details != "" {
									fmt.Printf("%sDetails:    %s\n", taskIndent, tt.Details)
								}
								if tt.Acceptance != "" {
									fmt.Printf("%sAcceptance: %s\n", taskIndent, tt.Acceptance)
								}
							}
						}
					}
					for _, child := range children[ph.ID] {
						renderPhase(child, indent+"  ")
					}
				}
				for _, root := range roots {
					renderPhase(root, "  ")
				}
			}
		}

		if body != "" {
			fmt.Printf("\n%s\n", body)
		}

		return nil
	},
}

var (
	planUpdateBackground string
	planUpdateObjective  string
	planUpdateScope      string
	planUpdateDesc       string
	planUpdatePurpose    string
	planUpdateDesign     string
)

var planUpdateCmd = &cobra.Command{
	Use:   "update <slug>",
	Short: "Update plan-level fields (background, objective, scope)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if cmd.Flags().Changed("background") {
			p.Background = planUpdateBackground
		}
		if cmd.Flags().Changed("objective") {
			p.Objective = planUpdateObjective
		}
		if cmd.Flags().Changed("scope") {
			p.PlanScope = planUpdateScope
		}
		if cmd.Flags().Changed("description") {
			p.Description = planUpdateDesc
		}
		if cmd.Flags().Changed("purpose") {
			p.Purpose = planUpdatePurpose
		}
		if cmd.Flags().Changed("design") {
			p.Design = planUpdateDesign
		}

		if _, err := store.Update(p, body); err != nil {
			return err
		}
		fmt.Printf("Updated plan: %s\n", p.Name)
		return nil
	},
}

var planApproveCmd = &cobra.Command{
	Use:   "approve <slug>",
	Short: "Approve a plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, slug)
		if err != nil {
			return fmt.Errorf("plan not found: %s", slug)
		}

		p := entity.(*model.PlanEntity)
		if err := validatePlanApprovalReadiness(p); err != nil {
			return err
		}

		p.PlanStatus = model.PlanApproved
		if p.ApprovedAt == "" {
			p.ApprovedAt = time.Now().UTC().Format(time.RFC3339)
		}

		req, created, reqFile, err := createRequirementIfMissing(store, requirementCapture{
			Name:       requirementName("Approved plan", p.Name),
			Statement:  planRequirementStatement(p),
			Source:     "plan",
			SourceRef:  "plan:" + p.CanonicalSlug(),
			Rationale:  "Captured automatically when the plan was approved.",
			ApprovedAt: p.ApprovedAt,
			Relationships: []model.Relationship{
				{
					Target: p.CanonicalSlug(),
					Type:   model.RelReferences,
					Label:  "approved_plan",
				},
			},
		})
		if err != nil {
			return fmt.Errorf("create plan requirement: %w", err)
		}
		p.Relationships = appendRelationshipOnce(p.Relationships, model.Relationship{
			Target: req.CanonicalSlug(),
			Type:   model.RelReferences,
			Label:  "requirement",
		})

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		fmt.Printf("Approved plan: %s\n", p.Name)
		if created {
			fmt.Printf("  Requirement: %s\n", reqFile)
		} else {
			fmt.Printf("  Requirement already exists: %s\n", reqFile)
		}
		return nil
	},
}

func validatePlanApprovalReadiness(p *model.PlanEntity) error {
	if len(p.Phases) == 0 {
		return fmt.Errorf("cannot approve plan %q: add at least one phase with tasks first", p.Name)
	}

	var emptyPhases []string
	for _, ph := range p.Phases {
		if len(ph.Tasks) > 0 {
			continue
		}
		label := ph.ID
		if strings.TrimSpace(ph.Name) != "" {
			label = fmt.Sprintf("%s (%s)", ph.ID, ph.Name)
		}
		if strings.TrimSpace(label) == "" {
			label = "(unnamed phase)"
		}
		emptyPhases = append(emptyPhases, label)
	}
	if len(emptyPhases) > 0 {
		return fmt.Errorf(
			"cannot approve plan %q: every phase must have at least one direct task before approval; empty phases: %s",
			p.Name,
			strings.Join(emptyPhases, ", "),
		)
	}
	return nil
}

func planRequirementStatement(p *model.PlanEntity) string {
	if strings.TrimSpace(p.Objective) != "" {
		return p.Objective
	}
	if strings.TrimSpace(p.Description) != "" {
		return p.Description
	}
	return "Approve and execute plan: " + p.Name
}

var (
	planPhaseStatus    string
	planPhaseName      string
	planPhaseParent    string
	planPhaseDesc      string
	planPhaseObjective string
	planPhaseChanges   string
	planPhaseDetails   string
	planPhaseNotes     string
)

var planPhaseCmd = &cobra.Command{
	Use:   "phase <plan-slug> <phase-id>",
	Short: "Update a plan phase",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug := args[0]
		phaseID := args[1]

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, planSlug)
		if err != nil {
			return fmt.Errorf("plan not found: %s", planSlug)
		}

		p := entity.(*model.PlanEntity)
		found := false
		for i := range p.Phases {
			if p.Phases[i].ID == phaseID {
				if cmd.Flags().Changed("name") {
					p.Phases[i].Name = planPhaseName
				}
				if cmd.Flags().Changed("parent") {
					p.Phases[i].ParentPhase = planPhaseParent
				}
				if cmd.Flags().Changed("status") {
					newStatus := model.PhaseStatus(planPhaseStatus)
					// Prevent completing a phase if children are not all completed
					if newStatus == model.PhaseCompleted {
						children := p.ChildPhases(phaseID)
						for _, child := range children {
							if child.Status != model.PhaseCompleted && child.Status != model.PhaseSkipped {
								return fmt.Errorf("cannot complete phase %s: child phase %s (%s) is still %s", phaseID, child.ID, child.Name, child.Status)
							}
						}
					}
					p.Phases[i].Status = newStatus
				}
				if cmd.Flags().Changed("description") {
					p.Phases[i].Description = planPhaseDesc
				}
				if cmd.Flags().Changed("objective") {
					p.Phases[i].Objective = planPhaseObjective
				}
				if cmd.Flags().Changed("changes") {
					p.Phases[i].Changes = planPhaseChanges
				}
				if cmd.Flags().Changed("details") {
					p.Phases[i].Details = planPhaseDetails
				}
				if cmd.Flags().Changed("notes") {
					p.Phases[i].Notes = planPhaseNotes
				}
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("phase not found: %s", phaseID)
		}

		// Auto-complete plan if all phases done
		allDone := true
		for _, ph := range p.Phases {
			if ph.Status != model.PhaseCompleted && ph.Status != model.PhaseSkipped {
				allDone = false
				break
			}
		}
		if allDone {
			p.PlanStatus = model.PlanCompleted
			p.CompletedAt = time.Now().UTC().Format(time.RFC3339)
		}

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		fmt.Printf("Updated phase %s → %s\n", phaseID, planPhaseStatus)
		if allDone {
			fmt.Printf("Plan '%s' is now completed!\n", p.Name)
		}
		return nil
	},
}

var planCompleteForce bool

var planCompleteCmd = &cobra.Command{
	Use:   "complete <plan-slug>",
	Short: "Mark a plan completed, blocking on plan-completion audit errors",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, body, err := loadPlanForChange(client, slug)
		if err != nil {
			return err
		}

		// Run full sync check — plan complete requires a clean model.
		report, err := client.c.SyncCheck(true)
		if err != nil {
			return fmt.Errorf("fetch audit report: %w", err)
		}

		// Show plan-specific completion findings.
		planSlug := plan.GetBase().CanonicalSlug()
		type finding = struct {
			Severity, Category, Message, EntitySlug string
		}
		var planFindings []finding
		for _, f := range report.Errors {
			if f.Category == "plan_completion" && f.EntitySlug == planSlug {
				planFindings = append(planFindings, finding{f.Severity, f.Category, f.Message, f.EntitySlug})
			}
		}
		for _, f := range report.Warnings {
			if f.Category == "plan_completion" && f.EntitySlug == planSlug {
				planFindings = append(planFindings, finding{f.Severity, f.Category, f.Message, f.EntitySlug})
			}
		}
		for _, f := range planFindings {
			icon := "!"
			if f.Severity == "error" {
				icon = "✗"
			}
			fmt.Printf("  %s [%s] %s\n", icon, f.Severity, f.Message)
		}

		// Count ALL errors across the full sync check (not just plan-scoped).
		totalErrors := len(report.Errors)
		planErrors := 0
		for _, f := range planFindings {
			if f.Severity == "error" {
				planErrors++
			}
		}

		if totalErrors > 0 && !planCompleteForce {
			if planErrors > 0 {
				fmt.Printf("\n  Plan completion findings: %d error(s)\n", planErrors)
			}
			if totalErrors > planErrors {
				fmt.Printf("  Sync check has %d additional error(s) — run 'syde sync check' for details\n", totalErrors-planErrors)
			}
			return fmt.Errorf("plan completion blocked: syde sync check reports %d total error(s) — fix all errors or re-run with --force", totalErrors)
		}

		plan.PlanStatus = model.PlanCompleted
		plan.CompletedAt = time.Now().UTC().Format(time.RFC3339)
		if _, err := client.Update(plan, body); err != nil {
			return err
		}
		fmt.Printf("Plan %s → completed\n", plan.Name)
		if totalErrors > 0 {
			fmt.Printf("  (forced despite %d audit error(s))\n", totalErrors)
		}
		return nil
	},
}

var planOpenCmd = &cobra.Command{
	Use:   "open <plan-slug>",
	Short: "Open a plan in the dashboard",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := openWriteClient()
		if err != nil {
			return err
		}
		defer wc.Close()

		plan, _, err := loadPlanForChange(wc, args[0])
		if err != nil {
			return err
		}

		path := fmt.Sprintf("/%s/plan/%s", wc.c.ProjectSlug(), plan.GetBase().CanonicalSlug())
		resp, err := wc.c.Navigate(path)
		if err != nil {
			return fmt.Errorf("send navigate event: %w", err)
		}

		url := wc.c.DashboardURL(path)
		if resp.Clients > 0 {
			fmt.Println("opened in existing dashboard tab")
			fmt.Println(url)
			return nil
		}

		openBrowser(url)
		fmt.Println(url)
		return nil
	},
}

var planReviewCmd = &cobra.Command{
	Use:   "review <plan-slug>",
	Short: "Print a plan-reviewer subagent prompt with the plan interpolated",
	Long: `Loads the plan markdown and the reviewer-prompt template bundled
with the skill, interpolates the plan content, and prints the rendered
prompt to stdout. Paste the output into a fresh subagent (Claude Code
sub-agent, Codex agent, claude --print session, etc.) to receive a
calibrated Approved | Issues Found verdict.

Output is the prompt only — the CLI does not invoke a reviewer itself.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := openWriteClient()
		if err != nil {
			return err
		}
		defer wc.Close()

		plan, _, err := loadPlanForChange(wc, args[0])
		if err != nil {
			return err
		}
		planSlug := plan.GetBase().CanonicalSlug()

		sydeDir := wc.FS.Root
		planPath := filepath.Join(sydeDir, "plans", planSlug+".md")
		planBytes, err := os.ReadFile(planPath)
		if err != nil {
			return fmt.Errorf("read plan markdown: %w", err)
		}

		tmplBytes, err := skill.FS.ReadFile("references/plan-review-prompt.md")
		if err != nil {
			return fmt.Errorf("load reviewer prompt template: %w", err)
		}
		tmpl := string(tmplBytes)
		rendered := strings.ReplaceAll(tmpl, "{{plan_slug}}", planSlug)
		rendered = strings.ReplaceAll(rendered, "{{plan_markdown}}", string(planBytes))

		fmt.Println(rendered)
		return nil
	},
}

var planCheckCmd = &cobra.Command{
	Use:   "check <plan-slug>",
	Short: "Check plan authoring and completion findings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		wc, err := openWriteClient()
		if err != nil {
			return err
		}
		defer wc.Close()

		plan, _, err := loadPlanForChange(wc, args[0])
		if err != nil {
			return err
		}

		findings, err := runPlanCheck(wc, plan)
		if err != nil {
			return err
		}

		errorCount := printPlanCheckFindings(findings)
		if errorCount > 0 {
			return fmt.Errorf("plan check failed: %d error(s)", errorCount)
		}
		return nil
	},
}

func runPlanCheck(wc *writeClient, plan *model.PlanEntity) ([]client.Finding, error) {
	report, err := wc.c.SyncCheck(true)
	if err != nil {
		return nil, fmt.Errorf("fetch audit report: %w", err)
	}
	planSlug := plan.GetBase().CanonicalSlug()
	var scoped []client.Finding
	for _, bucket := range [][]client.Finding{report.Errors, report.Warnings, report.Hints} {
		for _, f := range bucket {
			if f.EntitySlug != planSlug {
				continue
			}
			if f.Category != "plan_authoring" && f.Category != "plan_completion" {
				continue
			}
			scoped = append(scoped, f)
		}
	}
	return scoped, nil
}

func printPlanCheckFindings(findings []client.Finding) int {
	if len(findings) == 0 {
		fmt.Println("No plan check findings.")
		return 0
	}
	order := []string{"error", "warning", "hint"}
	errorCount := 0
	for _, severity := range order {
		var bucket []client.Finding
		for _, f := range findings {
			if f.Severity == severity {
				bucket = append(bucket, f)
			}
		}
		if len(bucket) == 0 {
			continue
		}
		fmt.Printf("%s (%d)\n", strings.ToUpper(severity), len(bucket))
		for _, f := range bucket {
			if severity == "error" {
				errorCount++
			}
			field := ""
			if f.Field != "" {
				field = " [" + f.Field + "]"
			}
			fmt.Printf("  - %s%s: %s\n", f.Category, field, f.Message)
		}
	}
	return errorCount
}

// runPlanCompletionCheck calls the syded audit endpoint and filters
// the findings down to plan_completion entries scoped to the given
// plan. The HealthReport groups by severity; we flatten both errors
// and warnings so the caller can print them in the same loop.
func runPlanCompletionCheck(wc *writeClient, plan *model.PlanEntity) ([]client.Finding, error) {
	report, err := wc.c.SyncCheck(true)
	if err != nil {
		return nil, fmt.Errorf("fetch audit report: %w", err)
	}
	planSlug := plan.GetBase().CanonicalSlug()
	var scoped []client.Finding
	for _, bucket := range [][]client.Finding{report.Errors, report.Warnings} {
		for _, f := range bucket {
			if f.Category != "plan_completion" {
				continue
			}
			if f.EntitySlug != planSlug {
				continue
			}
			scoped = append(scoped, f)
		}
	}
	return scoped, nil
}

var planSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync from most recent Claude Code plan file",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		// Find most recent .claude/plans/*.md
		dir := sydeDir
		if dir == "" {
			dir, _ = findSydeDirHelper()
		}
		projectRoot := filepath.Dir(dir)
		plansDir := filepath.Join(projectRoot, ".claude", "plans")

		entries, err := os.ReadDir(plansDir)
		if err != nil {
			return fmt.Errorf("no .claude/plans/ directory found")
		}

		var latestFile string
		var latestTime time.Time
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
				continue
			}
			info, _ := e.Info()
			if info.ModTime().After(latestTime) {
				latestTime = info.ModTime()
				latestFile = filepath.Join(plansDir, e.Name())
			}
		}

		if latestFile == "" {
			return fmt.Errorf("no plan files found in .claude/plans/")
		}

		data, err := os.ReadFile(latestFile)
		if err != nil {
			return err
		}

		// Create a syde plan from the Claude plan
		planName := strings.TrimSuffix(filepath.Base(latestFile), ".md")
		planName = strings.ReplaceAll(planName, "-", " ")

		plan := &model.PlanEntity{
			BaseEntity: model.BaseEntity{
				Kind: model.KindPlan,
				Name: planName,
			},
			PlanStatus:     model.PlanDraft,
			Source:         "claude-plan",
			ClaudePlanFile: latestFile,
			CreatedAt:      time.Now().UTC().Format(time.RFC3339),
		}

		filePath, err := store.Create(plan, string(data))
		if err != nil {
			return err
		}

		fmt.Printf("Synced plan from: %s\n", filepath.Base(latestFile))
		fmt.Printf("  Created: %s\n", filePath)
		fmt.Println("  Edit the plan file to add structured phases (phases: field in frontmatter)")
		return nil
	},
}

var planExecuteCmd = &cobra.Command{
	Use:   "execute <slug>",
	Short: "Scaffold entity files for pending plan phases",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if p.PlanStatus != model.PlanApproved && p.PlanStatus != model.PlanInProgress {
			return fmt.Errorf("plan must be approved first (current status: %s)", p.PlanStatus)
		}

		p.PlanStatus = model.PlanInProgress
		store.Update(p, body)
		fmt.Printf("Plan '%s' is now in-progress.\n", p.Name)
		fmt.Println()
		fmt.Println("Plans no longer materialize draft entities on execute.")
		fmt.Println("Create entities with `syde add` as part of each task's implementation.")
		fmt.Println("If a task needs new entities, list them in its --note field, then run:")
		fmt.Println("    syde task start <slug>")
		return nil
	},
}

var (
	addPhaseName      string
	addPhaseParent    string
	addPhaseDesc      string
	addPhaseObjective string
	addPhaseChanges   string
	addPhaseDetails   string
	addPhaseNotes     string
)

var planAddPhaseCmd = &cobra.Command{
	Use:   "add-phase <plan-slug>",
	Short: "Add a structured phase to a plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if strings.TrimSpace(addPhaseName) == "" {
			return fmt.Errorf("phase --name is required (short label identifying this phase)")
		}

		// Compute a unique phase_N id that does not collide with any
		// existing phase, even if the on-disk data is malformed (e.g.
		// gaps from a prior buggy writer or hand edits). Starting from
		// len(Phases)+1 and walking forward guarantees a non-empty
		// non-colliding ID, defending the write path against producing
		// another corrupt plan.
		existingIDs := make(map[string]bool, len(p.Phases))
		for _, ph := range p.Phases {
			if ph.ID != "" {
				existingIDs[ph.ID] = true
			}
		}
		phaseID := ""
		for n := len(p.Phases) + 1; n < len(p.Phases)+100; n++ {
			candidate := fmt.Sprintf("phase_%d", n)
			if !existingIDs[candidate] {
				phaseID = candidate
				break
			}
		}
		if phaseID == "" {
			return fmt.Errorf("could not allocate a unique phase id for plan %q", p.Name)
		}
		phase := model.PlanPhase{
			ID:          phaseID,
			Name:        addPhaseName,
			ParentPhase: addPhaseParent,
			Status:      model.PhasePending,
			Description: addPhaseDesc,
			Objective:   addPhaseObjective,
			Changes:     addPhaseChanges,
			Details:     addPhaseDetails,
			Notes:       addPhaseNotes,
		}
		p.Phases = append(p.Phases, phase)

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		displayName := addPhaseName
		if displayName == "" {
			displayName = addPhaseDesc
		}
		fmt.Printf("Added phase %s to plan '%s'\n", phaseID, p.Name)
		if addPhaseParent != "" {
			fmt.Printf("  %s (parent: %s): %s\n", phaseID, addPhaseParent, displayName)
		} else {
			fmt.Printf("  %s: %s\n", phaseID, displayName)
		}
		fmt.Printf("  Total phases: %d\n", len(p.Phases))
		return nil
	},
}

var planEstimateCmd = &cobra.Command{
	Use:   "estimate <slug>",
	Short: "Estimate plan size and suggest splitting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, _, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if len(p.Phases) == 0 {
			fmt.Printf("Plan '%s' has no structured phases.\n", p.Name)
			fmt.Println("Add phases with: syde plan add-phase")
			return nil
		}

		// Count tasks across all phases for a rough estimate
		totalTasks := 0
		for _, ph := range p.Phases {
			totalTasks += len(ph.Tasks)
		}
		cmdPerTask := 5 // start, query, edit files, verify, done
		totalEstimated := totalTasks * cmdPerTask

		fmt.Printf("Plan: %s (%d phases, %d tasks)\n", p.Name, len(p.Phases), totalTasks)
		fmt.Printf("\nEstimated commands: ~%d (at ~%d per task)\n", totalEstimated, cmdPerTask)

		if totalEstimated > 25 {
			subPlanSize := 5
			numSubPlans := (len(p.Phases) + subPlanSize - 1) / subPlanSize
			fmt.Printf("\n⚠ Recommendation: SPLIT into %d sub-plans of ~%d phases each\n", numSubPlans, subPlanSize)
			for i := 0; i < numSubPlans; i++ {
				start := i * subPlanSize
				end := start + subPlanSize
				if end > len(p.Phases) {
					end = len(p.Phases)
				}
				fmt.Printf("  Sub-plan %d: phases %d-%d\n", i+1, start+1, end)
			}
		} else {
			fmt.Println("\n✓ Plan fits in one session turn.")
		}

		return nil
	},
}

var (
	planChangeWhy          string
	planChangeWhat         string
	planChangeName         string
	planChangeFields       []string
	planChangeTasks        []string
	planChangeDraftFlags   = map[string]string{}
	planChangeDraftListFlg = map[string][]string{}
)

// newChangeID returns a 4-char lowercase alphanumeric id for a
// plan-change entry. Unique within a single plan is enough; the
// caller can address the entry via remove-change <plan> <id>.
func newChangeID() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 4)
	r := make([]byte, 4)
	_, _ = crand.Read(r)
	for i := range b {
		b[i] = alphabet[int(r[i])%len(alphabet)]
	}
	return string(b)
}

// planChangeLane returns a pointer to the named lane so handlers can
// mutate it in place. Kind must be one of the six coverage kinds
// (requirement/system/concept/component/contract/flow).
func planChangeLane(changes *model.PlanChanges, kind string) (*model.ChangeLane, error) {
	switch strings.ToLower(kind) {
	case "requirement", "requirements":
		return &changes.Requirements, nil
	case "system", "systems":
		return &changes.Systems, nil
	case "concept", "concepts":
		return &changes.Concepts, nil
	case "component", "components":
		return &changes.Components, nil
	case "contract", "contracts":
		return &changes.Contracts, nil
	case "flow", "flows":
		return &changes.Flows, nil
	}
	return nil, fmt.Errorf("unknown kind %q (expected requirement|system|concept|component|contract|flow)", kind)
}

// loadPlanForChange resolves a plan by slug and returns the entity +
// body so the caller can append a change and store.Update it back.
func loadPlanForChange(client *writeClient, slug string) (*model.PlanEntity, string, error) {
	entity, body, err := client.Get(slug)
	if err != nil {
		return nil, "", err
	}
	plan, ok := entity.(*model.PlanEntity)
	if !ok {
		return nil, "", fmt.Errorf("%s is not a plan", slug)
	}
	return plan, body, nil
}

var planAddChangeDeleteCmd = &cobra.Command{
	Use:   "delete <plan-slug> <kind> <target-slug>",
	Short: "Append a Deleted entry to a plan's changes lane",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug, kind, targetSlug := args[0], args[1], args[2]
		if strings.TrimSpace(planChangeWhy) == "" {
			return fmt.Errorf("--why is required")
		}

		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, body, err := loadPlanForChange(client, planSlug)
		if err != nil {
			return err
		}
		lane, err := planChangeLane(&plan.Changes, kind)
		if err != nil {
			return err
		}
		lane.Deleted = append(lane.Deleted, model.DeletedChange{
			ID:    newChangeID(),
			Slug:  targetSlug,
			Why:   planChangeWhy,
			Tasks: planChangeTasks,
		})

		if _, err := client.Update(plan, body); err != nil {
			return fmt.Errorf("update plan: %w", err)
		}
		fmt.Printf("Added delete-change to %s (%s lane): %s\n", planSlug, kind, targetSlug)
		return nil
	},
}

var planAddChangeExtendCmd = &cobra.Command{
	Use:   "extend <plan-slug> <kind> <target-slug>",
	Short: "Append an Extended entry to a plan's changes lane",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug, kind, targetSlug := args[0], args[1], args[2]
		if strings.TrimSpace(planChangeWhat) == "" || strings.TrimSpace(planChangeWhy) == "" {
			return fmt.Errorf("--what and --why are required")
		}
		fieldChanges, err := parseFieldChanges(planChangeFields)
		if err != nil {
			return err
		}

		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, body, err := loadPlanForChange(client, planSlug)
		if err != nil {
			return err
		}
		lane, err := planChangeLane(&plan.Changes, kind)
		if err != nil {
			return err
		}
		lane.Extended = append(lane.Extended, model.ExtendedChange{
			ID:           newChangeID(),
			Slug:         targetSlug,
			What:         planChangeWhat,
			Why:          planChangeWhy,
			FieldChanges: fieldChanges,
			Tasks:        planChangeTasks,
		})

		if _, err := client.Update(plan, body); err != nil {
			return fmt.Errorf("update plan: %w", err)
		}
		fmt.Printf("Added extend-change to %s (%s lane): %s (%d field changes)\n", planSlug, kind, targetSlug, len(fieldChanges))
		return nil
	},
}

// parseFieldChanges turns repeatable --field key=value flags into a
// map. Keys must be non-empty; "DELETE" is the sentinel value that
// tells the completion validator the target field should be cleared.
func parseFieldChanges(flags []string) (map[string]string, error) {
	if len(flags) == 0 {
		return nil, nil
	}
	out := make(map[string]string, len(flags))
	for _, f := range flags {
		idx := strings.Index(f, "=")
		if idx <= 0 {
			return nil, fmt.Errorf("invalid --field %q: expected key=value", f)
		}
		k := strings.TrimSpace(f[:idx])
		v := f[idx+1:]
		if k == "" {
			return nil, fmt.Errorf("invalid --field %q: empty key", f)
		}
		out[k] = v
	}
	return out, nil
}

var planAddChangeNewCmd = &cobra.Command{
	Use:   "new <plan-slug> <kind>",
	Short: "Append a NewChange draft to a plan's changes lane",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug, kind := args[0], args[1]
		if strings.TrimSpace(planChangeName) == "" || strings.TrimSpace(planChangeWhat) == "" || strings.TrimSpace(planChangeWhy) == "" {
			return fmt.Errorf("--name, --what, and --why are required")
		}
		draft, err := parseDraftFields(planChangeFields)
		if err != nil {
			return err
		}
		if len(draft) == 0 {
			return fmt.Errorf("at least one --draft key=value is required — NewChange must carry kind-specific fields")
		}

		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, body, err := loadPlanForChange(client, planSlug)
		if err != nil {
			return err
		}
		lane, err := planChangeLane(&plan.Changes, kind)
		if err != nil {
			return err
		}
		lane.New = append(lane.New, model.NewChange{
			ID:    newChangeID(),
			Name:  planChangeName,
			What:  planChangeWhat,
			Why:   planChangeWhy,
			Draft: draft,
			Tasks: planChangeTasks,
		})

		if _, err := client.Update(plan, body); err != nil {
			return fmt.Errorf("update plan: %w", err)
		}
		fmt.Printf("Added new-change to %s (%s lane): %s (%d draft fields)\n", planSlug, kind, planChangeName, len(draft))
		return nil
	},
}

// parseDraftFields parses repeatable --draft key=value pairs. Values
// that parse as JSON arrays/objects/numbers/booleans are decoded into
// their native Go types so per-kind list fields (capabilities, tags,
// attributes) round-trip cleanly. Everything else stays a string.
func parseDraftFields(flags []string) (map[string]interface{}, error) {
	if len(flags) == 0 {
		return nil, nil
	}
	out := make(map[string]interface{}, len(flags))
	for _, f := range flags {
		idx := strings.Index(f, "=")
		if idx <= 0 {
			return nil, fmt.Errorf("invalid --draft %q: expected key=value", f)
		}
		k := strings.TrimSpace(f[:idx])
		v := f[idx+1:]
		if k == "" {
			return nil, fmt.Errorf("invalid --draft %q: empty key", f)
		}
		// Try JSON-decode first so arrays/objects/numbers/bools
		// round-trip into native Go types. Fall back to raw string.
		trimmed := strings.TrimSpace(v)
		if trimmed != "" && (trimmed[0] == '[' || trimmed[0] == '{' || trimmed[0] == '"' || trimmed == "true" || trimmed == "false") {
			var decoded interface{}
			if err := json.Unmarshal([]byte(trimmed), &decoded); err == nil {
				out[k] = decoded
				continue
			}
		}
		out[k] = v
	}
	return out, nil
}

var planAddChangeCmd = &cobra.Command{
	Use:   "add-change",
	Short: "Append a structured change entry to a plan",
}

// removeChangeFromLane drops any entry whose ID matches id from the
// three sub-slices and reports whether it removed anything.
func removeChangeFromLane(lane *model.ChangeLane, id string) bool {
	removed := false
	var keptDeleted []model.DeletedChange
	for _, d := range lane.Deleted {
		if d.ID == id {
			removed = true
			continue
		}
		keptDeleted = append(keptDeleted, d)
	}
	var keptExtended []model.ExtendedChange
	for _, e := range lane.Extended {
		if e.ID == id {
			removed = true
			continue
		}
		keptExtended = append(keptExtended, e)
	}
	var keptNew []model.NewChange
	for _, n := range lane.New {
		if n.ID == id {
			removed = true
			continue
		}
		keptNew = append(keptNew, n)
	}
	if removed {
		lane.Deleted = keptDeleted
		lane.Extended = keptExtended
		lane.New = keptNew
	}
	return removed
}

var planRemoveChangeCmd = &cobra.Command{
	Use:   "remove-change <plan-slug> <change-id>",
	Short: "Remove a change entry from a plan by id",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug, changeID := args[0], args[1]

		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, body, err := loadPlanForChange(client, planSlug)
		if err != nil {
			return err
		}
		lanes := []*model.ChangeLane{
			&plan.Changes.Requirements,
			&plan.Changes.Systems,
			&plan.Changes.Concepts,
			&plan.Changes.Components,
			&plan.Changes.Contracts,
			&plan.Changes.Flows,
		}
		found := false
		for _, lane := range lanes {
			if removeChangeFromLane(lane, changeID) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("no change with id %q in plan %q", changeID, planSlug)
		}
		if _, err := client.Update(plan, body); err != nil {
			return fmt.Errorf("update plan: %w", err)
		}
		fmt.Printf("Removed change %s from plan %s\n", changeID, planSlug)
		return nil
	},
}

var planShowChangesFormat string

var planShowChangesCmd = &cobra.Command{
	Use:   "show-changes <plan-slug>",
	Short: "Print a plan's structured change diff",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := openWriteClient()
		if err != nil {
			return err
		}
		defer client.Close()

		plan, _, err := loadPlanForChange(client, args[0])
		if err != nil {
			return err
		}

		if planShowChangesFormat == "json" {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(plan.Changes)
		}

		lanes := map[string]model.ChangeLane{
			"requirements": plan.Changes.Requirements,
			"systems":      plan.Changes.Systems,
			"concepts":     plan.Changes.Concepts,
			"components":   plan.Changes.Components,
			"contracts":    plan.Changes.Contracts,
			"flows":        plan.Changes.Flows,
		}
		laneOrder := []string{"requirements", "systems", "concepts", "components", "contracts", "flows"}
		empty := true
		for _, name := range laneOrder {
			lane := lanes[name]
			if len(lane.Deleted) == 0 && len(lane.Extended) == 0 && len(lane.New) == 0 {
				continue
			}
			empty = false
			fmt.Printf("── %s ── (del:%d ext:%d new:%d)\n", name, len(lane.Deleted), len(lane.Extended), len(lane.New))
			for _, d := range lane.Deleted {
				taskSuffix := formatChangeTasks(d.Tasks)
				fmt.Printf("  ✗ [%s] delete %s — %s%s\n", d.ID, d.Slug, d.Why, taskSuffix)
			}
			for _, e := range lane.Extended {
				fields := ""
				if len(e.FieldChanges) > 0 {
					var kv []string
					for k, v := range e.FieldChanges {
						kv = append(kv, fmt.Sprintf("%s=%q", k, v))
					}
					fields = " {" + strings.Join(kv, ", ") + "}"
				}
				taskSuffix := formatChangeTasks(e.Tasks)
				fmt.Printf("  ± [%s] extend %s — %s%s%s\n     why: %s\n", e.ID, e.Slug, e.What, fields, taskSuffix, e.Why)
			}
			for _, n := range lane.New {
				taskSuffix := formatChangeTasks(n.Tasks)
				fmt.Printf("  + [%s] new %q — %s%s\n     why: %s\n     draft: %v\n", n.ID, n.Name, n.What, taskSuffix, n.Why, n.Draft)
			}
			fmt.Println()
		}
		if empty {
			fmt.Println("(no changes)")
		}
		return nil
	},
}

func formatChangeTasks(tasks []string) string {
	if len(tasks) == 0 {
		return ""
	}
	return fmt.Sprintf(" [tasks: %s]", strings.Join(tasks, ","))
}

var planSyncFile string

func init() {
	_ = utils.GenerateID
	planShowCmd.Flags().BoolVar(&planShowFull, "full", false, "show entity draft data, phase details, and notes")

	planCreateCmd.Flags().StringVar(&planCreateBackground, "background", "", "why this plan exists (context, motivation)")
	planCreateCmd.Flags().StringVar(&planCreateObjective, "objective", "", "what success looks like")
	planCreateCmd.Flags().StringVar(&planCreateScope, "scope", "", "what's in-scope / out-of-scope")
	planCreateCmd.Flags().StringVar(&planCreateDesign, "design", "", "detailed implementation design prose (rendered as markdown in the dashboard)")

	planUpdateCmd.Flags().StringVar(&planUpdateBackground, "background", "", "update background")
	planUpdateCmd.Flags().StringVar(&planUpdateObjective, "objective", "", "update objective")
	planUpdateCmd.Flags().StringVar(&planUpdateScope, "scope", "", "update scope")
	planUpdateCmd.Flags().StringVar(&planUpdateDesc, "description", "", "update description")
	planUpdateCmd.Flags().StringVar(&planUpdatePurpose, "purpose", "", "update purpose")
	planUpdateCmd.Flags().StringVar(&planUpdateDesign, "design", "", "update implementation design prose")

	planPhaseCmd.Flags().StringVar(&planPhaseName, "name", "", "update phase name")
	planPhaseCmd.Flags().StringVar(&planPhaseParent, "parent", "", "set parent phase ID for nesting")
	planPhaseCmd.Flags().StringVar(&planPhaseStatus, "status", "", "phase status (pending, in_progress, completed, skipped)")
	planPhaseCmd.Flags().StringVar(&planPhaseDesc, "description", "", "update phase description")
	planPhaseCmd.Flags().StringVar(&planPhaseObjective, "objective", "", "update phase objective (what it achieves)")
	planPhaseCmd.Flags().StringVar(&planPhaseChanges, "changes", "", "update phase changes (what concretely changes)")
	planPhaseCmd.Flags().StringVar(&planPhaseDetails, "details", "", "update phase details (implementation walkthrough)")
	planPhaseCmd.Flags().StringVar(&planPhaseNotes, "notes", "", "update phase notes (risks, reminders)")

	planAddPhaseCmd.Flags().StringVar(&addPhaseName, "name", "", "phase name (short label)")
	planAddPhaseCmd.Flags().StringVar(&addPhaseParent, "parent", "", "parent phase ID for nesting (e.g., phase_1)")
	planAddPhaseCmd.Flags().StringVar(&addPhaseDesc, "description", "", "phase description")
	planAddPhaseCmd.Flags().StringVar(&addPhaseObjective, "objective", "", "what this phase achieves")
	planAddPhaseCmd.Flags().StringVar(&addPhaseChanges, "changes", "", "what concretely changes in this phase")
	planAddPhaseCmd.Flags().StringVar(&addPhaseDetails, "details", "", "implementation walkthrough (how to build)")
	planAddPhaseCmd.Flags().StringVar(&addPhaseNotes, "notes", "", "additional notes, risks, reminders. Mention new entities the phase will require here as free text — the task's --note flag takes it from there.")

	planSyncCmd.Flags().StringVar(&planSyncFile, "file", "", "path to plan file (instead of scanning .claude/plans/)")

	planAddChangeDeleteCmd.Flags().StringVar(&planChangeWhy, "why", "", "why the entity will be deleted (required)")
	planAddChangeDeleteCmd.Flags().StringArrayVar(&planChangeTasks, "task", nil, "task slug implementing this change (repeatable)")

	planAddChangeExtendCmd.Flags().StringVar(&planChangeWhat, "what", "", "what concretely changes (required)")
	planAddChangeExtendCmd.Flags().StringVar(&planChangeWhy, "why", "", "why the change is being made (required)")
	planAddChangeExtendCmd.Flags().StringArrayVar(&planChangeFields, "field", nil, "declared field change in key=value form (repeatable; value 'DELETE' clears the field)")
	planAddChangeExtendCmd.Flags().StringArrayVar(&planChangeTasks, "task", nil, "task slug implementing this change (repeatable)")

	planAddChangeNewCmd.Flags().StringVar(&planChangeName, "name", "", "human-readable name for the new entity (required)")
	planAddChangeNewCmd.Flags().StringVar(&planChangeWhat, "what", "", "one-line description of what the new entity does (required)")
	planAddChangeNewCmd.Flags().StringVar(&planChangeWhy, "why", "", "why it needs to exist (required)")
	planAddChangeNewCmd.Flags().StringArrayVar(&planChangeFields, "draft", nil, "kind-specific draft field in key=value form (repeatable; at least one required)")
	planAddChangeNewCmd.Flags().StringArrayVar(&planChangeTasks, "task", nil, "task slug implementing this change (repeatable)")

	planAddChangeCmd.AddCommand(planAddChangeDeleteCmd, planAddChangeExtendCmd, planAddChangeNewCmd)

	planShowChangesCmd.Flags().StringVar(&planShowChangesFormat, "format", "rich", "output format (rich|json)")

	planCompleteCmd.Flags().BoolVar(&planCompleteForce, "force", false, "complete the plan even when plan_completion audit reports errors")

	planCmd.AddCommand(planReviewCmd)
	planCmd.AddCommand(planCreateCmd, planListCmd, planShowCmd, planApproveCmd, planUpdateCmd,
		planPhaseCmd, planSyncCmd, planExecuteCmd, planAddPhaseCmd, planEstimateCmd,
		planAddChangeCmd, planRemoveChangeCmd, planShowChangesCmd, planCompleteCmd, planOpenCmd, planCheckCmd)
	rootCmd.AddCommand(planCmd)
}
