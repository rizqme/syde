// Kind-specialized preview for a NewChange entity draft. Dispatches on the
// lane kind (the key under plan.changes) and renders the fields that matter
// for that kind. Contract drafts are delegated to NewContractDraftView which
// further specializes on contract_kind.

import { NewContractDraftView } from './NewContractDraftView';

interface NewEntityDraftViewProps {
  kind: string;
  draft: Record<string, any>;
}

export function NewEntityDraftView({ kind, draft }: NewEntityDraftViewProps) {
  if (!draft) return null;
  switch (kind) {
    case 'contract':
    case 'contracts':
      return <NewContractDraftView draft={draft} />;
    case 'component':
    case 'components':
      return <ComponentDraft draft={draft} />;
    case 'requirement':
    case 'requirements':
      return <RequirementDraft draft={draft} />;
    case 'concept':
    case 'concepts':
      return <ConceptDraft draft={draft} />;
    case 'flow':
    case 'flows':
      return <FlowDraft draft={draft} />;
    case 'system':
    case 'systems':
      return <SystemDraft draft={draft} />;
    default:
      return <GenericDraft draft={draft} />;
  }
}

function ComponentDraft({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-3">
      {draft.responsibility && (
        <Field label="Responsibility" value={draft.responsibility} highlight />
      )}
      {Array.isArray(draft.capabilities) && draft.capabilities.length > 0 && (
        <div>
          <FieldLabel>Capabilities</FieldLabel>
          <ul className="space-y-1.5">
            {draft.capabilities.map((c: string, i: number) => (
              <li
                key={i}
                className="flex items-center gap-2.5 text-xs leading-relaxed rounded-md border border-border bg-card px-3 py-2 text-foreground"
              >
                <span className="inline-block h-1.5 w-1.5 shrink-0 rounded-full bg-kind-component" />
                <span className="min-w-0">{c}</span>
              </li>
            ))}
          </ul>
        </div>
      )}
      {draft.boundaries && <Field label="Boundaries" value={draft.boundaries} muted />}
      {draft.purpose && (
        <div className="text-xs italic text-muted-foreground">{draft.purpose}</div>
      )}
    </div>
  );
}

function RequirementDraft({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-3">
      {draft.statement && (
        <blockquote className="border-l-2 border-kind-requirement pl-3 text-sm italic">
          {draft.statement}
        </blockquote>
      )}
      <div className="flex flex-wrap items-center gap-2 text-xs">
        {draft.req_type && (
          <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">{draft.req_type}</span>
        )}
        {draft.priority && (
          <span className="px-2 py-0.5 rounded bg-muted text-muted-foreground">
            priority: {draft.priority}
          </span>
        )}
      </div>
      {draft.verification && <Field label="Verification" value={draft.verification} muted />}
      {draft.rationale && <Field label="Rationale" value={draft.rationale} />}
    </div>
  );
}

function ConceptDraft({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-3">
      {draft.meaning && (
        <div className="text-xs italic text-foreground">{draft.meaning}</div>
      )}
      {draft.invariants && <Field label="Invariants" value={draft.invariants} />}
      {Array.isArray(draft.attributes) && draft.attributes.length > 0 && (
        <div>
          <FieldLabel>Attributes</FieldLabel>
          <div className="rounded border border-border bg-card divide-y divide-border">
            {draft.attributes.map((a: { name: string; description?: string }, i: number) => (
              <div key={i} className="p-2 text-xs">
                <div className="font-mono text-foreground">{a.name}</div>
                {a.description && (
                  <div className="text-muted-foreground mt-0.5 leading-relaxed">
                    {a.description}
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>
      )}
      {Array.isArray(draft.actions) && draft.actions.length > 0 && (
        <div>
          <FieldLabel>Actions</FieldLabel>
          <ul className="space-y-1 text-xs">
            {draft.actions.map((a: any, i: number) => (
              <li key={i} className="font-mono text-foreground">
                {typeof a === 'string' ? a : a.name}
                {typeof a !== 'string' && a.description && (
                  <span className="ml-2 text-muted-foreground font-sans">{a.description}</span>
                )}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

function FlowDraft({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-3">
      {(draft.trigger || draft.goal) && (
        <div className="flex items-center gap-2 text-sm">
          {draft.trigger && <span className="text-muted-foreground">{draft.trigger}</span>}
          {draft.trigger && draft.goal && <span className="text-kind-flow">→</span>}
          {draft.goal && <span className="font-medium">{draft.goal}</span>}
        </div>
      )}
      {draft.happy_path && <Field label="Happy Path" value={draft.happy_path} />}
      {draft.edge_cases && <Field label="Edge Cases" value={draft.edge_cases} />}
    </div>
  );
}

function SystemDraft({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-3">
      {draft.purpose && <Field label="Purpose" value={draft.purpose} highlight />}
      {draft.quality_goals && <Field label="Quality Goals" value={draft.quality_goals} />}
      {draft.design_principles && (
        <Field label="Design Principles" value={draft.design_principles} muted />
      )}
    </div>
  );
}

function GenericDraft({ draft }: { draft: Record<string, any> }) {
  const entries = Object.entries(draft);
  if (entries.length === 0) return null;
  return (
    <div className="rounded border border-border bg-card/40 divide-y divide-border">
      {entries.map(([k, v]) => (
        <div key={k} className="p-2 text-xs">
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-0.5">
            {k}
          </div>
          <pre className="font-mono text-[11px] whitespace-pre-wrap break-words text-foreground">
            {typeof v === 'string' ? v : JSON.stringify(v, null, 2)}
          </pre>
        </div>
      ))}
    </div>
  );
}

function FieldLabel({ children }: { children: React.ReactNode }) {
  return (
    <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
      {children}
    </div>
  );
}

function Field({
  label,
  value,
  highlight,
  muted,
}: {
  label: string;
  value: string;
  highlight?: boolean;
  muted?: boolean;
}) {
  return (
    <div>
      <FieldLabel>{label}</FieldLabel>
      <div
        className={`text-xs leading-relaxed rounded p-2 ${
          highlight
            ? 'bg-card border border-border'
            : muted
              ? 'bg-muted/50 text-muted-foreground'
              : 'text-muted-foreground'
        }`}
      >
        {value}
      </div>
    </div>
  );
}
