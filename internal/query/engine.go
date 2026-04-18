package query

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/feedloop/syde/internal/audit"
	"github.com/feedloop/syde/internal/graph"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
	"github.com/feedloop/syde/internal/utils"
)

// Engine executes rich queries against the store.
type Engine struct {
	Store *storage.Store
}

// NewEngine creates a query engine.
func NewEngine(store *storage.Store) *Engine {
	return &Engine{Store: store}
}

// Lookup returns a fully resolved entity by slug.
func (e *Engine) Lookup(slug string) (*ResolvedEntity, error) {
	return Resolve(e.Store, slug)
}

// Filter returns entities matching kind/tag filters.
func (e *Engine) Filter(kind model.EntityKind, tag string) ([]EntitySummary, error) {
	var kinds []model.EntityKind
	if kind != "" {
		kinds = []model.EntityKind{kind}
	} else {
		kinds = model.AllEntityKinds()
	}

	var results []EntitySummary
	for _, k := range kinds {
		entities, err := e.Store.List(k)
		if err != nil {
			continue
		}
		for _, ewb := range entities {
			b := ewb.Entity.GetBase()
			if tag != "" && !hasTag(b.Tags, tag) {
				continue
			}

			relCount := len(b.Relationships)
			canonSlug := b.CanonicalSlug()
			desc := b.Description

			rels := make([]SummaryRelation, 0, len(b.Relationships))
			for _, r := range b.Relationships {
				rels = append(rels, SummaryRelation{Type: r.Type, Target: r.Target, Label: r.Label})
			}
			summary := EntitySummary{
				ID:            b.ID,
				Kind:          string(b.Kind),
				Name:          b.Name,
				Slug:          canonSlug,
				Description:   desc,
				File:          e.Store.FS.RelativePath(b.Kind, canonSlug),
				RelCount:      relCount,
				Tags:          b.Tags,
				Files:         b.Files,
				Relationships: rels,
			}
			if ct, ok := ewb.Entity.(*model.ContractEntity); ok {
				summary.ContractKind = ct.ContractKind
				summary.InteractionPattern = ct.InteractionPattern
			}
			if p, ok := ewb.Entity.(*model.PlanEntity); ok {
				summary.PlanStatus = string(p.PlanStatus)
				summary.UpdatedAt = b.UpdatedAt
			}
			if req, ok := ewb.Entity.(*model.RequirementEntity); ok {
				if summary.Description == "" {
					summary.Description = req.Statement
				}
				summary.Statement = req.Statement
				summary.RequirementStatus = string(req.RequirementStatus)
				summary.Source = req.Source
				summary.SourceRef = req.SourceRef
			}
			results = append(results, summary)
		}
	}
	return results, nil
}

// EntitySummary is a compact entity listing.
type EntitySummary struct {
	ID            string   `json:"id"`
	Kind          string   `json:"kind"`
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Description   string   `json:"description"`
	File          string   `json:"file"`
	RelCount      int      `json:"relationship_count"`
	Tags          []string `json:"tags,omitempty"`
	Files         []string `json:"files,omitempty"`
	// Outbound relationships only — enough for client-side filtering by
	// belongs_to / depends_on / references / etc.
	Relationships []SummaryRelation `json:"relationships,omitempty"`
	// Contract-specific fields surfaced in the list view so the
	// filter bar can narrow the Contracts page by kind and pattern
	// without a second per-entity fetch.
	ContractKind       string `json:"contract_kind,omitempty"`
	InteractionPattern string `json:"interaction_pattern,omitempty"`
	// Requirement-specific fields surfaced in list/API output so
	// traceability views can inspect requirement status without a
	// second per-entity fetch.
	Statement         string `json:"statement,omitempty"`
	RequirementStatus string `json:"requirement_status,omitempty"`
	Source            string `json:"source,omitempty"`
	SourceRef         string `json:"source_ref,omitempty"`
	// Plan-specific fields surfaced in the list view so the frontend
	// can sort by date and filter by status without a detail fetch.
	PlanStatus string `json:"plan_status,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

// SummaryRelation is a flat outbound relationship reference. Label
// carries the free-form annotation from Relationship.Label. Empty
// when no label was set at creation time.
type SummaryRelation struct {
	Type   string `json:"type"`
	Target string `json:"target"`
	Label  string `json:"label,omitempty"`
}

// Impacts returns transitive impact analysis grouped by hop distance.
type ImpactResult struct {
	EntityID   string                  `json:"entity_id"`
	EntityName string                  `json:"entity_name"`
	Hops       map[int][]EntitySummary `json:"hops"`
	Total      int                     `json:"total"`
}

func (e *Engine) Impacts(slug string, maxDepth int) (*ImpactResult, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	hops, err := graph.ImpactAnalysis(e.Store.Idx, b.ID, maxDepth)
	if err != nil {
		return nil, err
	}

	result := &ImpactResult{
		EntityID:   b.ID,
		EntityName: b.Name,
		Hops:       make(map[int][]EntitySummary),
	}

	for depth, nodes := range hops {
		for _, node := range nodes {
			summary := EntitySummary{ID: node.ID}
			if kind, ok := utils.ParseIDKind(node.ID); ok {
				if ref, err := e.Store.Idx.LookupByID(kind, node.ID); err == nil {
					summary.Kind = string(ref.Kind)
					summary.Name = ref.Name
					summary.Slug = utils.Slugify(ref.Name)
					summary.File = ref.File
				}
			}
			result.Hops[depth] = append(result.Hops[depth], summary)
			result.Total++
		}
	}
	return result, nil
}

// FlowDecomposition shows a flow with its participating components.
type FlowDecomposition struct {
	FlowName   string          `json:"flow_name"`
	Trigger    string          `json:"trigger"`
	Goal       string          `json:"goal"`
	Components []EntitySummary `json:"components"`
	Contracts  []EntitySummary `json:"contracts"`
	Concepts   []EntitySummary `json:"concepts"`
}

func (e *Engine) FlowComponents(slug string) (*FlowDecomposition, error) {
	entity, _, err := e.Store.GetByKind(model.KindFlow, slug)
	if err != nil {
		return nil, err
	}
	f := entity.(*model.FlowEntity)

	result := &FlowDecomposition{
		FlowName: f.Name,
		Trigger:  f.Trigger,
		Goal:     f.Goal,
	}

	for _, rel := range f.Relationships {
		kind, ok := utils.ParseIDKind(rel.Target)
		if !ok {
			continue
		}
		ref, err := e.Store.Idx.LookupByID(kind, rel.Target)
		if err != nil {
			continue
		}
		summary := EntitySummary{
			ID:   ref.ID,
			Kind: string(ref.Kind),
			Name: ref.Name,
			Slug: utils.Slugify(ref.Name),
			File: ref.File,
		}
		switch kind {
		case model.KindComponent:
			result.Components = append(result.Components, summary)
		case model.KindContract:
			result.Contracts = append(result.Contracts, summary)
		case model.KindConcept:
			result.Concepts = append(result.Concepts, summary)
		}
	}
	return result, nil
}

// SearchHit is a single search result with enough context for an
// agent to decide whether the entity is worth opening: the exact
// token that matched, the field it came from, and a short snippet of
// the surrounding text. Broadened is set when the result came from
// the AND→OR fallback path — i.e. no entity matched every token, so
// Search retried with OR semantics. Agents should treat broadened
// results as relevant-but-loose: review the matched_tokens field to
// see which subset of the query actually hit.
type SearchHit struct {
	ID        string   `json:"id"`
	Kind      string   `json:"kind"`
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	File      string   `json:"file"`
	Field     string   `json:"field,omitempty"`
	Word      string   `json:"word,omitempty"`
	Snippet   string   `json:"snippet,omitempty"`
	Matched   []string `json:"matched_tokens,omitempty"`
	Broadened bool     `json:"broadened,omitempty"`
}

// SearchOptions configures Engine.Search. Zero-value Limit means
// unbounded; Any=false performs AND-intersection across tokens;
// Kind/Tag narrow the candidate set before scoring.
type SearchOptions struct {
	Query string
	Kind  model.EntityKind
	Tag   string
	Any   bool
	Limit int
}

// Search runs a filtered full-text search and returns one SearchHit
// per surviving entity, ranked by the number of distinct tokens that
// matched (descending). AND is the default — a multi-word query only
// returns entities that matched every token — flip Any=true to OR.
//
// Snippets come from the entity's description, purpose, or body
// (whichever has the first occurrence of a matched token). This lets
// agents skim results without opening the markdown files.
func (e *Engine) Search(opts SearchOptions) ([]SearchHit, error) {
	tokensMap, err := e.Store.Idx.SearchTokens(opts.Query)
	if err != nil {
		return nil, err
	}
	if len(tokensMap) == 0 {
		return nil, nil
	}

	// Group by entity ID: collect every (token, field, word) hit per
	// entity so we can both score (# distinct tokens) and pick a
	// snippet source (the first field a token appeared in).
	type entityHits struct {
		ref     storage.FileRef
		tokens  map[string]bool
		fields  map[string]string // field → first word seen there
		primary storage.WordRef
	}
	buckets := make(map[string]*entityHits)
	for token, hits := range tokensMap {
		for _, h := range hits {
			b, ok := buckets[h.Ref.ID]
			if !ok {
				b = &entityHits{
					ref:     h.Ref,
					tokens:  map[string]bool{},
					fields:  map[string]string{},
					primary: h,
				}
				buckets[h.Ref.ID] = b
			}
			b.tokens[token] = true
			if _, exists := b.fields[h.Field]; !exists {
				b.fields[h.Field] = h.Word
			}
		}
	}

	tokenCount := len(tokensMap)

	// Apply AND-intersection, kind filter, and tag filter in one pass.
	var tagHits map[string]bool
	if opts.Tag != "" {
		tagHits = make(map[string]bool)
		if refs, terr := e.Store.Idx.ListByTag(opts.Tag); terr == nil {
			for _, r := range refs {
				tagHits[r.ID] = true
			}
		}
	}

	// applyFilters returns the buckets surviving kind/tag and the
	// requested AND/OR token rule. Extracted as a closure so we can
	// run it once with AND, then re-run with OR if the first pass
	// returns nothing.
	applyFilters := func(requireAll bool) []*entityHits {
		out := make([]*entityHits, 0, len(buckets))
		for _, b := range buckets {
			if requireAll && len(b.tokens) < tokenCount {
				continue
			}
			if opts.Kind != "" && b.ref.Kind != opts.Kind {
				continue
			}
			if opts.Tag != "" && !tagHits[b.ref.ID] {
				continue
			}
			out = append(out, b)
		}
		return out
	}

	requireAll := !opts.Any
	filtered := applyFilters(requireAll)

	// AND→OR fallback: if the strict pass yielded nothing AND the
	// caller did not already ask for OR, broaden silently and mark
	// every resulting hit so the renderer can flag the relaxation.
	// This rescues loose human queries like "relationship label"
	// that no single entity contains in full but several entities
	// match partially.
	broadened := false
	if requireAll && len(filtered) == 0 {
		filtered = applyFilters(false)
		broadened = len(filtered) > 0
	}

	// Rank by distinct tokens matched (desc), then by name for a
	// stable order when ties occur.
	sort.SliceStable(filtered, func(i, j int) bool {
		if len(filtered[i].tokens) != len(filtered[j].tokens) {
			return len(filtered[i].tokens) > len(filtered[j].tokens)
		}
		return filtered[i].ref.Name < filtered[j].ref.Name
	})

	if opts.Limit > 0 && len(filtered) > opts.Limit {
		filtered = filtered[:opts.Limit]
	}

	// Build SearchHit with snippet. Snippet priority: description,
	// purpose, name, notes, body — whichever field the word was
	// indexed from, we load the entity and slice a window around the
	// first match.
	hits := make([]SearchHit, 0, len(filtered))
	for _, b := range filtered {
		hit := SearchHit{
			ID:        b.ref.ID,
			Kind:      string(b.ref.Kind),
			Name:      b.ref.Name,
			Slug:      slugFromFilePath(b.ref.File),
			File:      b.ref.File,
			Field:     b.primary.Field,
			Word:      b.primary.Word,
			Broadened: broadened,
		}
		matched := make([]string, 0, len(b.tokens))
		for t := range b.tokens {
			matched = append(matched, t)
		}
		sort.Strings(matched)
		hit.Matched = matched
		hit.Snippet = e.buildSnippet(b.ref, matched)
		hits = append(hits, hit)
	}
	return hits, nil
}

// slugFromFilePath mirrors storage.slugFromFile — stripping the
// directory and .md extension. Duplicated here because the query
// package cannot reach the unexported storage helper.
func slugFromFilePath(file string) string {
	base := file
	if idx := strings.LastIndex(file, "/"); idx >= 0 {
		base = file[idx+1:]
	}
	return strings.TrimSuffix(base, ".md")
}

// buildSnippet loads the on-disk markdown for a hit and returns a
// 120-character window centered on the first matched token. Falls
// back to the entity description if nothing matches in the body.
func (e *Engine) buildSnippet(ref storage.FileRef, tokens []string) string {
	entity, body, err := e.Store.GetByKind(ref.Kind, slugFromFilePath(ref.File))
	if err != nil {
		return ""
	}
	b := entity.GetBase()

	// Search description first (it's the agent-friendly elevator
	// pitch), then body. Pick the earliest match across all tokens.
	haystacks := []string{b.Description, b.Purpose, body}
	for _, hay := range haystacks {
		if hay == "" {
			continue
		}
		low := strings.ToLower(hay)
		bestIdx := -1
		for _, tok := range tokens {
			if i := strings.Index(low, tok); i >= 0 {
				if bestIdx < 0 || i < bestIdx {
					bestIdx = i
				}
			}
		}
		if bestIdx >= 0 {
			return snippetAround(hay, bestIdx, 120)
		}
	}
	return b.Description
}

// snippetAround returns up to `width` characters centered on `idx`
// within `text`, collapsing whitespace/newlines and adding ellipses
// when the window is a true substring.
func snippetAround(text string, idx, width int) string {
	// Normalize whitespace so markdown newlines don't ruin the window.
	clean := strings.Join(strings.Fields(text), " ")
	// Re-find the approximate position in the cleaned string. We
	// cannot use the raw idx because whitespace collapsing shifted it.
	low := strings.ToLower(clean)
	// Seed with the first token that still appears in the clean text.
	start := idx
	if start > len(clean) {
		start = 0
	}
	_ = low
	half := width / 2
	s := start - half
	if s < 0 {
		s = 0
	}
	e2 := s + width
	if e2 > len(clean) {
		e2 = len(clean)
	}
	out := clean[s:e2]
	if s > 0 {
		out = "… " + out
	}
	if e2 < len(clean) {
		out = out + " …"
	}
	return out
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

// FullContext is alias for Lookup — returns everything.
func (e *Engine) FullContext(slug string) (*ResolvedEntity, error) {
	return e.Lookup(slug)
}

// RelatedTo returns all directly connected entities.
func (e *Engine) RelatedTo(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	seen := make(map[string]bool)
	var results []EntitySummary

	// Outbound
	for _, rel := range b.Relationships {
		if seen[rel.Target] {
			continue
		}
		seen[rel.Target] = true
		if kind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, rel.Target); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, Slug: utils.Slugify(ref.Name), File: ref.File})
			}
		}
	}

	// Inbound
	inbound, _ := e.Store.Idx.GetInbound(b.ID)
	for _, in := range inbound {
		if seen[in.Source] {
			continue
		}
		seen[in.Source] = true
		if kind, ok := utils.ParseIDKind(in.Source); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, in.Source); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, Slug: utils.Slugify(ref.Name), File: ref.File})
			}
		}
	}

	return results, nil
}

// DependsOn returns entities this slug depends on.
func (e *Engine) DependsOn(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	var results []EntitySummary
	for _, rel := range b.Relationships {
		if rel.Type != model.RelDependsOn {
			continue
		}
		if kind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, rel.Target); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, Slug: utils.Slugify(ref.Name), File: ref.File})
			}
		}
	}
	return results, nil
}

// DependedBy returns entities that depend on this slug.
func (e *Engine) DependedBy(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	var results []EntitySummary
	inbound, _ := e.Store.Idx.GetInbound(b.ID)
	for _, in := range inbound {
		if in.Type != model.RelDependsOn {
			continue
		}
		if kind, ok := utils.ParseIDKind(in.Source); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, in.Source); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, Slug: utils.Slugify(ref.Name), File: ref.File})
			}
		}
	}
	return results, nil
}

// RefinedBy returns active requirements that refine the named component
// (via a `refines` relationship). The inbound index uses IDs while YAML
// targets may be slugs, so fall back to a full scan of requirement
// relationships when inbound lookup misses — mirrors the four-way
// alias tolerance used by the audit engine.
func (e *Engine) RefinedBy(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()
	aliases := map[string]bool{
		b.ID:                              true,
		b.CanonicalSlug():                 true,
		utils.BaseSlug(b.CanonicalSlug()): true,
		utils.Slugify(b.Name):             true,
	}

	all, err := e.Store.ListAll()
	if err != nil {
		return nil, err
	}
	var results []EntitySummary
	seen := map[string]bool{}
	for _, ewb := range all {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok || req.RequirementStatus != model.RequirementActive {
			continue
		}
		for _, rel := range req.GetBase().Relationships {
			if rel.Type != model.RelRefines || !aliases[rel.Target] {
				continue
			}
			if seen[req.ID] {
				break
			}
			seen[req.ID] = true
			rb := req.GetBase()
			results = append(results, EntitySummary{
				ID:   rb.ID,
				Kind: string(rb.Kind),
				Name: rb.Name,
				Slug: rb.CanonicalSlug(),
			})
			break
		}
	}
	return results, nil
}

// unused import guard
var _ = fmt.Sprintf

// ByFileResult answers "what does this source file participate in?".
// Owners are the entities whose --file list contains (or prefix-
// matches) the given path; Related is the one-hop neighborhood
// around those owners (outbound relationships + inbound
// relationships, de-duped by ID, owners themselves excluded).
type ByFileResult struct {
	Path             string          `json:"path"`
	Mode             string          `json:"mode"` // "exact" or "prefix"
	Owners           []EntitySummary `json:"owners"`
	Related          []EntitySummary `json:"related,omitempty"`
	Content          string          `json:"content,omitempty"`
	ContentBytes     int             `json:"content_bytes,omitempty"`
	ContentTruncated bool            `json:"content_truncated,omitempty"`
}

// SearchCodeOptions configures Engine.SearchCode. Pattern is matched
// literally (no regex by default). Limit caps the result set.
type SearchCodeOptions struct {
	Pattern string
	Limit   int
}

// CodeHit is one source-file match returned by SearchCode. The owner
// fields are populated when the file is mapped to an entity in the
// design model — that is the architecture↔code bridge: every hit
// arrives with its owning component (or "" if the file is an orphan).
type CodeHit struct {
	Path      string `json:"path"`
	Line      int    `json:"line"`
	Snippet   string `json:"snippet"`
	OwnerKind string `json:"owner_kind,omitempty"`
	OwnerSlug string `json:"owner_slug,omitempty"`
	OwnerName string `json:"owner_name,omitempty"`
}

// CodeSearchResult bundles the hit list with metadata about how the
// search ran (whether ripgrep was used, how many files were scanned).
type CodeSearchResult struct {
	Pattern      string    `json:"pattern"`
	Engine       string    `json:"engine"` // "rg" or "go"
	FilesScanned int       `json:"files_scanned"`
	Hits         []CodeHit `json:"hits"`
}

// maxCodeHits is the default limit when SearchCodeOptions.Limit is 0.
const maxCodeHits = 50

// maxByFileContentBytes caps the inlined file content for ByFile
// --content. 100KB is enough for almost any source file while
// preventing pathological binaries or generated assets from
// flooding the response.
const maxByFileContentBytes = 100 * 1024

// projectRoot returns the absolute path of the project root — the
// parent of FileStore.Root, which itself points at .syde/. Tree
// paths are stored relative to the project root, so any os.Open or
// exec.Command that needs to resolve a tracked file must join
// against this, not against FS.Root.
func (e *Engine) projectRoot() string {
	return filepath.Dir(e.Store.FS.Root)
}

// SearchCode runs a literal-string pattern search across every file
// in the summary tree (skipping ignored entries), uses ripgrep when
// available for speed, and falls back to a Go-native walker so
// fresh clones / CI environments without rg still get correct
// results. Each hit is annotated with the entity that owns the file
// (when one exists) so agents see the architectural framing alongside
// the code match.
func (e *Engine) SearchCode(opts SearchCodeOptions) (*CodeSearchResult, error) {
	if strings.TrimSpace(opts.Pattern) == "" {
		return &CodeSearchResult{Pattern: opts.Pattern, Engine: "noop"}, nil
	}
	limit := opts.Limit
	if limit <= 0 {
		limit = maxCodeHits
	}

	t, err := tree.Load(e.Store.FS.Root)
	if err != nil {
		return nil, fmt.Errorf("load tree: %w", err)
	}
	files := make([]string, 0, len(t.Nodes))
	for p, n := range t.Nodes {
		if n.Type != tree.TypeFile || n.Ignored {
			continue
		}
		files = append(files, p)
	}
	sort.Strings(files)

	all, _ := e.Store.ListAll()
	cov := audit.FileCoverage(all, t)
	annotate := func(h *CodeHit) {
		owners, ok := cov[h.Path]
		if !ok || len(owners) == 0 {
			return
		}
		o := owners[0]
		h.OwnerKind = string(o.Kind)
		h.OwnerSlug = o.Slug
		h.OwnerName = o.Name
	}

	result := &CodeSearchResult{Pattern: opts.Pattern, FilesScanned: len(files)}

	if rgPath, lookErr := exec.LookPath("rg"); lookErr == nil {
		hits, runErr := e.runRipgrep(rgPath, opts.Pattern, files, limit)
		if runErr == nil {
			result.Engine = "rg"
			for i := range hits {
				annotate(&hits[i])
			}
			result.Hits = hits
			return result, nil
		}
		// rg blew up — fall through to the Go walker so the user
		// still gets an answer.
	}

	hits := e.scanFilesGo(opts.Pattern, files, limit)
	for i := range hits {
		annotate(&hits[i])
	}
	result.Engine = "go"
	result.Hits = hits
	return result, nil
}

// runRipgrep shells out to ripgrep with literal-match flags and
// parses its line-numbered output. Pattern is treated literally
// (no regex) so callers do not accidentally hit metacharacter
// surprises like '.' or '('. Returns the parsed hits or an error;
// an exit-code-1 from ripgrep means "no matches" and is returned as
// an empty slice without error.
func (e *Engine) runRipgrep(rgPath, pattern string, files []string, limit int) ([]CodeHit, error) {
	if len(files) == 0 {
		return nil, nil
	}
	args := []string{
		"--line-number",
		"--no-heading",
		"--color", "never",
		"--max-count", fmt.Sprintf("%d", limit),
		"-F",
		"--",
		pattern,
	}
	args = append(args, files...)
	cmd := exec.Command(rgPath, args...)
	cmd.Dir = e.projectRoot()
	out, err := cmd.Output()
	if err != nil {
		// Exit code 1 means "no matches" — that is success for us.
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, err
	}
	hits := make([]CodeHit, 0)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		// rg format: "path:line:content"
		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}
		var lineNo int
		fmt.Sscanf(parts[1], "%d", &lineNo)
		hits = append(hits, CodeHit{
			Path:    parts[0],
			Line:    lineNo,
			Snippet: strings.TrimSpace(parts[2]),
		})
		if len(hits) >= limit {
			break
		}
	}
	return hits, nil
}

// scanFilesGo is the rg-free fallback. It walks the file list in
// alphabetical order, opens each one, and scans line by line for
// the literal pattern. Slower than rg but dependency-free so fresh
// clones / restricted CI environments still get usable results.
func (e *Engine) scanFilesGo(pattern string, files []string, limit int) []CodeHit {
	hits := make([]CodeHit, 0)
	root := e.projectRoot()
	for _, rel := range files {
		abs := filepath.Join(root, rel)
		f, err := os.Open(abs)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(f)
		scanner.Buffer(make([]byte, 64*1024), 1024*1024)
		lineNo := 0
		for scanner.Scan() {
			lineNo++
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				hits = append(hits, CodeHit{
					Path:    rel,
					Line:    lineNo,
					Snippet: strings.TrimSpace(line),
				})
				if len(hits) >= limit {
					f.Close()
					return hits
				}
			}
		}
		f.Close()
	}
	return hits
}

// ByFile resolves a source file or directory prefix to the entities
// that own it plus their one-hop neighbors. The match is exact when
// the path is a concrete file; if the caller passes a directory
// (trailing slash) or a path with no exact owner, ByFile falls back
// to a prefix scan across the tree and returns the union of owners
// for every matched file.
//
// withContent, when true and the match is exact, also inlines the
// file's content (capped at maxByFileContentBytes) so a single call
// answers both "what does this file participate in?" and "what does
// it actually say?" — the architecture↔code bridge in one round-trip.
func (e *Engine) ByFile(path string, withRelated bool) (*ByFileResult, error) {
	return e.ByFileWith(path, withRelated, false)
}

// ByFileWith is the extended form of ByFile that also accepts a
// withContent flag. The two-arg ByFile stays for backward
// compatibility; new callers use this one.
func (e *Engine) ByFileWith(path string, withRelated, withContent bool) (*ByFileResult, error) {
	all, err := e.Store.ListAll()
	if err != nil {
		return nil, err
	}
	t, err := tree.Load(e.Store.FS.Root)
	if err != nil {
		// No tree yet — fall back to walking entity --file lists
		// directly. Prefix matching degrades gracefully.
		return e.byFileWithoutTree(path, all, withRelated), nil
	}
	cov := audit.FileCoverage(all, t)

	result := &ByFileResult{Path: path, Mode: "exact"}
	ownerSet := make(map[string]bool) // keyed by entity ID for de-dup
	ownerEntities := []EntitySummary{}
	addOwner := func(o audit.FileOwner) {
		key := string(o.Kind) + "/" + o.Slug
		if ownerSet[key] {
			return
		}
		ownerSet[key] = true
		ownerEntities = append(ownerEntities, EntitySummary{
			Kind: string(o.Kind),
			Name: o.Name,
			Slug: o.Slug,
			File: "", // path varies per owner, not shown
		})
	}

	if owners, ok := cov[path]; ok && len(owners) > 0 {
		for _, o := range owners {
			addOwner(o)
		}
	} else {
		// Prefix match: treat path as a directory scope. Normalize
		// trailing slash so "internal/cli" and "internal/cli/" both
		// work the same way.
		prefix := strings.TrimSuffix(path, "/")
		result.Mode = "prefix"
		for p, owners := range cov {
			if p == prefix || strings.HasPrefix(p, prefix+"/") {
				for _, o := range owners {
					addOwner(o)
				}
			}
		}
	}
	result.Owners = ownerEntities

	if withRelated && len(result.Owners) > 0 {
		result.Related = e.oneHopRelated(result.Owners, ownerSet)
	}
	if withContent && result.Mode == "exact" {
		e.attachContent(result, path)
	}
	return result, nil
}

// attachContent reads the file at path (relative to the project
// root), caps it at maxByFileContentBytes, and populates the
// Content / ContentBytes / ContentTruncated fields on the result.
// Silent on errors — missing file just leaves Content empty.
func (e *Engine) attachContent(result *ByFileResult, path string) {
	abs := filepath.Join(e.projectRoot(), path)
	info, err := os.Stat(abs)
	if err != nil || info.IsDir() {
		return
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		return
	}
	result.ContentBytes = len(data)
	if len(data) > maxByFileContentBytes {
		data = data[:maxByFileContentBytes]
		result.ContentTruncated = true
	}
	result.Content = string(data)
}

// byFileWithoutTree is the fallback path when the summary tree is
// missing or unreadable. We walk every entity and check if the given
// path matches any of its --file entries directly — exact match when
// possible, prefix match otherwise.
func (e *Engine) byFileWithoutTree(path string, all []model.EntityWithBody, withRelated bool) *ByFileResult {
	result := &ByFileResult{Path: path, Mode: "exact"}
	ownerSet := make(map[string]bool)
	prefix := strings.TrimSuffix(path, "/")
	exact := false

	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		for _, fp := range b.Files {
			if fp == path {
				exact = true
				key := string(b.Kind) + "/" + b.CanonicalSlug()
				if ownerSet[key] {
					continue
				}
				ownerSet[key] = true
				result.Owners = append(result.Owners, EntitySummary{
					ID: b.ID, Kind: string(b.Kind), Name: b.Name,
					Slug: b.CanonicalSlug(),
				})
			}
		}
	}
	if !exact {
		result.Mode = "prefix"
		for _, ewb := range all {
			b := ewb.Entity.GetBase()
			for _, fp := range b.Files {
				if fp == prefix || strings.HasPrefix(fp, prefix+"/") {
					key := string(b.Kind) + "/" + b.CanonicalSlug()
					if ownerSet[key] {
						continue
					}
					ownerSet[key] = true
					result.Owners = append(result.Owners, EntitySummary{
						ID: b.ID, Kind: string(b.Kind), Name: b.Name,
						Slug: b.CanonicalSlug(),
					})
				}
			}
		}
	}
	if withRelated && len(result.Owners) > 0 {
		result.Related = e.oneHopRelated(result.Owners, ownerSet)
	}
	return result
}

// oneHopRelated expands a set of owners by one step in the
// relationship graph: every outbound target and every inbound source
// of every owner, resolved to EntitySummary, de-duped by kind/slug,
// excluding the owners themselves.
func (e *Engine) oneHopRelated(owners []EntitySummary, ownerSet map[string]bool) []EntitySummary {
	seen := make(map[string]bool)
	for k := range ownerSet {
		seen[k] = true
	}
	var related []EntitySummary
	// addByTarget resolves a relationship target that may be stored as
	// either an entity ID (COM-0004) or a slug (full "storage-engine-
	// ahgm" / bare "storage-engine"). The raw index only knows IDs via
	// ParseIDKind, so if that fails we fall back to Store.Get which
	// handles every slug form.
	addByTarget := func(target string) {
		if kind, ok := utils.ParseIDKind(target); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, target); err == nil {
				slug := slugFromFilePath(ref.File)
				key := string(ref.Kind) + "/" + slug
				if seen[key] {
					return
				}
				seen[key] = true
				related = append(related, EntitySummary{
					ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name,
					Slug: slug, File: ref.File,
				})
				return
			}
		}
		ent, _, err := e.Store.Get(target)
		if err != nil {
			return
		}
		tb := ent.GetBase()
		key := string(tb.Kind) + "/" + tb.CanonicalSlug()
		if seen[key] {
			return
		}
		seen[key] = true
		related = append(related, EntitySummary{
			ID: tb.ID, Kind: string(tb.Kind), Name: tb.Name,
			Slug: tb.CanonicalSlug(),
			File: e.Store.FS.RelativePath(tb.Kind, tb.CanonicalSlug()),
		})
	}
	for _, o := range owners {
		// Need the full entity to read outbound rels and its aliases
		// for inbound lookups — the summary we built does not carry
		// them.
		ent, _, err := e.Store.GetByKind(model.EntityKind(o.Kind), o.Slug)
		if err != nil {
			continue
		}
		b := ent.GetBase()
		for _, rel := range b.Relationships {
			addByTarget(rel.Target)
		}
		// Inbound lookups must match every form the index might have
		// used (ID, full slug, bare slug), same as the resolver.
		aliases := []string{b.ID}
		if b.Slug != "" {
			aliases = append(aliases, b.Slug)
			if base := utils.BaseSlug(b.Slug); base != "" && base != b.Slug {
				aliases = append(aliases, base)
			}
		}
		for _, alias := range aliases {
			inbound, _ := e.Store.Idx.GetInbound(alias)
			for _, in := range inbound {
				addByTarget(in.Source)
			}
		}
	}
	return related
}
