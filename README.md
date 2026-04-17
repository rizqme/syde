# syde

Text-first software design model that turns your AI coding agent (Claude Code, Codex, etc.) into a disciplined collaborator. The agent does the work; syde keeps it honest.

syde stores your system's architecture as markdown files in `.syde/` — requirements, components, contracts, flows, concepts, plans, and tasks. When you load the skill, the agent gets the full architecture in its first message, must clarify before writing, must produce an approved structured plan before any edit, must track the work as tasks with affected entities and files, and must pass a strict single-Finding audit before the session can end.

You don't run syde commands by hand. You talk to the agent.

## Install

```bash
make install
```

Builds the React SPA, compiles `syde` and `syded`, installs them into
`$HOME/.local/bin` (override with `PREFIX=/usr/local make install`), and
— when run from inside a project that already has a `.syde/` directory —
also installs the skill into `.claude/`, `.agents/`, `.codex/`,
`CLAUDE.md`, and `AGENTS.md`. Requires Go and [Bun](https://bun.sh).

For a fresh project, run `syde init --install-skill` once after install
to create `.syde/` and bootstrap the skill in one command.

## Quick Start (inside Claude Code or Codex)

In a session in your project, just say:

```
load skill syde
```

That activates the SessionStart hook (architecture context auto-loaded), the PreToolUse hook (blocks code edits without an approved plan + active task), and the PostToolUse hooks (overlap resolution + strict gate at session end).

Then drive the agent in plain English:

```
You: I want to add user authentication.

Agent (syde-driven):
  1. Asks clarifying questions — auth method, session duration, password
     reset flow, rate limits — with recommended answers. Waits for your
     reply.
  2. Drafts a plan with background, objective, scope, design, and a
     structured change diff (which requirements / contracts / components
     / flows are added, extended, deleted). Opens it in the dashboard.
     "Approve to proceed."
  3. You: "approve". Agent runs syde plan approve.
  4. Agent works one task at a time: starts the task, writes code, marks
     done with the real list of affected entities and files. Repeats
     until every phase is complete.
  5. Agent runs syde sync check. Every finding blocks; agent fixes them
     all before reporting done.
```

You never run `syde plan create` or `syde task done` yourself. The agent does. Your job is to confirm the clarifications, approve the plan, and review the result.

## What the skill enforces

| Gate | Where | What |
|------|-------|------|
| **Architecture context** | SessionStart hook | `syde context` injected so the agent starts every session knowing the system. |
| **Clarify before code** | Skill rules | Agent must list requirements, propose constraints, ask "what happens when this fails?" — and **wait for your confirmation**. |
| **Plan before edit** | PreToolUse hook on Write / Edit / MultiEdit / NotebookEdit | Code edits **blocked** unless a plan is approved and a task is in_progress. Hard exit code 2. |
| **Structured plan diff** | `syde plan check` | Every plan declares `delete` / `extend` / `new` lanes per kind with what / why / field_changes / draft. Plan check exits non-zero on any gap. |
| **Explicit chat approval** | Skill rules | The agent runs `syde plan approve` only after you say "approve" in chat — not on "ok" or "sure". |
| **Task-tracked work** | `syde task start` / `syde task done` | Each task names its affected entities + files; done-time flags merge into reality so drift is visible. |
| **Requirement overlap = semantic review** | `syde add requirement` gate + PostToolUse hook | TF-IDF surfaces candidate overlaps; the CLI **blocks** unless the agent resolves each by MERGE (drop the new one), RENAME (rewrite to read distinctly), or DISTINCT (`--audited slug:"why these mean different things"`, ≥20 chars of substantive reasoning). The post-plan audit re-checks every acknowledgement and fails on rubber stamps. |
| **Contract surface coverage** | Audit (planning + post-plan) | Every requirement that names a CLI command, REST path, dashboard screen, or pub-sub topic must have a matching active contract. Caught at plan time and against the corpus at rest. |
| **Flow coverage** | Audit (planning + post-plan) | Every active contract must be referenced by at least one flow step. Plans that introduce contracts must touch flows in the same diff. |
| **Strict audit** | `syde sync check` | One severity level. Every finding blocks. No `--strict` flag, no warning tier. |
| **Plan completion** | `syde plan complete` | Refuses to mark completed unless every declared change matches actual entity state — deletes really gone, news really created, extended fields really equal the declared values. |
| **Session-end gate** | Stop hook | Re-runs `syde sync check`; blocks session end on any finding. |
| **Planning ↔ post-plan symmetry** | `internal/audit/symmetry.go` parity registry + Go test | Every planning rule has an equivalent post-plan rule, so an intent missed at plan time is still caught against entity state. Adding a one-sided rule fails the test. |

## The design model

The agent works with **8 entity kinds**:

| Kind | What it represents |
|------|-------------------|
| **system** | A standalone process / app / service. Sub-systems nest via `belongs_to`. |
| **component** | An internal module of a process — `purpose`, `responsibility`, `capability`s, `boundaries`, `--file` paths. |
| **contract** | One process boundary — CLI command, REST endpoint, event, screen, storage schema, RPC, GraphQL, WebSocket. Carries invocation signature + input/output parameters. |
| **concept** | A domain glossary entry — `meaning`, `invariants`, `lifecycle`. Wired to its realisation via `implemented_by` (component), `exposed_via` (contract), `used_in` (flow). |
| **flow** | One user goal as an ordered step list. Each step references a contract by slug. |
| **requirement** | Append-only EARS shall-form intent. The five EARS patterns (Ubiquitous, Event-driven, State-driven, Unwanted-behavior, Optional-feature) are validated on save. |
| **plan** | A tracked implementation plan with structured change diff (deleted / extended / new per lane), phases, tasks. |
| **task** | A work item with affected entities + files. Done-time `--affected-*` clears drift on the listed entities. |

Each entity is a markdown file with YAML frontmatter in `.syde/<kind-plural>/<slug>.md`. Human-readable, git-friendly, editable by hand if you want to.

### Relationship types

`belongs_to`, `depends_on`, `exposes`, `consumes`, `uses`, `involves`, `references`, `relates_to`, `implements`, `applies_to`, `modifies`, `visualizes`, `refines`, `derives_from`, `implemented_by`, `exposed_via`, `used_in`.

## Dashboard

```bash
syde open
```

Starts `syded` and opens `http://localhost:5703/<project-slug>`. You get entity inboxes (2-column list + detail), plans with the structured-change diff and per-phase task progress, flow charts, file tree, force-directed entity graph, and concept glossary cards. The agent uses `syde plan open <plan>` to navigate an already-open dashboard tab via WebSocket so you can review plans without copy-pasting URLs.

## `.syde/` directory layout

```
.syde/
├── syde.yaml              # Project config
├── counters.yaml          # Per-kind ID counters (SYS-0001, COM-0002, ...)
├── tree.yaml              # Summary tree mirroring the source tree
├── index/                 # BadgerDB cache (gitignored, rebuildable via syde reindex)
├── systems/               # System entities
├── components/            # Component entities
├── contracts/             # Contract entities
├── concepts/              # Concept entities
├── flows/                 # Flow entities
├── requirements/          # Requirement entities (EARS shall-form, append-only)
├── plans/                 # Plan entities (with structured change diff + phases)
└── tasks/                 # Task entities (affected entities + files)
```

Add to `.gitignore`:
```
.syde/index/
```

Everything else in `.syde/` should be committed — it's your architecture documentation.

## CLI reference

The CLI exists for the agent and for power users / contributors. In normal use you don't invoke it directly. If you need it: every command supports `--help`, and `skill/references/commands.md` (also installed at `.claude/skills/syde/references/commands.md`) is the full reference the skill consumes.

## License

MIT
