---
id: REQ-0294
kind: requirement
name: Remove Entity Invocation
slug: remove-entity-invocation-mdjr
relationships:
    - target: remove-entity-t21l
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde remove <slug>, the syde CLI shall delete the named entity and print a removal confirmation line.
req_type: interface
priority: must
verification: integration test invoking syde remove
source: manual
source_ref: contract:remove-entity-t21l
requirement_status: active
rationale: Remove is the canonical deletion command and must provide clear confirmation.
---
