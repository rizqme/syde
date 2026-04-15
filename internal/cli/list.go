package cli

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/client"
	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var (
	listTag    string
	listFormat string
)

var listCmd = &cobra.Command{
	Use:   "list [kind]",
	Short: "List entities",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}

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

		var items []client.EntitySummary
		for _, kind := range kinds {
			hits, err := c.List(string(kind), listTag)
			if err != nil {
				return err
			}
			items = append(items, hits...)
		}

		rich := func() {
			for _, it := range items {
				if listFormat == FormatCompact {
					fmt.Printf("%s  %-12s %s\n", it.ID, it.Kind, it.Name)
					continue
				}
				desc := it.Description
				if len(desc) > 60 {
					desc = desc[:57] + "..."
				}
				fmt.Printf("  %-12s %-25s %s\n", it.Kind, it.Name, desc)
			}
			if len(items) == 0 {
				fmt.Println("No entities found.")
			} else {
				fmt.Printf("\n%d entities\n", len(items))
			}
		}
		return Emit("list", listFormat, items, &Meta{Count: len(items)}, rich)
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
	listCmd.Flags().StringVar(&listTag, "tag", "", "filter by tag")
	listCmd.Flags().StringVar(&listFormat, "format", FormatRich, "output format (rich, compact, json)")
	rootCmd.AddCommand(listCmd)
}
