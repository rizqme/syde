---
id: REQ-0078
kind: requirement
name: Component Depends-On DAG
slug: component-depends-on-dag-43z6
relationships:
    - target: relationship-hjgt
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:07Z"
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
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:07Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:07Z"
---
