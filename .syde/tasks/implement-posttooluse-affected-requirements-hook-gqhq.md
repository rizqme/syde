---
id: TSK-0251
kind: task
name: Implement PostToolUse affected-requirements hook
slug: implement-posttooluse-affected-requirements-hook-gqhq
updated_at: '2026-04-18T08:20:18Z'
task_status: completed
priority: medium
objective: After Edit/Write touches a path mapped to a component's files, the hook emits the affected active requirements as a context block for the next agent turn
details: 'Modify or create skill/hooks.go (or hooks.json + handler). PostToolUse handler receives tool input (path). Calls daemon /api/<proj>/constraints-check?path=<p> to find owning component, then /api/<proj>/query?refined-by=<comp-slug>. Builds context block: ''Affected requirements (re-verify with syde requirement verify <slug>): ...''. No-op when no component owns the path.'
acceptance: Editing internal/audit/audit.go in a session causes the next prompt context to include the audit-engine-4ktg refining requirements
affected_entities:
- skill-installer-wbmu
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_5
created_at: '2026-04-18T08:00:26Z'
completed_at: '2026-04-18T08:20:18Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: posttooluse-hook-shall-surface-affected-requirements-when-a-file-mapped-to-a-component-is-edited-ofeu
---
