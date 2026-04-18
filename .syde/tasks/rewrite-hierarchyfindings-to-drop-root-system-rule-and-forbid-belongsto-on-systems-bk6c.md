---
id: TSK-0252
kind: task
name: Rewrite hierarchyFindings to drop root-system rule and forbid belongs_to on systems
slug: rewrite-hierarchyfindings-to-drop-root-system-rule-and-forbid-belongsto-on-systems-bk6c
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: system-entities-shall-not-carry-belongs-to-6hg2
      type: implements
updated_at: "2026-04-18T09:44:35Z"
task_status: completed
priority: high
objective: hierarchyFindings emits a finding for any system carrying belongs_to and for any non-system non-req entity lacking belongs_to; no root-system selection happens.
acceptance: Adding a belongs_to edge to a system entity triggers a sync-check finding; sync check tolerates 0/1/N systems without belongs_to.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_1
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:27:34Z"
---
