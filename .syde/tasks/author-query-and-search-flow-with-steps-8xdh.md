---
id: TSK-0066
kind: task
name: Author Query and Search flow with steps
slug: author-query-and-search-flow-with-steps-8xdh
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: 'Author authoring + querying flows (8): Add Entity, View Entity, Update Entity, Remove Entity, List Entities, Search Model, View Relationships, Check Constraints'
details: Batch script. Covers entity CRUD contracts + storage key contracts (counter, entity-index, slug-index, tag-index, word-index, incoming-rel, outgoing-rel) + query/search/graph/constraints contracts
acceptance: All authoring and querying contracts appear in flow steps
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_4
created_at: "2026-04-16T09:23:28Z"
completed_at: "2026-04-16T10:53:00Z"
---
