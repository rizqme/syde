---
id: REQ-0263
kind: requirement
name: Constraints Check File Invocation
slug: constraints-check-file-invocation-59ps
relationships:
    - target: constraints-for-file-ld34
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:51Z"
statement: When the user runs syde constraints check <file>, the syde CLI shall return the applicable component slug and architecture decisions for that file.
req_type: interface
priority: must
verification: integration test invoking syde constraints check against a mapped source file
source: manual
source_ref: contract:constraints-for-file-ld34
requirement_status: active
rationale: File-scoped constraint lookup lets authors verify edits against architectural decisions.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:51Z"
---
