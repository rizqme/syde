package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"gopkg.in/yaml.v3"
)

// writeRequest is the envelope CLI write commands send. The entity
// payload is a YAML frontmatter blob (what storage.FileStore already
// produces and consumes) so we avoid hand-writing JSON schemas for
// every kind. Body is the markdown body.
type writeRequest struct {
	Kind string `json:"kind"`
	// Frontmatter is the YAML serialization of the typed entity — same
	// bytes storage.FileStore.Save would write. Sending it as a string
	// avoids having to reflect the polymorphic entity across the JSON
	// boundary.
	Frontmatter string `json:"frontmatter"`
	Body        string `json:"body"`
}

// writeResponse mirrors what the CLI needs to render a "created/updated"
// message and stay in sync with syded's allocated ID + slug.
type writeResponse struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Kind     string `json:"kind"`
	FilePath string `json:"file_path"`
}

// handleEntityWrite handles POST (create) and PUT (update) on
// /api/<project>/entity. syded is the single writer — it allocates
// the next ID, writes markdown via FileStore, updates the BadgerDB
// index, and returns the saved entity's metadata.
func handleEntityWrite(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		jsonError(w, "POST or PUT required", 405)
		return
	}

	var req writeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "bad body: "+err.Error(), 400)
		return
	}

	kind, ok := model.ParseEntityKind(req.Kind)
	if !ok {
		jsonError(w, "unknown kind: "+req.Kind, 400)
		return
	}

	e := model.NewEntityForKind(kind)
	if err := yaml.Unmarshal([]byte(req.Frontmatter), e); err != nil {
		jsonError(w, "decode frontmatter: "+err.Error(), 400)
		return
	}
	e.GetBase().Kind = kind

	// CreateCascade / UpdateCascade propagate the write up the
	// belongs_to chain so every ancestor's UpdatedAt is bumped — the
	// validator's drift check stays clean when a child changes.
	// Traversal is cycle-safe via a visited-ID set inside Store.
	var filePath string
	var err error
	if r.Method == http.MethodPost {
		filePath, err = store.CreateCascade(e, req.Body)
	} else {
		filePath, err = store.UpdateCascade(e, req.Body)
	}
	if err != nil {
		jsonError(w, err.Error(), 500)
		return
	}

	b := e.GetBase()
	json.NewEncoder(w).Encode(writeResponse{
		ID:       b.ID,
		Slug:     b.CanonicalSlug(),
		Kind:     string(kind),
		FilePath: filePath,
	})
}

// handleEntityDelete backs DELETE /api/<project>/entity/<kind>/<slug>.
// Path form is /entity/<kind>/<slug> to avoid slug-without-kind
// ambiguity during deletes. Uses DeleteCascade so the former parents
// get their UpdatedAt bumped after a child disappears.
func handleEntityDelete(w http.ResponseWriter, kindStr, slug string, store *storage.Store) {
	kind, ok := model.ParseEntityKind(kindStr)
	if !ok {
		jsonError(w, "unknown kind: "+kindStr, 400)
		return
	}
	if kind == model.KindRequirement {
		jsonError(w, "requirements are append-only; mark the requirement superseded or obsolete instead", 400)
		return
	}
	if err := store.DeleteCascade(kind, slug); err != nil {
		jsonError(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"ok": "1", "slug": slug})
}

// entityYAML returns the YAML frontmatter bytes for an entity, matching
// what FileStore.Save would write. Used by the CLI client to build the
// writeRequest payload without duplicating the serialization logic.
func entityYAML(e model.Entity) ([]byte, error) {
	return yaml.Marshal(e)
}

// handleEntityRaw returns the raw markdown file bytes (frontmatter +
// body) for an entity. Unlike the /entity/{slug} read endpoint which
// goes through query.FormatJSON (lossy — drops kind-specific fields),
// this is the lossless path the CLI write-client uses before mutating
// and posting back. Format: same on-disk bytes as FileStore.Save
// produced, so round-tripping through storage.UnmarshalAuto preserves
// every field.
func handleEntityRaw(w http.ResponseWriter, slug string, store *storage.Store) {
	e, body, err := store.Get(slug)
	if err != nil {
		jsonError(w, err.Error(), 404)
		return
	}
	raw, err := serializeEntityMarkdown(e, body)
	if err != nil {
		jsonError(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(raw)
}

func serializeEntityMarkdown(e model.Entity, body string) ([]byte, error) {
	// Reuse the shared serializer so the bytes match what FileStore
	// writes to disk — guarantees a lossless round-trip.
	return storage.Marshal(e, body)
}

var _ = fmt.Sprintf // keep import until used
