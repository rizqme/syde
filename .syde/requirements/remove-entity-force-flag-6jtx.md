---
id: REQ-0295
kind: requirement
name: Remove Entity Force Flag
slug: remove-entity-force-flag-6jtx
relationships:
    - target: remove-entity-t21l
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:41Z"
statement: Where --force is passed to syde remove, the syde CLI shall skip the interactive confirmation prompt before deletion.
req_type: interface
priority: must
verification: integration test invoking syde remove --force
source: manual
source_ref: contract:remove-entity-t21l
requirement_status: active
rationale: Force mode supports scripted removal in automation contexts.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:41Z"
---
