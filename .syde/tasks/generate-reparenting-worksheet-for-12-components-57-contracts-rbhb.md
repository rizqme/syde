---
id: TSK-0256
kind: task
name: Generate reparenting worksheet for 12 components + 57 contracts
slug: generate-reparenting-worksheet-for-12-components-57-contracts-rbhb
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: component-shall-be-allowed-to-belong-to-multiple-systems-qd6u
      type: implements
updated_at: "2026-04-18T10:00:03Z"
task_status: completed
priority: high
objective: Produce a triage worksheet listing every component/contract currently belongs_to syde-cli-2478 with a heuristic-based proposed assignment to syde-5tdt and/or syded-dashboard-e82c.
details: 'Components: static import trace from cmd/syde/main.go and cmd/syded/main.go to component.files. Contracts: classify by input signature (syde prefix, /api/ path, ws:// prefix) and contract_kind. Confidence: high if heuristic gives a single answer, medium if shared, low if ambiguous.'
acceptance: /tmp/syde-system-reparent-worksheet.json exists with one row per entity; every row has 1+ proposed targets.
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_3
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T10:00:03Z"
---
