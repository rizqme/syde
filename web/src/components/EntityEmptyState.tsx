import { iconForKind } from './icons';

// Per-kind accent for the illustration glow + featured card stroke. Keeps
// the empty state on-theme with whichever list the user is browsing.
const KIND_ACCENT: Record<string, { stroke: string; glow: string; icon: string }> = {
  system:    { stroke: 'stroke-kind-system',    glow: 'fill-kind-system/10',    icon: 'text-kind-system' },
  component: { stroke: 'stroke-kind-component', glow: 'fill-kind-component/10', icon: 'text-kind-component' },
  contract:  { stroke: 'stroke-kind-contract',  glow: 'fill-kind-contract/10',  icon: 'text-kind-contract' },
  concept:   { stroke: 'stroke-kind-concept',   glow: 'fill-kind-concept/10',   icon: 'text-kind-concept' },
  flow:      { stroke: 'stroke-kind-flow',      glow: 'fill-kind-flow/10',      icon: 'text-kind-flow' },
  decision:  { stroke: 'stroke-kind-decision',  glow: 'fill-kind-decision/10',  icon: 'text-kind-decision' },
  plan:      { stroke: 'stroke-kind-plan',      glow: 'fill-kind-plan/10',      icon: 'text-kind-plan' },
  task:      { stroke: 'stroke-kind-task',      glow: 'fill-kind-task/10',      icon: 'text-kind-task' },
  learning:  { stroke: 'stroke-kind-learning',  glow: 'fill-kind-learning/10',  icon: 'text-kind-learning' },
};

const FALLBACK = { stroke: 'stroke-muted-foreground', glow: 'fill-muted/40', icon: 'text-muted-foreground' };

export function EntityEmptyState({ kind }: { kind: string }) {
  const Icon = iconForKind(kind);
  const accent = KIND_ACCENT[kind] || FALLBACK;
  const label = kind.charAt(0).toUpperCase() + kind.slice(1);

  return (
    <div className="h-full flex flex-col items-center justify-center px-8 text-center">
      {/* Stacked-cards illustration. Two muted background cards + a
          highlighted "selected" card; the kind icon floats on top via an
          absolutely-positioned wrapper so we don't have to embed React
          components via <foreignObject>. */}
      <div className="relative mb-8 w-[200px] h-[160px] opacity-90">
      <svg
        width="200"
        height="160"
        viewBox="0 0 200 160"
        fill="none"
        className="absolute inset-0"
        aria-hidden="true"
      >
        {/* Soft glow behind the highlighted card */}
        <ellipse cx="100" cy="135" rx="70" ry="6" className="fill-foreground/5" />

        {/* Background card #1 (furthest) */}
        <rect
          x="32"
          y="22"
          width="120"
          height="76"
          rx="10"
          className="stroke-border fill-card"
          strokeWidth="1.5"
        />
        <rect x="44" y="38" width="60" height="5" rx="2" className="fill-muted" />
        <rect x="44" y="50" width="90" height="3" rx="1.5" className="fill-muted/60" />
        <rect x="44" y="58" width="76" height="3" rx="1.5" className="fill-muted/60" />

        {/* Background card #2 */}
        <rect
          x="42"
          y="36"
          width="120"
          height="76"
          rx="10"
          className="stroke-border fill-card"
          strokeWidth="1.5"
        />
        <rect x="54" y="52" width="48" height="5" rx="2" className="fill-muted" />
        <rect x="54" y="64" width="92" height="3" rx="1.5" className="fill-muted/60" />
        <rect x="54" y="72" width="70" height="3" rx="1.5" className="fill-muted/60" />

        {/* Foreground (highlighted) card with kind accent */}
        <rect
          x="22"
          y="56"
          width="156"
          height="86"
          rx="12"
          className={`${accent.stroke} ${accent.glow}`}
          strokeWidth="2"
        />
        {/* Icon badge inside the foreground card */}
        <rect
          x="36"
          y="72"
          width="32"
          height="32"
          rx="8"
          className={`${accent.stroke}`}
          strokeWidth="1.5"
          fill="none"
        />
        {/* Title bar */}
        <rect x="80" y="78" width="72" height="6" rx="2" className="fill-foreground/70" />
        {/* Subtitle bars */}
        <rect x="80" y="92" width="92" height="3" rx="1.5" className="fill-muted-foreground/60" />
        <rect x="80" y="100" width="64" height="3" rx="1.5" className="fill-muted-foreground/60" />

        {/* Connector dots from the list (left) into the card */}
        <circle cx="8" cy="100" r="2" className="fill-muted-foreground/40" />
        <circle cx="14" cy="100" r="2" className="fill-muted-foreground/60" />
        <circle cx="20" cy="100" r="2" className={accent.icon.replace('text-', 'fill-')} />
      </svg>
        {/* Kind icon overlay positioned on top of the badge rect at (36,72)
            in 200x160 viewBox space. Coordinates: badge centre is (52,88) →
            in container px that's (52,88) since the SVG fills the box. */}
        {Icon && (
          <div
            className={`absolute ${accent.icon} pointer-events-none`}
            style={{ left: 42, top: 78, width: 20, height: 20 }}
          >
            <Icon className="w-5 h-5" />
          </div>
        )}
      </div>

      <h3 className="text-base font-medium text-foreground">
        Pick a {label.toLowerCase()} to inspect
      </h3>
      <p className="text-sm text-muted-foreground mt-2 max-w-sm">
        Choose an entry from the list on the left to see its full details, files,
        relationships, and learnings.
      </p>
    </div>
  );
}
