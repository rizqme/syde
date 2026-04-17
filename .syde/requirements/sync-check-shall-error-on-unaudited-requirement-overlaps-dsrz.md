---
id: REQ-0382
kind: requirement
name: Sync check shall error on unaudited requirement overlaps
slug: sync-check-shall-error-on-unaudited-requirement-overlaps-dsrz
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: If two active requirements overlap above 50 percent similarity without mutual acknowledgement, then the syde sync check engine shall report an error.
req_type: constraint
priority: must
verification: sync check errors when two requirements overlap and neither acknowledges the other.
source: plan
requirement_status: active
rationale: Mandatory acknowledgement forces the author to confirm or resolve overlaps.
---
