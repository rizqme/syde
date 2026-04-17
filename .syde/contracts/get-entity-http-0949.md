---
id: CON-0055
kind: contract
name: Get Entity (HTTP)
slug: get-entity-http-0949
description: GET /api/<project>/entity/<slug> — full entity payload for the detail panel.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/entity/<slug>
input_parameters:
    - path: project
      type: string
      description: path parameter
    - path: slug
      type: string
      description: path parameter; full or bare slug
output: 200 OK application/json full entity; 404 if not found
output_parameters:
    - path: entity
      type: object
      description: full entity including relationships, files, body
---
