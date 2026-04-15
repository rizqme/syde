---
contract_kind: cli
description: Walk the source tree, hash files, and cascade stale to ancestors.
id: CON-0043
input: syde tree scan
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Scan Tree
output: Diff summary on stdout; updates .syde/tree.yaml in place
output_parameters:
    - description: newly tracked paths
      path: added
      type: int
    - description: paths whose hash changed
      path: changed
      type: int
    - description: paths removed from disk
      path: deleted
      type: int
    - description: total stale node count after cascade
      path: stale
      type: int
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: scan-tree-vmkd
updated_at: "2026-04-14T03:27:05Z"
---
