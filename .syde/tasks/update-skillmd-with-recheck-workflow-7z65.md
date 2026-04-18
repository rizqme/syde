---
id: TSK-0248
kind: task
name: Update SKILL.md with recheck workflow
slug: update-skillmd-with-recheck-workflow-7z65
updated_at: '2026-04-18T08:20:55Z'
task_status: completed
priority: low
objective: SKILL.md gains a 'Recheck affected requirements' subsection in the workflow rules
acceptance: 'SKILL.md contains explicit guidance: edit a file → read affected requirements list from hook → re-verify each with syde requirement verify <slug>'
affected_files:
- skill/SKILL.md
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_5
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:20:55Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: posttooluse-hook-shall-surface-affected-requirements-when-a-file-mapped-to-a-component-is-edited-ofeu
---
