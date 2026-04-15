---
id: REQ-0312
kind: requirement
name: Tree Context Max Bytes Flag
slug: tree-context-max-bytes-flag-q2zt
relationships:
    - target: tree-context-bundle-3co6
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When syde tree context is invoked with --max-bytes, the syde CLI shall truncate inlined file content to at most the specified byte count.
req_type: interface
priority: must
verification: integration test invoking syde tree context --max-bytes on a large file
source: manual
source_ref: contract:tree-context-bundle-3co6
requirement_status: active
rationale: Byte capping prevents the command from blowing up agent context windows.
---
