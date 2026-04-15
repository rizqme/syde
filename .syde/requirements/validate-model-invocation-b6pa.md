---
id: REQ-0322
kind: requirement
name: Validate Model Invocation
slug: validate-model-invocation-b6pa
relationships:
    - target: validate-model-tjzs
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde validate, the syde CLI shall print errors and warnings grouped by severity and exit with code 0 on success or 1 on errors.
req_type: interface
priority: must
verification: integration test invoking syde validate against valid and invalid models
source: manual
source_ref: contract:validate-model-tjzs
requirement_status: active
rationale: Validation is the final gate before finishing any design session.
---
