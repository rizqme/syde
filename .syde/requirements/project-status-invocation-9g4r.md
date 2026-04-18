---
id: REQ-0288
kind: requirement
name: Project Status Invocation
slug: project-status-invocation-9g4r
relationships:
    - target: project-status-zo3d
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:26Z"
statement: When the user runs syde status, the syde CLI shall print entity counts grouped by kind on stdout.
req_type: interface
priority: must
verification: integration test invoking syde status
source: manual
source_ref: contract:project-status-zo3d
requirement_status: active
rationale: Status is the shortest health-check command available to operators.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:26Z"
---
