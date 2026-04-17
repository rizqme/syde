---
id: REQ-0332
kind: requirement
name: 'User request: load syde skill and then can you check remaining task in current'
slug: user-request-load-syde-skill-and-then-can-you-check-remaining-task-in-current-r9vt
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-16T01:09:09Z"
statement: When the user asks Codex to inspect remaining syde task state, the syde workflow shall load syde context and report the current pending plan tasks.
req_type: functional
priority: must
verification: Manual inspection confirms the response used syde context and listed pending current-plan tasks.
source: user
source_ref: codex:019d93cd-a768-7703-b52c-5b1b7e204988:019d93cd-d822-7f03-8cbc-5f04779b9b0d
requirement_status: active
rationale: The user explicitly asked to load the syde skill and check remaining tasks in the current work.
approved_at: "2026-04-16T01:00:34Z"
---
