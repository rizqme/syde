package tree

import (
	"fmt"
	"path"
	"sort"
	"time"
)

// ScanResult summarizes what changed in the last scan.
type ScanResult struct {
	Added   int
	Changed int
	Deleted int
	Stale   int // total nodes with summary_stale=true after scan
}

func (r ScanResult) String() string {
	return fmt.Sprintf("added: %d, changed: %d, deleted: %d, stale: %d",
		r.Added, r.Changed, r.Deleted, r.Stale)
}

// Scan merges a fresh file walk into the existing tree. For each file:
//   - new → create node, mark stale, mark ancestors stale
//   - changed (hash mismatch) → update node, mark stale, mark ancestors stale
//   - unchanged → leave alone
//
// After processing walked files, any node still present from the previous
// scan that wasn't in the walk is deleted, and its parent is marked stale.
//
// Ancestor folder nodes are auto-created as needed. Binary and large
// files never get marked stale (they're auto-summarized by the walker).
func Scan(t *Tree, walked []WalkedFile) ScanResult {
	if t.Nodes == nil {
		t.Nodes = make(map[string]*Node)
	}
	t.ScannedAt = time.Now().UTC().Format(time.RFC3339)

	// Ensure root exists
	if t.Get(".") == nil {
		t.Nodes["."] = &Node{Path: ".", Type: TypeDir, SummaryStale: true}
	}

	result := ScanResult{}
	seen := make(map[string]bool, len(walked))
	seen["."] = true

	// Sort walked files for deterministic iteration (and deterministic
	// children lists on insert).
	sort.Slice(walked, func(i, j int) bool { return walked[i].Path < walked[j].Path })

	for _, w := range walked {
		// Ensure all parent directories exist
		ensureParents(t, w.Path, seen)
		seen[w.Path] = true

		existing := t.Get(w.Path)
		autoSummary, autoStaleFalse := autoSummarize(w)

		if existing == nil {
			// New file
			parent := parentPath(w.Path)
			n := &Node{
				Path:   w.Path,
				Type:   TypeFile,
				Parent: parent,
				Size:   w.Size,
				Hash:   w.Hash,
				Mtime:  w.Mtime,
				Binary: w.Binary,
			}
			if autoSummary != "" {
				n.Summary = autoSummary
				n.SummaryStale = !autoStaleFalse
			} else {
				n.SummaryStale = true
			}
			t.Nodes[w.Path] = n
			t.AddChild(parent, w.Path)
			t.MarkAncestorsStale(w.Path)
			result.Added++
			continue
		}

		// Existing file — check for changes. Preserve the Ignored flag
		// across rescans so a hand-marked ignore sticks.
		if existing.Hash != w.Hash {
			existing.Size = w.Size
			existing.Hash = w.Hash
			existing.Mtime = w.Mtime
			existing.Binary = w.Binary
			if existing.Ignored {
				// Ignored files don't need re-summarization, but they
				// should still track ancestor staleness so folder
				// summaries pick up structural changes.
			} else if autoSummary != "" {
				existing.Summary = autoSummary
				existing.SummaryStale = !autoStaleFalse
			} else {
				existing.SummaryStale = true
			}
			t.MarkAncestorsStale(w.Path)
			result.Changed++
		}
	}

	// Find deletions: any node present in the tree that wasn't seen in
	// the walk and isn't a directory we auto-created along the way.
	var toDelete []string
	for p, n := range t.Nodes {
		if n.Type == TypeFile && !seen[p] {
			toDelete = append(toDelete, p)
		}
	}
	for _, p := range toDelete {
		parent := t.Delete(p)
		if parent != "" {
			// Mark the immediate parent + all higher ancestors stale.
			if pn := t.Get(parent); pn != nil {
				pn.SummaryStale = true
			}
			t.MarkAncestorsStale(parent)
		}
		result.Deleted++
	}

	// Also prune empty directories that no longer have any children and
	// didn't exist in the previous tree's root listing. Keep root.
	pruneEmptyDirs(t)

	// Count stale
	for _, n := range t.Nodes {
		if n.SummaryStale {
			result.Stale++
		}
	}

	return result
}

// ensureParents walks from the root down to path's parent, creating any
// missing directory nodes along the way and marking newly created ones
// stale.
func ensureParents(t *Tree, p string, seen map[string]bool) {
	parent := parentPath(p)
	if parent == "." || parent == "" {
		seen["."] = true
		return
	}
	// Recurse upward first so ancestors exist before descendants
	ensureParents(t, parent, seen)
	seen[parent] = true

	if t.Get(parent) != nil {
		return
	}
	grand := parentPath(parent)
	n := &Node{
		Path:         parent,
		Type:         TypeDir,
		Parent:       grand,
		SummaryStale: true,
	}
	t.Nodes[parent] = n
	t.AddChild(grand, parent)
}

// parentPath returns the parent directory of p, using "." for top-level.
func parentPath(p string) string {
	dir := path.Dir(p)
	if dir == "" || dir == "." || dir == "/" {
		return "."
	}
	return dir
}

// autoSummarize returns a synthesized summary for binary / large files so
// they don't clutter the stale-changes list. Returns ("", false) for
// normal files that need a human-written summary.
func autoSummarize(w WalkedFile) (summary string, cleared bool) {
	if w.Binary {
		return fmt.Sprintf("<binary file, %d bytes>", w.Size), true
	}
	if w.Size > MaxSummarizableBytes {
		return fmt.Sprintf("<large file, %d bytes>", w.Size), true
	}
	return "", false
}

// pruneEmptyDirs removes dir nodes that have no children and aren't root.
// Repeats until no more changes (so cascading empty chains collapse).
func pruneEmptyDirs(t *Tree) {
	for {
		var toDelete []string
		for p, n := range t.Nodes {
			if p == "." {
				continue
			}
			if n.Type == TypeDir && len(n.Children) == 0 {
				toDelete = append(toDelete, p)
			}
		}
		if len(toDelete) == 0 {
			return
		}
		for _, p := range toDelete {
			parent := t.Delete(p)
			if parent != "" {
				if pn := t.Get(parent); pn != nil {
					pn.SummaryStale = true
				}
			}
		}
	}
}
