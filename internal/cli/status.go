package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var statusFormat string

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show design model overview",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		st, err := c.Status()
		if err != nil {
			return err
		}

		projectName, projectVersion := loadProjectConfig()

		data := struct {
			Project string         `json:"project,omitempty"`
			Version string         `json:"version,omitempty"`
			Counts  map[string]int `json:"counts"`
			Total   int            `json:"total"`
		}{
			Project: projectName,
			Version: projectVersion,
			Counts:  st.Counts,
			Total:   st.Total,
		}

		rich := func() {
			if projectName != "" {
				fmt.Printf("Project: %s (v%s)\n", projectName, projectVersion)
			}
			fmt.Println()
			// Iterate kinds in canonical order so output is stable.
			for _, kind := range model.AllEntityKinds() {
				if st.Counts[string(kind)] == 0 {
					continue
				}
				fmt.Printf("  %-12s %d\n", kind.KindPlural(), st.Counts[string(kind)])
			}
			if st.Total == 0 {
				fmt.Println("  (no entities yet)")
			} else {
				fmt.Printf("\n  Total: %d entities\n", st.Total)
			}
		}
		return Emit("status", statusFormat, data, &Meta{Count: st.Total}, rich)
	},
}

func init() {
	statusCmd.Flags().StringVar(&statusFormat, "format", FormatRich, "output format (rich, json)")
	rootCmd.AddCommand(statusCmd)
}
