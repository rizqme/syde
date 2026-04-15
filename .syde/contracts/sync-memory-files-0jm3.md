---
contract_kind: cli
description: Write per-entity Claude memory files under .claude/.
id: CON-0065
input: syde memory sync
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Sync Memory Files
output: Writes per-entity memory files under .claude/
output_parameters:
    - description: number of memory files written
      path: written
      type: int
relationships:
    - target: syde-cli
      type: belongs_to
    - target: memory-sync
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: sync-memory-files-0jm3
updated_at: "2026-04-14T03:27:06Z"
---
