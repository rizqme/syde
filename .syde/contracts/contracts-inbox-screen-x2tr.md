---
contract_kind: screen
description: 2-column inbox listing contracts with kind+pattern filters and detail panel
id: CON-0079
input: /contract
input_parameters:
    - description: 'optional filter DSL query (supports contract_kind: and pattern:)'
      path: filter
      type: string
interaction_pattern: render
kind: contract
name: Contracts Inbox Screen
output: rendered contracts list and detail
output_parameters:
    - description: select contract in detail panel
      path: contract-click
      type: event
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: contracts-inbox-screen-x2tr
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Contracts Inbox" direction="horizontal"><sidebar name="Kinds" width="200"><menu><item>Systems</item><item>Components</item><item>Contracts</item><item>Concepts</item><item>Flows</item><item>Decisions</item></menu></sidebar><panel name="List" width="360"><heading>Contracts</heading><button-group><button>Filter</button><button>Sort</button></button-group><list><item active="true"><image/><label>Item one</label><label>meta</label></item><item><image/><label>Item two</label><label>meta</label></item><item><image/><label>Item three</label><label>meta</label></item></list></panel><main name="Detail"><heading>Selected entity</heading><text/><section title="Files"><text/></section><section title="Relationships"><list><item><image/><label>Related entity</label><label>belongs_to</label></item></list></section><button-group><button>Edit</button><button>Delete</button></button-group></main></screen>
---
