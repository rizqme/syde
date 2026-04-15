---
acceptance: curl each endpoint against live syded returns valid JSON; covers all read cmds
affected_entities:
    - http-api
affected_files:
    - internal/dashboard/api.go
completed_at: "2026-04-14T06:44:16Z"
created_at: "2026-04-14T06:38:16Z"
details: 'Add: GET list, get, query, status, validate, sync-check (+strict), context, constraints-check (+path), health. Reuse query.Engine; each handler builds the same payload the CLI currently prints.'
id: TSK-0015
kind: task
name: Add read endpoints covering every CLI command family
objective: Every syde read command has a matching API endpoint returning JSON
plan_phase: phase_1
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-read-endpoints-covering-every-cli-command-family-n4u9
task_status: completed
updated_at: "2026-04-14T06:44:16Z"
---
