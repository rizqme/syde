---
category: gotcha
confidence: medium
description: 'React Flow v12 controlled mode (<ReactFlow nodes={x} edges={y} />) REQUIRES matching onNodesChange + onEdgesChange handlers. Without them, drag/select/remove events fire but never apply — node positions revert instantly and it looks like dragging is disabled. Fix: use useNodesState/useEdgesState hooks (which return [state, setter, onChange]) and pass the onChange prop. If the source data changes (e.g. concepts refetched), useEffect setNodes(initialNodes) to sync. This is the second-most common React Flow v12 footgun after the button-as-nodrag rule.'
discovered_at: "2026-04-14T10:57:41Z"
entity_refs:
    - web-spa
id: LRN-0021
kind: learning
name: React Flow v12 controlled mode (<ReactFlow nodes={x} edges={
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: react-flow-v12-controlled-mode-reactflow-nodesx-edges-kunm
source: session-observation
updated_at: "2026-04-14T10:57:41Z"
---
