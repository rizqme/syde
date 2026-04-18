---
id: REQ-0080
kind: requirement
name: System Belongs-To Acyclic
slug: system-belongs-to-acyclic-79sv
relationships:
    - target: relationship-hjgt
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:01Z"
statement: If adding a belongs_to relationship between systems would create a cycle, then the syde CLI shall reject the save with a cycle error.
req_type: constraint
priority: must
verification: integration test adding a belongs_to link that closes a cycle
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: Cyclic system nesting breaks ownership trees used by belongs_to queries.
audited_overlaps:
    - slug: component-depends-on-dag-43z6
      distinction: Systems form an acyclic belongs_to hierarchy; components form an acyclic depends_on graph, different entity kinds and relationship types.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:38:01Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:01Z"
---
