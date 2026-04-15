---
id: REQ-0115
kind: requirement
name: Task Affected Entities Must Resolve
slug: task-affected-entities-must-resolve-wxl0
relationships:
    - target: task-d3oc
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:33Z"
statement: If a task lists an affected entity slug that does not resolve, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a task with an unknown affected entity
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Dangling affected entities break traceability and auto-bump on task done.
---
