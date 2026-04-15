package storage

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/feedloop/syde/internal/model"
)

// IndexStats tracks reindex results.
type IndexStats struct {
	Entities      int
	Relationships int
	Tags          int
	Words         int
}

// Reindex rebuilds the entire BadgerDB index from markdown files.
// It also recalculates the per-kind ID counter by scanning every
// entity's ID (form: "PFX-NNNN") and keeping the maximum observed
// number. This means deleting counters.yaml or dropping the index
// never corrupts the ID allocator — the next reindex restores it
// from the entity files themselves.
func Reindex(fs *FileStore, idx *Index) (*IndexStats, error) {
	if err := idx.DropAll(); err != nil {
		return nil, fmt.Errorf("drop index: %w", err)
	}

	stats := &IndexStats{}
	maxCounter := make(map[model.EntityKind]int)

	for _, kind := range model.AllEntityKinds() {
		files, err := fs.ListFiles(kind)
		if err != nil {
			continue
		}

		for _, filePath := range files {
			data, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			entity, body, err := UnmarshalAuto(data)
			if err != nil {
				continue
			}

			b := entity.GetBase()
			relPath := fs.RelativePath(kind, slugFromFile(filePath))
			lineMap := ComputeLineMap(data)

			ref := FileRef{
				File:  relPath,
				ID:    b.ID,
				Name:  b.Name,
				Kind:  b.Kind,
				Lines: lineMap,
			}

			if err := idx.IndexEntity(ref); err != nil {
				continue
			}
			stats.Entities++

			// Track highest counter for this kind, if the ID matches
			// the new "PFX-NNNN" form.
			if n, ok := parseCounterFromID(b.ID); ok {
				if n > maxCounter[kind] {
					maxCounter[kind] = n
				}
			}

			if len(b.Tags) > 0 {
				idx.IndexTags(ref, b.Tags)
				stats.Tags += len(b.Tags)
			}

			if len(b.Relationships) > 0 {
				idx.IndexRelationships(b.ID, relPath, b.Relationships)
				stats.Relationships += len(b.Relationships)
			}

			// Build word index from key text fields
			words := extractWords(b, body)
			if len(words) > 0 {
				idx.IndexWords(ref, words)
				stats.Words += len(words)
			}
		}
	}

	// Persist recalculated counters so subsequent NextID calls don't
	// collide with existing IDs.
	for kind, n := range maxCounter {
		if err := idx.SetCounter(kind, n); err != nil {
			return stats, fmt.Errorf("restore counter for %s: %w", kind, err)
		}
	}

	return stats, nil
}

// parseCounterFromID extracts the NNNN number from an ID of the form
// "PFX-NNNN". Returns (0, false) on any legacy format.
func parseCounterFromID(id string) (int, bool) {
	idx := strings.LastIndex(id, "-")
	if idx < 0 || idx == len(id)-1 {
		return 0, false
	}
	n, err := strconv.Atoi(id[idx+1:])
	if err != nil {
		return 0, false
	}
	return n, true
}

func extractWords(b *model.BaseEntity, body string) map[string]string {
	words := make(map[string]string)

	addWords := func(text, field string) {
		for _, token := range tokenize(text) {
			if _, exists := words[token]; !exists {
				words[token] = field
			}
		}
	}

	addWords(b.Name, "name")
	addWords(b.Description, "description")
	addWords(b.Purpose, "purpose")
	addWords(strings.Join(b.Notes, " "), "notes")

	// Body indexing: skip plan/task because their bodies are procedural
	// execution notes that churn every write and pollute search results
	// with transient implementation detail. Components, contracts,
	// concepts, decisions, flows, learnings, and systems hold the
	// substantive prose agents actually want to find.
	if b.Kind != model.KindPlan && b.Kind != model.KindTask {
		addWords(body, "body")
	}

	return words
}
