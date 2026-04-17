---
id: REQ-0071
kind: requirement
name: Plan Parent Phase Completion Gate
slug: plan-parent-phase-completion-gate-ubab
relationships:
    - target: plan-sk33
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:01Z"
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
---
