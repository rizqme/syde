---
id: REQ-0182
kind: requirement
name: Storage Engine Cascades UpdatedAt Through belongs_to
slug: storage-engine-cascades-updatedat-through-belongsto-k97y
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:28Z"
statement: When CreateCascade or UpdateCascade or DeleteCascade is called, the storage engine shall propagate UpdatedAt bumps up the belongs_to chain using a visited-ID map so cycles terminate.
req_type: functional
priority: must
verification: unit test of cascade helpers in store.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Parent entities must surface downstream edits without infinite loops.
---
