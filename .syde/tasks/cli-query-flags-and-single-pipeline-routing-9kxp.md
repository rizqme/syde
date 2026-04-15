---
acceptance: ./syde query --search badger --kind component prints rich list with snippets; ./syde query --file internal/storage/index.go prints owners + related.
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/query.go
    - internal/client/client.go
completed_at: "2026-04-14T08:24:42Z"
created_at: "2026-04-14T08:07:48Z"
details: Add queryFile, queryLimit, queryAny, queryNoRelated flag vars. --search branch builds extra url.Values (kind, tag, limit, any) and routes mode=search. --file branch routes mode=by-file with path + with_related. Drop the old c.Search() call. When --full + --search, take the top hit and do a follow-up mode=full lookup.
id: TSK-0046
kind: task
name: CLI query flags and single pipeline routing
objective: syde query supports --search, --file, --limit, --any, --no-related; no c.Search() shortcut
plan_phase: phase_3
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: cli-query-flags-and-single-pipeline-routing-9kxp
task_status: completed
updated_at: "2026-04-14T08:24:42Z"
---
