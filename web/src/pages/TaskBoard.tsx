import { useApi } from '../hooks/useApi';
import { api, Task } from '../lib/api';

interface TaskBoardProps {
  onSelectEntity: (slug: string) => void;
}

const COLUMNS = [
  { status: 'pending', label: 'Pending', icon: '○' },
  { status: 'in_progress', label: 'In Progress', icon: '●' },
  { status: 'completed', label: 'Completed', icon: '✓' },
  { status: 'blocked', label: 'Blocked', icon: '✗' },
];

const PRIORITY_COLORS: Record<string, string> = {
  high: 'text-red-400 border-red-400/30',
  medium: 'text-yellow-400 border-yellow-400/30',
  low: 'text-zinc-400 border-zinc-400/30',
};

export function TaskBoard({ onSelectEntity }: TaskBoardProps) {
  const { data, loading } = useApi(() => api.tasks());
  const tasks = data?.tasks || [];

  if (loading) return <div className="text-muted-foreground text-sm">Loading tasks...</div>;
  if (!tasks.length) return <EmptyTasks />;

  const byStatus: Record<string, Task[]> = {};
  tasks.forEach((t) => { (byStatus[t.status] = byStatus[t.status] || []).push(t); });

  return (
    <div>
      <h1 className="text-xl font-semibold tracking-tight mb-6">Tasks</h1>
      <div className="grid grid-cols-4 gap-4">
        {COLUMNS.map((col) => {
          const items = byStatus[col.status] || [];
          return (
            <div key={col.status}>
              <div className="flex items-center gap-2 mb-3">
                <span className="text-xs">{col.icon}</span>
                <span className="text-xs font-medium text-muted-foreground uppercase tracking-widest">
                  {col.label}
                </span>
                <span className="text-[10px] bg-muted px-1.5 py-0.5 rounded-full text-muted-foreground">
                  {items.length}
                </span>
              </div>
              <div className="space-y-2">
                {items.map((task) => (
                  <div
                    key={task.name}
                    className="bg-card border border-border rounded-lg p-3 hover:border-muted-foreground transition-colors cursor-pointer"
                    onClick={() => onSelectEntity(task.name.toLowerCase().replace(/\s+/g, '-'))}
                  >
                    <div className="text-xs font-medium mb-1">{task.name}</div>
                    <div className="flex items-center gap-2">
                      <span className={`text-[10px] px-1.5 py-0.5 rounded border ${PRIORITY_COLORS[task.priority] || 'text-muted-foreground border-border'}`}>
                        {task.priority}
                      </span>
                      {task.plan_ref && (
                        <span className="text-[10px] text-muted-foreground truncate">
                          {task.plan_ref}
                        </span>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}

function EmptyTasks() {
  return (
    <div className="text-center py-16">
      <div className="text-2xl mb-2">☐</div>
      <div className="text-sm font-medium mb-1">No tasks yet</div>
      <div className="text-xs text-muted-foreground">
        Run: <code className="font-mono bg-card px-1.5 py-0.5 rounded border border-border">syde task create &lt;name&gt;</code>
      </div>
    </div>
  );
}
