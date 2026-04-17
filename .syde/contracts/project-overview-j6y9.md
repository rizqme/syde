---
id: CON-0053
kind: contract
name: Project Overview
slug: project-overview-j6y9
description: GET /api/<project>/overview — entity counts and recent file changes.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/overview
input_parameters:
    - path: project
      type: string
      description: path parameter, project identifier
output: 200 OK application/json
output_parameters:
    - path: name
      type: string
      description: project name
    - path: entity_counts
      type: map<string,int>
      description: entities per kind
    - path: recent_changes
      type: '[]object'
      description: recent file changes from git log
---
