---
id: REQ-0274
kind: requirement
name: Get Tree Summary Invocation
slug: get-tree-summary-invocation-1j26
relationships:
    - target: get-tree-summary-2vyd
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde tree get <path>, the syde CLI shall return the stored summary text for the named tree node.
req_type: interface
priority: must
verification: integration test invoking syde tree get on a summarized path
source: manual
source_ref: contract:get-tree-summary-2vyd
requirement_status: active
rationale: Targeted summary retrieval lets agents inspect single nodes without rendering the full tree.
---
