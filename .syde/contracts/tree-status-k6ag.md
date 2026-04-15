---
contract_kind: cli
description: Show tree counts and last scan time, with --strict gate for hooks.
id: CON-0049
input: syde tree status [--strict]
input_parameters:
    - description: exit code 1 if anything is stale (for CI and session-end hooks)
      path: --strict
      type: bool
interaction_pattern: request-response
kind: contract
name: Tree Status
output: Counts line + optional stale listing
output_parameters:
    - description: total tracked files
      path: files
      type: int
    - description: total tracked dirs
      path: dirs
      type: int
    - description: stale file count
      path: stale_files
      type: int
    - description: stale dir count
      path: stale_dirs
      type: int
    - description: ISO8601 timestamp of last scan
      path: last_scan
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: tree-status-k6ag
updated_at: "2026-04-14T03:27:05Z"
---
