package cli

import (
	"fmt"
	"os"

	"github.com/feedloop/syde/internal/client"
	"github.com/spf13/cobra"
)

var validateFormat string

// syde validate is now a thin deprecated alias over the syded HTTP
// API — it calls /validate and filters out everything but errors
// to preserve its legacy errors-only semantics. The canonical health
// gate is `syde sync check`.
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check model integrity (deprecated — use 'syde sync check')",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(os.Stderr, "DEPRECATED: 'syde validate' is now an alias for 'syde sync check --errors-only'. Switch to 'syde sync check' for the full health report.")

		c, err := openClient()
		if err != nil {
			return err
		}
		rep, err := c.Validate()
		if err != nil {
			return err
		}

		// Errors-only filter.
		rep.Warnings = nil
		rep.Hints = nil

		rich := func() { printHealthReport(rep) }
		if err := Emit("health", validateFormat, rep, &Meta{Count: rep.Entities}, rich); err != nil {
			return err
		}

		if len(rep.Errors) > 0 {
			os.Exit(1)
		}
		return nil
	},
}

// printHealthReport renders a client.HealthReport the same way the old
// server-side printReport did. Used by both `syde validate` and
// `syde sync check`.
func printHealthReport(rep *client.HealthReport) {
	printFindings("ERROR", rep.Errors)
	printFindings("WARN ", rep.Warnings)
	printFindings("HINT ", rep.Hints)

	total := len(rep.Errors) + len(rep.Warnings) + len(rep.Hints)
	if total == 0 {
		fmt.Printf("Validation passed. %d entities checked.\n", rep.Entities)
		return
	}
	fmt.Printf("\n%d errors, %d warnings across %d entities\n", len(rep.Errors), len(rep.Warnings), rep.Entities)
}

func printFindings(prefix string, findings []client.Finding) {
	for _, f := range findings {
		scope := ""
		if f.EntityKind != "" {
			scope = fmt.Sprintf(" %s/%s:", f.EntityKind, f.EntityName)
		} else if f.Path != "" {
			scope = fmt.Sprintf(" tree/%s:", f.Path)
		}
		fmt.Printf("  %s%s %s\n", prefix, scope, f.Message)
	}
}

func init() {
	validateCmd.Flags().StringVar(&validateFormat, "format", FormatRich, "output format (rich, json)")
	rootCmd.AddCommand(validateCmd)
}
