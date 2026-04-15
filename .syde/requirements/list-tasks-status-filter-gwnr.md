---
id: REQ-0284
kind: requirement
name: List Tasks Status Filter
slug: list-tasks-status-filter-gwnr
relationships:
    - target: list-tasks-0sgj
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When syde task list is invoked with --status, the syde CLI shall filter the listing to tasks matching one of pending, in_progress, completed, blocked, or cancelled.
req_type: interface
priority: must
verification: integration test invoking syde task list --status
source: manual
source_ref: contract:list-tasks-0sgj
requirement_status: active
rationale: Status filtering lets operators focus on actionable or stuck work.
---
