---
id: CON-0009
kind: contract
name: Get Entity
slug: get-entity-0lzq
description: Show full entity details for a slug, ID, or parent/child path.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde get <slug>
input_parameters:
    - path: slug
      type: string
      description: positional, required. Full slug, bare slug, or parent/child path
output: human-readable entity dump on stdout
output_parameters:
    - path: body
      type: string
      description: YAML frontmatter + markdown body rendered with ANSI colors
---
