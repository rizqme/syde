---
id: CON-0088
kind: contract
name: Plans Inbox Screen
slug: plans-inbox-screen-72gg
description: 2-column dashboard layout with plan list on the left and plan detail on the right
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
      type: references
updated_at: "2026-04-17T10:45:59Z"
contract_kind: screen
interaction_pattern: render
input: /syde-<proj>/plan
input_parameters:
    - path: filter
      type: string
      description: optional filter DSL query passed via ?filter=
output: rendered plans list and detail panel
output_parameters:
    - path: click
      type: event
      description: select a plan to view its details
wireframe: <screen name="Plans Inbox" direction="horizontal"><panel name="List" width="360"><heading>Plans</heading><list><item active="true"><label>Plan one</label><label>approved</label></item><item><label>Plan two</label><label>completed</label></item></list></panel><main name="Detail"><heading>Plan one</heading><tabs><tab active="true">Plan</tab><tab>Tasks</tab></tabs><section title="Background"><text/></section><section title="Objective"><text/></section><section title="Scope"><text/></section><section title="Design"><text/></section><section title="Changes"><text/></section></main></screen>
---
