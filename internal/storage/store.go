package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// Store provides unified access to the file store and index.
type Store struct {
	FS  *FileStore
	Idx *Index
}

// NewStore creates a Store with file store and BadgerDB index.
func NewStore(sydeDir string) (*Store, error) {
	fs := NewFileStore(sydeDir)
	idx, err := OpenIndex(filepath.Join(sydeDir, "index"))
	if err != nil {
		return nil, err
	}
	return &Store{FS: fs, Idx: idx}, nil
}

// Close closes the index.
func (s *Store) Close() error {
	return s.Idx.Close()
}

// Create saves a new entity and indexes it.
func (s *Store) Create(e model.Entity, body string) (string, error) {
	b := e.GetBase()
	if b.ID == "" {
		b.ID = utils.GenerateID(b.Kind)
	}
	if b.Status == "" {
		b.Status = model.StatusDraft
	}

	filePath, err := s.FS.Save(e, body)
	if err != nil {
		return "", err
	}

	if err := s.indexEntity(e, body); err != nil {
		return filePath, fmt.Errorf("index after create: %w", err)
	}

	return filePath, nil
}

// Update saves an existing entity and reindexes it.
func (s *Store) Update(e model.Entity, body string) (string, error) {
	b := e.GetBase()

	// Remove old index entries
	s.Idx.RemoveEntity(b.Kind, b.ID)

	filePath, err := s.FS.Save(e, body)
	if err != nil {
		return "", err
	}

	if err := s.indexEntity(e, body); err != nil {
		return filePath, fmt.Errorf("index after update: %w", err)
	}

	return filePath, nil
}

// Delete removes an entity file and its index entries.
func (s *Store) Delete(kind model.EntityKind, slug string) error {
	// Load to get the ID first
	e, _, err := s.FS.Load(kind, slug)
	if err != nil {
		return err
	}

	s.Idx.RemoveEntity(kind, e.GetBase().ID)
	return s.FS.Delete(kind, slug)
}

// Get loads an entity by slug, trying index first for kind resolution.
func (s *Store) Get(slug string) (model.Entity, string, error) {
	// Try each kind
	for _, kind := range model.AllEntityKinds() {
		if s.FS.Exists(kind, slug) {
			return s.FS.Load(kind, slug)
		}
	}
	return nil, "", fmt.Errorf("entity not found: %s", slug)
}

// GetByKind loads an entity by kind and slug.
func (s *Store) GetByKind(kind model.EntityKind, slug string) (model.Entity, string, error) {
	return s.FS.Load(kind, slug)
}

// List returns entities of a given kind, optionally filtered.
func (s *Store) List(kind model.EntityKind) ([]model.EntityWithBody, error) {
	return s.FS.LoadAll(kind)
}

// ListAll returns all entities across all kinds.
func (s *Store) ListAll() ([]model.EntityWithBody, error) {
	var all []model.EntityWithBody
	for _, kind := range model.AllEntityKinds() {
		entities, err := s.FS.LoadAll(kind)
		if err != nil {
			continue
		}
		all = append(all, entities...)
	}
	return all, nil
}

// Reindex rebuilds the BadgerDB index from all markdown files.
func (s *Store) Reindex() (*IndexStats, error) {
	return Reindex(s.FS, s.Idx)
}

func (s *Store) indexEntity(e model.Entity, body string) error {
	b := e.GetBase()
	relPath := s.FS.RelativePath(b.Kind, utils.Slugify(b.Name))

	// Read the file back to compute line numbers
	filePath := s.FS.FilePath(b.Kind, utils.Slugify(b.Name))
	data, _ := readFileBytes(filePath)
	lineMap := ComputeLineMap(data)

	ref := FileRef{
		File:  relPath,
		ID:    b.ID,
		Name:  b.Name,
		Kind:  b.Kind,
		Lines: lineMap,
	}

	if err := s.Idx.IndexEntity(ref); err != nil {
		return err
	}

	if len(b.Tags) > 0 {
		s.Idx.IndexTags(ref, b.Tags)
	}

	if b.Status != "" {
		s.Idx.IndexStatus(ref, b.Status)
	}

	if len(b.Relationships) > 0 {
		s.Idx.IndexRelationships(b.ID, relPath, b.Relationships)
	}

	words := extractWords(b, body)
	if len(words) > 0 {
		s.Idx.IndexWords(ref, words)
	}

	return nil
}

func readFileBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}
