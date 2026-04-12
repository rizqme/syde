package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var (
	contextJSON bool
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Full architecture snapshot in one call",
	Long:  "Outputs project architecture: entities, decisions, learnings, plans, tasks. Replaces separate status + constraints + list + plan list calls.",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		projectName, _ := loadProjectConfig()

		// Collect entities by kind
		type entry struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Stat string `json:"status"`
			Desc string `json:"description,omitempty"`
		}
		entitiesByKind := make(map[string][]entry)

		var decisions []map[string]string
		var learnings []map[string]string
		var plans []map[string]interface{}
		taskStats := map[string]int{"completed": 0, "in_progress": 0, "pending": 0, "blocked": 0}
		total := 0

		for _, kind := range model.AllEntityKinds() {
			entities, _ := store.List(kind)
			if len(entities) == 0 {
				continue
			}
			total += len(entities)

			for _, ewb := range entities {
				b := ewb.Entity.GetBase()

				switch v := ewb.Entity.(type) {
				case *model.DecisionEntity:
					if b.Status == model.StatusActive || b.Status == model.StatusDraft || b.Status == "" {
						decisions = append(decisions, map[string]string{
							"name": b.Name, "statement": v.Statement, "category": v.Category,
						})
					}
				case *model.LearningEntity:
					if v.ConfLevel == model.ConfidenceHigh || v.Category == model.CatGotcha || v.Category == model.CatConstraint {
						learnings = append(learnings, map[string]string{
							"name": b.Name, "category": string(v.Category),
							"confidence": string(v.ConfLevel), "description": v.Description,
						})
					}
				case *model.PlanEntity:
					done := 0
					for _, s := range v.Steps {
						if s.Status == model.StepCompleted || s.Status == model.StepSkipped {
							done++
						}
					}
					plans = append(plans, map[string]interface{}{
						"name": b.Name, "status": string(b.Status),
						"progress": v.Progress(), "total_steps": len(v.Steps), "done_steps": done,
					})
				case *model.TaskEntity:
					switch model.TaskStatus(b.Status) {
					case model.TaskCompleted:
						taskStats["completed"]++
					case model.TaskInProgress:
						taskStats["in_progress"]++
					case model.TaskBlocked:
						taskStats["blocked"]++
					default:
						taskStats["pending"]++
					}
				default:
					entitiesByKind[string(kind)] = append(entitiesByKind[string(kind)], entry{
						ID: b.ID, Name: b.Name, Stat: string(b.Status), Desc: b.Description,
					})
				}
			}
		}

		if contextJSON {
			out := map[string]interface{}{
				"project": projectName, "entities": entitiesByKind,
				"decisions": decisions, "learnings": learnings,
				"plans": plans, "tasks": taskStats, "total": total,
			}
			data, _ := json.Marshal(out)
			fmt.Println(string(data))
			return nil
		}

		// Human-readable
		fmt.Printf("Project: %s\n\n", projectName)

		for _, kindName := range []string{"system", "component", "contract", "concept", "flow", "design"} {
			entries := entitiesByKind[kindName]
			if len(entries) == 0 {
				continue
			}
			title := strings.ToUpper(kindName[:1]) + kindName[1:] + "s"
			if kindName == "system" && len(entries) == 1 {
				fmt.Printf("%s: %s — %s\n", title, entries[0].Name, entries[0].Desc)
				continue
			}
			names := make([]string, len(entries))
			for i, e := range entries {
				names[i] = fmt.Sprintf("%s [%s]", e.Name, e.Stat)
			}
			fmt.Printf("%s (%d): %s\n", title, len(entries), strings.Join(names, ", "))
		}

		if len(decisions) > 0 {
			fmt.Printf("Decisions (%d):\n", len(decisions))
			for _, d := range decisions {
				s := d["statement"]
				if s == "" {
					s = d["name"]
				}
				fmt.Printf("  • %s\n", s)
			}
		}

		if len(learnings) > 0 {
			fmt.Printf("Learnings (%d):\n", len(learnings))
			for _, l := range learnings {
				fmt.Printf("  ⚠ %s: %s\n", strings.ToUpper(l["category"]), l["description"])
			}
		}

		for _, p := range plans {
			fmt.Printf("Plan: %s — %v/%v steps (%.0f%%) [%s]\n",
				p["name"], p["done_steps"], p["total_steps"], p["progress"], p["status"])
		}

		totalTasks := taskStats["completed"] + taskStats["in_progress"] + taskStats["pending"] + taskStats["blocked"]
		if totalTasks > 0 {
			fmt.Printf("Tasks: %d done, %d active, %d pending\n",
				taskStats["completed"], taskStats["in_progress"], taskStats["pending"])
		}

		return nil
	},
}

func init() {
	contextCmd.Flags().BoolVar(&contextJSON, "json", false, "output as JSON")
	rootCmd.AddCommand(contextCmd)
}
