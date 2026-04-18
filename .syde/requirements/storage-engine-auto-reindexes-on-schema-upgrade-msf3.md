---
id: REQ-0180
kind: requirement
name: Storage Engine Auto Reindexes On Schema Upgrade
slug: storage-engine-auto-reindexes-on-schema-upgrade-msf3
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:46Z"
statement: When the persisted index schema version is older than the current version, the storage engine shall run a full reindex at store open and update the stored schema version.
req_type: functional
priority: must
verification: inspection of NewStore schema version handling in store.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Self-healing index upgrades prevent stale layouts after binary updates.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:46Z"
---
