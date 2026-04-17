---
id: CON-0052
kind: contract
name: List Projects
slug: list-projects-eilx
description: GET /api/projects — list dashboard-registered projects.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/projects
input_parameters:
    - path: _
      type: '-'
      description: no request parameters
output: 200 OK application/json with project list
output_parameters:
    - path: projects
      type: '[]object'
      description: list of registered project entries
    - path: projects[].name
      type: string
      description: project name
    - path: projects[].path
      type: string
      description: absolute filesystem path
---
