package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// FileStore handles reading and writing entity markdown files in .syde/.
type FileStore struct {
	Root string // Path to .syde/ directory
}

// NewFileStore creates a FileStore for the given .syde/ directory.
func NewFileStore(root string) *FileStore {
	return &FileStore{Root: root}
}

// Save writes an entity to a markdown file. Uses b.Slug (the full slug
// with random suffix) as the filename; falls back to Slugify(name) for
// legacy entities created before the slug field existed.
func (fs *FileStore) Save(e model.Entity, body string) (string, error) {
	b := e.GetBase()
	dir := filepath.Join(fs.Root, b.Kind.KindPlural())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create dir: %w", err)
	}

	slug := b.Slug
	if slug == "" {
		slug = utils.Slugify(b.Name)
	}
	filePath := filepath.Join(dir, slug+".md")

	data, err := Marshal(e, body)
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	return filePath, nil
}

// Load reads an entity from a markdown file.
func (fs *FileStore) Load(kind model.EntityKind, slug string) (model.Entity, string, error) {
	filePath := filepath.Join(fs.Root, kind.KindPlural(), slug+".md")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("read file: %w", err)
	}

	return Unmarshal(data, kind)
}

// LoadFile reads an entity from a specific file path.
func (fs *FileStore) LoadFile(path string) (model.Entity, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("read file: %w", err)
	}

	return UnmarshalAuto(data)
}

// Delete removes an entity file.
func (fs *FileStore) Delete(kind model.EntityKind, slug string) error {
	filePath := filepath.Join(fs.Root, kind.KindPlural(), slug+".md")
	return os.Remove(filePath)
}

// ListFiles returns all markdown files for a given entity kind.
func (fs *FileStore) ListFiles(kind model.EntityKind) ([]string, error) {
	dir := filepath.Join(fs.Root, kind.KindPlural())
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, nil
}

// LoadAll reads all entities of a given kind.
func (fs *FileStore) LoadAll(kind model.EntityKind) ([]model.EntityWithBody, error) {
	files, err := fs.ListFiles(kind)
	if err != nil {
		return nil, err
	}

	var results []model.EntityWithBody
	for _, f := range files {
		e, body, err := fs.LoadFile(f)
		if err != nil {
			continue // skip malformed files
		}
		results = append(results, model.EntityWithBody{Entity: e, Body: body})
	}
	return results, nil
}

// FilePath returns the expected file path for an entity.
func (fs *FileStore) FilePath(kind model.EntityKind, slug string) string {
	return filepath.Join(fs.Root, kind.KindPlural(), slug+".md")
}

// RelativePath returns the path relative to .syde/ root.
func (fs *FileStore) RelativePath(kind model.EntityKind, slug string) string {
	return filepath.Join(kind.KindPlural(), slug+".md")
}

// Exists checks if an entity file exists.
func (fs *FileStore) Exists(kind model.EntityKind, slug string) bool {
	_, err := os.Stat(fs.FilePath(kind, slug))
	return err == nil
}

// AllKinds returns all entity kinds that have directories.
func (fs *FileStore) AllKinds() []model.EntityKind {
	var kinds []model.EntityKind
	for _, k := range model.AllEntityKinds() {
		dir := filepath.Join(fs.Root, k.KindPlural())
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			kinds = append(kinds, k)
		}
	}
	return kinds
}
