import { useApi } from '../hooks/useApi';
import { api, StatusResponse, Plan, Learning } from '../lib/api';

interface OverviewProps {
  onNavigate: (kind: string) => void;
}

export function Overview({ onNavigate }: OverviewProps) {
  const { data: status } = useApi<StatusResponse>(() => api.status());
  const { data: plansData } = useApi(() => api.plans());
  const { data: learningsData } = useApi(() => api.learnings());

  const counts = status?.counts || {};
  const plans = plansData?.plans || [];
  const learnings = (learningsData?.learnings || []).slice(0, 5);

  const kindOrder = ['component', 'contract', 'concept', 'flow', 'decision', 'plan', 'task', 'learning'];

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-semibold tracking-tight">Overview</h1>
        <p className="text-muted-foreground text-sm mt-1">
          {status?.total || 0} entities in the design model
        </p>
      </div>

      <div className="grid grid-cols-4 gap-3 mb-8">
        {kindOrder.map((kind) => (
          <button
            key={kind}
            onClick={() => onNavigate(kind)}
            className="bg-card border border-border rounded-lg p-3 text-left hover:border-muted-foreground transition-colors"
          >
            <div className="text-[10px] text-muted-foreground uppercase tracking-widest">{kind}s</div>
            <div className="text-2xl font-semibold mt-1">{counts[kind] || 0}</div>
          </button>
        ))}
      </div>

      <div className="grid grid-cols-2 gap-6">
        <div>
          <h2 className="text-xs font-medium text-muted-foreground uppercase tracking-widest mb-3">Plans</h2>
          {plans.length === 0 ? (
            <div className="text-sm text-muted-foreground">No plans yet.</div>
          ) : (
            <div className="space-y-3">
              {plans.map((p: Plan) => (
                <div key={p.name} className="bg-card border border-border rounded-lg p-3">
                  <div className="flex items-center justify-between mb-2">
                    <span className="text-sm font-medium">{p.name}</span>
                    <span className="text-[10px] px-1.5 py-0.5 rounded border border-border text-muted-foreground">
                      {p.status}
                    </span>
                  </div>
                  <div className="h-1.5 bg-muted rounded-full overflow-hidden">
                    <div className="h-full bg-foreground rounded-full" style={{ width: `${Math.round(p.progress)}%` }} />
                  </div>
                  <div className="text-[10px] text-muted-foreground mt-1">{Math.round(p.progress)}% complete</div>
                </div>
              ))}
            </div>
          )}
        </div>

        <div>
          <h2 className="text-xs font-medium text-muted-foreground uppercase tracking-widest mb-3">Recent Learnings</h2>
          {learnings.length === 0 ? (
            <div className="text-sm text-muted-foreground">No learnings yet.</div>
          ) : (
            <div className="space-y-2">
              {learnings.map((l: Learning, i: number) => (
                <div key={i} className="flex items-start gap-2 py-1.5">
                  <span>{l.category === 'gotcha' || l.category === 'constraint' ? '⚠' : 'ℹ'}</span>
                  <div>
                    <span className="text-[10px] px-1.5 py-0.5 rounded border border-kind-learning/30 text-kind-learning mr-1">
                      {l.category}
                    </span>
                    <span className="text-xs text-muted-foreground">{l.description}</span>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
