---
id: REQ-0145
kind: requirement
name: Entity Model Does Not Resolve Relationships
slug: entity-model-does-not-resolve-relationships-7aww
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:36Z"
statement: The entity model shall not resolve relationship targets and shall leave traversal to the query and graph engines.
req_type: constraint
priority: must
verification: code review of internal/model for store imports
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Keeping the model acyclic with respect to storage avoids bootstrap loops.
---
