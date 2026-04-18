---
id: TSK-0235
kind: task
name: Generate migration triage worksheet
slug: generate-migration-triage-worksheet-icp8
updated_at: '2026-04-18T08:01:18Z'
task_status: completed
priority: high
objective: Produce a structured worksheet listing all 207 active reqs lacking refines:component, with heuristic-based component suggestion(s) and confidence per row
details: 'Walk every active requirement. Skip those already carrying refines:component. For each remaining: collect (a) components in the system named by belongs_to:system, (b) components whose name (or any word in name) appears as a substring of the requirement statement, (c) any component file paths mentioned in statement/rationale. Emit JSON or markdown table with columns: req_slug, statement_excerpt, current_belongs_to, suggested_components (ranked), confidence (high/medium/low).'
acceptance: /tmp/syde-triage-worksheet.md or .json exists with one row per req-needing-migration; confidence labels assigned.
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_1
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:01:18Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: active-requirement-shall-refine-at-least-one-component-mke4
---
