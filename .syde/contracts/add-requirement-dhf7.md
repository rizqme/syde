---
id: CON-0083
kind: contract
name: Add Requirement
slug: add-requirement-dhf7
description: CLI invocation that authors a new requirement entity with EARS-validated statement, overlap detection, and mandatory acknowledgement for surfaced duplicates.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
      type: references
updated_at: "2026-04-17T10:45:59Z"
contract_kind: cli
interaction_pattern: request-response
input: syde add requirement <name> [--statement --type --priority --verification --source --rationale --audited slug:reason --force]
input_parameters:
    - path: name
      type: string
      description: positional, required. Human-readable requirement title
    - path: --statement
      type: string
      description: EARS shall-form text (required)
    - path: --type
      type: string
      description: functional/non-functional/constraint/interface/performance/security/usability (required)
    - path: --priority
      type: string
      description: must/should/could/wont (required)
    - path: --verification
      type: string
      description: how the requirement is verified (required)
    - path: --source
      type: string
      description: user/plan/migration/manual
    - path: --rationale
      type: string
      description: why this requirement exists
    - path: --audited
      type: array<string>
      description: slug[:distinction] (repeatable) acknowledging TF-IDF overlap pairs
    - path: --force
      type: bool
      description: bypass the overlap gate (rare)
output: exit 0 on success; exit non-zero when unacknowledged overlaps exist and --force not passed
output_parameters:
    - path: id
      type: string
      description: generated REQ-NNNN ID
    - path: file
      type: string
      description: path to the new .md file
    - path: overlaps
      type: array<string>
      description: overlap candidates printed to stdout when similarity > 0.6
---
