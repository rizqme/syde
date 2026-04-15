package tree

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DefaultContextMaxBytes is the default cap on embedded file content in
// a ContextBundle. Matches the plan's 64 KiB default.
const DefaultContextMaxBytes = 64 * 1024

// BreadcrumbEntry is one link in the root → parent trail for a node.
type BreadcrumbEntry struct {
	Path    string `json:"path"`
	Summary string `json:"summary"`
	Stale   bool   `json:"stale"`
}

// ChildEntry is a summary-only listing of a folder's direct children.
type ChildEntry struct {
	Path    string `json:"path"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Stale   bool   `json:"stale"`
}

// ContextBundle is everything an agent needs to reason about a file or
// folder in one call: the ancestor breadcrumb, the node's own summary,
// and (for files) the inlined content.
type ContextBundle struct {
	Path       string            `json:"path"`
	Type       string            `json:"type"`
	Breadcrumb []BreadcrumbEntry `json:"breadcrumb"`
	Summary    string            `json:"summary"`
	Stale      bool               `json:"stale"`

	// File-only
	Size       int64  `json:"size,omitempty"`
	Hash       string `json:"hash,omitempty"`
	Mtime      string `json:"mtime,omitempty"`
	Binary     bool   `json:"binary,omitempty"`
	Content    string `json:"content,omitempty"`
	Truncated  bool   `json:"truncated,omitempty"`
	TotalBytes int64  `json:"total_bytes,omitempty"`

	// Dir-only
	Children []ChildEntry `json:"children,omitempty"`
}

// ContextOptions tunes how content is loaded.
type ContextOptions struct {
	IncludeContent bool  // default true
	MaxBytes       int64 // <=0 means DefaultContextMaxBytes
	ProjectRoot    string
}

// BuildContext assembles a ContextBundle for a given path.
func BuildContext(t *Tree, path string, opts ContextOptions) (*ContextBundle, error) {
	n := t.Get(path)
	if n == nil {
		return nil, fmt.Errorf("path not found in tree: %s", path)
	}

	bundle := &ContextBundle{
		Path:    n.Path,
		Type:    string(n.Type),
		Summary: n.Summary,
		Stale:   n.SummaryStale,
	}

	// Build breadcrumb root → direct-parent (ancestors returns
	// parent-first, so we reverse to root-first).
	ancestors := t.Ancestors(path)
	for i := len(ancestors) - 1; i >= 0; i-- {
		an := t.Get(ancestors[i])
		if an == nil {
			continue
		}
		bundle.Breadcrumb = append(bundle.Breadcrumb, BreadcrumbEntry{
			Path:    an.Path,
			Summary: an.Summary,
			Stale:   an.SummaryStale,
		})
	}

	switch n.Type {
	case TypeFile:
		bundle.Size = n.Size
		bundle.Hash = n.Hash
		bundle.Mtime = n.Mtime
		bundle.Binary = n.Binary

		if opts.IncludeContent && !n.Binary && opts.ProjectRoot != "" {
			maxBytes := opts.MaxBytes
			if maxBytes <= 0 {
				maxBytes = DefaultContextMaxBytes
			}
			content, total, truncated, err := readFileCapped(filepath.Join(opts.ProjectRoot, n.Path), maxBytes)
			if err == nil {
				bundle.Content = string(content)
				bundle.TotalBytes = total
				bundle.Truncated = truncated
			}
		}

	case TypeDir:
		for _, c := range n.Children {
			child := t.Get(c)
			if child == nil {
				continue
			}
			bundle.Children = append(bundle.Children, ChildEntry{
				Path:    child.Path,
				Type:    string(child.Type),
				Summary: child.Summary,
				Stale:   child.SummaryStale,
			})
		}
	}

	return bundle, nil
}

// readFileCapped reads up to maxBytes from a file, returning the bytes,
// the total file size, and whether truncation occurred.
func readFileCapped(path string, maxBytes int64) ([]byte, int64, bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, 0, false, err
	}
	total := info.Size()

	f, err := os.Open(path)
	if err != nil {
		return nil, total, false, err
	}
	defer f.Close()

	lr := io.LimitReader(f, maxBytes)
	data, err := io.ReadAll(lr)
	if err != nil {
		return nil, total, false, err
	}
	return data, total, total > maxBytes, nil
}
