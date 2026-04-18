---
id: TSK-0265
kind: task
name: Draft references/plan-authoring.md
slug: draft-referencesplan-authoringmd-tybp
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: plan-phases-shall-list-created-and-modified-files-before-task-breakdown-hhvm
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: medium
objective: 'A new skill/references/plan-authoring.md exists and codifies the superpowers-derived plan-authoring discipline: Files section, bite-sized checkbox steps, No Placeholders blacklist, Self-Review checklist, Execution Handoff prompt.'
details: 'Write a single reference doc at skill/references/plan-authoring.md. Structure it into five sections in order: (1) Audience framing (assume executor has no context), (2) Files section per phase (Create/Modify/Test), (3) Bite-sized checkbox steps (2-5 min; test-first sequence where applicable; exact code blocks; exact commands; expected output), (4) No Placeholders blacklist (verbatim list with examples of good rewrites), (5) Self-Review checklist (spec coverage, placeholder scan, type consistency), (6) Execution Handoff (subagent-driven vs inline). Include one complete end-to-end example showing a Files section + two tasks with bite-sized steps.'
acceptance: skill/references/plan-authoring.md exists, is bundled via go:embed, and contains all five sections plus the end-to-end example. 'syde install-skill' copies it into .claude/skills/syde/references/.
affected_entities:
    - skill-installer-wbmu
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:24:42Z"
completed_at: "2026-04-18T09:34:03Z"
---
