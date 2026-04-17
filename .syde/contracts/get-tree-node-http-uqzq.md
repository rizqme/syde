---
id: CON-0058
kind: contract
name: Get Tree Node (HTTP)
slug: get-tree-node-http-uqzq
description: GET /api/<project>/tree/<path> — context bundle for one tree node.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/tree/<path>
input_parameters:
    - path: project
      type: string
      description: path parameter
    - path: path
      type: string
      description: path parameter; relative source path
output: 200 OK application/json ContextBundle
output_parameters:
    - path: breadcrumb
      type: '[]object'
      description: ancestor folder summaries
    - path: summary
      type: string
      description: node summary
    - path: content
      type: string
      description: raw file content (files only)
    - path: children
      type: '[]object'
      description: direct children (folders only)
---
