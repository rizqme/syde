import { useMemo } from 'react';

interface FlowStep {
  id: string;
  action: string;
  contract?: string;
  description?: string;
  on_success?: string;
  on_failure?: string;
}

interface FlowChartProps {
  steps: FlowStep[];
  onContractClick?: (slug: string) => void;
}

interface LayoutNode {
  step: FlowStep;
  x: number;
  y: number;
  col: number;
  row: number;
}

interface Edge {
  from: string;
  to: string;
  type: 'success' | 'failure';
}

const NODE_W = 260;
const NODE_H = 72;
const GAP_X = 40;
const GAP_Y = 32;
const PAD = 24;

export function FlowChart({ steps, onContractClick }: FlowChartProps) {
  const { nodes, edges, width, height } = useMemo(() => layoutGraph(steps), [steps]);

  if (steps.length === 0) {
    return (
      <div className="text-sm text-muted-foreground italic py-4">
        No steps defined yet.
      </div>
    );
  }

  return (
    <div className="overflow-x-auto">
      <svg
        width={width}
        height={height}
        className="block"
        style={{ minWidth: width }}
      >
        <defs>
          <marker id="arrow-success" markerWidth="8" markerHeight="6" refX="8" refY="3" orient="auto">
            <path d="M0,0 L8,3 L0,6" fill="#22c55e" />
          </marker>
          <marker id="arrow-failure" markerWidth="8" markerHeight="6" refX="8" refY="3" orient="auto">
            <path d="M0,0 L8,3 L0,6" fill="#ef4444" />
          </marker>
          <marker id="arrow-default" markerWidth="8" markerHeight="6" refX="8" refY="3" orient="auto">
            <path d="M0,0 L8,3 L0,6" fill="#52525b" />
          </marker>
        </defs>

        {/* Edges */}
        {edges.map((edge, i) => {
          const from = nodes.find((n) => n.step.id === edge.from);
          const to = nodes.find((n) => n.step.id === edge.to);
          if (!from || !to) return null;

          const x1 = from.x + NODE_W / 2;
          const y1 = from.y + NODE_H;
          const x2 = to.x + NODE_W / 2;
          const y2 = to.y;

          const isSuccess = edge.type === 'success';
          const color = isSuccess ? '#22c55e' : '#ef4444';
          const markerId = isSuccess ? 'arrow-success' : 'arrow-failure';
          const dashArray = isSuccess ? undefined : '6,4';

          // Curved path for non-straight connections
          if (Math.abs(x1 - x2) < 2) {
            // Straight vertical
            return (
              <line
                key={i}
                x1={x1} y1={y1}
                x2={x2} y2={y2 - 2}
                stroke={color}
                strokeWidth={1.5}
                strokeDasharray={dashArray}
                markerEnd={`url(#${markerId})`}
              />
            );
          }

          // Curved path
          const midY = (y1 + y2) / 2;
          return (
            <path
              key={i}
              d={`M${x1},${y1} C${x1},${midY} ${x2},${midY} ${x2},${y2 - 2}`}
              fill="none"
              stroke={color}
              strokeWidth={1.5}
              strokeDasharray={dashArray}
              markerEnd={`url(#${markerId})`}
            />
          );
        })}

        {/* Terminal edges (to "done"/"abort") */}
        {nodes.map((node) => {
          const terminalTargets = [
            { val: node.step.on_success, type: 'success' as const },
            { val: node.step.on_failure, type: 'failure' as const },
          ].filter((t) => t.val === 'done' || t.val === 'abort');

          return terminalTargets.map((t, i) => {
            const x = node.x + NODE_W / 2 + (t.type === 'failure' ? 40 : 0);
            const y = node.y + NODE_H;
            const color = t.type === 'success' ? '#22c55e' : '#ef4444';
            const dashArray = t.type === 'success' ? undefined : '6,4';
            return (
              <g key={`${node.step.id}-term-${i}`}>
                <line
                  x1={x} y1={y}
                  x2={x} y2={y + 20}
                  stroke={color}
                  strokeWidth={1.5}
                  strokeDasharray={dashArray}
                />
                <rect
                  x={x - 20} y={y + 20}
                  width={40} height={18}
                  rx={9}
                  fill="none"
                  stroke={color}
                  strokeWidth={1}
                />
                <text
                  x={x} y={y + 33}
                  textAnchor="middle"
                  fill={color}
                  fontSize={9}
                  fontWeight={500}
                >
                  {t.val}
                </text>
              </g>
            );
          });
        })}

        {/* Nodes */}
        {nodes.map((node) => (
          <g key={node.step.id}>
            <rect
              x={node.x}
              y={node.y}
              width={NODE_W}
              height={NODE_H}
              rx={8}
              fill="#18181b"
              stroke="#3f3f46"
              strokeWidth={1}
            />
            {/* Action text */}
            <text
              x={node.x + 12}
              y={node.y + 20}
              fill="#fafafa"
              fontSize={12}
              fontWeight={500}
            >
              {truncate(node.step.action, 34)}
            </text>
            {/* Description */}
            {node.step.description && (
              <text
                x={node.x + 12}
                y={node.y + 36}
                fill="#a1a1aa"
                fontSize={10}
              >
                {truncate(node.step.description, 40)}
              </text>
            )}
            {/* Contract chip */}
            {node.step.contract && (
              <g
                className="cursor-pointer"
                onClick={(e) => {
                  e.stopPropagation();
                  onContractClick?.(node.step.contract!);
                }}
              >
                <rect
                  x={node.x + 12}
                  y={node.y + NODE_H - 24}
                  width={Math.min(node.step.contract.length * 6.5 + 16, NODE_W - 24)}
                  height={18}
                  rx={4}
                  fill="#22c55e15"
                  stroke="#22c55e40"
                  strokeWidth={0.5}
                />
                <text
                  x={node.x + 20}
                  y={node.y + NODE_H - 12}
                  fill="#22c55e"
                  fontSize={9}
                  fontWeight={500}
                >
                  {truncate(node.step.contract, 36)}
                </text>
              </g>
            )}
            {/* Step ID badge */}
            <text
              x={node.x + NODE_W - 12}
              y={node.y + 16}
              textAnchor="end"
              fill="#52525b"
              fontSize={9}
              fontWeight={500}
            >
              {node.step.id}
            </text>
          </g>
        ))}
      </svg>
    </div>
  );
}

function truncate(s: string, max: number): string {
  return s.length > max ? s.slice(0, max - 1) + '\u2026' : s;
}

function layoutGraph(steps: FlowStep[]): {
  nodes: LayoutNode[];
  edges: Edge[];
  width: number;
  height: number;
} {
  if (steps.length === 0) return { nodes: [], edges: [], width: 0, height: 0 };

  const idSet = new Set(steps.map((s) => s.id));
  const edges: Edge[] = [];

  // Build edges from on_success/on_failure
  for (let i = 0; i < steps.length; i++) {
    const s = steps[i];
    const successTarget = s.on_success || (i < steps.length - 1 ? steps[i + 1].id : 'done');

    if (successTarget && successTarget !== 'done' && successTarget !== 'abort' && idSet.has(successTarget)) {
      edges.push({ from: s.id, to: successTarget, type: 'success' });
    }
    if (s.on_failure && s.on_failure !== 'done' && s.on_failure !== 'abort' && idSet.has(s.on_failure)) {
      edges.push({ from: s.id, to: s.on_failure, type: 'failure' });
    }
  }

  // Simple layout: main path in col 0, failure branches in col 1
  const mainPath = new Set<string>();
  let current = steps[0]?.id;
  while (current && idSet.has(current) && !mainPath.has(current)) {
    mainPath.add(current);
    const step = steps.find((s) => s.id === current);
    if (!step) break;
    current = step.on_success || steps[steps.indexOf(step) + 1]?.id;
    if (current === 'done' || current === 'abort') break;
  }

  const placed = new Map<string, { col: number; row: number }>();
  let row = 0;

  // Place main path
  for (const s of steps) {
    if (mainPath.has(s.id)) {
      placed.set(s.id, { col: 0, row });
      row++;
    }
  }

  // Place branching nodes
  for (const s of steps) {
    if (!placed.has(s.id)) {
      // Find the parent that branches to this node
      const parent = steps.find(
        (p) => p.on_failure === s.id || p.on_success === s.id
      );
      const parentPos = parent ? placed.get(parent.id) : undefined;
      const r = parentPos ? parentPos.row + 1 : row;
      placed.set(s.id, { col: 1, row: r });
      row = Math.max(row, r + 1);
    }
  }

  const nodes: LayoutNode[] = steps
    .filter((s) => placed.has(s.id))
    .map((s) => {
      const pos = placed.get(s.id)!;
      return {
        step: s,
        col: pos.col,
        row: pos.row,
        x: PAD + pos.col * (NODE_W + GAP_X),
        y: PAD + pos.row * (NODE_H + GAP_Y),
      };
    });

  const maxCol = Math.max(...nodes.map((n) => n.col));
  const maxRow = Math.max(...nodes.map((n) => n.row));

  return {
    nodes,
    edges,
    width: PAD * 2 + (maxCol + 1) * NODE_W + maxCol * GAP_X,
    height: PAD * 2 + (maxRow + 1) * NODE_H + maxRow * GAP_Y + 50, // extra for terminal markers
  };
}
