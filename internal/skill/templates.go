package skill

// SkillMD is the SKILL.md content for the Claude Code skill.
var SkillMD = `---
name: syde
description: >
  Always active when installed. Manages the project's software design model
  in .syde/. Enforces architectural constraints, tracks plans and tasks,
  captures learnings. Triggers on any code modification, design discussion,
  plan creation, or architecture query.
tools: Read, Write, Edit, Bash, Glob, Grep
---

# syde — Software Design Model

This project has a syde design model. Architecture context is auto-loaded at
session start via the SessionStart hook. You already have the full entity map,
decisions, and learnings in your context.

## Workflow (MANDATORY — DO NOT SKIP ANY STEP)

### Step 1: CLARIFY BEFORE DOING ANYTHING
- Before ANY implementation, list ALL clarifying questions
- For each question, provide your recommended answer
- Present to the user and WAIT for confirmation
- Do NOT proceed without explicit user approval
- Example: "Before I start, I have these questions:
  1. Should the API use JSON or protobuf? (Recommended: JSON)
  2. Do we need authentication? (Recommended: No, keep it simple)
  Please confirm or adjust."

### Step 2: CREATE AND PRESENT PLAN
1. Create design entities first (use ` + "`syde batch`" + ` for >3 entities):
   ` + "```" + `
   # Prepare batch.yaml with all entities, then:
   syde batch batch.yaml
   ` + "```" + `
   Or individually: ` + "`syde add <kind> <name> --description \"...\" [kind-specific flags]`" + `
2. Create plan: ` + "`syde plan create \"<name>\"`" + `
3. Add steps: ` + "`syde plan add-step <slug> --description \"...\" --action create --entity-kind <kind> --entity-name \"...\"`" + `
4. Check size: ` + "`syde plan estimate <slug>`" + `
   - If >10 steps: split into sub-plans (one per session turn)
5. Present: ` + "`syde plan show <slug>`" + `
6. Tell user: "Here is the plan. Please review and approve with ` + "`syde plan approve <slug>`" + `"
7. **STOP. Do NOT implement until plan is approved.**

### Step 3: IMPLEMENT (one sub-plan per turn)
For each task:
1. ` + "`syde task start <slug>`" + `
2. Use ` + "`syde query <entity> --full`" + ` only if you need deep detail about a specific entity
3. Write the code
4. ` + "`syde constraints check <file>`" + ` for each new file — verify it maps to a component
5. ` + "`syde task done <slug>`" + `
6. ` + "`syde plan step <plan> <step-id> --status completed`" + `
7. If you discovered anything undocumented: ` + "`syde remember \"<text>\" --category <type> --entity <name>`" + `

### Step 4: FINISH
- ` + "`syde validate`" + ` — check model integrity
- ` + "`syde status`" + ` — confirm counts
- ` + "`syde plan show <slug>`" + ` — confirm 100% complete
- You MUST complete ALL tasks in the current sub-plan. Never leave tasks in_progress.

## Commands Quick Reference
- ` + "`syde context`" + ` — full architecture snapshot (auto-loaded at session start)
- ` + "`syde query <slug> --full`" + ` — deep dive on one entity (relationships, learnings, tasks)
- ` + "`syde add <kind> <name> [flags]`" + ` — create entity
- ` + "`syde update <slug> [flags]`" + ` — update entity (supports all kind-specific fields)
- ` + "`syde batch <file.yaml>`" + ` — bulk create entities
- ` + "`syde plan add-step <plan> --description \"...\"`" + ` — add structured step
- ` + "`syde plan estimate <plan>`" + ` — check if plan needs splitting
- ` + "`syde task start/done <slug>`" + ` — track implementation progress
- ` + "`syde constraints check <file>`" + ` — verify source file is mapped to component
- ` + "`syde remember \"<text>\" --category <type>`" + ` — capture learning
- ` + "`syde validate`" + ` — check model integrity

## Key Rules
1. NEVER read .syde/ files directly — always use syde CLI commands
2. NEVER skip the clarify → plan → approve → implement sequence
3. ALWAYS verify new source files are mapped to components after writing
4. ALWAYS capture learnings when you discover undocumented behavior
5. Use ` + "`syde batch`" + ` when creating >3 entities to save turns

## Reference Files
- ` + "`references/entity-spec.md`" + ` — field specification for all entity kinds
- ` + "`references/plan-workflow.md`" + ` — detailed plan workflow
- ` + "`references/constraints.md`" + ` — constraint enforcement protocol
- ` + "`references/scan-workflow.md`" + ` — bootstrapping from existing code
- ` + "`references/commands.md`" + ` — full CLI reference
`

// EntitySpecRef is the entity specification reference content.
var EntitySpecRef = `# Entity Specification

## BaseEntity (all types)
- id: stable unique identifier (auto-generated)
- kind: entity type
- name: human-readable name
- description: short explanation
- purpose: why it exists
- status: draft, active, deprecated, proposed, superseded
- tags: labels
- notes: informal notes
- relationships: typed links to other entities

## System
- context, scope, design_principles, quality_goals, assumptions
- Flags: --context-text, --scope, --design-principles, --quality-goals, --assumptions

## Component
- responsibility, boundaries, behavior_summary, interaction_summary
- data_handling, scaling_notes, failure_modes
- Flags: --responsibility, --boundaries, --behavior-summary, --interaction-summary, --data-handling, --scaling-notes

## Contract
- contract_kind (api/event/command/query), interaction_pattern (sync/async/request-response/pub-sub)
- protocol_notes, input_description, output_description, constraints, versioning_notes
- Flags: --contract-kind, --interaction-pattern, --protocol-notes, --input-desc, --output-desc, --constraints-text, --versioning-notes

## Concept
- meaning, structure_notes, lifecycle, invariants, data_sensitivity
- Flags: --meaning, --structure-notes, --lifecycle, --invariants, --data-sensitivity

## Flow
- trigger, goal, narrative, happy_path, edge_cases, failure_modes
- observability_notes, performance_notes
- Flags: --trigger, --goal, --narrative, --happy-path, --edge-cases, --failure-modes, --performance-notes

## Decision
- category, statement, rationale, alternatives_considered, tradeoffs, consequences
- Flags: --category, --statement, --rationale, --alternatives, --tradeoffs, --consequences

## Learning
- category (gotcha/constraint/convention/context/dependency/performance/workaround)
- entity_refs, source, confidence (high/medium/low)

## All entities also support
- --body "markdown text" — set the detail body
- --add-rel target:type — add a relationship
- --remove-rel target — remove a relationship

## Batch creation (batch.yaml)
` + "```yaml" + `
entities:
  - kind: component
    name: API Server
    description: HTTP server
    responsibility: Request routing
  - kind: contract
    name: Todo API
    description: REST CRUD
    contract_kind: api
` + "```" + `
Run: syde batch batch.yaml
`

// CommandsRef is the CLI command reference.
var CommandsRef = `# syde CLI Commands

## Core
` + "```" + `
syde context [--json]              # Full architecture snapshot
syde init [--install-skill]        # Initialize .syde/
syde add <kind> <name> [flags]     # Create entity (all kind-specific flags supported)
syde get <id-or-slug>              # Show entity
syde list [kind] [--status X]      # List entities
syde update <slug> [flags]         # Update entity (all kind-specific flags supported)
syde remove <slug> [--force]       # Delete entity
syde batch <file.yaml>             # Bulk create entities
syde search <query>                # Full-text search
syde status                        # Model overview
syde validate                      # Check integrity
syde reindex                       # Rebuild index
syde constraints [--json]          # Active decisions + learnings
syde constraints check <file>      # Map file → component → constraints
` + "```" + `

## Query
` + "```" + `
syde query <slug> [--full]         # Rich entity lookup
syde query --kind X --tag Y        # Filter entities
syde query --impacts <slug>        # Transitive impact analysis
syde query --flow <slug> --components  # Flow decomposition
syde query --related-to <slug>     # Direct connections
syde query --search "text"         # Full-text search
syde query --diff <slug> --since 7d    # Git change history
syde query --format json|compact|refs  # Output format
` + "```" + `

## Plans
` + "```" + `
syde plan create <name>            # Create plan
syde plan add-step <slug> --description "..." [--action X --entity-kind Y --entity-name Z]
syde plan estimate <slug>          # Size check + split recommendation
syde plan show <slug>              # Show with step progress
syde plan approve <slug>           # Mark approved
syde plan step <slug> <step-id> --status completed
syde plan list                     # List all plans
` + "```" + `

## Tasks
` + "```" + `
syde task create <name> [--plan P --entity E --priority high|medium|low]
syde task start/done/block <slug>
syde task list [--status X]
syde task sub <parent> <name>      # Create subtask
syde task link <task> <entity>     # Link to entity
` + "```" + `

## Learnings + Memory
` + "```" + `
syde remember "<text>" --category C [--entity E --confidence high|medium|low]
syde learn list/about/search/promote/stale
syde memory sync/list/clean
` + "```" + `

## Other
` + "```" + `
syde graph [entity] [--format dot]
syde design create/show/preview/export/validate/link
syde scan [--dry-run --coverage]
syde open / syde server start/stop
syde install-skill
` + "```" + `
`

// PlanWorkflowRef is the plan workflow reference.
var PlanWorkflowRef = `# Plan Workflow (MANDATORY)

## Before Starting: Clarify Requirements
1. List ALL questions about the task with your recommended answer for each
2. Present to the user: "Before I begin, here are my questions: [list]"
3. WAIT for the user to confirm or adjust
4. Do NOT proceed without explicit confirmation

## Step 1: Design Entities
- If creating >3 entities: prepare batch.yaml and run ` + "`syde batch batch.yaml`" + `
- Otherwise: ` + "`syde add <kind> <name> --description ... [kind-specific flags]`" + `
- Run ` + "`syde validate`" + ` to check integrity

## Step 2: Create Plan with Structured Steps
1. ` + "`syde plan create \"<plan name>\"`" + `
2. For each implementation step:
   ` + "`syde plan add-step <slug> --description \"...\" --action create --entity-kind component --entity-name \"...\"`" + `
3. ` + "`syde plan estimate <slug>`" + ` — check if plan needs splitting
4. If estimate says >10 steps or >25 estimated commands: create sub-plans

## Step 3: Present Plan and Wait for Approval
1. ` + "`syde plan show <slug>`" + `
2. Tell the user: "Plan ready. Run ` + "`syde plan approve <slug>`" + ` to proceed."
3. **STOP HERE.** Do not implement until the user approves.

## Step 4: Implement (One Sub-Plan Per Turn)
For each task in the plan:
1. ` + "`syde task start <task-slug>`" + `
2. Write code (use ` + "`syde query <entity> --full`" + ` for targeted lookups)
3. After each new file: ` + "`syde constraints check <file>`" + ` — verify mapping
4. ` + "`syde task done <task-slug>`" + `
5. ` + "`syde plan step <plan> <step-id> --status completed`" + `
6. If you learned something: ` + "`syde remember \"...\" --category <type> --entity <name>`" + `

## Step 5: Verify Completion
- ` + "`syde validate`" + ` — 0 errors
- ` + "`syde plan show <slug>`" + ` — 100% complete
- ` + "`syde status`" + ` — entity counts match expectations
- All tasks must be completed. Never leave tasks in in_progress state.
`

// ConstraintsRef is the constraint enforcement protocol.
var ConstraintsRef = `# Constraint Enforcement

## Architecture is auto-loaded
The SessionStart hook injects ` + "`syde context --json`" + ` at the beginning of every session.
You already have all entities, decisions, and learnings in your context.
Do NOT re-run syde context/constraints/status manually — it's already loaded.

## When you need deep detail
Use ` + "`syde query <slug> --full`" + ` for a specific entity. This returns:
- All structured fields
- Resolved relationships (with target names and files)
- Related learnings
- Linked tasks
- Applicable decisions

## After writing source files
Run ` + "`syde constraints check <file>`" + ` to verify the file maps to a component.
If it doesn't map, use ` + "`syde update <component> --add-rel <file-ref>:<type>`" + `.

## If your change violates a constraint
1. Flag it to the user with the specific decision/learning that's violated
2. Suggest the correct approach
3. Only proceed if the user explicitly approves the violation
`

// ScanWorkflowRef is the scan workflow reference.
var ScanWorkflowRef = `# Scan Workflow (Bootstrap from Existing Code)

Run ` + "`syde scan`" + ` to generate a scan guide, then follow 5 agent rounds.

## Round 1: System + Components (parallel)
One agent per directory in scan-guide.json. Each reads ALL files and creates a component.

## Round 2: Contracts + Concepts (parallel)
Find API routes, proto files, event schemas. Create contracts. Find domain models. Create concepts.

## Round 3: Flows (parallel)
Trace execution from entry points across components. Write step-by-step narratives.

## Round 4: Decisions (single)
Read all entities + architecture docs. Extract implicit decisions.

## Round 5: Completeness (single)
Wire relationships. Validate. Fill gaps. Capture learnings.
Run ` + "`syde scan --coverage`" + ` to verify all directories are mapped.
`
