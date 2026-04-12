package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/query"
	"github.com/spf13/cobra"
)

var (
	queryKind      string
	queryTag       string
	queryStatus    string
	queryFull      bool
	queryRelatedTo string
	queryDependsOn string
	queryDependedBy string
	queryImpacts   string
	queryFlow      string
	queryFlowComps bool
	querySearch    string
	queryDiff      string
	queryDiffSince string
	queryFormat    string
)

var queryCmd = &cobra.Command{
	Use:   "query [slug]",
	Short: "Rich query system for targeted information",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()
		eng := query.NewEngine(store)

		// Entity lookup
		if len(args) > 0 {
			slug := args[0]
			if queryFull {
				r, err := eng.FullContext(slug)
				if err != nil {
					return err
				}
				if queryFormat == "json" {
					fmt.Println(query.FormatJSON(r))
				} else {
					fmt.Print(query.FormatRich(r))
				}
				return nil
			}
			r, err := eng.Lookup(slug)
			if err != nil {
				return err
			}
			if queryFormat == "json" {
				fmt.Println(query.FormatJSON(r))
			} else {
				fmt.Print(query.FormatRich(r))
			}
			return nil
		}

		// Impact analysis
		if queryImpacts != "" {
			r, err := eng.Impacts(queryImpacts, 3)
			if err != nil {
				return err
			}
			fmt.Print(query.FormatImpact(r))
			return nil
		}

		// Flow decomposition
		if queryFlow != "" && queryFlowComps {
			fd, err := eng.FlowComponents(queryFlow)
			if err != nil {
				return err
			}
			fmt.Print(query.FormatFlowDecomposition(fd))
			return nil
		}

		// Related to
		if queryRelatedTo != "" {
			results, err := eng.RelatedTo(queryRelatedTo)
			if err != nil {
				return err
			}
			fmt.Printf("Related to %s:\n\n", queryRelatedTo)
			printSummaries(results)
			return nil
		}

		// Depends on
		if queryDependsOn != "" {
			results, err := eng.DependsOn(queryDependsOn)
			if err != nil {
				return err
			}
			fmt.Printf("%s depends on:\n\n", queryDependsOn)
			printSummaries(results)
			return nil
		}

		// Depended by
		if queryDependedBy != "" {
			results, err := eng.DependedBy(queryDependedBy)
			if err != nil {
				return err
			}
			fmt.Printf("Depended by %s:\n\n", queryDependedBy)
			printSummaries(results)
			return nil
		}

		// Search
		if querySearch != "" {
			hits, err := eng.Search(querySearch)
			if err != nil {
				return err
			}
			if len(hits) == 0 {
				fmt.Printf("No results for '%s'\n", querySearch)
				return nil
			}
			fmt.Printf("Search results for '%s':\n\n", querySearch)
			for _, h := range hits {
				fmt.Printf("  %-12s %-25s %s\n", h.Kind, h.Name, h.File)
			}
			fmt.Printf("\n%d results\n", len(hits))
			return nil
		}

		// Diff
		if queryDiff != "" {
			entries, err := query.EntityDiff(store, queryDiff, queryDiffSince)
			if err != nil {
				return err
			}
			if len(entries) == 0 {
				fmt.Printf("No changes found for %s", queryDiff)
				if queryDiffSince != "" {
					fmt.Printf(" since %s", queryDiffSince)
				}
				fmt.Println()
				return nil
			}
			fmt.Printf("Changes to %s:\n\n", queryDiff)
			for _, e := range entries {
				fmt.Printf("  %s %s  %s\n", e.Date, e.Hash, e.Subject)
			}
			return nil
		}

		// Filter
		var kind model.EntityKind
		if queryKind != "" {
			k, ok := model.ParseEntityKind(queryKind)
			if !ok {
				return fmt.Errorf("unknown kind: %s", queryKind)
			}
			kind = k
		}
		results, err := eng.Filter(kind, queryTag, queryStatus)
		if err != nil {
			return err
		}
		if len(results) == 0 {
			fmt.Println("No entities found.")
			return nil
		}

		switch queryFormat {
		case "compact":
			fmt.Print(query.FormatCompact(results))
		case "refs":
			fmt.Print(query.FormatRefs(results))
		default:
			fmt.Print(query.FormatCompact(results))
		}
		fmt.Printf("\n%d entities\n", len(results))
		return nil
	},
}

func printSummaries(results []query.EntitySummary) {
	if len(results) == 0 {
		fmt.Println("  (none)")
		return
	}
	switch queryFormat {
	case "compact":
		fmt.Print(query.FormatCompact(results))
	case "refs":
		fmt.Print(query.FormatRefs(results))
	default:
		fmt.Print(query.FormatCompact(results))
	}
	fmt.Printf("\n%d entities\n", len(results))
}

func init() {
	queryCmd.Flags().StringVar(&queryKind, "kind", "", "filter by entity kind")
	queryCmd.Flags().StringVar(&queryTag, "tag", "", "filter by tag")
	queryCmd.Flags().StringVar(&queryStatus, "status", "", "filter by status")
	queryCmd.Flags().BoolVar(&queryFull, "full", false, "full context dump")
	queryCmd.Flags().StringVar(&queryRelatedTo, "related-to", "", "all entities related to slug")
	queryCmd.Flags().StringVar(&queryDependsOn, "depends-on", "", "what slug depends on")
	queryCmd.Flags().StringVar(&queryDependedBy, "depended-by", "", "what depends on slug")
	queryCmd.Flags().StringVar(&queryImpacts, "impacts", "", "transitive impact analysis")
	queryCmd.Flags().StringVar(&queryFlow, "flow", "", "flow slug for decomposition")
	queryCmd.Flags().BoolVar(&queryFlowComps, "components", false, "show flow components")
	queryCmd.Flags().StringVar(&querySearch, "search", "", "full-text search")
	queryCmd.Flags().StringVar(&queryDiff, "diff", "", "entity slug for git change history")
	queryCmd.Flags().StringVar(&queryDiffSince, "since", "", "git diff since (e.g. 7d, 2026-04-01)")
	queryCmd.Flags().StringVar(&queryFormat, "format", "rich", "output format (rich, json, compact, refs)")
	rootCmd.AddCommand(queryCmd)
}
