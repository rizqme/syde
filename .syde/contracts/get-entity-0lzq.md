---
contract_kind: cli
description: Show full entity details for a slug, ID, or parent/child path.
id: CON-0009
input: syde get <slug>
input_parameters:
    - description: positional, required. Full slug, bare slug, or parent/child path
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Get Entity
output: human-readable entity dump on stdout
output_parameters:
    - description: YAML frontmatter + markdown body rendered with ANSI colors
      path: body
      type: string
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
slug: get-entity-0lzq
updated_at: "2026-04-14T03:27:03Z"
---
