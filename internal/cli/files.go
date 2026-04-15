package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Inspect source file ↔ entity coverage",
}

var filesOrphansCmd = &cobra.Command{
	Use:   "orphans",
	Short: "List non-ignored source files with no owning entity",
	Long: `Prints one path per line — every file in the summary tree (not marked
ignored) that has zero entities claiming it in their --file list.

Exits non-zero when any orphan exists, so the command composes in shell
pipelines and CI gates.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}
		orphans, err := c.FilesOrphans()
		if err != nil {
			return err
		}
		for _, p := range orphans {
			fmt.Println(p)
		}
		if len(orphans) > 0 {
			os.Exit(1)
		}
		return nil
	},
}

var filesCoverageCmd = &cobra.Command{
	Use:   "coverage [path]",
	Short: "Show which entities own a given file (or every file)",
	Long: `With a path argument, prints the entities that claim that file in
their --file list, one per line. Without arguments, prints every
non-ignored file in the tree followed by its owners (or "<orphan>" if
none). Useful for spotting duplicated ownership or gaps.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}

		if len(args) == 1 {
			raw, err := c.FilesCoverage(args[0])
			if err != nil {
				return err
			}
			var body struct {
				Path   string   `json:"path"`
				Owners []string `json:"owners"`
			}
			if err := json.Unmarshal(raw, &body); err != nil {
				return err
			}
			if len(body.Owners) == 0 {
				fmt.Println("<orphan>")
				os.Exit(1)
			}
			for _, o := range body.Owners {
				fmt.Println(o)
			}
			return nil
		}

		raw, err := c.FilesCoverage("")
		if err != nil {
			return err
		}
		var body struct {
			Coverage map[string][]string `json:"coverage"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			return err
		}
		paths := make([]string, 0, len(body.Coverage))
		for p := range body.Coverage {
			paths = append(paths, p)
		}
		sort.Strings(paths)
		for _, p := range paths {
			owners := body.Coverage[p]
			if len(owners) == 0 {
				fmt.Printf("%s\t<orphan>\n", p)
				continue
			}
			for _, o := range owners {
				fmt.Printf("%s\t%s\n", p, o)
			}
		}
		return nil
	},
}

func init() {
	filesCmd.AddCommand(filesOrphansCmd)
	filesCmd.AddCommand(filesCoverageCmd)
	rootCmd.AddCommand(filesCmd)
}
