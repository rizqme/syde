---
id: REQ-0480
kind: requirement
name: Requirement shall not refine or belong to a system
slug: requirement-shall-not-refine-or-belong-to-a-system-m4c2
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:45:46Z"
statement: If a requirement carries a refines edge to a system or a belongs_to edge to a system, then the syde audit engine shall report a finding.
req_type: functional
priority: must
verification: syde sync check errors when any active requirement targets an entity of kind system via refines or belongs_to
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Cross-cutting reqs hide which component is actually responsible; force decomposition.
audited_overlaps:
    - slug: active-requirement-shall-refine-at-least-one-component-mke4
      distinction: This rule fires when refines/belongs_to points at a system (disallowed target); the counterpart fires when no component target exists at all. Distinct failure modes with different fix actions.
    - slug: system-entities-shall-not-carry-belongs-to-6hg2
      distinction: distinct — requirement-shall-not-refine targets the requirement entity rule; system-entities targets the system entity rule
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:48Z"
---
