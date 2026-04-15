---
id: TSK-0109
kind: task
name: Document outgoing traceability invariant
slug: document-outgoing-traceability-invariant-urel
relationships:
    - target: enforce-outgoing-requirement-traceability-lo0c
      type: belongs_to
    - target: strict-sync-requires-outgoing-requirement-traceability-yvdj
      type: references
updated_at: "2026-04-15T08:09:36Z"
task_status: completed
priority: medium
objective: Align agent-facing docs with strict sync's outgoing requirement-link rule.
details: Update skill/SKILL.md and skill/references/entity-spec.md to say every non-requirement entity carries an outbound relationship to a requirement, including components and contracts.
acceptance: Skill docs explicitly say outbound requirement relationship and name components/contracts as covered.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/codex/SKILL.md
plan_ref: enforce-outgoing-requirement-traceability-lo0c
plan_phase: phase_2
created_at: "2026-04-15T08:08:13Z"
completed_at: "2026-04-15T08:09:36Z"
---
