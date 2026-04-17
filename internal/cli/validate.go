package cli

import (
	"fmt"
	"os"

	"github.com/feedloop/syde/internal/client"
	"github.com/spf13/cobra"
)

var validateFormat string

// syde validate is a deprecated alias over the syded HTTP API that
// calls /validate. The canonical health gate is `syde sync check`.
// Both share the same single-severity Finding model; any finding
// blocks.
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check model integrity (deprecated — use 'syde sync check')",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(os.Stderr, "DEPRECATED: 'syde validate' is an alias for 'syde sync check'. Switch to 'syde sync check' for the full health report.")

		c, err := openClient()
		if err != nil {
			return err
		}
		rep, err := c.Validate()
		if err != nil {
			return err
		}

		rich := func() { printHealthReport(rep) }
		if err := Emit("health", validateFormat, rep, &Meta{Count: rep.Entities}, rich); err != nil {
			return err
		}

		if len(rep.Errors) > 0 || len(rep.Warnings) > 0 || len(rep.Hints) > 0 {
			os.Exit(1)
		}
		return nil
	},
}

// printHealthReport renders a client.HealthReport. The audit engine
// now emits a single severity level; the Warnings/Hints slots are
// always empty but the renderer still iterates them for backward
// compatibility with older server payloads.
func printHealthReport(rep *client.HealthReport) {
	printFindings("FINDING", rep.Errors)
	printFindings("FINDING", rep.Warnings)
	printFindings("FINDING", rep.Hints)

	total := len(rep.Errors) + len(rep.Warnings) + len(rep.Hints)
	if total == 0 {
		fmt.Printf("Validation passed. %d entities checked.\n", rep.Entities)
		return
	}
	fmt.Printf("\n%d finding(s) across %d entities — every finding blocks\n", total, rep.Entities)
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
