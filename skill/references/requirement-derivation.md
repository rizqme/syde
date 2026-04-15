# Requirement Derivation Algorithm

This is the deterministic procedure for backfilling EARS requirements
from existing design entities. Subagents handling phase-4 backfill
batches MUST follow this algorithm exactly. The goal is dense,
narrow, verifiable requirements — not prose summaries.

## Inputs

You will be given a list of entity slugs to walk. For each entity,
read its full content with `syde get <slug> --full --format json`.
**Do NOT use grep or Read on the markdown file directly** — `syde
get` is the canonical accessor and gives you body + relationships
in one call.

## Output

For every requirement you create, run:

```
syde add requirement "<short Title Case name>" \
  --statement "<EARS-compliant shall sentence ending in a period>" \
  --type <functional|non-functional|constraint|interface|performance|security|usability> \
  --priority <must|should|could|wont> \
  --verification "<one short sentence describing how the property is verified>" \
  --rationale "<one short sentence on why this matters>" \
  --source manual \
  --source-ref "<source-entity-slug>" \
  --add-rel "<source-entity-slug>:refines" \
  --add-rel "syde-5tdt:belongs_to"
```

**Requirements never carry a `--file` flag.** They link only to
other entities (refines → component/contract/concept/system,
belongs_to → system, derives_from → parent requirement). A
requirement that lists files is a bug.

## The five EARS patterns

Every statement must match exactly one of these templates. The save
validator rejects anything that doesn't.

| Pattern | Template | Use when… |
|---|---|---|
| Ubiquitous | `The <subject> shall <action>.` | The property always holds. |
| Event-driven | `When <trigger>, the <subject> shall <action>.` | The property holds in response to a discrete event. |
| State-driven | `While <state>, the <subject> shall <action>.` | The property holds while the system is in a given state. |
| Unwanted-behavior | `If <unwanted condition>, then the <subject> shall <action>.` | The property describes how the system handles a failure or invalid input. |
| Optional-feature | `Where <feature is included>, the <subject> shall <action>.` | The property holds only when an optional feature/config is enabled. |

`<subject>` should be the system or component name being constrained
(e.g. "the syde CLI", "the syded daemon", "the storage layer"). Don't
use generic "the system" if a more specific subject is available.

## Per-kind derivation procedure

### Component

For a component, walk these fields in order:

1. **`responsibility`** — generate **one Ubiquitous requirement**.
   - `req_type`: `functional`
   - `priority`: `must`
   - `statement` template: `"The <component subject> shall <responsibility rephrased as verb phrase>."`
   - `verification`: if any file in `files[]` is a `_test.go` file or
     there's a known integration test command for the component, use
     `"automated test in <file>"`; otherwise `"inspection"`.

2. **`capabilities[]`** — one requirement per item.
   - `req_type`: `functional`
   - `priority`: `must`
   - `statement` template: `"The <subject> shall <capability rephrased>."`
   - Drop bullet points and rephrase as a single shall sentence.

3. **`boundaries`** — split on "does NOT" / "must not" / "never" /
   "shall not" sentences. One **constraint** requirement per phrase.
   - `req_type`: `constraint`
   - `priority`: `must`
   - `statement` template: `"The <subject> shall not <forbidden behavior>."`

4. **`failure_modes[]`** — one **Unwanted-behavior** requirement per item.
   - `req_type`: `functional`
   - `priority`: `must` (if mitigation is well-defined) or `should`
   - `statement` template: `"If <failure trigger>, then the <subject> shall <mitigation>."`

5. **`scaling_notes` / `performance_notes`** — split into individual
   quantitative claims, one **Non-functional** requirement each.
   - `req_type`: `performance`
   - `priority`: `should` unless the original text says "must" / "required"
   - `statement` template: `"The <subject> shall <action> within <quantitative bound>."`

6. **`data_handling`** — if it mentions security, encryption, secrets,
   or PII, one **security** requirement.

Target: 5–12 requirements per component.

### Contract

For a contract, walk these fields:

1. **`input` + `output`** — one **interface** requirement covering
   the call signature. Use Event-driven form.
   - `req_type`: `interface`
   - `priority`: `must`
   - `statement` template: `"When <input invocation>, the <contract owner> shall return <output shape>."`
   - For CLI contracts: `"When the user runs <command>, the syde CLI shall <produce output>."`
   - For HTTP contracts: `"When a client invokes <METHOD path>, the syded daemon shall respond with <status + body shape>."`
   - For screen contracts: `"When the user navigates to <route>, the dashboard shall render <visible elements>."`

2. **`input_parameters[]`** — one interface requirement per parameter
   with a non-trivial type or required flag.
   - `req_type`: `interface`
   - `priority`: `must`
   - `statement` template: `"When <input> is invoked, the <contract owner> shall accept <param name> as <type>."`
   - Skip purely cosmetic flags (e.g. `--format`, `--quiet`).

3. **`output_parameters[]`** — one interface requirement per parameter
   that's part of the success contract (skip optional/error fields).
   - `statement` template: `"When <input> succeeds, the <contract owner> shall return <param name> as <type>."`

4. **`constraints`** — one constraint requirement per phrase.
   - `req_type`: `constraint`
   - `priority`: `must`

5. **`protocol_notes`** — if it describes a protocol invariant
   (idempotency, ordering, timeouts), one constraint requirement.

Verification for contracts:
- CLI contract → `"integration test invoking <command>"`
- REST/RPC contract → `"integration test against /api/<path>"`
- Screen contract → `"manual inspection of <route>"`
- Event/WebSocket contract → `"end-to-end test publishing <event>"`

Target: 1–4 requirements per contract.

### System

For the system entity:

1. **`quality_goals`** — split on lines or sentences. One
   **non-functional** requirement per goal. Classify by keyword:
   - "fast", "latency", "ms", "throughput", "scale" → `performance`
   - "secure", "secret", "encrypt", "auth" → `security`
   - "easy", "discoverable", "ergonomic", "intuitive" → `usability`
   - "available", "reliable", "uptime" → `non-functional`
   - default → `non-functional`
   - `priority`: `must`
   - `statement` template: `"The <system> shall <goal restated as a property>."`

2. **`design_principles`** — one **constraint** requirement per
   principle.
   - `req_type`: `constraint`
   - `priority`: `must`
   - `statement` template: `"The <system> shall <principle restated>."`

3. **`assumptions`** — DO NOT generate requirements. Assumptions are
   not requirements; they are context for understanding the system's
   environment. Skip.

Target: 5–15 requirements per system.

### Concept

For a concept entity:

1. **`invariants`** — split on sentences. One **constraint**
   requirement per invariant.
   - `req_type`: `constraint`
   - `priority`: `must`
   - `statement` template: `"The <system> shall ensure <invariant>."` (Ubiquitous)
     or `"If <violation condition>, then the <system> shall <reject/repair>."` (Unwanted-behavior)

2. **`lifecycle`** — identify each named state transition. One
   **state-driven** requirement per transition.
   - `req_type`: `functional`
   - `priority`: `must`
   - `statement` template: `"While <concept> is in <state>, the <system> shall <allow/disallow action>."`

3. **`data_sensitivity`** — if non-trivial, one **security**
   requirement.
   - `req_type`: `security`
   - `priority`: `must`

4. **Required attributes** — from `attributes[]` filter to those
   marked required or implied required. One **constraint** requirement
   per attribute.
   - `statement` template: `"The <system> shall require <attribute name> on every <concept> instance."`

Target: 3–8 requirements per concept.

## Quality checks

Before submitting each `syde add requirement` call, verify:

1. **EARS pattern match.** The statement matches one of the five
   templates exactly. The save validator will enforce this — if it
   rejects you, the statement is wrong.
2. **Length.** Statement is 10–250 characters. Anything shorter is
   probably vague; anything longer is two requirements jammed
   together — split it.
3. **Uniqueness.** Don't create two requirements that overlap. If a
   capability and a constraint say the same thing from different
   angles, keep only the constraint.
4. **No task verbs.** Reject statements that start with "Add", "Create",
   "Implement", "Build", "Refactor", "Update", "Fix" applied to the
   *system itself*. Those are tasks. ("The system shall create X" is
   fine — that's the system creating something at runtime.)
5. **Back-link.** Every generated requirement carries
   `--add-rel <source-slug>:refines` so the coverage audit
   can find it.
6. **No `--file`.** Never carry a files list on a requirement.
7. **Subject is concrete.** Replace "the system" with the actual
   subsystem name when the requirement applies to a specific
   component or contract owner.
8. **Verifiable.** A reviewer should be able to look at the
   `verification` field and know exactly what to run or inspect to
   prove the property holds. "Manual" with no detail is a smell.

## What NOT to derive

- **Plan/task entities.** Never generate requirements about plans
  or tasks. Those are work artefacts, not system properties.
- **Implementation details.** "The Go code shall use a sync.Mutex"
  is wrong — that's an implementation choice, not a property.
  Express the property: "If two writers race, the storage layer
  shall serialize them and not lose data."
- **Aspirations without a verification path.** If you can't write a
  verification field, the requirement is too vague.
- **Catch-all requirements.** Each requirement should be linked by
  ≤10 entities of the same kind (the fanout cap audit enforces this).
  If you find yourself writing "The system shall do everything good",
  split it.

## Worked examples

**Component: Audit Engine**
- Responsibility: "Validate the syde model on demand and during sync"
- Capabilities: ["Run entity field validation", "Detect requirement traceability gaps", "Cap requirement fanout per kind"]
- Boundaries: "Does NOT mutate entities. Does NOT call LLMs."
- Failure modes: ["Cycle in plan phase graph", "Missing relationship target"]

Generated:
1. `"The audit engine shall validate every entity in the syde model when invoked by sync check or validate."` — functional/must
2. `"The audit engine shall detect entities that lack outgoing requirement traceability."` — functional/must
3. `"The audit engine shall reject requirements linked by more than ten entities of the same source kind."` — constraint/must
4. `"The audit engine shall not mutate any entity it inspects."` — constraint/must
5. `"The audit engine shall not invoke any LLM API."` — constraint/must
6. `"If a plan phase graph contains a cycle, then the audit engine shall report a cycle finding and skip that plan."` — functional/must
7. `"If a relationship target slug does not exist, then the audit engine shall report a missing-target warning."` — functional/should

**Contract: syde add requirement (CLI)**
- Input: `syde add requirement <name> [flags]`
- Input parameters: `--statement (string, required)`, `--type (enum, required)`, `--priority (enum, required)`
- Output: created requirement file path

Generated:
1. `"When the user runs syde add requirement <name>, the syde CLI shall create a new requirement entity file under .syde/requirements/."` — interface/must
2. `"When syde add requirement is invoked, the syde CLI shall require a non-empty --statement that matches an EARS pattern."` — interface/must
3. `"When syde add requirement is invoked, the syde CLI shall require --type to be one of functional, non-functional, constraint, interface, performance, security, or usability."` — interface/must
4. `"If --statement does not match any EARS pattern, then the syde CLI shall reject the command and emit a validation error naming the five allowed patterns."` — functional/must

## When to ask for help

If an entity's responsibility or capabilities are too vague to
derive concrete requirements (e.g. a placeholder component with one
sentence of prose), STOP. Don't invent properties. Report the
entity slug back to the orchestrator so a human can flesh it out
first.
