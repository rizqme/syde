package storage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/feedloop/syde/internal/model"
)

// IndexSchemaVersion is bumped whenever the on-disk BadgerDB layout
// changes in a way that requires a full reindex. Store.NewStore reads
// the persisted meta:schema key at open time and triggers Reindex when
// the value is missing or lower than this constant.
//
// Version history:
//
//	1 — initial layout (w: keys stored only {file, field})
//	2 — w: values now carry the full FileRef + field + word so search
//	    hits resolve back to a complete entity identity in one lookup.
//	3 — tokenizer also emits CamelCase / snake_case / dash-cased sub-
//	    tokens so loose human queries like "ConceptEntity" or
//	    "add-rel" find their targets without exact phrasing.
const IndexSchemaVersion = 3

// FileRef is the index value pointing to a location in a markdown file.
type FileRef struct {
	File  string         `json:"file"`
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Kind  model.EntityKind `json:"kind"`
	Lines map[string][2]int `json:"lines,omitempty"`
}

// WordRef is the value written to every w:<token>:<kind>:<id> key in
// the inverted index. It carries enough entity identity (via Ref) to
// rebuild a SearchHit without a second LookupByID, plus the field name
// where the match came from and the exact indexed token so the
// formatter can highlight or build a snippet.
type WordRef struct {
	Ref   FileRef `json:"ref"`
	Field string  `json:"field"`
	Word  string  `json:"word"`
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

// SchemaVersion reads the persisted meta:schema value. Missing key
// returns 0 so first-time opens look like version 0 and trigger a
// reindex up to IndexSchemaVersion.
func (idx *Index) SchemaVersion() (int, error) {
	var v int
	err := idx.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("meta:schema"))
		if err == badger.ErrKeyNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			n, perr := strconv.Atoi(string(val))
			if perr != nil {
				return nil
			}
			v = n
			return nil
		})
	})
	return v, err
}

// SetSchemaVersion writes the current IndexSchemaVersion sentinel.
// Called after a successful reindex so subsequent opens skip the
// rebuild path.
func (idx *Index) SetSchemaVersion(v int) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("meta:schema"), []byte(strconv.Itoa(v)))
	})
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

// IndexWords writes inverted index word keys. Each value is a
// WordRef carrying the full FileRef plus the field name where the
// match came from — search hits resolve back to a complete entity
// identity without a second LookupByID.
func (idx *Index) IndexWords(ref FileRef, words map[string]string) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		for word, field := range words {
			key := fmt.Sprintf("w:%s:%s:%s", word, ref.Kind, ref.ID)
			wr := WordRef{Ref: ref, Field: field, Word: word}
			val, _ := json.Marshal(wr)
			if err := txn.Set([]byte(key), val); err != nil {
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

// SearchTokens returns the raw WordRef hits for every token in the
// query, grouped by token. Callers decide whether to AND (intersect
// by entity ID) or OR (union) the per-token slices. Empty query
// returns a nil map. Each hit carries the full FileRef, the field
// name the word came from, and the token itself.
func (idx *Index) SearchTokens(query string) (map[string][]WordRef, error) {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil, nil
	}
	out := make(map[string][]WordRef, len(tokens))
	for _, token := range tokens {
		hits, err := idx.scanWordPrefix([]byte(fmt.Sprintf("w:%s:", token)))
		if err != nil {
			continue
		}
		out[token] = hits
	}
	return out, nil
}

// Search keeps a FileRef-shaped return for back-compat with callers
// that only need entity identity. It performs an OR union across
// tokens, de-duping by entity ID. Prefer SearchTokens in new code —
// it preserves field and word metadata needed for snippets.
func (idx *Index) Search(query string) ([]FileRef, error) {
	tokensMap, err := idx.SearchTokens(query)
	if err != nil {
		return nil, err
	}
	seen := make(map[string]FileRef)
	for _, hits := range tokensMap {
		for _, h := range hits {
			seen[h.Ref.ID] = h.Ref
		}
	}
	results := make([]FileRef, 0, len(seen))
	for _, ref := range seen {
		results = append(results, ref)
	}
	return results, nil
}

// scanWordPrefix walks every w:<token>:* key and unmarshals the
// stored WordRef. Separated from scanPrefix because w: values are
// WordRef-shaped while e:/s:/t: values are FileRef-shaped — the same
// iterator cannot handle both.
func (idx *Index) scanWordPrefix(prefix []byte) ([]WordRef, error) {
	var hits []WordRef
	err := idx.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var wr WordRef
			if err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &wr)
			}); err != nil {
				continue
			}
			hits = append(hits, wr)
		}
		return nil
	})
	return hits, err
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

// tokenize splits text into search tokens while preserving CamelCase
// information that pure-lowercase splitting would lose. The pipeline:
//
//  1. Split on non-alphanumeric runes — _ , - , space, punctuation
//     all become boundaries, so "snake_case", "kebab-case", and
//     "free text" all get split here.
//  2. For each raw word, emit the full lowercased form (so an exact
//     query for "ConceptEntity" still finds "conceptentity") AND
//     every CamelCase sub-token (so the same indexed value also
//     answers a search for "concept" or "entity" alone).
//  3. Skip stop words and 1-char tokens, de-dupe.
//
// This is the v3 tokenizer (IndexSchemaVersion=3). Existing indexes
// auto-rebuild on syded open via the schema-version gate.
func tokenize(text string) []string {
	rawWords := strings.FieldsFunc(text, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
	})
	stop := map[string]bool{"the": true, "a": true, "an": true, "is": true, "it": true, "of": true, "to": true, "in": true, "and": true, "or": true, "for": true}
	seen := make(map[string]bool)
	var result []string
	add := func(w string) {
		w = strings.ToLower(w)
		if len(w) <= 1 || stop[w] || seen[w] {
			return
		}
		seen[w] = true
		result = append(result, w)
	}
	for _, raw := range rawWords {
		add(raw)
		for _, sub := range splitCamel(raw) {
			add(sub)
		}
	}
	return result
}

// splitCamel emits sub-tokens at lowercase→uppercase boundaries.
// "ConceptEntity"      → ["Concept", "Entity"]
// "IndexSchemaVersion" → ["Index", "Schema", "Version"]
// "addRel"             → ["add", "Rel"]
// Pure-lowercase or all-uppercase input returns a single-element slice.
// Acronym-then-lowercase boundaries ("URLPath") are intentionally not
// split — keeps the implementation simple at the cost of one edge
// case; the full lowercased form is still indexed so exact matches
// still work.
func splitCamel(s string) []string {
	if len(s) < 2 {
		return []string{s}
	}
	var out []string
	start := 0
	for i := 1; i < len(s); i++ {
		prev := s[i-1]
		cur := s[i]
		isPrevLower := prev >= 'a' && prev <= 'z'
		isCurUpper := cur >= 'A' && cur <= 'Z'
		if isPrevLower && isCurUpper {
			out = append(out, s[start:i])
			start = i
		}
	}
	out = append(out, s[start:])
	return out
}
