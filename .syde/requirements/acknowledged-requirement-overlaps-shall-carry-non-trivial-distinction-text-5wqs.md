---
id: REQ-0401
kind: requirement
name: Acknowledged requirement overlaps shall carry non-trivial distinction text
slug: acknowledged-requirement-overlaps-shall-carry-non-trivial-distinction-text-5wqs
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:36:42Z"
statement: If an audited overlap entry on a requirement carries a distinction rationale shorter than 20 characters, then the syde audit engine shall report an error.
req_type: constraint
priority: must
verification: sync check reports errors for every acknowledgement whose distinction is empty or below 20 chars
source: plan
requirement_status: active
rationale: Audited acknowledgements must document semantic distinction or be treated as unresolved.
supersedes:
    - audited-overlap-acknowledgements-shall-carry-a-distinction-rationale-yruo
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:42Z"
---
