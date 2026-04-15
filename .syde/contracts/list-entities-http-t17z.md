---
contract_kind: rest
description: GET /api/<project>/entities — list entities for the dashboard.
id: CON-0054
input: GET /api/<project>/entities?kind=<kind>&tag=<tag>
input_parameters:
    - description: path parameter
      path: project
      type: string
    - description: query, optional kind filter
      path: kind
      type: string
    - description: query, optional tag filter
      path: tag
      type: string
interaction_pattern: request-response
kind: contract
name: List Entities (HTTP)
output: 200 OK application/json list of entities
output_parameters:
    - description: entity summaries
      path: entities
      type: '[]object'
    - description: entity ID
      path: entities[].id
      type: string
    - description: file slug
      path: entities[].slug
      type: string
    - description: entity kind
      path: entities[].kind
      type: string
    - description: display name
      path: entities[].name
      type: string
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: list-entities-http-t17z
updated_at: "2026-04-14T03:27:05Z"
---
