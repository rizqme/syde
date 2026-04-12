package cli

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var (
	listStatus string
	listTag    string
	listFormat string
)

var listCmd = &cobra.Command{
	Use:   "list [kind]",
	Short: "List entities",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		var kinds []model.EntityKind
		if len(args) > 0 {
			kind, ok := model.ParseEntityKind(args[0])
			if !ok {
				return fmt.Errorf("unknown entity kind: %s", args[0])
			}
			kinds = []model.EntityKind{kind}
		} else {
			kinds = model.AllEntityKinds()
		}

		total := 0
		for _, kind := range kinds {
			entities, err := store.List(kind)
			if err != nil || len(entities) == 0 {
				continue
			}

			for _, ewb := range entities {
				b := ewb.Entity.GetBase()

				// Apply filters
				if listStatus != "" && string(b.Status) != listStatus {
					continue
				}
				if listTag != "" && !containsTag(b.Tags, listTag) {
					continue
				}

				if listFormat == "compact" {
					fmt.Printf("%s  %-12s %-30s %s\n", b.ID, b.Kind, b.Name, b.Status)
				} else {
					desc := b.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					fmt.Printf("  %-12s %-25s %-10s %s\n", b.Kind, b.Name, b.Status, desc)
				}
				total++
			}
		}

		if total == 0 {
			fmt.Println("No entities found.")
		} else {
			fmt.Printf("\n%d entities\n", total)
		}

		return nil
	},
}

func containsTag(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

func init() {
	listCmd.Flags().StringVar(&listStatus, "status", "", "filter by status")
	listCmd.Flags().StringVar(&listTag, "tag", "", "filter by tag")
	listCmd.Flags().StringVar(&listFormat, "format", "", "output format (compact)")
	rootCmd.AddCommand(listCmd)
}
