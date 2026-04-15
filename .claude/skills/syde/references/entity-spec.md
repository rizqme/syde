# Entity Specification

## Entity Kinds

| Kind | What it represents |
|------|-------------------|
| `system` | The overall product/service |
| `component` | A service, module, or major part |
| `contract` | An API, event, or protocol boundary |
| `concept` | A domain object or business entity |
| `flow` | An end-to-end workflow or user journey |
| `decision` | An architectural decision (ADR) |
| `requirement` | Append-only user/plan/migration requirement |
| `plan` | A tracked implementation plan |
| `task` | A work item linked to plans/entities |
| `design` | A UI mockup in UIML format |
| `learning` | Captured design knowledge |

## Shared Fields (all entity kinds)

| Field | YAML key | Description |
|-------|----------|-------------|
| ID | `id` | Auto-assigned counter-based identifier: `<PFX>-<NNNN>` (e.g. `SYS-0001`, `COM-0002`, `CON-0003`). Never reused. |
| Kind | `kind` | Entity type |
| Name | `name` | Human-readable name |
| Slug | `slug` | File-level addressable slug: `<name-slugified>-<rand4>` (e.g. `cli-a3f2`). Generated on create, stable across renames. Commands accept the full slug, the bare name-slugified base (when unambiguous), or a parent/child path (`<parent-slug>/<child-slug>`). |
| Description | `description` | Short one-sentence identification of what the entity is. **REQUIRED on every kind except plans/tasks and requirements with `statement`** — `syde validate` errors when missing. Keep distinct from `purpose` (the why) and `responsibility` (the what does). |
| Purpose | `purpose` | Why it exists |
| Tags | `tags` | Label list |
| Notes | `notes` | Informal notes list (`--note`, repeatable) |
| Files | `files` | Source file paths (repeatable `--file`). **Concrete paths only — NO wildcards.** Each entry must resolve to a node in `.syde/tree.yaml`. The validator rejects globs like `internal/storage/*.go`. Use `syde tree show <folder>` to enumerate, then pass each path individually. |
| Updated At | `updated_at` | Auto-bumped when `syde task done` completes a task that lists this entity in `--affected-entity`, or when a task's `--affected-file` matches an entity's `files`. Used by the drift validator — if a file's `mtime` in the tree is newer than the owning entity's `updated_at`, a warning is raised. |
| Relationships | `relationships` | Typed links to other entities |

## Kind-Specific Fields

### System
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `context` | `--context-text` | System context |
| `scope` | `--scope` | System scope |
| `design_principles` | `--design-principles` | Design principles |
| `quality_goals` | `--quality-goals` | Quality goals |
| `assumptions` | `--assumptions` | Assumptions |

### Component (purpose REQUIRED from base)
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `responsibility` | `--responsibility` | What it does (**required**) |
| `capabilities` | `--capability` | Concrete capabilities (repeatable, **at least one required**) |
| `boundaries` | `--boundaries` | What it does NOT do |
| `behavior_summary` | `--behavior-summary` | Behavior summary |
| `interaction_summary` | `--interaction-summary` | Interaction summary |
| `data_handling` | `--data-handling` | Data handling |
| `scaling_notes` | `--scaling-notes` | Scaling notes |
| `failure_modes` | N/A (use body) | Failure modes list |

### Contract
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `contract_kind` | `--contract-kind` | `rest`, `cli`, `event`, `rpc`, `graphql`, `websocket`, `pubsub`, `storage` (data schemas — KV prefix, SQL table, proto, cache key), `screen` (UI page; requires `wireframe`) |
| `interaction_pattern` | `--interaction-pattern` | `sync`, `async`, `request-response`, `pub-sub`, `streaming`, `schema` (with `storage`), `render` (with `screen`) |
| `wireframe` | `--wireframe` | UIML source describing the screen layout. Required when `contract_kind=screen`. Validator parses via `uiml.Parse`; dashboard renders via the dark-mode wireframe renderer (`uiml.RenderWireframeHTML`). Use the full UIML vocabulary including attributes: `<screen direction="horizontal">`, `<layout direction="horizontal">`, `<grid cols="N">`, `<sidebar width="200">`, `<panel width="360">`, `<item active="true">`. The classic inbox pattern is `<screen direction="horizontal"><sidebar/><panel/><main/></screen>`. Render from the terminal with `syde wireframe render <slug> [--format html\|ascii\|image] [--out path] [--open]`. |
| `protocol_notes` | `--protocol-notes` | Protocol details |
| `input` | `--input` | Invocation signature, **required** |
| `input_parameters` | `--input-parameter` | List of `{path, type, description}` entries (repeatable flag, spec `"path\|type\|description"`), **required ≥1** |
| `output` | `--output` | Output signature / response shape, **required** |
| `output_parameters` | `--output-parameter` | List of `{path, type, description}` entries, **required ≥1** |
| `constraints` | `--constraints-text` | Constraints |
| `versioning_notes` | `--versioning-notes` | Versioning notes |

### Concept

Concepts are first-class ERD entities — high-level domain objects
with named attributes, domain actions, and cardinality-labelled
`relates_to` relationships. Concepts are a **design-level lens** —
attributes carry name and description only, no concrete types. The
dashboard renders them in two places:

- **Concept detail panel** — accessed by clicking any concept in the
  Concepts list. Shows attribute and action tables alongside the
  usual entity fields.
- **ERD view** — accessed via the **List / ERD** toggle at the
  top-right of the Concepts page (not a separate sidebar entry).
  Renders every concept as a draggable React Flow card with
  attribute rows, cardinality-labelled aggregate edges, and
  FK-style per-attribute edges for attributes that carry `refs`.

| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `meaning` | `--meaning` | **Required** — one-line domain meaning |
| `attributes` | `--attribute` | **Required (≥1)** — repeatable, `name\|description\|refs` |
| `actions` | `--action` | Optional, repeatable, `name\|description` |
| `invariants` | `--invariants` | Recommended — rules that must always hold |
| `structure_notes` | `--structure-notes` | Free-form structure notes (prefer structured attributes) |
| `lifecycle` | `--lifecycle` | Lifecycle description (state machine, etc.) |
| `data_sensitivity` | `--data-sensitivity` | Sensitivity (e.g., PII, public) |

**Attribute spec** — pipe-separated, 1–3 parts:

- `name` (required) — the field / column / property name
- `description` — optional prose for what the attribute means, its
  invariants, constraints, or intent. Rendered as the secondary
  line in the ERD card and in the concept detail panel.
- `refs` — optional comma-separated list of concept slugs this
  attribute references (FK-style). The ERD view draws a dashed
  arrow from this attribute row directly to each referenced
  concept's card, labelled with the attribute name.

There is **no type field**. Concrete types belong in code or in
contract schemas, not in the design model. If you find yourself
writing `uuid` or `timestamptz` in an attribute, move that detail
into the description prose or the component that implements the
concept.

Examples:

- `id|primary key`
- `total|must be > 0`
- `status|lifecycle state: draft, placed, paid, shipped`
- `customer_id|foreign key|customer` (attribute-level FK)
- `tag_ids|tag references|tag,label` (multiple refs)

**Action spec** — pipe-separated, 1–2 parts:

- `name` (required) — the verb (e.g. `place`, `cancel`, `ship`)
- `description` — one-line explanation of what the action does

Examples: `place|transitions from draft to placed`, `cancel|reverts to draft if not yet shipped`.

**Cardinality labels on `relates_to`** — `--add-rel` accepts an
optional third part for relationship cardinality. Syntax:
`<target>:relates_to:<cardinality>`. The cardinality value MUST be
one of the four canonical values — anything else is rejected by
`syde sync check --strict`:

| Label          | Meaning                                          |
|----------------|--------------------------------------------------|
| `one-to-one`   | Each row on the left has exactly one on the right |
| `one-to-many`  | Each row on the left has many on the right       |
| `many-to-one`  | Many rows on the left share one on the right     |
| `many-to-many` | Many-to-many via an explicit join                |

Cardinality is optional — a two-part `--add-rel "x:relates_to"`
stays valid and renders an unlabeled edge in the ERD view.

**Validator rules (enforced by `syde sync check`):**

- `meaning` is required (ERROR)
- `attributes` must have at least one entry (ERROR)
- Each attribute needs a non-empty `name` (ERROR)
- `invariants` is recommended (WARN)
- `relates_to.label` must be in the cardinality enum when non-empty (ERROR)

### Flow
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `trigger` | `--trigger` | What starts this flow |
| `goal` | `--goal` | What this flow achieves |
| `narrative` | `--narrative` | Step-by-step narrative |
| `happy_path` | `--happy-path` | Happy path description |
| `edge_cases` | `--edge-cases` | Edge cases |
| `failure_modes` | `--failure-modes` | Failure modes |
| `performance_notes` | `--performance-notes` | Performance notes |

### Decision
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `category` | `--category` | Category (e.g., data, api, security) |
| `statement` | `--statement` | The decision itself |
| `rationale` | `--rationale` | Why |
| `alternatives_considered` | `--alternatives` | Alternatives |
| `tradeoffs` | `--tradeoffs` | Tradeoffs |
| `consequences` | `--consequences` | Consequences |

### Requirement
Requirements preserve user intent and approved plan intent as append-only
records. Do not delete requirements. If a later requirement conflicts with an
older one, create the newer requirement and mark the older requirement
`superseded` or `obsolete`.

| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `statement` | `--statement` | Required requirement text; can satisfy the base description rule |
| `source` | `--source` | `user`, `plan`, `migration`, or `manual` |
| `source_ref` | `--source-ref` | Stable pointer to the prompt, plan, issue, migration, etc. |
| `requirement_status` | `--requirement-status` | `active`, `superseded`, or `obsolete` |
| `rationale` | `--rationale` | Why this requirement exists |
| `acceptance_criteria` | `--acceptance` | How fulfillment is recognized |
| `supersedes` | `--supersedes` | Comma-separated requirement refs replaced by this one |
| `superseded_by` | `--superseded-by` | Comma-separated newer requirement refs |
| `obsolete_reason` | `--obsolete-reason` | Required when status is `obsolete` |
| `approved_at` | `--approved-at` | Approval/capture timestamp |

Validation rules:

- `statement`, `source`, and valid `requirement_status` are required.
- `superseded` requirements must have `superseded_by`.
- `obsolete` requirements must have `obsolete_reason`.
- Supersede links must point to requirement entities and be reciprocal.
- Requirements are append-only; `syde remove` refuses to delete them.

### Plan
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `plan_status` | (auto) | `draft`, `approved`, `in-progress`, `completed` |
| `background` | `--background` | Why the plan exists (context, motivation) |
| `objective` | `--objective` | What success looks like |
| `scope` | `--scope` | In-scope / out-of-scope summary |
| `phases` | (add-phase) | Ordered phase list |

Each `PlanPhase` has:
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `name` | `--name` | Short label |
| `parent_phase` | `--parent` | Parent phase ID for nesting |
| `description` | `--description` | One-line phase summary |
| `objective` | `--objective` | What this phase achieves |
| `changes` | `--changes` | What concretely changes |
| `details` | `--details` | Implementation walkthrough (HOW) |
| `notes` | `--notes` | Risks, reminders, context |
| `status` | (auto) | `pending`, `in_progress`, `completed`, `skipped` |

### Task
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `task_status` | (auto) | `pending`, `in_progress`, `completed`, `blocked`, `cancelled` |
| `priority` | `--priority` | `high`, `medium`, `low` |
| `objective` | `--objective` | What the task achieves |
| `details` | `--details` | Implementation specifics |
| `acceptance` | `--acceptance` | Observable done-condition |
| `plan_ref` | `--plan` | Parent plan slug |
| `plan_phase` | `--phase` | Parent phase ID |
| `entity_refs` | `--entity` | Linked entities |

### Learning
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `category` | `--category` | `gotcha`, `constraint`, `convention`, `context`, `dependency`, `performance`, `workaround` |
| `entity_refs` | `--entity` | Linked entity slugs |
| `confidence` | `--confidence` | `high`, `medium`, `low` |
| `source` | `--source` | Where this was learned |

## Relationships

Add with `--add-rel "target-slug:type"`. Valid types:

| Type | Meaning |
|------|---------|
| `belongs_to` | Is part of (system→system for sub-systems, component→system, contract→system, concept→system) |
| `depends_on` | Requires (component→component, must be acyclic) |
| `exposes` | Provides interface |
| `consumes` | Uses interface |
| `uses` | General usage |
| `involves` | Participates in (flow→component) |
| `references` | Points to (contract→concept, concept→component) |
| `relates_to` | ERD-style relationship (concept→concept) |
| `implements` | Implements (component→concept) |
| `applies_to` | Applies to target (decision→component) |
| `modifies` | Changes target |
| `visualizes` | UI for target |

Every entity except the single root system must have a `belongs_to` parent.
Every non-requirement entity must link to at least one requirement, usually
with `references` or `implements`. Every contract must participate in at least
one flow through either an outbound or inbound relationship.

## Creating Entities via Plans

Use `syde plan add-entity` to add draft entities to a plan. All kind-specific
flags are supported. Entities are created when the plan is executed.

```bash
syde plan add-entity <plan-slug> component "API Server" \
  --description "HTTP server for REST API" \
  --responsibility "Request routing, validation" \
  --boundaries "No direct DB access"
```

For entities created outside a plan (e.g., during sync), use `syde add` directly.
