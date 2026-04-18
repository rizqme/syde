---
id: REQ-0185
kind: requirement
name: Storage Engine Does Not Own Schema
slug: storage-engine-does-not-own-schema-b7ui
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:02Z"
statement: The storage engine shall not own entity schema definitions and shall delegate per-kind struct definitions to the entity model.
req_type: constraint
priority: must
verification: code review of internal/storage for struct definitions
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Schema lives in model so storage can remain kind-agnostic.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:02Z"
---
