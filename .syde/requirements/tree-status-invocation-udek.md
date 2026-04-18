---
id: REQ-0313
kind: requirement
name: Tree Status Invocation
slug: tree-status-invocation-udek
relationships:
    - target: tree-status-k6ag
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:39Z"
statement: When the user runs syde tree status, the syde CLI shall print counts of total files, total dirs, stale files, stale dirs, and the last scan timestamp.
req_type: interface
priority: must
verification: integration test invoking syde tree status
source: manual
source_ref: contract:tree-status-k6ag
requirement_status: active
rationale: Tree status is the quick-check command that tells operators whether a scan is needed.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:39Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:39Z"
---
