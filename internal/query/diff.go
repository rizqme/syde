package query

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/utils"
)

// ChangeEntry represents a git-backed change to an entity file.
type ChangeEntry struct {
	Date    string `json:"date"`
	Hash    string `json:"hash"`
	Subject string `json:"subject"`
	File    string `json:"file"`
}

// EntityDiff returns recent git changes to an entity file.
func EntityDiff(store *storage.Store, slug, since string) ([]ChangeEntry, error) {
	entity, _, err := store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()
	relPath := store.FS.RelativePath(b.Kind, utils.Slugify(b.Name))
	absPath := filepath.Join(store.FS.Root, relPath)

	// Find the git root
	projectRoot := filepath.Dir(store.FS.Root)

	args := []string{"log", "--format=%H|%ad|%s", "--date=short"}
	if since != "" {
		args = append(args, "--since="+since)
	}
	args = append(args, "--", absPath)

	cmd := exec.Command("git", args...)
	cmd.Dir = projectRoot
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git log: %w", err)
	}

	var entries []ChangeEntry
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 3)
		if len(parts) < 3 {
			continue
		}
		entries = append(entries, ChangeEntry{
			Hash:    parts[0][:8],
			Date:    parts[1],
			Subject: parts[2],
			File:    relPath,
		})
	}
	return entries, nil
}
