package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "Rebuild the BadgerDB index from markdown files",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		fmt.Println("Rebuilding index...")
		stats, err := store.Reindex()
		if err != nil {
			return fmt.Errorf("reindex: %w", err)
		}

		fmt.Printf("  Entities:      %d\n", stats.Entities)
		fmt.Printf("  Relationships: %d\n", stats.Relationships)
		fmt.Printf("  Tags:          %d\n", stats.Tags)
		fmt.Printf("  Words:         %d\n", stats.Words)
		fmt.Println("Done.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(reindexCmd)
}
