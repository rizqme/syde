---
contract_kind: cli
description: Sync/audit the design model against the codebase with optional gap check.
id: CON-0069
input: syde sync [--dry-run --coverage --check]
input_parameters:
    - description: report without writing
      path: --dry-run
      type: bool
    - description: file-to-component coverage audit
      path: --coverage
      type: bool
    - description: exit non-zero on completeness gaps
      path: --check
      type: bool
interaction_pattern: request-response
kind: contract
name: Sync Codebase
output: sync/audit report on stdout
output_parameters:
    - description: missing components, files, relationships
      path: gaps
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: sync-codebase-b4jw
updated_at: "2026-04-14T03:27:06Z"
---
