package cli

import (
	"fmt"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/tree"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

// validateTaskReferences checks that each affected entity slug resolves
// to an existing entity and each affected file path exists as a node in
// .syde/tree.yaml. Returns an error listing every problem found.
func validateTaskReferences(store *writeClient, affectedEntities, affectedFiles []string) error {
	var problems []string

	for _, slug := range affectedEntities {
		if slug == "" {
			continue
		}
		if _, _, err := store.Get(slug); err != nil {
			problems = append(problems, fmt.Sprintf("affected entity %q does not exist", slug))
		}
	}

	if len(affectedFiles) > 0 {
		t, err := tree.Load(store.FS.Root)
		if err != nil {
			problems = append(problems, fmt.Sprintf("cannot load tree to verify affected files: %v", err))
		} else if t != nil {
			for _, fp := range affectedFiles {
				if fp == "" {
					continue
				}
				if t.Get(fp) == nil {
					problems = append(problems, fmt.Sprintf("affected file %q not in tree (run 'syde tree scan')", fp))
				}
			}
		}
	}

	if len(problems) > 0 {
		return fmt.Errorf("task reference validation failed:\n  - %s", joinLines(problems))
	}
	return nil
}

func joinLines(items []string) string {
	out := ""
	for i, s := range items {
		if i > 0 {
			out += "\n  - "
		}
		out += s
	}
	return out
}

// dedupeAppend unions two string slices in-order and drops duplicates.
// Preserves the original order of existing first, then the additions
// in the order they were passed. Used by the task-done merge so
// repeated --affected-entity flags collapse cleanly without reordering
// the historical list.
func dedupeAppend(existing, additions []string) []string {
	seen := make(map[string]bool, len(existing)+len(additions))
	out := make([]string, 0, len(existing)+len(additions))
	for _, s := range existing {
		if s == "" || seen[s] {
			continue
		}
		seen[s] = true
		out = append(out, s)
	}
	for _, s := range additions {
		if s == "" || seen[s] {
			continue
		}
		seen[s] = true
		out = append(out, s)
	}
	return out
}

// touchAffectedEntities bumps UpdatedAt on every entity that this task
// affected — either directly via task.AffectedEntities, or indirectly
// because one of the entity's Files appears in task.AffectedFiles.
// Returns the labels of entities that were touched (kind/name).
func touchAffectedEntities(store *writeClient, t *model.TaskEntity) []string {
	now := time.Now().UTC().Format(time.RFC3339)

	// Deduplicated set of entity IDs to touch.
	toTouch := make(map[string]bool)

	// 1. Direct listings
	for _, slug := range t.AffectedEntities {
		if slug == "" {
			continue
		}
		if e, _, err := store.Get(slug); err == nil {
			toTouch[e.GetBase().ID] = true
		}
	}

	// 2. Reverse lookup by affected file → owning entities
	if len(t.AffectedFiles) > 0 {
		fileSet := make(map[string]bool, len(t.AffectedFiles))
		for _, fp := range t.AffectedFiles {
			fileSet[fp] = true
		}
		all, _ := store.ListAll()
		for _, ewb := range all {
			b := ewb.Entity.GetBase()
			for _, fp := range b.Files {
				if fileSet[fp] {
					toTouch[b.ID] = true
					break
				}
			}
		}
	}

	// Apply UpdatedAt to each and persist.
	var labels []string
	if len(toTouch) == 0 {
		return labels
	}
	all, _ := store.ListAll()
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if !toTouch[b.ID] {
			continue
		}
		b.UpdatedAt = now
		store.Update(ewb.Entity, ewb.Body)
		labels = append(labels, fmt.Sprintf("%s/%s", b.Kind, b.Name))
	}
	return labels
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var (
	taskPlan             string
	taskPhase            string
	taskEntity           []string
	taskPriority         string
	taskReason           string
	taskObjective        string
	taskDetails          string
	taskAcceptance       string
	taskAffectedEntities []string
	taskAffectedFiles    []string
)

var taskCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		// Validate affected entities and files exist before creating.
		if err := validateTaskReferences(store, taskAffectedEntities, taskAffectedFiles); err != nil {
			return err
		}

		task := &model.TaskEntity{
			BaseEntity: model.BaseEntity{
				Kind: model.KindTask,
				Name: name,
			},
			TaskStatus:       model.TaskPending,
			Priority:         model.Priority(taskPriority),
			Objective:        taskObjective,
			Details:          taskDetails,
			Acceptance:       taskAcceptance,
			AffectedEntities: taskAffectedEntities,
			AffectedFiles:    taskAffectedFiles,
			PlanRef:          taskPlan,
			PlanPhase:        taskPhase,
			EntityRefs:       taskEntity,
			CreatedAt:        time.Now().UTC().Format(time.RFC3339),
		}

		filePath, err := store.Create(task, "")
		if err != nil {
			return err
		}

		// Auto-register task in plan phase
		if taskPlan != "" && taskPhase != "" {
			planSlug := utils.Slugify(taskPlan)
			if pe, pb, err := store.GetByKind(model.KindPlan, planSlug); err == nil {
				p := pe.(*model.PlanEntity)
				for i := range p.Phases {
					if p.Phases[i].ID == taskPhase {
						p.Phases[i].Tasks = append(p.Phases[i].Tasks, utils.Slugify(name))
						store.Update(p, pb)
						break
					}
				}
			}
		}

		fmt.Printf("Created task: %s\n", name)
		fmt.Printf("  ID: %s\n", task.ID)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		tasks, err := store.List(model.KindTask)
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		for _, ewb := range tasks {
			t := ewb.Entity.(*model.TaskEntity)
			icon := "○"
			switch t.TaskStatus {
			case model.TaskCompleted:
				icon = "✓"
			case model.TaskInProgress:
				icon = "●"
			case model.TaskBlocked:
				icon = "✗"
			case model.TaskCancelled:
				icon = "–"
			}
			pri := ""
			if t.Priority != "" {
				pri = fmt.Sprintf(" [%s]", t.Priority)
			}
			fmt.Printf("  %s %-40s %s%s\n", icon, t.Name, t.TaskStatus, pri)
		}
		return nil
	},
}

var taskShowCmd = &cobra.Command{
	Use:   "show <slug>",
	Short: "Show task details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindTask, args[0])
		if err != nil {
			return fmt.Errorf("task not found: %s", args[0])
		}
		t := entity.(*model.TaskEntity)
		fmt.Printf("═══ Task: %s ═══\n", t.Name)
		fmt.Printf("Status: %s | Priority: %s\n", t.TaskStatus, t.Priority)
		if t.PlanRef != "" {
			fmt.Printf("Plan: %s (phase: %s)\n", t.PlanRef, t.PlanPhase)
		}
		if t.Description != "" {
			fmt.Printf("\n%s\n", t.Description)
		}
		if t.Objective != "" {
			fmt.Printf("\nOBJECTIVE\n  %s\n", t.Objective)
		}
		if t.Details != "" {
			fmt.Printf("\nDETAILS\n  %s\n", t.Details)
		}
		if t.Acceptance != "" {
			fmt.Printf("\nACCEPTANCE\n  %s\n", t.Acceptance)
		}
		if len(t.AffectedEntities) > 0 {
			fmt.Printf("\nAFFECTED ENTITIES\n")
			for _, s := range t.AffectedEntities {
				fmt.Printf("  - %s\n", s)
			}
		}
		if len(t.AffectedFiles) > 0 {
			fmt.Printf("\nAFFECTED FILES\n")
			for _, s := range t.AffectedFiles {
				fmt.Printf("  - %s\n", s)
			}
		}
		if len(t.EntityRefs) > 0 {
			fmt.Printf("\nLegacy entity refs: %v\n", t.EntityRefs)
		}
		if t.ParentTask != "" {
			fmt.Printf("Parent: %s\n", t.ParentTask)
		}
		if len(t.Subtasks) > 0 {
			fmt.Printf("Subtasks: %v\n", t.Subtasks)
		}
		if t.BlockReason != "" {
			fmt.Printf("Blocked: %s\n", t.BlockReason)
		}
		if body != "" {
			fmt.Printf("\n%s\n", body)
		}
		return nil
	},
}

// Done-specific flags so the completing agent can declare the REAL set
// of entities and files the task touched. The create-time set is a
// prediction; the done-time set is reality. Merged into the task's
// existing affected lists, validated, persisted, then fed to
// touchAffectedEntities so drift clears for the full set.
var (
	taskDoneAffectedEntities []string
	taskDoneAffectedFiles    []string
)

var taskDoneCmd = &cobra.Command{
	Use:   "done <slug>",
	Short: "Mark task completed; optionally declare the real affected entities/files touched during implementation",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return setTaskStatus(args[0], model.TaskCompleted, taskDoneAffectedEntities, taskDoneAffectedFiles)
	},
}

var taskStartCmd = &cobra.Command{
	Use:   "start <slug>",
	Short: "Mark task in progress",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return setTaskStatus(args[0], model.TaskInProgress, nil, nil)
	},
}

var taskBlockCmd = &cobra.Command{
	Use:   "block <slug>",
	Short: "Mark task blocked",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()
		entity, body, err := store.GetByKind(model.KindTask, args[0])
		if err != nil {
			return fmt.Errorf("task not found: %s", args[0])
		}
		t := entity.(*model.TaskEntity)
		t.TaskStatus = model.TaskBlocked
		t.BlockReason = taskReason
		store.Update(t, body)
		fmt.Printf("Blocked: %s\n", t.Name)
		if taskReason != "" {
			fmt.Printf("  Reason: %s\n", taskReason)
		}
		return nil
	},
}

var taskSubCmd = &cobra.Command{
	Use:   "sub <parent-slug> <name>",
	Short: "Create subtask under parent",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		parentSlug, name := args[0], args[1]
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		parent, parentBody, err := store.GetByKind(model.KindTask, parentSlug)
		if err != nil {
			return fmt.Errorf("parent task not found: %s", parentSlug)
		}
		pt := parent.(*model.TaskEntity)

		sub := &model.TaskEntity{
			BaseEntity: model.BaseEntity{
				Kind: model.KindTask,
				Name: name,
			},
			TaskStatus: model.TaskPending,
			ParentTask: pt.ID,
			Priority:   pt.Priority,
			PlanRef:    pt.PlanRef,
			CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		}
		filePath, err := store.Create(sub, "")
		if err != nil {
			return err
		}

		pt.Subtasks = append(pt.Subtasks, sub.ID)
		store.Update(pt, parentBody)

		fmt.Printf("Created subtask: %s\n", name)
		fmt.Printf("  Parent: %s\n", pt.Name)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

var taskLinkCmd = &cobra.Command{
	Use:   "link <task-slug> <entity-slug>",
	Short: "Link task to design entity",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindTask, args[0])
		if err != nil {
			return fmt.Errorf("task not found: %s", args[0])
		}
		t := entity.(*model.TaskEntity)

		target, _, err := store.Get(args[1])
		if err != nil {
			return fmt.Errorf("entity not found: %s", args[1])
		}

		t.EntityRefs = append(t.EntityRefs, target.GetBase().ID)
		store.Update(t, body)
		fmt.Printf("Linked task '%s' to %s '%s'\n", t.Name, target.GetBase().Kind, target.GetBase().Name)
		return nil
	},
}

var taskUpdateCmd = &cobra.Command{
	Use:   "update <slug>",
	Short: "Update task fields (objective, details, acceptance, priority)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindTask, args[0])
		if err != nil {
			return fmt.Errorf("task not found: %s", args[0])
		}
		t := entity.(*model.TaskEntity)

		if cmd.Flags().Changed("objective") {
			t.Objective = taskObjective
		}
		if cmd.Flags().Changed("details") {
			t.Details = taskDetails
		}
		if cmd.Flags().Changed("acceptance") {
			t.Acceptance = taskAcceptance
		}
		if cmd.Flags().Changed("priority") {
			t.Priority = model.Priority(taskPriority)
		}
		if cmd.Flags().Changed("description") {
			t.Description = cmd.Flag("description").Value.String()
		}
		if cmd.Flags().Changed("affected-entity") {
			if err := validateTaskReferences(store, taskAffectedEntities, nil); err != nil {
				return err
			}
			t.AffectedEntities = taskAffectedEntities
		}
		if cmd.Flags().Changed("affected-file") {
			if err := validateTaskReferences(store, nil, taskAffectedFiles); err != nil {
				return err
			}
			t.AffectedFiles = taskAffectedFiles
		}

		if _, err := store.Update(t, body); err != nil {
			return err
		}
		fmt.Printf("Updated task: %s\n", t.Name)
		return nil
	},
}

// setTaskStatus flips a task to the given status. When addAffected* are
// non-empty (only the done path passes them today) they are merged
// into the task's existing AffectedEntities / AffectedFiles, deduped,
// validated, and persisted as part of the same Update — touchAffected
// Entities below then sees the full set so the drift cascade clears
// every entity the task actually touched, not just what was predicted
// at create time.
func setTaskStatus(slug string, status model.TaskStatus, addAffectedEntities, addAffectedFiles []string) error {
	store, err := openWriteClient()
	if err != nil {
		return err
	}
	defer store.Close()

	entity, body, err := store.GetByKind(model.KindTask, slug)
	if err != nil {
		return fmt.Errorf("task not found: %s", slug)
	}
	t := entity.(*model.TaskEntity)

	// Merge the done-time additions into the task BEFORE we flip status,
	// so the validator runs over only the new entries (existing ones
	// were already validated at create or a previous done).
	if len(addAffectedEntities) > 0 || len(addAffectedFiles) > 0 {
		newEntities := dedupeAppend(t.AffectedEntities, addAffectedEntities)
		newFiles := dedupeAppend(t.AffectedFiles, addAffectedFiles)
		if err := validateTaskReferences(store, addAffectedEntities, addAffectedFiles); err != nil {
			return err
		}
		t.AffectedEntities = newEntities
		t.AffectedFiles = newFiles
	}

	t.TaskStatus = status
	if status == model.TaskCompleted {
		t.CompletedAt = time.Now().UTC().Format(time.RFC3339)
	}
	store.Update(t, body)

	// On completion, touch every entity the task affected so the drift
	// validator knows they've been reviewed as part of this task. Two
	// sources:
	//   1. t.AffectedEntities — the entities the task explicitly listed
	//   2. Any entity whose .Files overlaps t.AffectedFiles — auto-
	//      discovered so listing affected_files is sufficient.
	if status == model.TaskCompleted {
		touched := touchAffectedEntities(store, t)
		if len(touched) > 0 {
			fmt.Printf("Touched %d entity/entities (drift cleared):\n", len(touched))
			for _, name := range touched {
				fmt.Printf("  - %s\n", name)
			}
		}
	}

	// Auto-propagate to plan phase if linked
	if status == model.TaskCompleted && t.PlanRef != "" && t.PlanPhase != "" {
		planSlug := utils.Slugify(t.PlanRef)
		if pe, pb, err := store.GetByKind(model.KindPlan, planSlug); err == nil {
			p := pe.(*model.PlanEntity)

			// Check if ALL tasks on this phase are completed
			phase := p.PhaseByID(t.PlanPhase)
			if phase != nil {
				allTasksDone := true
				taskEntities, _ := store.List(model.KindTask)
				for _, tSlug := range phase.Tasks {
					for _, tewb := range taskEntities {
						te := tewb.Entity.(*model.TaskEntity)
						if utils.Slugify(te.Name) == tSlug && te.TaskStatus != model.TaskCompleted && te.TaskStatus != model.TaskCancelled {
							allTasksDone = false
							break
						}
					}
				}
				// Only auto-complete the phase if ALL its tasks are done
				if allTasksDone {
					for i := range p.Phases {
						if p.Phases[i].ID == t.PlanPhase {
							p.Phases[i].Status = model.PhaseCompleted
							fmt.Printf("Phase %s auto-completed (all tasks done)\n", t.PlanPhase)
							break
						}
					}
				}
			}

			// Check if ALL phases done → auto-complete plan
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
			store.Update(p, pb)
		}
	}

	fmt.Printf("%s → %s\n", t.Name, status)
	return nil
}

func init() {
	taskCreateCmd.Flags().StringVar(&taskPlan, "plan", "", "linked plan slug")
	taskCreateCmd.Flags().StringVar(&taskPhase, "phase", "", "linked plan phase ID")
	taskCreateCmd.Flags().StringSliceVar(&taskEntity, "entity", nil, "legacy: linked entity slug/ID (prefer --affected-entity)")
	taskCreateCmd.Flags().StringVar(&taskPriority, "priority", "medium", "priority (high, medium, low)")
	taskCreateCmd.Flags().StringVar(&taskObjective, "objective", "", "what this task achieves")
	taskCreateCmd.Flags().StringVar(&taskDetails, "details", "", "implementation specifics (files, approach)")
	taskCreateCmd.Flags().StringVar(&taskAcceptance, "acceptance", "", "how to know the task is done")
	taskCreateCmd.Flags().StringSliceVar(&taskAffectedEntities, "affected-entity", nil, "existing entity slug this task will modify (repeatable)")
	taskCreateCmd.Flags().StringSliceVar(&taskAffectedFiles, "affected-file", nil, "concrete source file path this task will modify; must exist in tree (repeatable)")
	taskBlockCmd.Flags().StringVar(&taskReason, "reason", "", "block reason")

	// Done-time affected declarations — merged into the task's existing
	// affected lists, validated, and used by touchAffectedEntities so
	// drift clears for the real set of entities/files the task touched.
	taskDoneCmd.Flags().StringArrayVar(&taskDoneAffectedEntities, "affected-entity", nil, "entity slug modified by this task (repeatable, merged with existing affected_entities)")
	taskDoneCmd.Flags().StringArrayVar(&taskDoneAffectedFiles, "affected-file", nil, "concrete file path modified by this task (repeatable, merged with existing affected_files)")

	taskUpdateCmd.Flags().StringVar(&taskObjective, "objective", "", "update task objective")
	taskUpdateCmd.Flags().StringVar(&taskDetails, "details", "", "update task details")
	taskUpdateCmd.Flags().StringVar(&taskAcceptance, "acceptance", "", "update task acceptance criteria")
	taskUpdateCmd.Flags().StringVar(&taskPriority, "priority", "", "update priority")
	taskUpdateCmd.Flags().String("description", "", "update description")
	taskUpdateCmd.Flags().StringSliceVar(&taskAffectedEntities, "affected-entity", nil, "replace affected entities list (repeatable)")
	taskUpdateCmd.Flags().StringSliceVar(&taskAffectedFiles, "affected-file", nil, "replace affected files list (repeatable)")

	taskCmd.AddCommand(taskCreateCmd, taskListCmd, taskShowCmd, taskDoneCmd, taskStartCmd, taskBlockCmd, taskSubCmd, taskLinkCmd, taskUpdateCmd)
	rootCmd.AddCommand(taskCmd)
}
