package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search entities by keyword",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		c, err := openClient()
		if err != nil {
			return err
		}
		raw, err := c.Search(query)
		if err != nil {
			return fmt.Errorf("search: %w", err)
		}

		// Server returns {query, results, count} — decode and render.
		var body struct {
			Query   string `json:"query"`
			Count   int    `json:"count"`
			Results []struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
				File string `json:"file"`
			} `json:"results"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			return err
		}
		if body.Count == 0 {
			fmt.Printf("No results for '%s'\n", query)
			return nil
		}
		fmt.Printf("Search results for '%s':\n\n", query)
		for _, hit := range body.Results {
			fmt.Printf("  %-12s %-25s %s\n", hit.Kind, hit.Name, hit.File)
		}
		fmt.Printf("\n%d results\n", body.Count)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
