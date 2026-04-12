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

var planCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		plan := &model.PlanEntity{
			BaseEntity: model.BaseEntity{
				Kind:   model.KindPlan,
				Name:   name,
				Status: model.StatusDraft,
			},
			Source:    "manual",
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
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
		store, err := openStore()
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
			completedSteps := 0
			for _, s := range p.Steps {
				if s.Status == model.StepCompleted {
					completedSteps++
				}
			}
			fmt.Printf("  %-30s %-12s %d/%d steps (%.0f%%)\n",
				p.Name, p.Status, completedSteps, len(p.Steps), progress)
		}
		return nil
	},
}

var planShowCmd = &cobra.Command{
	Use:   "show <slug>",
	Short: "Show plan details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openStore()
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
		fmt.Printf("Status: %s | Progress: %.0f%%\n", p.Status, p.Progress())
		if p.CreatedAt != "" {
			fmt.Printf("Created: %s\n", p.CreatedAt)
		}
		if p.ApprovedAt != "" {
			fmt.Printf("Approved: %s\n", p.ApprovedAt)
		}
		fmt.Println()

		if len(p.Steps) > 0 {
			fmt.Println("Steps:")
			for _, s := range p.Steps {
				icon := "○"
				switch s.Status {
				case model.StepCompleted:
					icon = "✓"
				case model.StepInProgress:
					icon = "●"
				case model.StepSkipped:
					icon = "–"
				}
				fmt.Printf("  %s %s (%s %s) — %s\n", icon, s.Description, s.Action, s.EntityKind, s.Status)
			}
		}

		if body != "" {
			fmt.Printf("\n%s\n", body)
		}

		return nil
	},
}

var planApproveCmd = &cobra.Command{
	Use:   "approve <slug>",
	Short: "Approve a plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, slug)
		if err != nil {
			return fmt.Errorf("plan not found: %s", slug)
		}

		p := entity.(*model.PlanEntity)
		p.Status = model.Status("approved")
		p.ApprovedAt = time.Now().UTC().Format(time.RFC3339)

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		fmt.Printf("Approved plan: %s\n", p.Name)
		return nil
	},
}

var planStepStatus string

var planStepCmd = &cobra.Command{
	Use:   "step <plan-slug> <step-id>",
	Short: "Update a plan step status",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		planSlug := args[0]
		stepID := args[1]

		store, err := openStore()
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
		for i := range p.Steps {
			if p.Steps[i].ID == stepID {
				p.Steps[i].Status = model.PlanStepStatus(planStepStatus)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("step not found: %s", stepID)
		}

		// Auto-complete plan if all steps done
		allDone := true
		for _, s := range p.Steps {
			if s.Status != model.StepCompleted && s.Status != model.StepSkipped {
				allDone = false
				break
			}
		}
		if allDone {
			p.Status = model.Status("completed")
			p.CompletedAt = time.Now().UTC().Format(time.RFC3339)
		}

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		fmt.Printf("Updated step %s → %s\n", stepID, planStepStatus)
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
		store, err := openStore()
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
				Kind:   model.KindPlan,
				Name:   planName,
				Status: model.StatusDraft,
			},
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
		fmt.Println("  Edit the plan file to add structured steps (steps: field in frontmatter)")
		return nil
	},
}

var planExecuteCmd = &cobra.Command{
	Use:   "execute <slug>",
	Short: "Scaffold entity files for pending plan steps",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if string(p.Status) != "approved" && string(p.Status) != "in-progress" {
			return fmt.Errorf("plan must be approved first (current status: %s)", p.Status)
		}

		p.Status = model.Status("in-progress")
		created := 0

		for i, step := range p.Steps {
			if step.Status != model.StepPending {
				continue
			}
			if step.Action != model.ActionCreate {
				continue
			}

			kind, ok := model.ParseEntityKind(string(step.EntityKind))
			if !ok {
				fmt.Printf("  SKIP: unknown kind '%s' in step %s\n", step.EntityKind, step.ID)
				continue
			}

			newEntity := model.NewEntityForKind(kind)
			nb := newEntity.GetBase()
			nb.Name = step.EntityName
			nb.Description = step.Description

			fp, err := store.Create(newEntity, "")
			if err != nil {
				fmt.Printf("  ERROR: %s: %v\n", step.EntityName, err)
				continue
			}

			p.Steps[i].Status = model.StepCompleted
			fmt.Printf("  ✓ Created %s '%s' → %s\n", kind, step.EntityName, fp)
			created++
		}

		store.Update(p, body)

		if created == 0 {
			fmt.Println("No pending create steps to execute.")
		} else {
			fmt.Printf("\nExecuted %d steps.\n", created)
		}
		return nil
	},
}

var (
	addStepDesc       string
	addStepAction     string
	addStepEntityKind string
	addStepEntityName string
)

var planAddStepCmd = &cobra.Command{
	Use:   "add-step <plan-slug>",
	Short: "Add a structured step to a plan",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		stepID := fmt.Sprintf("step_%d", len(p.Steps)+1)
		step := model.PlanStep{
			ID:          stepID,
			Action:      model.PlanStepAction(addStepAction),
			EntityKind:  model.EntityKind(addStepEntityKind),
			EntityName:  addStepEntityName,
			Status:      model.StepPending,
			Description: addStepDesc,
		}
		p.Steps = append(p.Steps, step)

		if _, err := store.Update(p, body); err != nil {
			return err
		}

		fmt.Printf("Added step %s to plan '%s'\n", stepID, p.Name)
		fmt.Printf("  %s %s %s: %s\n", stepID, addStepAction, addStepEntityKind, addStepDesc)
		fmt.Printf("  Total steps: %d\n", len(p.Steps))
		return nil
	},
}

var planEstimateCmd = &cobra.Command{
	Use:   "estimate <slug>",
	Short: "Estimate plan size and suggest splitting",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, _, err := store.GetByKind(model.KindPlan, args[0])
		if err != nil {
			return fmt.Errorf("plan not found: %s", args[0])
		}
		p := entity.(*model.PlanEntity)

		if len(p.Steps) == 0 {
			fmt.Printf("Plan '%s' has no structured steps.\n", p.Name)
			fmt.Println("Add steps with: syde plan add-step")
			return nil
		}

		// Count by action and kind
		actionCounts := make(map[string]int)
		kindCounts := make(map[string]int)
		for _, s := range p.Steps {
			actionCounts[string(s.Action)]++
			if s.EntityKind != "" {
				kindCounts[string(s.EntityKind)]++
			}
		}

		cmdPerStep := 5 // start, query, write, verify, done
		totalEstimated := len(p.Steps) * cmdPerStep

		fmt.Printf("Plan: %s (%d steps)\n", p.Name, len(p.Steps))
		for action, count := range actionCounts {
			kinds := []string{}
			for kind, c := range kindCounts {
				kinds = append(kinds, fmt.Sprintf("%s × %d", kind, c))
			}
			fmt.Printf("  %s: %s\n", action, strings.Join(kinds, ", "))
			_ = count
		}
		fmt.Printf("\nEstimated commands: ~%d (at ~%d per step)\n", totalEstimated, cmdPerStep)

		if totalEstimated > 25 {
			subPlanSize := 5
			numSubPlans := (len(p.Steps) + subPlanSize - 1) / subPlanSize
			fmt.Printf("\n⚠ Recommendation: SPLIT into %d sub-plans of ~%d steps each\n", numSubPlans, subPlanSize)
			for i := 0; i < numSubPlans; i++ {
				start := i * subPlanSize
				end := start + subPlanSize
				if end > len(p.Steps) {
					end = len(p.Steps)
				}
				fmt.Printf("  Sub-plan %d: steps %d-%d (~%d commands)\n", i+1, start+1, end, (end-start)*cmdPerStep)
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
	planStepCmd.Flags().StringVar(&planStepStatus, "status", "", "new step status (pending, in_progress, completed, skipped)")
	planStepCmd.MarkFlagRequired("status")

	planAddStepCmd.Flags().StringVar(&addStepDesc, "description", "", "step description")
	planAddStepCmd.Flags().StringVar(&addStepAction, "action", "create", "action (create/update/delete)")
	planAddStepCmd.Flags().StringVar(&addStepEntityKind, "entity-kind", "", "target entity kind")
	planAddStepCmd.Flags().StringVar(&addStepEntityName, "entity-name", "", "target entity name")
	planAddStepCmd.MarkFlagRequired("description")

	planSyncCmd.Flags().StringVar(&planSyncFile, "file", "", "path to plan file (instead of scanning .claude/plans/)")

	planCmd.AddCommand(planCreateCmd, planListCmd, planShowCmd, planApproveCmd,
		planStepCmd, planSyncCmd, planExecuteCmd, planAddStepCmd, planEstimateCmd)
	rootCmd.AddCommand(planCmd)
}
