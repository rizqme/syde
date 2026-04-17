---
id: CON-0071
kind: contract
name: File Tree Screen
slug: file-tree-screen-fez2
description: Source-tree browser with summaries on every node and inline file content
files:
    - web/src/pages/FileTree.tsx
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: screen
interaction_pattern: render
input: /__tree__
input_parameters:
    - path: path
      type: string
      description: optional preselected node path
output: rendered file tree page
output_parameters:
    - path: node-click
      type: event
      description: select tree node
wireframe: <screen name="File Tree" direction="horizontal"><sidebar name="Tree" width="280"><heading>Source tree</heading><list><item><label>internal/</label></item><item><label>web/</label></item><item active="true"><label>scripts/</label></item><item><label>skill/</label></item></list></sidebar><main name="File"><heading>Selected file</heading><section title="Summary"><text/></section><section title="Content"><placeholder/></section></main></screen>
---
