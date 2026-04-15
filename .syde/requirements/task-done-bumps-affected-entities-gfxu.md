---
id: REQ-0118
kind: requirement
name: Task Done Bumps Affected Entities
slug: task-done-bumps-affected-entities-gfxu
relationships:
    - target: task-d3oc
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:37Z"
statement: When a task transitions to completed, the syde CLI shall update the updated_at timestamp on every one of its affected entities.
req_type: functional
priority: must
verification: integration test running syde task done and asserting updated_at changes
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Auto-bump keeps staleness tracking honest when work touches design nodes.
---
