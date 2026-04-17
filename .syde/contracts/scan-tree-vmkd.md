---
id: CON-0043
kind: contract
name: Scan Tree
slug: scan-tree-vmkd
description: Walk the source tree, hash files, and cascade stale to ancestors.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree scan
input_parameters:
    - path: _
      type: '-'
      description: no arguments
output: Diff summary on stdout; updates .syde/tree.yaml in place
output_parameters:
    - path: added
      type: int
      description: newly tracked paths
    - path: changed
      type: int
      description: paths whose hash changed
    - path: deleted
      type: int
      description: paths removed from disk
    - path: stale
      type: int
      description: total stale node count after cascade
---
