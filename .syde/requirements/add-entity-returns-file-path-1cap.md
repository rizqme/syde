---
id: REQ-0245
kind: requirement
name: Add Entity Returns File Path
slug: add-entity-returns-file-path-1cap
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:45Z"
statement: When syde add succeeds, the syde CLI shall return the absolute file path of the newly created markdown entity.
req_type: interface
priority: must
verification: integration test invoking syde add and asserting the printed path exists
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: Downstream tooling and the operator need the exact location of the created entity.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:45Z"
---
