---
id: CON-0049
kind: contract
name: Tree Status
slug: tree-status-k6ag
description: Show tree counts and last scan time, with --strict gate for hooks.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree status [--strict]
input_parameters:
    - path: --strict
      type: bool
      description: exit code 1 if anything is stale (for CI and session-end hooks)
output: Counts line + optional stale listing
output_parameters:
    - path: files
      type: int
      description: total tracked files
    - path: dirs
      type: int
      description: total tracked dirs
    - path: stale_files
      type: int
      description: stale file count
    - path: stale_dirs
      type: int
      description: stale dir count
    - path: last_scan
      type: string
      description: ISO8601 timestamp of last scan
---
