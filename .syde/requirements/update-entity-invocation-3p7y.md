---
id: REQ-0316
kind: requirement
name: Update Entity Invocation
slug: update-entity-invocation-3p7y
relationships:
    - target: update-entity-zpnh
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-17T10:46:07Z"
statement: When the user runs syde update <slug>, the syde CLI shall apply the provided field flags to the named entity and print its updated slug.
req_type: interface
priority: must
verification: integration test invoking syde update with various field flags
source: manual
source_ref: contract:update-entity-zpnh
requirement_status: active
rationale: Update is the canonical mutation command for existing entities.
audited_overlaps:
    - slug: update-plan-invocation-vrec
      distinction: syde update targets generic entities; syde plan update is the plan-specific command with distinct invocation syntax.
    - slug: update-task-invocation-q9hh
      distinction: syde update targets generic entities; syde task update is the task-specific command with distinct invocation syntax.
---
