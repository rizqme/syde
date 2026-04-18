---
id: REQ-0261
kind: requirement
name: Complete Task Auto-Bumps Touched Entities
slug: complete-task-auto-bumps-touched-entities-0vhb
relationships:
    - target: complete-task-k8je
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:31Z"
statement: When syde task done succeeds, the syde CLI shall return touched_entities as a list of entity slugs whose updated_at was bumped.
req_type: interface
priority: must
verification: integration test invoking syde task done and inspecting touched_entities output
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Clearing drift warnings depends on accurate reporting of which entities were touched.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:31Z"
---
