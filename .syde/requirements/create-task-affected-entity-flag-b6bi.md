---
id: REQ-0267
kind: requirement
name: Create Task Affected Entity Flag
slug: create-task-affected-entity-flag-b6bi
relationships:
    - target: create-task-23f4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When syde task create is invoked, the syde CLI shall accept --affected-entity as a repeatable string whose values are validated against existing entity slugs.
req_type: interface
priority: must
verification: integration test invoking syde task create --affected-entity with valid and invalid slugs
source: manual
source_ref: contract:create-task-23f4
requirement_status: active
rationale: Tracking affected entities enables coverage audits and drift detection.
---
