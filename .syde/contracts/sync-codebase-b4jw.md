---
id: CON-0069
kind: contract
name: Sync Codebase
slug: sync-codebase-b4jw
description: Sync/audit the design model against the codebase with optional gap check.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde sync [--dry-run --coverage --check]
input_parameters:
    - path: --dry-run
      type: bool
      description: report without writing
    - path: --coverage
      type: bool
      description: file-to-component coverage audit
    - path: --check
      type: bool
      description: exit non-zero on completeness gaps
output: sync/audit report on stdout
output_parameters:
    - path: gaps
      type: '[]string'
      description: missing components, files, relationships
---
