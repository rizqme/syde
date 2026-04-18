---
id: REQ-0211
kind: requirement
name: Project Overview Returns Recent Changes
slug: project-overview-returns-recent-changes-d5cz
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:57Z"
statement: When GET /api/<project>/overview succeeds, the syded daemon shall return recent_changes as an array of recent file change objects sourced from the git log.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The overview page surfaces a git-backed activity feed.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:57Z"
---
