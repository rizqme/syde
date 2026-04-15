---
acceptance: 'ERD view: concepts are laid out in columns, edges route right-to-left via smoothstep right-angle paths with minimal crossings, drag still works.'
affected_entities:
    - web-spa
affected_files:
    - web/src/pages/ERD.tsx
    - web/package.json
completed_at: "2026-04-14T11:01:07Z"
created_at: "2026-04-14T10:59:16Z"
details: 'cd web && bun add @dagrejs/dagre. web/src/pages/ERD.tsx: import dagre from ''@dagrejs/dagre''. Write getLayoutedElements(nodes, edges): build dagre graph with rankdir=LR nodesep=80 ranksep=240, setNode per node with width=300 height=280+40*attr, setEdge per edge, dagre.layout, return nodes mapped to {...n, position: {x: pos.x-150, y: pos.y-(h/2)}, targetPosition: Position.Left, sourcePosition: Position.Right}. Apply getLayoutedElements to initialNodes+initialEdges before useNodesState. ConceptNode: delete top Handle, delete bottom Handle, add <Handle type=target position=Position.Left /> and <Handle type=source position=Position.Right /> on the card root. Per-attribute handles stay as-is on Position.Right. Edges: add type: ''smoothstep'' to every edge object.'
id: TSK-0073
kind: task
name: Dagre layout + LR handles + smoothstep edges
objective: ERD canvas computes dagre LR positions on mount, ConceptNode uses left/right handles, edges are smoothstep
plan_phase: phase_1
plan_ref: erd-auto-layout
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dagre-layout-lr-handles-smoothstep-edges-a22c
task_status: completed
updated_at: "2026-04-14T11:01:07Z"
---
