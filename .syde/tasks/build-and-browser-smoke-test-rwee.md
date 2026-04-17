---
id: TSK-0071
kind: task
name: Build and browser smoke test
slug: build-and-browser-smoke-test-rwee
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: go build and bun run build succeed; dashboard renders a flow with flowchart
details: go build ./cmd/syde/ && go build ./cmd/syded/ && cd web && bun run build. Open dashboard, navigate to a flow, verify flowchart renders.
acceptance: Both builds green; flowchart visible in browser
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_5
created_at: "2026-04-16T09:23:43Z"
completed_at: "2026-04-16T10:56:36Z"
---
