---
id: REQ-0481
kind: requirement
name: Requirement shall be marked stale when refining component file content changes
slug: requirement-shall-be-marked-stale-when-refining-component-file-content-changes-85v0
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:38:01Z"
statement: Where any file in a component refining a requirement has a SHA-256 content hash differing from the requirement's stored verified_against entry, the syde audit engine shall report a finding.
req_type: functional
priority: must
verification: Editing a file in a refining component without subsequently running syde requirement verify on the requirement causes syde sync check to error
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Content-hash chosen over mtime for resilience against copy/touch operations and to ignore format-only changes.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:38:01Z"
---
