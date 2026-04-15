package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "Rebuild the BadgerDB index from markdown files (via syded)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		fmt.Println("Rebuilding index via syded...")
		stats, err := c.Reindex(nil, true)
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
