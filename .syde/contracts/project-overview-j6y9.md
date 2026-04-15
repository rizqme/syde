---
contract_kind: rest
description: GET /api/<project>/overview — entity counts and recent file changes.
id: CON-0053
input: GET /api/<project>/overview
input_parameters:
    - description: path parameter, project identifier
      path: project
      type: string
interaction_pattern: request-response
kind: contract
name: Project Overview
output: 200 OK application/json
output_parameters:
    - description: project name
      path: name
      type: string
    - description: entities per kind
      path: entity_counts
      type: map<string,int>
    - description: recent file changes from git log
      path: recent_changes
      type: '[]object'
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: project-overview-j6y9
updated_at: "2026-04-14T03:27:05Z"
---
