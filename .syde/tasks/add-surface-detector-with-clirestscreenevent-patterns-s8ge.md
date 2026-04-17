---
id: TSK-0117
kind: task
name: Add surface detector with CLI/REST/screen/event patterns
slug: add-surface-detector-with-clirestscreenevent-patterns-s8ge
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: internal/audit/surfaces.go extracts contract-surface mentions from any requirement statement with passing unit tests
details: Implement SurfaceKind enum (CLI/REST/Screen/Event), type Surface{Kind, Raw, Normalised}, function ExtractSurfaces(statement string) []Surface using four regexes. Unit tests in surfaces_test.go cover each pattern positively and negatively, plus a requirement-corpus round-trip against a handful of representative active reqs.
acceptance: go test ./internal/audit/... passes; ExtractSurfaces returns expected Surface slice for at least two reqs per kind
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/surfaces.go
    - internal/audit/surfaces_test.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_8
created_at: "2026-04-17T09:46:36Z"
completed_at: "2026-04-17T10:20:47Z"
---
