---
id: REQ-0004
kind: requirement
name: Clarification must use ask-user-question tool
slug: clarification-must-use-ask-user-question-tool-97sm
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T06:53:47Z"
statement: Agents must clarify requirements before planning by using the ask-user-question/request_user_input tool when available, and must not proceed until the user answers.
source: user
source_ref: user:clarification-ask-user-question-tool
requirement_status: active
rationale: The user explicitly requested enforcement of clarification through the question tool in the syde skill.
acceptance_criteria: The syde skills instruct agents to use the question tool for clarification before plan creation, with a clear fallback only when the runtime does not expose that tool.
---
