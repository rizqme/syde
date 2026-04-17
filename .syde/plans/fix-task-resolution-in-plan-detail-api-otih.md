---
id: PLN-0008
kind: plan
name: Fix task resolution in plan detail API
slug: fix-task-resolution-in-plan-detail-api-otih
relationships:
    - target: approved-plan-fix-task-resolution-in-plan-detail-api-pmu9
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T01:38:15Z"
plan_status: completed
background: handlePlanDetail resolves tasks by exact slug match but phases store bare slugs from utils.Slugify(name). Tasks with common names like 'Build and verify' get suffix-disambiguated slugs (build-and-verify-1s89) that don't match the bare slug stored in the phase.
objective: Plan detail API resolves tasks by bare slug fallback when exact match fails.
scope: internal/dashboard/api.go handlePlanDetail only.
design: When resolving phase task slugs, try store.Get(slug) first. If not found, iterate task entities and match via utils.BaseSlug. Same pattern as the phase auto-completion fix.
source: manual
created_at: "2026-04-16T11:49:56Z"
approved_at: "2026-04-16T11:49:56Z"
completed_at: "2026-04-16T11:52:05Z"
phases:
    - id: phase_1
      name: Fix
      status: completed
      description: Add BaseSlug fallback to task resolution
      objective: Plan detail API finds tasks regardless of bare vs full slug
      changes: internal/dashboard/api.go
      details: In the task resolution loop, when store.Get fails, try matching by BaseSlug
      tasks:
        - add-baseslug-fallback-to-handleplandetail-task-resolution
---
