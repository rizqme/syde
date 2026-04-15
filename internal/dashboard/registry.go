package dashboard

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/feedloop/syde/internal/storage"
)

// Per-project Store handle cache. syded is now the sole owner of the
// on-disk BadgerDB index — the CLI talks to syded over HTTP instead of
// opening its own handle, so there's no more lock contention and the
// in-memory / fsnotify scaffolding from the previous phase is gone.
// Writes come in through the /api/<project>/entity endpoints, which
// call Store.Create/Update/Delete directly under syded's single
// writer.
var (
	storeCacheMu sync.Mutex
	storeCache   = make(map[string]*storage.Store)
)

// GetStore returns a cached Store for the given .syde/ directory,
// opening the persistent BadgerDB index on first access. Callers MUST
// NOT Close() the returned store — the cache owns its lifetime.
func GetStore(sydeDir string) (*storage.Store, error) {
	storeCacheMu.Lock()
	defer storeCacheMu.Unlock()
	if s, ok := storeCache[sydeDir]; ok {
		return s, nil
	}
	s, err := storage.NewStore(sydeDir)
	if err != nil {
		return nil, err
	}
	storeCache[sydeDir] = s
	return s, nil
}

// ProjectEntry is a registered project in the dashboard.
type ProjectEntry struct {
	Slug       string `json:"slug"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	LastOpened string `json:"last_opened"`
}

// MakeProjectSlug creates a slug in the format {project-name}-{4char-id}.
// The 4-char ID is derived from a hash of the absolute path for uniqueness.
func MakeProjectSlug(projectName, absPath string) string {
	name := strings.ToLower(strings.TrimSpace(projectName))
	name = strings.ReplaceAll(name, " ", "-")

	h := sha256.Sum256([]byte(absPath))
	shortID := hex.EncodeToString(h[:])[:4]

	return fmt.Sprintf("%s-%s", name, shortID)
}

// RegisterProject adds or updates a project in the global registry.
func RegisterProject(slug, path, name string) error {
	projects, _ := loadProjectRegistry()

	found := false
	for i, p := range projects {
		if p.Slug == slug || p.Path == path {
			projects[i].Slug = slug
			projects[i].Path = path
			projects[i].Name = name
			projects[i].LastOpened = time.Now().UTC().Format(time.RFC3339)
			found = true
			break
		}
	}
	if !found {
		projects = append(projects, ProjectEntry{
			Slug:       slug,
			Path:       path,
			Name:       name,
			LastOpened: time.Now().UTC().Format(time.RFC3339),
		})
	}

	return saveProjectRegistry(projects)
}

// FindProjectBySlug looks up a project by its slug.
func FindProjectBySlug(slug string) (*ProjectEntry, error) {
	projects, err := loadProjectRegistry()
	if err != nil {
		return nil, err
	}
	for _, p := range projects {
		if p.Slug == slug {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("project not found: %s", slug)
}

func loadProjectRegistry() ([]ProjectEntry, error) {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".syde", "projects.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var reg struct {
		Projects []ProjectEntry `json:"projects"`
	}
	if err := json.Unmarshal(data, &reg); err != nil {
		return nil, err
	}
	return reg.Projects, nil
}

func saveProjectRegistry(projects []ProjectEntry) error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".syde")
	os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, "projects.json")
	data, err := json.MarshalIndent(map[string]interface{}{"projects": projects}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
