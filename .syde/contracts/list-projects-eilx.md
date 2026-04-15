---
contract_kind: rest
description: GET /api/projects — list dashboard-registered projects.
id: CON-0052
input: GET /api/projects
input_parameters:
    - description: no request parameters
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: List Projects
output: 200 OK application/json with project list
output_parameters:
    - description: list of registered project entries
      path: projects
      type: '[]object'
    - description: project name
      path: projects[].name
      type: string
    - description: absolute filesystem path
      path: projects[].path
      type: string
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: list-projects-eilx
updated_at: "2026-04-14T03:27:05Z"
---
