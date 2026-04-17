---
id: CON-0079
kind: contract
name: Contracts Inbox Screen
slug: contracts-inbox-screen-x2tr
description: 2-column inbox listing contracts with kind+pattern filters and detail panel
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: screen
interaction_pattern: render
input: /contract
input_parameters:
    - path: filter
      type: string
      description: 'optional filter DSL query (supports contract_kind: and pattern:)'
output: rendered contracts list and detail
output_parameters:
    - path: contract-click
      type: event
      description: select contract in detail panel
wireframe: <screen name="Contracts Inbox" direction="horizontal"><sidebar name="Kinds" width="200"><menu><item>Systems</item><item>Components</item><item>Contracts</item><item>Concepts</item><item>Flows</item><item>Decisions</item></menu></sidebar><panel name="List" width="360"><heading>Contracts</heading><button-group><button>Filter</button><button>Sort</button></button-group><list><item active="true"><image/><label>Item one</label><label>meta</label></item><item><image/><label>Item two</label><label>meta</label></item><item><image/><label>Item three</label><label>meta</label></item></list></panel><main name="Detail"><heading>Selected entity</heading><text/><section title="Files"><text/></section><section title="Relationships"><list><item><image/><label>Related entity</label><label>belongs_to</label></item></list></section><button-group><button>Edit</button><button>Delete</button></button-group></main></screen>
---
