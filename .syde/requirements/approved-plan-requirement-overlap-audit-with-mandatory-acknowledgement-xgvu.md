---
id: REQ-0364
kind: requirement
name: 'Approved plan: Requirement overlap audit with mandatory acknowledgement'
slug: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
relationships:
    - target: requirement-overlap-audit-with-mandatory-acknowledgement-u5lj
      type: references
      label: approved_plan
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:52Z"
statement: When syde add requirement creates a new requirement, the syde CLI shall compute TF-IDF similarity against every active requirement, print each overlap candidate, and require --audited acknowledgement for every match.
req_type: functional
priority: must
verification: syde add requirement prints overlap candidates; sync check errors on unacknowledged overlaps and clears on --audited entries
source: plan
source_ref: plan:requirement-overlap-audit-with-mandatory-acknowledgement-u5lj
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-16T11:33:51Z"
audited_overlaps:
    - slug: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      distinction: Triggers on syde add requirement requiring --audited acknowledgement for TF-IDF matches vs. triggers on syde plan complete blocking on sync errors — different command and different subject entity.
    - slug: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      distinction: 'Different plan scope: overlap audit governs requirement creation, while the phase-collision plan fixes task-done phase auto-completion logic.'
    - slug: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      distinction: 'Different plan scope: overlap audit runs at requirement add-time in CLI, while the task-resolution plan fixes a syded dashboard API fallback.'
    - slug: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      distinction: 'Different plan scope: overlap audit enforces requirement dedup, while the flow-steps plan adds ordered contract-linked steps and flowchart rendering.'
    - slug: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      distinction: This plan enforces mandatory --audited acknowledgement at requirement creation; the coverage plan only emits advisory warnings for lane imbalance and TF-IDF overlap.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:52Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:52Z"
---
