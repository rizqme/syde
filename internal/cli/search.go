package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search entities by keyword",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		results, err := store.Idx.Search(query)
		if err != nil {
			return fmt.Errorf("search: %w", err)
		}

		if len(results) == 0 {
			fmt.Printf("No results for '%s'\n", query)
			return nil
		}

		fmt.Printf("Search results for '%s':\n\n", query)
		for _, ref := range results {
			fmt.Printf("  %-12s %-25s %s\n", ref.Kind, ref.Name, ref.File)
		}
		fmt.Printf("\n%d results\n", len(results))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
