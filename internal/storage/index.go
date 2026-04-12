package storage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/feedloop/syde/internal/model"
)

// FileRef is the index value pointing to a location in a markdown file.
type FileRef struct {
	File  string         `json:"file"`
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Kind  model.EntityKind `json:"kind"`
	Lines map[string][2]int `json:"lines,omitempty"`
}

// RelRef is a relationship index value with source/target file info.
type RelRef struct {
	Label      string `json:"label,omitempty"`
	SourceFile string `json:"source_file"`
	TargetFile string `json:"target_file,omitempty"`
	Line       int    `json:"line,omitempty"`
}

// Index wraps BadgerDB for entity indexing.
type Index struct {
	db *badger.DB
}

// OpenIndex opens or creates a BadgerDB index at the given path.
func OpenIndex(path string) (*Index, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil // Silence BadgerDB logs
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("open index: %w", err)
	}
	return &Index{db: db}, nil
}

// Close closes the BadgerDB index.
func (idx *Index) Close() error {
	return idx.db.Close()
}

// IndexEntity writes all index keys for an entity.
func (idx *Index) IndexEntity(ref FileRef) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		refBytes, _ := json.Marshal(ref)

		// e:{kind}:{id}
		if err := txn.Set([]byte(fmt.Sprintf("e:%s:%s", ref.Kind, ref.ID)), refBytes); err != nil {
			return err
		}

		// s:{kind}:{slug}
		slug := slugFromFile(ref.File)
		if err := txn.Set([]byte(fmt.Sprintf("s:%s:%s", ref.Kind, slug)), refBytes); err != nil {
			return err
		}

		return nil
	})
}

// IndexTags writes tag index keys for an entity.
func (idx *Index) IndexTags(ref FileRef, tags []string) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		compact, _ := json.Marshal(FileRef{File: ref.File, ID: ref.ID, Name: ref.Name, Kind: ref.Kind})
		for _, tag := range tags {
			key := fmt.Sprintf("t:%s:%s:%s", tag, ref.Kind, ref.ID)
			if err := txn.Set([]byte(key), compact); err != nil {
				return err
			}
		}
		return nil
	})
}

// IndexStatus writes status index key for an entity.
func (idx *Index) IndexStatus(ref FileRef, status model.Status) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		compact, _ := json.Marshal(FileRef{File: ref.File, ID: ref.ID, Name: ref.Name, Kind: ref.Kind})
		key := fmt.Sprintf("st:%s:%s:%s", status, ref.Kind, ref.ID)
		return txn.Set([]byte(key), compact)
	})
}

// IndexRelationships writes relationship index keys.
func (idx *Index) IndexRelationships(sourceID, sourceFile string, rels []model.Relationship) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		for _, rel := range rels {
			rr, _ := json.Marshal(RelRef{
				Label:      rel.Label,
				SourceFile: sourceFile,
			})
			outKey := fmt.Sprintf("r:out:%s:%s:%s", sourceID, rel.Type, rel.Target)
			if err := txn.Set([]byte(outKey), rr); err != nil {
				return err
			}
			inKey := fmt.Sprintf("r:in:%s:%s:%s", rel.Target, rel.Type, sourceID)
			if err := txn.Set([]byte(inKey), rr); err != nil {
				return err
			}
		}
		return nil
	})
}

// IndexWords writes inverted index word keys.
func (idx *Index) IndexWords(ref FileRef, words map[string]string) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		for word, field := range words {
			key := fmt.Sprintf("w:%s:%s:%s", word, ref.Kind, ref.ID)
			val := fmt.Sprintf(`{"file":"%s","field":"%s"}`, ref.File, field)
			if err := txn.Set([]byte(key), []byte(val)); err != nil {
				return err
			}
		}
		return nil
	})
}

// RemoveEntity removes all index keys for an entity.
func (idx *Index) RemoveEntity(kind model.EntityKind, id string) error {
	prefixes := []string{
		fmt.Sprintf("e:%s:%s", kind, id),
	}

	return idx.db.Update(func(txn *badger.Txn) error {
		for _, prefix := range prefixes {
			if err := txn.Delete([]byte(prefix)); err != nil && err != badger.ErrKeyNotFound {
				return err
			}
		}
		// Scan and delete tag, status, relationship, and word keys
		return idx.deleteByPrefix(txn, fmt.Sprintf("t:"), kind, id)
	})
}

func (idx *Index) deleteByPrefix(txn *badger.Txn, _ string, kind model.EntityKind, id string) error {
	// Delete all keys ending with :{kind}:{id}
	suffix := fmt.Sprintf(":%s:%s", kind, id)
	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false
	it := txn.NewIterator(opts)
	defer it.Close()

	var toDelete [][]byte
	for it.Rewind(); it.Valid(); it.Next() {
		key := it.Item().Key()
		if strings.HasSuffix(string(key), suffix) {
			cp := make([]byte, len(key))
			copy(cp, key)
			toDelete = append(toDelete, cp)
		}
	}
	for _, key := range toDelete {
		if err := txn.Delete(key); err != nil {
			return err
		}
	}
	return nil
}

// LookupByID finds an entity by kind and ID.
func (idx *Index) LookupByID(kind model.EntityKind, id string) (*FileRef, error) {
	var ref FileRef
	err := idx.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("e:%s:%s", kind, id)))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &ref)
		})
	})
	if err != nil {
		return nil, err
	}
	return &ref, nil
}

// LookupBySlug finds an entity by kind and slug.
func (idx *Index) LookupBySlug(kind model.EntityKind, slug string) (*FileRef, error) {
	var ref FileRef
	err := idx.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("s:%s:%s", kind, slug)))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &ref)
		})
	})
	if err != nil {
		return nil, err
	}
	return &ref, nil
}

// ListByKind returns all FileRefs for a given entity kind.
func (idx *Index) ListByKind(kind model.EntityKind) ([]FileRef, error) {
	prefix := []byte(fmt.Sprintf("e:%s:", kind))
	return idx.scanPrefix(prefix)
}

// ListByTag returns all FileRefs for a given tag.
func (idx *Index) ListByTag(tag string) ([]FileRef, error) {
	prefix := []byte(fmt.Sprintf("t:%s:", tag))
	return idx.scanPrefix(prefix)
}

// ListByStatus returns all FileRefs for a given status and optional kind.
func (idx *Index) ListByStatus(status model.Status, kind model.EntityKind) ([]FileRef, error) {
	var prefix []byte
	if kind != "" {
		prefix = []byte(fmt.Sprintf("st:%s:%s:", status, kind))
	} else {
		prefix = []byte(fmt.Sprintf("st:%s:", status))
	}
	return idx.scanPrefix(prefix)
}

// GetOutbound returns all outbound relationships for an entity.
func (idx *Index) GetOutbound(entityID string) ([]struct {
	Type   string
	Target string
	Rel    RelRef
}, error) {
	prefix := []byte(fmt.Sprintf("r:out:%s:", entityID))
	var results []struct {
		Type   string
		Target string
		Rel    RelRef
	}

	err := idx.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			key := string(item.Key())
			// key format: r:out:{source}:{type}:{target}
			parts := strings.SplitN(key, ":", 5)
			if len(parts) < 5 {
				continue
			}
			var rr RelRef
			if err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &rr)
			}); err != nil {
				continue
			}
			results = append(results, struct {
				Type   string
				Target string
				Rel    RelRef
			}{Type: parts[3], Target: parts[4], Rel: rr})
		}
		return nil
	})
	return results, err
}

// GetInbound returns all inbound relationships for an entity.
func (idx *Index) GetInbound(entityID string) ([]struct {
	Type   string
	Source string
	Rel    RelRef
}, error) {
	prefix := []byte(fmt.Sprintf("r:in:%s:", entityID))
	var results []struct {
		Type   string
		Source string
		Rel    RelRef
	}

	err := idx.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			key := string(item.Key())
			parts := strings.SplitN(key, ":", 5)
			if len(parts) < 5 {
				continue
			}
			var rr RelRef
			if err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &rr)
			}); err != nil {
				continue
			}
			results = append(results, struct {
				Type   string
				Source string
				Rel    RelRef
			}{Type: parts[3], Source: parts[4], Rel: rr})
		}
		return nil
	})
	return results, err
}

// Search performs a simple inverted index search.
func (idx *Index) Search(query string) ([]FileRef, error) {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil, nil
	}

	// For simplicity, OR search: return all entities matching any token
	seen := make(map[string]FileRef)
	for _, token := range tokens {
		refs, err := idx.scanPrefix([]byte(fmt.Sprintf("w:%s:", token)))
		if err != nil {
			continue
		}
		for _, ref := range refs {
			seen[ref.ID] = ref
		}
	}

	var results []FileRef
	for _, ref := range seen {
		results = append(results, ref)
	}
	return results, nil
}

// DropAll removes all keys from the index.
func (idx *Index) DropAll() error {
	return idx.db.DropAll()
}

func (idx *Index) scanPrefix(prefix []byte) ([]FileRef, error) {
	var refs []FileRef
	err := idx.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var ref FileRef
			if err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &ref)
			}); err != nil {
				continue
			}
			refs = append(refs, ref)
		}
		return nil
	})
	return refs, err
}

func slugFromFile(file string) string {
	base := file
	if idx := strings.LastIndex(file, "/"); idx >= 0 {
		base = file[idx+1:]
	}
	return strings.TrimSuffix(base, ".md")
}

func tokenize(text string) []string {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'))
	})
	// Filter stop words
	stop := map[string]bool{"the": true, "a": true, "an": true, "is": true, "it": true, "of": true, "to": true, "in": true, "and": true, "or": true, "for": true}
	var result []string
	for _, w := range words {
		if len(w) > 1 && !stop[w] {
			result = append(result, w)
		}
	}
	return result
}
