// Package client is the CLI's thin HTTP wrapper around syded's JSON
// API. Every `syde` read command that previously opened BadgerDB now
// goes through Client.* methods instead — syded is the single owner of
// the on-disk index and we never contend with it for the dir lock.
//
// Auto-launch: New() calls daemon.EnsureRunning before returning, so
// callers can assume "syded is up" from the first method call. On
// fresh clones / CI this forks a detached syded transparently.
package client

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/daemon"
)

// Client talks to a local syded over HTTP. Every project is keyed by
// the same hash-derived slug the dashboard uses, so all projects
// registered with syded are addressable.
type Client struct {
	base        string // e.g. http://localhost:5703
	projectSlug string
	http        *http.Client
}

// Options tweaks Client construction. Zero value uses defaults.
type Options struct {
	Port        int    // 0 -> daemon.DefaultPort
	ProjectSlug string // override auto-detection
	SydeDir     string // path to .syde/ dir; required unless ProjectSlug is set
}

// New constructs a Client, ensuring syded is running, and derives the
// project slug from the given sydeDir. Most CLI commands can call
// New(Options{SydeDir: dir}) and be done.
func New(opts Options) (*Client, error) {
	port := opts.Port
	if port == 0 {
		port = daemon.DefaultPort
	}

	if err := daemon.EnsureRunning(port, func(f string, a ...interface{}) {
		fmt.Fprintf(os.Stderr, f, a...)
	}); err != nil {
		return nil, err
	}

	slug := opts.ProjectSlug
	if slug == "" {
		if opts.SydeDir == "" {
			return nil, fmt.Errorf("client: need SydeDir or ProjectSlug")
		}
		var err error
		slug, err = DeriveProjectSlug(opts.SydeDir)
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		base:        fmt.Sprintf("http://localhost:%d", port),
		projectSlug: slug,
		http:        &http.Client{Timeout: 30 * time.Second},
	}, nil
}

// DeriveProjectSlug mirrors dashboard.MakeProjectSlug so the CLI and
// syded agree on the project key for a given .syde/ dir. Format:
// {project-name}-{4char-sha256-of-abs-path}. If syde.yaml names the
// project, that configured name is part of the slug.
func DeriveProjectSlug(sydeDir string) (string, error) {
	abs, err := filepath.Abs(sydeDir)
	if err != nil {
		return "", err
	}
	projectRoot := filepath.Dir(abs)
	name := filepath.Base(projectRoot)
	if cfg, err := config.Load(abs); err == nil && cfg != nil && cfg.Project != "" {
		name = cfg.Project
	}
	name = strings.ToLower(strings.TrimSpace(name))
	name = strings.ReplaceAll(name, " ", "-")
	h := sha256.Sum256([]byte(projectRoot))
	return fmt.Sprintf("%s-%s", name, hex.EncodeToString(h[:])[:4]), nil
}

// ProjectSlug returns the slug this client talks to.
func (c *Client) ProjectSlug() string { return c.projectSlug }

// url builds a /api/<slug>/<path> URL with optional query parameters.
func (c *Client) url(path string, q url.Values) string {
	u := fmt.Sprintf("%s/api/%s/%s", c.base, c.projectSlug, strings.TrimPrefix(path, "/"))
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	return u
}

// getJSON issues a GET and decodes into out. Errors for non-2xx.
func (c *Client) getJSON(path string, q url.Values, out interface{}) error {
	resp, err := c.http.Get(c.url(path, q))
	if err != nil {
		return fmt.Errorf("GET %s: %w", path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GET %s: %s — %s", path, resp.Status, strings.TrimSpace(string(body)))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// getRaw returns the raw response body. Used for endpoints that return
// pre-formatted JSON (e.g. query).
func (c *Client) getRaw(path string, q url.Values) ([]byte, error) {
	resp, err := c.http.Get(c.url(path, q))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GET %s: %s — %s", path, resp.Status, strings.TrimSpace(string(body)))
	}
	return io.ReadAll(resp.Body)
}

// postJSON issues a POST with an application/json body.
func (c *Client) postJSON(path string, body, out interface{}) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return err
		}
	}
	resp, err := c.http.Post(c.url(path, nil), "application/json", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("POST %s: %s — %s", path, resp.Status, strings.TrimSpace(string(b)))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// --- Typed response shapes --------------------------------------------

type StatusResponse struct {
	Counts map[string]int `json:"counts"`
	Total  int            `json:"total"`
}

type HealthResponse struct {
	OK        bool   `json:"ok"`
	Version   string `json:"version"`
	UptimeSec int64  `json:"uptime_sec"`
	LastReqTS int64  `json:"last_request_ts"`
}

type HealthReport struct {
	Errors   []Finding `json:"errors"`
	Warnings []Finding `json:"warnings"`
	Hints    []Finding `json:"hints"`
	Entities int       `json:"entities"`
}

type Finding struct {
	Severity   string `json:"severity"`
	Category   string `json:"category"`
	Message    string `json:"message"`
	EntityKind string `json:"entity_kind,omitempty"`
	EntitySlug string `json:"entity_slug,omitempty"`
	EntityName string `json:"entity_name,omitempty"`
	Field      string `json:"field,omitempty"`
	Path       string `json:"path,omitempty"`
}

type EntitySummary struct {
	ID                string   `json:"id"`
	Kind              string   `json:"kind"`
	Name              string   `json:"name"`
	Slug              string   `json:"slug"`
	Description       string   `json:"description"`
	File              string   `json:"file"`
	RelationshipCount int      `json:"relationship_count"`
	Tags              []string `json:"tags,omitempty"`
	Statement         string   `json:"statement,omitempty"`
	RequirementStatus string   `json:"requirement_status,omitempty"`
	Source            string   `json:"source,omitempty"`
	SourceRef         string   `json:"source_ref,omitempty"`
}

type ReindexResponse struct {
	Mode          string   `json:"mode"`
	Indexed       int      `json:"indexed,omitempty"`
	Failed        []string `json:"failed,omitempty"`
	Entities      int      `json:"entities,omitempty"`
	Tags          int      `json:"tags,omitempty"`
	Relationships int      `json:"relationships,omitempty"`
	Words         int      `json:"words,omitempty"`
}

// --- Methods ----------------------------------------------------------

func (c *Client) Health() (*HealthResponse, error) {
	resp, err := c.http.Get(c.base + "/health")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out HealthResponse
	return &out, json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) Status() (*StatusResponse, error) {
	var out StatusResponse
	return &out, c.getJSON("status", nil, &out)
}

func (c *Client) List(kind, tag string) ([]EntitySummary, error) {
	q := url.Values{}
	if kind != "" {
		q.Set("kind", kind)
	}
	if tag != "" {
		q.Set("tag", tag)
	}
	var body struct {
		Entities []EntitySummary `json:"entities"`
		Count    int             `json:"count"`
	}
	if err := c.getJSON("entities", q, &body); err != nil {
		return nil, err
	}
	return body.Entities, nil
}

// Get returns the raw JSON for an entity — callers parse into the type
// they want. The server side uses query.FormatJSON which is a flat
// snapshot with every field.
func (c *Client) Get(slug string) ([]byte, error) {
	return c.getRaw("entity/"+slug, nil)
}

// GetRaw returns the raw markdown file bytes (frontmatter + body) for
// an entity. Lossless — unlike Get, this preserves every field because
// it bypasses query.FormatJSON and returns whatever FileStore.Save
// would have written. Used by the write-client Get path so mutate+PUT
// round-trips don't drop fields like affected_entities, created_at.
func (c *Client) GetRaw(slug string) ([]byte, error) {
	return c.getRaw("entity-raw/"+slug, nil)
}

// Query dispatches to a query mode on the server.
// mode = "" | "lookup" | "full" | "impacts" | "related-to" | "depends-on" | "depended-by" | "flow-components" | "filter"
// format = "" (json) | "rich" | "compact" | "refs"
// extra = optional query params (used by filter mode for kind/tag)
func (c *Client) Query(mode, slug, format string, extra url.Values) ([]byte, error) {
	q := url.Values{}
	q.Set("mode", mode)
	if slug != "" {
		q.Set("slug", slug)
	}
	if format != "" {
		q.Set("format", format)
	}
	for k, vs := range extra {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	return c.getRaw("query", q)
}

func (c *Client) Validate() (*HealthReport, error) {
	var out HealthReport
	return &out, c.getJSON("validate", nil, &out)
}

func (c *Client) SyncCheck(strict bool) (*HealthReport, error) {
	q := url.Values{}
	if strict {
		q.Set("strict", "true")
	}
	var out HealthReport
	return &out, c.getJSON("sync-check", q, &out)
}

func (c *Client) Context() ([]byte, error) {
	return c.getRaw("context", nil)
}

// Search uses the existing /search?q= endpoint backed by the index
// full-text engine.
func (c *Client) Search(q string) ([]byte, error) {
	qv := url.Values{}
	qv.Set("q", q)
	return c.getRaw("search", qv)
}

func (c *Client) ConstraintsCheck(path string) ([]byte, error) {
	q := url.Values{}
	q.Set("path", path)
	return c.getRaw("constraints-check", q)
}

// Constraints fetches the active-constraints list (decisions) for
// the current project.
func (c *Client) Constraints() ([]byte, error) {
	return c.getRaw("constraints", nil)
}

func (c *Client) FilesOrphans() ([]string, error) {
	var body struct {
		Orphans []string `json:"orphans"`
	}
	if err := c.getJSON("files/orphans", nil, &body); err != nil {
		return nil, err
	}
	return body.Orphans, nil
}

func (c *Client) FilesCoverage(path string) ([]byte, error) {
	q := url.Values{}
	if path != "" {
		q.Set("path", path)
	}
	return c.getRaw("files/coverage", q)
}

// Reindex is the write-side companion. CLI write commands save markdown
// via FileStore then call Reindex(paths) to refresh just those entries.
// Pass full=true (and nil paths) for a full rebuild.
func (c *Client) Reindex(paths []string, full bool) (*ReindexResponse, error) {
	body := map[string]interface{}{"paths": paths, "full": full}
	var out ReindexResponse
	return &out, c.postJSON("reindex", body, &out)
}

// WriteResponse is the metadata returned after a create/update.
type WriteResponse struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Kind     string `json:"kind"`
	FilePath string `json:"file_path"`
}

// CreateEntity sends a new entity to syded. `frontmatter` is the YAML
// serialization of the typed entity struct (same bytes FileStore.Save
// would write), `body` is the markdown body. syded allocates the ID +
// slug, writes the file, and updates the index in one transaction.
func (c *Client) CreateEntity(kind, frontmatter, body string) (*WriteResponse, error) {
	payload := map[string]string{"kind": kind, "frontmatter": frontmatter, "body": body}
	var out WriteResponse
	return &out, c.postJSON("entity", payload, &out)
}

// UpdateEntity sends a full entity replacement. The entity's ID + slug
// must already be set — syded looks them up to locate the file.
func (c *Client) UpdateEntity(kind, frontmatter, body string) (*WriteResponse, error) {
	// PUT /entity uses the same handler as POST on syded — just the
	// method differs so we forward it through the client here.
	var out WriteResponse
	return &out, c.putJSON("entity", map[string]string{
		"kind":        kind,
		"frontmatter": frontmatter,
		"body":        body,
	}, &out)
}

// DeleteEntity removes an entity by kind + slug.
func (c *Client) DeleteEntity(kind, slug string) error {
	req, err := http.NewRequest(http.MethodDelete, c.url("entity/"+kind+"/"+slug, nil), nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("DELETE entity/%s/%s: %s — %s", kind, slug, resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

func (c *Client) putJSON(path string, body, out interface{}) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return err
		}
	}
	req, err := http.NewRequest(http.MethodPut, c.url(path, nil), &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("PUT %s: %s — %s", path, resp.Status, strings.TrimSpace(string(b)))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
