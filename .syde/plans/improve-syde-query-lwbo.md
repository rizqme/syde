---
approved_at: "2026-04-14T08:10:33Z"
background: 'Agents rely on syde query as the primary way to find project info, but --search is broken: the inverted index stores only {file,field} so hits return empty identity; search ignores --kind/--tag; there is no file→entities query mode; CLI bypasses the format/filter pipeline via c.Search(). Agents fall back to naive Read and grep, burning context.'
completed_at: "2026-04-14T08:25:54Z"
created_at: "2026-04-14T08:06:11Z"
id: PLN-0004
kind: plan
name: improve-syde-query
objective: Agents can find relevant entities via syde query with three reliable access paths — keyword (filtered by kind/tag, with snippets), file path (owners + one-hop related), and rich format — all in one command. Index auto-heals on schema bumps so upgrades are seamless.
phases:
    - changes: 'internal/storage/index.go (w: value shape, schema version key, IndexWords signature, Search return type); syded startup path triggers full reindex when the schema version is missing or older than current'
      description: 'Store full FileRef in w: values, tokenize body content, auto-reindex on schema bump'
      details: Add const IndexSchemaVersion=2; on Store.Open read key 'meta:schema' — if missing/stale, call Reindex and write current version. Rewrite IndexWords to store JSON {FileRef, field, word}. Change Search to return []SearchHit (kind, id, name, slug, file, field, word) instead of []FileRef. Tokenize entity body markdown in indexer.go for component/contract/concept/decision/flow/learning (not plan/task which churn). Keep word tokenizer but bump to include 3-char words (agents want specificity).
      id: phase_1
      name: Index schema v2 + auto reindex
      notes: Existing dev DBs will auto-reindex once on first syded launch — harmless, one-shot. The old legacy /search endpoint must keep working; it just calls the same new engine path.
      objective: Searching any indexed word returns a complete FileRef (id, kind, name, slug, file) plus the field/excerpt it matched, and fresh installs auto-rebuild without user action
      status: completed
      tasks:
        - add-schema-version-key-and-auto-reindex-on-mismatch
        - reshape-w-index-values-to-carry-full-fileref-field-word
        - tokenize-entity-body-markdown-in-the-indexer
    - changes: 'internal/query/engine.go: new SearchOptions struct (Query, Kind, Tag, Any, Limit); new ByFileResult struct + Engine.ByFile; SearchHit gains Field, Word, Snippet; Engine.Search applies kind/tag filter, AND-merge default, OR via Any, Limit truncates; internal/query/resolver.go gains a helper for one-hop expansion around a set of owners'
      description: Engine-level API supporting filtered search with snippets and file→entities with one-hop expansion
      details: 'Search pipeline: tokenize → per-token prefix scan → intersect (AND) or union (OR) → filter by kind/tag → load entity descriptions for snippets → limit. ByFile pipeline: resolve path via storage.FileCoverage → if path ends with / or has no exact owner, prefix-match tree nodes → union owners → if withRelated, add one-hop outbound + inbound (belongs_to, depends_on, references, exposes) de-duped by ID. Snippets: 60-char window around first match in description or body, with ellipses.'
      id: phase_2
      name: 'Query engine: filtered search + ByFile'
      notes: Keep Filter(kind,tag) intact — Search is a separate method, not a replacement. ByFile should cheap-path when the index already has a file→entity reverse map; if it does not, fall back to scanning FileCoverage from storage.
      objective: Engine exposes Search(opts) and ByFile(path, withRelated) returning typed results the CLI and dashboard can render in any format
      status: completed
      tasks:
        - enginesearch-with-filters-andor-snippets-limit
        - enginebyfile-with-exactprefix-match-and-one-hop-expansion
    - changes: 'internal/cli/query.go: drop c.Search(); add queryFile/queryLimit/queryAny/queryNoRelated flags; --search routes mode=search; --file routes mode=by-file. internal/client: add Query(mode,...) support for new modes. internal/dashboard/api_readall.go: handleQueryAPI adds case ''search'' and case ''by-file'', honoring kind/tag/limit/any/with_related; api.go handleSearchAPI becomes a thin wrapper that calls the same engine with the legacy response shape for back-compat.'
      description: Route all query paths through one handler; add --file, --limit, --any, --no-related
      details: 'CLI flags: --file STR, --limit INT (default 0 = unbounded for lookups, 20 for search), --any, --no-related. Format=rich prints a compact list with snippets; refs prints ''kind/name file:line''; json prints full SearchHit structs. For --full + --search, expand the top hit into a ResolvedEntity automatically.'
      id: phase_3
      name: CLI + dashboard wiring
      notes: Verify the legacy /search HTTP endpoint still works (it is referenced by the React SPA). The React app gets no changes in this plan.
      objective: syde query --search respects --kind/--tag/--format/--limit; syde query --file <path> returns owners + related entities; no code path bypasses the shared pipeline
      status: completed
      tasks:
        - dashboard-handlequeryapi-add-search-by-file-modes
        - cli-query-flags-and-single-pipeline-routing
    - changes: skill/SKILL.md (expand 'Finding Files to Read' with the new recipes); skill/references/commands.md (add flag rows for --file/--limit/--any/--no-related and a 'Search & discovery' cookbook section)
      description: Teach the agent the new query surface so it uses the CLI instead of naive Read/grep
      details: 'Examples to include: (a) ''find all components mentioning badger'' → syde query --search badger --kind component, (b) ''what owns this file and what does it connect to'' → syde query --file internal/storage/index.go, (c) ''narrow to tag + keyword'' → syde query --search migration --tag critical --limit 5, (d) ''broad directory scope'' → syde query --file internal/cli/. Stress that agents should prefer these over Grep/Read for architecture-level lookups.'
      id: phase_4
      name: Skill documentation
      notes: Skill files are embedded via go:embed in skill/embed.go — no regen step needed, just edit the .md and rebuild.
      objective: SKILL.md 'Finding Files to Read' section and references/commands.md accurately describe --search filters, --file, --limit, --any, --no-related with worked examples
      status: completed
      tasks:
        - update-skillmd-and-commandsmd-with-new-query-surface
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) bump index schema version with auto-reindex on mismatch, (2) re-shape w: values with full FileRef + field + body tokenization, (3) Engine search honors filters/AND/OR/limit/snippets, (4) Engine.ByFile with one-hop expansion, (5) CLI query --search routed through Query pipeline, new --file/--limit/--any/--no-related flags, (6) skill docs refreshed. Out-of-scope: new top-level CLI commands, dashboard UI, tree changes, rewriting legacy /search endpoint payload shape.'
slug: improve-syde-query-lwbo
source: manual
updated_at: "2026-04-14T08:25:54Z"
---
