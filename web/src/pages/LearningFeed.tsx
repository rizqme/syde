import { useApi } from '../hooks/useApi';
import { api, Learning } from '../lib/api';

interface LearningFeedProps {
  onSelectEntity: (slug: string) => void;
}

const CATEGORY_ICONS: Record<string, string> = {
  gotcha: '⚠',
  constraint: '⚠',
  convention: 'ℹ',
  context: 'ℹ',
  dependency: '🔗',
  performance: '⚡',
  workaround: '🔧',
};

export function LearningFeed({ onSelectEntity }: LearningFeedProps) {
  const { data, loading } = useApi(() => api.learnings());
  const learnings = data?.learnings || [];

  if (loading) return <div className="text-muted-foreground text-sm">Loading learnings...</div>;
  if (!learnings.length) return <EmptyLearnings />;

  // Group by category
  const grouped: Record<string, Learning[]> = {};
  learnings.forEach((l) => { (grouped[l.category] = grouped[l.category] || []).push(l); });

  // Sort categories: gotcha/constraint first
  const sortedCategories = Object.keys(grouped).sort((a, b) => {
    const priority: Record<string, number> = { gotcha: 0, constraint: 1 };
    return (priority[a] ?? 5) - (priority[b] ?? 5);
  });

  return (
    <div>
      <h1 className="text-xl font-semibold tracking-tight mb-6">Learnings</h1>
      <div className="space-y-6">
        {sortedCategories.map((category) => (
          <div key={category}>
            <div className="flex items-center gap-2 mb-3">
              <span>{CATEGORY_ICONS[category] || 'ℹ'}</span>
              <span className="text-xs font-medium text-muted-foreground uppercase tracking-widest">
                {category}
              </span>
              <span className="text-[10px] bg-muted px-1.5 py-0.5 rounded-full text-muted-foreground">
                {grouped[category].length}
              </span>
            </div>
            <div className="space-y-2">
              {grouped[category].map((learning, i) => (
                <div key={i} className="bg-card border border-border rounded-lg p-4">
                  <div className="flex items-center gap-2 mb-2">
                    <span className="text-[10px] px-1.5 py-0.5 rounded border border-kind-learning/30 text-kind-learning">
                      {learning.category}
                    </span>
                    <span className="text-[10px] px-1.5 py-0.5 rounded border border-border text-muted-foreground">
                      {learning.confidence}
                    </span>
                  </div>
                  <p className="text-xs text-muted-foreground leading-relaxed">{learning.description}</p>
                  {learning.entity_refs?.length > 0 && (
                    <div className="flex flex-wrap gap-1 mt-2">
                      {learning.entity_refs.map((ref) => (
                        <button
                          key={ref}
                          onClick={() => onSelectEntity(ref.toLowerCase().replace(/\s+/g, '-'))}
                          className="text-[10px] px-1.5 py-0.5 rounded border border-kind-component/30 text-kind-component hover:bg-accent transition-colors"
                        >
                          {ref}
                        </button>
                      ))}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

function EmptyLearnings() {
  return (
    <div className="text-center py-16">
      <div className="text-2xl mb-2">💡</div>
      <div className="text-sm font-medium mb-1">No learnings yet</div>
      <div className="text-xs text-muted-foreground">
        Run: <code className="font-mono bg-card px-1.5 py-0.5 rounded border border-border">syde remember "&lt;text&gt;" --category gotcha</code>
      </div>
    </div>
  );
}
