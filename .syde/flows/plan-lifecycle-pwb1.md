---
id: FLW-0002
kind: flow
name: Plan Lifecycle
slug: plan-lifecycle-pwb1
description: Clarify → plan → approve → implement → finish workflow for a non-trivial change.
relationships:
    - target: syde
      type: belongs_to
    - target: plan
      type: involves
    - target: task
      type: involves
    - target: cli-commands
      type: involves
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T10:59:47Z"
trigger: User asks the agent to implement a non-trivial change
goal: Produce an approved plan with phases and tasks, then implement it phase-by-phase with the model staying in sync
steps:
    - id: s1
      action: User describes change
      description: Agent identifies the request
      on_success: s2
    - id: s2
      action: Agent clarifies requirements
      description: Probes for missing info
      on_success: s3
    - id: s3
      action: Agent creates plan
      contract: create-plan
      description: Creates with background/objective/scope/design
      on_success: s4
    - id: s4
      action: Agent adds phases
      contract: add-plan-phase
      description: Adds with objective/changes/details
      on_success: s5
    - id: s5
      action: Agent creates tasks
      contract: create-task
      description: Creates for each phase
      on_success: s6
    - id: s6
      action: Agent opens plan in dashboard
      contract: plan-view-screen
      description: Opens browser tab
      on_success: s7
    - id: s7
      action: User approves plan
      contract: approve-plan
      description: Transitions to approved
      on_success: s8
    - id: s8
      action: Agent starts task
      contract: start-task
      description: Marks in_progress
      on_success: s9
    - id: s9
      action: Agent implements code
      description: Writes files
      on_success: s10
      on_failure: s11
    - id: s10
      action: Agent completes task
      contract: complete-task
      description: Marks done with affected entities/files
      on_success: s8
      on_failure: done
    - id: s11
      action: Agent blocks task
      contract: block-task
      description: Records reason, seeks resolution
      on_success: s8
narrative: 'Phase 1: agent clarifies requirements critically and waits for user confirmation. Phase 2: agent identifies the underlying requirement (search existing requirements first; mark conflicting ones superseded or obsolete), then drafts a plan with background/objective/scope/design, populates the structured Changes block lane-by-lane (Requirements first, then Components/Contracts/Concepts/Flows), adds phases with objective/changes/details, adds tasks with objective/details/acceptance/affected-entities/affected-files. Phase 2.5: agent runs syde plan check, addresses every ERROR and reviews every WARN, then runs syde plan open to surface the plan in the user''s existing dashboard tab. Phase 3: agent presents plan with summary of caught gaps, user approves via syde plan approve. Phase 4: agent executes tasks one phase at a time without pausing for permission. Phase 5: agent runs syde plan complete which invokes planCompletionFindings to verify every declared change against actual entity state.'
happy_path: clarify → cascade requirement-first → draft plan with Design + Changes (Requirements/Components/Contracts/Concepts/Flows lanes) + phases + tasks → syde plan check (must exit 0) → syde plan open → user approves → execute task-by-task → syde plan complete (validator-gated) → finish with clean tree
edge_cases: Plan discovered mid-work to be wrong → 'syde plan update' or abort and create new plan. Task needs to create new entity → notes it in phase, runs 'syde add' during implementation, then 'syde task update --affected-entity'. User wants to jump ahead → agent refuses, reminds of phase order.
failure_modes: Approval given before clarification is complete → bad plan. Skipping 'syde task done' leaves affected entities flagged stale. Skipping tree refresh at session end leaves rot for the next session.
---
