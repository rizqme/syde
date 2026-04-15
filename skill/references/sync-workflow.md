# Sync Workflow

`syde sync` keeps the design model aligned with the codebase. Use it in two scenarios:

1. **New project** — bootstrap a design model from existing code
2. **Existing model** — verify entities match the implementation, detect drift

Always run `syde sync` after `syde init` on a project that already has code.

## Prerequisite: the summary tree must be clean

`syde sync` **automatically runs `syde tree scan` as its first step**, so
you don't need to remember to do it separately. But syde only *scans*
the tree — it does not *summarize* stale nodes for you. If the scan
reports stale paths, you must run the leaves-first summarize loop
before entity extraction:

```
syde sync                         # runs tree scan then sync guide
# if sync reports "N paths need summaries":
syde tree changes --leaves-only
syde tree summarize <path> --summary "..."
# repeat until:
syde tree status --strict         # must exit 0
```

Every round below pulls file content through `syde tree context
<path>` — **never naive `Read`**. The tree context bundles the ancestor
breadcrumb (which system/folder the file lives in), the file's stored
summary, and the raw content in one call. That's the right framing for
setting `--purpose`, `--responsibility`, `--boundaries`, and picking
the correct `belongs_to` target. Cheaper on tokens than chained `Read`s
and avoids mis-assigning entities to the wrong sub-system.

> **Rule**: if you're turning a source file into a syde entity, pull
> it through `syde tree context <file>` first. Only fall back to
> `Read` for verification of specific details.

---

## Scenario 1: New Project (no entities yet)

Run `syde sync` to generate a sync guide, then follow 5 agent rounds.

### Round 1: System + Components (parallel agents)

One agent per top-level directory in `.syde/scan-guide.json`.

Each agent:
1. Calls `syde tree context <directory>` — returns the folder summary
   (written in Phase 0), a listing of its children with their
   summaries, and the architectural breadcrumb. That's enough to name
   the component and write its description WITHOUT reading any source
   file cold.
2. For each child file that's relevant, calls `syde tree context <file>`
   to get content + breadcrumb in one shot. Identifies the module's
   responsibility, boundaries, data handling from the bundled summary +
   content.
3. Creates a component:
   ```bash
   syde add component "<name>" \
     --description "..." \
     --responsibility "..." \
     --boundaries "..."
   ```

Also create the system entity:
```bash
syde add system "<project name>" \
  --description "..." \
  --scope "..." \
  --design-principles "..."
```

### Round 2: Contracts + Concepts (parallel agents)

Find API boundaries and domain models:
- **Contracts**: API route files, proto files, event schemas, OpenAPI specs, GraphQL schemas
- **Concepts**: Domain models, database schemas, shared types

```bash
# Contract name must be a descriptive noun phrase (e.g. "User Login"),
# NOT the raw invocation. The invocation goes in --input.
syde add contract "<descriptive name>" \
  --contract-kind rest \
  --interaction-pattern request-response \
  --input "POST /api/resource" \
  --input-parameter "field|type|description" \
  --input-parameter "other.path|type|description" \
  --output "200 OK application/json" \
  --output-parameter "data.id|string|resource id" \
  --output-parameter "data.name|string|resource name" \
  --add-rel "<system-slug>:belongs_to"

syde add concept "<Model name>" \
  --meaning "..." \
  --invariants "..." \
  --lifecycle "..." \
  --add-rel "<system-slug>:belongs_to"
```

All four contract fields (`--input`, `--input-parameter`, `--output`,
`--output-parameter`) are REQUIRED — validator errors on missing ones.

Create each entity individually with `syde add`.

### Round 3: Flows (parallel agents)

Trace execution from entry points across components:
- HTTP request handlers end-to-end
- Background job execution paths
- User-facing workflows

```bash
syde add flow "<flow name>" \
  --trigger "..." \
  --goal "..." \
  --narrative "Step 1: ... Step 2: ..."
```

### Round 4: Decisions (single agent)

Read all entities created so far plus any architecture docs (ADRs, README, design docs).
Extract implicit architectural decisions:

```bash
syde add decision "<decision>" \
  --statement "..." \
  --rationale "..." \
  --category "..."
```

### Round 5: Completeness (single agent)

1. Wire relationships between all entities (**one `--add-rel` per command call**):
   ```bash
   syde update <component> --add-rel "<other>:depends_on"
   syde update <component> --add-rel "<contract>:exposes"
   syde update <contract> --add-rel "<component>:belongs_to"
   syde update <flow> --add-rel "<component>:involves"
   syde update <decision> --add-rel "<component>:applies_to"
   syde update <concept> --add-rel "<component>:references"
   ```
   **Important**: `--add-rel` only accepts one relationship per call. Run separate
   `syde update` commands for each relationship.

2. Run completeness audit: `syde sync --check`
   This checks:
   - System entity exists
   - All components have file references
   - All components expose at least one contract
   - All entities have relationships
   - All entity kinds are represented
   - All descriptions are substantial (>20 chars)
   - All components have responsibility and boundaries

3. Check directory coverage: `syde sync --coverage`

4. Fix all gaps reported by `--check` and `--coverage`.

5. Validate: `syde validate` — must show 0 errors.

6. Capture learnings:
   ```bash
   syde remember "..." --category gotcha --entity <name>
   ```

---

## Scenario 2: Existing Model (verify and update)

When the project already has syde entities, sync verifies them against the code.

### Step 1: Generate sync guide

```bash
syde sync
```

This compares the current codebase structure against existing entities.

### Step 2: Check coverage

```bash
syde sync --coverage
```

Review the report:
- `✓` = directory is mapped to a component
- `✗` = directory has no matching component (may need one)
- `(infrastructure)` = non-source directory (migrations, deploy, etc.)

### Step 3: Detect drift (agent-driven)

For each component in the model, read the actual source files and verify:

1. **Description still accurate?** — Does the component still do what the description says?
   If not: `syde update <slug> --description "..."`

2. **Responsibility changed?** — Has the component taken on new responsibilities or shed old ones?
   If so: `syde update <slug> --responsibility "..."`

3. **Boundaries violated?** — Is the component doing things its `boundaries` says it shouldn't?
   If so: either fix the code or update the boundary.

4. **New dependencies?** — Does the component import modules it didn't before?
   If so: `syde update <slug> --add-rel "<other>:depends_on"`

5. **Contracts match implementation?** — Do API contracts (input/output descriptions) match the actual endpoints?
   If not: `syde update <contract-slug> --input "..." --input-parameter "path|type|desc" --output "..." --output-parameter "path|type|desc"`

6. **Flows still valid?** — Do flow narratives match the actual execution paths?
   If not: `syde update <flow-slug> --narrative "..."`

7. **Missing entities?** — Are there new modules, APIs, or domain models that have no entity?
   If so: create them with `syde add`.

8. **Stale entities?** — Are there entities for code that was removed?
   If so: `syde remove <slug>`

### Step 4: Completeness audit

```bash
syde sync --check
```

This verifies the model is comprehensive:
- Every component has file references and exposes contracts
- Every entity has relationships to other entities
- All entity kinds are represented (system, component, contract, concept, flow, decision)
- Descriptions are substantial, not placeholder text
- Components have responsibility and boundaries defined

Fix all reported gaps before proceeding.

### Step 5: Check learnings

```bash
syde learn stale
```

This shows learnings that reference entities whose files have changed since the
learning was captured. Review each one — it may be outdated.

### Step 5: Validate

```bash
syde validate
```

Must show 0 errors.

### Step 6: Confirm context accuracy

```bash
syde context
```

Read the full output. It should accurately describe the system as it exists now.
If anything is wrong or missing, fix it.
