---
id: REQ-0382
kind: requirement
name: Sync check shall error on unaudited requirement overlaps
slug: sync-check-shall-error-on-unaudited-requirement-overlaps-dsrz
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:38:09Z"
statement: If two active requirements overlap above 50 percent similarity without mutual acknowledgement, then the syde sync check engine shall report an error.
req_type: constraint
priority: must
verification: sync check errors when two requirements overlap and neither acknowledges the other.
source: plan
requirement_status: active
rationale: Mandatory acknowledgement forces the author to confirm or resolve overlaps.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:38:09Z"
---
