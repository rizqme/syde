package cli

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var contextJSON bool

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Full architecture snapshot in one call",
	Long:  "Outputs project architecture: entities, decisions, learnings, plans, tasks. Replaces separate status + constraints + list + plan list calls.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		raw, err := c.Context()
		if err != nil {
			return err
		}

		if contextJSON {
			fmt.Println(string(raw))
			return nil
		}

		// Human-readable: decode the payload and print the same layout
		// the old context command produced.
		var payload struct {
			Project  string `json:"project"`
			Entities map[string][]struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Desc string `json:"description,omitempty"`
			} `json:"entities"`
			Decisions []map[string]string      `json:"decisions"`
			Learnings []map[string]string      `json:"learnings"`
			Plans     []map[string]interface{} `json:"plans"`
			Tasks     map[string]int           `json:"tasks"`
			Total     int                      `json:"total"`
		}
		if err := json.Unmarshal(raw, &payload); err != nil {
			return fmt.Errorf("decode context: %w", err)
		}

		fmt.Printf("Project: %s\n\n", payload.Project)

		for _, kindName := range []string{"system", "component", "contract", "concept", "flow", "design"} {
			entries := payload.Entities[kindName]
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
				names[i] = e.Name
			}
			sort.Strings(names)
			fmt.Printf("%s (%d): %s\n", title, len(entries), strings.Join(names, ", "))
		}

		if len(payload.Decisions) > 0 {
			fmt.Printf("Decisions (%d):\n", len(payload.Decisions))
			for _, d := range payload.Decisions {
				s := d["statement"]
				if s == "" {
					s = d["name"]
				}
				fmt.Printf("  • %s\n", s)
			}
		}

		if len(payload.Learnings) > 0 {
			fmt.Printf("Learnings (%d):\n", len(payload.Learnings))
			for _, l := range payload.Learnings {
				fmt.Printf("  ⚠ %s: %s\n", strings.ToUpper(l["category"]), l["description"])
			}
		}

		for _, p := range payload.Plans {
			fmt.Printf("Plan: %s — %v/%v phases (%.0f%%) [%s]\n",
				p["name"], p["done_phases"], p["total_phases"], p["progress"], p["status"])
		}

		totalTasks := payload.Tasks["completed"] + payload.Tasks["in_progress"] + payload.Tasks["pending"] + payload.Tasks["blocked"]
		if totalTasks > 0 {
			fmt.Printf("Tasks: %d done, %d active, %d pending\n",
				payload.Tasks["completed"], payload.Tasks["in_progress"], payload.Tasks["pending"])
		}

		return nil
	},
}

func init() {
	contextCmd.Flags().BoolVar(&contextJSON, "json", false, "output as JSON")
	rootCmd.AddCommand(contextCmd)
}
