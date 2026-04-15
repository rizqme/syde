---
contract_kind: cli
description: Start syded in the background and open the project page in a browser.
id: CON-0024
input: syde open
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Open Dashboard
output: Starts syded in the background and opens the browser to the project page
output_parameters:
    - description: opened URL
      path: url
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: open-dashboard-cupm
updated_at: "2026-04-14T03:27:04Z"
---
