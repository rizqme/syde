---
contract_kind: rest
description: GET /api/<project>/tree/<path> — context bundle for one tree node.
id: CON-0058
input: GET /api/<project>/tree/<path>
input_parameters:
    - description: path parameter
      path: project
      type: string
    - description: path parameter; relative source path
      path: path
      type: string
interaction_pattern: request-response
kind: contract
name: Get Tree Node (HTTP)
output: 200 OK application/json ContextBundle
output_parameters:
    - description: ancestor folder summaries
      path: breadcrumb
      type: '[]object'
    - description: node summary
      path: summary
      type: string
    - description: raw file content (files only)
      path: content
      type: string
    - description: direct children (folders only)
      path: children
      type: '[]object'
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: get-tree-node-http-uqzq
updated_at: "2026-04-14T03:27:06Z"
---
