package tree

import (
	"fmt"
	"strings"
)

// RenderOptions controls Render output.
type RenderOptions struct {
	Root         string // starting path (default ".")
	MaxDepth     int    // -1 = unlimited, 0 = just the root
	WithSummary  bool   // include summary text inline
	ShowStale    bool   // prefix stale nodes with "!"
}

// Render produces an ASCII tree of the subtree rooted at opts.Root.
func Render(t *Tree, opts RenderOptions) string {
	root := opts.Root
	if root == "" {
		root = "."
	}
	var b strings.Builder
	renderNode(&b, t, root, "", true, 0, opts)
	return b.String()
}

func renderNode(b *strings.Builder, t *Tree, path, prefix string, isLast bool, depth int, opts RenderOptions) {
	n := t.Get(path)
	if n == nil {
		return
	}

	// Print this node
	connector := "├── "
	childPrefix := prefix + "│   "
	if isLast {
		connector = "└── "
		childPrefix = prefix + "    "
	}
	if depth == 0 {
		connector = ""
		childPrefix = ""
	}

	staleMark := ""
	if opts.ShowStale && n.SummaryStale {
		staleMark = "! "
	}
	label := displayName(path)
	if n.Type == TypeDir {
		label += "/"
	}
	b.WriteString(prefix)
	b.WriteString(connector)
	b.WriteString(staleMark)
	b.WriteString(label)

	if opts.WithSummary && n.Summary != "" {
		// Wrap long summaries
		summary := truncate(n.Summary, 80)
		fmt.Fprintf(b, "  — %s", summary)
	}
	b.WriteString("\n")

	// Recurse
	if opts.MaxDepth >= 0 && depth >= opts.MaxDepth {
		return
	}
	if n.Type != TypeDir {
		return
	}
	for i, c := range n.Children {
		last := i == len(n.Children)-1
		renderNode(b, t, c, childPrefix, last, depth+1, opts)
	}
}

func displayName(path string) string {
	if path == "." {
		return "."
	}
	if idx := strings.LastIndex(path, "/"); idx >= 0 {
		return path[idx+1:]
	}
	return path
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
