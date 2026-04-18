---
id: REQ-0258
kind: requirement
name: Block Task Invocation
slug: block-task-invocation-t3z7
relationships:
    - target: block-task-egd4
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:02Z"
statement: When the user runs syde task block <slug>, the syde CLI shall transition the task's task_status field to blocked.
req_type: interface
priority: must
verification: integration test invoking syde task block and checking task_status
source: manual
source_ref: contract:block-task-egd4
requirement_status: active
rationale: Blocked state communicates that work cannot proceed without unblocking.
audited_overlaps:
    - slug: complete-task-invocation-dm0b
      distinction: 'Different command and end state: block transitions task_status to blocked via syde task block, while done transitions it to completed via syde task done.'
    - slug: start-task-invocation-gdej
      distinction: 'Different command and end state: block transitions task_status to blocked, while start transitions it to in_progress via syde task start.'
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:02Z"
---
