---
id: REQ-0284
kind: requirement
name: List Tasks Status Filter
slug: list-tasks-status-filter-gwnr
relationships:
    - target: list-tasks-0sgj
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:13Z"
statement: When syde task list is invoked with --status, the syde CLI shall filter the listing to tasks matching one of pending, in_progress, completed, blocked, or cancelled.
req_type: interface
priority: must
verification: integration test invoking syde task list --status
source: manual
source_ref: contract:list-tasks-0sgj
requirement_status: active
rationale: Status filtering lets operators focus on actionable or stuck work.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:13Z"
---
