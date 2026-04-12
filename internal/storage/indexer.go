package storage

import (
	"fmt"
	"os"
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
func Reindex(fs *FileStore, idx *Index) (*IndexStats, error) {
	if err := idx.DropAll(); err != nil {
		return nil, fmt.Errorf("drop index: %w", err)
	}

	stats := &IndexStats{}

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

			if len(b.Tags) > 0 {
				idx.IndexTags(ref, b.Tags)
				stats.Tags += len(b.Tags)
			}

			if b.Status != "" {
				idx.IndexStatus(ref, b.Status)
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

	return stats, nil
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
	addWords(body, "body")

	return words
}
