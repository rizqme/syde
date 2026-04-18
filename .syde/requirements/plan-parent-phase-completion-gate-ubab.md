---
id: REQ-0071
kind: requirement
name: Plan Parent Phase Completion Gate
slug: plan-parent-phase-completion-gate-ubab
relationships:
    - target: plan-sk33
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:42Z"
statement: If any child phase of a parent phase is not completed, then the syde CLI shall not allow the parent phase to transition to completed.
req_type: constraint
priority: must
verification: integration test completing a parent phase with an incomplete child
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: Preserves bottom-up completion semantics in the phase tree.
audited_overlaps:
    - slug: plan-phase-completion-gate-1jat
      distinction: Blocks parent phase completion based on child phase status; target blocks phase completion based on task status.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:42Z"
---
