---
id: REQ-0279
kind: requirement
name: Link Task To Design Invocation
slug: link-task-to-design-invocation-lae8
relationships:
    - target: link-task-to-design-ooqq
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:42Z"
statement: When the user runs syde task link <task-slug> <entity-slug>, the syde CLI shall create a link between the named task and the target design entity.
req_type: interface
priority: must
verification: integration test invoking syde task link and inspecting relationships
source: manual
source_ref: contract:link-task-to-design-ooqq
requirement_status: active
rationale: Explicit task-to-design links preserve traceability between work and architecture.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:42Z"
---
