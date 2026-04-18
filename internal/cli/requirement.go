package cli

import (
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var requirementCmd = &cobra.Command{
	Use:   "requirement",
	Short: "Manage requirements",
}

var requirementVerifyCmd = &cobra.Command{
	Use:   "verify <slug>",
	Short: "Snapshot current SHA-256 of every file in each refining component",
	Long: `Updates the requirement's verified_against map with the current content
hash of each refining component's files. Run this after re-reading the
requirement and confirming it still holds against the current code —
it clears the requirement_stale finding for that requirement.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.GetByKind(model.KindRequirement, slug)
		if err != nil {
			return fmt.Errorf("requirement not found: %s", slug)
		}
		req := entity.(*model.RequirementEntity)

		if req.RequirementStatus != "" && req.RequirementStatus != model.RequirementActive {
			return fmt.Errorf("requirement %s is %s, not active — verify is only meaningful for active requirements", req.CanonicalSlug(), req.RequirementStatus)
		}

		projectRoot := filepath.Dir(store.FS.Root)

		// Index components once so we can resolve refines targets.
		all, err := store.ListAll()
		if err != nil {
			return fmt.Errorf("list entities: %w", err)
		}
		components := map[string]*model.ComponentEntity{}
		for _, ewb := range all {
			c, ok := ewb.Entity.(*model.ComponentEntity)
			if !ok {
				continue
			}
			b := c.GetBase()
			components[b.ID] = c
			components[b.CanonicalSlug()] = c
			components[utils.BaseSlug(b.CanonicalSlug())] = c
			components[utils.Slugify(b.Name)] = c
		}

		snapshots := map[string]model.VerifiedSnapshot{}
		var refinedComps []*model.ComponentEntity
		seen := map[string]bool{}
		for _, rel := range req.GetBase().Relationships {
			if rel.Type != model.RelRefines {
				continue
			}
			c, ok := components[rel.Target]
			if !ok {
				continue
			}
			slug := c.GetBase().CanonicalSlug()
			if seen[slug] {
				continue
			}
			seen[slug] = true
			refinedComps = append(refinedComps, c)
		}
		if len(refinedComps) == 0 {
			return fmt.Errorf("requirement %s has no refines:component edges — add at least one before verifying", req.CanonicalSlug())
		}

		now := time.Now().UTC().Format(time.RFC3339)
		for _, c := range refinedComps {
			b := c.GetBase()
			if len(b.Files) == 0 {
				fmt.Printf("  skip %s (no files mapped — design-phase component)\n", b.CanonicalSlug())
				continue
			}
			paths := make([]string, len(b.Files))
			for i, f := range b.Files {
				paths[i] = filepath.Join(projectRoot, f)
			}
			sort.Strings(paths)
			hash, err := utils.CombinedFilesSHA256(paths)
			if err != nil {
				return fmt.Errorf("hash component %s files: %w", b.CanonicalSlug(), err)
			}
			snapshots[b.CanonicalSlug()] = model.VerifiedSnapshot{Hash: hash, At: now}
			fmt.Printf("  snapshot %s → %s\n", b.CanonicalSlug(), hash[:12])
		}

		if req.VerifiedAgainst == nil {
			req.VerifiedAgainst = map[string]model.VerifiedSnapshot{}
		}
		for k, v := range snapshots {
			req.VerifiedAgainst[k] = v
		}
		if _, err := store.Update(req, body); err != nil {
			return fmt.Errorf("save requirement: %w", err)
		}
		fmt.Printf("Verified: %s (%d component snapshot(s))\n", req.Name, len(snapshots))
		return nil
	},
}

func init() {
	requirementCmd.AddCommand(requirementVerifyCmd)
	rootCmd.AddCommand(requirementCmd)
}
