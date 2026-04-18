---
id: REQ-0311
kind: requirement
name: Tree Context Bundle Invocation
slug: tree-context-bundle-invocation-iqrq
relationships:
    - target: tree-context-bundle-3co6
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:59Z"
statement: When the user runs syde tree context <path>, the syde CLI shall return the ancestor breadcrumb, node summary, and file content or children listing for the named node.
req_type: interface
priority: must
verification: integration test invoking syde tree context on a file and a folder
source: manual
source_ref: contract:tree-context-bundle-3co6
requirement_status: active
rationale: Bundled context is the framing used when creating entities against an existing codebase.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:59Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:59Z"
---
