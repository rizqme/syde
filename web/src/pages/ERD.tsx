import { useEffect, useMemo } from 'react';
import {
  ReactFlow,
  ReactFlowProvider,
  Background,
  Controls,
  Handle,
  Position,
  MarkerType,
  useNodesState,
  useEdgesState,
  type Node,
  type Edge,
  type NodeProps,
} from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import dagre from '@dagrejs/dagre';
import { useApi } from '../hooks/useApi';
import { api, EntitiesResponse, EntitySummary } from '../lib/api';

// ERD page — renders every concept entity as a React Flow node with
// a dagre-computed layered layout so edges route cleanly along the
// left→right flow of the graph. Two edge kinds:
//
//  - relates_to edges (concept → concept) carry a cardinality label
//    like "one-to-many" and connect the right of a source card to
//    the left of the target card.
//  - attribute refs edges (attribute row → concept) render a dashed
//    arrow from a specific attribute row of a concept to the card
//    of the concept it references (FK-style), labelled with the
//    attribute name.
//
// All edges are smoothstep (right-angle) so they look like ERD lines
// rather than bezier spaghetti.

interface ERDProps {
  onSelectEntity: (slug: string, kind?: string) => void;
}

interface ConceptAttribute {
  name: string;
  description?: string;
  refs?: string[];
}

interface ConceptAction {
  name: string;
  description?: string;
}

interface ConceptNodeData extends Record<string, unknown> {
  name: string;
  description?: string;
  slug: string;
  attributes: ConceptAttribute[];
  actions: ConceptAction[];
  onSelect: () => void;
}

// Node size estimates fed to dagre so it can allocate enough space
// per layer. Height grows with attribute count so tall concept cards
// don't overlap their neighbours.
const NODE_WIDTH = 300;
const NODE_HEADER_HEIGHT = 70;
const NODE_ATTR_ROW_HEIGHT = 42;
const NODE_ACTION_FOOTER = 44;

function estimateNodeHeight(c: EntitySummary): number {
  const attrs = (c.attributes || []).length;
  const hasActions = (c.actions || []).length > 0;
  return NODE_HEADER_HEIGHT + attrs * NODE_ATTR_ROW_HEIGHT + (hasActions ? NODE_ACTION_FOOTER : 0);
}

// getLayoutedElements runs a dagre layered graph layout on the
// concept graph and returns nodes with computed positions + the
// left/right handle positions that React Flow needs for smooth
// edge routing. Direction is left-to-right because the attribute
// refs edges always exit from the right side of an attribute row,
// which aligns naturally with LR flow.
function getLayoutedElements(
  nodes: Node<ConceptNodeData>[],
  edges: Edge[],
  concepts: EntitySummary[],
) {
  const g = new dagre.graphlib.Graph();
  g.setDefaultEdgeLabel(() => ({}));
  g.setGraph({
    rankdir: 'LR',
    nodesep: 60,
    ranksep: 220,
    marginx: 40,
    marginy: 40,
  });

  const heightBySlug: Record<string, number> = {};
  for (const c of concepts) {
    heightBySlug[c.slug] = estimateNodeHeight(c);
  }

  for (const n of nodes) {
    g.setNode(n.id, {
      width: NODE_WIDTH,
      height: heightBySlug[n.id] || NODE_HEADER_HEIGHT,
    });
  }
  for (const e of edges) {
    g.setEdge(e.source, e.target);
  }

  dagre.layout(g);

  const laidOut = nodes.map((n): Node<ConceptNodeData> => {
    const pos = g.node(n.id);
    const h = heightBySlug[n.id] || NODE_HEADER_HEIGHT;
    return {
      ...n,
      // Dagre positions are node centers; React Flow expects the
      // top-left corner, so offset by half width / height.
      position: { x: pos.x - NODE_WIDTH / 2, y: pos.y - h / 2 },
      targetPosition: Position.Left,
      sourcePosition: Position.Right,
    };
  });

  return { nodes: laidOut, edges };
}

function ConceptNode({ data }: NodeProps<Node<ConceptNodeData>>) {
  // Left = inbound, right = outbound. Attribute refs also exit from
  // the right side via per-row handles so they stay consistent with
  // the LR flow direction.
  return (
    <div
      className="rounded-lg border border-border bg-card shadow-md overflow-hidden text-xs cursor-grab active:cursor-grabbing"
      style={{ width: NODE_WIDTH }}
    >
      <Handle
        type="target"
        position={Position.Left}
        className="!bg-kind-concept !w-2 !h-2 !border-0"
        style={{ pointerEvents: 'none' }}
      />
      <div className="px-3 py-2 bg-muted/40 border-b border-border">
        <div
          role="button"
          tabIndex={0}
          onClick={data.onSelect}
          onKeyDown={(e) => {
            if (e.key === 'Enter' || e.key === ' ') {
              e.preventDefault();
              data.onSelect();
            }
          }}
          className="nodrag text-sm font-semibold text-foreground hover:underline cursor-pointer truncate"
          title={data.name}
        >
          {data.name}
        </div>
        {data.description && (
          <div className="text-[10px] text-muted-foreground leading-snug mt-0.5 line-clamp-2">
            {data.description}
          </div>
        )}
      </div>
      {data.attributes.length > 0 && (
        <div className="divide-y divide-border/60">
          {data.attributes.map((a, i) => (
            <div key={i} className="relative px-3 py-1.5">
              <div className="font-mono text-foreground truncate">{a.name}</div>
              {a.description && (
                <div className="text-[10px] text-muted-foreground leading-snug line-clamp-2">
                  {a.description}
                </div>
              )}
              {a.refs && a.refs.length > 0 && (
                <Handle
                  type="source"
                  position={Position.Right}
                  id={`attr-${a.name}`}
                  className="!bg-kind-concept !w-2 !h-2 !border-0"
                  style={{ top: '50%', pointerEvents: 'none' }}
                />
              )}
            </div>
          ))}
        </div>
      )}
      {data.actions.length > 0 && (
        <div className="border-t border-border/60 px-3 py-1.5 bg-muted/20">
          <div className="text-[9px] uppercase tracking-wider text-muted-foreground mb-1">actions</div>
          <div className="flex flex-wrap gap-1">
            {data.actions.map((a, i) => (
              <span
                key={i}
                className="font-mono text-[10px] px-1.5 py-0.5 rounded bg-muted text-foreground"
                title={a.description}
              >
                {a.name}()
              </span>
            ))}
          </div>
        </div>
      )}
      <Handle
        type="source"
        position={Position.Right}
        className="!bg-kind-concept !w-2 !h-2 !border-0"
        style={{ pointerEvents: 'none' }}
      />
    </div>
  );
}

const nodeTypes = { concept: ConceptNode };

function ERDCanvas({ concepts, onSelectEntity }: { concepts: EntitySummary[]; onSelectEntity: ERDProps['onSelectEntity'] }) {
  // Slug→node-id map so rel targets can resolve by any slug form.
  const slugToId = useMemo(() => {
    const map: Record<string, string> = {};
    for (const c of concepts) {
      map[c.slug] = c.slug;
      const base = c.slug.replace(/-[a-z0-9]{4}$/, '');
      if (base && base !== c.slug && !(base in map)) {
        map[base] = c.slug;
      }
    }
    return map;
  }, [concepts]);

  // Seed nodes and edges from the concepts prop. Dagre lays them
  // out; then useNodesState owns the state so drag events mutate
  // positions in place.
  const { seedNodes, seedEdges } = useMemo(() => {
    const rawNodes: Node<ConceptNodeData>[] = concepts.map((c) => ({
      id: c.slug,
      type: 'concept',
      position: { x: 0, y: 0 },
      data: {
        name: c.name,
        description: c.description,
        slug: c.slug,
        attributes: (c.attributes || []) as ConceptAttribute[],
        actions: (c.actions || []) as ConceptAction[],
        onSelect: () => onSelectEntity(c.slug, 'concept'),
      },
    }));
    const rawEdges: Edge[] = [];
    // Aggregate relates_to edges between concept cards.
    for (const c of concepts) {
      for (const rel of c.relationships || []) {
        if (rel.type !== 'relates_to') continue;
        const targetId = slugToId[rel.target];
        if (!targetId) continue;
        rawEdges.push({
          id: `rel:${c.slug}->${targetId}:${rel.label || ''}`,
          source: c.slug,
          target: targetId,
          label: rel.label || '',
          type: 'smoothstep',
          labelStyle: { fontSize: 10, fill: '#d4d4d8' },
          labelBgStyle: { fill: '#18181b', fillOpacity: 0.85 },
          style: { stroke: '#a78bfa', strokeWidth: 2 },
          markerEnd: { type: MarkerType.ArrowClosed, color: '#a78bfa' },
        });
      }
    }
    // Attribute refs edges from per-attribute handles to target cards.
    for (const c of concepts) {
      for (const a of (c.attributes as ConceptAttribute[]) || []) {
        if (!a.refs || a.refs.length === 0) continue;
        for (const ref of a.refs) {
          const targetId = slugToId[ref];
          if (!targetId) continue;
          rawEdges.push({
            id: `attr:${c.slug}.${a.name}->${targetId}`,
            source: c.slug,
            sourceHandle: `attr-${a.name}`,
            target: targetId,
            label: a.name,
            type: 'smoothstep',
            labelStyle: { fontSize: 9, fill: '#a1a1aa' },
            labelBgStyle: { fill: '#18181b', fillOpacity: 0.85 },
            style: { stroke: '#6b7280', strokeWidth: 1, strokeDasharray: '4 3' },
            markerEnd: { type: MarkerType.ArrowClosed, color: '#6b7280' },
          });
        }
      }
    }
    const laid = getLayoutedElements(rawNodes, rawEdges, concepts);
    return { seedNodes: laid.nodes, seedEdges: laid.edges };
  }, [concepts, slugToId, onSelectEntity]);

  const [nodes, setNodes, onNodesChange] = useNodesState(seedNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(seedEdges);

  useEffect(() => {
    setNodes(seedNodes);
  }, [seedNodes, setNodes]);
  useEffect(() => {
    setEdges(seedEdges);
  }, [seedEdges, setEdges]);

  return (
    <ReactFlow
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      nodeTypes={nodeTypes}
      fitView
      fitViewOptions={{ padding: 0.2 }}
      proOptions={{ hideAttribution: true }}
      nodesDraggable
      nodesConnectable={false}
      elementsSelectable
      panOnDrag
      zoomOnScroll
      zoomOnPinch
      minZoom={0.2}
      maxZoom={2}
    >
      <Background gap={24} color="#27272a" />
      <Controls showInteractive={false} />
    </ReactFlow>
  );
}

export function ERD({ onSelectEntity }: ERDProps) {
  const { data, loading, error } = useApi<EntitiesResponse>(() => api.entities('concept'), []);

  if (loading) {
    return <div className="p-8 text-muted-foreground">Loading concepts…</div>;
  }
  if (error) {
    return <div className="p-8 text-red-400">Failed to load concepts: {String(error)}</div>;
  }
  const concepts = data?.entities || [];
  if (concepts.length === 0) {
    return (
      <div className="p-8 text-muted-foreground">
        No concepts yet. Create one with
        <code className="mx-1 font-mono text-xs px-1.5 py-0.5 rounded bg-muted">
          syde add concept &lt;name&gt;
        </code>
        and it will appear here.
      </div>
    );
  }
  return (
    <div className="h-full w-full">
      <ReactFlowProvider>
        <ERDCanvas concepts={concepts} onSelectEntity={onSelectEntity} />
      </ReactFlowProvider>
    </div>
  );
}
