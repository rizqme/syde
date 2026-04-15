const KIND_COLORS: Record<string, string> = {
  system: 'border-foreground/30 text-foreground',
  component: 'border-kind-component/40 text-kind-component',
  contract: 'border-kind-contract/40 text-kind-contract',
  concept: 'border-kind-concept/40 text-kind-concept',
  flow: 'border-kind-flow/40 text-kind-flow',
  decision: 'border-kind-decision/40 text-kind-decision',
  plan: 'border-kind-plan/40 text-kind-plan',
  task: 'border-kind-task/40 text-kind-task',
  learning: 'border-kind-learning/40 text-kind-learning',
  design: 'border-kind-design/40 text-kind-design',
};

export function KindBadge({ kind }: { kind: string }) {
  const colors = KIND_COLORS[kind] || 'border-border text-muted-foreground';
  return (
    <span className={`inline-flex items-center px-1.5 py-0.5 text-[10px] font-medium rounded border ${colors}`}>
      {kind}
    </span>
  );
}
