---
id: REQ-0092
kind: requirement
name: Relationship Requires Source And Target
slug: relationship-requires-source-and-target-n4bp
relationships:
    - target: relationship-hjgt
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:56Z"
statement: The syde CLI shall require a non-empty source and target on every relationship instance.
req_type: constraint
priority: must
verification: unit test rejecting a relationship with an empty source or target
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: A relationship without both endpoints cannot form a graph edge.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:56Z"
---
