import { Relationship } from '../lib/api';
import { iconForKind } from './icons';

// Per-kind palette: subtle tinted background + matching border + accent dot.
// Keeps the kind connection visible without screaming colour at full saturation.
const KIND_PALETTE: Record<string, { border: string; bg: string; ring: string; accent: string; text: string }> = {
  system:    { border: 'border-kind-system/30',    bg: 'bg-kind-system/5',    ring: 'hover:border-kind-system/60',    accent: 'text-kind-system',    text: 'text-foreground' },
  component: { border: 'border-kind-component/30', bg: 'bg-kind-component/5', ring: 'hover:border-kind-component/60', accent: 'text-kind-component', text: 'text-foreground' },
  contract:  { border: 'border-kind-contract/30',  bg: 'bg-kind-contract/5',  ring: 'hover:border-kind-contract/60',  accent: 'text-kind-contract',  text: 'text-foreground' },
  concept:   { border: 'border-kind-concept/30',   bg: 'bg-kind-concept/5',   ring: 'hover:border-kind-concept/60',   accent: 'text-kind-concept',   text: 'text-foreground' },
  flow:      { border: 'border-kind-flow/30',      bg: 'bg-kind-flow/5',      ring: 'hover:border-kind-flow/60',      accent: 'text-kind-flow',      text: 'text-foreground' },
  decision:  { border: 'border-kind-decision/30',  bg: 'bg-kind-decision/5',  ring: 'hover:border-kind-decision/60',  accent: 'text-kind-decision',  text: 'text-foreground' },
  plan:      { border: 'border-kind-plan/30',      bg: 'bg-kind-plan/5',      ring: 'hover:border-kind-plan/60',      accent: 'text-kind-plan',      text: 'text-foreground' },
  task:      { border: 'border-kind-task/30',      bg: 'bg-kind-task/5',      ring: 'hover:border-kind-task/60',      accent: 'text-kind-task',      text: 'text-foreground' },
  learning:  { border: 'border-kind-learning/30',  bg: 'bg-kind-learning/5',  ring: 'hover:border-kind-learning/60',  accent: 'text-kind-learning',  text: 'text-foreground' },
};

const FALLBACK = { border: 'border-border', bg: 'bg-card', ring: 'hover:border-muted-foreground', accent: 'text-muted-foreground', text: 'text-foreground' };

// Pretty-print a relationship type for display: belongs_to → "belongs to".
function prettyType(type: string): string {
  return type.replace(/_/g, ' ');
}

interface RelationshipChipProps {
  rel: Relationship;
  onClick: () => void;
}

export function RelationshipChip({ rel, onClick }: RelationshipChipProps) {
  const Icon = iconForKind(rel.target_kind);
  const p = KIND_PALETTE[rel.target_kind] || FALLBACK;
  const name = rel.target_name || rel.target_id;

  return (
    <button
      onClick={onClick}
      title={`${rel.type}: ${name}`}
      className={`group inline-flex items-center gap-2.5 pl-2.5 pr-3 py-1.5 rounded-lg border transition-colors ${p.border} ${p.bg} ${p.ring}`}
    >
      {Icon && (
        <span className={`flex items-center justify-center w-5 h-5 rounded-md ${p.accent}`}>
          <Icon className="w-3.5 h-3.5" />
        </span>
      )}
      <span className="flex flex-col items-start leading-tight min-w-0">
        <span className={`text-[9px] font-medium uppercase tracking-wider ${p.accent}`}>
          {prettyType(rel.type)}
        </span>
        <span className={`text-[11px] font-medium truncate ${p.text}`}>
          {name}
        </span>
      </span>
    </button>
  );
}
