---
id: TSK-0254
kind: task
name: Rename syded-dashboard-e82c name field to 'syded' and remove belongs_to:syde
slug: rename-syded-dashboard-e82c-name-field-to-syded-and-remove-belongstosyde-kglv
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: design-model-shall-contain-exactly-two-systems-named-syde-and-syded-f2q8
      type: implements
updated_at: "2026-04-18T09:44:53Z"
task_status: completed
priority: high
objective: The syded system's name field reads 'syded' and it carries no belongs_to relationship.
acceptance: syde query syded-dashboard-e82c shows name 'syded' and zero belongs_to edges.
affected_entities:
    - syded-dashboard-e82c
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_2
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:28:46Z"
---
