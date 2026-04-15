import { useEffect, useMemo, useState } from 'react';
import { api, EntitySummary } from '../lib/api';
import { KindBadge } from './KindBadge';
import { EntityFilterBar, applyFilter } from './EntityFilterBar';

interface SearchPaletteProps {
  open: boolean;
  onClose: () => void;
  onSelect: (slug: string, kind?: string) => void;
}

const RESULT_LIMIT = 60;

// Order kinds appear in the grouped result list — architectural first,
// workflow second.
const KIND_ORDER = [
  'system',
  'component',
  'contract',
  'concept',
  'flow',
  'decision',
  'plan',
  'task',
  'design',
  'learning',
];

export function SearchPalette({ open, onClose, onSelect }: SearchPaletteProps) {
  const [query, setQuery] = useState('');
  const [entities, setEntities] = useState<EntitySummary[]>([]);
  const [selected, setSelected] = useState(0);

  // Fetch every entity once whenever the palette opens. Client-side
  // filtering via the EntityFilterBar DSL gives us the full key:value
  // grammar (kind:, belongs_to:, etc.) without round-tripping each query.
  useEffect(() => {
    if (!open) return;
    setQuery('');
    setSelected(0);
    let cancelled = false;
    api
      .entities()
      .then((res) => {
        if (cancelled) return;
        setEntities(res.entities || []);
      })
      .catch(() => {
        if (cancelled) return;
        setEntities([]);
      });
    return () => {
      cancelled = true;
    };
  }, [open]);

  const filtered = useMemo(() => {
    if (!query.trim()) return [] as EntitySummary[];
    return applyFilter(entities, query).slice(0, RESULT_LIMIT);
  }, [entities, query]);

  // Group filtered hits by kind, preserving KIND_ORDER.
  const grouped = useMemo(() => {
    const map: Record<string, EntitySummary[]> = {};
    for (const e of filtered) {
      (map[e.kind] = map[e.kind] || []).push(e);
    }
    return KIND_ORDER.filter((k) => map[k]?.length).map((k) => ({ kind: k, hits: map[k] }));
  }, [filtered]);

  // Flat index for keyboard navigation across groups.
  const flat = useMemo(() => grouped.flatMap((g) => g.hits), [grouped]);

  useEffect(() => {
    setSelected((s) => Math.min(s, Math.max(0, flat.length - 1)));
  }, [flat.length]);

  if (!open) return null;

  function commit(hit: EntitySummary) {
    const slug = hit.slug || hit.name.toLowerCase().replace(/\s+/g, '-');
    onSelect(slug, hit.kind);
    onClose();
  }

  function handleKey(ev: React.KeyboardEvent) {
    if (ev.key === 'Escape') {
      onClose();
      return;
    }
    if (flat.length === 0) return;
    if (ev.key === 'ArrowDown') {
      ev.preventDefault();
      setSelected((s) => (s + 1) % flat.length);
    } else if (ev.key === 'ArrowUp') {
      ev.preventDefault();
      setSelected((s) => (s - 1 + flat.length) % flat.length);
    } else if (ev.key === 'Enter') {
      ev.preventDefault();
      const hit = flat[selected];
      if (hit) commit(hit);
    }
  }

  let runningIdx = -1;

  return (
    <div
      className="fixed inset-0 z-50 flex items-start justify-center pt-[16vh] bg-background/50 backdrop-blur-md"
      onClick={onClose}
      onKeyDown={handleKey}
    >
      <div
        className="relative w-[720px] max-h-[68vh] flex flex-col bg-black border border-white/10 rounded-[28px] shadow-[0_30px_90px_-15px_rgba(0,0,0,0.95),0_0_0_1px_rgba(255,255,255,0.05),0_0_60px_rgba(0,0,0,0.4)] overflow-hidden"
        onClick={(e) => e.stopPropagation()}
      >
        {/* Filter bar uses the same DSL as the per-kind list pages. The
            palette uses the `large` variant — borderless, big text, the bar
            IS the chrome of the card. */}
        <EntityFilterBar
          value={query}
          onChange={setQuery}
          entities={entities}
          large
          autoFocus
          placeholder="search entities — name, kind:component, belongs_to:syde-cli …"
        />

        {/* Results — only rendered when there's a query. Empty input
            collapses the card down to a single pill-shaped search bar. */}
        {query.trim() && (
          <>
            <div className="overflow-y-auto flex-1 border-t border-white/10">
              {grouped.length === 0 ? (
                <div className="px-5 py-10 text-center text-muted-foreground text-xs">
                  No matches.
                </div>
              ) : (
                grouped.map(({ kind, hits }) => (
                  <div key={kind}>
                    <div className="px-5 py-1.5 text-[9px] font-medium text-muted-foreground uppercase tracking-widest bg-white/[0.02] sticky top-0">
                      {kind}s · {hits.length}
                    </div>
                    {hits.map((hit) => {
                      runningIdx++;
                      const idx = runningIdx;
                      const isSel = idx === selected;
                      return (
                        <button
                          key={hit.id}
                          onMouseEnter={() => setSelected(idx)}
                          onClick={() => commit(hit)}
                          className={`w-full text-left px-5 py-2.5 flex items-start gap-3 ${
                            isSel ? 'bg-white/[0.06]' : 'hover:bg-white/[0.03]'
                          }`}
                        >
                          <KindBadge kind={hit.kind} />
                          <div className="min-w-0 flex-1">
                            <div className="text-[13px] font-medium text-foreground truncate">
                              {hit.name}
                            </div>
                            {hit.description && (
                              <div className="text-[11px] text-muted-foreground truncate mt-0.5">
                                {hit.description}
                              </div>
                            )}
                          </div>
                        </button>
                      );
                    })}
                  </div>
                ))
              )}
            </div>
            <div className="px-5 py-2 border-t border-white/10 text-[10px] text-muted-foreground flex justify-between">
              <span>↑↓ navigate · ⏎ open · esc close</span>
              <span>{filtered.length === RESULT_LIMIT ? `showing ${RESULT_LIMIT}+` : ''}</span>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
