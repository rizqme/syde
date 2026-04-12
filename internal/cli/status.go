package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show design model overview",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		dir := sydeDir
		if dir == "" {
			dir, _ = config.FindSydeDir()
		}

		cfg, _ := config.Load(dir)
		if cfg != nil {
			fmt.Printf("Project: %s (v%s)\n", cfg.Project, cfg.Version)
		}
		fmt.Println()

		total := 0
		for _, kind := range model.AllEntityKinds() {
			entities, err := store.List(kind)
			if err != nil || len(entities) == 0 {
				continue
			}
			fmt.Printf("  %-12s %d\n", kind.KindPlural(), len(entities))
			total += len(entities)
		}

		if total == 0 {
			fmt.Println("  (no entities yet)")
		} else {
			fmt.Printf("\n  Total: %d entities\n", total)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
