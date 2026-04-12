package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var removeForce bool

var removeCmd = &cobra.Command{
	Use:     "remove <id-or-slug>",
	Aliases: []string{"rm"},
	Short:   "Remove an entity",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, _, err := store.Get(slug)
		if err != nil {
			return fmt.Errorf("entity not found: %s", slug)
		}

		b := entity.GetBase()

		if !removeForce {
			fmt.Printf("Remove %s '%s' (%s)? [y/N] ", b.Kind, b.Name, b.ID)
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(answer)), "y") {
				fmt.Println("Cancelled.")
				return nil
			}
		}

		entitySlug := utils.Slugify(b.Name)
		if err := store.Delete(b.Kind, entitySlug); err != nil {
			return fmt.Errorf("delete: %w", err)
		}

		fmt.Printf("Removed %s: %s\n", b.Kind, b.Name)
		return nil
	},
}

func init() {
	removeCmd.Flags().BoolVar(&removeForce, "force", false, "skip confirmation")
	rootCmd.AddCommand(removeCmd)
}
