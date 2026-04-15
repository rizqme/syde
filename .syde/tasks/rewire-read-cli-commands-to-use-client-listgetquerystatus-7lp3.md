---
acceptance: syde list/get/query/status/context all work end-to-end via HTTP
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/list.go
    - internal/cli/get.go
    - internal/cli/query.go
    - internal/cli/search.go
    - internal/cli/status.go
    - internal/cli/context.go
completed_at: "2026-04-14T06:51:10Z"
created_at: "2026-04-14T06:38:16Z"
details: Replace openStore() + direct index calls with client.New() + client.<Method>(). Output formatting stays in CLI. Apply to list, get, query, search, status, context.
id: TSK-0022
kind: task
name: Rewire read CLI commands to use client (list/get/query/status)
objective: Pure-read commands stop calling openStore(); use client instead
plan_phase: phase_4
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-read-cli-commands-to-use-client-listgetquerystatus-7lp3
task_status: completed
updated_at: "2026-04-14T06:51:10Z"
---
