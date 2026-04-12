package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var getCmd = &cobra.Command{
	Use:   "get <id-or-slug>",
	Short: "Show entity details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.Get(slug)
		if err != nil {
			return fmt.Errorf("entity not found: %s", slug)
		}

		b := entity.GetBase()
		kindStr := string(b.Kind)
		if len(kindStr) > 0 {
			kindStr = strings.ToUpper(kindStr[:1]) + kindStr[1:]
		}
		fmt.Printf("═══ %s: %s ═══\n", kindStr, b.Name)
		fmt.Printf("ID: %s | Status: %s", b.ID, b.Status)
		if len(b.Tags) > 0 {
			fmt.Printf(" | Tags: %s", strings.Join(b.Tags, ", "))
		}
		fmt.Println()

		if b.Description != "" {
			fmt.Printf("\n  %s\n", b.Description)
		}
		if b.Purpose != "" {
			fmt.Printf("  Purpose: %s\n", b.Purpose)
		}

		// Print kind-specific fields
		fmBytes, _ := yaml.Marshal(entity)
		fmt.Printf("\n%s", string(fmBytes))

		if body != "" {
			fmt.Printf("\n%s\n", body)
		}

		if len(b.Relationships) > 0 {
			fmt.Println("\n── Relationships ──")
			for _, rel := range b.Relationships {
				label := ""
				if rel.Label != "" {
					label = " — " + rel.Label
				}
				fmt.Printf("  → %s: %s%s\n", rel.Type, rel.Target, label)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
