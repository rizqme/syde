package cli

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var constraintsJSON bool

var constraintsCmd = &cobra.Command{
	Use:   "constraints",
	Short: "Show active architectural constraints",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		// Collect active decisions
		decisions, _ := store.List(model.KindDecision)
		var activeDecisions []map[string]string
		for _, ewb := range decisions {
			d := ewb.Entity.(*model.DecisionEntity)
			if d.Status == model.StatusActive || d.Status == "" {
				activeDecisions = append(activeDecisions, map[string]string{
					"name":      d.Name,
					"statement": d.Statement,
					"category":  d.Category,
				})
			}
		}

		// Collect high-confidence gotchas/constraints
		learnings, _ := store.List(model.KindLearning)
		var criticalLearnings []map[string]string
		for _, ewb := range learnings {
			l := ewb.Entity.(*model.LearningEntity)
			if l.ConfLevel == model.ConfidenceHigh &&
				(l.Category == model.CatGotcha || l.Category == model.CatConstraint) {
				criticalLearnings = append(criticalLearnings, map[string]string{
					"name":        l.Name,
					"category":    string(l.Category),
					"description": l.Description,
				})
			}
		}

		if constraintsJSON {
			result := map[string]interface{}{
				"decisions": activeDecisions,
				"learnings": criticalLearnings,
			}
			data, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(data))
			return nil
		}

		// Human-readable output
		if len(activeDecisions) > 0 {
			fmt.Println("Active Decisions:")
			for _, d := range activeDecisions {
				fmt.Printf("  • %s", d["name"])
				if d["statement"] != "" {
					fmt.Printf(": %s", d["statement"])
				}
				fmt.Println()
			}
			fmt.Println()
		}

		if len(criticalLearnings) > 0 {
			fmt.Println("Critical Learnings:")
			for _, l := range criticalLearnings {
				fmt.Printf("  ⚠ %s: %s\n", l["category"], l["description"])
			}
			fmt.Println()
		}

		if len(activeDecisions) == 0 && len(criticalLearnings) == 0 {
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

		cfg, err := config.Load(dir)
		if err != nil {
			return err
		}

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		// Resolve file to component via component_paths
		absFile, _ := filepath.Abs(filePath)
		projectRoot := filepath.Dir(dir)
		relFile, _ := filepath.Rel(projectRoot, absFile)

		componentSlug := ""
		if cfg.ComponentPaths != nil {
			for slug, patterns := range cfg.ComponentPaths {
				for _, pattern := range patterns {
					if matched, _ := filepath.Match(pattern, relFile); matched {
						componentSlug = slug
						break
					}
					// Also try prefix match for ** patterns
					if strings.HasSuffix(pattern, "/**") {
						prefix := strings.TrimSuffix(pattern, "/**")
						if strings.HasPrefix(relFile, prefix+"/") || strings.HasPrefix(relFile, prefix+"\\") {
							componentSlug = slug
							break
						}
					}
				}
				if componentSlug != "" {
					break
				}
			}
		}

		if componentSlug == "" {
			if constraintsJSON {
				fmt.Println("{}")
			} else {
				fmt.Printf("No component mapping for: %s\n", relFile)
				fmt.Println("Add component_paths to syde.yaml to enable file-to-component mapping.")
			}
			return nil
		}

		entity, _, err := store.GetByKind(model.KindComponent, componentSlug)
		if err != nil {
			if constraintsJSON {
				fmt.Println("{}")
			} else {
				fmt.Printf("Component '%s' not found in .syde/\n", componentSlug)
			}
			return nil
		}

		comp := entity.(*model.ComponentEntity)

		if constraintsJSON {
			result := map[string]interface{}{
				"file":         relFile,
				"component":    comp.Name,
				"boundaries":   comp.Boundaries,
				"responsibility": comp.Responsibility,
			}

			// Get learnings
			learnings, _ := store.List(model.KindLearning)
			var compLearnings []map[string]string
			for _, ewb := range learnings {
				l := ewb.Entity.(*model.LearningEntity)
				for _, ref := range l.EntityRefs {
					if ref == comp.ID || ref == componentSlug {
						compLearnings = append(compLearnings, map[string]string{
							"category":    string(l.Category),
							"description": l.Description,
							"confidence":  string(l.ConfLevel),
						})
						break
					}
				}
			}
			result["learnings"] = compLearnings

			data, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(data))
		} else {
			fmt.Printf("File: %s\n", relFile)
			fmt.Printf("Component: %s\n\n", comp.Name)

			if comp.Boundaries != "" {
				fmt.Printf("  Boundaries: %s\n", comp.Boundaries)
			}
			if comp.Responsibility != "" {
				fmt.Printf("  Responsibility: %s\n", comp.Responsibility)
			}

			// Show learnings
			learnings, _ := store.List(model.KindLearning)
			for _, ewb := range learnings {
				l := ewb.Entity.(*model.LearningEntity)
				for _, ref := range l.EntityRefs {
					if ref == comp.ID || ref == componentSlug {
						fmt.Printf("\n  ⚠ %s: %s\n", l.Category, l.Description)
						break
					}
				}
			}
		}

		_ = utils.Slugify
		return nil
	},
}

func init() {
	constraintsCmd.Flags().BoolVar(&constraintsJSON, "json", false, "output as JSON")
	constraintsCheckCmd.Flags().BoolVar(&constraintsJSON, "json", false, "output as JSON")
	constraintsCmd.AddCommand(constraintsCheckCmd)
	rootCmd.AddCommand(constraintsCmd)
}
