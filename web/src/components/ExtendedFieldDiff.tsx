import { useState } from 'react';

// Side-by-side view of field-level changes declared on an ExtendedChange.
// The plan model stores `field_changes` as a map of field name -> proposed
// string value (the sentinel "DELETE" means "drop the field"), and the API
// resolves the current entity snapshot into `current_values`. This component
// renders a two-column table so reviewers can see exactly what flips.

interface ExtendedFieldDiffProps {
  currentValues: Record<string, any>;
  fieldChanges: Record<string, string>;
  proposedValuesHTML?: Record<string, string>;
}

export function ExtendedFieldDiff({
  currentValues,
  fieldChanges,
  proposedValuesHTML,
}: ExtendedFieldDiffProps) {
  const fields = Object.keys(fieldChanges);
  if (fields.length === 0) return null;

  return (
    <div className="rounded border border-border bg-card/40 overflow-hidden">
      <div className="grid grid-cols-[auto_1fr_1fr] text-[10px] font-medium text-muted-foreground uppercase tracking-widest border-b border-border">
        <div className="px-3 py-1.5">Field</div>
        <div className="px-3 py-1.5 border-l border-border">Current</div>
        <div className="px-3 py-1.5 border-l border-border">Proposed</div>
      </div>
      {fields.map((name) => {
        const current = currentValues ? currentValues[name] : undefined;
        const proposed = fieldChanges[name];
        if (
          name === 'wireframe' &&
          (currentValues?.wireframe_html || proposedValuesHTML?.wireframe)
        ) {
          return (
            <div key={name} className="border-t border-border first:border-t-0">
              <div className="px-3 py-2 font-mono text-xs text-foreground border-b border-border">
                {name}
              </div>
              <WireframeDiff
                currentHTML={currentValues?.wireframe_html}
                currentSource={current}
                proposedHTML={proposedValuesHTML?.wireframe}
                proposedSource={proposed}
              />
            </div>
          );
        }
        return (
          <div
            key={name}
            className="grid grid-cols-[auto_1fr_1fr] text-xs border-t border-border first:border-t-0"
          >
            <div className="px-3 py-2 font-mono text-foreground whitespace-nowrap">{name}</div>
            <div className="px-3 py-2 border-l border-border">
              <ValueBlock value={current} />
            </div>
            <div className="px-3 py-2 border-l border-border">
              {proposed === 'DELETE' ? (
                <span className="text-red-400 line-through italic">(remove)</span>
              ) : (
                <ValueBlock value={proposed} />
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
}

function WireframeDiff({
  currentHTML,
  currentSource,
  proposedHTML,
  proposedSource,
}: {
  currentHTML?: string;
  currentSource: any;
  proposedHTML?: string;
  proposedSource: string;
}) {
  const [tab, setTab] = useState<'current' | 'proposed'>('current');
  const [showSource, setShowSource] = useState(false);
  const html = tab === 'current' ? currentHTML : proposedHTML;
  const source = tab === 'current' ? currentSource : proposedSource;

  return (
    <div className="p-3 space-y-3">
      <div className="flex items-center justify-between gap-3">
        <div className="inline-flex rounded border border-border overflow-hidden text-xs">
          <button
            type="button"
            onClick={() => setTab('current')}
            className={`px-3 py-1.5 ${
              tab === 'current'
                ? 'bg-muted text-foreground'
                : 'text-muted-foreground hover:text-foreground'
            }`}
          >
            Current
          </button>
          <button
            type="button"
            onClick={() => setTab('proposed')}
            className={`px-3 py-1.5 border-l border-border ${
              tab === 'proposed'
                ? 'bg-muted text-foreground'
                : 'text-muted-foreground hover:text-foreground'
            }`}
          >
            Proposed
          </button>
        </div>
        <button
          type="button"
          onClick={() => setShowSource((v) => !v)}
          className="text-[11px] text-muted-foreground hover:text-foreground"
        >
          {showSource ? 'Hide source' : 'View source'}
        </button>
      </div>
      {tab === 'proposed' && proposedSource === 'DELETE' ? (
        <span className="text-red-400 line-through italic">(remove)</span>
      ) : showSource || !html ? (
        <ValueBlock value={source} />
      ) : (
        <div
          className="rounded border border-border bg-card p-3 overflow-x-auto text-xs"
          dangerouslySetInnerHTML={{ __html: html }}
        />
      )}
    </div>
  );
}

function ValueBlock({ value }: { value: any }) {
  if (value === undefined || value === null) {
    return <span className="italic text-muted-foreground">(not set)</span>;
  }
  if (Array.isArray(value)) {
    if (value.length === 0) {
      return <span className="italic text-muted-foreground">(empty list)</span>;
    }
    return (
      <ul className="space-y-1 font-mono text-[11px]">
        {value.map((item, i) => (
          <li key={i} className="flex gap-1.5">
            <span className="text-muted-foreground">-</span>
            <span className="min-w-0 break-words">{stringify(item)}</span>
          </li>
        ))}
      </ul>
    );
  }
  if (typeof value === 'object') {
    return (
      <pre className="font-mono text-[11px] whitespace-pre-wrap break-words">
        {JSON.stringify(value, null, 2)}
      </pre>
    );
  }
  const str = String(value);
  return (
    <pre className="font-mono text-[11px] whitespace-pre-wrap break-words">{str}</pre>
  );
}

function stringify(v: any): string {
  if (v === null || v === undefined) return '';
  if (typeof v === 'object') return JSON.stringify(v);
  return String(v);
}
