---
id: REQ-0061
kind: requirement
name: CLI Commands Expose Every Feature As Subcommand
slug: cli-commands-expose-every-feature-as-subcommand-9hs6
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: The syde CLI shall expose every syde feature as a cobra subcommand under a single root binary.
req_type: functional
priority: must
verification: inspection of internal/cli/root.go command tree
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: A unified command tree is the sole user interface for the syde model.
---
