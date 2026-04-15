---
approved_at: "2026-04-14T06:07:45Z"
background: 'Four CLI pain points surfaced while auditing orphan files: (1) syde validate claims to fail on orphan files but doesn''t — the promise is broken, (2) no machine-readable (JSON) output on list/query/get/validate/sync --check so any audit must parse markdown directly, (3) syde and syded can''t coexist because both try to take the BadgerDB directory lock — every CLI command while dashboard is running requires pkill syded, (4) sync --check and validate overlap without either being authoritative, so ''did I finish properly'' is ambiguous. The skill''s session-end guarantees depend on a single, trustworthy health command.'
completed_at: "2026-04-14T06:24:45Z"
created_at: "2026-04-14T06:02:11Z"
id: PLN-0001
kind: plan
name: CLI Health & Daemon Coexistence
objective: 'One command (syde sync check) is the authoritative session-end health gate: structural integrity, file orphans, missing fields, drift, and descriptions all checked, exits non-zero on any gap. syded reads without holding the BadgerDB lock so CLI and dashboard coexist. All read commands emit --format json. SKILL.md Phase 5 points at the one canonical finish command.'
phases:
    - changes: New internal/audit/audit.go with Run(store, tree) returning []Finding; new internal/audit/orphans.go; syde validate rewired; new syde files parent command + files orphans + files coverage subcommands
      description: Shared internal/audit/ package that both validate and sync --check call into; make validate actually fail on orphan files
      details: 'Define Finding{Severity, Category, EntityOrPath, Message}. Categories: orphan_file, missing_field, broken_rel, cycle, drift, short_description, missing_contract, missing_rel. Implement FindOrphans walking tree.yaml non-ignored files vs union of entity.files. Implement FileCoverage returning map[path][]ownerSlug. Rewire internal/cli/validate.go to call audit.Run and print findings grouped by severity, exit 1 if any Error. Add internal/cli/files.go with parent files cmd + files orphans + files coverage. Reuse tree loader from internal/tree/store.go.'
      id: phase_1
      name: Audit Core + Orphan Enforcement
      notes: Must preserve current validate error messages for backward compat (relationship cycles, missing required fields). Orphan check is additive.
      objective: syde validate exits non-zero when any non-ignored tree file has no owning entity; audit logic lives in one place
      status: completed
      tasks:
        - rewire-syde-validate-to-call-auditrun
        - create-internalaudit-package-with-finding-run
        - implement-orphan-coverage-checks-in-internalaudit
        - add-syde-files-parent-cmd-orphans-coverage
    - changes: syde sync check calls audit.Run with all categories; syde validate becomes alias for syde sync check --errors-only with deprecation notice; Stop/SessionEnd hook updated; SKILL.md Phase 5 rewritten
      description: Make syde sync check the single authoritative health gate; deprecate syde validate to a thin alias
      details: 'Add internal/cli/sync.go check subcommand (or extend existing sync --check flag into real subcommand). Output: grouped Errors/Warnings/Hints sections with counts, color, and a final ''N errors, M warnings'' line. Exit codes: 0 clean; 1 any error; 2 --strict and any warning. validate.go now prints ''syde validate is deprecated, use syde sync check'' on stderr. skill/hooks.json SessionEnd runs syde sync check --strict. SKILL.md Phase 5 single bullet: ''run syde sync check --strict — it enforces validate + orphans + drift + tree-status in one shot''.'
      id: phase_2
      name: Consolidate Health Commands into syde sync check
      notes: Keep both entrypoints working for one release — don't break existing muscle memory
      objective: One command answers 'can I finish the session?'; SKILL.md Phase 5 points at it; Stop hook runs it
      status: completed
      tasks:
        - deprecate-validate-to-alias-of-sync-check-errors-only
        - wire-stop-hook-to-syde-sync-check-strict
        - add-syde-sync-check-subcommand-as-canonical-health-gate
    - changes: New internal/cli/output.go with Emit(any, format); list/query/get/status/validate/sync check all call Emit; references/commands.md updated
      description: Unify --format json across list, query, get, status, validate, sync check so scripted audits don't need markdown parsing
      details: 'Shared schema: {kind, data, findings?, meta{count, elapsed_ms}}. For list: {kind: ''list'', data: []EntitySummary}. For query: {kind:''query'', data: EntityDetail}. For findings-carrying commands (validate, sync check): {kind:''health'', data: {errors, warnings, hints}}. Keep existing rich/compact formats. Wire --format flag on commands that don''t have it (list is currently --format compact only; add json).'
      id: phase_3
      name: JSON Output Across Read Commands
      notes: Don't break existing default rich output. --format json must be lossless — every field that rich shows is in the JSON.
      objective: Every read command supports --format json with a consistent schema; documented in references/commands.md
      status: completed
      tasks:
        - add-format-json-to-list-query-get-status
        - create-internalclioutputgo-shared-emit-formatter
        - json-output-for-validate-sync-check
    - changes: New internal/storage/memstore.go implementing the Store read interface backed by an in-memory map; internal/dashboard/registry.go GetStore switched to memstore; fsnotify watcher rebuilds memstore incrementally on .syde/*.md change; syded never calls badger.Open
      description: syded stops opening BadgerDB so CLI + daemon coexist without lock contention
      details: 'Extract Store reader interface (Get, List, Query, GetInbound, etc.) so memstore can satisfy it without re-implementing filestore. memstore walks .syde/<kind>/*.md on startup, builds per-kind map + reverse relationship index. Subscribe to fsnotify on .syde/ and rebuild affected entries (full rebuild is acceptable for v1 — project sizes are small). Dashboard API handlers already go through registry.GetStore — swap the impl, handlers unchanged. Write path (dashboard POSTs) stays blocked or goes through CLI exec for v1 — dashboard is read-only while CLI writes. Tradeoff: slightly stale reads between fsnotify events; acceptable.'
      id: phase_4
      name: syded BadgerDB-free Read Mode
      notes: BadgerDB remains the CLI's source of fast indexed lookup; memstore is a syded-only read path. No schema change to .syde/ files.
      objective: Running syded never blocks syde CLI; lock contention errors are gone; dashboard reads are served from an in-memory index rebuilt from markdown
      status: completed
      tasks:
        - extract-store-reader-interface-for-memstore-impl
        - switch-syded-registry-to-memstore-fsnotify-rebuild
        - implement-internalstoragememstorego
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'IN: internal/audit/ package for orphan + completeness checks; rewire syde validate + sync --check to share it; add syde files orphans and syde files coverage subcommands; JSON output across list, query, get, status, validate, sync; consolidate validate into sync check (validate becomes a thin alias with deprecation notice); make syded use an in-memory index rebuilt from markdown via fsnotify instead of BadgerDB so the daemon never touches the dir lock; update SKILL.md Phase 5 + hooks to run syde sync check --strict. OUT: new GUI features; changes to entity schemas; dashboard API changes beyond read-path; gitignore parsing for tree (already deferred).'
slug: cli-health-daemon-coexistence-p25w
source: manual
updated_at: "2026-04-14T06:24:45Z"
---
