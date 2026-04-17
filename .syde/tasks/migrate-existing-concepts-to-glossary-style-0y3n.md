---
id: TSK-0084
kind: task
name: Migrate existing concepts to glossary style
slug: migrate-existing-concepts-to-glossary-style-0y3n
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: All 10 concepts have meaning+invariants, no attributes/actions, with role-based links
details: 'Batch script: syde update each concept to strip old fields and add implemented_by/exposed_via/used_in relationships'
acceptance: syde query --kind concept shows all concepts with role-based relationships
affected_entities:
    - decision-m2um
    - entity-8x6p
    - fileref-7vac
    - plan-phase-23bb
    - plan-sk33
    - relationship-hjgt
    - skill-7fmf
    - summary-tree-u2fo
    - task-d3oc
    - tree-node-iutv
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_4
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T08:25:52Z"
---
