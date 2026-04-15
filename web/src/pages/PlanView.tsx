import { useApi } from '../hooks/useApi';
import { api, Plan, PlanPhase } from '../lib/api';

interface PlanViewProps {
  onSelectEntity: (slug: string) => void;
}

export function PlanView({ onSelectEntity }: PlanViewProps) {
  const { data, loading } = useApi(() => api.plans());
  const plans = data?.plans || [];

  if (loading) return <div className="text-muted-foreground text-sm">Loading plans...</div>;
  if (!plans.length) return <EmptyPlans />;

  return (
    <div>
      <h1 className="text-xl font-semibold tracking-tight mb-6">Plans</h1>
      <div className="space-y-6">
        {plans.map((plan) => (
          <PlanCard key={plan.name} plan={plan} onSelectEntity={onSelectEntity} />
        ))}
      </div>
    </div>
  );
}

function PlanCard({ plan, onSelectEntity }: { plan: Plan; onSelectEntity: (slug: string) => void }) {
  const pct = Math.round(plan.progress);
  const phases = plan.phases || [];

  // Build parent→children map
  const children: Record<string, PlanPhase[]> = {};
  const roots: PlanPhase[] = [];
  phases.forEach((ph) => {
    if (ph.parent_phase) {
      (children[ph.parent_phase] = children[ph.parent_phase] || []).push(ph);
    } else {
      roots.push(ph);
    }
  });

  return (
    <div className="bg-card border border-border rounded-lg p-5">
      <div className="flex items-center justify-between mb-3">
        <h2 className="text-base font-semibold">{plan.name}</h2>
        <span className="text-[10px] px-2 py-0.5 rounded-full border border-border text-muted-foreground">
          {plan.status}
        </span>
      </div>

      <div className="h-1.5 bg-muted rounded-full overflow-hidden mb-1">
        <div className="h-full bg-foreground rounded-full transition-all" style={{ width: `${pct}%` }} />
      </div>
      <div className="text-[10px] text-muted-foreground mb-4">{pct}% complete</div>

      {/* Entity drafts */}
      {phases.some((ph) => ph.entities?.length) && (
        <div className="mb-4">
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-2">
            Draft Entities
          </div>
          <div className="flex flex-wrap gap-1.5">
            {phases.flatMap((ph) =>
              (ph.entities || []).map((e, i) => (
                <button
                  key={`${ph.id}-${i}`}
                  onClick={() => onSelectEntity(e.name.toLowerCase().replace(/\s+/g, '-'))}
                  className="text-[10px] px-2 py-1 rounded border border-border text-muted-foreground hover:text-foreground hover:border-muted-foreground transition-colors"
                >
                  [{e.kind}] {e.name}
                </button>
              ))
            )}
          </div>
        </div>
      )}

      {/* Phase tree */}
      <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-2">
        Phases
      </div>
      <div className="space-y-0.5">
        {roots.map((ph) => (
          <PhaseNode key={ph.id} phase={ph} children={children} indent={0} />
        ))}
      </div>
    </div>
  );
}

function PhaseNode({ phase, children, indent }: { phase: PlanPhase; children: Record<string, PlanPhase[]>; indent: number }) {
  const icon = phase.status === 'completed' ? '✓' : phase.status === 'in_progress' ? '●' : '○';
  const iconColor = phase.status === 'completed' ? 'text-green-400' : phase.status === 'in_progress' ? 'text-foreground' : 'text-muted-foreground';
  const name = phase.name || phase.description;
  const childPhases = children[phase.id] || [];

  return (
    <div style={{ paddingLeft: indent * 16 }}>
      <div className="flex items-center gap-2 py-1">
        <span className={`text-xs ${iconColor}`}>{icon}</span>
        <span className={`text-xs ${phase.status === 'completed' ? 'text-muted-foreground' : 'text-foreground'}`}>
          {name}
        </span>
        {(phase.tasks?.length || 0) > 0 && (
          <span className="text-[10px] text-muted-foreground">
            {phase.tasks?.length} tasks
          </span>
        )}
      </div>
      {/* Tasks */}
      {phase.tasks?.map((task) => (
        <div key={task} className="flex items-center gap-2 py-0.5" style={{ paddingLeft: 16 }}>
          <span className="text-[10px] text-muted-foreground">○</span>
          <span className="text-[10px] text-muted-foreground font-mono">{task}</span>
        </div>
      ))}
      {childPhases.map((child) => (
        <PhaseNode key={child.id} phase={child} children={children} indent={indent + 1} />
      ))}
    </div>
  );
}

function EmptyPlans() {
  return (
    <div className="text-center py-16">
      <div className="text-2xl mb-2">📋</div>
      <div className="text-sm font-medium mb-1">No plans yet</div>
      <div className="text-xs text-muted-foreground">
        Run: <code className="font-mono bg-card px-1.5 py-0.5 rounded border border-border">syde plan create &lt;name&gt;</code>
      </div>
    </div>
  );
}
