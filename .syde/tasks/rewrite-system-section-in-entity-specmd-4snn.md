---
id: TSK-0260
kind: task
name: Rewrite System section in entity-spec.md
slug: rewrite-system-section-in-entity-specmd-4snn
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: skill-documentation-shall-describe-systems-as-standalone-processes-rsf8
      type: implements
updated_at: "2026-04-18T09:44:53Z"
task_status: completed
priority: medium
objective: entity-spec.md System section matches the standalone-processes semantics.
acceptance: entity-spec.md System section no longer mentions 'root' or 'sub-system' and documents multi-system belongs_to on components.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/references/entity-spec.md
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:32:53Z"
---
