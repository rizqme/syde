---
id: REQ-0132
kind: requirement
name: Ignored Tree Node Exempt From Gates
slug: ignored-tree-node-exempt-from-gates-afho
relationships:
    - target: tree-node-iutv
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:47Z"
statement: While a tree node has ignored set to true, the syde CLI shall exclude it from orphan and stale-tree status gates.
req_type: functional
priority: must
verification: integration test marking a node ignored and running syde tree status --strict
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: Ignored nodes are intentional exclusions that must not block sync gates.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:47Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:47Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:47Z"
---
