---
id: REQ-0311
kind: requirement
name: Tree Context Bundle Invocation
slug: tree-context-bundle-invocation-iqrq
relationships:
    - target: tree-context-bundle-3co6
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde tree context <path>, the syde CLI shall return the ancestor breadcrumb, node summary, and file content or children listing for the named node.
req_type: interface
priority: must
verification: integration test invoking syde tree context on a file and a folder
source: manual
source_ref: contract:tree-context-bundle-3co6
requirement_status: active
rationale: Bundled context is the framing used when creating entities against an existing codebase.
---
