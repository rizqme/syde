---
id: REQ-0478
kind: requirement
name: Active requirement shall refine at least one component
slug: active-requirement-shall-refine-at-least-one-component-mke4
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:04Z"
statement: When an active requirement has no refines edge to a component, the syde audit engine shall report a finding.
req_type: functional
priority: must
verification: syde sync check errors when an active requirement is created or migrated without a refines:component edge
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Bidirectional invariant requires every requirement to constrain a concrete component.
audited_overlaps:
    - slug: component-with-mapped-files-shall-have-at-least-one-refining-requirement-300f
      distinction: This rule fires per-requirement (forward direction — req lacks refines:component); the counterpart fires per-component (reverse — component lacks incoming refines). Both halves of the bidirectional invariant, triggered by different entity kinds.
    - slug: requirement-shall-not-refine-or-belong-to-a-system-m4c2
      distinction: This rule fires when refines:component is missing entirely; the counterpart fires only when refines/belongs_to points at a system. Distinct failure modes with different fix actions (add edge vs remove edge).
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:04Z"
---
