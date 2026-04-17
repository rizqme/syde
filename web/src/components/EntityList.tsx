import { useEffect, useMemo, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { EntitySummary } from '../lib/api';
import { KindBadge } from './KindBadge';
import { EntityFilterBar, applyFilter } from './EntityFilterBar';
import { EyeIcon, EyeOffIcon } from './icons';

interface EntityListProps {
  entities: EntitySummary[];
  kind: string;
  onSelect: (slug: string) => void;
  selectedSlug?: string;
}

const STATUS_ORDER: Record<string, number> = {
  in_progress: 0,
  'in-progress': 0,
  approved: 1,
  draft: 2,
  completed: 3,
};

const STATUS_DOT: Record<string, string> = {
  draft: 'bg-zinc-500',
  approved: 'bg-blue-400',
  'in-progress': 'bg-yellow-400',
  in_progress: 'bg-yellow-400',
  completed: 'bg-green-400',
};

export function EntityList({ entities, kind, onSelect, selectedSlug }: EntityListProps) {
  const location = useLocation();
  const isPlan = kind === 'plan';
  const initialFilter = useMemo(() => {
    const params = new URLSearchParams(location.search);
    return params.get('filter') || '';
  }, [location.search]);
  const [filter, setFilter] = useState(initialFilter);
  const [hideCompleted, setHideCompleted] = useState(false);
  useEffect(() => {
    setFilter(initialFilter);
  }, [initialFilter]);

  const filtered = useMemo(() => {
    let list = applyFilter(entities, filter);
    if (isPlan && hideCompleted) {
      list = list.filter((e) => e.plan_status !== 'completed');
    }
    if (isPlan) {
      list = [...list].sort((a, b) => {
        const sa = STATUS_ORDER[a.plan_status || 'draft'] ?? 2;
        const sb = STATUS_ORDER[b.plan_status || 'draft'] ?? 2;
        if (sa !== sb) return sa - sb;
        const da = a.updated_at || '';
        const db = b.updated_at || '';
        return db.localeCompare(da);
      });
    }
    return list;
  }, [entities, filter, isPlan, hideCompleted]);

  const kindLabel = kind.charAt(0).toUpperCase() + kind.slice(1) + 's';

  return (
    <div>
      <div className="mb-4 flex items-baseline justify-between gap-3">
        <h1 className="text-xl font-semibold tracking-tight">{kindLabel}</h1>
        <div className="flex items-center gap-3">
          {isPlan && (
            <button
              onClick={() => setHideCompleted((h) => !h)}
              className={`p-1.5 rounded transition-colors ${
                hideCompleted
                  ? 'text-foreground'
                  : 'text-muted-foreground hover:text-foreground'
              }`}
              title={hideCompleted ? 'Show completed plans' : 'Hide completed plans'}
            >
              {hideCompleted ? (
                <EyeOffIcon className="w-4 h-4" />
              ) : (
                <EyeIcon className="w-4 h-4" />
              )}
            </button>
          )}
          <p className="text-muted-foreground text-xs shrink-0">
            {filtered.length}
            {filtered.length !== entities.length && (
              <span className="text-muted-foreground/60"> / {entities.length}</span>
            )}
          </p>
        </div>
      </div>

      {entities.length > 5 && (
        <div className="mb-3">
          <EntityFilterBar
            value={filter}
            onChange={setFilter}
            entities={entities}
            placeholder={`Filter ${kindLabel.toLowerCase()} — e.g. belongs_to:syde-cli, tag:internal, -has:files`}
          />
        </div>
      )}

      <div className="space-y-1.5">
        {filtered.map((entity) => {
          const slug = entity.slug || entity.name.toLowerCase().replace(/\s+/g, '-');
          const isSelected = selectedSlug === slug;
          return (
            <button
              key={entity.id}
              onClick={() => onSelect(slug)}
              className={`w-full text-left px-4 py-3 rounded-lg border transition-colors ${
                isSelected
                  ? 'bg-accent border-accent-foreground/20'
                  : 'bg-card border-border hover:border-muted-foreground'
              }`}
            >
              <div className="flex items-start gap-2.5">
                {isPlan && (
                  <span className={`mt-1.5 w-2 h-2 rounded-full shrink-0 ${STATUS_DOT[entity.plan_status || 'draft'] || 'bg-zinc-500'}`} />
                )}
                <div className="min-w-0 flex-1">
                  <div className="flex items-center gap-2">
                    <span className="text-sm font-medium text-foreground truncate">{entity.name}</span>
                    {isPlan && entity.plan_status && (
                      <span className="text-[10px] text-muted-foreground shrink-0">{entity.plan_status}</span>
                    )}
                    <span className="ml-auto shrink-0 flex items-center gap-2">
                      {isPlan && entity.updated_at && (
                        <span className="text-[10px] text-muted-foreground/60">{formatRelative(entity.updated_at)}</span>
                      )}
                      {!isPlan && <KindBadge kind={entity.kind} />}
                    </span>
                  </div>
                  {entity.description && (
                    <p className="text-xs text-muted-foreground line-clamp-1 mt-0.5">
                      {entity.description}
                    </p>
                  )}
                  {!isPlan && entity.relationship_count > 0 && (
                    <div className="flex gap-3 mt-1 text-[10px] text-muted-foreground">
                      <span>{entity.relationship_count} rels</span>
                    </div>
                  )}
                </div>
              </div>
            </button>
          );
        })}
        {filtered.length === 0 && (
          <div className="text-center py-12 text-muted-foreground text-sm">
            {filter
              ? 'No matches'
              : hideCompleted
                ? 'All plans are completed.'
                : `No ${kindLabel.toLowerCase()} yet.`}
          </div>
        )}
      </div>
    </div>
  );
}

function formatRelative(iso: string): string {
  try {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return '';
    const now = Date.now();
    const diff = now - d.getTime();
    const mins = Math.floor(diff / 60000);
    if (mins < 1) return 'just now';
    if (mins < 60) return `${mins}m ago`;
    const hours = Math.floor(mins / 60);
    if (hours < 24) return `${hours}h ago`;
    const days = Math.floor(hours / 24);
    if (days < 30) return `${days}d ago`;
    return d.toLocaleDateString(undefined, { month: 'short', day: 'numeric' });
  } catch {
    return '';
  }
}
