---
id: TSK-0119
kind: task
name: Convert decisions to EARS requirements
slug: convert-decisions-to-ears-requirements-yn4f
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:15:40Z"
task_status: completed
priority: high
objective: Every existing DecisionEntity is replaced by a new RequirementEntity file whose statement is EARS-compliant.
details: Read each decision's statement, rationale, category, alternatives_considered, tradeoffs, consequences. Rewrite statement into EARS form. Map category to req_type (cross-cutting/architecture -> constraint, behavior -> functional, performance -> non-functional). Preserve rationale verbatim. Use consequences as verification. Bundle alternatives into body. Re-link any entity pointing at the old decision to the new requirement.
acceptance: .syde/decisions is empty; new requirement files exist with EARS statements; no inbound references to deleted decisions.
affected_entities:
    - storage-engine-ahgm
    - entity-model-f28o
    - slug-and-id-utils-8kr7
    - summary-tree-fq6u
    - syde-cli-2478
    - syde-5tdt
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_1
created_at: "2026-04-15T09:53:21Z"
completed_at: "2026-04-15T10:15:40Z"
---
