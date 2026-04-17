import { useEffect, useMemo, useState } from 'react';
import Markdown from 'react-markdown';
import { useApi } from '../hooks/useApi';
import { api, PlanDetail } from '../lib/api';
import { PhaseTaskList } from './PhaseTaskList';
import { PlanChangesView } from './PlanChangesView';

interface PlanDetailPanelProps {
  slug: string;
  onClose: () => void;
  onNavigate: (slug: string, kind?: string, filter?: string) => void;
  onOpenFile: (path: string) => void;
}

type PlanTab = 'plan' | 'tasks';

export function PlanDetailPanel({
  slug,
  onClose,
  onNavigate,
}: PlanDetailPanelProps) {
  const { data, loading, error } = useApi<PlanDetail>(() => api.planDetail(slug), [slug]);
  const [tab, setTab] = useState<PlanTab>(() => readTabFromURL());

  useEffect(() => {
    setTab(readTabFromURL());
  }, [slug]);

  const setActiveTab = (next: PlanTab) => {
    setTab(next);
    const url = new URL(window.location.href);
    if (next === 'plan') {
      url.searchParams.delete('tab');
    } else {
      url.searchParams.set('tab', next);
    }
    window.history.replaceState(window.history.state, '', url);
  };

  if (loading) {
    return (
      <div className="px-8 py-8 text-muted-foreground text-sm">
        Loading plan...
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="px-8 py-8 text-sm text-red-400">
        Failed to load plan: {error ? String(error) : 'not found'}
      </div>
    );
  }

  const pct = Math.round(data.progress || 0);

  return (
    <article className="px-8 py-8 max-w-5xl">
      <header className="mb-5">
        <div className="flex items-start gap-3">
          <div className="min-w-0 flex-1">
            <div className="flex flex-wrap items-center gap-3 mb-2">
              <h1 className="text-xl font-semibold tracking-tight truncate">
                {data.name}
              </h1>
              <PlanStatusBadge status={data.status} />
              <span className="text-xs text-muted-foreground">{pct}% complete</span>
            </div>
            {data.approved_at && (
              <div className="text-[11px] text-muted-foreground">
                Approved {formatDate(data.approved_at)}
              </div>
            )}
          </div>
        </div>
        <div className="mt-3 h-1.5 bg-muted rounded-full overflow-hidden max-w-md">
          <div
            className="h-full bg-foreground rounded-full transition-all"
            style={{ width: `${pct}%` }}
          />
        </div>
      </header>

      <div className="flex items-center gap-1 mb-5 border-b border-border">
        <TabButton
          label="Plan"
          active={tab === 'plan'}
          onClick={() => setActiveTab('plan')}
        />
        <TabButton
          label="Tasks"
          active={tab === 'tasks'}
          onClick={() => setActiveTab('tasks')}
        />
      </div>

      {tab === 'plan' ? (
        <PlanOverviewTab plan={data} />
      ) : (
        <PhaseTaskList
          phases={data.phases || []}
          taskIndex={data.task_index || {}}
          onSelectTask={(taskSlug) => onNavigate(taskSlug, 'task')}
        />
      )}
    </article>
  );
}

function PlanOverviewTab({ plan }: { plan: PlanDetail }) {
  const hasChanges = useMemo(() => {
    if (!plan.changes) return false;
    return Object.values(plan.changes).some((lane) => {
      if (!lane) return false;
      return (
        (lane.deleted?.length || 0) +
          (lane.extended?.length || 0) +
          (lane.new?.length || 0) >
        0
      );
    });
  }, [plan.changes]);

  return (
    <div className="space-y-6">
      {plan.background && (
        <Section title="Background">
          <ProseMarkdown>{plan.background}</ProseMarkdown>
        </Section>
      )}
      {plan.objective && (
        <Section title="Objective">
          <ProseMarkdown>{plan.objective}</ProseMarkdown>
        </Section>
      )}
      {plan.scope && (
        <Section title="Scope">
          <ProseMarkdown>{plan.scope}</ProseMarkdown>
        </Section>
      )}
      {plan.design && (
        <Section title="Design">
          <ProseMarkdown>{plan.design}</ProseMarkdown>
        </Section>
      )}
      <Section title="Changes">
        <PlanChangesView changes={plan.changes} />
      </Section>
      {!plan.design &&
        !plan.background &&
        !plan.objective &&
        !plan.scope &&
        !hasChanges && (
          <div className="text-sm text-muted-foreground italic">
            This plan has no design prose or structured changes yet.
          </div>
        )}
    </div>
  );
}

function ProseMarkdown({ children }: { children: string }) {
  return (
    <div className="prose-plan text-sm leading-relaxed text-muted-foreground [&_h1]:text-base [&_h1]:font-semibold [&_h1]:text-foreground [&_h1]:mt-4 [&_h1]:mb-2 [&_h2]:text-sm [&_h2]:font-semibold [&_h2]:text-foreground [&_h2]:mt-3 [&_h2]:mb-1.5 [&_h3]:text-sm [&_h3]:font-medium [&_h3]:text-foreground [&_h3]:mt-2 [&_h3]:mb-1 [&_p]:mb-2 [&_ul]:list-disc [&_ul]:pl-5 [&_ul]:mb-2 [&_ol]:list-decimal [&_ol]:pl-5 [&_ol]:mb-2 [&_li]:mb-0.5 [&_code]:text-xs [&_code]:bg-muted [&_code]:px-1 [&_code]:py-0.5 [&_code]:rounded [&_code]:text-foreground [&_pre]:bg-muted [&_pre]:p-3 [&_pre]:rounded [&_pre]:mb-2 [&_pre]:overflow-x-auto [&_pre_code]:bg-transparent [&_pre_code]:p-0 [&_strong]:text-foreground [&_strong]:font-semibold [&_a]:text-blue-400 [&_a]:underline [&_blockquote]:border-l-2 [&_blockquote]:border-border [&_blockquote]:pl-3 [&_blockquote]:italic [&_table]:w-full [&_table]:text-xs [&_table]:mb-2 [&_th]:text-left [&_th]:border-b [&_th]:border-border [&_th]:pb-1 [&_th]:pr-3 [&_th]:text-foreground [&_td]:border-b [&_td]:border-border/50 [&_td]:py-1 [&_td]:pr-3">
      <Markdown>{children}</Markdown>
    </div>
  );
}

function Section({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <section>
      <h2 className="text-[11px] font-medium text-muted-foreground uppercase tracking-widest mb-2">
        {title}
      </h2>
      {children}
    </section>
  );
}

function TabButton({
  label,
  active,
  onClick,
}: {
  label: string;
  active: boolean;
  onClick: () => void;
}) {
  return (
    <button
      onClick={onClick}
      className={`px-3 py-1.5 text-xs font-medium border-b-2 transition-colors -mb-px ${
        active
          ? 'border-foreground text-foreground'
          : 'border-transparent text-muted-foreground hover:text-foreground'
      }`}
    >
      {label}
    </button>
  );
}

const STATUS_COLORS: Record<string, string> = {
  draft: 'border-border text-muted-foreground',
  approved: 'border-blue-400/40 text-blue-400',
  'in-progress': 'border-yellow-400/40 text-yellow-400',
  in_progress: 'border-yellow-400/40 text-yellow-400',
  completed: 'border-green-400/40 text-green-400',
};

function PlanStatusBadge({ status }: { status: string }) {
  const color = STATUS_COLORS[status] || 'border-border text-muted-foreground';
  return (
    <span
      className={`inline-flex items-center px-1.5 py-0.5 text-[10px] font-medium rounded border ${color}`}
    >
      {status}
    </span>
  );
}

function readTabFromURL(): PlanTab {
  const params = new URLSearchParams(window.location.search);
  return params.get('tab') === 'tasks' ? 'tasks' : 'plan';
}

function formatDate(iso: string): string {
  try {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return iso;
    return d.toLocaleDateString(undefined, {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  } catch {
    return iso;
  }
}
