---
id: REQ-0324
kind: requirement
name: Counter key shall track per-kind ID allocation
slug: counter-key-shall-track-per-kind-id-allocation-bnpe
relationships:
    - target: counter-key-4b8h
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:46Z"
statement: The syde storage layer shall persist the highest issued entity counter per kind under BadgerDB key 'c:<kind>'.
req_type: interface
priority: must
verification: Integration test allocating consecutive entity IDs and inspecting the counter key
source: manual
source_ref: contract:counter-key-4b8h
requirement_status: active
rationale: Counter keys are the source of truth for monotonic ID allocation; reindex recomputes from the highest observed ID.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:46Z"
---
