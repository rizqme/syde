---
id: REQ-0313
kind: requirement
name: Tree Status Invocation
slug: tree-status-invocation-udek
relationships:
    - target: tree-status-k6ag
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree status, the syde CLI shall print counts of total files, total dirs, stale files, stale dirs, and the last scan timestamp.
req_type: interface
priority: must
verification: integration test invoking syde tree status
source: manual
source_ref: contract:tree-status-k6ag
requirement_status: active
rationale: Tree status is the quick-check command that tells operators whether a scan is needed.
---
