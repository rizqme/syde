// Top-level view for PlanDetail.changes. Renders kind tabs (Requirements /
// Systems / Concepts / Components / Contracts / Flows); tabs whose lane is
// entirely empty are hidden. Each tab shows Deleted / Extended / New
// subsections in that order. Active tab is local state — the outer
// Plan/Tasks URL tab is managed by PlanDetailPanel.

import { useMemo, useState } from 'react';
import {
  ChangeLane,
  DeletedChange,
  ExtendedChange,
  NewChange,
  PlanChanges,
} from '../lib/api';
import { ExtendedFieldDiff } from './ExtendedFieldDiff';
import { NewEntityDraftView } from './NewEntityDraftView';

interface PlanChangesViewProps {
  changes: PlanChanges;
}

// The tab key is the plural lane name so we can read lanes off PlanChanges
// directly. The singular form is used when handing the lane to
// NewEntityDraftView, which dispatches on the entity kind.
type LaneKey = 'requirements' | 'systems' | 'concepts' | 'components' | 'contracts' | 'flows';

const LANE_ORDER: { key: LaneKey; label: string; singular: string }[] = [
  { key: 'requirements', label: 'Requirements', singular: 'requirement' },
  { key: 'systems', label: 'Systems', singular: 'system' },
  { key: 'concepts', label: 'Concepts', singular: 'concept' },
  { key: 'components', label: 'Components', singular: 'component' },
  { key: 'contracts', label: 'Contracts', singular: 'contract' },
  { key: 'flows', label: 'Flows', singular: 'flow' },
];

export function PlanChangesView({ changes }: PlanChangesViewProps) {
  const visible = useMemo(() => {
    return LANE_ORDER
      .map((l) => ({ ...l, lane: (changes?.[l.key] ?? {}) as ChangeLane }))
      .filter((l) => laneCount(l.lane).total > 0);
  }, [changes]);

  const [active, setActive] = useState<LaneKey | null>(
    visible.length > 0 ? visible[0].key : null,
  );

  if (visible.length === 0) {
    return (
      <div className="rounded border border-dashed border-border bg-card/30 px-4 py-8 text-center">
        <div className="text-sm text-muted-foreground">No changes declared yet</div>
        <div className="text-xs text-muted-foreground/70 mt-1">
          Use syde plan add-change to record deletions, extensions, or new entities.
        </div>
      </div>
    );
  }

  // If the previously-active tab disappeared (shouldn't happen but guard
  // anyway), fall back to the first visible tab.
  const activeEntry =
    visible.find((l) => l.key === active) ?? visible[0];

  return (
    <div>
      <div className="flex flex-wrap items-center gap-1 mb-4 border-b border-border">
        {visible.map((l) => {
          const counts = laneCount(l.lane);
          const isActive = activeEntry.key === l.key;
          return (
            <button
              key={l.key}
              onClick={() => setActive(l.key)}
              className={`px-3 py-1.5 text-xs font-medium border-b-2 transition-colors -mb-px ${
                isActive
                  ? 'border-foreground text-foreground'
                  : 'border-transparent text-muted-foreground hover:text-foreground'
              }`}
            >
              {l.label}
              <span className="ml-1.5 text-[10px] text-muted-foreground">
                ({countsLabel(counts)})
              </span>
            </button>
          );
        })}
      </div>

      <LaneView lane={activeEntry.lane} kindSingular={activeEntry.singular} />
    </div>
  );
}

function LaneView({ lane, kindSingular }: { lane: ChangeLane; kindSingular: string }) {
  const deleted = lane.deleted || [];
  const extended = lane.extended || [];
  const added = lane.new || [];
  const groupedExtended = groupExtendedBySlug(extended);

  return (
    <div className="space-y-6">
      {deleted.length > 0 && (
        <Subsection title={`Deleted (${deleted.length})`}>
          <div className="space-y-2">
            {deleted.map((d) => (
              <DeletedCard key={d.id} change={d} />
            ))}
          </div>
        </Subsection>
      )}
      {extended.length > 0 && (
        <Subsection title={`Extended (${extended.length})`}>
          <div className="space-y-2">
            {groupedExtended.map((group) => (
              <ExtendedCard key={group.slug} slug={group.slug} changes={group.changes} />
            ))}
          </div>
        </Subsection>
      )}
      {added.length > 0 && (
        <Subsection title={`New (${added.length})`}>
          <div className="space-y-2">
            {added.map((n) => (
              <NewCard key={n.id} change={n} kindSingular={kindSingular} />
            ))}
          </div>
        </Subsection>
      )}
    </div>
  );
}

function DeletedCard({ change }: { change: DeletedChange }) {
  return (
    <div className="rounded border border-red-600/30 bg-red-900/10 p-3 space-y-1.5">
      <div className="flex items-center gap-2">
        <span className="text-[10px] uppercase tracking-wider px-1.5 py-0.5 rounded bg-red-900/30 text-red-300">
          delete
        </span>
        <code className="font-mono text-xs text-foreground">{change.slug}</code>
      </div>
      {change.why && (
        <div className="text-xs text-muted-foreground leading-relaxed">
          <span className="uppercase tracking-widest text-[10px] mr-1">Why:</span>
          {change.why}
        </div>
      )}
    </div>
  );
}

function ExtendedCard({
  slug,
  changes,
}: {
  slug: string;
  changes: ExtendedChange[];
}) {
  const fields: Record<string, string> = {};
  const currentValues: Record<string, any> = {};
  const proposedValuesHTML: Record<string, string> = {};
  for (const change of changes) {
    Object.assign(fields, change.field_changes || {});
    Object.assign(currentValues, change.current_values || {});
    Object.assign(proposedValuesHTML, change.proposed_values_html || {});
  }
  const hasFields = Object.keys(fields).length > 0;
  const wireframeHTML = currentValues.wireframe_html;
  return (
    <div className="rounded border border-amber-600/30 bg-yellow-900/10 p-3 space-y-2">
      <div className="flex items-center gap-2">
        <span className="text-[10px] uppercase tracking-wider px-1.5 py-0.5 rounded bg-amber-900/30 text-amber-300">
          extend
        </span>
        <code className="font-mono text-xs text-foreground">{slug}</code>
        {changes.length > 1 && (
          <span className="text-[10px] text-muted-foreground">
            {changes.length} entries
          </span>
        )}
      </div>
      {wireframeHTML && (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Current wireframe
          </div>
          <div
            className="rounded border border-border bg-card p-3 overflow-x-auto text-xs"
            dangerouslySetInnerHTML={{ __html: wireframeHTML }}
          />
        </div>
      )}
      {changes.map((change) => (
        <div
          key={change.id}
          className={changes.length > 1 ? 'rounded border border-border/60 bg-card/30 p-2 space-y-1' : 'space-y-1'}
        >
          {changes.length > 1 && (
            <div className="text-[10px] font-mono text-muted-foreground">
              change {change.id}
            </div>
          )}
          {change.what && (
            <div className="text-xs text-foreground leading-relaxed font-medium">
              {change.what}
            </div>
          )}
          {change.why && (
            <div className="text-xs text-muted-foreground leading-relaxed">
              <span className="uppercase tracking-widest text-[10px] mr-1">Why:</span>
              {change.why}
            </div>
          )}
        </div>
      ))}
      {hasFields && (
        <ExtendedFieldDiff
          currentValues={currentValues}
          fieldChanges={fields}
          proposedValuesHTML={proposedValuesHTML}
        />
      )}
    </div>
  );
}

function NewCard({
  change,
  kindSingular,
}: {
  change: NewChange;
  kindSingular: string;
}) {
  return (
    <div className="rounded border border-green-600/30 bg-green-900/10 p-3 space-y-2">
      <div className="flex items-center gap-2">
        <span className="text-[10px] uppercase tracking-wider px-1.5 py-0.5 rounded bg-green-900/30 text-green-300">
          new
        </span>
        <h4 className="text-sm font-semibold text-foreground">{change.name}</h4>
      </div>
      {change.what && (
        <div className="text-xs text-foreground leading-relaxed font-medium">
          {change.what}
        </div>
      )}
      {change.why && (
        <div className="text-xs text-muted-foreground leading-relaxed">
          <span className="uppercase tracking-widest text-[10px] mr-1">Why:</span>
          {change.why}
        </div>
      )}
      {change.draft && (
        <div className="pt-1">
          <NewEntityDraftView kind={kindSingular} draft={change.draft} />
        </div>
      )}
    </div>
  );
}

function Subsection({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <div>
      <h3 className="text-[11px] font-medium text-muted-foreground uppercase tracking-widest mb-2">
        {title}
      </h3>
      {children}
    </div>
  );
}

function laneCount(lane: ChangeLane): { del: number; ext: number; add: number; total: number } {
  const del = lane?.deleted?.length || 0;
  const ext = lane?.extended?.length || 0;
  const add = lane?.new?.length || 0;
  return { del, ext, add, total: del + ext + add };
}

function groupExtendedBySlug(changes: ExtendedChange[]): { slug: string; changes: ExtendedChange[] }[] {
  const order: string[] = [];
  const groups: Record<string, ExtendedChange[]> = {};
  for (const change of changes) {
    if (!groups[change.slug]) {
      groups[change.slug] = [];
      order.push(change.slug);
    }
    groups[change.slug].push(change);
  }
  return order.map((slug) => ({ slug, changes: groups[slug] }));
}

function countsLabel(c: { del: number; ext: number; add: number }): string {
  const parts: string[] = [];
  if (c.del) parts.push(`${c.del} del`);
  if (c.ext) parts.push(`${c.ext} ext`);
  if (c.add) parts.push(`${c.add} new`);
  return parts.join(', ');
}
