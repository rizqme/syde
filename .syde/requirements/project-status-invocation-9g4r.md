---
id: REQ-0288
kind: requirement
name: Project Status Invocation
slug: project-status-invocation-9g4r
relationships:
    - target: project-status-zo3d
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde status, the syde CLI shall print entity counts grouped by kind on stdout.
req_type: interface
priority: must
verification: integration test invoking syde status
source: manual
source_ref: contract:project-status-zo3d
requirement_status: active
rationale: Status is the shortest health-check command available to operators.
---
