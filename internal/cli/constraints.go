package cli

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/spf13/cobra"
)

var constraintsJSON bool

var constraintsCmd = &cobra.Command{
	Use:   "constraints",
	Short: "Show active architectural constraints",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		body, err := c.Constraints()
		if err != nil {
			return err
		}

		if constraintsJSON {
			fmt.Println(string(body))
			return nil
		}
		var payload struct {
			Decisions []map[string]string `json:"decisions"`
			Learnings []map[string]string `json:"learnings"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		if len(payload.Decisions) > 0 {
			fmt.Println("Active Decisions:")
			for _, d := range payload.Decisions {
				fmt.Printf("  • %s", d["name"])
				if d["statement"] != "" {
					fmt.Printf(": %s", d["statement"])
				}
				fmt.Println()
			}
			fmt.Println()
		}
		if len(payload.Learnings) > 0 {
			fmt.Println("Critical Learnings:")
			for _, l := range payload.Learnings {
				fmt.Printf("  ⚠ %s: %s\n", l["category"], l["description"])
			}
			fmt.Println()
		}
		if len(payload.Decisions) == 0 && len(payload.Learnings) == 0 {
			fmt.Println("No active constraints.")
		}
		return nil
	},
}

var constraintsCheckCmd = &cobra.Command{
	Use:   "check <file>",
	Short: "Map source file to component and show relevant constraints",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		dir := sydeDir
		if dir == "" {
			dir, _ = config.FindSydeDir()
		}
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}

		// Normalize to project-relative path so syded matches against
		// the same component_paths keys `syde.yaml` uses.
		absFile, _ := filepath.Abs(filePath)
		projectRoot := filepath.Dir(dir)
		relFile, _ := filepath.Rel(projectRoot, absFile)

		c, err := openClient()
		if err != nil {
			return err
		}
		raw, err := c.ConstraintsCheck(relFile)
		if err != nil {
			return err
		}

		if constraintsJSON {
			fmt.Println(string(raw))
			return nil
		}

		var payload struct {
			File           string              `json:"file"`
			Component      string              `json:"component"`
			Responsibility string              `json:"responsibility"`
			Boundaries     string              `json:"boundaries"`
			Learnings      []map[string]string `json:"learnings"`
		}
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}
		if payload.Component == "" {
			fmt.Printf("No component mapping for: %s\n", relFile)
			fmt.Println("Add component_paths to syde.yaml to enable file-to-component mapping.")
			return nil
		}
		fmt.Printf("File: %s\n", payload.File)
		fmt.Printf("Component: %s\n\n", payload.Component)
		if payload.Boundaries != "" {
			fmt.Printf("  Boundaries: %s\n", payload.Boundaries)
		}
		if payload.Responsibility != "" {
			fmt.Printf("  Responsibility: %s\n", payload.Responsibility)
		}
		for _, l := range payload.Learnings {
			fmt.Printf("\n  ⚠ %s: %s\n", l["category"], l["description"])
		}
		return nil
	},
}

func init() {
	constraintsCmd.Flags().BoolVar(&constraintsJSON, "json", false, "output as JSON")
	constraintsCheckCmd.Flags().BoolVar(&constraintsJSON, "json", false, "output as JSON")
	constraintsCmd.AddCommand(constraintsCheckCmd)
	rootCmd.AddCommand(constraintsCmd)
}
