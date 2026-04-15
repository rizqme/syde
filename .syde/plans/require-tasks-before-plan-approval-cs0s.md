---
id: PLN-0016
kind: plan
name: Require Tasks Before Plan Approval
slug: require-tasks-before-plan-approval-cs0s
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: clarification-must-use-ask-user-question-tool-97sm
      type: references
    - target: plan-phases-need-granular-tasks-dk2b
      type: references
    - target: approved-plan-require-tasks-before-plan-approval-80kv
      type: references
      label: requirement
updated_at: "2026-04-15T07:08:44Z"
plan_status: completed
background: The agent workflow allowed plans to be approved with phases but no concrete tasks, causing Codex-visible todos and syde plan tasks to drift.
objective: Make clarification questions mandatory through the ask-user-question tool where available, require granular phase task lists before plan approval, and teach agents to mirror syde tasks into Codex/Claude visible todo tracking.
scope: Update syde skill guidance, Codex skill guidance, command references, and plan approval validation. Include clarification-tool usage and granular task decomposition. Do not redesign task execution semantics beyond requiring phase tasks before approval.
source: manual
created_at: "2026-04-15T06:51:17Z"
approved_at: "2026-04-15T06:58:27Z"
completed_at: "2026-04-15T07:08:44Z"
phases:
    - id: phase_1
      name: Skill workflow update
      status: completed
      description: Make pre-approval task creation and visible todo synchronization explicit in the skill instructions.
      objective: Agents clarify through the question tool, decompose every phase into granular syde tasks before approval, and mirror syde task status into visible todo tracking.
      changes: Update full syde skill, Codex skill, installed copies, and command references.
      details: Document request_user_input/ask-user-question usage during clarification; require granular tasks before presenting a plan; describe how Codex update_plan and Claude TodoWrite mirror syde task statuses.
      notes: Tasks in this phase are intentionally split by document surface so task progress stays visible and specific.
      tasks:
        - document-task-first-planning-workflow
        - update-codex-skill-clarification-rules
        - update-command-references-for-approval-gate
        - refresh-installed-skill-copies
    - id: phase_2
      name: Approval validation
      status: completed
      description: Reject plan approval when phases have no tasks.
      objective: syde plan approve fails until every phase has granular task coverage.
      changes: Update approval validation code and command docs; smoke test both failing and passing approval paths.
      details: Add a focused helper that inspects all plan phases for task coverage before status changes or requirement capture; wire it into plan approve; document and verify behavior.
      notes: Tasks in this phase separate validation code, command documentation, and smoke verification.
      tasks:
        - enforce-tasks-during-plan-approval
        - wire-approval-validation-before-mutation
        - smoke-test-plan-approval-task-gate
---
