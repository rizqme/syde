---
contract_kind: rest
description: GET /api/<project>/tree — flat tree node map for the FileTree page.
id: CON-0057
input: GET /api/<project>/tree
input_parameters:
    - description: path parameter
      path: project
      type: string
interaction_pattern: request-response
kind: contract
name: Get Full Tree (HTTP)
output: 200 OK application/json flat node map
output_parameters:
    - description: path-keyed tree node map
      path: nodes
      type: map<string,object>
    - description: ISO8601 timestamp
      path: scanned_at
      type: string
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: get-full-tree-http-c0vt
updated_at: "2026-04-14T03:27:06Z"
---
