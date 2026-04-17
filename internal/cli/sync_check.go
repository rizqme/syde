package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var syncCheckFormat string

// syde sync check is the canonical session-end health gate. The audit
// engine emits a single severity level — every finding blocks. Exit
// codes:
//
//	0  clean (no findings)
//	1  any finding
var syncCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Canonical health gate: audit + tree status + completeness",
	Long: `Runs every syde health check via syded and prints findings. The
audit engine uses a single strict severity level; the gate exits
non-zero on any finding. Suitable for session-end hooks and CI.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		// Server still accepts the legacy strict parameter but it now
		// has no effect — always true is the only mode.
		rep, err := c.SyncCheck(true)
		if err != nil {
			return err
		}

		rich := func() { printHealthReport(rep) }
		if err := Emit("health", syncCheckFormat, rep, &Meta{Count: rep.Entities}, rich); err != nil {
			return err
		}

		if len(rep.Errors) > 0 || len(rep.Warnings) > 0 || len(rep.Hints) > 0 {
			os.Exit(1)
		}
		return nil
	},
}

func init() {
	syncCheckCmd.Flags().StringVar(&syncCheckFormat, "format", FormatRich, "output format (rich, json)")
	syncCmd.AddCommand(syncCheckCmd)

	_ = fmt.Sprintf
}
