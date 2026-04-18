---
id: REQ-0118
kind: requirement
name: Task Done Bumps Affected Entities
slug: task-done-bumps-affected-entities-gfxu
relationships:
    - target: task-d3oc
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:59Z"
statement: When a task transitions to completed, the syde CLI shall update the updated_at timestamp on every one of its affected entities.
req_type: functional
priority: must
verification: integration test running syde task done and asserting updated_at changes
source: manual
source_ref: concept:task-d3oc
requirement_status: active
rationale: Auto-bump keeps staleness tracking honest when work touches design nodes.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:59Z"
---
