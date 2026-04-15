---
id: REQ-0259
kind: requirement
name: Block Task Reason Flag
slug: block-task-reason-flag-28g4
relationships:
    - target: block-task-egd4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When syde task block is invoked, the syde CLI shall accept --reason as an optional string stored in the task notes.
req_type: interface
priority: must
verification: integration test invoking syde task block --reason
source: manual
source_ref: contract:block-task-egd4
requirement_status: active
rationale: Capturing the block reason preserves context for future reviewers.
---
