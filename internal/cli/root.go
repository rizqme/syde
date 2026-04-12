package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sydeDir string

var rootCmd = &cobra.Command{
	Use:   "syde",
	Short: "Text-first software design model CLI",
	Long:  "syde manages software design models stored as markdown files in .syde/",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&sydeDir, "dir", "", "path to .syde/ directory (auto-detected if empty)")
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
