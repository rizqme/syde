import { useEffect, useMemo, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { EntitySummary } from '../lib/api';
import { KindBadge } from './KindBadge';
import { EntityFilterBar, applyFilter } from './EntityFilterBar';

interface EntityListProps {
  entities: EntitySummary[];
  kind: string;
  onSelect: (slug: string) => void;
  selectedSlug?: string;
}

export function EntityList({ entities, kind, onSelect, selectedSlug }: EntityListProps) {
  const location = useLocation();
  // Seed the filter from ?filter=… so a child clicked from a parent system
  // lands here with the list already scoped to siblings.
  const initialFilter = useMemo(() => {
    const params = new URLSearchParams(location.search);
    return params.get('filter') || '';
  }, [location.search]);
  const [filter, setFilter] = useState(initialFilter);
  // Re-seed when navigating between filtered URLs without remounting
  // (e.g. clicking a different child while still on the same kind page).
  useEffect(() => {
    setFilter(initialFilter);
  }, [initialFilter]);
  const filtered = useMemo(() => applyFilter(entities, filter), [entities, filter]);

  const kindLabel = kind.charAt(0).toUpperCase() + kind.slice(1) + 's';

  return (
    <div>
      <div className="mb-4 flex items-baseline justify-between gap-3">
        <h1 className="text-xl font-semibold tracking-tight">{kindLabel}</h1>
        <p className="text-muted-foreground text-xs shrink-0">
          {filtered.length}
          {filtered.length !== entities.length && (
            <span className="text-muted-foreground/60"> of {entities.length}</span>
          )}{' '}
          entities
        </p>
      </div>

      {entities.length > 5 && (
        <EntityFilterBar
          value={filter}
          onChange={setFilter}
          entities={entities}
          placeholder={`Filter ${kindLabel.toLowerCase()} — e.g. belongs_to:syde-cli, tag:internal, -has:files`}
        />
      )}

      <div className="space-y-2">
        {filtered.map((entity) => {
          const slug = entity.slug || entity.name.toLowerCase().replace(/\s+/g, '-');
          const isSelected = selectedSlug === slug;
          return (
            <button
              key={entity.id}
              onClick={() => onSelect(slug)}
              className={`w-full text-left p-4 rounded-lg border transition-colors text-foreground ${
                isSelected
                  ? 'bg-accent border-accent-foreground/20'
                  : 'bg-card border-border hover:border-muted-foreground'
              }`}
            >
              <div className="flex items-center gap-2 mb-1">
                <KindBadge kind={entity.kind} />
                <span className="text-sm font-medium text-foreground">{entity.name}</span>
              </div>
              {entity.description && (
                <p className="text-xs text-muted-foreground line-clamp-2 mt-1">
                  {entity.description}
                </p>
              )}
              <div className="flex gap-3 mt-2 text-[10px] text-muted-foreground">
                {entity.relationship_count > 0 && (
                  <span>{entity.relationship_count} rels</span>
                )}
                {entity.learning_count > 0 && (
                  <span>{entity.learning_count} learnings</span>
                )}
              </div>
            </button>
          );
        })}
        {filtered.length === 0 && (
          <div className="text-center py-12 text-muted-foreground text-sm">
            {filter ? 'No matches' : `No ${kindLabel.toLowerCase()} yet.`}
          </div>
        )}
      </div>
    </div>
  );
}
