---
id: REQ-0272
kind: requirement
name: Get Entity Invocation
slug: get-entity-invocation-il5p
relationships:
    - target: get-entity-0lzq
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde get <slug>, the syde CLI shall render the entity's YAML frontmatter and markdown body on stdout.
req_type: interface
priority: must
verification: integration test invoking syde get on an existing slug
source: manual
source_ref: contract:get-entity-0lzq
requirement_status: active
rationale: syde get is the canonical accessor used by agents and humans alike.
---
