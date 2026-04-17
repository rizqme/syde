import { useMemo, useState } from 'react';
import { PlanDetailPhase, TaskSummary } from '../lib/api';
import { LayersIcon, ChevronRightIcon } from './icons';

interface PhaseTaskListProps {
  phases: PlanDetailPhase[];
  taskIndex: Record<string, TaskSummary>;
  onSelectTask: (slug: string) => void;
}

// Renders each phase as a collapsible section, with nested sub-phases
// rendered recursively inside their parents. Tasks inside a phase are
// resolved via taskIndex so every row shows name/status/priority/objective.
export function PhaseTaskList({ phases, taskIndex, onSelectTask }: PhaseTaskListProps) {
  // parent → children index for sub-phase nesting.
  const childrenByParent = useMemo(() => {
    const m: Record<string, PlanDetailPhase[]> = {};
    phases.forEach((ph) => {
      const parent = ph.ParentPhase || '';
      (m[parent] = m[parent] || []).push(ph);
    });
    return m;
  }, [phases]);

  const roots = childrenByParent[''] || [];

  if (!roots.length) {
    return (
      <div className="text-sm text-muted-foreground py-6">
        This plan has no phases yet.
      </div>
    );
  }

  return (
    <div className="space-y-3">
      {roots.map((phase, index) => (
        <PhaseSection
          key={phase.ID}
          phase={phase}
          childrenByParent={childrenByParent}
          taskIndex={taskIndex}
          onSelectTask={onSelectTask}
          depth={0}
          phaseLabel={`${index + 1}`}
        />
      ))}
    </div>
  );
}

function PhaseSection({
  phase,
  childrenByParent,
  taskIndex,
  onSelectTask,
  depth,
  phaseLabel,
}: {
  phase: PlanDetailPhase;
  childrenByParent: Record<string, PlanDetailPhase[]>;
  taskIndex: Record<string, TaskSummary>;
  onSelectTask: (slug: string) => void;
  depth: number;
  phaseLabel: string;
}) {
  const [open, setOpen] = useState(true);
  const subPhases = childrenByParent[phase.ID] || [];
  const tasks = phase.Tasks || [];
  const taskSummaries = tasks
    .map((slug) => taskIndex[slug])
    .filter((t): t is TaskSummary => Boolean(t));
  const done = taskSummaries.filter((t) => t.status === 'completed').length;

  const name = `Phase ${phaseLabel}: ${phase.Name || phase.Description || phase.ID}`;

  return (
    <div
      className={`rounded-lg border border-border bg-card ${depth > 0 ? 'ml-4' : ''}`}
    >
      <button
        onClick={() => setOpen((o) => !o)}
        className="w-full flex items-center gap-2 px-4 py-3 text-left"
      >
        <ChevronRightIcon className={`w-3 h-3 shrink-0 text-muted-foreground transition-transform ${open ? 'rotate-90' : ''}`} />
        <LayersIcon className={`w-3.5 h-3.5 shrink-0 ${phaseIconClass(phase.Status)}`} />
        <span className="text-sm font-medium flex-1 truncate">{name}</span>
        <PhaseStatusPill status={phase.Status} />
        {tasks.length > 0 && (
          <span className="text-[10px] text-muted-foreground shrink-0">
            {done}/{tasks.length} tasks done
          </span>
        )}
      </button>
      {open && (
        <div className="border-t border-border px-4 py-3 space-y-2">
          {taskSummaries.length > 0 && (
            <div className="space-y-1.5">
              {tasks.map((slug) => {
                const task = taskIndex[slug];
                if (!task) {
                  return (
                    <div
                      key={slug}
                      className="text-[11px] text-muted-foreground/60 font-mono"
                    >
                      {slug} (missing)
                    </div>
                  );
                }
                return (
                  <TaskRow
                    key={slug}
                    task={task}
                    onSelect={() => onSelectTask(task.slug)}
                  />
                );
              })}
            </div>
          )}
          {tasks.length === 0 && subPhases.length === 0 && (
            <div className="text-[11px] text-muted-foreground">No tasks in this phase.</div>
          )}
          {subPhases.length > 0 && (
            <div className="space-y-2 pt-1">
              {subPhases.map((sub, index) => (
                <PhaseSection
                  key={sub.ID}
                  phase={sub}
                  childrenByParent={childrenByParent}
                  taskIndex={taskIndex}
                  onSelectTask={onSelectTask}
                  depth={depth + 1}
                  phaseLabel={`${phaseLabel}.${index + 1}`}
                />
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  );
}

function TaskRow({ task, onSelect }: { task: TaskSummary; onSelect: () => void }) {
  const snippet =
    task.objective && task.objective.length > 80
      ? task.objective.slice(0, 80).trimEnd() + '…'
      : task.objective;
  return (
    <button
      onClick={onSelect}
      className="w-full text-left flex items-start gap-2 px-2 py-1.5 rounded hover:bg-muted/50 transition-colors"
    >
      <TaskStatusIcon status={task.status} />
      <div className="flex-1 min-w-0">
        <div className="flex items-center gap-2">
          <span className="text-xs font-medium truncate">{task.name}</span>
          <PriorityBadge priority={task.priority} />
        </div>
        {snippet && (
          <div className="text-[11px] text-muted-foreground mt-0.5 line-clamp-1">
            {snippet}
          </div>
        )}
      </div>
    </button>
  );
}

function PhaseStatusPill({ status }: { status: string }) {
  const cls = phaseIconClass(status);
  return (
    <span className={`text-[10px] px-1.5 py-0.5 rounded border border-current/30 shrink-0 ${cls}`}>
      {status}
    </span>
  );
}

function phaseIconClass(status: string) {
  const map: Record<string, string> = {
    completed: 'text-green-400',
    in_progress: 'text-yellow-400',
    skipped: 'text-muted-foreground/60',
    pending: 'text-muted-foreground',
  };
  return map[status] || map.pending;
}

function TaskStatusIcon({ status }: { status: string }) {
  const map: Record<string, { icon: string; cls: string }> = {
    completed: { icon: '✓', cls: 'text-green-400' },
    in_progress: { icon: '●', cls: 'text-yellow-400' },
    blocked: { icon: '✗', cls: 'text-red-400' },
    pending: { icon: '○', cls: 'text-muted-foreground' },
  };
  const s = map[status] || map.pending;
  return <span className={`text-xs mt-0.5 ${s.cls}`}>{s.icon}</span>;
}

const PRIORITY_COLORS: Record<string, string> = {
  high: 'text-red-400 border-red-400/30',
  medium: 'text-yellow-400 border-yellow-400/30',
  low: 'text-zinc-400 border-zinc-400/30',
};

function PriorityBadge({ priority }: { priority: string }) {
  const cls = PRIORITY_COLORS[priority] || 'text-muted-foreground border-border';
  return (
    <span
      className={`text-[10px] px-1.5 py-0.5 rounded border shrink-0 ${cls}`}
    >
      {priority}
    </span>
  );
}
