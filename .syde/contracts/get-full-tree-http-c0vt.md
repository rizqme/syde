---
id: CON-0057
kind: contract
name: Get Full Tree (HTTP)
slug: get-full-tree-http-c0vt
description: GET /api/<project>/tree — flat tree node map for the FileTree page.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: http-api
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: rest
interaction_pattern: request-response
input: GET /api/<project>/tree
input_parameters:
    - path: project
      type: string
      description: path parameter
output: 200 OK application/json flat node map
output_parameters:
    - path: nodes
      type: map<string,object>
      description: path-keyed tree node map
    - path: scanned_at
      type: string
      description: ISO8601 timestamp
---
