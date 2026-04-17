---
id: CON-0056
kind: contract
name: Search Entities (HTTP)
slug: search-entities-http-1th3
description: GET /api/<project>/search — full-text search for the dashboard palette.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/search?q=<query>
input_parameters:
    - path: project
      type: string
      description: path parameter
    - path: q
      type: string
      description: query, required. Search text
output: 200 OK application/json ranked results
output_parameters:
    - path: hits
      type: '[]object'
      description: matching entities with score + snippet
---
