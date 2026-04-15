---
acceptance: Build syde, run ./syde install-skill in scratch dir, verify updated files land in .claude/skills/syde/.
affected_entities:
    - skill-installer
affected_files:
    - skill/SKILL.md
    - skill/references/commands.md
completed_at: "2026-04-14T08:25:53Z"
created_at: "2026-04-14T08:08:10Z"
details: 'skill/SKILL.md: expand ''Finding Files to Read'' with four recipes (search+kind, --file, --search+--tag+--limit, directory prefix). skill/references/commands.md: document --search, --file, --limit, --any, --no-related with examples. Mention body indexing.'
id: TSK-0047
kind: task
name: Update SKILL.md and commands.md with new query surface
objective: Agents reading the skill learn to use syde query for keyword and file-based lookups instead of naive Read/grep
plan_phase: phase_4
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: update-skillmd-and-commandsmd-with-new-query-surface-jvr0
task_status: completed
updated_at: "2026-04-14T08:25:53Z"
---
