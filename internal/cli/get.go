package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var getFormat string

var getCmd = &cobra.Command{
	Use:   "get <id-or-slug>",
	Short: "Show entity details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		raw, err := c.Get(args[0])
		if err != nil {
			return err
		}

		if getFormat == FormatJSON {
			fmt.Println(string(raw))
			return nil
		}

		// Rich mode: decode the server-side JSON (produced by
		// query.FormatJSON on the ResolvedEntity) and pretty-print the
		// salient fields.
		var payload map[string]interface{}
		if err := json.Unmarshal(raw, &payload); err != nil {
			return fmt.Errorf("decode: %w", err)
		}
		entity, _ := payload["entity"].(map[string]interface{})
		if entity == nil {
			fmt.Println(string(raw))
			return nil
		}
		kindStr, _ := entity["kind"].(string)
		name, _ := entity["name"].(string)
		id, _ := entity["id"].(string)
		desc, _ := entity["description"].(string)

		title := kindStr
		if len(title) > 0 {
			title = strings.ToUpper(title[:1]) + title[1:]
		}
		fmt.Printf("═══ %s: %s ═══\n", title, name)
		fmt.Printf("ID: %s\n", id)
		if desc != "" {
			fmt.Printf("\n  %s\n", desc)
		}

		// Relationships section if present.
		if rels, ok := payload["relationships"].([]interface{}); ok && len(rels) > 0 {
			fmt.Println("\n── Relationships ──")
			for _, r := range rels {
				rm, _ := r.(map[string]interface{})
				typ, _ := rm["type"].(string)
				target, _ := rm["target"].(string)
				label, _ := rm["label"].(string)
				if label != "" {
					fmt.Printf("  → %s: %s — %s\n", typ, target, label)
				} else {
					fmt.Printf("  → %s: %s\n", typ, target)
				}
			}
		}
		return nil
	},
}

func init() {
	getCmd.Flags().StringVar(&getFormat, "format", FormatRich, "output format (rich, json)")
	rootCmd.AddCommand(getCmd)
}
