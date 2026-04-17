// Kind-specialized preview for a NewChange contract draft. The plan detail
// API returns the raw draft fields (contract_kind, input, input_parameters,
// output, output_parameters, wireframe, ...) and this component switches on
// contract_kind to render the most useful view.
//
interface NewContractDraftViewProps {
  draft: Record<string, any>;
}

type Param = { path?: string; name?: string; type?: string; description?: string };

export function NewContractDraftView({ draft }: NewContractDraftViewProps) {
  const kind = (draft.contract_kind || '').toLowerCase();
  const pattern = draft.interaction_pattern;

  return (
    <div className="space-y-3">
      <div className="flex flex-wrap items-center gap-2">
        {draft.contract_kind && (
          <span className="text-[10px] uppercase tracking-wider px-1.5 py-0.5 rounded bg-muted text-muted-foreground">
            {draft.contract_kind}
          </span>
        )}
        {pattern && (
          <span className="text-[10px] uppercase tracking-wider px-1.5 py-0.5 rounded bg-muted text-muted-foreground">
            {pattern}
          </span>
        )}
      </div>

      {renderBody(kind, draft)}

      {draft.protocol_notes && (
        <div className="text-[11px] text-muted-foreground leading-relaxed">
          <span className="uppercase tracking-widest text-[10px] mr-1">Protocol:</span>
          {draft.protocol_notes}
        </div>
      )}
      {draft.constraints && (
        <div className="text-[11px] text-muted-foreground leading-relaxed">
          <span className="uppercase tracking-widest text-[10px] mr-1">Constraints:</span>
          {draft.constraints}
        </div>
      )}
    </div>
  );
}

function renderBody(kind: string, draft: Record<string, any>) {
  switch (kind) {
    case 'screen':
      return <ScreenBody draft={draft} />;
    case 'cli':
      return <CliBody draft={draft} />;
    case 'rest':
    case 'rpc':
      return <RestBody draft={draft} />;
    case 'storage':
      return <StorageBody draft={draft} />;
    case 'event':
    case 'websocket':
      return <EventBody draft={draft} />;
    default:
      return <GenericKV draft={draft} />;
  }
}

function ScreenBody({ draft }: { draft: Record<string, any> }) {
  const wireframe: string | undefined = draft.wireframe;
  const wireframeHTML: string | undefined = draft.wireframe_html;
  return (
    <div className="space-y-2">
      {draft.input && <HeaderLine label="Route">{draft.input}</HeaderLine>}
      {wireframeHTML ? (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Wireframe
          </div>
          <div
            className="rounded border border-border bg-card p-3 overflow-x-auto text-xs"
            dangerouslySetInnerHTML={{ __html: wireframeHTML }}
          />
        </div>
      ) : wireframe ? (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Wireframe (uiml source)
          </div>
          <pre className="text-[11px] overflow-auto bg-muted p-2 rounded font-mono whitespace-pre-wrap">
            {wireframe}
          </pre>
        </div>
      ) : null}
      {draft.output && <HeaderLine label="Output">{draft.output}</HeaderLine>}
    </div>
  );
}

function CliBody({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-2">
      {draft.input && (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Command
          </div>
          <code className="block font-mono text-xs bg-card border border-border rounded px-2 py-1.5">
            {draft.input}
          </code>
        </div>
      )}
      <ParamTable label="Flags" params={draft.input_parameters} />
      {draft.output && <HeaderLine label="Output">{draft.output}</HeaderLine>}
    </div>
  );
}

function RestBody({ draft }: { draft: Record<string, any> }) {
  const { method, path } = parseMethodPath(draft.input);
  return (
    <div className="space-y-2">
      <div className="flex items-baseline gap-2">
        {method && (
          <span className="font-mono text-[11px] px-2 py-0.5 rounded bg-muted text-foreground">
            {method}
          </span>
        )}
        {path && <code className="font-mono text-xs text-foreground">{path}</code>}
        {!method && !path && draft.input && (
          <code className="font-mono text-xs text-foreground">{draft.input}</code>
        )}
      </div>
      <ParamTable label="Request Parameters" params={draft.input_parameters} />
      <ParamTable label="Response Parameters" params={draft.output_parameters} />
      {draft.output && !draft.output_parameters && <HeaderLine label="Response">{draft.output}</HeaderLine>}
    </div>
  );
}

function StorageBody({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-2">
      {draft.input && (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Key Pattern
          </div>
          <code className="block font-mono text-xs bg-card border border-border rounded px-2 py-1.5">
            {draft.input}
          </code>
        </div>
      )}
      <ParamTable label="Fields" params={draft.output_parameters} />
    </div>
  );
}

function EventBody({ draft }: { draft: Record<string, any> }) {
  return (
    <div className="space-y-2">
      {draft.input && (
        <div>
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
            Event
          </div>
          <code className="block font-mono text-xs bg-card border border-border rounded px-2 py-1.5">
            {draft.input}
          </code>
        </div>
      )}
      <ParamTable label="Payload" params={draft.input_parameters} />
    </div>
  );
}

function GenericKV({ draft }: { draft: Record<string, any> }) {
  const entries = Object.entries(draft).filter(([k]) => k !== 'contract_kind' && k !== 'interaction_pattern' && k !== 'protocol_notes' && k !== 'constraints');
  if (entries.length === 0) return null;
  return (
    <div className="rounded border border-border bg-card/40 divide-y divide-border">
      {entries.map(([k, v]) => (
        <div key={k} className="p-2 text-xs">
          <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-0.5">{k}</div>
          <pre className="font-mono text-[11px] whitespace-pre-wrap break-words text-foreground">
            {typeof v === 'string' ? v : JSON.stringify(v, null, 2)}
          </pre>
        </div>
      ))}
    </div>
  );
}

function HeaderLine({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <div>
      <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
        {label}
      </div>
      <code className="block font-mono text-xs bg-card border border-border rounded px-2 py-1.5">
        {children}
      </code>
    </div>
  );
}

function ParamTable({ label, params }: { label: string; params: Param[] | undefined }) {
  if (!params || params.length === 0) return null;
  return (
    <div>
      <div className="text-[10px] font-medium text-muted-foreground uppercase tracking-widest mb-1">
        {label}
      </div>
      <div className="rounded border border-border bg-card divide-y divide-border">
        {params.map((p, i) => {
          const name = p.path || p.name || '';
          return (
            <div key={i} className="p-2 text-xs">
              <div className="flex items-baseline gap-2">
                <span className="font-mono text-foreground">{name}</span>
                {p.type && (
                  <span className="font-mono text-[10px] px-1.5 py-0.5 rounded bg-muted text-muted-foreground">
                    {p.type}
                  </span>
                )}
              </div>
              {p.description && (
                <div className="text-muted-foreground mt-0.5 leading-relaxed">
                  {p.description}
                </div>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
}

// Parse "GET /api/foo" -> { method: "GET", path: "/api/foo" }.
// If the string doesn't match the pattern, return the whole thing as path.
function parseMethodPath(input: string | undefined): { method?: string; path?: string } {
  if (!input) return {};
  const trimmed = input.trim();
  const m = trimmed.match(/^([A-Z]+)\s+(\S.*)$/);
  if (m) return { method: m[1], path: m[2] };
  return { path: trimmed };
}
