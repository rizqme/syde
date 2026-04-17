---
name: syde
description: >
  Always active when installed. Manages the project's software design model
  in .syde/. Enforces architectural constraints, tracks plans and tasks.
  Triggers on any code modification, design discussion, plan creation, or
  architecture query.
tools: Read, Write, Edit, Bash, Glob, Grep
---

# syde — Software Design Model

This project has a syde design model in `.syde/`. Architecture context is
**auto-loaded at session start** via the SessionStart hook — you already have
the full entity map and requirements in your context. Do NOT re-run
`syde context` manually.

## Getting Context — syde first, always

**syde is the single entry point for understanding any file, symbol, or
entity in this project.** It is not "an architecture tool you use alongside
grep" — it is the unified surface where design entities, source files, and
free-text content meet. Every query you run also tells you whether
architecture and code are in sync. Bypassing syde to grep or read raw is
not just inefficient — it silently disconnects the two and lets the
design model rot.

### The three-question checklist (run before any Grep / Read)

When you need context on something — a symbol, a file, a concept, a
behaviour — ask these three questions in order:

1. **What entity owns this?** → `syde query --file <path>` or
   `syde query <slug>`
2. **What does syde know about this term?** → `syde query --search "<term>"`
   (then `--any` if AND yields nothing — the engine actually retries
   automatically and labels the results "broadened")
3. **What code references it?** → `syde query --code "<symbol>"` —
   literal-string search across every tracked source file, with the
   owning entity attached to every hit

Three syde calls beat one wrong grep. The owning-entity annotation on
every result is the architecture↔code bridge: you don't have to ask "is
this file mapped?" separately, the answer arrives with the search.

### Architecture↔code sync is a feedback loop, not a side effect

Two specific signals show up in syde output and must be acted on
immediately:

- **`syde query --file <path>` reports `⚠ DRIFT: no owning entities`** —
  this tracked file is an orphan. Either add it to a component's `--file`
  list, or `syde tree ignore <path>` if it is genuinely not part of the
  design model. Don't move on without resolving it.
- **`syde query --code <pattern>` reports a hit on a file with `⚠ orphan`** —
  same deal: code lives in a file the design model doesn't claim. Map it
  or ignore it before continuing.

These warnings are the model telling you it has rotted relative to the
code. Acting on them is part of the workflow, not a separate cleanup step.

### Tools to use, tools to avoid

**Use `syde query` for everything tracked by the summary tree.** Including:
finding a Go symbol, reading a source file with framing, listing all
entities of a kind, finding what a file participates in, finding what
depends on what, searching documentation, and triaging orphans / drift.

**Reserve `Grep` and `Read` for files the summary tree intentionally
ignores** — vendor/, node_modules/, generated assets, .git/, build
artifacts, binary blobs. If a file is tracked by `syde tree scan`, you
should not be reading it with `Read` — use `syde query --file <path>
--content` instead, which gives you the same content plus the
architectural framing in one call.

If you find yourself reaching for `Grep` to find a symbol in code, stop
and use `syde query --code` instead. If you find yourself reaching for
`Read` to open a tracked file, use `syde query --file <path> --content`.

### Recipe cookbook

```bash
# 1. SYMBOL LOOKUP — find every reference to a Go identifier in code,
#    each hit annotated with its owning entity.
syde query --code ConceptEntity
syde query --code "func NewStore("

# 2. FILE READ WITH FRAMING — read a source file and learn which
#    entities own it + their one-hop neighbours, all in one call.
syde query --file internal/storage/index.go --content

# 3. FILE PARTICIPATION ONLY — same as #2 without the body, when
#    you only want the architectural framing.
syde query --file internal/storage/index.go

# 4. DIRECTORY SWEEP — owners of every file under a directory.
syde query --file internal/cli/

# 5. KEYWORD SEARCH IN ENTITY PROSE — name, description, purpose,
#    notes, body. Multi-word AND by default; auto-broadens to OR
#    when AND yields zero, results labelled "broadened".
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

The two new flags introduced by this version are `--code <pattern>` and
`--content` (paired with `--file`). They cover everything that used to
require `Grep` or `Read` on a tracked file. Use them.

## Phase 0: Survey the Codebase (always run first)

Before planning or writing any code, refresh the **summary tree** so you
understand the project without re-reading every file. The tree is a
mirror of the source directory with a short human-written summary on
every file and folder. syde tracks file hashes so changed files bubble
up as "stale" automatically.

```
syde tree scan              # walk fs, mark changed files stale (respects .gitignore)
syde tree status            # see how many nodes are stale
```

If any nodes are stale, iterate leaves-first until the tree is clean.
**Delegate the summarizing work to subagents** — do NOT read stale files
into the main session context. Reading every file in the project just to
write summaries burns the main session's token budget for nothing.

**The right pattern**:

```
# 1. Main session: list stale leaves
syde tree changes --leaves-only --format json

# 2. Main session: dispatch subagents in parallel (via the Task/Agent tool).
#    Batch 5-10 stale leaves per subagent call. Each subagent's prompt:
#
#      "Summarize these syde tree nodes, one at a time. For each path:
#       1. Run `syde tree context <path>` to read the file content + parent breadcrumb
#       2. Write a 1-3 sentence summary describing what the file does
#       3. Run `syde tree summarize <path> --summary '...'` to store it
#       Return only a short confirmation when done — no file contents, no commentary."
#
#    The subagent reads files, the main session never sees them.

# 3. Main session: once subagents return, re-list
syde tree changes --leaves-only

# 4. When all file leaves are clean, summarize the folder layer.
#    Folder summaries DON'T need subagents — they derive from children's
#    summaries, which the main session can read cheaply via:
syde tree show <folder>
syde tree summarize <folder> --summary "..."

# 5. Repeat until `syde tree changes` prints "clean"
```

A file summary answers "what does this file do?" in 1-3 sentences. A
folder summary answers "what is this folder about, and how do its
children relate?" derived from the children's summaries — you do NOT
need to re-read every file inside a folder; `syde tree show <folder>`
is enough for the main session to write a folder summary directly.

**Why subagents for files, not folders?** File summaries require the
raw file content (large, noisy, one-off use). Folder summaries require
only the children's already-stored summaries (small, curated, main
session already has them from `tree show`). Delegating files to
subagents keeps main context lean; writing folder summaries in the main
session avoids unnecessary subagent overhead.

Once the tree is clean, consult `syde tree show` and `syde tree context
<path>` **instead of** blind `Read` calls when you need to understand a
file. That's cheaper on context and gives you the architectural framing
(the ancestor breadcrumb) alongside the content.

## First Session: Sync with Existing Code

If this is a newly initialized syde project on an existing codebase (i.e., there
is code but few or no syde entities), **after completing Phase 0 above**, run sync:

```
syde sync
```

Then follow `references/sync-workflow.md` to bootstrap or verify the design model.
During bootstrap, creating entities (components / contracts / concepts) MUST
source their understanding from `syde tree context <path>`, NOT naive
`Read` calls. The tree context returns breadcrumb + file summary + content in
one shot — exactly the framing needed to fill in `--purpose`,
`--responsibility`, `--boundaries`, and the like. Rationale: (a) the
summaries encode the architectural framing Claude just wrote in Phase 0,
(b) the breadcrumb tells you which sub-system / folder the file belongs
to so `belongs_to` relationships land on the right target, (c) it's
cheaper on tokens than re-reading files cold.

**Skip the sync step if** the model is already populated and you're continuing normal work.
Phase 0 is NEVER skipped — always run `syde tree scan` at session start.

### Bulk entity creation: use a shell script, not one tool call per entity

Bootstrapping a non-trivial codebase creates dozens to hundreds of entities
(one system + sub-systems, a component per package, a contract per CLI
command / HTTP endpoint / schema prefix, concepts per domain type, flows,
requirements). Doing this as individual tool calls burns tool-use budget
for nothing.

**Instead, write a shell script to `/tmp/syde-<batch>.sh`** containing
chained `syde add` / `syde update` commands, then execute it with one
`bash` call. Group related entities per script (systems, components,
contracts, concepts, flows, requirements) and run them in dependency order
so `belongs_to` targets already exist when children reference them.

```bash
cat > /tmp/syde-components.sh <<'EOF'
#!/bin/bash
set -e
cd /path/to/project
syde add component "Storage Engine" --purpose ... --add-rel "my-app:belongs_to"
syde add component "Query Engine"   --purpose ... --add-rel "my-app:belongs_to" \
  --add-rel "storage-engine:depends_on"
# ... many more
EOF
bash /tmp/syde-components.sh
```

This pattern makes bulk bootstrap 10-50x cheaper in tool calls. Individual
`syde add` calls are still fine when you're creating one entity during
task implementation — batch scripts are specifically for bootstrap.

### Concurrency caveat for parallel subagents writing the tree

`syde tree summarize` is not safe for truly concurrent writers — several
parallel subagents writing to `.syde/tree.yaml` can collide on the atomic
rename and a single update may be lost. Mitigations:

- Keep subagent batches modest (8-10 files each) so writes are spread
  out and collisions are rare.
- After the subagent wave returns, run `syde tree changes --leaves-only`
  — any straggler that was lost to a race will resurface here. Re-summarize
  it in the main session (or a single retry subagent).
- Folder summaries are always written serially in the main session, so
  they never hit this race.

Be aware of this on first encounter so future sessions stay alert — the
tool does not emit an error, the update just silently fails to stick.

### Orphan files: three valid resolutions

`syde validate` raises an ERROR for any non-ignored source file that has
zero owning entities. During bootstrap this is expected — every new
component needs its files attached. Three ways to clear:

1. **Add to an entity's file list** (the usual answer):
   `syde update <component> --file path/one.go --file path/two.go`
   `syde update` replaces the files list, so pass every file the
   component owns, not just the new one. Re-enumerate with
   `syde tree show <folder>` if the component spans a whole directory.

2. **Ignore the file** if it's not part of the design model (README,
   Makefile, go.mod, generated assets, public/favicon.svg, editor
   configs): `syde tree ignore <path>`. Ignored files stay tracked on
   disk but are exempt from the orphan validator and from the
   `--leaves-only` stale list.

3. **Create a new component** if the orphan belongs to an undocumented
   module. Prefer this over stuffing unrelated files into an existing
   component.

## Mandatory Workflow

Follow these phases in order. Do NOT skip or reorder.

### Phase 1: CLARIFY (before anything else)

**STOP. Do NOT run `syde plan create` or `syde add` until this phase is complete
and the user has confirmed.**

You must be **critical and thorough** when gathering requirements. Do not accept
vague requests — probe for missing information, challenge assumptions, and
propose constraints the user hasn't considered.

1. Identify the project type (web, backend, CLI, mobile, library, full-stack)
2. Use the checklist in `references/clarify-guide.md` for that project type
3. For each question, provide your **recommended answer** with reasoning
4. Ask "what happens when this fails?" for every feature
5. Flag any requirements the user hasn't mentioned but should decide on
6. Present **ALL questions through the runtime's ask-user-question tool**
   and WAIT for the user's answer. In Codex Plan mode this is
   `request_user_input`; in other runtimes use the equivalent
   AskUserQuestion / user-input tool when it exists.
7. **Do NOT proceed without explicit user approval**

If the runtime does not expose an ask-user-question tool, stop and ask the
questions in a plain assistant message. Do not create a plan, task, or entity
until the user answers. Clarification is not optional; only the transport
changes.

**Common mistake**: jumping straight to `syde plan create` because the request
seems clear. Even "build a dashboard" has 20+ hidden choices (framework,
routing, state management, responsiveness, accessibility, etc.). Always clarify.

Example:
```
Before I start, I need to clarify several things:

1. Auth strategy — you said "add login". Which method?
   Recommended: JWT with refresh tokens (stateless, fits your existing API)
   You haven't mentioned: password reset flow, session expiry, rate limiting on login

2. Error handling — no error format specified.
   Recommended: RFC 7807 Problem Details JSON
   This affects every endpoint, so we should decide now.

3. Database migrations — you have no migration tool yet.
   Recommended: golang-migrate (matches your Go stack)
   Risk: if we skip this, schema changes will be manual and error-prone.

Please confirm or adjust each point.
```

## Entity IDs and Slugs

**IDs** are counter-based: `SYS-0001`, `COM-0002`, `CON-0003`, etc.
One counter per kind, stored in `.syde/counters.yaml`. Never reused even
when entities are deleted — the counter only goes up. IDs are stable
and globally unique within a project.

**Slugs** have a random 4-char alphanumeric suffix appended to the
name-slugified base: `cli-a3f2`, `user-login-b7k9`, `syde-cli-hitk`.
The suffix makes every slug unique by construction so two entities with
the same name can coexist. The slug is the filename under `.syde/<kind>/`.

**Addressing** — three equivalent ways to identify an entity in any
command (`syde get`, `syde update`, relationship targets, `--affected-entity`):

| Form | Example | Notes |
|---|---|---|
| Full slug | `cli-a3f2` | Always unique. Preferred for scripts. |
| Bare name slug | `cli` | Works when only one entity matches. Errors "ambiguous" if multiple. |
| Parent/child path | `syde-cli/cli` | Disambiguates by walking `belongs_to`. Each segment can be full or bare. |

Relationship targets stored in `.yaml` files should prefer the **bare
name slug** for readability, unless ambiguous — in which case use the
full slug or a parent/child path.

## Required For Every Entity

Every entity of every kind MUST have a `--description` — a one-sentence
identification of what the entity is. The audit emits a blocking Finding
on entities with no description. Without it the dashboard list
views render empty cards, `syde get` shows a stub, and the model becomes
unreviewable.

Description is a short noun phrase or one-liner ("Cobra command tree
exposing every syde feature as a CLI subcommand", "BadgerDB key for the
per-entity FileRef index", "Mirror of the source tree with cascading
stale tracking"). Keep it distinct from `purpose` (the *why*) and
`responsibility` (the *what does*) — description is the elevator pitch
that explains what the entity *is* in one breath.

## Entity Hierarchy

The design model has a strict hierarchy. Always follow it.

```
system                                   ← a standalone process / app / service
├── sub-systems (system belongs_to system)  ← another standalone process inside
├── components (belongs_to system)        ← internal modules of this process
│   └── depends_on other components (NO CYCLES)
├── contracts (belongs_to system)         ← the process's external boundaries
├── concepts (belongs_to system)          ← domain glossary terms
│   ├── implemented_by components / exposed_via contracts / used_in flows
│   └── relates_to other concepts
└── flows, requirements, plans, tasks
```

**What is a system?** A system is anything with its own process / binary /
service / standalone app. Examples: a CLI binary, an HTTP API daemon, a
background worker, a mobile app, a desktop app, a web frontend served to a
browser. If you can `./run-it` independently, it's a system.

**Sub-systems.** A project may ship multiple standalone processes. Each is
its own `system`, and they nest via `belongs_to` pointing at the parent
system. For example, a project might have a top-level `my-app` system with
two sub-systems `my-app API` (HTTP daemon) and `my-app Worker` (background
process). Each sub-system has its own components and contracts. System
nesting MUST be acyclic.

**Entity `files` are concrete paths only — NO wildcards.** Every `--file`
entry on a component / contract / etc. must be a literal path that
exists as a node in the summary tree. Globs like `internal/cli/*.go`
are rejected by the validator — enumerate each file. This is what lets
the dashboard and `syde query` render the tree summary next to every
file ref, and what guarantees `syde constraints check <file>` has a
deterministic answer. If a component spans many files, list them all
(the tree already knows them — use `syde tree show <folder>` to get
the list cheaply, then pass each as a separate `--file` flag).

**Contracts live on systems, not components.** A contract is a *process
boundary* — an HTTP endpoint, a CLI command, a published event, a WebSocket
message. It is exposed BY the process (system), even if a specific component
inside the process implements it. Put `belongs_to` on the contract pointing
at the system that exposes it.

**Every entity supports `--note` (repeatable).** Use `--note` to attach
informal reminders, quirks, or operational notes that don't fit into
structured fields. Notes render as a bullet list in the dashboard and in
`syde get`. Example: `syde update cli --note "fsnotify fires twice on save"
--note "Windows line endings break the YAML parser"`.

### Component rules

Every component MUST have:
- **`--purpose`** — WHY it exists (the problem it solves)
- **`--responsibility`** — WHAT it does (one-liner)
- **`--capability`** — concrete capabilities it provides (repeatable, e.g., "Store entities", "Index entities", "Query by tag")
- **`--boundaries`** — what it does NOT do
- **`--file`** — source file paths
- Relationship: `belongs_to:<system-slug>`
- Zero or more: `depends_on:<component-slug>` (must be acyclic)

Example:
```
syde add component "Storage Engine" \
  --purpose "Persist entities as markdown with fast lookups" \
  --responsibility "CRUD for entities with BadgerDB index" \
  --capability "Serialize entities to YAML+markdown" \
  --capability "Index entities in BadgerDB" \
  --capability "Query by ID, slug, tag, relationship" \
  --capability "Full-text search across entity fields" \
  --boundaries "Does NOT define entity models. Does NOT resolve cross-entity graphs." \
  --file "internal/storage/*.go" \
  --add-rel "syde:belongs_to"
```

### Contract rules — ONE contract per endpoint/command/event

**Every API endpoint, CLI command, WebSocket event, RPC call, or pub/sub event
is its own contract.** Do NOT create a single "API" contract covering many
endpoints. Fine-grained contracts are the rule.

Every contract MUST have:
- **Descriptive name** — a noun phrase describing what the boundary is ("User Login", "List Projects", "User Registered"). NOT the raw invocation.
- **`--contract-kind`** — one of:
  - `rest` — HTTP REST endpoint (use with `--interaction-pattern request-response` or `async`)
  - `cli` — command-line invocation (use with `request-response`)
  - `event` — pub/sub event emitted on a topic (use with `pub-sub`)
  - `rpc` — remote procedure call (use with `request-response` or `streaming`)
  - `graphql` — GraphQL operation (use with `request-response`)
  - `websocket` — WebSocket message (use with `streaming` or `async`)
  - `pubsub` — pub/sub topic (use with `pub-sub`)
  - `storage` — data schema (KV key prefix, SQL table, protobuf message, cache key, queue payload). Use with `--interaction-pattern schema`. Input is the key/table/type template, input parameters are the fields/columns, output parameters are indexes / constraints / foreign keys.
  - `screen` — UI screen / page rendered to a user (a React route, a desktop window, a TUI view). Use with `--interaction-pattern render` and **`--wireframe`** carrying a UIML source describing the layout. Input is the route path / launch trigger; input parameters are route/query params; output is "rendered UI"; output parameters are user actions exposed by the screen. The dashboard renders the wireframe as a preview in the contract detail panel.
- **`--interaction-pattern`** — one of `sync`, `async`, `request-response`, `pub-sub`, `streaming`, `schema`, `render`. Use `schema` only with `--contract-kind storage` and `render` only with `--contract-kind screen`.

**Screen contracts** also require:
- **`--wireframe`** — UIML source. Validator runs `uiml.Parse` on it and rejects malformed wireframes as ERRORs. The dashboard renders wireframes server-side via the dark-mode wireframe renderer (charcoal-on-dark, region badges, ✕-rect placeholders, line-bar text). You can also render wireframes from the terminal with `syde wireframe render <slug>` which supports `--format html|ascii|image` and `--out <path>` / `--open`.

**UIML wireframe vocabulary** (attribute-rich — the lexer handles attributes correctly):

| Tag | Effect | Common attributes |
|---|---|---|
| `<screen>` | Outer container, optional `name` chip | `name`, `direction="horizontal\|vertical"` |
| `<layout>` | Flex container | `direction="horizontal\|vertical"` |
| `<grid>` | CSS grid | `cols="N"` |
| `<stack>` | Vertical stack | (none) |
| `<sidebar>` | Side rail with `KINDS`-style label | `name`, `width="200"` |
| `<main>` | Main content region | `name` |
| `<panel>` | Sub-region (use beside `<sidebar>` and `<main>` for inboxes) | `name`, `width="360"` |
| `<card>` | Bordered card | `name` |
| `<section>` | Sub-section with optional title | `title="..."` |
| `<navbar>` | Horizontal top bar | (none) |
| `<heading>` | Bold heading text | (text content) |
| `<text>` / `<paragraph>` | If empty → 3 line-bar placeholders; if filled → grey text | (text content) |
| `<button>` | Bordered rounded label | (text content) |
| `<button-group>` | Horizontal row of buttons | (none) |
| `<list>` / `<item>` | Vertical list with optional `<image/>` thumbnail + labels | `<item active="true">` for selected row |
| `<menu>` | Vertical nav stack | (none) |
| `<image/>` / `<placeholder/>` | ✕-rect media stub | (self-closing) |
| `<input>` | Label + thin underline | `label`, `placeholder` |
| `<search>` | Bordered search box | `placeholder` |
| `<metric>` | Stat label + value | `label`, `value` |
| `<icon>` | Small bordered glyph | `glyph="✓"` |
| `<divider/>` | Thick horizontal bar | (self-closing) |
| `<tabs>` / `<tab>` | Horizontal tab strip | `<tab active="true">` |

Use `<layout direction="horizontal">` to lay regions side-by-side.
The classic inbox pattern is `<screen direction="horizontal"><sidebar/><panel width="360"/><main/></screen>`.

Worked example for a screen contract:

```bash
syde add contract "Components Inbox" \
  --description "2-column inbox with sidebar + list + detail" \
  --contract-kind screen --interaction-pattern render \
  --input "/component" \
  --input-parameter "filter|string|optional filter DSL query" \
  --output "rendered components list and detail" \
  --output-parameter "click|event|select component in detail panel" \
  --wireframe '<screen name="Components Inbox" direction="horizontal"><sidebar name="Kinds" width="200"><menu><item>Systems</item><item active="true">Components</item><item>Contracts</item></menu></sidebar><panel name="List" width="360"><heading>Components</heading><button-group><button>Filter</button><button>Sort</button></button-group><list><item active="true"><image/><label>CLI Commands</label><label>42 files</label></item><item><image/><label>Storage Engine</label><label>8 files</label></item></list></panel><main name="Detail"><heading>CLI Commands</heading><text/><section title="Files"><text/></section><button-group><button>Edit</button><button>Delete</button></button-group></main></screen>' \
  --add-rel "syded-dashboard:belongs_to" \
  --add-rel "web-spa:references"
```

Once created, render the wireframe from the terminal:

```bash
# HTML to stdout
syde wireframe render components-inbox

# Open in browser
syde wireframe render components-inbox --out /tmp/x.html --open

# PNG screenshot via headless Chrome
syde wireframe render components-inbox --format image --out /tmp/x.png

# Plain ASCII to stdout
syde wireframe render components-inbox --format ascii
```
- **`--input`** — the invocation signature (the raw command/path, e.g. `GET /api/projects`, `syde plan create <name>`, `users.created` topic)
- **`--input-parameter`** — structured parameter entry: `"path|type|description"` (repeatable, at least one required)
- **`--output`** — output signature / response shape (e.g. `200 OK application/json`, `stdout text`, `event payload`)
- **`--output-parameter`** — structured parameter entry: `"path|type|description"` (repeatable, at least one required)

ALL FOUR fields are REQUIRED — validator will error on missing ones.

Parameter format: each `--input-parameter` / `--output-parameter` takes a
pipe-separated spec `path|type|description`. Path uses JSON-path notation
for nested fields: `file.path`, `user.email`, `items[].id`. Type is a short
hint: `string`, `int`, `bool`, `array<User>`. Description is free text
(commas inside descriptions are preserved — the flag is literal-valued).
Example: `"user.email|string|email address of the requesting user"`.

**CLI contracts — enumerate every flag.** For a `cli` contract, document
each positional argument AND each CLI flag as its own `--input-parameter`.
Use the flag name as the path (e.g. `--description`) so readers see the
exact invocation surface. Group related flags by kind/cluster in the
description (e.g. `"component: boundaries"`, `"flow: narrative"`). The
`--input` field holds the one-line invocation signature; the parameters
are the full flag catalogue. A CLI contract with only positional params
documented is incomplete.

**Schema contracts — one per concrete schema, not per engine.** When the
contract is a data schema (KV key prefix, SQL table, protobuf message,
queue topic, cache key), create ONE contract per concrete schema, not a
single umbrella contract for the whole store.

- **Key-value indexes**: one contract per key prefix. `Input` is the key
  format template (e.g. `t:<tag>:<kind>:<id>`), `input_parameters` are
  each placeholder in the key, `output` is the value shape, and
  `output_parameters` are the fields of the value.
- **SQL tables**: one contract per table. `Input` is the fully-qualified
  table name (e.g. `public.users`), `input_parameters` list the columns
  (`path|type|description`) with type including constraints (`int PRIMARY KEY`,
  `text NOT NULL`, `timestamptz DEFAULT now()`), and `output_parameters`
  list the indexes / foreign keys / constraints.
- **Topics / channels**: one contract per topic.
- **Protobuf / JSON schemas**: one contract per message type.

A single "Index Schema" or "Database Schema" contract covering many
prefixes / tables is incomplete — split it.
- Relationship: **`belongs_to:<system-slug>`** — the system (process) that exposes this boundary
- Optional: `references:<concept-slug>` for concepts used in input/output
- Optional: `references:<component-slug>` for the internal component that implements it

Examples (name is descriptive; `--input` carries the invocation; `belongs_to` points at a **system**):
```
# REST endpoint
syde add contract "User Login" \
  --contract-kind rest --interaction-pattern request-response \
  --input "POST /auth/login" \
  --input-parameter "email|string|user email address" \
  --input-parameter "password|string|plaintext password (TLS only)" \
  --output "200 OK application/json; 401 on invalid credentials" \
  --output-parameter "access_token|string|short-lived JWT" \
  --output-parameter "refresh_token|string|long-lived token" \
  --output-parameter "user.id|string|user unique ID" \
  --output-parameter "user.email|string|confirmed email" \
  --add-rel "auth-api:belongs_to" \
  --add-rel "auth-middleware:references" \
  --add-rel "user:references"

# CLI command — enumerate every positional AND every flag as input parameters
syde add contract "Create Plan" \
  --contract-kind cli --interaction-pattern request-response \
  --input "myapp plan create <name> [--background --objective --scope]" \
  --input-parameter "name|string|positional, required. Plan name" \
  --input-parameter "--background|string|why this plan exists (context, motivation)" \
  --input-parameter "--objective|string|what success looks like when done" \
  --input-parameter "--scope|string|in-scope and out-of-scope summary" \
  --output "exit 0 on success; stdout prints plan ID and file path" \
  --output-parameter "plan_id|string|generated plan ID" \
  --output-parameter "file_path|string|absolute path to new plan file" \
  --add-rel "myapp-cli:belongs_to"

# Event
syde add contract "User Registered" \
  --contract-kind event --interaction-pattern pub-sub \
  --input "users.registered topic" \
  --input-parameter "user_id|string|unique user identifier" \
  --input-parameter "email|string|user email" \
  --input-parameter "timestamp|iso8601|registration time" \
  --output "fire-and-forget; consumed asynchronously" \
  --output-parameter "consumed_by|array<string>|list of downstream consumer services" \
  --add-rel "auth-api:belongs_to" \
  --add-rel "user:references"
```

### Concept rules (glossary)

Concepts are **domain glossary entries** — high-level terms that
explain what something *means* in the domain, not how it's shaped in
code. They are a design-level lens: they carry names, descriptions,
and prose fields explaining meaning, lifecycle, and invariants. They
do not carry typed attributes, actions, or cardinality-labelled
relationships — those belong in code (struct fields, methods) or in
contract schemas. The dashboard renders concepts as a standard
2-column inbox with a detail panel showing the prose fields plus
relationship chips grouped by role.

**Every concept MUST have:**

- `--meaning` — what the concept represents in the domain (enforced: blocking Finding)
- Relationship: `belongs_to:<system-slug>`

**Every concept SHOULD have:**

- `--invariants` — rules that must always hold (recommended: blocking Finding)
- `--lifecycle` — the state machine or progression through stages

**Every concept MAY have role-based relationships:**

| Relationship type | Target kind | Meaning |
|-------------------|-------------|---------|
| `implemented_by`  | component   | Which component implements the concept in code |
| `exposed_via`     | contract    | Which contract exposes the concept externally   |
| `used_in`         | flow        | Which flow operates on or produces this concept |
| `relates_to`      | concept     | Another concept this one relates to             |

`relates_to` is a plain edge with **no cardinality label** — prose
relationships, not schema-level cardinality. Keep multiplicity detail
inside the target component or contract schema.

**Worked example (Order / LineItem / Customer):**

```bash
# Parent domain
syde add system "Ecommerce" --description "Online store"

# Supporting scaffolding (components, contracts, flows)
syde add component "Order Service" \
  --purpose "Place, fulfil, and track customer orders" \
  --responsibility "CRUD for orders and line items, state machine enforcement" \
  --capability "Place order" --capability "Cancel order" --capability "Ship order" \
  --boundaries "Does not handle payments or shipping labels" \
  --add-rel "ecommerce:belongs_to"

syde add contract "Place Order" \
  --contract-kind rest --interaction-pattern request-response \
  --input "POST /orders" \
  --input-parameter "customer_id|string|placing customer" \
  --input-parameter "items|array<LineItem>|items to order" \
  --output "201 Created application/json" \
  --output-parameter "order.id|string|newly allocated order ID" \
  --add-rel "ecommerce:belongs_to" \
  --add-rel "order-service:references"

syde add flow "Place Order Flow" \
  --description "End-to-end placement from cart to confirmation" \
  --trigger "Customer submits checkout" \
  --goal "Order is persisted and confirmation sent" \
  --add-rel "ecommerce:belongs_to"

# Core glossary entries — prose only, no attributes or actions.
syde add concept "Order" \
  --description "A customer's purchase request" \
  --meaning "Groups line items into a billable transaction with a single delivery and payment." \
  --lifecycle "draft → placed → paid → shipped → delivered. Status transitions are forward-only." \
  --invariants "Must have at least one line item. Total > 0. Once shipped, cannot be cancelled." \
  --add-rel "ecommerce:belongs_to" \
  --add-rel "order-service:implemented_by" \
  --add-rel "place-order:exposed_via" \
  --add-rel "place-order-flow:used_in" \
  --add-rel "line-item:relates_to" \
  --add-rel "customer:relates_to"

syde add concept "LineItem" \
  --description "One product in an order" \
  --meaning "A single product-and-quantity entry within an order." \
  --invariants "Quantity > 0. Subtotal = quantity * unit price captured at purchase time." \
  --add-rel "ecommerce:belongs_to" \
  --add-rel "order-service:implemented_by" \
  --add-rel "order:relates_to"

syde add concept "Customer" \
  --description "A registered buyer" \
  --meaning "Person or organisation that can place orders." \
  --lifecycle "registered → active → deactivated" \
  --invariants "Email is unique and non-empty across active customers." \
  --add-rel "ecommerce:belongs_to" \
  --add-rel "order-service:implemented_by" \
  --add-rel "order:relates_to"
```

After running the above, open the dashboard and click **Concepts** in
the sidebar. The page is a standard 2-column inbox: a list on the left
and a glossary detail panel on the right. The detail panel shows
`Meaning`, `Lifecycle`, and `Invariants` as prose blocks, followed by
relationship chips grouped by role (`implemented by`, `exposed via`,
`used in`, `relates to`). Clicking any chip navigates to the target
entity.

**Migrating an older concept** that was written before the glossary
redesign? It may still carry attribute/action detail in YAML. Just
run `syde update <slug>` (any no-op call works) — the re-save strips
fields that no longer exist on the struct. Then add role-based
relationships with `--add-rel`, e.g. `syde update order-sk3a
--add-rel "order-service:implemented_by"`.

**Common mistakes:**

- Writing schema detail in `--meaning`. If you find yourself typing
  field names, types, or column lists, stop — those belong on the
  implementing component, the exposing contract's `input_parameters`,
  or a `storage` contract.
- Using the generic `references` relationship to a component. Use
  `implemented_by` instead so the role is explicit in the graph.
- Creating a concept that has no meaningful `implemented_by` or
  `exposed_via` link. A domain term that no code, contract, or flow
  touches is either too abstract to be useful or a sign that the
  implementing entity is missing from the model.

### System rules

A `system` represents a standalone process/app. Most projects have a single
top-level system, but a project shipping multiple binaries / daemons / apps
should model each as its own sub-system.

- **Top system**: no `belongs_to` — it is the root of the project.
- **Sub-system**: `belongs_to:<parent-system-slug>` — nests under another system.
- System nesting is acyclic (validator enforces).
- Components, contracts, and concepts `belongs_to:<system-slug>` of the
  system they are part of — for a sub-system, that's the sub-system, not the
  top-level one.
- Every entity except the single root system needs a `belongs_to` parent.
- Every non-requirement entity, including components and contracts, must carry
  an outbound relationship to at least one requirement.
- Every contract must participate in at least one flow.

Example: a project with a CLI binary and a daemon:
```
syde add system "MyApp"  # top-level
syde add system "MyApp CLI" \
  --description "The myapp CLI binary" \
  --add-rel "myapp:belongs_to"
syde add system "MyApp Daemon" \
  --description "The myappd HTTP daemon" \
  --add-rel "myapp:belongs_to"
```

### Requirement rules

Requirements capture architectural intent in **EARS shall-form** and are the
traceability backbone of the model. Every non-requirement entity must carry an
outbound relationship to at least one requirement.

**Statements MUST match one of the five EARS patterns.** The save validator
parses every statement and rejects any non-conforming requirement — there is
no way to persist a statement that does not match. The five patterns:

| Pattern | Template |
|---------|----------|
| Ubiquitous | `The <subject> shall <action>.` |
| Event-driven | `When <trigger>, the <subject> shall <action>.` |
| State-driven | `While <state>, the <subject> shall <action>.` |
| Unwanted-behavior | `If <unwanted condition>, then the <subject> shall <action>.` |
| Optional-feature | `Where <feature is included>, the <subject> shall <action>.` |

**Required fields on every requirement:**

- `--statement` — EARS shall-form text (validated on save)
- `--type` — one of `functional`, `non-functional`, `constraint`, `interface`,
  `performance`, `security`, `usability`
- `--priority` — MoSCoW: `must`, `should`, `could`, `wont`
- `--verification` — short free-form description of how the requirement is
  verified (e.g. "integration test hits /api/projects and asserts 200",
  "manual review of CLI help output")
- `--source` — `user`, `plan`, `migration`, or `manual`
- `--rationale` — why the requirement exists

The legacy acceptance field has been **removed** from requirements.
`--acceptance` is no longer a flag on `syde add requirement` (it still exists
on `syde task`). Use `--verification` instead.

**Requirements never carry a `files` list.** They are pure design intent and
link only to other entities via `refines`, `derives_from`, and `belongs_to`.
Do not pass `--file` when creating a requirement.

**Requirement relationships:**

- `refines` — requirement → component / contract / concept / system, or
  requirement → requirement. Use when a requirement narrows a higher-level
  intent against a specific design target. Example: a performance requirement
  `refines` the HTTP API component.
- `derives_from` — requirement → parent requirement. Use when a child
  requirement is logically implied by a parent (derivation chain).
- `belongs_to` — requirement → system, same as every other entity.

**Backfilling requirements from an existing codebase.** When the design model
already has components / contracts / concepts / systems but lacks requirements,
follow the deterministic algorithm in
`references/requirement-derivation.md`. It produces one EARS statement per
architectural property with stable slugs and the correct `refines` /
`derives_from` chain, and is the canonical way subagents backfill the
requirement layer.

### Validation

`syde validate` enforces:
- Components must have `purpose`, `responsibility`, `capabilities`
- Contracts must have `contract_kind`
- Requirements are append-only and must be `active`, `superseded`, or `obsolete`
- Requirement statements must match one of the five EARS patterns (save-time)
- Requirements must have `req_type`, `priority`, and `verification`
- Requirements must not carry a `files` list
- Superseded requirements must point to replacements; obsolete requirements must say why
- Every non-requirement entity has an outbound relationship to a requirement
- Every entity except the root system has `belongs_to`
- Every contract has at least one flow relationship
- No cyclic `depends_on` relationships between components
- All relationship targets must exist

#### Requirement overlap resolution (MERGE / RENAME / DISTINCT)

`syde add requirement` surfaces any existing active requirement whose
statement is above 0.6 TF-IDF similarity to the new one. **Text
similarity is only a candidate signal — semantic overlap is your job
to verify.** The CLI blocks the create (non-zero exit) unless every
surfaced overlap is resolved by one of three outcomes:

| Outcome | When | How |
|---------|------|-----|
| **MERGE** | The two requirements mean the same thing. | Abandon the new one; reuse the existing slug. Update inbound `refines` / `derives_from` / `references` to point at the survivor. |
| **RENAME** | They are semantically distinct but the statements accidentally share vocabulary. | Rewrite the new statement with different domain terms so the TF-IDF match drops below 0.6. Retry the create. |
| **DISTINCT** | They are genuinely close cousins that must both exist. | Retry with `--audited <slug>:"why these two mean different things"` for each overlap. The distinction text is persisted on the requirement and must be ≥20 characters of substantive reasoning. |

The installed PostToolUse hook also injects a system reminder into
the session whenever `syde add requirement` prints overlap
candidates, so the decision can't be skipped silently.

**Audit enforcement.** `syde sync check` errors on any acknowledged
overlap whose distinction is empty or shorter than 20 chars — rubber
stamps are treated as unresolved. If you acknowledged an overlap with
a throwaway string, revisit it and write a real semantic reason.

**Contract coverage.** When a requirement's statement names a
CLI invocation (`syde <sub>`), REST path (`GET /api/...`), dashboard
screen, or pub-sub event topic, an active contract must cover that
surface — or the plan must introduce/extend one in the same diff.
The planning-time and sync-check Findings are symmetric: the
gap is caught before the plan is approved and against the corpus at
rest.

**Flow coverage.** When a plan introduces or extends a contract, the
plan's flow lane must introduce or extend a flow whose `steps`
reference that contract. The planning-time Finding mirrors the
existing post-plan `contractFlowFindings` rule so no contract ever
lives without at least one flow traversing it.

### Phase 2: CREATE PLAN

A plan has five levels of detail:
1. **Plan header**: `background`, `objective`, `scope` — why, what success is, what changes at a high level
2. **Design**: prose implementation design — the detailed "how we're going to build it" narrative
3. **Changes**: structured diff of every entity that will be deleted, extended, or newly created, organized into six per-kind lanes (requirements, systems, concepts, components, contracts, flows)
4. **Phases**: `objective`, `changes`, `details`, `notes` — per-milestone plan
5. **Tasks**: `objective`, `details`, `acceptance` — per-work-item plan, referencing real entities and files

Every plan MUST answer: **WHY** (background), **WHAT SUCCESS LOOKS LIKE**
(objective), **WHAT CHANGES AT A HIGH LEVEL** (scope), **HOW WE'LL BUILD IT**
(design), and **WHICH ENTITIES MOVE** (structured changes). Without these the
plan is unreviewable.

Per REQ-0331, every plan records its entity changes as a structured diff with
deleted / extended / new entries, and the completion validator verifies that
diff against actual entity state before the plan can be marked `completed`.

#### Requirement-first cascade

Before authoring any implementation change, identify the underlying requirements
first. **Every design decision in the plan must trace back to a requirement.**
A plan with 5 phases and 2 requirements is a red flag — split each behavioral
property, constraint, and UI expectation into its own EARS statement.

**Step 1 — Search existing requirements.** Use `syde query --kind requirement`
and `syde query --search "<term>" --kind requirement`. If an existing
requirement overlaps, conflicts, or is superseded by the user's request, mark
the old requirement `superseded` / `obsolete` instead of silently replacing
intent.

**Step 2 — Decompose into granular requirements.** For each aspect of the
request, ask: "what specific property must the system have?" Each answer is
one requirement. Categories to check:

- **Model properties** — new fields, struct shapes, cardinality constraints
- **Audit / validation rules** — each rule is its own requirement (what triggers it, what it checks; every audit Finding blocks)
- **UI / rendering behavior** — each visual expectation (colors, layouts, interactions)
- **Workflow constraints** — ordering rules, one-per-X rules, naming conventions
- **Organization / grouping** — how entities are categorized, filtered, tagged
- **Migration / cleanup** — what gets deleted, what replaces it

A good heuristic: if the implementation needs a separate `if` branch or a
distinct test case, it deserves its own requirement.

**Step 3 — Add requirements to the plan's Requirements lane** before declaring
implementation changes. Then cascade the implementation lanes in this order:
Requirements → Components → Contracts → Concepts → Flows. Requirements are
the why; the other lanes are the how. Never invert that order.

**Step 4 — Link every change to requirements.** Each component/contract/flow
change should trace back to one or more requirements via `--task` links.
Orphan implementation changes with no requirement backing are a Finding
in `syde plan check`.

Always check existing flows. If the request changes behavior captured by a
flow entity, extend that flow in the plan with `field_changes` for the changed
`happy_path`, `narrative`, or `edge_cases`; flows rot silently when authors
forget to update them.

#### Step 1: Create plan with background / objective / scope / design

```
syde plan create "<name>" \
  --background "Why does this plan exist? What problem or context drives it?" \
  --objective  "What does success look like when the plan is done?" \
  --scope      "What's in-scope and out-of-scope at a high level" \
  --design     "Detailed implementation design: architecture choices, data flow, key tradeoffs, migration steps"
```

You can update these later with `syde plan update <slug> --background ... --objective ... --scope ... --design ...`.

The `--design` field is free-form prose. Use it to capture the implementation
walkthrough the phases gloss over: the overall approach, the key data-shape
choices, what gets built in which order, and why. Reviewers read the design
section before they skim phases, so write it like a short RFC, not like a
commit message.

The background/objective/scope/design sections render at the top of `syde plan
show` and in the dashboard plan detail page so reviewers see the why/what/how
before any phase detail.

#### Step 2: Declare the structured change diff

Before creating phases or tasks, spell out **every entity that will be touched**
as a structured diff. syde models this as six per-kind lanes (requirements,
systems, concepts, components, contracts, flows), each carrying three lists:
`deleted`, `extended`, `new`. Every entry has `what` and `why` prose; extended
entries may also carry programmatically-verified `field_changes`; new entries
carry a kind-specific `draft` map that seeds the entity when it's created.

Build the diff with `syde plan add-change`:

```
# Mark an entity for deletion
syde plan add-change delete <plan-slug> component legacy-auth \
  --why "Replaced by new JWT middleware; no callers remain after phase 2"

# Extend an existing entity; field_changes are verified at completion time
syde plan add-change extend <plan-slug> component api-server \
  --what "Add /api/plans routes and wire the plan completion handler" \
  --why  "REQ-0331 requires a dashboard-visible plan inbox" \
  --field responsibility="Request routing, validation, plan lifecycle endpoints" \
  --field boundaries="No direct DB access; delegates to storage layer"

# An extended change with no --field is a hand-review note (completion validator skips programmatic verification for it)
syde plan add-change extend <plan-slug> flow session-start-bootstrap \
  --what "Mention the plan inbox in the bootstrap narrative" \
  --why  "Reviewers land on the inbox first once this ships"

# Use the sentinel DELETE to assert a field must become empty
syde plan add-change extend <plan-slug> component memory-sync \
  --what "Retire the learnings pipeline" \
  --why  "Learnings entity was removed in REQ-0319" \
  --field files=DELETE

# Declare a brand-new entity; --draft seeds it at entity-create time
syde plan add-change new <plan-slug> contract \
  --name "Plan Completion API" \
  --what "POST /api/plans/:slug/complete returning the validator findings" \
  --why  "Dashboard reviewers need a one-click gate" \
  --draft contract_kind=rest \
  --draft interaction_pattern=request-response \
  --draft input="POST /api/plans/:slug/complete" \
  --draft 'input_parameters=[{"path":"slug","type":"string","description":"plan slug"}]' \
  --draft output="200 OK application/json"
```

Both `--field` and `--draft` are **repeatable**. `--draft` auto-decodes JSON
literal values (arrays, objects, numbers, booleans), so structured fields like
`input_parameters` or `files` can be passed as JSON without escaping ceremony.
The sentinel value `DELETE` on a `--field` means "this field must be the zero
value when the plan completes".

Once tasks exist, every deleted / extended / new change must list the task
slug(s) that implement it with repeatable `--task <task-slug>` flags. If you
draft changes before creating tasks, return to the change list after Step 3 and
make sure every change is claimed by at least one task in the same plan.

You can drop a change that's no longer needed:
```
syde plan remove-change <plan-slug> <change-id>
```

Review the full diff at any time:
```
syde plan show-changes <plan-slug>                  # rich, grouped by kind
syde plan show-changes <plan-slug> --format json    # machine-readable
```

The dashboard surfaces this diff on the per-plan detail page (see Step 5) with
side-by-side field renders for extended entries and draft previews for new
entries. Reviewers should be able to read the plan, design, and diff in the
browser before they approve.

#### Step 3: Create phases with tasks

Each phase is a **deliverable milestone** with multiple **tasks** (the actual work).

```
syde plan add-phase <plan-slug> --name "Scaffolding" \
  --description "Get React app serving from Go binary with dark theme" \
  --objective  "React app is served from the Go binary with a dark theme applied" \
  --changes    "New web/ dir, go:embed dist/, Makefile targets, html.go replaced" \
  --details    "Implementation: create Vite project, configure Tailwind, update embed.go, wire Makefile" \
  --notes      "Bun must be installed; skill/hooks.json must still load"

syde task create "Create Vite project and install deps" --plan <plan-slug> --phase phase_1 \
  --objective  "Vite + React + TS scaffolded in web/" \
  --details    "bun create vite web --template react-ts; install tailwind, react-router" \
  --acceptance "web/package.json exists; bun run dev serves on :5173"

syde task create "Configure Tailwind dark theme" --plan <plan-slug> --phase phase_1 \
  --objective  "Dark theme tokens apply globally" \
  --details    "Tailwind v4; index.css with CSS variables for bg/fg/muted/border" \
  --acceptance "Root page renders dark; kind colors visible"
```

The plan shows two levels with inline objectives: **phase → tasks with objective**.

```
○ Scaffolding — pending [4 tasks]
    Objective: React app is served from the Go binary with a dark theme applied
    Changes:   New web/ dir, go:embed dist/, Makefile targets, html.go replaced
  ○ create-vite-project — pending
      Objective: Vite + React + TS scaffolded in web/
  ○ configure-tailwind — pending
      Objective: Dark theme tokens apply globally
```

**Every phase MUST have:**
- `--name` — short label
- `--description` — what this phase delivers
- `--objective` — what this phase achieves (1 sentence, success condition)
- `--changes` — concrete list of things that change (files, entities, behavior)
- `--details` — implementation walkthrough (HOW to build)
- **Granular tasks** — concrete work items, not broad placeholders. Aim for
  several tasks per phase when a phase changes more than one concern. Each task
  should be small enough to start, finish, verify, and mark done independently.

Before presenting a plan for approval, create the full task list with
`syde task create --plan <plan-slug> --phase <phase-id> ...`. A phase with no
tasks is invalid and `syde plan approve` rejects it. Do not ask for approval
while any phase still has zero tasks.

**Every task MUST have:**
- `--objective` — what the task achieves
- `--details` — approach / files to touch
- `--acceptance` — how to know it's done (observable)

A phase cannot be completed until ALL its tasks are done.

Use `--notes` for reminders, risks, or context.

#### Large plans: use 3 levels

For large plans (>20 tasks), use **parent phase → child phase → tasks**, but
do not leave the parent phase taskless. Parent phases need at least one direct
coordination/integration task before approval.

```
# Parent phase (milestone)
syde plan add-phase <plan-slug> --name "Frontend" \
  --description "Complete React SPA with all views" \
  --details "Milestone: all frontend views working against live API"
syde task create "Coordinate frontend milestone integration" --plan <plan-slug> --phase phase_1 \
  --objective "Keep child phases integrated and verify the whole frontend milestone"

# Child phases (deliverables within the milestone)
syde plan add-phase <plan-slug> --name "Layout + Sidebar" --parent phase_1 \
  --description "Three-column layout with sidebar navigation" \
  --details "Sidebar with kind groups, main area, detail panel"

# Tasks on the child phase
syde task create "App layout component" --plan <plan-slug> --phase phase_2
syde task create "Sidebar with kind groups" --plan <plan-slug> --phase phase_2
```

This shows three levels:
```
○ Frontend — pending [9 tasks]
  ○ coordinate-frontend-milestone-integration — pending
  ○ Layout + Sidebar — pending [2 tasks]
    ○ app-layout-component — pending
    ○ sidebar-with-kind-groups — pending
  ○ Entity Views — pending [3 tasks]
    ○ entity-list — pending
    ○ entity-detail — pending
    ○ relationship-chips — pending
```

Parent phases aggregate all descendant tasks and entities, while still carrying
at least one direct task of their own. A parent cannot be completed until ALL
children are completed. Children cannot be completed until ALL their tasks are
done.

**When to use 3 levels:** >20 tasks, or the plan spans multiple sessions.
**When to use 2 levels:** <20 tasks, single session plan.

#### Step 4: Declare what each task affects

Each task references the **entities** and **source files** it will touch. The
structured change diff from Step 2 is the plan-level statement of intent; the
per-task `--affected-*` flags are what the drift validator uses at `task done`
time to bump `updated_at` on the right entities.

For every task you create, set:
- `--affected-entity <slug>` — a slug of an entity this task will modify. If
  the change diff declares a `new` entry for this kind + name, the slug will
  be the one produced when `syde add` runs during implementation; pass it
  back into the task via `syde task update --affected-entity` once it exists.
  Repeatable. Validator rejects slugs that don't resolve.
- `--affected-file <path>` — a concrete source file this task will touch.
  Must exist as a node in `.syde/tree.yaml` (run `syde tree scan` first).
  Repeatable.

Example:
```
syde task create "Harden JWT middleware" \
  --plan add-auth --phase phase_1 \
  --objective "Rotate signing key without downtime" \
  --affected-entity auth-middleware \
  --affected-entity jwt-config \
  --affected-file internal/auth/middleware.go \
  --affected-file internal/auth/keys.go
```

**If a task needs to CREATE a brand-new entity**, it MUST be declared in the
plan's structured diff via `syde plan add-change new <plan> <kind> --name
"..." --draft ...` (Step 2). That's the plan-level record of intent. When the
task executes, run `syde add <kind> "<name>" ...` using the same name and the
draft fields as defaults, then attach the freshly-created slug to the task
with `syde task update <task> --affected-entity <new-slug>`. The plan
completion validator will refuse to close the plan unless every declared
`new` change corresponds to a real entity with a matching name + kind.

#### Step 5: Estimate and present

```
syde plan estimate <plan-slug>
syde plan check <plan-slug>
syde plan open <plan-slug>
syde plan show <plan-slug> --full
syde plan show-changes <plan-slug>
```

Always use `--full` when presenting — shows phase details, notes, and
task-by-task status with affected entities / files. Pair it with
`show-changes` so the structured diff is visible alongside the phase tree.

Run `syde plan check <plan-slug>` after drafting the structured diff, phases,
and tasks. Address every Finding before presenting for approval — there is
no non-blocking severity tier. `syde plan check` must exit 0.

Run `syde plan open <plan-slug>` before asking for approval. This reuses an
existing dashboard tab via WebSocket navigation when one is connected, or opens
a new browser tab when none are listening. The plan must be visible in the
browser before you ask the user to approve it.

The dashboard has a **Plans inbox** at `/<project>/plan` and a per-plan
detail page at `/<project>/plan/<slug>` with two tabs: **Plan** (design +
structured changes grouped by kind, with side-by-side diffs for extended
entries and draft previews for new entries) and **Tasks** (phases collapsed
with nested task rows). The standalone Tasks sidebar nav is gone — tasks now
live only inside the plan detail page. Point reviewers at the plan detail URL
so they can read design + diff in the browser before approving in chat.

Tell the user: "Plan ready. Review in the dashboard or via `syde plan show`.
Approve to proceed, or suggest changes."

**Approval rule — explicit chat approval required.** You may run `syde plan
approve <plan-slug>` yourself, BUT only after the user has explicitly approved
the plan in chat (e.g. "approve", "ok approve", "looks good, approve", "go
ahead"). Ambiguous or neutral acknowledgements ("ok", "sure", "nice") are NOT
approval — ask for confirmation. This chat approval is the explicit
human-in-the-loop gate for every plan. Never approve your own plan without
the user's explicit go-ahead.

Approval also requires task coverage: every phase must already have at least
one concrete task. If `syde plan approve` reports an empty phase, add granular
tasks for that phase, show the updated plan, and ask for approval again.

Approving a plan creates a plan-sourced requirement and links the plan to it.
Treat all user prompts and approved plans as durable requirements: if intent
changes later, create a new requirement and mark the older one superseded or
obsolete; never delete requirement history.

**STOP. Do NOT implement until the plan is approved.** Check the plan status
with `syde plan show <plan-slug>` — if it says `draft`, you are not allowed to
start any tasks. Only after the user approves in chat, run `syde plan approve
<plan-slug>` and then begin phase 1.

### Phase 3: IMPLEMENT (one phase at a time)

Work through the plan **one phase at a time**. Complete all tasks in a phase
before moving to the next. Never jump ahead.

**Do not stop until every task in every phase of the approved plan is done.**
Once the user has approved a plan, executing it through to completion is a
commitment, not a series of mid-flight check-ins. Do NOT pause to ask "should
I continue?" between tasks or phases. Do NOT batch tasks and present them for
review unless the task itself explicitly requires user input. The only valid
reasons to stop mid-plan are:

1. A task hits a real blocker (failing build you can't fix, missing
   credential, ambiguous requirement that needs clarification).
2. You discover the plan itself was wrong and needs revision — in which case,
   say so explicitly and propose the revised plan.
3. The user interrupts you.

If you finish a task and the next one is straightforward, just start it.
Status updates are fine; permission requests are noise. The user already
approved the plan — running it is what they're paying for.

**CRITICAL: Write / Edit / MultiEdit / NotebookEdit are HARD-BLOCKED by the
PreToolUse hook when either condition is true:**

- **No approved plan** — the hook refuses all writes until you run
  `syde plan create <name>`, present it to the user, receive explicit
  chat approval, and run `syde plan approve <slug>`. The plan does not
  have to be the *only* plan — any plan in state `approved` or
  `in-progress` unlocks writes for the session.
- **No active task** — even with an approved plan, every write requires
  an `in_progress` task (`●` in `syde task list`). Run
  `syde task start <slug>` before editing code.

The block fires with exit code 2 and a stderr message explaining which
condition failed. Excluded paths (safe to write without a task): files
under `.syde/`, `.claude/`, `node_modules/`, `vendor/`, `web/dist/`,
`web/node_modules/`, `.git/`, and `/tmp/`. Everything else — source
code, documentation, tests, configs — goes through the plan → task
gate.

This is non-negotiable. If you hit the block, do not work around it by
moving files to excluded paths. Create a task, start it, then write.

**Keep syde tasks and the visible todo/checklist tool in sync.** The syde plan
is the source of truth, and the runtime todo tool mirrors it for visibility:
Claude uses TodoWrite; Codex uses `update_plan`. Before starting a phase, mirror
all of that phase's syde tasks into the visible todo list with matching task
names. When you run `syde task start`, mark that todo `in_progress`. When you
run `syde task done`, mark that todo `completed`. If you add, split, block, or
rename a syde task, immediately update the visible todo list to match.

#### For each task in the phase:

**BEFORE — gather context:**
1. Read the phase details: `syde plan show <plan-slug> --full`
2. Read related entities: `syde query <entity> --full` for each entity involved
3. Check files to read: look at the entity `files` field for relevant source paths
4. Start the task:
   ```
   syde task start <task-slug>
   ```

**DURING — write code:**
5. Write the code
6. Verify new files: `syde constraints check <file>`
7. Update entities as you go — don't wait until later:
   - New entity needed? → `syde add <kind> <name>` with `--file` references
     and outbound requirement traceability such as
     `--add-rel <requirement>:references`
   - Entity changed? → `syde update <slug>` with updated fields
   - New relationship? → `syde update <slug> --add-rel "<target>:<type>"`
   - Architectural choice made? → `syde add requirement "<name>" --statement "The system shall ..." --type functional --priority must --verification "..." --source manual --rationale "..."`

**AFTER — complete and verify:**
8. Complete the task and **declare the real affected set** — this is
   how drift clears correctly at the end of the session:
   ```
   syde task done <task-slug> \
     --affected-entity <slug1> --affected-entity <slug2> \
     --affected-file path/one.go --affected-file path/two.go
   ```
   The `--affected-entity` / `--affected-file` flags at done time are
   MERGED into whatever was predicted at create time. Pass every entity
   you actually modified and every file you actually touched — this
   is the authoritative set the drift cascade uses to bump `UpdatedAt`
   on each affected entity, and it's what makes `syde sync check`
   pass at the end of the session. Create-time affected lists are
   predictions; done-time affected lists are reality, and reality
   wins. Never call `syde task done` without these flags if the task
   touched anything — the omission will surface as drift warnings
   during the finish gate, and you will have to go back and update
   entities manually.
9. Verify entity completeness:
   - Are all new/changed files referenced? → `syde update <entity> --file "path/*.go"`
   - Are all relationships wired? → `syde update <entity> --add-rel "<target>:<type>"`
   - Do descriptions still match reality? → `syde update <entity> --description "..."`
   - Run `syde sync --check` if unsure — must show 0 gaps
10. Check the plan: `syde plan show <plan-slug>`
    - Task should show ✓
    - When ALL tasks in a phase are done, the phase auto-completes

**Do NOT batch completions.** Complete each task as soon as the work is done.
The plan should reflect reality at all times.

#### Adapting tasks on the fly

Tasks are flexible. Adjust as you learn:

- **Add new tasks**: `syde task create "<new task>" --plan <plan-slug> --phase <phase-id>`
- **Split a task**: `syde task sub <parent-slug> "<subtask name>"`
- **Block a task**: `syde task block <slug> --reason "Waiting for X"`
- **Update phase notes**: `syde plan phase <plan-slug> <phase-id> --notes "Discovered X"`

The plan is a living document — adapt based on what you discover.

**Never block on missing tools or dependencies.** If a task requires installing
a tool, library, or dependency — install it and continue. Do NOT mark tasks as
blocked just because something needs to be installed. Only block when the
blocker is outside your control (e.g., waiting on user input, external API key,
another team's work). If you can resolve it yourself, resolve it and keep going.

#### Completing a phase

When all tasks are done, the phase auto-completes. Then:
1. Review: do all entities reflect what was actually built? Update any drift.
2. Sync the model (Phase 4) before starting the next phase.

### Phase 4: SYNC MODEL (after every code change)

The design model must stay in sync with the code at all times. After implementing
each phase, update the model before moving on:

1. **New component or module created** → `syde add component` or `syde update` existing one
   - Set `responsibility`, `boundaries`, `data_handling` at minimum
   - Add relationships: `--add-rel <other>:depends_on`, `--add-rel <contract>:exposes`
2. **API endpoint or interface added/changed** → `syde add contract` or `syde update`
   - Name must be descriptive ("List Projects"), invocation goes in `--input`
   - Set `contract_kind`, `interaction_pattern`, `input`, `input_parameters`, `output`, `output_parameters` (all required)
3. **New domain model or data type** → `syde add concept` or `syde update`
   - Set `meaning`, `invariants`, `lifecycle`
   - Set `data_sensitivity` for anything containing PII or secrets
4. **New user-facing workflow** → `syde add flow` or `syde update`
   - **One flow per user goal** — not broad categories. "Create Plan", not "Planning".
   - Set `trigger`, `goal`, and **structured steps** via `--step`
   - Each step: `--step "id|action|contract|description|on_success|on_failure"`
   - Steps describe what user and system do; action text names the actor
   - One contract per step; if a behavior touches two contracts, use two steps
   - Tag flows by category: `--tag planning`, `--tag authoring`, `--tag dashboard`, etc.
   - The audit ERRORs on contracts not referenced by any flow step
   - Legacy prose fields (`narrative`, `happy_path`, `edge_cases`, `failure_modes`) are still supported but prefer structured steps
5. **Architectural choice made during implementation** → `syde add requirement`
   - Set `statement` (EARS shall-form, save-validated), `--type`, `--priority`,
     `--verification`, `--source`, `--rationale`
   - Link with `--add-rel <target>:refines` (component/contract/concept/system
     or a parent requirement) and/or `--add-rel <parent-req>:derives_from`
   - Never pass `--file` on a requirement — requirements have no files list
   - Even implicit choices count (e.g., choosing a library, picking a pattern) — capture them as requirements so intent is preserved
6. **Relationship changes** → `syde update <entity> --add-rel <target>:<type>`
   - New dependency? Add `depends_on`
   - Component exposes a contract? Add `exposes`
   - Flow involves a component? Add `involves`

**The test**: after every phase, run `syde context` and ask yourself:
"If someone new read this model, would they understand the system as it exists
right now?" If the answer is no, the model is incomplete — update it before
proceeding.

### Phase 5: FINISH

1. `syde task list` — **ALL tasks must be completed**. No `pending` or `in_progress`
   tasks may remain. If any exist, complete them now.
2. `syde plan show <plan-slug>` — must show 100% complete.
3. **Close the plan with the completion validator:**
   ```
   syde plan complete <plan-slug>
   ```
   This runs **the full `syde sync check`** and refuses to mark the plan
   `completed` if ANY errors exist anywhere in the model — not just
   plan-specific findings. The gate checks:
   - Plan-completion diff: `delete` targets must not exist, `new` targets
     must exist, `extend` field_changes must match actual values.
   - **All sync check errors**: broken relationships, cycles, missing fields,
     orphan files, TF-IDF overlaps without `--audited`, missing `belongs_to`,
     etc. The entire model must be clean before a plan can complete.
   Fix every error `syde sync check` reports before running `plan complete`.
   `--force` overrides with a warning — only use it if you've explained to
   the user why the errors are intentionally accepted.

   **Do NOT call `syde plan complete` until `syde sync check` exits 0.**
   Run `syde sync check` first, fix all findings, then complete.
4. **Refresh the summary tree** (the rest of the gate reads from it):
   ```
   syde tree scan
   syde tree changes --leaves-only --format json
   # dispatch subagents in parallel to summarize stale files (see Phase 0)
   # then summarize stale folders in the main session using `tree show`
   ```
5. **`syde sync check` — the one gate, always strict.** This is the
   canonical session-end command. The audit emits one severity level
   only — every Finding blocks. Exit 0 means the session can end;
   exit 1 means at least one Finding exists. There is no `--strict`
   flag and no non-strict mode. Fix every Finding it prints, then
   rerun until it exits 0:
   ```
   syde sync check
   ```
   Under the hood `sync check` bundles five previously-separate checks
   so you only need one command at the end:
   - **Structural audit** — required fields on every entity, recommended fields, broken relationship targets, cyclic `belongs_to` / `depends_on`.
   - **Tree ↔ entity consistency** — every `--file` ref exists as a real tree node.
   - **Orphan detection** — every non-ignored source file is owned by at least one entity. Use `syde files orphans` for a targeted list and `syde files coverage <path>` to see who owns a specific file.
   - **File drift** — file mtime newer than the owner entity's `updated_at` becomes a Finding (fixed by `syde task done` with `--affected-entity/--affected-file`, or a direct `syde update`).
   - **Summary-tree staleness** — any stale file or folder node in `tree.yaml`.
   - **Plan completion drift** (`planCompletionFindings`) — any approved or in-progress plan whose declared change diff no longer matches entity state. Fix by running `syde plan complete` and resolving its Findings.
   - **Requirement overlap distinction** — any acknowledged overlap whose distinction text is empty or shorter than 20 characters fails as a rubber stamp.
   - **Contract surface coverage** — any active requirement whose statement names a CLI / REST / screen / event surface that no active contract covers.
6. `syde context` — final sanity scan. Read through the snapshot and
   confirm it describes the current state of the code.

**You cannot commit or push code with incomplete tasks or a failing
`syde sync check`.** The PostToolUse Bash hook warns on `git commit`
/ `git push` when tasks are pending or the gate fails, and the Stop
hook blocks session end with the same gate.

> `syde validate` still works as a deprecated alias for
> `syde sync check`. Prefer `syde sync check`.

## Rules

1. **Never read `.syde/` files directly** — always use syde CLI commands
2. **Never skip phases** — Phase 0 (tree survey) → clarify → plan (tasks reference existing entities and files via `--affected-*`) → approve → implement → sync → finish (with tree refresh)
3. **Never implement before plan approval and an active task** — the
   PreToolUse hook HARD-BLOCKS every Write / Edit / MultiEdit /
   NotebookEdit when either (a) no plan is in `approved` or
   `in-progress` state, or (b) no task is `in_progress`. You cannot
   work around this by editing files; you must create a plan, get it
   approved in chat, run `syde plan approve`, then `syde task start
   <slug>` before any code change. The block prints a stderr message
   explaining which condition failed.
4. **Never leave the model out of sync** — every code change that affects architecture must be reflected in syde entities. No exceptions. If you added a file, a function, an endpoint, a dependency, or made a design choice — update the model. At task completion time, always pass `syde task done <slug> --affected-entity ... --affected-file ...` declaring the **real** set of entities and files you touched. The done-time flags are merged into the task's existing affected lists, validated, and used by the drift cascade to bump `UpdatedAt` on each affected entity. Omitting them leaves drift Findings in `syde sync check` that you'll have to clear manually. See Phase 3 step 8 for the full invocation.
5. **Always run `syde tree scan` at session start** and `syde sync check` at session end. Between them, keep the tree clean via the leaves-first summarize loop. Stale tree = rotted understanding for the next session.
6. **Delegate file summarization to subagents, never burn main context on it.** When summarizing stale files for the tree, dispatch subagents in parallel and give each a batch of paths. Each subagent calls `syde tree context <path>` + `syde tree summarize`. Main session only handles folder summaries (cheap, derived from children via `syde tree show`).
7. **When creating entities on an existing codebase, use `syde tree context <path>`, never naive `Read`.** The tree context returns the ancestor breadcrumb + file summary + content in one call — that's the right framing for `--purpose`, `--responsibility`, `--boundaries`. `Read` is only for verification.
8. **Always verify new source files** — run `syde constraints check` after writing
9. **Always document thoroughly** — entity descriptions must be specific, not vague. "Handles auth" is not enough. "Manages JWT token issuance, refresh, and revocation. Validates tokens on every request via middleware. Does NOT own user profiles — delegates to user-service." is what we need.
11. **Use `syde query` for targeted lookups** — architecture is already in your context from session start, use query only when you need full detail on a specific entity
12. **Bulk entity creation goes through a shell script, not individual tool calls.** For bootstrap (many entities at once), write a `/tmp/syde-<batch>.sh` and run it in one bash call. One-off creates during task implementation can still be direct `syde add` calls.
13. **Clear orphan files deliberately** — every ERROR from `syde sync check` is either (a) a file that should join an existing component's `--file` list, (b) a file that warrants a new component, or (c) a file that isn't part of the design model and should be `syde tree ignore`d. Use `syde files orphans` to list them and `syde files coverage <path>` to check ownership. Do not mass-ignore to silence the gate.

## Reference Files

- `references/clarify-guide.md` — critical requirement gathering by project type
- `references/entity-spec.md` — entity kinds, fields, valid values
- `references/commands.md` — complete CLI reference with all flags
- `references/sync-workflow.md` — syncing design model with codebase (new projects + existing models)
- `references/requirement-derivation.md` — deterministic algorithm for backfilling EARS requirements from existing components / contracts / concepts / systems
