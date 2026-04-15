---
contract_kind: rest
description: GET /api/<project>/search — full-text search for the dashboard palette.
id: CON-0056
input: GET /api/<project>/search?q=<query>
input_parameters:
    - description: path parameter
      path: project
      type: string
    - description: query, required. Search text
      path: q
      type: string
interaction_pattern: request-response
kind: contract
name: Search Entities (HTTP)
output: 200 OK application/json ranked results
output_parameters:
    - description: matching entities with score + snippet
      path: hits
      type: '[]object'
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: search-entities-http-1th3
updated_at: "2026-04-14T03:27:05Z"
---
