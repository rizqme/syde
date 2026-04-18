---
id: REQ-0275
kind: requirement
name: Ignore Tree Node Invocation
slug: ignore-tree-node-invocation-0kg2
relationships:
    - target: ignore-tree-node-m460
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:34Z"
statement: When the user runs syde tree ignore <path>, the syde CLI shall flag the named node as ignored so the orphan validator exempts it.
req_type: interface
priority: must
verification: integration test invoking syde tree ignore and running the orphan validator
source: manual
source_ref: contract:ignore-tree-node-m460
requirement_status: active
rationale: Ignore lists let operators silence intentional gaps without breaking validation.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:34Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:34Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:34Z"
---
