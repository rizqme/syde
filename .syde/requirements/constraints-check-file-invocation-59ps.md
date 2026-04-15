---
id: REQ-0263
kind: requirement
name: Constraints Check File Invocation
slug: constraints-check-file-invocation-59ps
relationships:
    - target: constraints-for-file-ld34
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:20Z"
statement: When the user runs syde constraints check <file>, the syde CLI shall return the applicable component slug and architecture decisions for that file.
req_type: interface
priority: must
verification: integration test invoking syde constraints check against a mapped source file
source: manual
source_ref: contract:constraints-for-file-ld34
requirement_status: active
rationale: File-scoped constraint lookup lets authors verify edits against architectural decisions.
---
