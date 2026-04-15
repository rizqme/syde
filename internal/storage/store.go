package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// Store provides unified access to the file store and index.
type Store struct {
	FS  *FileStore
	Idx *Index
}

// NewStore creates a Store with file store and BadgerDB index. It
// also checks the persisted index schema version and triggers a full
// reindex when the value is missing or older than IndexSchemaVersion
// — this is how upgrades that reshape index key/value layouts self-
// heal without user action.
func NewStore(sydeDir string) (*Store, error) {
	fs := NewFileStore(sydeDir)
	idx, err := OpenIndex(filepath.Join(sydeDir, "index"))
	if err != nil {
		return nil, err
	}
	s := &Store{FS: fs, Idx: idx}

	v, _ := idx.SchemaVersion()
	if v < IndexSchemaVersion {
		if _, err := s.Reindex(); err != nil {
			return nil, fmt.Errorf("auto-reindex to schema v%d: %w", IndexSchemaVersion, err)
		}
		if err := idx.SetSchemaVersion(IndexSchemaVersion); err != nil {
			return nil, fmt.Errorf("persist schema version: %w", err)
		}
	}
	return s, nil
}

// ReindexOne refreshes a single entity's index entries from its on-disk
// markdown file, without mutating the file itself. Used by syded's
// incremental /reindex endpoint so the CLI can write markdown and then
// ask syded to refresh its index without bumping UpdatedAt.
func (s *Store) ReindexOne(kind model.EntityKind, slug string) error {
	e, body, err := s.FS.Load(kind, slug)
	if err != nil {
		return err
	}
	s.Idx.RemoveEntity(kind, e.GetBase().ID)
	return s.indexEntity(e, body)
}

// Close closes the index.
func (s *Store) Close() error {
	return s.Idx.Close()
}

// Create saves a new entity and indexes it.
func (s *Store) Create(e model.Entity, body string) (string, error) {
	b := e.GetBase()
	if b.ID == "" {
		id, err := NextID(s.Idx, b.Kind)
		if err != nil {
			return "", fmt.Errorf("allocate id: %w", err)
		}
		b.ID = id
	}
	if b.Slug == "" {
		// Keep picking a suffix until we find one not on disk yet.
		// Four random alphanumeric chars = 1.6M possibilities, so this
		// almost always succeeds on the first try.
		for i := 0; i < 10; i++ {
			candidate := utils.SlugifyWithSuffix(b.Name)
			if !s.FS.Exists(b.Kind, candidate) {
				b.Slug = candidate
				break
			}
		}
		if b.Slug == "" {
			return "", fmt.Errorf("could not allocate unique slug for %q after 10 tries", b.Name)
		}
	}
	if b.UpdatedAt == "" {
		b.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
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

	b.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

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

// CreateCascade wraps Create and then bumps every ancestor's
// UpdatedAt via the belongs_to chain. Used by server-side write
// handlers so a new child always marks its parent chain fresh.
func (s *Store) CreateCascade(e model.Entity, body string) (string, error) {
	filePath, err := s.Create(e, body)
	if err != nil {
		return filePath, err
	}
	// Seed the visited set with the newly-created entity so a buggy
	// belongs_to-to-self (or a cycle where an ancestor points back at
	// the child) terminates on the first recurse.
	visited := map[string]bool{e.GetBase().ID: true}
	s.cascadeFromParent(e, visited)
	return filePath, nil
}

// UpdateCascade wraps Update and then bumps every ancestor's
// UpdatedAt via the belongs_to chain.
func (s *Store) UpdateCascade(e model.Entity, body string) (string, error) {
	filePath, err := s.Update(e, body)
	if err != nil {
		return filePath, err
	}
	visited := map[string]bool{e.GetBase().ID: true}
	s.cascadeFromParent(e, visited)
	return filePath, nil
}

// DeleteCascade loads the entity first (to capture its belongs_to
// targets), deletes it, and then cascades to every former parent so
// the parent chain reflects that a child has disappeared.
func (s *Store) DeleteCascade(kind model.EntityKind, slug string) error {
	e, _, err := s.FS.Load(kind, slug)
	if err != nil {
		return err
	}
	// Capture belongs_to before we remove the entity — after Delete
	// the in-memory copy is all we have.
	captured := *e.GetBase()
	if err := s.Delete(kind, slug); err != nil {
		return err
	}
	// Use a synthetic BaseEntity so cascadeFromParent can iterate
	// captured.Relationships without needing a live Entity value.
	visited := map[string]bool{captured.ID: true}
	s.cascadeFromBase(&captured, visited)
	return nil
}

// cascadeFromParent walks the belongs_to relationships on an entity
// and, for each parent, loads it, marks it visited, runs Update (to
// stamp fresh UpdatedAt and reindex), and recurses. Unknown parents
// (Get fails) are silently skipped — the validator will surface those
// as broken-reference errors separately. The visited-ID map makes the
// traversal cycle-safe: even a self-loop or A->B->A terminates after
// each distinct entity has been processed once.
func (s *Store) cascadeFromParent(e model.Entity, visited map[string]bool) {
	s.cascadeFromBase(e.GetBase(), visited)
}

// cascadeFromBase is the actual worker. Split out so DeleteCascade can
// pass a captured BaseEntity (without a live Entity shell) after the
// underlying file has been removed.
func (s *Store) cascadeFromBase(b *model.BaseEntity, visited map[string]bool) {
	for _, rel := range b.Relationships {
		if rel.Type != model.RelBelongsTo {
			continue
		}
		parent, parentBody, err := s.Get(rel.Target)
		if err != nil {
			// Broken belongs_to target — validator's job to report.
			continue
		}
		pb := parent.GetBase()
		if visited[pb.ID] {
			continue
		}
		visited[pb.ID] = true
		if _, err := s.Update(parent, parentBody); err != nil {
			// Best-effort cascade; log-free path because Store has no
			// logger. Callers that care can wrap and inspect.
			continue
		}
		s.cascadeFromBase(pb, visited)
	}
}

// Get loads an entity by slug. Supports three addressing forms:
//
//  1. Full slug with -XXXX suffix (always unique): "cli-a3f2"
//  2. Bare name slug: "cli" — matches any file whose slug base is "cli".
//     If multiple entities share the same bare slug across or within
//     kinds, returns an "ambiguous" error listing the candidates.
//  3. Parent/child path: "syde-cli/cli" — resolves the parent first
//     (recursively), then picks the child whose belongs_to points at
//     that parent. Works with any mix of full/bare slugs at each step.
func (s *Store) Get(slug string) (model.Entity, string, error) {
	// Parent/child form
	if strings.Contains(slug, "/") {
		return s.resolvePath(slug)
	}

	// Exact filename match (full slug or legacy bare slug)
	for _, kind := range model.AllEntityKinds() {
		if s.FS.Exists(kind, slug) {
			return s.FS.Load(kind, slug)
		}
	}

	// Bare slug fallback: scan every kind for files whose base slug matches.
	var matches []struct {
		kind model.EntityKind
		file string // slug without .md
	}
	for _, kind := range model.AllEntityKinds() {
		files, err := s.FS.ListFiles(kind)
		if err != nil {
			continue
		}
		for _, f := range files {
			name := filepath.Base(f)
			name = strings.TrimSuffix(name, ".md")
			if utils.BaseSlug(name) == slug {
				matches = append(matches, struct {
					kind model.EntityKind
					file string
				}{kind, name})
			}
		}
	}
	if len(matches) == 1 {
		return s.FS.Load(matches[0].kind, matches[0].file)
	}
	if len(matches) > 1 {
		var names []string
		for _, m := range matches {
			names = append(names, fmt.Sprintf("%s/%s", m.kind, m.file))
		}
		return nil, "", fmt.Errorf("ambiguous slug %q matches %d entities: %s — use the full slug or parent/child form",
			slug, len(matches), strings.Join(names, ", "))
	}
	return nil, "", fmt.Errorf("entity not found: %s", slug)
}

// resolvePath walks a "parent-slug/child-slug[/...]" addressing string.
// Each segment is resolved independently; subsequent segments must be a
// child whose belongs_to target is the previous segment's entity.
func (s *Store) resolvePath(path string) (model.Entity, string, error) {
	segs := strings.Split(path, "/")
	if len(segs) == 0 {
		return nil, "", fmt.Errorf("empty path")
	}

	// First segment: resolve as a normal slug.
	parent, body, err := s.Get(segs[0])
	if err != nil {
		return nil, "", fmt.Errorf("resolve %q: %w", segs[0], err)
	}

	// Walk remaining segments, each must be a child of the previous.
	for i := 1; i < len(segs); i++ {
		child := segs[i]
		found, foundBody, ok := s.findChildOf(parent, child)
		if !ok {
			return nil, "", fmt.Errorf("no child %q of %s/%s", child, parent.GetBase().Kind, parent.GetBase().Slug)
		}
		parent = found
		body = foundBody
	}
	return parent, body, nil
}

// findChildOf finds an entity whose slug matches `child` (full or bare)
// AND whose belongs_to relationship points at `parent`.
func (s *Store) findChildOf(parent model.Entity, child string) (model.Entity, string, bool) {
	pb := parent.GetBase()
	parentKeys := map[string]bool{
		pb.ID:   true,
		pb.Slug: true,
	}
	if pb.Slug != "" {
		parentKeys[utils.BaseSlug(pb.Slug)] = true
	}

	all, _ := s.ListAll()
	for _, ewb := range all {
		eb := ewb.Entity.GetBase()
		if eb.Slug != child && utils.BaseSlug(eb.Slug) != child {
			continue
		}
		for _, rel := range eb.Relationships {
			if rel.Type != model.RelBelongsTo {
				continue
			}
			if parentKeys[rel.Target] {
				return ewb.Entity, ewb.Body, true
			}
		}
	}
	return nil, "", false
}

// GetByKind loads an entity by kind and slug. Accepts any of the three
// slug forms (full suffixed slug, bare name slug, parent/child path).
// For the bare form, ambiguity within the kind is surfaced as an
// "ambiguous" error listing candidates.
func (s *Store) GetByKind(kind model.EntityKind, slug string) (model.Entity, string, error) {
	// Exact file match first (full slug or legacy bare filename)
	if s.FS.Exists(kind, slug) {
		return s.FS.Load(kind, slug)
	}

	// Parent/child path
	if strings.Contains(slug, "/") {
		e, body, err := s.resolvePath(slug)
		if err != nil {
			return nil, "", err
		}
		if e.GetBase().Kind != kind {
			return nil, "", fmt.Errorf("entity %s is kind %s, not %s", slug, e.GetBase().Kind, kind)
		}
		return e, body, nil
	}

	// Bare slug scan within this kind
	files, err := s.FS.ListFiles(kind)
	if err != nil {
		return nil, "", fmt.Errorf("list %s: %w", kind, err)
	}
	var matches []string
	for _, f := range files {
		name := strings.TrimSuffix(filepath.Base(f), ".md")
		if utils.BaseSlug(name) == slug {
			matches = append(matches, name)
		}
	}
	if len(matches) == 1 {
		return s.FS.Load(kind, matches[0])
	}
	if len(matches) > 1 {
		return nil, "", fmt.Errorf("ambiguous %s slug %q matches %d entities: %s — use the full slug",
			kind, slug, len(matches), strings.Join(matches, ", "))
	}
	return nil, "", fmt.Errorf("%s not found: %s", kind, slug)
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
	slug := b.Slug
	if slug == "" {
		slug = utils.Slugify(b.Name)
	}
	relPath := s.FS.RelativePath(b.Kind, slug)

	// Read the file back to compute line numbers
	filePath := s.FS.FilePath(b.Kind, slug)
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
