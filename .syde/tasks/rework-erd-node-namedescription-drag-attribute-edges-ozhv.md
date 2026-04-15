---
acceptance: 'Navigating to Concepts > ERD: nodes show concept name + description, attribute rows have no type badge, nodes can be dragged, canvas pans and zooms. A test concept with attribute ''parent|string|parent slug|summary-tree'' renders an edge from the parent attribute row to the Summary Tree node.'
affected_entities:
    - web-spa
    - entity-model
affected_files:
    - web/src/pages/ERD.tsx
    - web/src/lib/api.ts
completed_at: "2026-04-14T10:28:44Z"
created_at: "2026-04-14T10:22:14Z"
details: 'web/src/pages/ERD.tsx ConceptNode: header shows concept name + truncated meaning. Attribute rows render name bold + description muted (no type badge). Each attribute with refs.length > 0 gets a <Handle type=''source'' position={Position.Right} id={''attr-''+attr.name}> positioned in the row. Edges builder iterates every concept.attributes and emits one edge per ref (id unique, source=slug, sourceHandle=attr-<name>, target=targetSlug, label=attr.name, distinct color). Relates_to edges still render with cardinality labels but colored differently. ReactFlow props explicitly set nodesDraggable, panOnDrag, zoomOnScroll, nodesConnectable={false}, elementsSelectable. Import ErdIcon removal handled in phase 2.'
id: TSK-0068
kind: task
name: 'Rework ERD node: name+description, drag, attribute edges'
objective: ERD canvas is pannable/draggable, nodes show name + description (attribute descriptions, not types), and attributes with refs draw edges from per-row handles to the target concepts
plan_phase: phase_6
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rework-erd-node-namedescription-drag-attribute-edges-ozhv
task_status: completed
updated_at: "2026-04-14T10:28:44Z"
---
