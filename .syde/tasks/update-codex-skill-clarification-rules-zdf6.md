---
id: TSK-0102
kind: task
name: Update Codex skill clarification rules
slug: update-codex-skill-clarification-rules-zdf6
relationships:
    - target: require-tasks-before-plan-approval-cs0s
      type: belongs_to
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
updated_at: "2026-04-15T07:07:40Z"
task_status: completed
priority: high
objective: Make the Codex skill require the ask-user-question/request_user_input tool for clarification when available.
details: Edit skill/codex/SKILL.md to instruct Codex to use the question tool for clarification before planning, then mirror approved syde tasks into update_plan.
acceptance: Codex skill explicitly names the question tool behavior and task synchronization behavior.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/codex/SKILL.md
plan_ref: require-tasks-before-plan-approval-cs0s
plan_phase: phase_1
created_at: "2026-04-15T06:55:02Z"
completed_at: "2026-04-15T07:01:23Z"
---
