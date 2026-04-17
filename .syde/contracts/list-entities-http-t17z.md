---
id: CON-0054
kind: contract
name: List Entities (HTTP)
slug: list-entities-http-t17z
description: GET /api/<project>/entities — list entities for the dashboard.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/entities?kind=<kind>&tag=<tag>
input_parameters:
    - path: project
      type: string
      description: path parameter
    - path: kind
      type: string
      description: query, optional kind filter
    - path: tag
      type: string
      description: query, optional tag filter
output: 200 OK application/json list of entities
output_parameters:
    - path: entities
      type: '[]object'
      description: entity summaries
    - path: entities[].id
      type: string
      description: entity ID
    - path: entities[].slug
      type: string
      description: file slug
    - path: entities[].kind
      type: string
      description: entity kind
    - path: entities[].name
      type: string
      description: display name
---
