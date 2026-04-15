---
id: REQ-0286
kind: requirement
name: List Tree Changes Leaves Only Flag
slug: list-tree-changes-leaves-only-flag-3dma
relationships:
    - target: list-tree-changes-mqhv
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: Where --leaves-only is passed to syde tree changes, the syde CLI shall hide stale folders whose descendants are still stale.
req_type: interface
priority: must
verification: integration test invoking syde tree changes --leaves-only
source: manual
source_ref: contract:list-tree-changes-mqhv
requirement_status: active
rationale: The leaves-only mode drives the standard summarize loop in the sync workflow.
---
