---
id: REQ-0145
kind: requirement
name: Entity Model Does Not Resolve Relationships
slug: entity-model-does-not-resolve-relationships-7aww
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:36:47Z"
statement: The entity model shall not resolve relationship targets and shall leave traversal to the query and graph engines.
req_type: constraint
priority: must
verification: code review of internal/model for store imports
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Keeping the model acyclic with respect to storage avoids bootstrap loops.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:47Z"
---
