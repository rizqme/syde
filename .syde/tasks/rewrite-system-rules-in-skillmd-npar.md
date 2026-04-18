---
id: TSK-0259
kind: task
name: Rewrite system rules in SKILL.md
slug: rewrite-system-rules-in-skillmd-npar
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: skill-documentation-shall-describe-systems-as-standalone-processes-rsf8
      type: implements
updated_at: "2026-04-18T09:44:53Z"
task_status: completed
priority: medium
objective: SKILL.md system rules section describes systems as standalone top-level processes with no belongs_to and multi-system belongs_to on components.
acceptance: 'skill/SKILL.md mentions: (a) systems are top-level, (b) systems must not have belongs_to, (c) components may carry multiple belongs_to:system edges.'
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:32:04Z"
---
