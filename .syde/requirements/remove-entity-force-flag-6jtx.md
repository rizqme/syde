---
id: REQ-0295
kind: requirement
name: Remove Entity Force Flag
slug: remove-entity-force-flag-6jtx
relationships:
    - target: remove-entity-t21l
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: Where --force is passed to syde remove, the syde CLI shall skip the interactive confirmation prompt before deletion.
req_type: interface
priority: must
verification: integration test invoking syde remove --force
source: manual
source_ref: contract:remove-entity-t21l
requirement_status: active
rationale: Force mode supports scripted removal in automation contexts.
---
