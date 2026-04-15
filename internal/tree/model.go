// Package tree implements syde's file/folder summary tree with change
// tracking. It mirrors the source tree and lets Claude (or any other
// agent / human) attach short summaries to every file and folder. Summaries
// are marked stale automatically when file content changes, and that
// stale bit cascades up to ancestor folders so the whole tree stays in
// sync with reality.
package tree

// NodeType distinguishes files from directories.
type NodeType string

const (
	TypeFile NodeType = "file"
	TypeDir  NodeType = "dir"
)

// Node is a single entry in the summary tree. Nodes are stored flat in
// Tree.Nodes keyed by their path so mutations stay O(1) and deletion
// detection is easy. Tree-shaped rendering is handled by render.go.
type Node struct {
	Path   string   `yaml:"-" json:"path"` // key in Tree.Nodes; omitted from on-disk repr
	Type   NodeType `yaml:"type" json:"type"`
	Parent string   `yaml:"parent" json:"parent"`

	// Dir-only
	Children []string `yaml:"children,omitempty" json:"children,omitempty"`

	// File-only
	Size   int64  `yaml:"size,omitempty" json:"size,omitempty"`
	Hash   string `yaml:"hash,omitempty" json:"hash,omitempty"`
	Mtime  string `yaml:"mtime,omitempty" json:"mtime,omitempty"`
	Binary bool   `yaml:"binary,omitempty" json:"binary,omitempty"`

	// Summary
	Summary      string `yaml:"summary,omitempty" json:"summary,omitempty"`
	SummaryStale bool   `yaml:"summary_stale,omitempty" json:"summary_stale,omitempty"`
	UpdatedAt    string `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`

	// Ignored nodes are tracked in the tree (so scan still detects them)
	// but are exempt from both "needs summary" and "needs entity
	// reference" validation. Use for generated files, fixtures, large
	// assets, and anything that's part of the repo but not part of the
	// design model.
	Ignored bool `yaml:"ignored,omitempty" json:"ignored,omitempty"`
}

// Tree is the full summary tree for a project. It holds the node map
// plus a last-scan timestamp. Persisted as .syde/tree.yaml.
type Tree struct {
	ScannedAt string           `yaml:"scanned_at,omitempty" json:"scanned_at,omitempty"`
	Root      string           `yaml:"root" json:"root"`
	Nodes     map[string]*Node `yaml:"nodes" json:"nodes"`
}

// New returns an empty tree rooted at ".".
func New() *Tree {
	t := &Tree{
		Root:  ".",
		Nodes: make(map[string]*Node),
	}
	t.Nodes["."] = &Node{
		Path:         ".",
		Type:         TypeDir,
		Parent:       "",
		SummaryStale: true,
	}
	return t
}

// Get returns the node at path, or nil if not found.
func (t *Tree) Get(path string) *Node {
	if path == "" {
		path = "."
	}
	n, ok := t.Nodes[path]
	if !ok {
		return nil
	}
	return n
}

// Ancestors returns the chain of ancestor paths from the node's direct
// parent up to root. Useful for cascading stale bits and building
// breadcrumbs for the context bundle.
func (t *Tree) Ancestors(path string) []string {
	var out []string
	n := t.Get(path)
	if n == nil {
		return nil
	}
	p := n.Parent
	for p != "" {
		out = append(out, p)
		parent := t.Get(p)
		if parent == nil {
			break
		}
		p = parent.Parent
	}
	return out
}

// MarkAncestorsStale walks up from the given path and marks every
// ancestor (but not the node itself) as summary_stale=true.
func (t *Tree) MarkAncestorsStale(path string) {
	for _, a := range t.Ancestors(path) {
		if n := t.Get(a); n != nil {
			n.SummaryStale = true
		}
	}
}

// AddChild appends child to parent's children list if not already present,
// preserving lexical order.
func (t *Tree) AddChild(parentPath, childPath string) {
	p := t.Get(parentPath)
	if p == nil {
		return
	}
	for _, c := range p.Children {
		if c == childPath {
			return
		}
	}
	// Insert in sorted position
	inserted := false
	for i, c := range p.Children {
		if childPath < c {
			p.Children = append(p.Children[:i], append([]string{childPath}, p.Children[i:]...)...)
			inserted = true
			break
		}
	}
	if !inserted {
		p.Children = append(p.Children, childPath)
	}
}

// RemoveChild removes childPath from parent's children list.
func (t *Tree) RemoveChild(parentPath, childPath string) {
	p := t.Get(parentPath)
	if p == nil {
		return
	}
	kept := p.Children[:0]
	for _, c := range p.Children {
		if c != childPath {
			kept = append(kept, c)
		}
	}
	p.Children = kept
}

// Delete removes a node and unlinks it from its parent. Returns the
// parent path (for cascade) or "" if the node didn't exist.
func (t *Tree) Delete(path string) string {
	n := t.Get(path)
	if n == nil {
		return ""
	}
	parent := n.Parent
	t.RemoveChild(parent, path)
	delete(t.Nodes, path)
	return parent
}
