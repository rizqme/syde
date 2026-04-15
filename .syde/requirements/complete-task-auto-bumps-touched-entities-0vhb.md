---
id: REQ-0261
kind: requirement
name: Complete Task Auto-Bumps Touched Entities
slug: complete-task-auto-bumps-touched-entities-0vhb
relationships:
    - target: complete-task-k8je
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:19Z"
statement: When syde task done succeeds, the syde CLI shall return touched_entities as a list of entity slugs whose updated_at was bumped.
req_type: interface
priority: must
verification: integration test invoking syde task done and inspecting touched_entities output
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Clearing drift warnings depends on accurate reporting of which entities were touched.
---
