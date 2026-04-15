---
approved_at: "2026-04-14T08:56:37Z"
background: 'A stale plan (cli-syded-client-refactor) was found with 5 phases whose id/name/status were all empty strings. syde plan show --full crashed with a 1GB goroutine stack dump because PlanEntity.CollectTasks and ChildPhases have no cycle guard: when every phase has ParentPhase=='''' and ID=='''', CollectTasks('''') finds the first phase (ID=''''), then ChildPhases('''') matches every phase (they all share ParentPhase==''''), recursing into itself forever. The root cause is a combo: (a) something wrote empty-ID phases to disk — likely an older add-phase code path or a bulk import — and (b) the traversal has zero defense against malformed data. syde sync check never flagged it either, so the rot sat there silently.'
completed_at: "2026-04-14T08:59:42Z"
created_at: "2026-04-14T08:54:51Z"
id: PLN-0005
kind: plan
name: guard-plan-phase-integrity
objective: Malformed plan phase data can never crash the renderer and cannot be written via syde plan add-phase. syde sync check --strict surfaces any pre-existing corruption as an ERROR so abandoned/broken plans are caught at session end instead of silently lingering.
phases:
    - changes: internal/model/plan.go ChildPhases skips phases with empty ID and self-referential ParentPhase; CollectTasks threads a visited map[string]bool so any phase ID is traversed at most once
      description: Visited-set guard in PlanEntity.CollectTasks + empty-ID and self-parent filtering in ChildPhases
      details: 'ChildPhases: skip phase if ph.ID == '''' or ph.ID == parentID (self-parent). CollectTasks takes an internal helper func with (phaseID string, visited map[string]bool) — if visited[phaseID] return, mark visited, recurse. Preserve the public CollectTasks(phaseID string) signature by wrapping the helper. Same guard shape for AllTasks if it also recurses.'
      id: phase_1
      name: Model — cycle-safe phase traversal
      notes: 'This fix alone would have prevented the 1GB stack crash on the stale plan. Back-compat preserved: visited-set is a no-op on valid acyclic data.'
      objective: No plan phase data shape can cause infinite recursion in the model layer
      status: completed
      tasks:
        - guard-childphases-and-collecttasks-against-cycles-and-empty-ids
    - changes: internal/cli/plan.go planAddPhaseCmd now requires --name (error if empty after trim), asserts the generated phase_N id is non-empty and does not collide with an existing phase ID (defensive — should never trip given the counter, but blocks hand-edited races), and rejects writes when the resulting PlanPhase would have an empty ID field for any reason
      description: Reject malformed add-phase invocations at the CLI boundary
      details: Add the checks just before the append to p.Phases. If addPhaseName is blank, return fmt.Errorf. Loop existing phases to detect ID collisions. Error messages should be one line with the remediation hint.
      id: phase_2
      name: CLI — add-phase input validation
      notes: The corrupted stale plan was written by some older code path — we cannot know exactly what, but hardening add-phase closes the CLI side regardless.
      objective: syde plan add-phase cannot create phases with empty IDs or names
      status: completed
      tasks:
        - require-name-and-non-empty-id-in-syde-plan-add-phase
    - changes: 'internal/audit adds a planPhasesFindings() function called from Run(). Detects: (a) phase with empty ID, (b) duplicate phase IDs, (c) phase whose ParentPhase equals its own ID (self-loop), (d) ParentPhase cycles across two or more phases, (e) ParentPhase pointing at an unknown phase ID in the same plan. All as ERROR severity under a new category name like CatPlanPhase.'
      description: New audit finding for plans with malformed phase data, ERROR severity
      details: 'Walk all plans in audit input. For each plan, build idSet and parent map. Emit findings. Cycle detection: simple DFS from every phase following ParentPhase, using a local visited set + on-stack set; if we revisit on-stack we have a cycle. Integrate into Run() alongside existing checks.'
      id: phase_3
      name: Audit — sync check reports corrupt plan phases
      notes: Add a new CatPlanPhase constant. Make sure sync check --errors-only still exits non-zero on the new findings.
      objective: syde sync check --strict surfaces broken plan phase structure so the session-end gate blocks until the plan is fixed or deleted
      status: completed
      tasks:
        - audit-error-findings-for-corrupt-plan-phase-data
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) cycle-safe PlanEntity.ChildPhases + CollectTasks (visited set, self-parent and empty-ID filtering), (2) CLI add-phase input validation (require --name, assert non-empty resulting ID, reject duplicate IDs), (3) audit findings in sync check for empty-ID phases / self-parent / parent cycles / dangling parent references, ERROR severity. Out-of-scope: syde plan repair command, rewriting the plan renderer, changing phase YAML shape, back-fixing existing plans (user must delete + recreate if they have corrupt ones).'
slug: guard-plan-phase-integrity-5tio
source: manual
updated_at: "2026-04-14T08:59:42Z"
---
