package cli

import (
	"fmt"

	"github.com/feedloop/syde/internal/client"
	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"gopkg.in/yaml.v3"
)

// writeClient wraps an HTTP client so it can stand in where CLI write
// commands used to call storage.Store.Create / Update / Delete. Every
// write command previously opened a Store, mutated an Entity, and
// called a method on Store — now they build the same Entity and hand
// it to writeClient. The helper serializes to YAML frontmatter (the
// same bytes FileStore writes to disk) and POSTs to syded, which owns
// the BadgerDB index.
//
// This keeps the write-command diff small: only the `store` variable
// name + the construction line change; the rest of each command is
// identical.
type writeClient struct {
	c *client.Client
	// FS is a local FileStore used ONLY for path operations
	// (RelativePath, Root) that don't touch BadgerDB. Write paths go
	// through the HTTP client, not through FS.Save.
	FS *storage.FileStore
}

func openWriteClient() (*writeClient, error) {
	dir := sydeDir
	if dir == "" {
		var err error
		dir, err = config.FindSydeDir()
		if err != nil {
			return nil, fmt.Errorf("no .syde/ directory found (run 'syde init' first)")
		}
	}
	c, err := client.New(client.Options{SydeDir: dir})
	if err != nil {
		return nil, err
	}
	return &writeClient{c: c, FS: storage.NewFileStore(dir)}, nil
}

// Close is a no-op — kept so `defer wc.Close()` in write commands
// remains valid after the refactor without having to delete every
// defer line.
func (w *writeClient) Close() error { return nil }

// Create serializes the entity to YAML frontmatter and POSTs it to
// syded. The caller's entity is patched with the server-allocated ID
// and slug so subsequent code (error messages, status prints) has the
// right values.
func (w *writeClient) Create(e model.Entity, body string) (string, error) {
	fm, err := yaml.Marshal(e)
	if err != nil {
		return "", fmt.Errorf("marshal entity: %w", err)
	}
	resp, err := w.c.CreateEntity(string(e.GetBase().Kind), string(fm), body)
	if err != nil {
		return "", err
	}
	b := e.GetBase()
	b.ID = resp.ID
	b.Slug = resp.Slug
	return resp.FilePath, nil
}

// Update is the replacement for storage.Store.Update. Like Create it
// forwards a full entity snapshot to syded — syded overwrites the
// file and reindexes in one step.
func (w *writeClient) Update(e model.Entity, body string) (string, error) {
	fm, err := yaml.Marshal(e)
	if err != nil {
		return "", fmt.Errorf("marshal entity: %w", err)
	}
	resp, err := w.c.UpdateEntity(string(e.GetBase().Kind), string(fm), body)
	if err != nil {
		return "", err
	}
	return resp.FilePath, nil
}

// Get proxies to client.GetRaw (lossless markdown) and decodes into a
// typed entity via storage.UnmarshalAuto. Write commands commonly load
// an entity, mutate it, then Update it back — the raw round-trip
// preserves every field, unlike query.FormatJSON which only emits a
// curated subset.
func (w *writeClient) Get(slug string) (model.Entity, string, error) {
	raw, err := w.c.GetRaw(slug)
	if err != nil {
		return nil, "", err
	}
	return storage.UnmarshalAuto(raw)
}

// GetByKind is the kind-filtered variant. Uses the same raw markdown
// round-trip as Get, then type-checks.
func (w *writeClient) GetByKind(kind model.EntityKind, slug string) (model.Entity, string, error) {
	e, body, err := w.Get(slug)
	if err != nil {
		return nil, "", err
	}
	if e.GetBase().Kind != kind {
		return nil, "", fmt.Errorf("entity %s is kind %s, not %s", slug, e.GetBase().Kind, kind)
	}
	return e, body, nil
}

// Delete removes an entity.
func (w *writeClient) Delete(kind model.EntityKind, slug string) error {
	return w.c.DeleteEntity(string(kind), slug)
}

// List proxies to client.List then rehydrates typed entities by
// fetching each one via the lossless raw endpoint. N+1 HTTP calls
// are fine for write commands' "list + pick" pattern at typical
// project sizes (~150 entities); the alternative (raw bulk endpoint)
// is a future optimization.
func (w *writeClient) List(kind model.EntityKind) ([]model.EntityWithBody, error) {
	summaries, err := w.c.List(string(kind), "")
	if err != nil {
		return nil, err
	}
	var out []model.EntityWithBody
	for _, s := range summaries {
		e, body, err := w.Get(s.Slug)
		if err != nil {
			continue
		}
		out = append(out, model.EntityWithBody{Entity: e, Body: body})
	}
	return out, nil
}

// ListAll iterates every kind via List.
func (w *writeClient) ListAll() ([]model.EntityWithBody, error) {
	var out []model.EntityWithBody
	for _, k := range model.AllEntityKinds() {
		batch, err := w.List(k)
		if err != nil {
			continue
		}
		out = append(out, batch...)
	}
	return out, nil
}

// decodeEntityFromQueryJSON parses the output of query.FormatJSON into
// a typed entity + body. The server returns a flat map (via
// query.FormatJSON) where "entity" is the base payload — we re-marshal
// it through YAML into the typed struct so every kind-specific field
// round-trips cleanly.
func decodeEntityFromQueryJSON(raw []byte) (model.Entity, string, error) {
	// query.FormatJSON produces a JSON object with keys like:
	//   entity: { id, kind, name, ... }, body: "...", relationships: [...]
	// We unmarshal the "entity" sub-object back into a typed Entity
	// via YAML (safer than hand-matching fields) by re-serializing to
	// YAML and calling yaml.Unmarshal.
	var envelope struct {
		Entity map[string]interface{} `json:"entity"`
		Body   string                 `json:"body"`
	}
	if err := yamlJSONDecode(raw, &envelope); err != nil {
		return nil, "", fmt.Errorf("decode entity envelope: %w", err)
	}
	kindStr, _ := envelope.Entity["kind"].(string)
	kind, ok := model.ParseEntityKind(kindStr)
	if !ok {
		return nil, "", fmt.Errorf("unknown kind in response: %q", kindStr)
	}
	// Round-trip through YAML so typed sub-fields (relationships,
	// capabilities, phases, ...) decode properly.
	fm, err := yaml.Marshal(envelope.Entity)
	if err != nil {
		return nil, "", err
	}
	e := model.NewEntityForKind(kind)
	if err := yaml.Unmarshal(fm, e); err != nil {
		return nil, "", err
	}
	e.GetBase().Kind = kind
	return e, envelope.Body, nil
}

// yamlJSONDecode lets us decode a JSON payload via the YAML unmarshaler
// which tolerates both JSON and YAML input. Avoids pulling encoding/json
// into this helper file and keeps the round-trip single-tool.
func yamlJSONDecode(data []byte, out interface{}) error {
	return yaml.Unmarshal(data, out)
}
