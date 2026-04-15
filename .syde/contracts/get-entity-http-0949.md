---
contract_kind: rest
description: GET /api/<project>/entity/<slug> — full entity payload for the detail panel.
id: CON-0055
input: GET /api/<project>/entity/<slug>
input_parameters:
    - description: path parameter
      path: project
      type: string
    - description: path parameter; full or bare slug
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Get Entity (HTTP)
output: 200 OK application/json full entity; 404 if not found
output_parameters:
    - description: full entity including relationships, files, body
      path: entity
      type: object
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: get-entity-http-0949
updated_at: "2026-04-14T03:27:05Z"
---
