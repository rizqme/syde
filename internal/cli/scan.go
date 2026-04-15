package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/scan"
	"github.com/feedloop/syde/internal/tree"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var (
	syncDryRun   bool
	syncCoverage bool
	syncCheck    bool
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync design model with codebase",
	Long:  "Analyzes the codebase structure and checks syde entities against the implementation. For new projects, drives agent-powered extraction. For existing models, verifies coverage and detects drift.",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		if syncCoverage {
			guide, err := scan.LoadGuide(store.FS.Root)
			if err != nil {
				return fmt.Errorf("no sync guide found — run 'syde sync' first")
			}
			entries, err := scan.CheckCoverage(guide, store)
			if err != nil {
				return err
			}
			fmt.Print(scan.FormatCoverage(entries))
			return nil
		}

		if syncCheck {
			return runCompletenessCheck(store)
		}

		// Refresh the summary tree BEFORE producing the sync guide so
		// entity extraction (Round 1+) can rely on the tree being up-to-
		// date. The tree is the cheap way for agents to understand the
		// codebase without re-reading every file, and it's required
		// before syncing so entities pick up current file summaries and
		// architectural framing.
		sydeDir := store.FS.Root
		projectRoot := store.FS.Root + "/.."
		fmt.Println("Refreshing summary tree (syde tree scan)...")
		treeData, err := tree.Load(sydeDir)
		if err != nil {
			return fmt.Errorf("load tree: %w", err)
		}
		matcher := tree.NewMatcher(projectRoot, nil)
		walked, err := tree.WalkProject(projectRoot, matcher)
		if err != nil {
			return fmt.Errorf("walk project: %w", err)
		}
		treeResult := tree.Scan(treeData, walked)
		if err := tree.Save(sydeDir, treeData); err != nil {
			return fmt.Errorf("save tree: %w", err)
		}
		fmt.Printf("  tree: %s\n", treeResult.String())
		stalePaths := tree.StalePaths(treeData, false)
		if len(stalePaths) > 0 {
			fmt.Printf("  NOTE: %d paths need summaries. Run:\n", len(stalePaths))
			fmt.Println("    syde tree changes --leaves-only")
			fmt.Println("    syde tree summarize <path> --summary \"...\"")
			fmt.Println("  ...iterating until `syde tree status --strict` passes,")
			fmt.Println("  then use `syde tree context <path>` when creating entities.")
			fmt.Println()
		}

		guide, err := scan.GenerateGuide(store.FS.Root + "/..")
		if err != nil {
			return fmt.Errorf("generate sync guide: %w", err)
		}

		if err := scan.SaveGuide(store.FS.Root, guide); err != nil {
			return fmt.Errorf("save sync guide: %w", err)
		}

		fmt.Printf("Sync guide generated:\n")
		fmt.Printf("  Files: %d\n", guide.FileCount)
		fmt.Printf("  Languages: %d\n", len(guide.Languages))
		fmt.Printf("  Directories: %d\n", len(guide.DirectoryMap))
		fmt.Printf("  Saved: .syde/scan-guide.json\n")

		if syncDryRun {
			fmt.Println("\nDry run — no entities created.")
			fmt.Println("Run 'syde sync' without --dry-run to start.")
			return nil
		}

		// Check if there are existing entities
		total := 0
		designKinds := []model.EntityKind{
			model.KindSystem, model.KindComponent, model.KindContract,
			model.KindConcept, model.KindFlow, model.KindDecision,
		}
		for _, kind := range designKinds {
			entities, _ := store.List(kind)
			total += len(entities)
		}

		if total > 0 {
			fmt.Printf("\nExisting model found: %d entities\n", total)
			fmt.Println("\nRun 'syde sync --coverage' for directory coverage report.")
			fmt.Println("Run 'syde sync --check' for completeness audit.")
			fmt.Println("Use the syde skill in Claude Code to run full agent-driven sync.")
		} else {
			fmt.Println("\nNo existing entities. Use the syde skill in Claude Code to run 5-round agent extraction:")
			fmt.Println("  Round 1: System + Components")
			fmt.Println("  Round 2: Contracts + Concepts")
			fmt.Println("  Round 3: Flows")
			fmt.Println("  Round 4: Decisions")
			fmt.Println("  Round 5: Relationship wiring + validation")
		}

		return nil
	},
}

func runCompletenessCheck(store interface {
	List(model.EntityKind) ([]model.EntityWithBody, error)
}) error {
	fmt.Println("Completeness Audit")
	fmt.Println()

	gaps := 0

	// Collect entities by kind
	components, _ := store.List(model.KindComponent)
	contracts, _ := store.List(model.KindContract)
	concepts, _ := store.List(model.KindConcept)
	flows, _ := store.List(model.KindFlow)
	decisions, _ := store.List(model.KindDecision)
	systems, _ := store.List(model.KindSystem)

	// 1. System entity
	fmt.Println("1. System Entity")
	if len(systems) == 0 {
		fmt.Println("  ✗ No system entity — create one with: syde add system \"<name>\"")
		gaps++
	} else {
		fmt.Printf("  ✓ %s\n", systems[0].Entity.GetBase().Name)
	}

	// 2. Components without file references
	fmt.Println("\n2. Components without file references")
	for _, ewb := range components {
		b := ewb.Entity.GetBase()
		if len(b.Files) == 0 {
			fmt.Printf("  ✗ %s — no files field. Add with: syde update %s --file \"path/*.go\"\n", b.Name, utils.Slugify(b.Name))
			gaps++
		}
	}
	if gaps == 0 || func() bool {
		for _, ewb := range components {
			if len(ewb.Entity.GetBase().Files) == 0 {
				return false
			}
		}
		return true
	}() {
		fmt.Println("  ✓ All components have file references")
	}

	// 3. Components without contracts (exposed APIs)
	fmt.Println("\n3. Components without contracts")
	componentSlugs := map[string]string{}
	for _, ewb := range components {
		b := ewb.Entity.GetBase()
		componentSlugs[utils.Slugify(b.Name)] = b.Name
	}
	// Find which components have "exposes" relationships to contracts
	componentsWithContracts := map[string]bool{}
	for _, ewb := range components {
		b := ewb.Entity.GetBase()
		for _, rel := range b.Relationships {
			if rel.Type == "exposes" {
				componentsWithContracts[utils.Slugify(b.Name)] = true
			}
		}
	}
	compMissing := 0
	for slug, name := range componentSlugs {
		if !componentsWithContracts[slug] {
			fmt.Printf("  ? %s — no exposed contract. Does it have a public API?\n", name)
			compMissing++
		}
	}
	if compMissing == 0 {
		fmt.Println("  ✓ All components expose at least one contract")
	}

	// 4. Components without relationships
	fmt.Println("\n4. Entities without relationships")
	noRels := 0
	allKinds := []model.EntityKind{
		model.KindComponent, model.KindContract, model.KindConcept,
		model.KindFlow, model.KindDecision,
	}
	for _, kind := range allKinds {
		entities, _ := store.List(kind)
		for _, ewb := range entities {
			b := ewb.Entity.GetBase()
			if len(b.Relationships) == 0 {
				fmt.Printf("  ✗ %s/%s — no relationships\n", b.Kind, b.Name)
				noRels++
				gaps++
			}
		}
	}
	if noRels == 0 {
		fmt.Println("  ✓ All entities have relationships")
	}

	// 5. Entity kind coverage
	fmt.Println("\n5. Entity kind coverage")
	kindCounts := map[string]int{
		"system":    len(systems),
		"component": len(components),
		"contract":  len(contracts),
		"concept":   len(concepts),
		"flow":      len(flows),
		"decision":  len(decisions),
	}
	for kind, count := range kindCounts {
		if count == 0 {
			fmt.Printf("  ✗ No %s entities\n", kind)
			gaps++
		} else {
			fmt.Printf("  ✓ %s: %d\n", kind, count)
		}
	}

	// 6. Descriptions quality check
	fmt.Println("\n6. Short descriptions (< 20 chars)")
	shortDescs := 0
	for _, kind := range allKinds {
		entities, _ := store.List(kind)
		for _, ewb := range entities {
			b := ewb.Entity.GetBase()
			if len(b.Description) < 20 {
				fmt.Printf("  ? %s/%s — description too short: \"%s\"\n", b.Kind, b.Name, b.Description)
				shortDescs++
			}
		}
	}
	if shortDescs == 0 {
		fmt.Println("  ✓ All entities have substantial descriptions")
	}

	// 7. Component-specific field checks
	fmt.Println("\n7. Components missing key fields")
	compFieldGaps := 0
	for _, ewb := range components {
		c := ewb.Entity.(*model.ComponentEntity)
		missing := []string{}
		if c.Responsibility == "" {
			missing = append(missing, "responsibility")
		}
		if c.Boundaries == "" {
			missing = append(missing, "boundaries")
		}
		if len(missing) > 0 {
			fmt.Printf("  ✗ %s — missing: %v\n", c.Name, missing)
			compFieldGaps++
			gaps++
		}
	}
	if compFieldGaps == 0 {
		fmt.Println("  ✓ All components have responsibility and boundaries")
	}

	// Summary
	total := len(systems) + len(components) + len(contracts) + len(concepts) + len(flows) + len(decisions)
	fmt.Printf("\n═══ Summary ═══\n")
	fmt.Printf("Entities: %d | Gaps: %d\n", total, gaps)
	if gaps == 0 {
		fmt.Println("✓ Model is comprehensive")
	} else {
		fmt.Printf("Fix %d gaps to complete the model.\n", gaps)
	}

	return nil
}

func init() {
	syncCmd.Flags().BoolVar(&syncDryRun, "dry-run", false, "generate sync guide only")
	syncCmd.Flags().BoolVar(&syncCoverage, "coverage", false, "show coverage of codebase by components")
	syncCmd.Flags().BoolVar(&syncCheck, "check", false, "run completeness audit on existing model")
	rootCmd.AddCommand(syncCmd)
}
