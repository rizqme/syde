# Entity Specification

## Entity Kinds

| Kind | What it represents |
|------|-------------------|
| `system` | The overall product/service |
| `component` | A service, module, or major part |
| `contract` | An API, event, or protocol boundary |
| `concept` | A domain object or business entity |
| `flow` | An end-to-end workflow or user journey |
| `requirement` | Append-only user/plan/migration requirement (architectural intent in EARS shall-form) |
| `plan` | A tracked implementation plan |
| `task` | A work item linked to plans/entities |

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

Concepts are **domain glossary entries** — short prose definitions of
the terms that matter in the domain. They are a design-level lens and
carry no typed attributes, no actions, and no cardinality-labelled
relationships. Schema detail and implementation types belong in code
or in contract schemas, never on a concept. The dashboard renders
concepts as a standard 2-column inbox: a list on the left, a glossary
detail panel on the right showing the prose fields plus relationship
chips grouped by role.

| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `meaning` | `--meaning` | **Required** — what the term represents in the domain |
| `invariants` | `--invariants` | Recommended — rules that must always hold |
| `lifecycle` | `--lifecycle` | Optional — state machine or progression of stages |

There are no `attributes`, `actions`, `data_sensitivity`, or
`structure_notes` fields anymore. A concept that was authored under
the earlier schema can be migrated with a no-op `syde update <slug>`:
the re-save drops the obsolete keys from YAML because the struct no
longer carries them.

**Role-based relationships** — pick the one that matches how the
concept is realised in the rest of the model:

| Relationship type | Target kind | Meaning |
|-------------------|-------------|---------|
| `implemented_by`  | component   | Component that implements this domain concept in code |
| `exposed_via`     | contract    | Contract that exposes this concept at a process boundary |
| `used_in`         | flow        | Flow that operates on or produces this concept |
| `relates_to`      | concept     | Another concept this one relates to (no cardinality) |

Examples:

- `--add-rel "order-service:implemented_by"` — the Order concept is
  realised by the Order Service component.
- `--add-rel "place-order:exposed_via"` — the Order concept appears
  in the Place Order contract's schema.
- `--add-rel "place-order-flow:used_in"` — the Order concept is
  produced by the Place Order flow.
- `--add-rel "line-item:relates_to"` — Order and LineItem are
  related domain concepts. No cardinality label.

Use the generic `references` relationship only when none of the
role-based types fit (rare).

**Validator rules (enforced by `syde sync check`):**

- `meaning` is required (ERROR)
- `invariants` is recommended (Finding when empty)
- `relates_to` carries no cardinality label (audit does not validate one)

### Flow
Each flow represents **one user goal** (e.g. "Create Plan", "Browse Components").
Steps describe what user and system do at each point in the journey.

| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `trigger` | `--trigger` | What starts this flow |
| `goal` | `--goal` | What this flow achieves |
| `steps` | `--step` | **Structured steps** (repeatable). Format: `id\|action\|contract\|description\|on_success\|on_failure`. Min 2 fields (id + action). |
| `narrative` | `--narrative` | Prose narrative (legacy — prefer steps) |
| `happy_path` | `--happy-path` | Happy path description |
| `edge_cases` | `--edge-cases` | Edge cases |
| `failure_modes` | `--failure-modes` | Failure modes |
| `performance_notes` | `--performance-notes` | Performance notes |

#### FlowStep fields
| Field | Description |
|-------|-------------|
| `id` | Short intra-flow identifier (e.g. `s1`, `s2`). Must be unique within the flow. |
| `action` | What happens — verb phrase describing what user or system does. |
| `contract` | Slug of the contract this step exercises. Empty for internal steps (legitimate; the audit only fires on unresolvable refs, not on intentionally empty ones). |
| `description` | Optional elaboration. |
| `on_success` | Step ID to follow on success, `done`, or empty (= implicit next in array order). |
| `on_failure` | Step ID to follow on failure, `done`, `abort`, or empty (= no failure path). |

Audit rules:
- ERROR: contract not referenced by any flow step across all flows
- ERROR: step `contract` ref doesn't resolve to an existing contract
- ERROR: step `on_success`/`on_failure` ref doesn't resolve within the same flow
- ERROR: duplicate step ID within a flow
- (internal steps with an empty `contract` field are legitimate and silent — the audit only fires on unresolvable contract refs)

### Requirement
Requirements preserve user intent and approved plan intent as append-only
records in **EARS shall-form**. Do not delete requirements. If a later
requirement conflicts with an older one, create the newer requirement and
mark the older requirement `superseded` or `obsolete`.

| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `statement` | `--statement` | Required EARS shall-form text; save validator rejects non-conforming statements. Can satisfy the base description rule. |
| `req_type` | `--type` | **Required** — `functional`, `non-functional`, `constraint`, `interface`, `performance`, `security`, `usability` |
| `priority` | `--priority` | **Required** — MoSCoW: `must`, `should`, `could`, `wont` |
| `verification` | `--verification` | **Required** — short free-form description of how the requirement is verified (test, review, demo, metric, etc.) |
| `source` | `--source` | `user`, `plan`, `migration`, or `manual` |
| `source_ref` | `--source-ref` | Stable pointer to the prompt, plan, issue, migration, etc. |
| `requirement_status` | `--requirement-status` | `active`, `superseded`, or `obsolete` |
| `rationale` | `--rationale` | Why this requirement exists |
| `supersedes` | `--supersedes` | Comma-separated requirement refs replaced by this one |
| `superseded_by` | `--superseded-by` | Comma-separated newer requirement refs |
| `obsolete_reason` | `--obsolete-reason` | Required when status is `obsolete` |
| `approved_at` | `--approved-at` | Approval/capture timestamp |
| `audited_overlaps` | `--audited` | List of `{slug, distinction}` acknowledging TF-IDF overlap pairs. Pass as `slug:distinction text` (distinction ≥20 chars, must describe real semantic difference). |

**EARS patterns** — every `statement` must match exactly one of these:

| Pattern | Template |
|---------|----------|
| Ubiquitous | `The <subject> shall <action>.` |
| Event-driven | `When <trigger>, the <subject> shall <action>.` |
| State-driven | `While <state>, the <subject> shall <action>.` |
| Unwanted-behavior | `If <unwanted condition>, then the <subject> shall <action>.` |
| Optional-feature | `Where <feature is included>, the <subject> shall <action>.` |

The legacy acceptance field has been **removed** from requirements. Use
`verification` (a short description of how fulfillment is checked) instead.
`--acceptance` is no longer a flag on `syde add requirement` (it still exists
on `syde task`).

**Requirements never carry a `files` list.** They are pure design intent and
link only via the relationships below. Do not pass `--file` on a requirement.

**Requirement-specific relationships:**

- `refines` — requirement → component / contract / concept / system, or
  requirement → requirement. Used when a requirement narrows higher-level
  intent against a specific design target.
- `derives_from` — requirement → parent requirement. Used for derivation
  chains where a child requirement is logically implied by a parent.

Validation rules:

- `statement`, `source`, valid `requirement_status`, `req_type`, `priority`,
  and `verification` are required.
- `statement` must match one of the five EARS patterns (enforced on save).
- Requirements must not have a `files` list.
- `superseded` requirements must have `superseded_by`.
- `obsolete` requirements must have `obsolete_reason`.
- Supersede links must point to requirement entities and be reciprocal.
- Requirements are append-only; `syde remove` refuses to delete them.
- Any two active requirements with TF-IDF similarity ≥0.6 must acknowledge
  each other via `audited_overlaps`, and each acknowledgement must carry a
  `distinction` rationale of at least 20 characters. `syde add requirement`
  blocks non-zero on unacknowledged overlaps unless `--force` is passed.
- Requirement statements that name a CLI invocation (`syde <sub>`), REST
  path, dashboard screen, or pub-sub event topic must have a matching
  active contract whose `input` covers the surface. Planning-time and
  post-plan audits are symmetric.

**Backfilling from an existing codebase.** See
`references/requirement-derivation.md` for the deterministic algorithm
subagents use to derive EARS requirements from existing components,
contracts, concepts, and systems (with stable slugs and the correct
`refines` / `derives_from` chains).

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
| `references` | Points to (generic; prefer role-based types where applicable) |
| `relates_to` | Plain relationship between two concepts (no cardinality label) |
| `implements` | Implements (component→concept) |
| `implemented_by` | Concept → component that implements it in code |
| `exposed_via` | Concept → contract that exposes it externally |
| `used_in` | Concept → flow that operates on or produces it |
| `applies_to` | Applies to target (requirement→component) |
| `refines` | Requirement narrows a design target (requirement→component/contract/concept/system, or requirement→requirement) |
| `derives_from` | Requirement is derived from a parent requirement (requirement→requirement) |
| `modifies` | Changes target |
| `visualizes` | UI for target |

Every entity except the single root system must have a `belongs_to` parent.
Every non-requirement entity, including every component and contract, must
carry an outbound relationship to at least one requirement, usually with
`references` or `implements`. Every contract must participate in at least one
flow through either an outbound or inbound relationship.

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
