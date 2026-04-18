---
id: COM-0014
kind: component
name: HTTP API
slug: http-api-afos
description: REST handlers serving syde project data to the dashboard SPA.
purpose: Expose syde project data to the dashboard SPA via REST
notes:
- handleAPI reuses GetStore instead of NewStore/defer Close per request.
files:
- internal/dashboard/api.go
- internal/dashboard/api_readall.go
- internal/dashboard/api_write.go
- internal/dashboard/run.go
- internal/dashboard/html.go
relationships:
- target: syde-5tdt
  type: belongs_to
- target: syde
  type: belongs_to
updated_at: '2026-04-18T08:47:43Z'
responsibility: Handle HTTP routes for entities, tree, search, constraints, context
capabilities:
- Serve GET /api/<project>/entities and per-kind listings
- Serve GET /api/<project>/tree and /tree/<path> (context bundle)
- Serve search, constraints-check, and project metadata endpoints
- Serve the embedded React SPA at the root path
boundaries: Does NOT mutate entities (dashboard is read-only for now). Does NOT own project discovery (Project Registry does).
---
