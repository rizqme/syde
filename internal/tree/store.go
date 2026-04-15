package tree

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const treeFile = "tree.yaml"

// Path returns the path to the tree file inside a .syde/ directory.
func Path(sydeDir string) string {
	return filepath.Join(sydeDir, treeFile)
}

// Load reads .syde/tree.yaml. Returns an empty tree if the file doesn't
// exist yet (first scan).
func Load(sydeDir string) (*Tree, error) {
	data, err := os.ReadFile(Path(sydeDir))
	if err != nil {
		if os.IsNotExist(err) {
			return New(), nil
		}
		return nil, fmt.Errorf("read tree.yaml: %w", err)
	}
	var t Tree
	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("parse tree.yaml: %w", err)
	}
	if t.Nodes == nil {
		t.Nodes = make(map[string]*Node)
	}
	// Populate Path field (not persisted) so callers can range over Nodes
	// and still know each node's key.
	for path, n := range t.Nodes {
		if n != nil {
			n.Path = path
		}
	}
	// Ensure root exists
	if _, ok := t.Nodes["."]; !ok {
		t.Nodes["."] = &Node{Path: ".", Type: TypeDir, SummaryStale: true}
	}
	if t.Root == "" {
		t.Root = "."
	}
	return &t, nil
}

// Save writes the tree to .syde/tree.yaml atomically.
func Save(sydeDir string, t *Tree) error {
	if err := os.MkdirAll(sydeDir, 0755); err != nil {
		return err
	}
	data, err := yaml.Marshal(t)
	if err != nil {
		return fmt.Errorf("marshal tree.yaml: %w", err)
	}
	// Atomic replace
	target := Path(sydeDir)
	tmp := target + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmp, target)
}
