package cli

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	queryKind       string
	queryTag        string
	queryFull       bool
	queryRelatedTo  string
	queryDependsOn  string
	queryDependedBy string
	queryRefinedBy  string
	queryImpacts    string
	queryFlow       string
	queryFlowComps  bool
	querySearch     string
	queryFile       string
	queryCode       string
	queryContent    bool
	queryLimit      int
	queryAny        bool
	queryNoRelated  bool
	queryFormat     string
)

// queryCmd is a thin wrapper over syded's /api/<project>/query
// endpoint. The heavy lifting — graph walks, search, formatting — all
// happens server-side so the CLI never opens BadgerDB. We just pick
// the right mode + format and print what comes back.
var queryCmd = &cobra.Command{
	Use:   "query [slug]",
	Short: "Rich query system for targeted information",
	Long: `The single entry point for understanding any file, symbol, or entity.

syde query is the unified context surface where architecture (the entity
index) and code (the source files) meet. Every query also reveals whether
architecture and code are in sync — bypassing syde to grep or read raw
silently disconnects the two and lets the design model rot.

THE THREE-QUESTION CHECKLIST (run before any Grep / Read):

  1. What entity owns this?      → syde query --file <path>  /  syde query <slug>
  2. What does syde know about a term? → syde query --search "<term>"
  3. What code references it?    → syde query --code "<symbol>"

FIVE ACCESS PATHS:

  1. Entity lookup       — syde query <slug> [--full]
  2. Keyword search      — syde query --search "<terms>" [--kind K] [--tag T] [--limit N] [--any]
  3. Code search         — syde query --code "<pattern>" [--limit N]
  4. File → entities     — syde query --file <path> [--content] [--no-related]
  5. Graph walks         — --impacts, --related-to, --depends-on, --depended-by, --flow --components

Plus filter listings via --kind / --tag with no slug and no search term.

NOTES:

  - Search is AND by default; the engine auto-broadens to OR when AND
    yields zero and labels the resulting hits "broadened".
  - The tokenizer splits CamelCase and snake_case identifiers into
    sub-tokens, so 'ConceptEntity' indexes as concept + entity too.
  - --code uses ripgrep when available, falls back to a Go walker.
    Every code hit carries its owning entity (or a ⚠ orphan warning).
  - --file --content inlines the file body (capped at 100KB) alongside
    owners + related — replaces 'Read <tracked-file>' entirely.

DRIFT SIGNALS — act on these immediately:

  - --file reports '⚠ DRIFT: no owning entities' → the tracked file is
    an orphan. Map it (syde update <component> --file ...) or
    'syde tree ignore <path>' if it is not part of the design model.
  - --code reports '⚠ orphan' on any hit → same fix.

These are the model telling you architecture has rotted relative to code.

DO NOT use Grep or Read on files tracked by 'syde tree scan'. Use
'syde query --code' for symbol search and 'syde query --file --content'
for file reads. Reserve Grep / Read for vendor/, node_modules/, generated
assets, .git/, build artifacts, and binary blobs only.
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := openClient()
		if err != nil {
			return err
		}

		format := queryFormat
		if format == "" {
			format = "rich"
		}

		printBytes := func(b []byte) {
			if len(b) == 0 {
				return
			}
			fmt.Print(string(b))
			if b[len(b)-1] != '\n' {
				fmt.Println()
			}
		}

		// Entity lookup
		if len(args) > 0 {
			mode := "lookup"
			if queryFull {
				mode = "full"
			}
			raw, err := c.Query(mode, args[0], format, nil)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		}

		switch {
		case queryImpacts != "":
			raw, err := c.Query("impacts", queryImpacts, format, nil)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		case queryFlow != "" && queryFlowComps:
			raw, err := c.Query("flow-components", queryFlow, format, nil)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		case queryRelatedTo != "":
			raw, err := c.Query("related-to", queryRelatedTo, format, nil)
			if err != nil {
				return err
			}
			fmt.Printf("Related to %s:\n\n", queryRelatedTo)
			printBytes(raw)
			return nil
		case queryDependsOn != "":
			raw, err := c.Query("depends-on", queryDependsOn, format, nil)
			if err != nil {
				return err
			}
			fmt.Printf("%s depends on:\n\n", queryDependsOn)
			printBytes(raw)
			return nil
		case queryDependedBy != "":
			raw, err := c.Query("depended-by", queryDependedBy, format, nil)
			if err != nil {
				return err
			}
			fmt.Printf("Depended by %s:\n\n", queryDependedBy)
			printBytes(raw)
			return nil
		case queryRefinedBy != "":
			raw, err := c.Query("refined-by", queryRefinedBy, format, nil)
			if err != nil {
				return err
			}
			if format == "rich" {
				fmt.Printf("Active requirements refining %s:\n\n", queryRefinedBy)
			}
			printBytes(raw)
			return nil
		case querySearch != "":
			extra := url.Values{}
			extra.Set("q", querySearch)
			if queryKind != "" {
				extra.Set("kind", queryKind)
			}
			if queryTag != "" {
				extra.Set("tag", queryTag)
			}
			if queryAny {
				extra.Set("any", "true")
			}
			if queryLimit > 0 {
				extra.Set("limit", strconv.Itoa(queryLimit))
			}
			raw, err := c.Query("search", "", format, extra)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		case queryFile != "":
			extra := url.Values{}
			extra.Set("path", queryFile)
			if queryNoRelated {
				extra.Set("with_related", "false")
			}
			if queryContent {
				extra.Set("content", "true")
			}
			raw, err := c.Query("by-file", "", format, extra)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		case queryCode != "":
			extra := url.Values{}
			extra.Set("q", queryCode)
			if queryLimit > 0 {
				extra.Set("limit", strconv.Itoa(queryLimit))
			}
			raw, err := c.Query("code", "", format, extra)
			if err != nil {
				return err
			}
			printBytes(raw)
			return nil
		}

		// Default: filter by kind/tag
		extra := url.Values{}
		if queryKind != "" {
			extra.Set("kind", queryKind)
		}
		if queryTag != "" {
			extra.Set("tag", queryTag)
		}
		raw, err := c.Query("filter", "", format, extra)
		if err != nil {
			return err
		}
		printBytes(raw)
		return nil
	},
}

func init() {
	queryCmd.Flags().StringVar(&queryKind, "kind", "", "filter by entity kind")
	queryCmd.Flags().StringVar(&queryTag, "tag", "", "filter by tag")
	queryCmd.Flags().BoolVar(&queryFull, "full", false, "full context dump")
	queryCmd.Flags().StringVar(&queryRelatedTo, "related-to", "", "all entities related to slug")
	queryCmd.Flags().StringVar(&queryDependsOn, "depends-on", "", "what slug depends on")
	queryCmd.Flags().StringVar(&queryDependedBy, "depended-by", "", "what depends on slug")
	queryCmd.Flags().StringVar(&queryRefinedBy, "refined-by", "", "active requirements that refine this component slug")
	queryCmd.Flags().StringVar(&queryImpacts, "impacts", "", "transitive impact analysis")
	queryCmd.Flags().StringVar(&queryFlow, "flow", "", "flow slug for decomposition")
	queryCmd.Flags().BoolVar(&queryFlowComps, "components", false, "show flow components")
	queryCmd.Flags().StringVar(&querySearch, "search", "", "full-text search (honors --kind/--tag/--limit/--any)")
	queryCmd.Flags().StringVar(&queryFile, "file", "", "find entities owning a source file path (exact match, or directory prefix); includes one-hop related entities unless --no-related")
	queryCmd.Flags().StringVar(&queryCode, "code", "", "literal-string search across every tracked source file (uses ripgrep when available, Go fallback otherwise); each hit is annotated with its owning entity")
	queryCmd.Flags().BoolVar(&queryContent, "content", false, "--file: also inline the file content (capped at 100KB)")
	queryCmd.Flags().IntVar(&queryLimit, "limit", 0, "max results for search/code (0 = unbounded for search, 50 for code)")
	queryCmd.Flags().BoolVar(&queryAny, "any", false, "search: OR-merge tokens instead of the default AND")
	queryCmd.Flags().BoolVar(&queryNoRelated, "no-related", false, "--file: omit one-hop related entities")
	queryCmd.Flags().StringVar(&queryFormat, "format", "rich", "output format (rich, json, compact, refs)")
	rootCmd.AddCommand(queryCmd)
}
