---
id: TSK-0019
kind: task
name: Update skill docs for new plan workflow
slug: update-skill-docs-for-new-plan-workflow-uibg
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: skill/SKILL.md, skill/codex/SKILL.md, skill/references/commands.md, skill/references/entity-spec.md, and skill/references/sync-workflow.md all teach the new design+diff plan workflow.
details: 'Add a ''Planning'' section to SKILL.md describing: write --design, declare changes with syde plan add-change (delete/extend/new), review with show-changes, execute tasks, run syde plan complete which invokes the validator. entity-spec.md gets a Plan kind-specific field table showing Design and Changes with DeletedChange/ExtendedChange/NewChange sub-schemas. commands.md documents the new subcommands with examples. sync-workflow.md references the plan completion validator.'
acceptance: rg -l 'plan add-change|planCompletionFindings|Design field' skill/ returns matches in every expected file.
affected_entities:
    - skill-installer-wbmu
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_5
created_at: "2026-04-15T11:42:14Z"
completed_at: "2026-04-15T12:41:04Z"
---
