---
id: REQ-0479
kind: requirement
name: Component with mapped files shall have at least one refining requirement
slug: component-with-mapped-files-shall-have-at-least-one-refining-requirement-300f
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:36:52Z"
statement: When a component has at least one file in its files list and no active requirement refines it, the syde audit engine shall report a finding.
req_type: functional
priority: must
verification: syde sync check errors when a component with files mapped has zero incoming active-requirement refines edges
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Symmetric half of the bidirectional invariant; design-phase components without files are exempt.
audited_overlaps:
    - slug: active-requirement-shall-refine-at-least-one-component-mke4
      distinction: This rule fires per-component (reverse direction — component lacks incoming refines); the counterpart fires per-requirement (forward — req lacks refines:component). Both halves of the bidirectional invariant, triggered by different entity kinds.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:52Z"
---
