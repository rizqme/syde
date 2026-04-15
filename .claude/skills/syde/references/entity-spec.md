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
| Description | `description` | Short one-sentence identification of what the entity is. **REQUIRED on every kind** — `syde validate` errors when missing. Keep distinct from `purpose` (the why) and `responsibility` (the what does). |
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
| `contract_kind` | `--contract-kind` | `rest`, `cli`, `event`, `rpc`, `graphql`, `websocket`, `pubsub` |
| `interaction_pattern` | `--interaction-pattern` | `sync`, `async`, `request-response`, `pub-sub` |
| `protocol_notes` | `--protocol-notes` | Protocol details |
| `input` | `--input` | Invocation signature, **required** |
| `input_parameters` | `--input-parameter` | List of `{path, type, description}` entries (repeatable flag, spec `"path\|type\|description"`), **required ≥1** |
| `output` | `--output` | Output signature / response shape, **required** |
| `output_parameters` | `--output-parameter` | List of `{path, type, description}` entries, **required ≥1** |
| `constraints` | `--constraints-text` | Constraints |
| `versioning_notes` | `--versioning-notes` | Versioning notes |

### Concept
| YAML key | CLI flag | Description |
|----------|----------|-------------|
| `meaning` | `--meaning` | Domain meaning |
| `structure_notes` | `--structure-notes` | Structure notes |
| `lifecycle` | `--lifecycle` | Lifecycle description |
| `invariants` | `--invariants` | Rules that must always hold |
| `data_sensitivity` | `--data-sensitivity` | Sensitivity (e.g., PII, public) |

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
