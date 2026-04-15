---
id: REQ-0185
kind: requirement
name: Storage Engine Does Not Own Schema
slug: storage-engine-does-not-own-schema-b7ui
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:35Z"
statement: The storage engine shall not own entity schema definitions and shall delegate per-kind struct definitions to the entity model.
req_type: constraint
priority: must
verification: code review of internal/storage for struct definitions
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Schema lives in model so storage can remain kind-agnostic.
---
