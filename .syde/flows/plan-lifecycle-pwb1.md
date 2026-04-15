---
description: Clarify → plan → approve → implement → finish workflow for a non-trivial change.
edge_cases: Plan discovered mid-work to be wrong → 'syde plan update' or abort and create new plan. Task needs to create new entity → notes it in phase, runs 'syde add' during implementation, then 'syde task update --affected-entity'. User wants to jump ahead → agent refuses, reminds of phase order.
failure_modes: Approval given before clarification is complete → bad plan. Skipping 'syde task done' leaves affected entities flagged stale. Skipping tree refresh at session end leaves rot for the next session.
goal: Produce an approved plan with phases and tasks, then implement it phase-by-phase with the model staying in sync
happy_path: clarify → plan → approve → execute → task-by-task → finish with clean tree
id: FLW-0002
kind: flow
name: Plan Lifecycle
narrative: 'Phase 1: agent clarifies requirements critically and waits for user confirmation. Phase 2: agent creates plan with background/objective/scope, adds phases with objective/changes/details, adds tasks with objective/details/acceptance/affected-entities/affected-files. Phase 3: agent presents plan with ''syde plan show --full'' and stops. User approves in chat. Agent runs ''syde plan approve'' then ''syde plan execute''. Phase 4: for each task agent runs ''syde task start'', writes code, runs ''syde constraints check'' on new files, updates entities, runs ''syde task done'' (auto-bumps updated_at on affected entities). Phase 5: agent runs ''syde validate'', ''syde sync --check'', and ''syde tree scan'' + subagent summarize loop + ''syde tree status --strict''.'
relationships:
    - target: syde
      type: belongs_to
    - target: plan
      type: involves
    - target: task
      type: involves
    - target: cli-commands
      type: involves
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: plan-lifecycle-pwb1
trigger: User asks the agent to implement a non-trivial change
updated_at: "2026-04-15T06:30:48Z"
---
