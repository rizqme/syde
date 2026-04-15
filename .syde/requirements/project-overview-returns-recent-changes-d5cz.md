---
id: REQ-0211
kind: requirement
name: Project Overview Returns Recent Changes
slug: project-overview-returns-recent-changes-d5cz
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:20Z"
statement: When GET /api/<project>/overview succeeds, the syded daemon shall return recent_changes as an array of recent file change objects sourced from the git log.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The overview page surfaces a git-backed activity feed.
---
