package tree

import (
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

// DefaultIgnore is the built-in ignore list applied to every project
// regardless of .gitignore or syde.yaml config. These are paths that
// should never carry summaries (binary indexes, dep caches, build
// artefacts, editor junk, syde's own data).
var DefaultIgnore = []string{
	".git/",
	".syde/",
	".claude/",
	".codex/",
	".agents/",
	"AGENTS.md",
	"CLAUDE.md",
	"node_modules/",
	"vendor/",
	"dist/",
	"build/",
	"__pycache__/",
	".venv/",
	".pytest_cache/",
	".next/",
	".turbo/",
	"target/",
	"out/",
	"coverage/",
	".DS_Store",
	"*.lock",
	"*.sum",
	"*.pyc",
	"*.o",
	"*.class",
}

// Matcher decides whether a path should be ignored. It combines:
//  1. The DefaultIgnore list (always applied).
//  2. The project's .gitignore file at root (if present).
//  3. Extra patterns from syde.yaml's tree_ignore field.
type Matcher struct {
	defaults  *ignore.GitIgnore
	gitignore *ignore.GitIgnore
	extra     *ignore.GitIgnore
}

// NewMatcher builds a Matcher for the given project root and extra
// patterns from syde.yaml.
func NewMatcher(root string, extra []string) *Matcher {
	m := &Matcher{
		defaults: ignore.CompileIgnoreLines(DefaultIgnore...),
	}

	// Parse .gitignore if it exists
	gitignorePath := filepath.Join(root, ".gitignore")
	if data, err := os.ReadFile(gitignorePath); err == nil {
		lines := strings.Split(string(data), "\n")
		m.gitignore = ignore.CompileIgnoreLines(lines...)
	}

	if len(extra) > 0 {
		m.extra = ignore.CompileIgnoreLines(extra...)
	}

	return m
}

// Match reports whether a path (relative to the project root, forward
// slashes, no leading dot) should be ignored.
func (m *Matcher) Match(relPath string, isDir bool) bool {
	if relPath == "" || relPath == "." {
		return false
	}
	// go-gitignore expects forward slashes
	p := filepath.ToSlash(relPath)
	if isDir && !strings.HasSuffix(p, "/") {
		p += "/"
	}

	if m.defaults != nil && m.defaults.MatchesPath(p) {
		return true
	}
	if m.gitignore != nil && m.gitignore.MatchesPath(p) {
		return true
	}
	if m.extra != nil && m.extra.MatchesPath(p) {
		return true
	}
	return false
}
