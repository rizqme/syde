---
id: CON-0080
kind: contract
name: Concepts Inbox Screen
slug: concepts-inbox-screen-bpow
description: 2-column inbox listing concepts with detail panel (List mode of the Concepts page)
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: screen
interaction_pattern: render
input: /concept (List toggle active)
input_parameters:
    - path: view
      type: string
      description: fixed value list
output: rendered concepts list and detail
output_parameters:
    - path: concept-click
      type: event
      description: select concept in detail panel
wireframe: <screen name="Concepts Inbox" direction="horizontal"><sidebar name="Kinds" width="200"><menu><item>Systems</item><item>Components</item><item>Contracts</item><item>Concepts</item><item>Flows</item><item>Decisions</item></menu></sidebar><panel name="List" width="360"><heading>Concepts</heading><button-group><button>Filter</button><button>Sort</button></button-group><list><item active="true"><image/><label>Item one</label><label>meta</label></item><item><image/><label>Item two</label><label>meta</label></item><item><image/><label>Item three</label><label>meta</label></item></list></panel><main name="Detail"><heading>Selected entity</heading><text/><section title="Files"><text/></section><section title="Relationships"><list><item><image/><label>Related entity</label><label>belongs_to</label></item></list></section><button-group><button>Edit</button><button>Delete</button></button-group></main></screen>
---
