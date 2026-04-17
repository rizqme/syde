---
id: REQ-0349
kind: requirement
name: Flow steps shall describe user and system actions
slug: flow-steps-shall-describe-user-and-system-actions-gs18
description: Step action text names the actor
relationships:
    - target: syde
      type: belongs_to
updated_at: "2026-04-16T10:40:59Z"
statement: When a flow has structured steps, the syde skill shall require each step action to describe what the user or system does at that point in the journey.
req_type: usability
priority: should
verification: Skill docs teach action text conventions
source: plan
requirement_status: active
rationale: Steps like 'do thing' are useless; 'User runs syde plan create' is actionable
---
