package audit

import (
	"fmt"
	"sort"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/tree"
)

// FileOwner identifies an entity that owns a file via its --file list.
type FileOwner struct {
	Kind      model.EntityKind
	Slug      string
	Name      string
	UpdatedAt string
}

func (o FileOwner) Label() string {
	return fmt.Sprintf("%s/%s", o.Kind, o.Name)
}

// buildFileOwners returns path → list of entities that reference it.
func buildFileOwners(all []model.EntityWithBody) map[string][]FileOwner {
	out := make(map[string][]FileOwner)
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		owner := FileOwner{
			Kind:      b.Kind,
			Slug:      b.CanonicalSlug(),
			Name:      b.Name,
			UpdatedAt: b.UpdatedAt,
		}
		for _, fp := range b.Files {
			out[fp] = append(out[fp], owner)
		}
	}
	return out
}

// FileCoverage returns path → list of owner entities for every non-ignored
// file in the tree (empty list when orphaned). Useful for `syde files
// coverage` to answer "who owns this file?".
func FileCoverage(all []model.EntityWithBody, t *tree.Tree) map[string][]FileOwner {
	owners := buildFileOwners(all)
	out := make(map[string][]FileOwner)
	for p, n := range t.Nodes {
		if n.Type != tree.TypeFile || n.Ignored {
			continue
		}
		out[p] = owners[p]
	}
	return out
}

// Orphans returns the list of non-ignored file paths with no owning
// entity, sorted.
func Orphans(all []model.EntityWithBody, t *tree.Tree) []string {
	owners := buildFileOwners(all)
	var out []string
	for p, n := range t.Nodes {
		if n.Type != tree.TypeFile || n.Ignored {
			continue
		}
		if len(owners[p]) == 0 {
			out = append(out, p)
		}
	}
	sort.Strings(out)
	return out
}

// orphanFindings reports non-ignored files with zero owners as errors.
func orphanFindings(all []model.EntityWithBody, t *tree.Tree) []Finding {
	var out []Finding
	for _, p := range Orphans(all, t) {
		out = append(out, Finding{
			Severity: SeverityError,
			Category: CatOrphanFile,
			Message:  fmt.Sprintf("file is not referenced by any entity. Add to an entity's --file list or run 'syde tree ignore %s' if it is not part of the design model.", p),
			Path:     p,
		})
	}
	return out
}

// fileRefFindings reports entity --file entries that don't correspond to
// any concrete node in the summary tree.
func fileRefFindings(all []model.EntityWithBody, t *tree.Tree) []Finding {
	var out []Finding
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		for _, fp := range b.Files {
			if t.Get(fp) == nil {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatFileNotInTree,
					Message:    fmt.Sprintf("file %q not in tree (run 'syde tree scan')", fp),
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Path:       fp,
				})
			}
		}
	}
	return out
}

// driftFindings warns when a file has a newer mtime than its owning
// entity's UpdatedAt — the code changed but the owning entity wasn't
// reviewed after the change.
func driftFindings(all []model.EntityWithBody, t *tree.Tree) []Finding {
	owners := buildFileOwners(all)
	var out []Finding
	for p, n := range t.Nodes {
		if n.Type != tree.TypeFile || n.Ignored || n.Mtime == "" {
			continue
		}
		var stale []string
		for _, o := range owners[p] {
			if o.UpdatedAt == "" || o.UpdatedAt < n.Mtime {
				stale = append(stale, o.Label())
			}
		}
		if len(stale) > 0 {
			out = append(out, Finding{
				Severity: SeverityWarning,
				Category: CatFileDrift,
				Message: fmt.Sprintf(
					"file changed at %s but these entities were not updated since (%s). Run 'syde task done' with --affected-entity/--affected-file set, or update the entities manually.",
					n.Mtime, strings.Join(stale, ", "),
				),
				Path: p,
			})
		}
	}
	return out
}
