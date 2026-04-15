package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
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

var planSyncFile string

func init() {
	_ = utils.GenerateID
	planShowCmd.Flags().BoolVar(&planShowFull, "full", false, "show entity draft data, phase details, and notes")

	planCreateCmd.Flags().StringVar(&planCreateBackground, "background", "", "why this plan exists (context, motivation)")
	planCreateCmd.Flags().StringVar(&planCreateObjective, "objective", "", "what success looks like")
	planCreateCmd.Flags().StringVar(&planCreateScope, "scope", "", "what's in-scope / out-of-scope")

	planUpdateCmd.Flags().StringVar(&planUpdateBackground, "background", "", "update background")
	planUpdateCmd.Flags().StringVar(&planUpdateObjective, "objective", "", "update objective")
	planUpdateCmd.Flags().StringVar(&planUpdateScope, "scope", "", "update scope")
	planUpdateCmd.Flags().StringVar(&planUpdateDesc, "description", "", "update description")
	planUpdateCmd.Flags().StringVar(&planUpdatePurpose, "purpose", "", "update purpose")

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

	planCmd.AddCommand(planCreateCmd, planListCmd, planShowCmd, planApproveCmd, planUpdateCmd,
		planPhaseCmd, planSyncCmd, planExecuteCmd, planAddPhaseCmd, planEstimateCmd)
	rootCmd.AddCommand(planCmd)
}
