package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/tree"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "File/folder summary tree with change tracking",
}

// projectRootFromSydeDir returns the project root given a .syde/ path.
func projectRootFromSydeDir(sydeDir string) string {
	return filepath.Dir(sydeDir)
}

// loadTreeMatcher builds a tree.Matcher from the project's syde.yaml
// (for TreeIgnore extras) and the project root (for .gitignore).
func loadTreeMatcher(sydeDir string) *tree.Matcher {
	root := projectRootFromSydeDir(sydeDir)
	var extra []string
	if cfg, err := config.Load(sydeDir); err == nil && cfg != nil {
		extra = cfg.TreeIgnore
	}
	return tree.NewMatcher(root, extra)
}

// --- scan ------------------------------------------------------------------

var treeScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Walk the project tree, diff against .syde/tree.yaml, mark stale files",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found (run 'syde init' first)")
		}
		root := projectRootFromSydeDir(dir)

		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		m := loadTreeMatcher(dir)
		walked, err := tree.WalkProject(root, m)
		if err != nil {
			return err
		}
		result := tree.Scan(t, walked)
		if err := tree.Save(dir, t); err != nil {
			return err
		}
		fmt.Println(result.String())
		return nil
	},
}

// --- changes ---------------------------------------------------------------

var (
	treeChangesFormat     string
	treeChangesLeavesOnly bool
)

var treeChangesCmd = &cobra.Command{
	Use:   "changes",
	Short: "List paths with stale summaries (deepest-first)",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		stale := tree.StalePaths(t, treeChangesLeavesOnly)
		if treeChangesFormat == "json" {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(map[string]interface{}{
				"stale": stale,
				"count": len(stale),
			})
		}
		if len(stale) == 0 {
			fmt.Println("clean")
			return nil
		}
		for _, p := range stale {
			n := t.Get(p)
			kind := "file"
			if n != nil && n.Type == tree.TypeDir {
				kind = "dir"
			}
			fmt.Printf("%-4s %s\n", kind, p)
		}
		return nil
	},
}

// --- summarize -------------------------------------------------------------

var treeSummarizeSummary string

var treeSummarizeCmd = &cobra.Command{
	Use:   "summarize <path>",
	Short: "Set the summary for a file or folder (cascades stale to parent)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		summary := treeSummarizeSummary
		if summary == "-" {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("read stdin: %w", err)
			}
			summary = string(data)
		}
		if summary == "" {
			return fmt.Errorf("--summary is required (use '-' to read from stdin)")
		}
		if err := tree.SetSummary(t, args[0], summary); err != nil {
			return err
		}
		if err := tree.Save(dir, t); err != nil {
			return err
		}
		n := t.Get(args[0])
		fmt.Printf("Summarized %s (parent %s marked stale)\n", args[0], n.Parent)
		return nil
	},
}

// --- show ------------------------------------------------------------------

var (
	treeShowFull     bool
	treeShowMaxDepth int
	treeShowStale    bool
)

var treeShowCmd = &cobra.Command{
	Use:   "show [path]",
	Short: "Pretty-print the summary tree",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		root := "."
		if len(args) > 0 {
			root = args[0]
		}
		maxDepth := treeShowMaxDepth
		if treeShowFull {
			maxDepth = -1
		}
		out := tree.Render(t, tree.RenderOptions{
			Root:        root,
			MaxDepth:    maxDepth,
			WithSummary: true,
			ShowStale:   treeShowStale,
		})
		fmt.Print(out)
		return nil
	},
}

// --- get -------------------------------------------------------------------

var treeGetCmd = &cobra.Command{
	Use:   "get <path>",
	Short: "Print just the summary for a node",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		n := t.Get(args[0])
		if n == nil {
			return fmt.Errorf("path not found in tree: %s", args[0])
		}
		fmt.Println(n.Summary)
		return nil
	},
}

// --- context ---------------------------------------------------------------

var (
	treeContextFormat    string
	treeContextNoContent bool
	treeContextMaxBytes  int64
)

var treeContextCmd = &cobra.Command{
	Use:   "context <path>",
	Short: "Bundle breadcrumb + summary + file content for a path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		root := projectRootFromSydeDir(dir)
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		bundle, err := tree.BuildContext(t, args[0], tree.ContextOptions{
			IncludeContent: !treeContextNoContent,
			MaxBytes:       treeContextMaxBytes,
			ProjectRoot:    root,
		})
		if err != nil {
			return err
		}
		if treeContextFormat == "json" {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(bundle)
		}
		return renderContextPlain(bundle)
	},
}

func renderContextPlain(b *tree.ContextBundle) error {
	fmt.Printf("PATH: %s (%s)\n", b.Path, b.Type)
	fmt.Println()
	if len(b.Breadcrumb) > 0 {
		fmt.Println("TREE SUMMARY (root → parent):")
		for _, e := range b.Breadcrumb {
			staleMark := ""
			if e.Stale {
				staleMark = " [stale]"
			}
			summary := e.Summary
			if summary == "" {
				summary = "(no summary)"
			}
			fmt.Printf("  %s — %s%s\n", e.Path, summary, staleMark)
		}
		fmt.Println()
	}
	nodeSummary := b.Summary
	if nodeSummary == "" {
		nodeSummary = "(no summary)"
	}
	staleMark := ""
	if b.Stale {
		staleMark = " [stale]"
	}
	if b.Type == string(tree.TypeFile) {
		fmt.Printf("FILE SUMMARY: %s%s\n", nodeSummary, staleMark)
		fmt.Println()
		if b.Binary {
			fmt.Printf("CONTENT: <binary file, %d bytes>\n", b.Size)
		} else if b.Content != "" {
			label := fmt.Sprintf("CONTENT (%d bytes)", b.TotalBytes)
			if b.Truncated {
				label = fmt.Sprintf("CONTENT (%d bytes, truncated to first %d)", b.TotalBytes, len(b.Content))
			}
			fmt.Printf("%s:\n", label)
			fmt.Println(b.Content)
		}
	} else {
		fmt.Printf("FOLDER SUMMARY: %s%s\n", nodeSummary, staleMark)
		if len(b.Children) > 0 {
			fmt.Println()
			fmt.Printf("CHILDREN (%d):\n", len(b.Children))
			for _, c := range b.Children {
				cs := c.Summary
				if cs == "" {
					cs = "(no summary)"
				}
				cStale := ""
				if c.Stale {
					cStale = " [stale]"
				}
				fmt.Printf("  %-4s %s — %s%s\n", c.Type, c.Path, cs, cStale)
			}
		}
	}
	return nil
}

// --- ignore / unignore -----------------------------------------------------

var treeIgnoreCmd = &cobra.Command{
	Use:   "ignore <path>",
	Short: "Mark a tree node as ignored (exempt from summary + entity-ref checks)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		if err := tree.SetIgnored(t, args[0], true); err != nil {
			return err
		}
		if err := tree.Save(dir, t); err != nil {
			return err
		}
		fmt.Printf("Ignored: %s\n", args[0])
		return nil
	},
}

var treeUnignoreCmd = &cobra.Command{
	Use:   "unignore <path>",
	Short: "Remove the ignored flag from a tree node",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		if err := tree.SetIgnored(t, args[0], false); err != nil {
			return err
		}
		if err := tree.Save(dir, t); err != nil {
			return err
		}
		fmt.Printf("Unignored: %s (marked stale — summarize it)\n", args[0])
		return nil
	},
}

// --- status ----------------------------------------------------------------

var treeStatusStrict bool

var treeStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Summary tree status (counts + last scan time)",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := resolveSydeDir()
		if dir == "" {
			return fmt.Errorf("no .syde/ directory found")
		}
		t, err := tree.Load(dir)
		if err != nil {
			return err
		}
		var files, dirs, staleFiles, staleDirs int
		for _, n := range t.Nodes {
			if n.Type == tree.TypeFile {
				files++
				if n.SummaryStale {
					staleFiles++
				}
			} else {
				dirs++
				if n.SummaryStale {
					staleDirs++
				}
			}
		}
		fmt.Printf("Last scan: %s\n", nonEmpty(t.ScannedAt, "(never)"))
		fmt.Printf("Files:     %d (stale: %d)\n", files, staleFiles)
		fmt.Printf("Dirs:      %d (stale: %d)\n", dirs, staleDirs)
		fmt.Printf("Total:     %d (stale: %d)\n", files+dirs, staleFiles+staleDirs)

		if treeStatusStrict && (staleFiles+staleDirs) > 0 {
			return fmt.Errorf("tree has %d stale entries — run 'syde tree changes' and re-summarize", staleFiles+staleDirs)
		}
		return nil
	},
}

func nonEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

func init() {
	treeChangesCmd.Flags().StringVar(&treeChangesFormat, "format", "plain", "output format: plain|json")
	treeChangesCmd.Flags().BoolVar(&treeChangesLeavesOnly, "leaves-only", false, "only list stale files and folders whose stale descendants are all resolved")

	treeSummarizeCmd.Flags().StringVar(&treeSummarizeSummary, "summary", "", "summary text (use '-' to read from stdin)")
	_ = treeSummarizeCmd.MarkFlagRequired("summary")

	treeShowCmd.Flags().BoolVar(&treeShowFull, "full", false, "show the entire tree (overrides --max-depth)")
	treeShowCmd.Flags().IntVar(&treeShowMaxDepth, "max-depth", 2, "maximum tree depth to render")
	treeShowCmd.Flags().BoolVar(&treeShowStale, "stale", false, "prefix stale entries with '!'")

	treeContextCmd.Flags().StringVar(&treeContextFormat, "format", "plain", "output format: plain|json")
	treeContextCmd.Flags().BoolVar(&treeContextNoContent, "no-content", false, "omit inlined file content")
	treeContextCmd.Flags().Int64Var(&treeContextMaxBytes, "max-bytes", tree.DefaultContextMaxBytes, "cap on inlined file content bytes")

	treeStatusCmd.Flags().BoolVar(&treeStatusStrict, "strict", false, "exit 1 if any files or folders are stale")

	treeCmd.AddCommand(treeScanCmd, treeChangesCmd, treeSummarizeCmd, treeShowCmd, treeGetCmd, treeContextCmd, treeStatusCmd, treeIgnoreCmd, treeUnignoreCmd)
	rootCmd.AddCommand(treeCmd)
}
