---
id: CON-0024
kind: contract
name: Open Dashboard
slug: open-dashboard-cupm
description: Start syded in the background and open the project page in a browser.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde open
input_parameters:
- path: _
  type: '-'
  description: no arguments
output: Starts syded in the background and opens the browser to the project page
output_parameters:
- path: url
  type: string
  description: opened URL
---
