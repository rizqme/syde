---
id: REQ-0306
kind: requirement
name: Start Task Invocation
slug: start-task-invocation-gdej
relationships:
    - target: start-task-wa36
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:23Z"
statement: When the user runs syde task start <slug>, the syde CLI shall transition the task's task_status field to in_progress.
req_type: interface
priority: must
verification: integration test invoking syde task start
source: manual
source_ref: contract:start-task-wa36
requirement_status: active
rationale: Start transitions signal active work and gate subsequent completion operations.
audited_overlaps:
    - slug: block-task-invocation-t3z7
      distinction: syde task start transitions task_status to in_progress; syde task block transitions it to blocked, different commands and target states.
    - slug: complete-task-invocation-dm0b
      distinction: syde task start moves task_status to in_progress; syde task done moves it to completed, distinct commands and end states.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:23Z"
---
