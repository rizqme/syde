---
acceptance: ./syde query --code ConceptEntity prints a rich list with owner per hit. ./syde query --file internal/cli/query.go --content prints owners + related + file content.
affected_entities:
    - cli-commands
    - http-api
affected_files:
    - internal/cli/query.go
    - internal/dashboard/api_readall.go
    - internal/query/formatter.go
completed_at: "2026-04-14T09:26:57Z"
created_at: "2026-04-14T09:20:18Z"
details: 'internal/cli/query.go: queryCode flag (string), queryContent flag (bool). --code routes mode=code with extra params q + limit. --file branch passes content=true when --content. Update Long help. internal/dashboard/api_readall.go handleQueryAPI: case ''code'' parses q + limit and calls eng.SearchCode; case ''by-file'' reads content=true|false and threads to ByFile. Add FormatCodeHits in formatter.go for rich/refs output.'
id: TSK-0053
kind: task
name: CLI + dashboard wiring for --code and --content
objective: Both new modes reachable via syde CLI and the dashboard /query endpoint
plan_phase: phase_1
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: cli-dashboard-wiring-for-code-and-content-fz9h
task_status: completed
updated_at: "2026-04-14T09:26:57Z"
---
