---
id: REQ-0116
kind: requirement
name: Task Affected Files Must Exist In Tree
slug: task-affected-files-must-exist-in-tree-0ykl
relationships:
    - target: task-d3oc
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:35Z"
statement: If a task lists an affected file that is not present in the summary tree, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a task referencing a missing file path
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Affected files must be real so constraint checks and reports are accurate.
---
