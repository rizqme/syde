---
id: REQ-0312
kind: requirement
name: Tree Context Max Bytes Flag
slug: tree-context-max-bytes-flag-q2zt
relationships:
    - target: tree-context-bundle-3co6
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:22Z"
statement: When syde tree context is invoked with --max-bytes, the syde CLI shall truncate inlined file content to at most the specified byte count.
req_type: interface
priority: must
verification: integration test invoking syde tree context --max-bytes on a large file
source: manual
source_ref: contract:tree-context-bundle-3co6
requirement_status: active
rationale: Byte capping prevents the command from blowing up agent context windows.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:22Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:22Z"
---
