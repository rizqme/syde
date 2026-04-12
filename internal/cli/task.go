package cli

import (
	"fmt"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var (
	taskPlan     string
	taskStep     string
	taskEntity   []string
	taskPriority string
	taskReason   string
)

var taskCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		task := &model.TaskEntity{
			BaseEntity: model.BaseEntity{
				Kind:   model.KindTask,
				Name:   name,
				Status: model.Status(model.TaskPending),
			},
			Priority:   model.Priority(taskPriority),
			PlanRef:    taskPlan,
			PlanStep:   taskStep,
			EntityRefs: taskEntity,
			CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		}

		filePath, err := store.Create(task, "")
		if err != nil {
			return err
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
		store, err := openStore()
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
			switch model.TaskStatus(t.Status) {
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
			fmt.Printf("  %s %-40s %s%s\n", icon, t.Name, t.Status, pri)
		}
		return nil
	},
}

var taskShowCmd = &cobra.Command{
	Use:   "show <slug>",
	Short: "Show task details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
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
		fmt.Printf("Status: %s | Priority: %s\n", t.Status, t.Priority)
		if t.PlanRef != "" {
			fmt.Printf("Plan: %s (step: %s)\n", t.PlanRef, t.PlanStep)
		}
		if len(t.EntityRefs) > 0 {
			fmt.Printf("Entities: %v\n", t.EntityRefs)
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

var taskDoneCmd = &cobra.Command{
	Use:   "done <slug>",
	Short: "Mark task completed",
	Args:  cobra.ExactArgs(1),
	RunE:  func(cmd *cobra.Command, args []string) error { return setTaskStatus(args[0], model.TaskCompleted) },
}

var taskStartCmd = &cobra.Command{
	Use:   "start <slug>",
	Short: "Mark task in progress",
	Args:  cobra.ExactArgs(1),
	RunE:  func(cmd *cobra.Command, args []string) error { return setTaskStatus(args[0], model.TaskInProgress) },
}

var taskBlockCmd = &cobra.Command{
	Use:   "block <slug>",
	Short: "Mark task blocked",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()
		entity, body, err := store.GetByKind(model.KindTask, args[0])
		if err != nil {
			return fmt.Errorf("task not found: %s", args[0])
		}
		t := entity.(*model.TaskEntity)
		t.Status = model.Status(model.TaskBlocked)
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
		store, err := openStore()
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
				Kind:   model.KindTask,
				Name:   name,
				Status: model.Status(model.TaskPending),
			},
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
		store, err := openStore()
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

func setTaskStatus(slug string, status model.TaskStatus) error {
	store, err := openStore()
	if err != nil {
		return err
	}
	defer store.Close()

	entity, body, err := store.GetByKind(model.KindTask, slug)
	if err != nil {
		return fmt.Errorf("task not found: %s", slug)
	}
	t := entity.(*model.TaskEntity)
	t.Status = model.Status(status)
	if status == model.TaskCompleted {
		t.CompletedAt = time.Now().UTC().Format(time.RFC3339)
	}
	store.Update(t, body)

	// Auto-propagate to plan step if linked
	if status == model.TaskCompleted && t.PlanRef != "" && t.PlanStep != "" {
		planSlug := utils.Slugify(t.PlanRef)
		if pe, pb, err := store.GetByKind(model.KindPlan, planSlug); err == nil {
			p := pe.(*model.PlanEntity)
			for i := range p.Steps {
				if p.Steps[i].ID == t.PlanStep {
					p.Steps[i].Status = model.StepCompleted
					break
				}
			}
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
			store.Update(p, pb)
		}
	}

	fmt.Printf("%s → %s\n", t.Name, status)
	return nil
}

func init() {
	taskCreateCmd.Flags().StringVar(&taskPlan, "plan", "", "linked plan slug")
	taskCreateCmd.Flags().StringVar(&taskStep, "step", "", "linked plan step ID")
	taskCreateCmd.Flags().StringSliceVar(&taskEntity, "entity", nil, "linked entity slug/ID")
	taskCreateCmd.Flags().StringVar(&taskPriority, "priority", "medium", "priority (high, medium, low)")
	taskBlockCmd.Flags().StringVar(&taskReason, "reason", "", "block reason")

	taskCmd.AddCommand(taskCreateCmd, taskListCmd, taskShowCmd, taskDoneCmd, taskStartCmd, taskBlockCmd, taskSubCmd, taskLinkCmd)
	rootCmd.AddCommand(taskCmd)
}
