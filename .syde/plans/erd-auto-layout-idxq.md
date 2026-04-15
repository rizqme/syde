---
approved_at: "2026-04-14T10:59:21Z"
background: The ERD canvas renders concept nodes in a naive 5-column grid, which means edges cross heavily and the canvas looks like a tangle of overlapping lines. For a small project with ~11 concepts the result is already unreadable — there's no layer structure, aggregate cardinality edges and attribute-level FK refs criss-cross through node centers, and users can't trace which attribute points where.
completed_at: "2026-04-14T11:01:07Z"
created_at: "2026-04-14T10:58:52Z"
id: PLN-0010
kind: plan
name: erd-auto-layout
objective: Nodes are placed by a layered graph layout algorithm so edges route cleanly along the direction of flow with minimal crossings. Edges render as smoothstep (right-angle) arrows instead of bezier curves, and per-attribute handles are routed to the appropriate side of the node for the chosen direction.
phases:
    - changes: web/package.json adds @dagrejs/dagre. web/src/pages/ERD.tsx imports dagre, adds getLayoutedElements that builds a dagre.graphlib.Graph with rankdir=LR, nodesep, ranksep, sets node sizes, runs dagre.layout, and returns positioned nodes with targetPosition=Left + sourcePosition=Right. ConceptNode rewritten with Handle type=target on left and source on right; top/bottom aggregate handles removed; per-attribute handles stay on right. All edges set to type='smoothstep' for right-angle routing.
      description: Install dagre, compute layered positions for concept nodes, switch edges to smoothstep, use left/right handles for LR flow
      details: 'dagre config: rankdir LR, nodesep 80, ranksep 240 (attribute-ref edges need horizontal runway), marginx 40, marginy 40. Node dims: w=300, h depends on attribute count (280 + 40*attrCount approximate). Layout runs in a useMemo derived from initialNodes + initialEdges before handing to useNodesState. CSS: keep the cursor-grab on nodes; no other style changes needed.'
      id: phase_1
      name: Dagre layered layout + smoothstep edges
      notes: Aggregate relates_to edges route card-to-card via default left/right handles. Attribute refs edges still use sourceHandle='attr-<name>' anchored to the per-attribute handle on the right side.
      objective: ERD canvas renders concepts in clean layered columns with non-crossing right-angle edges
      status: completed
      tasks:
        - dagre-layout-lr-handles-smoothstep-edges
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) add @dagrejs/dagre as a web dependency; (2) write a getLayoutedElements helper that runs dagre layered layout on the concept graph; (3) update ConceptNode to use left (target) + right (source) handles for left-right flow, drop unused top/bottom aggregate handles; (4) switch edge type to smoothstep for clean right-angle routing; (5) apply the layout on mount and whenever concepts change. Out-of-scope: interactive layout controls (direction toggle, nodesep slider), persisting positions across reloads, alternative layout engines (elk, cola).'
slug: erd-auto-layout-idxq
source: manual
updated_at: "2026-04-14T11:01:07Z"
---
