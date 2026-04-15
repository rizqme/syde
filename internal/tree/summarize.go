package tree

import (
	"fmt"
	"time"
)

// SetIgnored toggles the ignored flag on a node. When ignored becomes
// true, the node's stale bit is cleared so it stops showing up in
// `tree changes` / `tree status --strict`. When ignored becomes false,
// the node's stale bit is set so the node is re-summarized on the next
// pass.
func SetIgnored(t *Tree, path string, ignored bool) error {
	n := t.Get(path)
	if n == nil {
		return fmt.Errorf("path not found in tree: %s (run 'syde tree scan' first)", path)
	}
	n.Ignored = ignored
	if ignored {
		n.SummaryStale = false
	} else if n.Summary == "" {
		n.SummaryStale = true
	}
	return nil
}

// SetSummary updates a node's summary text, clears its stale bit, and
// marks the direct parent as stale so it bubbles up on the next pass.
// Returns an error if the path doesn't exist in the tree.
func SetSummary(t *Tree, path, summary string) error {
	n := t.Get(path)
	if n == nil {
		return fmt.Errorf("path not found in tree: %s (run 'syde tree scan' first)", path)
	}
	n.Summary = summary
	n.SummaryStale = false
	n.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	// Mark direct parent stale (cascade happens naturally when the
	// parent is summarized later).
	if n.Parent != "" {
		if p := t.Get(n.Parent); p != nil {
			p.SummaryStale = true
		}
	}
	return nil
}

// StalePaths returns all node paths with summary_stale=true. Ignored
// nodes are skipped entirely — they're tracked in the tree but don't
// participate in the summarize/validate lifecycle. If leavesOnly is
// true, stale folders whose descendants still contain stale files/
// folders are excluded (so callers fix leaves first). The result is
// sorted deepest-first so leaves come before their parents — the right
// iteration order for bottom-up summarization.
func StalePaths(t *Tree, leavesOnly bool) []string {
	var stale []string
	for p, n := range t.Nodes {
		if n.Ignored {
			continue
		}
		if n.SummaryStale {
			stale = append(stale, p)
		}
	}
	// Sort deepest-first, then lexically for determinism
	sortByDepthDesc(stale)

	if !leavesOnly {
		return stale
	}

	// Filter out folders whose descendants still have stale nodes.
	staleSet := make(map[string]bool, len(stale))
	for _, p := range stale {
		staleSet[p] = true
	}
	var out []string
	for _, p := range stale {
		n := t.Get(p)
		if n == nil {
			continue
		}
		if n.Type == TypeFile {
			out = append(out, p)
			continue
		}
		// Folder: check if any descendant is still stale
		if !anyDescendantStale(t, p, staleSet) {
			out = append(out, p)
		}
	}
	return out
}

func anyDescendantStale(t *Tree, path string, staleSet map[string]bool) bool {
	n := t.Get(path)
	if n == nil {
		return false
	}
	for _, c := range n.Children {
		if staleSet[c] && c != path {
			return true
		}
		if anyDescendantStale(t, c, staleSet) {
			return true
		}
	}
	return false
}

// sortByDepthDesc sorts paths so deeper paths come first. Ties broken
// lexically so the output is deterministic.
func sortByDepthDesc(paths []string) {
	// Simple insertion sort is fine; typical slices are small and we
	// want stable, deterministic behavior without bringing in sort.
	for i := 1; i < len(paths); i++ {
		for j := i; j > 0 && less(paths[j], paths[j-1]); j-- {
			paths[j], paths[j-1] = paths[j-1], paths[j]
		}
	}
}

func less(a, b string) bool {
	da, db := depth(a), depth(b)
	if da != db {
		return da > db
	}
	return a < b
}

func depth(p string) int {
	if p == "." || p == "" {
		return 0
	}
	d := 1
	for _, c := range p {
		if c == '/' {
			d++
		}
	}
	return d
}
