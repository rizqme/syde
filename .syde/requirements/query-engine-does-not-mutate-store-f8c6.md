---
id: REQ-0159
kind: requirement
name: Query Engine Does Not Mutate Store
slug: query-engine-does-not-mutate-store-f8c6
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:14Z"
statement: The query engine shall not mutate entities in the store during read operations.
req_type: constraint
priority: must
verification: code review of internal/query for write calls
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Read operations must be side-effect free.
---
