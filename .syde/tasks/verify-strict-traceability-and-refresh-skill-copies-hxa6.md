---
id: TSK-0110
kind: task
name: Verify strict traceability and refresh skill copies
slug: verify-strict-traceability-and-refresh-skill-copies-hxa6
relationships:
    - target: enforce-outgoing-requirement-traceability-lo0c
      type: belongs_to
    - target: strict-sync-requires-outgoing-requirement-traceability-yvdj
      type: references
updated_at: "2026-04-15T08:12:23Z"
task_status: completed
priority: medium
objective: Rebuild/install and prove strict sync catches missing outbound requirement links while the real model passes.
details: Run go test, make install, syde install-skill --all, a temporary broken-link smoke test, tree status, and sync strict.
acceptance: Tests pass, smoke test fails as expected, and syde sync check --strict passes after restoration.
affected_entities:
    - audit-engine-4ktg
    - skill-installer-wbmu
affected_files:
    - internal/audit/graph_rules.go
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/codex/SKILL.md
plan_ref: enforce-outgoing-requirement-traceability-lo0c
plan_phase: phase_2
created_at: "2026-04-15T08:08:13Z"
completed_at: "2026-04-15T08:12:23Z"
---
