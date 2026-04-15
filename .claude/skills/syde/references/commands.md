# syde CLI Commands

## Entity CRUD

### syde add <kind> <name> [flags]
Create an entity. Kind must be one of: `system`, `component`, `contract`, `concept`, `flow`, `decision`, `plan`, `task`, `design`, `learning`

**Base flags (all kinds):**
```
--description     short one-sentence identification (REQUIRED on every kind — validator errors when missing)
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

**Contract flags:** (one contract per endpoint/command/event)
```
--contract-kind       rest|cli|event|rpc|graphql|websocket|pubsub
--interaction-pattern sync|async|request-response|pub-sub|streaming
--input               invocation signature, REQUIRED (e.g. "GET /api/projects", "syde plan create <name>")
--input-parameter     parameter spec "path|type|description" (repeatable, REQUIRED at least one)
--output              output signature / response shape, REQUIRED (e.g. "200 OK application/json")
--output-parameter    parameter spec "path|type|description" (repeatable, REQUIRED at least one)
--protocol-notes      protocol notes
--constraints-text    constraints
--versioning-notes    versioning notes
```
Contract **name** should be a descriptive noun phrase like "List Projects",
not the raw invocation. The raw invocation goes in `--input`.
Contracts MUST belong_to a **system** (the process exposing them): `--add-rel "<system-slug>:belongs_to"`
Contracts SHOULD reference the implementing component: `--add-rel "<component-slug>:references"`
Contracts SHOULD reference concepts used in input/output: `--add-rel "<concept-slug>:references"`

**Concept flags:**
```
--meaning           what it means in the domain (required)
--structure-notes   structure notes
--lifecycle         lifecycle description
--invariants        invariants (rules that must always hold)
--data-sensitivity  data sensitivity (e.g., PII, public)
```
Concepts MUST belong_to a system: `--add-rel "<system-slug>:belongs_to"`
Concepts SHOULD reference the component that owns them: `--add-rel "<component-slug>:references"`
Concepts SHOULD relate_to other concepts (ERD): `--add-rel "<concept-slug>:relates_to"`

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

**Decision flags:**
```
--category      decision category (e.g., data, api, security)
--statement     the decision itself
--rationale     why this was decided
--alternatives  alternatives considered
--tradeoffs     tradeoffs of this decision
--consequences  consequences of this decision
```

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
Delete an entity. Prompts for confirmation unless `--force`.

### syde search <query>
Full-text search across all entities.

## Query

### syde query <slug> [flags]
Rich entity lookup with relationships, learnings, and tasks.

```
--full              include body and all related data
--kind <kind>       filter by entity kind
--tag <tag>         filter by tag
--format            json|compact|refs (default: human-readable)
--impacts <slug>    transitive impact analysis (what breaks if this changes)
--flow <slug> --components   flow decomposition with component details
--related-to <slug>          all direct connections
--depends-on <slug>          entities this depends on
--depended-by <slug>         entities that depend on this
--search "text"              full-text search
--diff <slug> --since 7d     git change history
```

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
Mark plan as approved. Required before implementation.

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
Link a task to a design entity.

## Learnings

### syde remember "<text>" [flags]
```
--category    gotcha|constraint|convention|context|dependency|performance|workaround
--entity      linked entity slug
--confidence  high|medium|low (default: high)
```

### syde learn list
List all learnings.

### syde learn about <entity-slug>
Learnings for a specific entity.

### syde learn search "<text>"
Search learning text.

### syde learn stale
Learnings referencing changed entities.

### syde learn promote <slug> --to decision
Promote a learning to a formal decision.

## Constraints

### syde constraints [--json]
Show active decisions + critical learnings.

### syde constraints check <file> [--json]
Map source file to component via `component_paths` in `syde.yaml`, then show applicable constraints. Returns `{}` if file is not mapped to any component.

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
syde design create/show/preview/export/validate/link  # UIML designs
syde memory sync/list/clean    # Claude Code memory sync
syde open                      # Start dashboard + open browser
syde install-skill             # Install skill files into .claude/
```
