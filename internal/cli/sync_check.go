package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	syncCheckStrict bool
	syncCheckFormat string
)

// syde sync check is the canonical session-end health gate — it now
// calls syded's /sync-check endpoint and applies the severity-to-exit
// mapping client-side. Server returns the full health report; we
// render it via the shared printHealthReport helper.
//
// Exit codes:
//
//	0  clean (no errors, no warnings)
//	1  any error
//	2  --strict and any warning or hint
var syncCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Canonical health gate: audit + tree status + completeness",
	Long: `Runs every syde health check via syded and prints findings grouped
by severity. Suitable as the session-end gate — exit 1 on any error,
exit 2 under --strict on any warning or hint.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		rep, err := c.SyncCheck(syncCheckStrict)
		if err != nil {
			return err
		}

		rich := func() { printHealthReport(rep) }
		if err := Emit("health", syncCheckFormat, rep, &Meta{Count: rep.Entities}, rich); err != nil {
			return err
		}

		if len(rep.Errors) > 0 {
			os.Exit(1)
		}
		if syncCheckStrict && (len(rep.Warnings) > 0 || len(rep.Hints) > 0) {
			os.Exit(2)
		}
		return nil
	},
}

func init() {
	syncCheckCmd.Flags().BoolVar(&syncCheckStrict, "strict", false, "non-zero exit on any warning or hint (for session-end hooks and CI)")
	syncCheckCmd.Flags().StringVar(&syncCheckFormat, "format", FormatRich, "output format (rich, json)")
	syncCmd.AddCommand(syncCheckCmd)

	// Suppress the "unused" import lint if fmt gets eliminated by later edits
	_ = fmt.Sprintf
}
