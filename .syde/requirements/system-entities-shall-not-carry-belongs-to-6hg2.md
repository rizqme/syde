---
id: REQ-0485
kind: requirement
name: System entities shall not carry belongs to
slug: system-entities-shall-not-carry-belongs-to-6hg2
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:55:16Z"
statement: If a system entity carries a belongs_to relationship, then the syde audit engine shall report a finding.
req_type: functional
priority: must
verification: syde sync check errors when any system entity has a relationship of type belongs_to
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Root-system model is retired; every system is an independent process boundary.
audited_overlaps:
    - slug: requirement-shall-not-refine-or-belong-to-a-system-m4c2
      distinction: distinct — system-entities targets the system entity rule; requirement-shall-not-refine targets the requirement entity rule
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:55:16Z"
---
