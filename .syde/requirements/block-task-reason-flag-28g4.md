---
id: REQ-0259
kind: requirement
name: Block Task Reason Flag
slug: block-task-reason-flag-28g4
relationships:
    - target: block-task-egd4
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:02Z"
statement: When syde task block is invoked, the syde CLI shall accept --reason as an optional string stored in the task notes.
req_type: interface
priority: must
verification: integration test invoking syde task block --reason
source: manual
source_ref: contract:block-task-egd4
requirement_status: active
rationale: Capturing the block reason preserves context for future reviewers.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:02Z"
---
