---
id: REQ-0078
kind: requirement
name: Component Depends-On DAG
slug: component-depends-on-dag-43z6
relationships:
    - target: relationship-hjgt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:08Z"
statement: If adding a depends_on relationship between components would create a cycle, then the syde CLI shall reject the save with a cycle error.
req_type: constraint
priority: must
verification: integration test adding a depends_on link that closes a cycle
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: Cyclic component dependencies make layering and build order impossible to compute.
audited_overlaps:
    - slug: system-belongs-to-acyclic-79sv
      distinction: 'Different entity kind and relationship: this guards depends_on cycles between components, while the other guards belongs_to cycles between systems.'
---
