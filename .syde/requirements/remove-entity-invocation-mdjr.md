---
id: REQ-0294
kind: requirement
name: Remove Entity Invocation
slug: remove-entity-invocation-mdjr
relationships:
    - target: remove-entity-t21l
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:48Z"
statement: When the user runs syde remove <slug>, the syde CLI shall delete the named entity and print a removal confirmation line.
req_type: interface
priority: must
verification: integration test invoking syde remove
source: manual
source_ref: contract:remove-entity-t21l
requirement_status: active
rationale: Remove is the canonical deletion command and must provide clear confirmation.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:48Z"
---
