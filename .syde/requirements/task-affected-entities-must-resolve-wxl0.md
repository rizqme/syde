---
id: REQ-0115
kind: requirement
name: Task Affected Entities Must Resolve
slug: task-affected-entities-must-resolve-wxl0
relationships:
    - target: task-d3oc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:00Z"
statement: If a task lists an affected entity slug that does not resolve, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a task with an unknown affected entity
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Dangling affected entities break traceability and auto-bump on task done.
audited_overlaps:
    - slug: task-affected-files-must-exist-in-tree-0ykl
      distinction: Entity-slug resolution checks entity existence; file-presence check validates paths against the summary tree, different reference types.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:00Z"
---
