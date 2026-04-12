package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/scan"
	"github.com/spf13/cobra"
)

var (
	scanDryRun  bool
	scanCoverage bool
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Bootstrap design model from existing source code",
	Long:  "Generates a scan guide from the codebase structure, then uses Claude agents to extract a comprehensive design model.",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		if scanCoverage {
			guide, err := scan.LoadGuide(store.FS.Root)
			if err != nil {
				return fmt.Errorf("no scan guide found — run 'syde scan' first")
			}
			entries, err := scan.CheckCoverage(guide, store)
			if err != nil {
				return err
			}
			fmt.Print(scan.FormatCoverage(entries))
			return nil
		}

		guide, err := scan.GenerateGuide(store.FS.Root + "/..")
		if err != nil {
			return fmt.Errorf("generate scan guide: %w", err)
		}

		if err := scan.SaveGuide(store.FS.Root, guide); err != nil {
			return fmt.Errorf("save scan guide: %w", err)
		}

		fmt.Printf("Scan guide generated:\n")
		fmt.Printf("  Files: %d\n", guide.FileCount)
		fmt.Printf("  Languages: %d\n", len(guide.Languages))
		fmt.Printf("  Directories: %d\n", len(guide.DirectoryMap))
		fmt.Printf("  Saved: .syde/scan-guide.json\n")

		if scanDryRun {
			fmt.Println("\nDry run — no entities created.")
			fmt.Println("Run 'syde scan' without --dry-run to start agent-driven extraction.")
			return nil
		}

		fmt.Println("\nScan guide ready. Use the syde skill in Claude Code to run the 5-round agent extraction:")
		fmt.Println("  Round 1: System + Components")
		fmt.Println("  Round 2: Contracts + Concepts")
		fmt.Println("  Round 3: Flows")
		fmt.Println("  Round 4: Decisions")
		fmt.Println("  Round 5: Relationship wiring + validation")

		return nil
	},
}

func init() {
	scanCmd.Flags().BoolVar(&scanDryRun, "dry-run", false, "generate scan guide only")
	scanCmd.Flags().BoolVar(&scanCoverage, "coverage", false, "show coverage of codebase by components")
	rootCmd.AddCommand(scanCmd)
}
