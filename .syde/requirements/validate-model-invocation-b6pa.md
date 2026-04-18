---
id: REQ-0322
kind: requirement
name: Validate Model Invocation
slug: validate-model-invocation-b6pa
relationships:
    - target: validate-model-tjzs
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:01Z"
statement: When the user runs syde validate, the syde CLI shall print errors and warnings grouped by severity and exit with code 0 on success or 1 on errors.
req_type: interface
priority: must
verification: integration test invoking syde validate against valid and invalid models
source: manual
source_ref: contract:validate-model-tjzs
requirement_status: active
rationale: Validation is the final gate before finishing any design session.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:01Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:01Z"
---
