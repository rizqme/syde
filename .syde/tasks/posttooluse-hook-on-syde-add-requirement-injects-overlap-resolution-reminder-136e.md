---
id: TSK-0110
kind: task
name: PostToolUse hook on syde add requirement injects overlap resolution reminder
slug: posttooluse-hook-on-syde-add-requirement-injects-overlap-resolution-reminder-136e
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-data-model-cli-hook-docs-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: After syde add requirement prints an overlap banner, the session receives a system reminder that must be resolved before further tool use
details: 'Edit skill/hooks.json: add a PostToolUse entry with matcher on Bash containing ''syde add requirement'' that greps stdout for ''⚠ Similar to''. If matched, emit a system reminder listing the overlapped slugs and naming the three resolution paths (MERGE via syde remove + new survivor, RENAME via syde update --statement, DISTINCT via syde update --audited slug:reason). Reinstall via syde install-skill --all.'
acceptance: After syde install-skill --all, a syde add requirement call with high overlap triggers a visible system reminder in the session transcript
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/hooks.json
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_4
created_at: "2026-04-17T09:40:20Z"
completed_at: "2026-04-17T10:16:48Z"
---
