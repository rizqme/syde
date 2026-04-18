---
id: TSK-0246
kind: task
name: Implement ResolveRefinedBy in query engine
slug: implement-resolverefinedby-in-query-engine-a7ai
updated_at: '2026-04-18T08:16:53Z'
task_status: completed
priority: medium
objective: Query engine method that returns active reqs refining a given component, used by both CLI and skill hook
details: Resolve componentSlug via existing canonicalization. Walk all reqs (kind=requirement, status=active), filter those whose Relationships contain a refines edge whose resolved Target equals the component's canonical slug.
acceptance: 'Unit test or smoke: ResolveRefinedBy(''audit-engine-4ktg'') returns >0 reqs after Phase 1 migration'
affected_files:
- internal/query/engine.go
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_4
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:16:53Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: syde-query-shall-support-refined-by-component-slug-o23d
---
