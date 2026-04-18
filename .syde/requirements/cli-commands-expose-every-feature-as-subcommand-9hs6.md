---
id: REQ-0061
kind: requirement
name: CLI Commands Expose Every Feature As Subcommand
slug: cli-commands-expose-every-feature-as-subcommand-9hs6
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:27Z"
statement: The syde CLI shall expose every syde feature as a cobra subcommand under a single root binary.
req_type: functional
priority: must
verification: inspection of internal/cli/root.go command tree
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: A unified command tree is the sole user interface for the syde model.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:27Z"
---
