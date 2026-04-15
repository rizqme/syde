import { useState } from 'react';
import { useApi } from '../hooks/useApi';
import { api, EntityDetailResponse, Relationship } from '../lib/api';
import { KindBadge } from './KindBadge';
import { iconForKind, FileIcon, ChevronRightIcon } from './icons';

interface EntityDetailProps {
  slug: string;
  // Relationship clicks pass the target's kind so the parent can switch
  // pages when the link points at a different entity kind. The optional
  // filter param pre-seeds the destination kind page's filter bar — used
  // when navigating to a child entity to scope the list to siblings.
  onNavigate: (slug: string, kind?: string, filter?: string) => void;
  // File card clicks open the file in the FileTree page.
  onOpenFile?: (path: string) => void;
  onClose: () => void;
  // When true, render without the floating w-96 shell. Used by the
  // inbox-style 2-column entity view where the detail fills the right pane.
  inline?: boolean;
  // When true, hide the X close button in the top-right. Used by the
  // concept page where the List/ERD toggle occupies that spot.
  hideClose?: boolean;
}

export function EntityDetail({ slug, onNavigate, onOpenFile, onClose, inline, hideClose }: EntityDetailProps) {
  const { data, loading, error } = useApi<EntityDetailResponse>(() => api.entity(slug), [slug]);

  if (loading) return <DetailShell onClose={onClose} inline={inline} hideClose={hideClose}><div className="text-muted-foreground text-sm p-4">Loading...</div></DetailShell>;
  if (error || !data) return <DetailShell onClose={onClose} inline={inline} hideClose={hideClose}><div className="text-red-400 text-sm p-4">Entity not found: {slug}</div></DetailShell>;

  const e = data.entity;
  const kind = e.kind as string;

  return (
    <DetailShell onClose={onClose} inline={inline} hideClose={hideClose}>
      <div className="p-4 border-b border-border">
        <div className="flex items-center gap-2 mb-2">
          <KindBadge kind={kind} />
          <h2 className="text-lg font-semibold">{e.name}</h2>
        </div>
        {e.deprecated && (
          <div className="mb-2 px-3 py-2 bg-yellow-900/20 border border-yellow-600/30 rounded text-xs text-yellow-400">
            Deprecated{e.deprecated_reason ? `: ${e.deprecated_reason}` : ''}
            {e.replaced_by && (
              <button onClick={() => onNavigate(e.replaced_by, kind)} className="ml-1 underline">
                → {e.replaced_by}
              </button>
            )}
          </div>
        )}
        {e.description && <p className="text-sm text-muted-foreground">{e.description}</p>}
      </div>

      <div className="p-4 space-y-4 overflow-y-auto flex-1">
        <KindFields entity={e} kind={kind} />

        {kind === 'system' && data.relationships && (
          <SystemChildren
            relationships={data.relationships}
            parentSlug={baseSlug(e.slug || slug)}
            onNavigate={onNavigate}
          />
        )}

        {((data.file_refs && data.file_refs.length > 0) || e.files?.length > 0) && (
          <Section title={`Files · ${(data.file_refs?.length || e.files?.length || 0)}`}>
            <ul className="rounded-lg border border-border bg-card/40 divide-y divide-border/60 overflow-hidden">
              {(data.file_refs && data.file_refs.length > 0
                ? data.file_refs
                : (e.files || []).map((f: string) => ({ path: f, summary: '', in_tree: false, stale: false }))
              ).map((ref: any) => {
                const clickable = !!onOpenFile && ref.in_tree;
                const inner = (
                  <div className="flex items-start gap-3 px-3 py-2.5 min-w-0 w-full text-left">
                    <FileIcon className="w-3.5 h-3.5 mt-0.5 shrink-0 text-muted-foreground" />
                    <div className="min-w-0 flex-1">
                      <div className="flex items-center gap-2 flex-wrap">
                        <span className="text-[12px] font-mono text-foreground truncate">
                          {ref.path}
                        </span>
                        {!ref.in_tree && (
                          <span className="text-[9px] uppercase tracking-wider px-1.5 py-px rounded bg-yellow-500/15 text-yellow-400">
                            not in tree
                          </span>
                        )}
                        {ref.stale && (
                          <span className="text-[9px] uppercase tracking-wider px-1.5 py-px rounded bg-orange-500/15 text-orange-400">
                            stale
                          </span>
                        )}
                      </div>
                      {ref.summary && (
                        <div className="text-[11px] text-muted-foreground mt-1 leading-relaxed">
                          {ref.summary}
                        </div>
                      )}
                    </div>
                    {clickable && (
                      <ChevronRightIcon className="w-3 h-3 mt-1 shrink-0 text-muted-foreground/60" />
                    )}
                  </div>
                );
                return (
                  <li key={ref.path}>
                    {clickable ? (
                      <button
                        type="button"
                        onClick={() => onOpenFile!(ref.path)}
                        className="w-full hover:bg-accent/30 transition-colors"
                      >
                        {inner}
                      </button>
                    ) : (
                      inner
                    )}
                  </li>
                );
              })}
            </ul>
          </Section>
        )}

        {e.notes?.length > 0 && (
          <Section title="Notes">
            <ul className="text-xs space-y-1">
              {e.notes.map((n: string, i: number) => (
                <li key={i} className="text-muted-foreground leading-relaxed">• {n}</li>
              ))}
            </ul>
          </Section>
        )}

        {(() => {
          // For systems, the inbound belongs_to relationships are shown in
          // the "Children" block above — filter them out here so we don't
          // render the same chips twice.
          const visible = (data.relationships || []).filter((rel) => {
            if (kind !== 'system') return true;
            return !(rel.direction === 'inbound' && rel.type === 'belongs_to');
          });
          if (visible.length === 0) return null;
          return <RelationshipsSection rels={visible} onNavigate={onNavigate} />;
        })()}

        {data.tasks?.length > 0 && (
          <Section title="Tasks">
            {data.tasks.map((t: any, i: number) => (
              <div key={i} className="flex items-center gap-2 text-xs mb-1">
                <span>{t.status === 'completed' ? '✓' : t.status === 'in_progress' ? '●' : '○'}</span>
                <span>{t.name}</span>
                <span className="text-muted-foreground">{t.priority}</span>
              </div>
            ))}
          </Section>
        )}

        {e.tags?.length > 0 && (
          <Section title="Tags">
            <div className="flex flex-wrap gap-1">
              {e.tags.map((t: string) => (
                <span key={t} className="text-[10px] px-1.5 py-0.5 rounded bg-muted text-muted-foreground">{t}</span>
              ))}
            </div>
          </Section>
        )}
      </div>
    </DetailShell>
  );
}

function DetailShell({ children, onClose, inline, hideClose }: { children: React.ReactNode; onClose: () => void; inline?: boolean; hideClose?: boolean }) {
  // Inline: fills the right pane of the 2-column entity view, no border-l
  // (the list pane already has border-r), no fixed width.
  // Floating: legacy 384px panel for special views (plan/task).
  const wrapperClass = inline
    ? 'flex flex-col h-full bg-background min-h-0'
    : 'w-96 border-l border-border bg-background flex flex-col h-full shrink-0';
  return (
    <div className={wrapperClass}>
      {!hideClose && (
        <div className="flex justify-end p-2">
          <button onClick={onClose} className="text-muted-foreground hover:text-foreground text-sm px-2">
            ✕
          </button>
        </div>
      )}
      {children}
    </div>
  );
}

function Section({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <div>
      <h3 className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-2">{title}</h3>
      {children}
    </div>
  );
}

function KindFields({ entity: e, kind }: { entity: Record<string, any>; kind: string }) {
  switch (kind) {
    case 'component':
      return (
        <div className="space-y-3">
          {e.responsibility && <Field label="Responsibility" value={e.responsibility} highlight />}
          {e.capabilities?.length > 0 && (
            <div>
              <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-2">Capabilities</div>
              <ul className="space-y-1.5">
                {e.capabilities.map((c: string, i: number) => (
                  <li
                    key={i}
                    className="flex items-center gap-2.5 text-xs leading-relaxed rounded-md border border-border bg-card px-3 py-2 text-foreground"
                  >
                    <span className="inline-block h-1.5 w-1.5 shrink-0 rounded-full bg-kind-component" />
                    <span className="min-w-0">{c}</span>
                  </li>
                ))}
              </ul>
            </div>
          )}
          {e.boundaries && <Field label="Boundaries" value={e.boundaries} muted />}
          {e.behavior_summary && <Field label="Behavior" value={e.behavior_summary} />}
          {e.interaction_summary && <Field label="Interactions" value={e.interaction_summary} />}
          {e.data_handling && <Field label="Data Handling" value={e.data_handling} />}
          {e.scaling_notes && <Field label="Scaling" value={e.scaling_notes} />}
        </div>
      );
    case 'contract':
      return (
        <div className="space-y-3">
          <div className="flex gap-2">
            {e.contract_kind && <KindBadge kind={e.contract_kind} />}
            {e.interaction_pattern && <KindBadge kind={e.interaction_pattern} />}
          </div>
          {e.input && <Field label="Input" value={e.input} code />}
          {e.input_parameters?.length > 0 && (
            <ParamTable label="Input Parameters" params={e.input_parameters} />
          )}
          {e.output && <Field label="Output" value={e.output} code />}
          {e.output_parameters?.length > 0 && (
            <ParamTable label="Output Parameters" params={e.output_parameters} />
          )}
          {e.protocol_notes && <Field label="Protocol" value={e.protocol_notes} />}
          {e.constraints && <Field label="Constraints" value={e.constraints} />}
          {e.contract_kind === 'screen' && e.wireframe_html && (
            <div>
              <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
                Wireframe
              </div>
              <div
                className="rounded border border-border bg-card p-3 overflow-x-auto text-xs"
                // dangerouslySetInnerHTML is safe here: the HTML
                // comes from our own server-side uiml.RenderHTML,
                // not user-supplied markup.
                dangerouslySetInnerHTML={{ __html: e.wireframe_html }}
              />
            </div>
          )}
        </div>
      );
    case 'concept':
      return (
        <div className="space-y-3">
          {e.meaning && <Field label="Meaning" value={e.meaning} />}
          {e.lifecycle && <Field label="Lifecycle" value={e.lifecycle} />}
          {e.invariants && <Field label="Invariants" value={e.invariants} />}
          {e.data_sensitivity && <Field label="Data Sensitivity" value={e.data_sensitivity} />}
          {e.attributes?.length > 0 && (
            <ParamTable
              label="Attributes"
              params={e.attributes.map((a: { name: string; description?: string }) => ({
                path: a.name,
                description: a.description,
              }))}
            />
          )}
          {e.actions?.length > 0 && (
            <div>
              <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">Actions</div>
              <div className="rounded border border-border bg-card divide-y divide-border">
                {e.actions.map((a: { name: string; description?: string }, i: number) => (
                  <div key={i} className="p-2 text-xs">
                    <div className="font-mono text-foreground">{a.name}</div>
                    {a.description && (
                      <div className="text-muted-foreground mt-0.5 leading-relaxed">{a.description}</div>
                    )}
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      );
    case 'flow':
      return (
        <div className="space-y-3">
          {(e.trigger || e.goal) && (
            <div className="flex items-center gap-2 text-sm">
              {e.trigger && <span className="text-muted-foreground">{e.trigger}</span>}
              {e.trigger && e.goal && <span className="text-kind-flow">→</span>}
              {e.goal && <span className="font-medium">{e.goal}</span>}
            </div>
          )}
          {e.narrative && <Field label="Narrative" value={e.narrative} />}
          {e.happy_path && <Field label="Happy Path" value={e.happy_path} />}
          {e.edge_cases && <Field label="Edge Cases" value={e.edge_cases} />}
          {e.failure_modes && <Field label="Failure Modes" value={e.failure_modes} />}
        </div>
      );
    case 'requirement':
      return (
        <div className="space-y-3">
          {e.statement && (
            <blockquote className="border-l-2 border-kind-requirement pl-3 text-sm italic">{e.statement}</blockquote>
          )}
          <div className="flex flex-wrap items-center gap-2 text-xs">
            {e.req_type && (
              <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">{e.req_type}</span>
            )}
            {e.priority && (
              <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">priority: {e.priority}</span>
            )}
            {e.requirement_status && (
              <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">{e.requirement_status}</span>
            )}
          </div>
          {e.rationale && <Field label="Rationale" value={e.rationale} />}
          {e.verification && <Field label="Verification" value={e.verification} muted />}
          {e.source && (
            <div className="text-xs text-muted-foreground">
              Source: <span className="text-foreground">{e.source}</span>
              {e.source_ref && <> · ref: <span className="text-foreground">{e.source_ref}</span></>}
            </div>
          )}
          {Array.isArray(e.supersedes) && e.supersedes.length > 0 && (
            <div className="text-xs text-muted-foreground">
              Supersedes: <span className="text-foreground">{e.supersedes.join(', ')}</span>
            </div>
          )}
          {Array.isArray(e.superseded_by) && e.superseded_by.length > 0 && (
            <div className="text-xs text-muted-foreground">
              Superseded by: <span className="text-foreground">{e.superseded_by.join(', ')}</span>
            </div>
          )}
          {e.obsolete_reason && <Field label="Obsolete Reason" value={e.obsolete_reason} muted />}
        </div>
      );
    case 'plan':
      return (
        <div className="space-y-3">
          {e.plan_status && (
            <div className="flex items-center gap-2">
              <span className="text-[10px] uppercase text-muted-foreground">Status</span>
              <span className="text-xs px-2 py-0.5 rounded bg-muted">{e.plan_status}</span>
              {typeof e.progress === 'number' && (
                <span className="text-xs text-muted-foreground">{Math.round(e.progress)}%</span>
              )}
            </div>
          )}
          {e.background && <Field label="Background" value={e.background} highlight />}
          {e.objective && <Field label="Objective" value={e.objective} highlight />}
          {e.scope && <Field label="Scope" value={e.scope} />}
          {e.phases?.length > 0 && (
            <div>
              <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-2">
                Phases ({e.phases.length})
              </div>
              <div className="space-y-2">
                {e.phases.map((ph: any) => (
                  <div key={ph.id} className="rounded border border-border p-2 text-xs">
                    <div className="flex items-center gap-2">
                      <span>
                        {ph.status === 'completed' ? '✓' : ph.status === 'in_progress' ? '●' : '○'}
                      </span>
                      <span className="font-medium">{ph.name || ph.description || ph.id}</span>
                      <span className="text-muted-foreground text-[10px]">{ph.status}</span>
                    </div>
                    {ph.objective && (
                      <div className="mt-1 text-muted-foreground">
                        <span className="text-foreground">Objective:</span> {ph.objective}
                      </div>
                    )}
                    {ph.changes && (
                      <div className="mt-1 text-muted-foreground">
                        <span className="text-foreground">Changes:</span> {ph.changes}
                      </div>
                    )}
                    {ph.details && <div className="mt-1 text-muted-foreground">{ph.details}</div>}
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      );
    case 'task':
      return (
        <div className="space-y-3">
          <div className="flex items-center gap-2 text-xs">
            {e.task_status && (
              <span className="px-2 py-0.5 rounded bg-muted">{e.task_status}</span>
            )}
            {e.priority && (
              <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">{e.priority}</span>
            )}
          </div>
          {e.objective && <Field label="Objective" value={e.objective} highlight />}
          {e.details && <Field label="Details" value={e.details} />}
          {e.acceptance && <Field label="Acceptance" value={e.acceptance} muted />}
          {e.plan_ref && (
            <div className="text-xs text-muted-foreground">
              Plan: <span className="text-foreground">{e.plan_ref}</span>
              {e.plan_phase && <> · phase: <span className="text-foreground">{e.plan_phase}</span></>}
            </div>
          )}
          {e.block_reason && <Field label="Block Reason" value={e.block_reason} muted />}
        </div>
      );
    default:
      return null;
  }
}

function ParamTable({ label, params }: { label: string; params: { path: string; type?: string; description?: string }[] }) {
  return (
    <div>
      <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">{label}</div>
      <div className="rounded border border-border bg-card divide-y divide-border">
        {params.map((p, i) => (
          <div key={i} className="p-2 text-xs">
            <div className="flex items-baseline gap-2">
              <span className="font-mono text-foreground">{p.path}</span>
              {p.type && (
                <span className="font-mono text-[10px] px-1.5 py-0.5 rounded bg-muted text-muted-foreground">{p.type}</span>
              )}
            </div>
            {p.description && (
              <div className="text-muted-foreground mt-0.5 leading-relaxed">{p.description}</div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

// Group an entity's relationships by (direction, type, target_kind) so
// "outbound references → 60 contracts" collapses into one line instead of
// rendering 60 chips. Singletons stay inline.
function groupRelationships(rels: Relationship[]): {
  key: string;
  direction: string;
  type: string;
  targetKind: string;
  items: Relationship[];
}[] {
  const groups: Record<string, Relationship[]> = {};
  for (const r of rels) {
    const direction = r.direction || 'outbound';
    const key = `${direction}|${r.type}|${r.target_kind}`;
    (groups[key] = groups[key] || []).push(r);
  }
  // Order: outbound first, belongs_to first within direction, then alpha by
  // (type, kind).
  const order: { d: string; rank: number }[] = [
    { d: 'outbound', rank: 0 },
    { d: 'inbound', rank: 1 },
  ];
  return Object.entries(groups)
    .map(([key, items]) => {
      const [direction, type, targetKind] = key.split('|');
      return { key, direction, type, targetKind, items };
    })
    .sort((a, b) => {
      const da = order.find((o) => o.d === a.direction)?.rank ?? 9;
      const db = order.find((o) => o.d === b.direction)?.rank ?? 9;
      if (da !== db) return da - db;
      if (a.type === 'belongs_to' && b.type !== 'belongs_to') return -1;
      if (b.type === 'belongs_to' && a.type !== 'belongs_to') return 1;
      if (a.type !== b.type) return a.type.localeCompare(b.type);
      return a.targetKind.localeCompare(b.targetKind);
    });
}

const COLLAPSE_THRESHOLD = 5;

const TYPE_LABEL: Record<string, string> = {
  belongs_to: 'belongs to',
  depends_on: 'depends on',
  exposes: 'exposes',
  consumes: 'consumes',
  uses: 'uses',
  involves: 'involves',
  references: 'references',
  relates_to: 'relates to',
  implements: 'implements',
  applies_to: 'applies to',
  modifies: 'modifies',
  visualizes: 'visualizes',
};

function RelationshipsSection({
  rels,
  onNavigate,
}: {
  rels: Relationship[];
  onNavigate: (slug: string, kind?: string) => void;
}) {
  const groups = groupRelationships(rels);
  return (
    <Section title={`Relationships · ${rels.length}`}>
      <div className="space-y-3">
        {groups.map((g) => (
          <RelationshipGroup key={g.key} group={g} onNavigate={onNavigate} />
        ))}
      </div>
    </Section>
  );
}

function RelationshipGroup({
  group,
  onNavigate,
}: {
  group: {
    direction: string;
    type: string;
    targetKind: string;
    items: Relationship[];
  };
  onNavigate: (slug: string, kind?: string) => void;
}) {
  const [expanded, setExpanded] = useState(group.items.length <= COLLAPSE_THRESHOLD);
  const Icon = iconForKind(group.targetKind);
  const directionArrow = group.direction === 'inbound' ? '←' : '→';
  const typeLabel = TYPE_LABEL[group.type] || group.type.replace(/_/g, ' ');
  const kindLabel = group.targetKind + (group.items.length === 1 ? '' : 's');
  const accent = ACCENT_FOR_KIND[group.targetKind] || 'text-muted-foreground';
  return (
    <div className="rounded-lg border border-border/60 bg-card/40 overflow-hidden">
      <button
        type="button"
        onClick={() => setExpanded((v: boolean) => !v)}
        className="flex items-center gap-2.5 w-full text-left px-3.5 py-2.5 hover:bg-card transition-colors"
      >
        <span className="text-muted-foreground text-xs tabular-nums w-3 text-center">
          {directionArrow}
        </span>
        {Icon && <Icon className={`w-4 h-4 ${accent}`} />}
        <span className="text-[11px] font-medium uppercase tracking-wider text-foreground/80">
          {group.direction === 'inbound' ? 'inbound ' : ''}{typeLabel}
        </span>
        <span className="text-[11px] text-muted-foreground tabular-nums">
          · {group.items.length} linked {kindLabel}
        </span>
        <span className="ml-auto text-[11px] text-muted-foreground">
          {expanded ? '▾' : '▸'}
        </span>
      </button>
      {expanded && (
        <div className="border-t border-border/60 px-3 py-2.5 flex flex-wrap gap-1.5">
          {group.items
            .slice()
            .sort((a, b) => (a.target_name || '').localeCompare(b.target_name || ''))
            .map((rel, i) => (
              <CompactRelPill
                key={i}
                rel={rel}
                onClick={() =>
                  onNavigate(
                    rel.target_slug || rel.target_name?.toLowerCase().replace(/\s+/g, '-') || rel.target_id,
                    rel.target_kind,
                  )
                }
              />
            ))}
        </div>
      )}
    </div>
  );
}

// Static palette per kind. Tailwind's JIT only picks up class names that
// appear as full literals in source, so the dynamic kind variants must be
// enumerated here rather than built via template literals at render time.
const PILL_PALETTE: Record<string, string> = {
  system:    'border-kind-system/30 bg-kind-system/5 hover:border-kind-system/60 hover:bg-kind-system/10 text-kind-system',
  component: 'border-kind-component/30 bg-kind-component/5 hover:border-kind-component/60 hover:bg-kind-component/10 text-kind-component',
  contract:  'border-kind-contract/30 bg-kind-contract/5 hover:border-kind-contract/60 hover:bg-kind-contract/10 text-kind-contract',
  concept:   'border-kind-concept/30 bg-kind-concept/5 hover:border-kind-concept/60 hover:bg-kind-concept/10 text-kind-concept',
  flow:      'border-kind-flow/30 bg-kind-flow/5 hover:border-kind-flow/60 hover:bg-kind-flow/10 text-kind-flow',
  decision:  'border-kind-decision/30 bg-kind-decision/5 hover:border-kind-decision/60 hover:bg-kind-decision/10 text-kind-decision',
  plan:      'border-kind-plan/30 bg-kind-plan/5 hover:border-kind-plan/60 hover:bg-kind-plan/10 text-kind-plan',
  task:      'border-kind-task/30 bg-kind-task/5 hover:border-kind-task/60 hover:bg-kind-task/10 text-kind-task',
};

function CompactRelPill({ rel, onClick }: { rel: Relationship; onClick: () => void }) {
  const Icon = iconForKind(rel.target_kind);
  const palette = PILL_PALETTE[rel.target_kind] || 'border-border bg-card hover:border-muted-foreground text-muted-foreground';
  return (
    <button
      type="button"
      onClick={onClick}
      title={`${rel.type}: ${rel.target_name || rel.target_id}`}
      className={`inline-flex items-center gap-2 rounded-md border px-2.5 py-1.5 transition-colors ${palette}`}
    >
      {Icon && <Icon className="w-3.5 h-3.5 shrink-0" />}
      <span className="text-[11px] leading-none text-foreground">
        {rel.target_name || rel.target_id}
      </span>
    </button>
  );
}

const ACCENT_FOR_KIND: Record<string, string> = {
  system: 'text-kind-system',
  component: 'text-kind-component',
  contract: 'text-kind-contract',
  concept: 'text-kind-concept',
  flow: 'text-kind-flow',
  decision: 'text-kind-decision',
  plan: 'text-kind-plan',
  task: 'text-kind-task',
};

// Order in which child kinds appear under a system. Sub-systems first
// (most architectural), then components, contracts, concepts.
//
// Static palette per kind because Tailwind's JIT can't scan dynamic
// `text-kind-${kind}` template strings — every class string must appear
// literal in source.
const SYSTEM_CHILD_KINDS: {
  kind: string;
  label: string;
  cardClass: string;
  iconClass: string;
}[] = [
  {
    kind: 'system',
    label: 'Sub-systems',
    cardClass: 'border-kind-system/30 bg-kind-system/5 hover:border-kind-system/60',
    iconClass: 'text-kind-system',
  },
  {
    kind: 'component',
    label: 'Components',
    cardClass: 'border-kind-component/30 bg-kind-component/5 hover:border-kind-component/60',
    iconClass: 'text-kind-component',
  },
  {
    kind: 'contract',
    label: 'Contracts',
    cardClass: 'border-kind-contract/30 bg-kind-contract/5 hover:border-kind-contract/60',
    iconClass: 'text-kind-contract',
  },
  {
    kind: 'concept',
    label: 'Concepts',
    cardClass: 'border-kind-concept/30 bg-kind-concept/5 hover:border-kind-concept/60',
    iconClass: 'text-kind-concept',
  },
];

function SystemChildren({
  relationships,
  parentSlug,
  onNavigate,
}: {
  relationships: Relationship[];
  parentSlug: string;
  onNavigate: (slug: string, kind?: string, filter?: string) => void;
}) {
  // Children of a system are inbound `belongs_to` relationships from
  // sub-systems, components, contracts, concepts.
  const children = relationships.filter(
    (r) => r.direction === 'inbound' && r.type === 'belongs_to',
  );
  if (children.length === 0) return null;

  const grouped: Record<string, Relationship[]> = {};
  for (const c of children) {
    (grouped[c.target_kind] = grouped[c.target_kind] || []).push(c);
  }

  return (
    <div className="space-y-4">
      {SYSTEM_CHILD_KINDS.map(({ kind, label, cardClass, iconClass }) => {
        const items = grouped[kind];
        if (!items || items.length === 0) return null;
        const Icon = iconForKind(kind);
        return (
          <div key={kind}>
            <div className="flex items-center gap-2 mb-2">
              {Icon && <Icon className={`w-3.5 h-3.5 ${iconClass}`} />}
              <h3 className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest">
                {label}
              </h3>
              <span className="text-[10px] text-muted-foreground">{items.length}</span>
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-2">
              {items
                .slice()
                .sort((a, b) => (a.target_name || '').localeCompare(b.target_name || ''))
                .map((c, i) => (
                  <button
                    key={i}
                    onClick={() =>
                      onNavigate(
                        c.target_slug || c.target_name?.toLowerCase().replace(/\s+/g, '-') || c.target_id,
                        c.target_kind,
                        // Scope the destination kind page to only entities
                        // that share this parent system.
                        `belongs_to:${parentSlug}`,
                      )
                    }
                    className={`group flex items-center gap-2.5 rounded-lg border ${cardClass} transition-colors px-3 py-2 text-left min-w-0`}
                  >
                    {Icon && (
                      <span className={`flex items-center justify-center w-6 h-6 rounded-md ${iconClass} shrink-0`}>
                        <Icon className="w-3.5 h-3.5" />
                      </span>
                    )}
                    <span className="text-xs font-medium text-foreground truncate min-w-0">
                      {c.target_name || c.target_id}
                    </span>
                  </button>
                ))}
            </div>
          </div>
        );
      })}
    </div>
  );
}

// Strip the 4-char alphanumeric suffix from a slug so it matches the bare
// target form stored in entity relationships (`syde`, not `syde-5tdt`).
// Mirrors internal/utils/slug.go BaseSlug.
function baseSlug(slug: string): string {
  return slug.replace(/-[a-z0-9]{4}$/, '');
}

function Field({ label, value, highlight, muted, code }: { label: string; value: string; highlight?: boolean; muted?: boolean; code?: boolean }) {
  return (
    <div>
      <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">{label}</div>
      <div
        className={`text-xs leading-relaxed rounded p-2 ${
          highlight ? 'bg-card border border-border' :
          muted ? 'bg-muted/50 text-muted-foreground' :
          code ? 'font-mono bg-card border border-border' :
          'text-muted-foreground'
        }`}
      >
        {value}
      </div>
    </div>
  );
}
