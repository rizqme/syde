package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check model integrity",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		all, err := store.ListAll()
		if err != nil {
			return err
		}

		totalErrors := 0
		totalWarnings := 0

		for _, ewb := range all {
			errs := model.ValidateEntity(ewb.Entity)
			b := ewb.Entity.GetBase()

			for _, e := range errs {
				if e.Message == "required" {
					fmt.Printf("  ERROR  %s/%s: %s is %s\n", b.Kind, b.Name, e.Field, e.Message)
					totalErrors++
				} else {
					fmt.Printf("  WARN   %s/%s: %s is %s\n", b.Kind, b.Name, e.Field, e.Message)
					totalWarnings++
				}
			}

			// Check relationship targets exist
			for _, rel := range b.Relationships {
				found := false
				for _, other := range all {
					if other.Entity.GetBase().ID == rel.Target {
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("  WARN   %s/%s: relationship target %s not found\n", b.Kind, b.Name, rel.Target)
					totalWarnings++
				}
			}
		}

		if totalErrors == 0 && totalWarnings == 0 {
			fmt.Printf("Validation passed. %d entities checked.\n", len(all))
		} else {
			fmt.Printf("\n%d errors, %d warnings across %d entities\n", totalErrors, totalWarnings, len(all))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
