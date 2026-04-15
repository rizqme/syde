---
id: REQ-0279
kind: requirement
name: Link Task To Design Invocation
slug: link-task-to-design-invocation-lae8
relationships:
    - target: link-task-to-design-ooqq
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde task link <task-slug> <entity-slug>, the syde CLI shall create a link between the named task and the target design entity.
req_type: interface
priority: must
verification: integration test invoking syde task link and inspecting relationships
source: manual
source_ref: contract:link-task-to-design-ooqq
requirement_status: active
rationale: Explicit task-to-design links preserve traceability between work and architecture.
---
