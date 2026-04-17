# syde CLI Commands

## Entity CRUD

### syde add <kind> <name> [flags]
Create an entity. Kind must be one of: `system`, `component`, `contract`, `concept`, `flow`, `requirement`, `plan`, `task`

**Base flags (all kinds):**
```
--description     short one-sentence identification (required except plan/task and requirements with --statement)
--purpose         why it exists
--tag             labels (repeatable)
--file            concrete source file paths, no wildcards (repeatable)
--body            markdown body content
--note            append informal note (repeatable)
--add-rel         add relationship: "target-slug:type" (repeatable)
```

**System flags:** (a system = a standalone process/app/service)
```
--context-text       system context
--scope              system scope
--design-principles  design principles
--quality-goals      quality goals
--assumptions        assumptions
```
Top-level systems have no `belongs_to`. Sub-systems (another standalone
process nested inside a project) belong_to their parent system:
`--add-rel "<parent-system-slug>:belongs_to"`. System nesting must be acyclic.

**Component flags:** (purpose is base flag, required)
```
--responsibility       what it does (required)
--capability           one capability it provides (repeatable, at least one required)
--boundaries           what it does NOT do
--behavior-summary     behavior summary
--interaction-summary  how it interacts with others
--data-handling        data handling description
--scaling-notes        scaling notes
```
Components MUST belong_to a system: `--add-rel "<system-slug>:belongs_to"`

**Contract flags:** (one contract per endpoint/command/event/screen)
```
--contract-kind       rest|cli|event|rpc|graphql|websocket|pubsub|storage|screen
--interaction-pattern sync|async|request-response|pub-sub|streaming|schema|render
--input               invocation signature, REQUIRED (e.g. "GET /api/projects", "syde plan create <name>", "/concepts")
--input-parameter     parameter spec "path|type|description" (repeatable, REQUIRED at least one)
--output              output signature / response shape, REQUIRED (e.g. "200 OK application/json", "rendered UI")
--output-parameter    parameter spec "path|type|description" (repeatable, REQUIRED at least one)
--protocol-notes      protocol notes
--constraints-text    constraints
--versioning-notes    versioning notes
--wireframe           UIML source, REQUIRED when --contract-kind=screen.
                      Validator parses it; dashboard renders it via the
                      dark-mode wireframe renderer (charcoal-on-dark, region
                      badges, ✕-rect placeholders, line-bar text). Use full
                      UIML vocabulary including attributes:
                        <screen direction="horizontal">
                        <layout direction="horizontal">
                        <grid cols="4">
                        <sidebar width="200">
                        <panel width="360">
                        <item active="true">
                        <metric label="X" value="N">
                      The classic inbox pattern is
                        <screen direction="horizontal">
                          <sidebar/><panel width="360"/><main/>
                        </screen>
```

Render any screen contract from the terminal:

```
syde wireframe render <contract-slug>            # html to stdout
syde wireframe render <slug> --format ascii      # plain text
syde wireframe render <slug> --format image --out /tmp/x.png
syde wireframe render <slug> --out /tmp/x.html --open
```
Contract **name** should be a descriptive noun phrase like "List Projects",
not the raw invocation. The raw invocation goes in `--input`.
Contracts MUST belong_to a **system** (the process exposing them): `--add-rel "<system-slug>:belongs_to"`
Contracts SHOULD reference the implementing component: `--add-rel "<component-slug>:references"`
Contracts SHOULD reference concepts used in input/output: `--add-rel "<concept-slug>:references"`

**Concept flags:**
```
--meaning           what it means in the domain (REQUIRED — ERROR if missing)
--invariants        rules that must always hold (recommended — Finding if empty)
--lifecycle         lifecycle description (state machine, etc.) — optional
```

Concepts are **glossary entries**, not ERD tables. They carry no
attributes, actions, data sensitivity, or structure notes — schema
detail belongs on the implementing component or on a contract schema.

Concepts MUST belong_to a system: `--add-rel "<system-slug>:belongs_to"`

Concepts SHOULD carry **role-based** relationships describing how the
term is realised elsewhere in the model:

| Relationship    | Target kind | Meaning |
|-----------------|-------------|---------|
| `implemented_by`| component   | Component that implements the concept in code |
| `exposed_via`   | contract    | Contract that exposes the concept externally |
| `used_in`       | flow        | Flow that operates on or produces the concept |
| `relates_to`    | concept     | Another concept this one relates to (no cardinality label) |

Example:
```
syde add concept "Order" \
  --description "A customer purchase request" \
  --meaning "Groups line items into a billable transaction with a single delivery and payment." \
  --lifecycle "draft → placed → paid → shipped → delivered" \
  --invariants "Total > 0; must have ≥1 line item; forward-only state transitions." \
  --add-rel "ecommerce:belongs_to" \
  --add-rel "order-service:implemented_by" \
  --add-rel "place-order:exposed_via" \
  --add-rel "place-order-flow:used_in" \
  --add-rel "customer:relates_to" \
  --add-rel "line-item:relates_to"
```

**Flow flags:**
```
--trigger            what starts this flow
--goal               what this flow achieves
--narrative          step-by-step narrative
--happy-path         happy path description
--edge-cases         edge cases
--failure-modes      failure modes
--performance-notes  performance notes
```

**Requirement flags:** (statements are save-validated against EARS patterns)
```
--statement            EARS shall-form text (REQUIRED). Must match one of:
                         Ubiquitous:        "The <subject> shall <action>."
                         Event-driven:      "When <trigger>, the <subject> shall <action>."
                         State-driven:      "While <state>, the <subject> shall <action>."
                         Unwanted-behavior: "If <unwanted>, then the <subject> shall <action>."
                         Optional-feature:  "Where <feature>, the <subject> shall <action>."
--type                 functional|non-functional|constraint|interface|performance|security|usability  (REQUIRED)
--priority             must|should|could|wont  (MoSCoW, REQUIRED)
--verification         short description of how fulfillment is verified (REQUIRED)
--source               user|plan|migration|manual
--source-ref           prompt/plan/issue/migration reference
--requirement-status   active|superseded|obsolete
--rationale            why this requirement exists
--supersedes           comma-separated requirement refs this one replaces
--superseded-by        comma-separated newer requirement refs
--obsolete-reason      required when obsolete
--approved-at          capture/approval timestamp
--add-rel              "target:refines" or "parent-req:derives_from" (repeatable)
```
Requirements MUST belong_to a system: `--add-rel "<system-slug>:belongs_to"`.
Requirements **never** carry a `--file` list — they are pure design intent.
There is no `--acceptance` flag on `syde add requirement` (that flag still
applies to `syde task create`); use `--verification` instead.

See `references/requirement-derivation.md` for the deterministic algorithm
to backfill EARS requirements from existing components / contracts /
concepts / systems.

### syde get <id-or-slug>
Show full entity details.

### syde list [kind] [--tag X]
List entities. Optionally filter by kind or tag.

### syde update <slug> [flags]
Update an entity. Supports all kind-specific flags from `syde add`, plus:
```
--add-rel      add relationship: "target-slug:type" (repeatable)
--remove-rel   remove relationship by target-slug (repeatable)
--add-tag      add a tag
--remove-tag   remove a tag
```

### syde remove <slug> [--force]
Delete an entity. Prompts for confirmation unless `--force`. Requirements are
append-only: mark them superseded or obsolete with `syde update` instead.

### syde search <query>
Full-text search across all entities.

## Query

### syde query [slug] [flags]
**The single entry point for understanding any file, symbol, or entity.**
syde query is not "the design tool" — it is the unified context surface
where architecture (the entity index) and code (the source files) meet.
Every query also reveals whether architecture and code are in sync.

**Five access paths:**

1. **Entity lookup** — `syde query <slug>` or `syde query <slug> --full`
2. **Keyword search** (entity prose) — `syde query --search "<terms>" [filters]`
3. **Code search** (source files) — `syde query --code "<pattern>"`
4. **File → entities + content** — `syde query --file <path> [--content] [--no-related]`
5. **Graph walks** — `--impacts`, `--related-to`, `--depends-on`,
   `--depended-by`, `--flow --components`

Plus filter listings (`--kind`, `--tag` with no slug and no search term).

```
--full                include body and all related data
--kind <kind>         filter by entity kind (combines with --search/--file)
--tag <tag>           filter by tag (combines with --search)
--format rich|json|compact|refs   output format (default: rich)

Entity-prose search:
--search "text"       full-text search across name, description, purpose,
                      notes, and body (for component/contract/concept/
                      flow/system — plan/task bodies are
                      excluded because they churn). Tokens are
                      AND-merged by default; the engine auto-broadens to
                      OR if AND yields zero hits and labels the results
                      "broadened". CamelCase and snake_case identifiers
                      are split into sub-tokens by the indexer.
--any                 OR-merge tokens explicitly (skips the AND attempt)
--limit N             cap search results (0 = unbounded for --search;
                      default 50 for --code)

Code search (every tracked source file, owner annotated):
--code "<pattern>"    literal-string search across non-ignored files in
                      the summary tree. Uses ripgrep when available,
                      falls back to a Go walker. Each hit carries its
                      owning entity (or a ⚠ orphan warning if none).

File → entities:
--file <path>         return owners of a source file (exact match), or
                      every owner under a directory prefix if the path
                      is a directory or has no exact match. Surfaces a
                      ⚠ DRIFT banner when the file has no owner.
--content             also inline the file body (capped at 100KB) so a
                      single call answers both "who owns this?" and
                      "what does it say?". Only works in exact-match mode.
--no-related          omit one-hop related entities from --file results

Graph walks:
--impacts <slug>      transitive impact analysis
--flow <slug> --components   flow decomposition with component details
--related-to <slug>          all direct connections
--depends-on <slug>          entities this depends on
--depended-by <slug>         entities that depend on this
--diff <slug> --since 7d     git change history
```

**The three-question checklist** — run before any Grep / Read:

1. **What entity owns this?** → `syde query --file <path>` or `syde query <slug>`
2. **What does syde know about this term?** → `syde query --search "<term>"`
3. **What code references it?** → `syde query --code "<symbol>"`

Three syde calls beat one wrong grep. The owning-entity annotation on
every code hit is the architecture↔code bridge — you don't have to ask
"is this file mapped?" separately, the answer arrives with the search.

**Drift signals to act on immediately:**

- `--file <path>` reports `⚠ DRIFT: no owning entities` → map the file
  with `syde update <component> --file ...` or `syde tree ignore <path>`.
- `--code <pattern>` reports a hit on a file marked `⚠ orphan` → same fix.

Both warnings are the model telling you architecture has rotted relative
to the code. They are part of the workflow, not a separate cleanup step.

**Search & discovery cookbook** — every recipe an agent needs:

```bash
# 1. SYMBOL LOOKUP — find a Go identifier in source, with owners.
syde query --code ConceptEntity
syde query --code "func NewStore("

# 2. FILE READ WITH FRAMING — content + owners + related, one call.
syde query --file internal/storage/index.go --content

# 3. FILE PARTICIPATION ONLY — same as #2 without the body.
syde query --file internal/storage/index.go

# 4. DIRECTORY SWEEP — owners under a path prefix.
syde query --file internal/cli/

# 5. ENTITY-PROSE SEARCH — name, description, purpose, notes, body.
#    Multi-word AND by default; auto-broadens to OR if zero hits.
syde query --search "BadgerDB index"
syde query --search "concept entity"   # broadened example

# 6. SCOPED SEARCH — narrow by kind and/or tag.
syde query --search migration --kind requirement --tag critical --limit 5

# 7. ENTITY DETAIL — full context for one entity.
syde query storage-engine --full

# 8. LIST BY KIND — every entity of a given kind.
syde query --kind concept
syde query --kind requirement --format refs

# 9. ORPHAN TRIAGE — find files the design model does not claim.
syde files orphans
syde query --file <orphan-path>   # to decide map vs. ignore

# 10. GRAPH WALKS — impact analysis and dependency tracing.
syde query --impacts storage-engine
syde query --depends-on query-engine
syde query --depended-by query-engine
syde query --related-to query-engine

# 11. RECENT ACTIVITY ON ONE ENTITY — git history scoped to its files.
syde query --diff storage-engine --since 7d
```

**Tools to avoid for tracked files**: `Grep` and `Read` should be
reserved for files the summary tree intentionally ignores (vendor/,
node_modules/, generated assets, .git/, build artifacts, binaries). If
a file is tracked by `syde tree scan`, use `syde query` — bypassing it
silently disconnects architecture from code.

Each search hit carries the matched tokens, the field they came from,
and a 120-char snippet centered on the first match, so you can triage
without opening the markdown. Each code hit carries its owning entity.

## Plans

### syde plan create <name> [flags]
Create a new plan (status: draft).
```
--background  why this plan exists (context, motivation)
--objective   what success looks like
--scope       in-scope / out-of-scope summary
```

### syde plan update <slug> [flags]
Update plan-level fields. Only changed flags are applied.
```
--background  update background
--objective   update objective
--scope       update scope
--description short description
--purpose     why it exists
```

### syde plan add-phase <plan-slug> [flags]
Add a phase to a plan. Phases can be nested using `--parent`. Returns auto-generated phase ID. Phases no longer hold entity drafts — use `--notes` to describe new entities the phase will require, then list them in tasks' `--affected-entity` after they're created via `syde add`.
Before approval, each phase must also have at least one granular task created
with `syde task create --plan <plan-slug> --phase <phase-id> ...`.
```
--name           short phase label (shown in tree view)
--parent         parent phase ID for nesting (e.g., phase_1)
--description    phase description
--objective      what this phase achieves (success condition)
--changes        concrete list of what changes in this phase
--details        implementation walkthrough (HOW to build)
--notes          risks, reminders, new entities the phase will introduce
```

### syde plan estimate <slug>
Check plan size. Recommends splitting if >25 tasks (~5 commands per task).

### syde plan show <slug> [--full]
Show plan as tree. Parent phases show aggregated task counts from children.
`--full` expands phase details/notes and per-task objective/details/acceptance.

### syde plan approve <slug>
Mark plan as approved. Required before implementation. Approval creates a
plan-sourced requirement and links the plan to it; repeated approval reuses the
existing plan requirement instead of creating a duplicate.

Approval validates phase task coverage. If any phase has zero direct tasks,
approval fails and prints the empty phase IDs. Add concrete tasks with
`syde task create --plan <plan-slug> --phase <phase-id> ...`, show the updated
plan, and ask for user approval again before re-running approval.

### syde plan phase <plan-slug> <phase-id> [flags]
Update a plan phase. All fields are optional — only changed flags are applied.
```
--name           update phase name
--parent         set/change parent phase
--status         pending|in_progress|completed|skipped
--description    update phase description
--objective      update phase objective (what it achieves)
--changes        update phase changes (what concretely changes)
--details        update implementation details
--notes          update notes/context
```

### syde plan execute <slug>
Transition plan to in-progress. Plan must be approved first. Does NOT create entities — use `syde add` during task implementation.
Also handles legacy phases with `EntityKind` + `EntityName` fields.

### syde plan sync [--file <path>]
Import from most recent `.claude/plans/*.md` file into syde.

### syde plan list
List all plans with progress.

## Tasks

### syde task create <name> [flags]
```
--plan              linked plan slug
--phase             linked plan phase ID (e.g., phase_1)
--priority          high|medium|low (default: medium)
--objective         what this task achieves
--details           implementation specifics (files, approach)
--acceptance        how to know it's done (observable)
--affected-entity   existing entity slug this task will modify (repeatable). Validator rejects unknown slugs.
--affected-file     concrete source file path this task will touch (repeatable). Must exist in .syde/tree.yaml.
--entity            (legacy) linked entity slug/ID (prefer --affected-entity)
```

### syde task update <slug> [flags]
Update task fields. Only changed flags are applied.
```
--objective       update objective
--details         update details
--acceptance      update acceptance criteria
--priority        update priority
--description     update description
--affected-entity replace affected entities list (repeatable)
--affected-file   replace affected files list (repeatable)
```

### syde task start <slug>
Mark task as in_progress.

### syde task done <slug>
Mark task as completed. Automatically:
  1. If linked to a plan phase, auto-marks the phase completed (if all its tasks are done).
  2. Bumps `updated_at` on every entity in `affected_entities` AND every entity whose `files` overlap `affected_files`. This clears the drift validator warning so subsequent `syde validate` runs show the entities as reviewed.

### syde task block <slug> [--reason "..."]
Mark task as blocked.

### syde task list [--status X]
List tasks. Status values: `pending`, `in_progress`, `completed`, `blocked`, `cancelled`

### syde task sub <parent-slug> <name>
Create a subtask under a parent task.

### syde task link <task-slug> <entity-slug>
Link a task to an entity.

## Constraints

### syde constraints [--json]
Show active requirements.

### syde constraints check <file> [--json]
Map source file to component via `component_paths` in `syde.yaml`, then show applicable requirements. Returns `{}` if file is not mapped to any component.

## Summary Tree

syde maintains a file/folder summary tree in `.syde/tree.yaml` mirroring
the source directory. Every file and folder carries a human-written
summary; syde tracks SHA-256 hashes so changed files are marked stale
and the stale bit cascades to ancestor folders. `.gitignore` is
respected automatically, plus built-in defaults (`.git`, `.syde`,
`.claude`, `node_modules`, `vendor`, `dist`, `build`, `__pycache__`,
`target`, `*.lock`, `*.sum`, etc.) and optional `tree_ignore` in `syde.yaml`.

### syde tree scan
Walk the project, diff against `.syde/tree.yaml`, mark changed/new files stale,
cascade stale up to ancestor folders, remove deleted entries (marking parent
stale). Prints `added: N, changed: N, deleted: N, stale: N`. Always run this
at session start and after any source-file changes.

### syde tree status [--strict]
Counts total files, dirs, stale files, stale dirs, last scan time. With
`--strict`, exits non-zero if anything is stale — used by git pre-commit hooks,
CI, and the Claude Code SessionEnd hook to block "leaving the tree dirty".

### syde tree changes [--format plain|json] [--leaves-only]
List all stale paths. Sorted deepest-first so you summarize children before
parents. `--leaves-only` filters out stale folders whose descendants are still
stale — summarize those leaves first, then the folders bubble up naturally.

### syde tree summarize <path> --summary "..."
Set a node's summary text. Clears its stale bit and marks its direct parent
stale (cascade continues on each subsequent summarize). Pass `--summary -` to
read summary text from stdin for long inputs.

### syde tree show [path] [--full] [--max-depth N] [--stale]
Pretty ASCII tree with inline summaries. Default: root, depth 2. `--full`
shows everything. `--stale` prefixes stale entries with `!`.

### syde tree get <path>
Print just that node's summary (for piping into other tools).

### syde tree ignore <path>
Mark a tree node as ignored. Ignored nodes:
  - Stop appearing in `syde tree changes` (no summary required)
  - Are exempt from the "orphan file: not referenced by any entity" validator error
  - Stay tracked (hash + mtime still updated on scan) so you can unignore them later
Use for generated files, large fixtures, vendor drops, and anything that's committed but not part of the design model.

### syde tree unignore <path>
Remove the ignored flag, marking the node stale so the next summarize pass picks it up.

### syde tree context <path> [--format plain|json] [--no-content] [--max-bytes N]
**The go-to call for understanding a file.** Bundles in one response:
1. The **tree summary** — breadcrumb of ancestor folder summaries from root
   down to the file's parent (architectural framing)
2. The **file summary** — the node's own stored summary
3. The **file content** — inlined raw bytes (files only), capped at 64 KiB
   by default; pass `--max-bytes` to adjust or `--no-content` to omit.

For folders, returns the breadcrumb + folder summary + a listing of direct
children with their summaries. **Use this instead of naive `Read` calls when
creating syde entities from existing source files** — the breadcrumb tells
you which sub-system / folder the file belongs to, and the cached folder
summaries give you architectural context without re-reading every sibling.

## Other

```
syde context [--json]          # Full architecture snapshot (auto-loaded at session start)
syde status                    # Entity counts
syde validate                  # Check model integrity
syde reindex                   # Rebuild BadgerDB index from markdown files
syde graph [entity] [--format dot]   # Relationship graph (ASCII or Graphviz DOT)
syde sync [--dry-run --coverage --check]  # Sync/audit design model against codebase
syde memory sync/list/clean    # Claude Code memory sync
syde open                      # Start dashboard + open browser
syde install-skill             # Install skill files into .claude/
```
