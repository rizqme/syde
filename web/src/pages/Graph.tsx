import { useEffect, useMemo, useRef, useState } from 'react';
import {
  forceCenter,
  forceCollide,
  forceLink,
  forceManyBody,
  forceSimulation,
  forceX,
  forceY,
  Simulation,
  SimulationLinkDatum,
  SimulationNodeDatum,
} from 'd3-force';
import { select } from 'd3-selection';
import { zoom, zoomIdentity, ZoomTransform } from 'd3-zoom';
import { useApi } from '../hooks/useApi';
import { api, EntityDetailResponse, GraphResponse } from '../lib/api';
import { iconForKind } from '../components/icons';

// Only architectural kinds appear on the entity graph. Plan/task/design
// are workflow concerns and are kept out so the network reads as a
// pure architecture map.
const INCLUDED_KINDS = new Set(['system', 'component', 'contract', 'concept', 'flow', 'decision']);

// Systems are flat (PLN-0019 — no root, no sub-system hierarchy), so
// every system renders at one size. Components are the granular unit;
// the rest are leaves.
const KIND_RADIUS: Record<string, number> = {
  system: 28,
  component: 18,
  contract: 11,
  concept: 11,
  flow: 10,
  decision: 10,
};

// Muted, dark-mode-friendly palette. Saturation pulled way down from the
// vivid kind tokens so the graph reads like a calm architecture map
// rather than a candy-coloured network. Each kind is the base colour;
// hovered/active nodes intensify via opacity instead of swapping colours.
const KIND_COLOR: Record<string, string> = {
  system: '#cbd5e1',     // slate 300 — the brightest, reserved for systems
  component: '#7aa2c7',  // muted blue
  contract: '#7fb289',   // muted green / sage
  concept: '#9b8bbf',    // muted lavender
  flow: '#c89171',       // muted terracotta
  decision: '#bfa365',   // muted ochre
};

// Edge stroke colour per relationship type. Belongs_to is the structural
// backbone, depends_on the dependency wiring, the rest are weaker links.
// All hand-tuned to sit quietly on the #0a0c10 canvas.
const EDGE_COLOR: Record<string, string> = {
  belongs_to: 'rgba(203, 213, 225, 0.32)',
  depends_on: 'rgba(122, 162, 199, 0.38)',
  exposes: 'rgba(127, 178, 137, 0.38)',
  references: 'rgba(155, 139, 191, 0.30)',
  involves: 'rgba(200, 145, 113, 0.36)',
  applies_to: 'rgba(191, 163, 101, 0.36)',
  relates_to: 'rgba(155, 139, 191, 0.28)',
};

interface GraphNode extends SimulationNodeDatum {
  id: string;
  name: string;
  kind: string;
  sizeKey: string;
  radius: number;
  color: string;
}

interface GraphLink extends SimulationLinkDatum<GraphNode> {
  source: string | GraphNode;
  target: string | GraphNode;
  type: string;
}

interface GraphProps {
  onSelectEntity: (slug: string, kind?: string) => void;
}

export function Graph({ onSelectEntity }: GraphProps) {
  const { data, loading, error } = useApi<GraphResponse>(() => api.graph(), []);
  const containerRef = useRef<HTMLDivElement | null>(null);
  const svgRef = useRef<SVGSVGElement | null>(null);
  const viewportRef = useRef<SVGGElement | null>(null);
  // Transform lives in a ref, not state — d3-zoom drives the wrapping <g>
  // directly via setAttribute so we don't re-render 300+ SVG elements on
  // every pan tick. Dragged nodes still need the live transform to convert
  // pointer coords into graph space.
  const transformRef = useRef<ZoomTransform>(zoomIdentity);
  const [hovered, setHovered] = useState<string | null>(null);
  const [selectedNodeId, setSelectedNodeId] = useState<string | null>(null);
  const [selectedSlug, setSelectedSlug] = useState<string | null>(null);
  const [selectedKind, setSelectedKind] = useState<string | null>(null);
  // Highlight follows selection if pinned, otherwise hover.
  const focused = selectedNodeId ?? hovered;

  // Active node drag — grabs onto a single node and lets the user pull it
  // around. d3-zoom is filtered to ignore pointer events that originated on
  // a node so the canvas doesn't pan while dragging.
  const dragRef = useRef<{ node: GraphNode } | null>(null);

  function clearSelection() {
    setSelectedNodeId(null);
    setSelectedSlug(null);
    setSelectedKind(null);
  }
  const [size, setSize] = useState({ w: 800, h: 600 });

  // Keep SVG sized to its container.
  useEffect(() => {
    if (!containerRef.current) return;
    const el = containerRef.current;
    const ro = new ResizeObserver(() => {
      setSize({ w: el.clientWidth, h: el.clientHeight });
    });
    ro.observe(el);
    setSize({ w: el.clientWidth, h: el.clientHeight });
    return () => ro.disconnect();
  }, []);

  // Build nodes + links, run d3-force, and store positioned data in state.
  // Re-run only when the API payload changes; layout stays stable across
  // hover/zoom by living in a ref.
  const built = useMemo(() => {
    if (!data?.nodes) return null;

    // The backend resolves every relationship target to a canonical
    // entity ID before emitting edges, so a single id-keyed index is
    // enough to resolve both sources and targets.
    const allowedById = new Map<string, GraphResponse['nodes'][number]>();
    for (const n of data.nodes) {
      if (!INCLUDED_KINDS.has(n.kind)) continue;
      allowedById.set(n.id, n);
    }

    // Systems are flat under PLN-0019 — every system renders at one
    // tier with the same size and colour. No root vs sub-system
    // distinction is computed.

    const nodes: GraphNode[] = [];
    for (const n of data.nodes) {
      if (!INCLUDED_KINDS.has(n.kind)) continue;
      nodes.push({
        id: n.id,
        name: n.name,
        kind: n.kind,
        sizeKey: n.kind,
        radius: KIND_RADIUS[n.kind] ?? 8,
        color: KIND_COLOR[n.kind] ?? '#94a3b8',
      });
    }

    const nodeById = new Map(nodes.map((n) => [n.id, n]));

    // Build links — drop any that touch a filtered-out node.
    const links: GraphLink[] = [];
    for (const e of data.edges) {
      const src = nodeById.get(e.source);
      const tgt = nodeById.get(e.target);
      if (!src || !tgt) continue;
      links.push({ source: src.id, target: tgt.id, type: e.type });
    }

    return { nodes, links };
  }, [data]);

  // Run the simulation when nodes/links change. Hold the simulation in a
  // ref so we can stop / restart cleanly.
  const simRef = useRef<Simulation<GraphNode, GraphLink> | null>(null);
  const [, forceTick] = useState(0);
  useEffect(() => {
    if (!built) return;
    const w = size.w || 800;
    const h = size.h || 600;
    const sim = forceSimulation<GraphNode, GraphLink>(built.nodes)
      .force(
        'link',
        forceLink<GraphNode, GraphLink>(built.links)
          .id((d) => d.id)
          .distance((l) => {
            const tgtKind = (l.target as GraphNode).kind;
            const srcKind = (l.source as GraphNode).kind;
            // Mid-density tuning: roomy enough that labels don't collide,
            // tight enough that related entities still feel grouped.
            if (l.type === 'belongs_to') return 95;
            if (l.type === 'depends_on') return 120;
            if (srcKind === 'system' || tgtKind === 'system') return 150;
            return 130;
          })
          .strength(0.45),
      )
      .force('charge', forceManyBody<GraphNode>().strength((d) => -240 - d.radius * 9).distanceMax(700))
      .force('center', forceCenter(w / 2, h / 2))
      .force('x', forceX<GraphNode>(w / 2).strength(0.04))
      .force('y', forceY<GraphNode>(h / 2).strength(0.04))
      .force(
        'collide',
        forceCollide<GraphNode>().radius((d) => d.radius + 10).iterations(2),
      )
      // Heavy damping + fast cool-down: the layout snaps into place in
      // ~80 ticks instead of drifting for ~400. Stops feeling like a
      // restless lava-lamp without losing the natural force-directed look.
      .velocityDecay(0.55)
      .alpha(0.9)
      .alphaDecay(0.08)
      .alphaMin(0.01);

    sim.on('tick', () => forceTick((n) => n + 1));
    // Auto-stop hard once it's converged so React isn't re-rendering
    // forever for sub-pixel jitter.
    sim.on('end', () => sim.stop());

    simRef.current = sim;
    return () => {
      sim.stop();
      simRef.current = null;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [built]);

  // Re-centre the simulation when the container is resized so the layout
  // tracks the visible viewport. Update both the centre force and the
  // forceX/forceY containment targets.
  useEffect(() => {
    const sim = simRef.current;
    if (!sim) return;
    sim.force('center', forceCenter(size.w / 2, size.h / 2));
    const fx = sim.force('x') as ReturnType<typeof forceX<GraphNode>> | undefined;
    const fy = sim.force('y') as ReturnType<typeof forceY<GraphNode>> | undefined;
    fx?.x(size.w / 2);
    fy?.y(size.h / 2);
    sim.alpha(0.4).restart();
  }, [size.w, size.h]);

  // Wire d3-zoom for pan/zoom. The filter rejects pointer events whose
  // target sits inside a [data-graph-node] group so node dragging doesn't
  // also pan the whole canvas — but wheel events always pass so the user
  // can zoom even when the cursor is over a node.
  useEffect(() => {
    if (!svgRef.current) return;
    const svg = select(svgRef.current);
    const z = zoom<SVGSVGElement, unknown>()
      .scaleExtent([0.2, 4])
      .filter((event: any) => {
        if (event.type === 'wheel') return true;
        if (event.button) return false;
        const t = event.target as Element | null;
        if (t && t.closest && t.closest('[data-graph-node]')) return false;
        return true;
      })
      .on('zoom', (event) => {
        // Drive the viewport <g> imperatively. Bypassing React state here
        // keeps pan / zoom at 60 fps even with hundreds of SVG children;
        // re-rendering the React tree per pointermove was the real source
        // of the "drag doesn't work" feeling.
        transformRef.current = event.transform;
        if (viewportRef.current) {
          viewportRef.current.setAttribute(
            'transform',
            `translate(${event.transform.x},${event.transform.y}) scale(${event.transform.k})`,
          );
        }
      });
    svg.call(z);
    return () => {
      svg.on('.zoom', null);
    };
  }, []);

  // Begin a node drag: pin the node's fx/fy under the pointer and attach
  // window-level move/up listeners. The simulation is reheated so the
  // dragged node tugs its neighbours along. Selection happens on pointerup
  // unconditionally — whether the user dragged or just clicked, the inspector
  // for that node opens.
  function beginDrag(node: GraphNode, ev: React.PointerEvent) {
    if (!svgRef.current) return;
    // stopPropagation so d3-zoom (the canvas pan handler) doesn't also see
    // this gesture as a pan-start.
    ev.stopPropagation();

    dragRef.current = { node };
    node.fx = node.x;
    node.fy = node.y;
    simRef.current?.alphaTarget(0.3).restart();

    const onMove = (mv: PointerEvent) => {
      const drag = dragRef.current;
      if (!drag || !svgRef.current) return;
      const rect = svgRef.current.getBoundingClientRect();
      const t = transformRef.current;
      drag.node.fx = (mv.clientX - rect.left - t.x) / t.k;
      drag.node.fy = (mv.clientY - rect.top - t.y) / t.k;
    };
    const onUp = () => {
      simRef.current?.alphaTarget(0);
      dragRef.current = null;
      window.removeEventListener('pointermove', onMove);
      window.removeEventListener('pointerup', onUp);
      // Always promote the released node to the active selection — drag or
      // click, both surface the inspector.
      setSelectedNodeId(node.id);
      setSelectedSlug(slugifyName(node.name));
      setSelectedKind(node.kind);
    };
    window.addEventListener('pointermove', onMove);
    window.addEventListener('pointerup', onUp);
  }

  // Loading + error are rendered as overlays so the canvas div (and its
  // ref) always exists from first mount. Without this the d3-zoom binding
  // useEffect runs once on mount when the SVG doesn't yet exist, sees a
  // null ref, bails, and never re-runs because its deps are [] —
  // breaking pan / zoom forever.
  const counts = built ? countByKind(built.nodes) : {};

  return (
    <div className="h-full flex flex-col min-h-0">
      {/* Header */}
      <div className="px-6 py-4 border-b border-border shrink-0">
        <div className="flex items-baseline gap-3">
          <h1 className="text-lg font-semibold">Architecture Graph</h1>
          <div className="flex items-center gap-3 text-[11px] text-muted-foreground">
            <span><span className="text-foreground font-medium">{built?.nodes.length ?? 0}</span> entities</span>
            <span className="text-border">·</span>
            <span><span className="text-foreground font-medium">{built?.links.length ?? 0}</span> relationships</span>
          </div>
        </div>
      </div>

      {/* Canvas + legend grid. The canvas cell takes everything left of the
          fixed-width legend; the SVG is sized at 100% / 100% so it exactly
          fills its grid cell no matter the viewport. */}
      <div className="flex-1 min-h-0 grid grid-cols-[minmax(0,1fr)_220px]">
        <div ref={containerRef} className="relative bg-[#0a0c10] overflow-hidden min-w-0">
          <svg
            ref={svgRef}
            width="100%"
            height="100%"
            className="absolute inset-0 w-full h-full cursor-grab active:cursor-grabbing"
            onClick={(ev) => {
              // Only deselect when the click was on bare canvas, not a node.
              // Any descendant node carries data-graph-node up the DOM, so
              // bubbled clicks from inside one are ignored here.
              const t = ev.target as Element;
              if (t.closest && t.closest('[data-graph-node]')) return;
              clearSelection();
            }}
          >
            <defs>
              <radialGradient id="node-glow" cx="50%" cy="50%" r="50%">
                <stop offset="0%" stopColor="white" stopOpacity="0.18" />
                <stop offset="100%" stopColor="white" stopOpacity="0" />
              </radialGradient>
            </defs>
            <g ref={viewportRef}>
              {/* Edges */}
              {built?.links.map((l, i) => {
                const s = l.source as GraphNode;
                const t = l.target as GraphNode;
                if (typeof s !== 'object' || typeof t !== 'object') return null;
                const isHot = focused && (s.id === focused || t.id === focused);
                return (
                  <line
                    key={i}
                    x1={s.x ?? 0}
                    y1={s.y ?? 0}
                    x2={t.x ?? 0}
                    y2={t.y ?? 0}
                    stroke={EDGE_COLOR[l.type] ?? 'rgba(148, 163, 184, 0.3)'}
                    strokeWidth={isHot ? 1.6 : 0.9}
                    opacity={focused && !isHot ? 0.12 : 1}
                  />
                );
              })}

              {/* Nodes */}
              {built?.nodes.map((n) => {
                const isHot = focused === n.id;
                const isPinned = selectedNodeId === n.id;
                const dim = focused && !isHot && !isNeighbor(n.id, focused, built!.links);
                return (
                  <g
                    key={n.id}
                    data-graph-node
                    transform={`translate(${n.x ?? 0},${n.y ?? 0})`}
                    onMouseEnter={() => setHovered(n.id)}
                    onMouseLeave={() => setHovered(null)}
                    onPointerDown={(ev) => beginDrag(n, ev)}
                    className="cursor-grab active:cursor-grabbing"
                    opacity={dim ? 0.2 : 1}
                  >
                    {/* Glow halo for the larger nodes */}
                    {n.radius >= 16 && (
                      <circle r={n.radius + 8} fill="url(#node-glow)" />
                    )}
                    {/* Pinned ring — extra outline around the actively
                        selected node so it stays obvious after the cursor
                        moves away. */}
                    {isPinned && (
                      <circle
                        r={n.radius + 5}
                        fill="none"
                        stroke={n.color}
                        strokeOpacity={0.9}
                        strokeWidth={1.5}
                      />
                    )}
                    <circle
                      r={n.radius}
                      fill={n.color}
                      fillOpacity={n.kind === 'system' ? 0.95 : 0.85}
                      stroke={n.color}
                      strokeOpacity={isPinned ? 1 : 0.9}
                      strokeWidth={isPinned ? 3 : isHot ? 2.5 : 1.2}
                    />
                    <text
                      textAnchor="middle"
                      dy={n.radius + 12}
                      fill={isHot ? '#fafafa' : 'rgba(226,232,240,0.85)'}
                      fontSize={n.kind === 'system' ? 12 : n.kind === 'component' ? 10 : 9}
                      fontWeight={n.kind === 'system' ? 600 : 400}
                      style={{ pointerEvents: 'none', userSelect: 'none' }}
                    >
                      {n.name}
                    </text>
                  </g>
                );
              })}
            </g>
          </svg>

          {/* Loading / error overlays. Rendered on top of the canvas instead
              of replacing it so refs (and thus d3-zoom) bind on first mount. */}
          {loading && !built && (
            <div className="absolute inset-0 flex items-center justify-center text-muted-foreground text-sm pointer-events-none">
              Loading graph…
            </div>
          )}
          {error && !built && (
            <div className="absolute inset-0 flex items-center justify-center text-red-400 text-sm pointer-events-none">
              Failed to load graph.
            </div>
          )}

          {/* Floating help text */}
          <div className="absolute bottom-3 left-3 text-[10px] text-muted-foreground/80 pointer-events-none">
            scroll to zoom · drag canvas to pan · drag a node to move it · click to inspect
          </div>

          {/* Slide-in inspector panel — opens when a node is clicked. */}
          {selectedSlug && (
            <NodeInspector
              slug={selectedSlug}
              kind={selectedKind}
              onClose={clearSelection}
              onOpen={() => {
                if (selectedSlug) onSelectEntity(selectedSlug, selectedKind || undefined);
              }}
            />
          )}
        </div>

        {/* Legend pane */}
        <aside className="border-l border-border bg-background overflow-y-auto px-4 py-5">
          <h3 className="text-[10px] font-medium uppercase tracking-widest text-muted-foreground mb-3">
            Legend
          </h3>
          <ul className="space-y-2.5 mb-6">
            {LEGEND_ENTRIES.map((entry) => (
              <li key={entry.label} className="flex items-center gap-2.5 text-[11px]">
                <span
                  className="rounded-full shrink-0"
                  style={{
                    width: entry.dot,
                    height: entry.dot,
                    background: entry.color,
                    boxShadow: `0 0 8px ${entry.color}55`,
                  }}
                />
                <span className="text-foreground">{entry.label}</span>
                {counts[entry.kind] != null && (
                  <span className="ml-auto text-muted-foreground tabular-nums">
                    {counts[entry.kind]}
                  </span>
                )}
              </li>
            ))}
          </ul>

          <h3 className="text-[10px] font-medium uppercase tracking-widest text-muted-foreground mb-3">
            Edges
          </h3>
          <ul className="space-y-2 text-[11px]">
            {EDGE_LEGEND.map((e) => (
              <li key={e.type} className="flex items-center gap-2.5">
                <span
                  className="h-px w-6 shrink-0"
                  style={{ background: e.color, boxShadow: `0 0 6px ${e.color}` }}
                />
                <span className="text-foreground">{e.label}</span>
              </li>
            ))}
          </ul>
        </aside>
      </div>
    </div>
  );
}

const LEGEND_ENTRIES = [
  { kind: 'system', label: 'System', color: KIND_COLOR.system, dot: 15 },
  { kind: 'component', label: 'Component', color: KIND_COLOR.component, dot: 11 },
  { kind: 'contract', label: 'Contract', color: KIND_COLOR.contract, dot: 9 },
  { kind: 'concept', label: 'Concept', color: KIND_COLOR.concept, dot: 9 },
  { kind: 'flow', label: 'Flow', color: KIND_COLOR.flow, dot: 8 },
  { kind: 'decision', label: 'Decision', color: KIND_COLOR.decision, dot: 8 },
];

const EDGE_LEGEND = [
  { type: 'belongs_to', label: 'belongs to', color: EDGE_COLOR.belongs_to },
  { type: 'depends_on', label: 'depends on', color: EDGE_COLOR.depends_on },
  { type: 'exposes', label: 'exposes', color: EDGE_COLOR.exposes },
  { type: 'references', label: 'references', color: EDGE_COLOR.references },
  { type: 'involves', label: 'involves', color: EDGE_COLOR.involves },
  { type: 'applies_to', label: 'applies to', color: EDGE_COLOR.applies_to },
];

function NodeInspector({
  slug,
  kind,
  onClose,
  onOpen,
}: {
  slug: string;
  kind: string | null;
  onClose: () => void;
  onOpen: () => void;
}) {
  const { data, loading, error } = useApi<EntityDetailResponse>(() => api.entity(slug), [slug]);
  const Icon = kind ? iconForKind(kind) : null;
  const accent = kind ? KIND_COLOR[kind] : '#cbd5e1';
  const e = data?.entity;
  const name = e?.name || slug;
  const description = e?.description as string | undefined;
  const purpose = e?.purpose as string | undefined;
  const responsibility = e?.responsibility as string | undefined;
  const meaning = e?.meaning as string | undefined;
  const statement = e?.statement as string | undefined;
  const trigger = e?.trigger as string | undefined;
  const goal = e?.goal as string | undefined;

  return (
    <div
      className="absolute top-3 right-3 bottom-3 w-80 rounded-xl border border-border/80 bg-background/95 backdrop-blur-md shadow-2xl flex flex-col overflow-hidden animate-in fade-in slide-in-from-right-4 duration-150"
      onClick={(e) => e.stopPropagation()}
    >
      {/* Header */}
      <div className="flex items-start gap-3 px-4 py-4 border-b border-border/60">
        <span
          className="flex items-center justify-center w-9 h-9 rounded-lg shrink-0"
          style={{ background: `${accent}22`, color: accent, boxShadow: `0 0 14px ${accent}33` }}
        >
          {Icon ? <Icon className="w-5 h-5" /> : null}
        </span>
        <div className="min-w-0 flex-1">
          <div className="text-[10px] uppercase tracking-widest text-muted-foreground mb-0.5">
            {kind || 'entity'}
          </div>
          <button
            type="button"
            onClick={onOpen}
            title="Open in detail view"
            className="text-left text-base font-semibold text-foreground hover:underline truncate w-full"
          >
            {name}
          </button>
        </div>
        <button
          type="button"
          onClick={onClose}
          aria-label="Close"
          className="text-muted-foreground hover:text-foreground text-sm"
        >
          ✕
        </button>
      </div>

      {/* Body */}
      <div className="flex-1 overflow-y-auto px-4 py-4 space-y-4 text-sm">
        {loading && !data && (
          <div className="text-muted-foreground text-xs">Loading…</div>
        )}
        {error && (
          <div className="text-red-400 text-xs">Failed to load entity.</div>
        )}
        {description && (
          <p className="text-foreground/90 leading-relaxed">{description}</p>
        )}
        {responsibility && (
          <Field label="Responsibility" value={responsibility} />
        )}
        {meaning && <Field label="Meaning" value={meaning} />}
        {statement && <Field label="Decision" value={statement} />}
        {trigger && <Field label="Trigger" value={trigger} />}
        {goal && <Field label="Goal" value={goal} />}
        {purpose && !responsibility && !meaning && !statement && (
          <Field label="Purpose" value={purpose} />
        )}
      </div>

      {/* Footer */}
      <div className="px-4 py-3 border-t border-border/60 shrink-0">
        <button
          type="button"
          onClick={onOpen}
          className="w-full text-xs font-medium text-foreground bg-card border border-border hover:bg-accent hover:border-muted-foreground rounded-md px-3 py-2 transition-colors"
        >
          Open full {kind || 'entity'} →
        </button>
      </div>
    </div>
  );
}

function Field({ label, value }: { label: string; value: string }) {
  return (
    <div>
      <div className="text-[10px] font-medium uppercase tracking-widest text-muted-foreground mb-1">
        {label}
      </div>
      <p className="text-xs text-foreground/85 leading-relaxed">{value}</p>
    </div>
  );
}

// Mirror of internal/utils/slug.go's normaliser.
function slugifyName(name: string): string {
  return name
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '');
}

function countByKind(nodes: GraphNode[]): Record<string, number> {
  const out: Record<string, number> = {};
  for (const n of nodes) {
    out[n.sizeKey] = (out[n.sizeKey] ?? 0) + 1;
    out[n.kind] = (out[n.kind] ?? 0) + 1;
  }
  return out;
}

function isNeighbor(nodeId: string, hoveredId: string, links: GraphLink[]): boolean {
  for (const l of links) {
    const s = typeof l.source === 'object' ? (l.source as GraphNode).id : l.source;
    const t = typeof l.target === 'object' ? (l.target as GraphNode).id : l.target;
    if ((s === hoveredId && t === nodeId) || (t === hoveredId && s === nodeId)) return true;
  }
  return false;
}
