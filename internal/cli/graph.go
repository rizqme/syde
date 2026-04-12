package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/graph"
	"github.com/spf13/cobra"
)

var (
	graphFormat string
	graphDepth  int
)

var graphCmd = &cobra.Command{
	Use:   "graph [entity]",
	Short: "Show entity relationships",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		if len(args) == 0 {
			// Show summary of all relationships
			all, _ := store.ListAll()
			totalRels := 0
			for _, ewb := range all {
				totalRels += len(ewb.Entity.GetBase().Relationships)
			}
			fmt.Printf("Graph: %d entities, %d relationships\n", len(all), totalRels)
			fmt.Println("Specify an entity to see its connections: syde graph <slug>")
			return nil
		}

		slug := args[0]
		entity, _, err := store.Get(slug)
		if err != nil {
			return fmt.Errorf("entity not found: %s", slug)
		}

		b := entity.GetBase()
		result, err := graph.Neighbors(store.Idx, b.ID)
		if err != nil {
			return err
		}

		if graphFormat == "dot" {
			fmt.Print(graph.RenderDOT(b.Name, result))
		} else {
			fmt.Print(graph.RenderASCII(b.Name, result))
		}

		return nil
	},
}

func init() {
	graphCmd.Flags().StringVar(&graphFormat, "format", "ascii", "output format (ascii, dot)")
	graphCmd.Flags().IntVar(&graphDepth, "depth", 1, "traversal depth")
	rootCmd.AddCommand(graphCmd)
}
