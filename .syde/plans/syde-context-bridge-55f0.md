---
approved_at: "2026-04-14T09:21:14Z"
background: 'syde was supposed to be the single context surface bridging architecture and code, but in practice agents (including me, this session) keep falling back to Grep/Read for source-code questions. Three reasons: (1) syde query --search only indexes entity markdown so ''where is ConceptEntity?'' returns nothing, (2) the AND-by-default tokenizer with no CamelCase splitting silently kills loose human queries like ''Relationship Label'' or ''add-rel'', (3) the skill positions syde as ''design tool'' alongside grep instead of as the unified entry point for both. Result: agents bypass syde for code lookups, miss the architecture/code drift signal, and the model rots silently.'
completed_at: "2026-04-14T09:32:26Z"
created_at: "2026-04-14T09:19:02Z"
id: PLN-0006
kind: plan
name: syde-context-bridge
objective: syde is the unambiguous, sole entry point for understanding any tracked file, symbol, or entity. Source-code questions resolve through syde query --code. File reads happen via syde query --file --content. Loose human queries succeed via CamelCase/snake_case-aware tokenization and an AND-then-OR fallback. The skill explicitly forbids Grep/Read for anything syde can answer, and explains the architecture↔code sync feedback loop so agents understand why bypassing syde is harmful, not just inefficient.
phases:
    - changes: internal/query/engine.go gets SearchCodeOptions + Engine.SearchCode returning CodeHit list (path, line, snippet, owning entity if any). internal/query/engine.go ByFile gains a withContent flag and ByFileResult gains Content (string) + ContentLines (int). internal/dashboard/api_readall.go handleQueryAPI adds case 'code' and a content query param on by-file. internal/cli/query.go adds --code <pattern> and --content flags.
      description: Engine.SearchCode (rg + Go fallback) and --content flag for ByFile, exposed via CLI and dashboard
      details: 'SearchCode pipeline: build the file set from the summary tree (every file node, skipping ignored). Try exec.LookPath(''rg''); if found, run ''rg --line-number --no-heading --color never -F <pattern> -- <file...>'' (literal by default; add --regex flag later). If rg missing, walk the tree files in Go and use a simple Contains check per line (slow but correct). For each hit, look up the owning entity via audit.FileCoverage so the result carries entity name+slug. Cap hits at --limit (default 50). ByFile --content: when withContent && exact match, read the file (size cap 100KB to avoid runaway responses) and attach. Truncate beyond cap with a note.'
      id: phase_1
      name: Code search + content inlining
      notes: rg fallback is intentional — install-on-demand UX is hostile in CI / fresh clones. The Go walker handles it correctly, just slower.
      objective: syde query can answer 'where is this symbol used?' and 'what does this file say?' without ever touching Grep or Read
      status: completed
      tasks:
        - enginesearchcode-with-rg-fallback-to-go-walker
        - byfile-content-inlines-file-body-with-size-cap
        - cli-dashboard-wiring-for-code-and-content
    - changes: internal/storage/indexer.go tokenize() now also emits sub-tokens from CamelCase splits and snake_case/dash splits, while preserving the original concatenated form. internal/storage/index.go IndexSchemaVersion bumped to 3. internal/query/engine.go Engine.Search retries the query with Any=true when AND yields zero hits, and labels each broadened SearchHit with a Broadened bool flag. internal/query/formatter.go FormatSearchHits notes 'broadened (no exact match for all tokens)' in the output when Broadened is set.
      description: CamelCase / snake_case tokenization (IndexSchemaVersion v3) and AND→OR fallback in Engine.Search
      details: 'Tokenizer: keep current alnum-only split, then for each token also emit CamelCase parts (regex on uppercase boundaries) and snake/dash parts. Skip duplicates and stop words. Bump IndexSchemaVersion to 3 — NewStore auto-reindexes on next open. Search OR fallback: if !opts.Any and len(filtered)==0 after the AND pass, retry with opts.Any=true, mark every resulting hit Broadened=true. Limit still applies.'
      id: phase_2
      name: Tokenizer hardening + OR fallback
      notes: The IndexSchemaVersion bump triggers a one-shot reindex on next syded launch — same auto-heal pattern as v1→v2 last session.
      objective: Loose human queries like 'ConceptEntity', 'add-rel', 'relationship label' return relevant hits instead of zero
      status: completed
      tasks:
        - camelcase-snakecase-tokenizer-split-with-indexschemaversion-v3
        - enginesearch-andor-fallback-with-broadened-flag
    - changes: skill/SKILL.md 'Finding Files to Read' section reframed as 'Getting Context — syde first, always'. New subsection on the architecture↔code sync feedback loop (every query also surfaces drift). Explicit instruction to forbid Grep/Read for tracked files. skill/references/commands.md cookbook expanded to 10+ recipes covering --code, --content, --kind listing, orphan triage, drift detection, broadened search, file→owner reverse lookup, recent activity scoping.
      description: Reframe SKILL.md and commands.md so syde is the sole context surface for both architecture AND source code, with the architecture/code sync feedback loop made explicit
      details: 'SKILL.md changes: (a) rename section, (b) lead with the three-question checklist (what entity owns this? what does syde know about this term? what code references it?), (c) document the drift signal — ''no owners'' or ''code in untracked file'' = act on it now, (d) explicit forbid-Grep-on-tracked-files line. commands.md: rewrite cookbook section with new recipes, document --code, --content, broadened-results behaviour, and the three-question pattern. Update the syde query Long help in internal/cli/query.go to match.'
      id: phase_3
      name: Skill rewrite — syde as the context bridge
      notes: 'Tone matters: this is not ''syde is preferred'', it is ''syde is the entry point and bypassing it loses information''. Make that argument once, clearly.'
      objective: Agents reading the skill default to syde for every read/explore step, understand that bypassing syde silently disconnects architecture from code, and have copy-pasteable recipes for every common question
      status: completed
      tasks:
        - skill-rewrite-framing-syde-as-the-context-bridge
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) Engine.SearchCode using rg if available with a Go-native walker fallback, scoped to summary-tree files; (2) --content flag for ByFile that inlines file body alongside owners + related; (3) CamelCase/snake_case tokenizer split with IndexSchemaVersion bump to v3 (auto reindex); (4) AND→OR fallback in Engine.Search with broadened-results labeling; (5) CLI + dashboard wiring for --code, --content; (6) skill rewrite framing syde as the context bridge, expanded cookbook with 10+ recipes, and explicit Grep/Read prohibition for tracked files. Out-of-scope: hard-blocking Read via PreToolUse hooks; vendor-importing ripgrep; semantic / vector search; LSP-style symbol indexing; React ERD view (was a separate user request the user paused on); changing the entity model.'
slug: syde-context-bridge-55f0
source: manual
updated_at: "2026-04-14T09:32:26Z"
---
