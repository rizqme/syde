---
id: REQ-0260
kind: requirement
name: Complete Task Invocation
slug: complete-task-invocation-dm0b
relationships:
    - target: complete-task-k8je
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:56Z"
statement: When the user runs syde task done <slug>, the syde CLI shall transition the task's task_status field to completed.
req_type: interface
priority: must
verification: integration test invoking syde task done
source: manual
source_ref: contract:complete-task-k8je
requirement_status: active
rationale: Task completion is the canonical progress signal for plans.
audited_overlaps:
    - slug: block-task-invocation-t3z7
      distinction: 'Different command and end state: done transitions task_status to completed via syde task done, while block transitions it to blocked.'
    - slug: start-task-invocation-gdej
      distinction: 'Different command and end state: done transitions task_status to completed, while start transitions it to in_progress via syde task start.'
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:56Z"
---
