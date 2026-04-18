---
id: REQ-0116
kind: requirement
name: Task Affected Files Must Exist In Tree
slug: task-affected-files-must-exist-in-tree-0ykl
relationships:
    - target: task-d3oc
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:34Z"
statement: If a task lists an affected file that is not present in the summary tree, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a task referencing a missing file path
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Affected files must be real so constraint checks and reports are accurate.
audited_overlaps:
    - slug: task-affected-entities-must-resolve-wxl0
      distinction: File presence validation against the summary tree differs from entity slug resolution against the entity store.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:34Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:34Z"
---
