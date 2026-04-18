---
id: REQ-0267
kind: requirement
name: Create Task Affected Entity Flag
slug: create-task-affected-entity-flag-b6bi
relationships:
    - target: create-task-23f4
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:42Z"
statement: When syde task create is invoked, the syde CLI shall accept --affected-entity as a repeatable string whose values are validated against existing entity slugs.
req_type: interface
priority: must
verification: integration test invoking syde task create --affected-entity with valid and invalid slugs
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Tracking affected entities enables coverage audits and drift detection.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:42Z"
---
