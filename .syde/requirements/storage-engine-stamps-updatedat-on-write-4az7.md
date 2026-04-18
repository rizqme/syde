---
id: REQ-0181
kind: requirement
name: Storage Engine Stamps UpdatedAt On Write
slug: storage-engine-stamps-updatedat-on-write-4az7
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:38Z"
statement: When Create or Update is called, the storage engine shall stamp BaseEntity.UpdatedAt with the current time and reindex the affected entity.
req_type: functional
priority: must
verification: unit test of Create and Update in store.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: UpdatedAt drives drift detection against source files.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:38Z"
---
