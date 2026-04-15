---
contract_kind: screen
description: Source-tree browser with summaries on every node and inline file content
files:
    - web/src/pages/FileTree.tsx
id: CON-0071
input: /__tree__
input_parameters:
    - description: optional preselected node path
      path: path
      type: string
interaction_pattern: render
kind: contract
name: File Tree Screen
output: rendered file tree page
output_parameters:
    - description: select tree node
      path: node-click
      type: event
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: file-tree-screen-fez2
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="File Tree" direction="horizontal"><sidebar name="Tree" width="280"><heading>Source tree</heading><list><item><label>internal/</label></item><item><label>web/</label></item><item active="true"><label>scripts/</label></item><item><label>skill/</label></item></list></sidebar><main name="File"><heading>Selected file</heading><section title="Summary"><text/></section><section title="Content"><placeholder/></section></main></screen>
---
