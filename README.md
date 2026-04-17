# syde

Text-first software design model for source code repositories. Works as a standalone CLI and as a Claude Code / Codex skill that enforces architectural constraints during development.

syde stores your system's architecture as markdown files in `.syde/` — requirements, components, contracts, flows, concepts, plans, and tasks. A BadgerDB index enables fast queries; a summary tree mirrors your source tree with one-line summaries on every file and folder. When integrated with Claude Code or Codex, syde auto-loads your architecture at session start, blocks code edits without an approved plan and active task, and runs a strict audit that exits non-zero on any finding.

## Install

```bash
make install
```

Builds the React SPA, compiles `syde` and `syded`, and installs them into
`$HOME/.local/bin` (override with `PREFIX=/usr/local make install`).
Requires Go and [Bun](https://bun.sh).

## Quick Start

```bash
cd your-project
syde init --install-skill
```

This writes:
- `.syde/` — design model directory with entity subdirectories, BadgerDB index, summary tree
- `.claude/skills/syde/SKILL.md` — skill that controls Claude's behavior
- `.claude/skills/syde/references/*.md` — entity spec, command reference, requirement-derivation, sync-workflow
- `.claude/hooks/syde-hooks.json` — SessionStart, PreToolUse, PostToolUse, and Stop hooks
- `.codex/hooks.json` + `.codex/config.toml` — equivalent for Codex
- `CLAUDE.md` and `AGENTS.md` — mandatory rules appended idempotently

## What happens in a Claude Code / Codex session

**1. Session start — architecture auto-loads.** SessionStart hook runs `syde context` and injects your full design model into Claude. Claude starts every session already knowing the system.

**2. Phase 0 — refresh the summary tree.** `syde tree scan` walks the source tree, marks changed files stale. Claude (or subagents in parallel) writes a 1-3 sentence summary on every stale file/folder via `syde tree summarize`. Claude consults `syde tree show <path>` and `syde tree context <path>` instead of cold `Read` calls — cheaper on tokens, framed by the architectural breadcrumb.

**3. Clarify first.** The skill instructs Claude to be critical: list missing requirements with recommended answers, propose constraints, ask "what happens when this fails?". Claude waits for explicit confirmation.

**4. Plan — required before any code edit.** The PreToolUse hook **blocks** every Write / Edit / MultiEdit / NotebookEdit unless: (a) a plan is in `approved` or `in-progress` state, AND (b) a task is `in_progress`. Claude must:

```bash
syde plan create "Add User Auth"               \
  --background "..." --objective "..." --scope "..." --design "..."

# Declare the structured change diff (per-kind lanes: requirements,
# systems, concepts, components, contracts, flows). Each entry has
# what / why and either field_changes (for extends) or a draft map
# (for new entries). Tasks are linked to the changes they implement.
syde plan add-change new    <plan> requirement --name "..." --draft statement="The system shall ..." --draft req_type=functional --draft priority=must
syde plan add-change extend <plan> component   auth-middleware --field responsibility="..." --task harden-jwt-keys
syde plan add-change delete <plan> contract    legacy-login    --why "Replaced by JWT"

syde plan add-phase <plan> --name "Scaffolding" --details "..."
syde task create   "Harden JWT keys" --plan <plan> --phase phase_1 \
    --objective "Rotate signing key without downtime" \
    --affected-entity auth-middleware --affected-file internal/auth/keys.go

syde plan check  <plan>     # Must exit 0 — every Finding blocks
syde plan open   <plan>     # Open in dashboard for review
syde plan approve <plan>    # ONLY after explicit user approval in chat
```

**5. Implement with task tracking.**

```bash
syde task start <task>      # Marks active; PreToolUse hook now lets Write through
# ... write code ...
syde task done <task> \
    --affected-entity auth-middleware --affected-entity jwt-config \
    --affected-file internal/auth/middleware.go \
    --affected-file internal/auth/keys.go
```

The done-time `--affected-*` flags merge into the predicted set and bump `updated_at` on each affected entity (drift tracking).

**6. Finish — strict gate.**

```bash
syde tree scan                    # Refresh summaries for changed files
syde sync check                   # Single severity level: every Finding blocks
syde plan complete <plan>         # Refuses unless every declared change matches reality
```

The Stop hook re-runs `syde sync check` and blocks session end if anything fails.

## The design model

syde models your architecture with **8 entity kinds**:

| Kind | What it represents | Required fields |
|------|-------------------|-----------------|
| **system** | A standalone process / app / service | `description`; sub-systems via `belongs_to:<parent-system>` |
| **component** | An internal module of a process | `purpose`, `responsibility`, `capability` (≥1), `boundaries`, `--file` paths |
| **contract** | One process boundary (CLI cmd / REST endpoint / event / screen / storage schema) | `contract_kind`, `interaction_pattern`, `input`, `input_parameter` (≥1), `output`, `output_parameter` (≥1); screens add `--wireframe` UIML |
| **concept** | A domain glossary entry | `meaning`; `invariants` and `lifecycle` recommended |
| **flow** | One user goal as a step list | `trigger`, `goal`, `--step` entries (`id\|action\|contract\|description\|on_success\|on_failure`) |
| **requirement** | Append-only EARS shall-form intent | `statement` (EARS-validated on save), `req_type`, `priority`, `verification`, `source`, `rationale` |
| **plan** | A tracked implementation plan with structured change diff | `background`, `objective`, `scope`, `design`, change lanes per kind, phases with tasks |
| **task** | A work item with affected entities + files | `objective`, `details`, `acceptance`, `--affected-entity`, `--affected-file` |

Each entity is a markdown file with YAML frontmatter in `.syde/<kind-plural>/<slug>.md`. Human-readable, git-friendly, editable by hand or via CLI.

### Relationship types

`belongs_to`, `depends_on`, `exposes`, `consumes`, `uses`, `involves`, `references`, `relates_to`, `implements`, `applies_to`, `modifies`, `visualizes`, `refines`, `derives_from`, `implemented_by`, `exposed_via`, `used_in`.

The last three (`implemented_by`, `exposed_via`, `used_in`) wire concepts to the components, contracts, and flows that realise them — the dashboard renders them as grouped chips on the concept detail panel.

## Strict audit — single Finding severity

There is **one severity level**. Every audit finding blocks `syde sync check`, `syde plan check`, `syde plan complete`, and the session-end Stop hook. There is no `--strict` flag and no non-blocking warning tier. The skill teaches: "if it's worth mentioning, it's worth fixing."

`syde sync check` bundles:

- **Structural audit** — required fields, recommended fields, broken relationship targets, cycles in `belongs_to` and `depends_on`.
- **Tree ↔ entity consistency** — every `--file` ref exists in the summary tree.
- **Orphan detection** — every non-ignored source file is owned by at least one entity (`syde files orphans` lists them).
- **Drift** — file mtime newer than the owner entity's `updated_at` becomes a Finding (cleared by `syde task done` with `--affected-*`).
- **Summary-tree staleness** — any stale file or folder in `tree.yaml`.
- **Plan completion drift** — declared changes that no longer match entity state.
- **Requirement overlap distinction** — acknowledged TF-IDF overlaps without semantic distinction text fail as rubber stamps.
- **Contract surface coverage** — active requirements naming a CLI / REST / screen / event surface must have a matching active contract.
- **Flow coverage** — every active contract is referenced by at least one flow step.

## Requirement overlap workflow (MERGE / RENAME / DISTINCT)

`syde add requirement` computes TF-IDF similarity against every active requirement. Above the 0.6 threshold the CLI **blocks** the create unless every surfaced overlap is resolved by one of three outcomes:

| Outcome | When | How |
|---------|------|-----|
| **MERGE** | The two requirements mean the same thing. | Abandon the new one; reuse the existing slug. |
| **RENAME** | They are semantically distinct but the statements accidentally share vocabulary. | Rewrite the new statement and retry. |
| **DISTINCT** | They are genuinely close cousins that must both exist. | Retry with `--audited <slug>:"why these two mean different things"` for each overlap. The distinction text is persisted on the requirement and must be ≥20 characters of substantive reasoning. |

The post-plan audit re-checks every acknowledgement: empty / trivially short distinction text fails as a rubber stamp. A PostToolUse hook on `syde add requirement` injects a system reminder listing the resolution paths so Claude cannot silently skip the review.

`--force` overrides the gate but should be rare and explained.

## Plan structured-change diff

Every plan declares its work as a structured diff with three lists per kind lane (`deleted`, `extended`, `new`):

```bash
syde plan add-change delete <plan> component legacy-auth --why "Replaced by JWT in phase 2"

syde plan add-change extend <plan> component api-server \
  --what "Add /api/plans routes and wire the plan completion handler" \
  --why  "REQ-0331 requires a dashboard-visible plan inbox" \
  --field responsibility="Request routing, validation, plan lifecycle endpoints" \
  --field boundaries="No direct DB access; delegates to storage layer" \
  --task wire-plan-routes

syde plan add-change new $PLAN contract \
  --name "Plan Completion API" \
  --what "POST /api/plans/:slug/complete returning the validator findings" \
  --why  "Dashboard reviewers need a one-click gate" \
  --draft contract_kind=rest --draft interaction_pattern=request-response \
  --draft input="POST /api/plans/:slug/complete" \
  --draft 'input_parameters=[{"path":"slug","type":"string","description":"plan slug"}]' \
  --draft output="200 OK application/json"
```

`syde plan complete` runs the completion validator: every declared `delete` must no longer exist, every declared `new` must exist with the declared name + kind, every declared `extend` field_change must equal the declared value (or be empty if the sentinel `DELETE` was used). If anything fails, the plan stays `approved` / `in-progress`.

## Planning ↔ post-plan symmetry

Every planning-time rule has an equivalent post-plan rule encoded in `internal/audit/symmetry.go`. So intent flagged at `syde plan check` is also caught against actual entity state at `syde sync check`. A Go test asserts the parity registry stays populated as new rules are added.

Examples of paired rules:
- requirement overlap detection (planning + post-plan)
- contract surface coverage on requirements (planning + post-plan)
- contract-flow coverage (planning + post-plan)
- requirement traceability (planning lane coverage + post-plan link check)

## CLI reference

### Architecture overview

```bash
syde context                         # Full snapshot: entities, plans, tasks
syde context --json                  # Machine-readable (used by SessionStart hook)
syde status                          # Entity counts
syde sync check                      # Strict audit gate (every Finding blocks)
```

### Entity CRUD

```bash
syde add component "Auth Service" \
  --description "Handles authentication" \
  --purpose "Validate session tokens before request handlers run" \
  --responsibility "JWT verification, OAuth2 token exchange" \
  --capability "Verify JWT" --capability "Refresh access token" \
  --boundaries "Does NOT own user profiles" \
  --file internal/auth/middleware.go --file internal/auth/keys.go \
  --add-rel "myapp:belongs_to" \
  --add-rel "auth-token-validation-required:references"

syde get auth-service                # Full entity details
syde list components                 # All components
syde search "auth"                   # Full-text search
syde update auth-service --capability "Verify JWT" --capability "Refresh access token" --capability "Revoke session"
syde remove auth-service             # With confirmation; --force to skip
```

Every entity supports `--note` (repeatable) for informal reminders, quirks, operational notes.

### Rich queries

```bash
syde query <slug>                          # Entity + relationships + tasks
syde query <slug> --full                   # Everything including body
syde query <slug> --content                # Read source file with framing
syde query --kind component --tag security # Filter
syde query --search "BadgerDB index"       # Multi-word AND, auto-broadens to OR
syde query --code "func NewStore("         # Literal-string search across tracked source
syde query --file internal/storage/index.go --content  # Read file framed by owner
syde query --impacts auth-service          # Transitive impact analysis
syde query --depends-on query-engine       # Forward graph walk
syde query --depended-by query-engine      # Reverse graph walk
syde query --diff <slug> --since 7d        # Git change history
```

`syde query` is the single entry point — it surfaces orphan files (`⚠ DRIFT`) and uncovered code in the same call.

### Plans + tasks

```bash
syde plan create   "Add Payment Processing" --background ... --objective ... --scope ... --design ...
syde plan add-change new    <plan> requirement --name ... --draft ...
syde plan add-change extend <plan> component   <slug>  --field key=value --task <task>
syde plan add-change delete <plan> flow        <slug>
syde plan add-phase <plan> --name "Scaffolding" --details ...
syde plan show     <plan> --full
syde plan show-changes <plan>             # Structured diff grouped by kind
syde plan check    <plan>                 # Pre-approval audit; must exit 0
syde plan open     <plan>                 # Open in dashboard
syde plan approve  <plan>                 # After explicit user chat approval
syde plan complete <plan>                 # Strict completion gate
syde plan estimate <plan>                 # Size + split recommendation

syde task create "Build payment webhook" --plan add-payment --phase phase_1 \
    --objective ... --details ... --acceptance ... \
    --affected-entity payment-service --affected-file internal/payment/webhook.go
syde task start <task>
syde task done  <task> --affected-entity ... --affected-file ...
syde task block <task> --reason "Waiting for Stripe API key"
syde task sub   <parent> "Subtask name"
syde task list
```

### Files & summary tree

```bash
syde files orphans                   # Tracked files with no owning component
syde files coverage <path>           # Who owns this file?

syde tree scan                       # Walk fs, mark changed files stale (.gitignore honored)
syde tree status                     # Stale file/folder counts
syde tree changes --leaves-only      # List stale leaves for the summarize loop
syde tree show <folder>              # Children with their stored summaries
syde tree context <path>             # Breadcrumb + summary + content
syde tree summarize <path> --summary "..."
syde tree ignore <path>              # Exempt from orphan + stale checks
```

### Dashboard

```bash
syde server start                    # Start syded on :5703
syde open                            # Start + register project + open browser
syde plan open <plan>                # WebSocket-navigates an open dashboard tab, or opens new tab
syde wireframe render <screen-slug> --format html|ascii|image --out /tmp/x.html --open
```

The dashboard at `http://localhost:5703/<project-slug>` shows entity inboxes (2-column list + detail), plans with structured-change diffs and per-phase task progress, flow charts, file tree, force-directed graph, and concept glossary cards.

## Hooks

The installed `.claude/hooks/syde-hooks.json` (and `.codex/hooks.json`) wires:

- **SessionStart** — `syde context` injected into Claude.
- **PreToolUse: Write/Edit/MultiEdit/NotebookEdit** — blocks the edit unless an approved plan is active AND a task is `in_progress`. Excluded paths: `.syde/`, `.claude/`, `node_modules/`, `vendor/`, `web/dist/`, `.git/`, `/tmp/`.
- **PostToolUse: Write** — warns if the new file is not mapped to any component.
- **PostToolUse: Bash on `syde add requirement`** — when overlap candidates appear, injects a MERGE/RENAME/DISTINCT system reminder listing the surfaced slugs.
- **PostToolUse: Bash on `git commit/push`** — warns if tasks are pending or `syde sync check` fails.
- **Stop** — re-runs `syde sync check`; blocks session end on any Finding.

## `.syde/` directory layout

```
.syde/
├── syde.yaml              # Project config
├── counters.yaml          # Per-kind ID counters (SYS-0001, COM-0002, ...)
├── tree.yaml              # Summary tree mirroring the source tree
├── index/                 # BadgerDB cache (gitignored, rebuildable via syde reindex)
├── systems/               # System entities
├── components/            # Component entities
├── contracts/             # Contract entities (rest / cli / event / screen / pubsub / storage / rpc / graphql / websocket)
├── concepts/              # Concept entities (glossary)
├── flows/                 # Flow entities (steps reference contracts by slug)
├── requirements/          # Requirement entities (EARS shall-form, append-only)
├── plans/                 # Plan entities (background/objective/scope/design + structured change diff + phases)
└── tasks/                 # Task entities (affected entities + files)
```

Add to `.gitignore`:
```
.syde/index/
```

Everything else in `.syde/` should be committed — it's your architecture documentation.

## License

MIT
