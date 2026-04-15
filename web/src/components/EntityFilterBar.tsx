import { useEffect, useMemo, useRef, useState } from 'react';
import { EntitySummary } from '../lib/api';
import { SearchIcon } from './icons';

// ---------------------------------------------------------------------------
// Query DSL
// ---------------------------------------------------------------------------
//
// Free text → substring match against name / description.
// `key:value` tokens → field-scoped predicates. `-key:value` negates.
// Quoted values (`tag:"big idea"`) keep spaces. Multiple tokens are AND-ed.
//
// Recognised keys:
//   name:foo                → name substring
//   desc:foo / description: → description substring
//   tag:foo                 → has tag (exact)
//   kind:component          → entity kind
//   id:COM-0001             → exact entity id
//   slug:cli                → slug substring
//   file:internal/cli/      → any file path substring
//   has:files               → at least one file
//   has:rels
//   rel:depends_on          → has any outbound relationship of this type
//   belongs_to:syde-cli     → outbound belongs_to with this target
//   depends_on:storage      → outbound depends_on with this target
//   references:cli          → outbound references with this target
//   exposes / consumes / involves / applies_to / relates_to / implements
//                            same shape as belongs_to / depends_on
//   contract_kind:rest      → contract's contract_kind field (exact)
//   pattern:request-response → contract's interaction_pattern field (exact)
//
// Unknown keys fall back to substring on the joined searchable text.

export interface FilterToken {
  key: string | null;
  value: string;
  negate: boolean;
  // Position in the input string so the caret-aware suggestions know
  // which token the user is currently editing.
  start: number;
  end: number;
}

export function parseQuery(input: string): FilterToken[] {
  const tokens: FilterToken[] = [];
  let i = 0;
  while (i < input.length) {
    while (i < input.length && /\s/.test(input[i])) i++;
    if (i >= input.length) break;
    const start = i;
    let negate = false;
    if (input[i] === '-') {
      negate = true;
      i++;
    }
    // Read raw word until first whitespace or end, with quoted-string support
    // for the value half of key:value (including key:"value with spaces").
    let raw = '';
    let inQuotes = false;
    while (i < input.length) {
      const c = input[i];
      if (c === '"') {
        inQuotes = !inQuotes;
        i++;
        continue;
      }
      if (!inQuotes && /\s/.test(c)) break;
      raw += c;
      i++;
    }
    const colon = raw.indexOf(':');
    if (colon > 0) {
      tokens.push({
        key: raw.slice(0, colon).toLowerCase(),
        value: raw.slice(colon + 1),
        negate,
        start,
        end: i,
      });
    } else {
      tokens.push({ key: null, value: raw, negate, start, end: i });
    }
  }
  return tokens;
}

// Every key the DSL knows how to evaluate. Tokens whose key isn't in this
// set are NOT auto-committed into a pill on space — the user is probably
// mid-typo and we don't want to swallow characters into a non-functional
// pill.
const KNOWN_KEYS = new Set([
  'name',
  'desc',
  'description',
  'slug',
  'id',
  'kind',
  'tag',
  'file',
  'has',
  'rel',
  'belongs_to',
  'depends_on',
  'references',
  'exposes',
  'consumes',
  'involves',
  'applies_to',
  'relates_to',
  'implements',
  'modifies',
  'visualizes',
  'uses',
  'contract_kind',
  'pattern',
]);

const REL_TYPES = new Set([
  'belongs_to',
  'depends_on',
  'references',
  'exposes',
  'consumes',
  'involves',
  'applies_to',
  'relates_to',
  'implements',
  'modifies',
  'visualizes',
  'uses',
]);

export function applyFilter(entities: EntitySummary[], query: string): EntitySummary[] {
  const tokens = parseQuery(query).filter((t) => t.value.length > 0 || t.key);
  if (tokens.length === 0) return entities;

  return entities.filter((e) => {
    for (const tok of tokens) {
      const ok = matchToken(e, tok);
      if (tok.negate ? ok : !ok) return false;
    }
    return true;
  });
}

function matchToken(e: EntitySummary, tok: FilterToken): boolean {
  const v = tok.value.toLowerCase();
  const key = tok.key;

  // Free text → name OR description OR slug substring.
  if (!key) {
    if (!v) return true;
    return (
      e.name.toLowerCase().includes(v) ||
      (e.description || '').toLowerCase().includes(v) ||
      (e.slug || '').toLowerCase().includes(v)
    );
  }

  switch (key) {
    case 'name':
      return e.name.toLowerCase().includes(v);
    case 'desc':
    case 'description':
      return (e.description || '').toLowerCase().includes(v);
    case 'slug':
      return (e.slug || '').toLowerCase().includes(v);
    case 'id':
      return e.id.toLowerCase() === v;
    case 'kind':
      return e.kind === v;
    case 'tag':
      return (e.tags || []).some((t) => t.toLowerCase() === v);
    case 'file':
      return (e.files || []).some((f) => f.toLowerCase().includes(v));
    case 'has':
      if (v === 'files') return (e.files || []).length > 0;
      if (v === 'tags') return (e.tags || []).length > 0;
      if (v === 'rels' || v === 'relationships') return (e.relationships || []).length > 0;
      if (v === 'description') return !!e.description;
      return false;
    case 'rel':
      return (e.relationships || []).some((r) => r.type === v);
    case 'contract_kind':
      return (e.contract_kind || '').toLowerCase() === v;
    case 'pattern':
      return (e.interaction_pattern || '').toLowerCase() === v;
  }

  // Relationship type as the key itself, e.g. `belongs_to:syde-cli`.
  if (REL_TYPES.has(key)) {
    return (e.relationships || []).some(
      (r) => r.type === key && (!v || r.target.toLowerCase().includes(v)),
    );
  }

  // Unknown key → fall back to free-text match on the value.
  return (
    e.name.toLowerCase().includes(v) ||
    (e.description || '').toLowerCase().includes(v)
  );
}

// ---------------------------------------------------------------------------
// Suggestion engine
// ---------------------------------------------------------------------------

export interface Suggestion {
  // The chunk that replaces the active token in the input.
  insert: string;
  // What's shown to the user.
  label: string;
  hint?: string;
}

const KEY_SUGGESTIONS: { key: string; hint: string }[] = [
  { key: 'name:', hint: 'name substring' },
  { key: 'desc:', hint: 'description substring' },
  { key: 'tag:', hint: 'has tag' },
  { key: 'kind:', hint: 'entity kind' },
  { key: 'slug:', hint: 'slug substring' },
  { key: 'file:', hint: 'file path substring' },
  { key: 'contract_kind:', hint: 'contract_kind field (rest, cli, event, storage, …)' },
  { key: 'pattern:', hint: 'interaction_pattern field (request-response, pub-sub, schema, …)' },
  { key: 'belongs_to:', hint: 'outbound belongs_to target' },
  { key: 'depends_on:', hint: 'outbound depends_on target' },
  { key: 'references:', hint: 'outbound references target' },
  { key: 'exposes:', hint: 'outbound exposes target' },
  { key: 'involves:', hint: 'outbound involves target' },
  { key: 'applies_to:', hint: 'outbound applies_to target' },
  { key: 'rel:', hint: 'has any rel of type' },
  { key: 'has:', hint: 'has any (files|tags|rels)' },
];

const HAS_VALUES = ['files', 'tags', 'rels', 'description'];

function suggestionsForToken(
  token: FilterToken | null,
  entities: EntitySummary[],
): Suggestion[] {
  if (!token || (!token.key && token.value.length === 0)) {
    return KEY_SUGGESTIONS.map((s) => ({ insert: s.key, label: s.key, hint: s.hint }));
  }

  // Still typing the key — no colon yet.
  if (token.key === null) {
    const v = token.value.toLowerCase();
    const matches = KEY_SUGGESTIONS.filter((s) => s.key.startsWith(v));
    if (matches.length > 0) {
      return matches.map((s) => ({ insert: s.key, label: s.key, hint: s.hint }));
    }
    // Fall through to free-text completions on entity names.
    return uniqueValues(entities.map((e) => e.name))
      .filter((n) => n.toLowerCase().startsWith(v))
      .slice(0, 8)
      .map((n) => ({ insert: n, label: n, hint: 'name' }));
  }

  // Key is set — suggest values for that key.
  const k = token.key;
  const v = token.value.toLowerCase();
  const startsWith = (s: string) => s.toLowerCase().startsWith(v);

  if (k === 'kind') {
    const kinds = ['system', 'component', 'contract', 'concept', 'flow', 'decision', 'plan', 'task', 'design'];
    return kinds.filter(startsWith).map((kind) => ({
      insert: `kind:${kind}`,
      label: kind,
      hint: 'kind',
    }));
  }
  if (k === 'has') {
    return HAS_VALUES.filter(startsWith).map((val) => ({
      insert: `has:${val}`,
      label: val,
      hint: 'has',
    }));
  }
  if (k === 'rel') {
    return Array.from(REL_TYPES)
      .filter(startsWith)
      .sort()
      .map((t) => ({ insert: `rel:${t}`, label: t, hint: 'rel type' }));
  }
  if (k === 'tag') {
    const tags = uniqueValues(entities.flatMap((e) => e.tags || []));
    return tags
      .filter(startsWith)
      .slice(0, 12)
      .map((t) => ({ insert: `tag:${t}`, label: t, hint: 'tag' }));
  }
  if (k === 'contract_kind') {
    // Populate from the actual loaded contracts so newly-introduced
    // kinds like 'storage' show up automatically without hardcoding.
    const kinds = uniqueValues(entities.map((e) => e.contract_kind || ''));
    return kinds
      .filter(startsWith)
      .sort()
      .slice(0, 12)
      .map((val) => ({ insert: `contract_kind:${val}`, label: val, hint: 'contract_kind' }));
  }
  if (k === 'pattern') {
    const patterns = uniqueValues(entities.map((e) => e.interaction_pattern || ''));
    return patterns
      .filter(startsWith)
      .sort()
      .slice(0, 12)
      .map((val) => ({ insert: `pattern:${val}`, label: val, hint: 'interaction_pattern' }));
  }
  if (REL_TYPES.has(k)) {
    // Suggest entity slugs that exist as outbound targets of this rel type
    // anywhere in the project, plus any entity slug as a fallback.
    const seen = new Set<string>();
    for (const e of entities) {
      for (const r of e.relationships || []) {
        if (r.type === k) seen.add(r.target);
      }
    }
    const fromRels = Array.from(seen).filter(startsWith).sort();
    if (fromRels.length > 0) {
      return fromRels.slice(0, 12).map((target) => ({
        insert: `${k}:${target}`,
        label: target,
        hint: k.replace('_', ' '),
      }));
    }
    // Fall back to entity slugs.
    return entities
      .map((e) => e.slug || e.name.toLowerCase().replace(/\s+/g, '-'))
      .filter(startsWith)
      .slice(0, 12)
      .map((slug) => ({ insert: `${k}:${slug}`, label: slug, hint: k.replace('_', ' ') }));
  }
  // Generic — suggest entity names that match.
  const names = uniqueValues(entities.map((e) => e.name));
  return names
    .filter((n) => n.toLowerCase().includes(v))
    .slice(0, 8)
    .map((n) => ({ insert: `${k}:${n}`, label: n, hint: k }));
}

function uniqueValues(values: string[]): string[] {
  return Array.from(new Set(values.filter(Boolean)));
}

// ---------------------------------------------------------------------------
// Token serialization + colour palette
// ---------------------------------------------------------------------------

function tokenToString(t: FilterToken): string {
  const prefix = t.negate ? '-' : '';
  if (!t.key) return prefix + (t.value.includes(' ') ? `"${t.value}"` : t.value);
  const v = t.value.includes(' ') ? `"${t.value}"` : t.value;
  return `${prefix}${t.key}:${v}`;
}

const KIND_PILL_PALETTE: Record<string, string> = {
  system: 'border-kind-system/40 bg-kind-system/10 text-kind-system',
  component: 'border-kind-component/40 bg-kind-component/10 text-kind-component',
  contract: 'border-kind-contract/40 bg-kind-contract/10 text-kind-contract',
  concept: 'border-kind-concept/40 bg-kind-concept/10 text-kind-concept',
  flow: 'border-kind-flow/40 bg-kind-flow/10 text-kind-flow',
  decision: 'border-kind-decision/40 bg-kind-decision/10 text-kind-decision',
  plan: 'border-kind-plan/40 bg-kind-plan/10 text-kind-plan',
  task: 'border-kind-task/40 bg-kind-task/10 text-kind-task',
};

// Each relationship type rides the colour of the kind it most often points at.
const REL_PILL_PALETTE: Record<string, string> = {
  belongs_to: 'border-kind-system/40 bg-kind-system/10 text-kind-system',
  depends_on: 'border-kind-component/40 bg-kind-component/10 text-kind-component',
  exposes: 'border-kind-contract/40 bg-kind-contract/10 text-kind-contract',
  consumes: 'border-kind-contract/40 bg-kind-contract/10 text-kind-contract',
  references: 'border-kind-concept/40 bg-kind-concept/10 text-kind-concept',
  relates_to: 'border-kind-concept/40 bg-kind-concept/10 text-kind-concept',
  implements: 'border-kind-concept/40 bg-kind-concept/10 text-kind-concept',
  involves: 'border-kind-flow/40 bg-kind-flow/10 text-kind-flow',
  applies_to: 'border-kind-decision/40 bg-kind-decision/10 text-kind-decision',
  modifies: 'border-border bg-card text-foreground/80',
  visualizes: 'border-border bg-card text-foreground/80',
  uses: 'border-border bg-card text-foreground/80',
};

const NEUTRAL_PILL = 'border-border bg-card text-foreground/80';

function paletteForToken(t: FilterToken): string {
  if (t.negate) {
    return 'border-red-500/40 bg-red-500/10 text-red-300/90';
  }
  if (!t.key) return NEUTRAL_PILL;
  if (t.key === 'kind' && KIND_PILL_PALETTE[t.value.toLowerCase()]) {
    return KIND_PILL_PALETTE[t.value.toLowerCase()];
  }
  if (REL_PILL_PALETTE[t.key]) return REL_PILL_PALETTE[t.key];
  if (t.key === 'rel' && REL_PILL_PALETTE[t.value]) return REL_PILL_PALETTE[t.value];
  return NEUTRAL_PILL;
}

// ---------------------------------------------------------------------------
// Component
// ---------------------------------------------------------------------------

interface EntityFilterBarProps {
  value: string;
  onChange: (value: string) => void;
  entities: EntitySummary[];
  placeholder?: string;
  // Larger, borderless rendering used by the command palette where the
  // bar IS the chrome of its host card.
  large?: boolean;
  // Focus the input on mount. Used by the command palette so Cmd+K lands
  // straight in the search bar.
  autoFocus?: boolean;
}

export function EntityFilterBar({ value, onChange, entities, placeholder, large, autoFocus }: EntityFilterBarProps) {
  const inputRef = useRef<HTMLInputElement | null>(null);

  // Internal state: a list of committed pills + the draft string the user
  // is currently typing. The combined serialisation flows out to onChange so
  // the parent always sees one query string.
  const [pills, setPills] = useState<FilterToken[]>(() => parseQuery(value));
  const [draft, setDraft] = useState('');
  const [open, setOpen] = useState(false);
  // Suggestions only auto-show once the user has typed something. Empty
  // input → no popover. Ctrl/Cmd + Space forces it open even when empty.
  const [forceOpen, setForceOpen] = useState(false);
  const [activeIdx, setActiveIdx] = useState(0);

  // Recombine pills + draft and emit it as the canonical query string.
  const combined = useMemo(() => {
    const parts = pills.map(tokenToString);
    if (draft) parts.push(draft);
    return parts.join(' ');
  }, [pills, draft]);

  // Push the combined string up whenever it changes.
  const lastEmittedRef = useRef(combined);
  useEffect(() => {
    if (combined !== lastEmittedRef.current) {
      lastEmittedRef.current = combined;
      onChange(combined);
    }
  }, [combined, onChange]);

  // Re-seed from external value (e.g. URL ?filter=… changes) when it
  // diverges from what we last emitted.
  useEffect(() => {
    if (value !== combined) {
      lastEmittedRef.current = value;
      setPills(parseQuery(value));
      setDraft('');
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [value]);

  // Suggestion engine works on the in-progress draft as a single token.
  const draftToken = useMemo(() => {
    const tok = parseQuery(draft)[0];
    return tok || ({ key: null, value: '', negate: false, start: 0, end: 0 } as FilterToken);
  }, [draft]);

  const suggestions = useMemo(
    () => suggestionsForToken(draftToken, entities),
    [draftToken, entities],
  );

  useEffect(() => {
    setActiveIdx((i) => Math.min(i, Math.max(0, suggestions.length - 1)));
  }, [suggestions]);

  // Once the user actually types, the forced-open hint isn't needed anymore;
  // the normal "draft has content" rule takes over.
  useEffect(() => {
    if (draft.length > 0 && forceOpen) setForceOpen(false);
  }, [draft, forceOpen]);

  // Focus the input on mount when requested. The command palette uses this
  // so opening it lands you typing immediately. requestAnimationFrame
  // ensures the input is in the DOM before .focus() runs.
  useEffect(() => {
    if (!autoFocus) return;
    const id = requestAnimationFrame(() => inputRef.current?.focus());
    return () => cancelAnimationFrame(id);
  }, [autoFocus]);

  const showDropdown = open && (draft.length > 0 || forceOpen) && suggestions.length > 0;

  // A token only auto-commits to a pill when it's a structured `key:value`
  // pair with a recognised key and a non-empty value. Free text (no key)
  // stays in the draft so the user can keep typing multi-word substring
  // searches; pressing Enter still commits the whole query for filtering.
  // Bare `key:` and `unknownkey:foo` are also rejected so the user doesn't
  // lose characters mid-typo.
  function isValidToken(t: FilterToken): boolean {
    if (t.value.length === 0) return false;
    if (t.key === null) return false;
    return KNOWN_KEYS.has(t.key);
  }

  // Try to commit the current draft into pills. Only valid tokens are
  // promoted; anything left over (bare keys, unknown keys, partial input)
  // stays in the draft.
  function commitDraft(text?: string) {
    const src = text ?? draft;
    if (!src.trim()) return false;
    const toks = parseQuery(src);
    if (toks.length === 0) return false;
    const newPills: FilterToken[] = [];
    let newDraft = '';
    for (const t of toks) {
      if (isValidToken(t)) {
        newPills.push(t);
      } else {
        // First non-pill token becomes the new draft (preserve everything
        // from this token onward so the user keeps their typing position).
        newDraft = src.slice(t.start).trimEnd();
        break;
      }
    }
    if (newPills.length === 0) return false;
    setPills((prev) => [...prev, ...newPills]);
    setDraft(newDraft);
    return true;
  }

  function applySuggestion(s: Suggestion) {
    if (s.insert.endsWith(':')) {
      // Picked a key — keep it in the draft so the user can add a value.
      const insert = draftToken.negate ? `-${s.insert}` : s.insert;
      setDraft(insert);
      requestAnimationFrame(() => {
        inputRef.current?.focus();
        const len = insert.length;
        inputRef.current?.setSelectionRange(len, len);
      });
      return;
    }
    // Complete suggestion → split on the FIRST colon so we can quote the
    // value half if it contains spaces. parseQuery would pre-split on the
    // unquoted space and lose the rest of the value, so we sidestep it here.
    const colon = s.insert.indexOf(':');
    let insert: string;
    if (colon > 0) {
      const k = s.insert.slice(0, colon);
      const v = s.insert.slice(colon + 1);
      const quoted = /\s/.test(v) ? `"${v}"` : v;
      insert = `${k}:${quoted}`;
    } else {
      insert = s.insert;
    }
    if (draftToken.negate) insert = `-${insert}`;
    commitDraft(insert);
    requestAnimationFrame(() => inputRef.current?.focus());
  }

  function handleKeyDown(ev: React.KeyboardEvent<HTMLInputElement>) {
    // Ctrl/Cmd + Space → open the suggestions even when the input is empty.
    if ((ev.ctrlKey || ev.metaKey) && ev.key === ' ') {
      ev.preventDefault();
      setForceOpen(true);
      setOpen(true);
      return;
    }
    if (ev.key === 'Enter') {
      ev.preventDefault();
      if (!commitDraft()) {
        setOpen(false);
        setForceOpen(false);
      }
      return;
    }
    if (ev.key === ' ') {
      // Don't commit if the user is mid-quote — they're typing a value with
      // spaces (`slug:"CLI Commands"`) and the trailing quote hasn't landed
      // yet. Count `"` chars in the draft; an odd count means there's an
      // unmatched opening quote.
      const openQuote = (draft.match(/"/g) || []).length % 2 === 1;
      if (openQuote) return;
      // Commit on space. If the draft is just `key:` (no value yet) leave it
      // in place — commitDraft returns false in that case.
      if (commitDraft()) {
        ev.preventDefault();
      }
      return;
    }
    if (ev.key === 'Backspace' && draft.length === 0 && pills.length > 0) {
      ev.preventDefault();
      // Pop the last pill back into the draft so the user can edit it.
      const last = pills[pills.length - 1];
      setPills((prev) => prev.slice(0, -1));
      setDraft(tokenToString(last));
      return;
    }
    if (ev.key === 'Escape') {
      setOpen(false);
      setForceOpen(false);
      return;
    }
    if (!open || suggestions.length === 0) return;
    if (ev.key === 'ArrowDown') {
      ev.preventDefault();
      setActiveIdx((i) => (i + 1) % suggestions.length);
    } else if (ev.key === 'ArrowUp') {
      ev.preventDefault();
      setActiveIdx((i) => (i - 1 + suggestions.length) % suggestions.length);
    } else if (ev.key === 'Tab' && suggestions[activeIdx]) {
      ev.preventDefault();
      applySuggestion(suggestions[activeIdx]);
    }
  }

  function removePill(idx: number) {
    setPills((prev) => prev.filter((_, i) => i !== idx));
    inputRef.current?.focus();
  }

  const wrapperClass = large
    ? 'flex items-center flex-wrap gap-2 px-4 bg-transparent'
    : 'flex items-center flex-wrap gap-1 px-1.5 bg-card border border-border rounded-md focus-within:border-muted-foreground transition-colors';
  const wrapperStyle = large ? { minHeight: 56 } : { minHeight: 30 };
  const inputStyle = large
    ? { fontSize: 16, lineHeight: 1.2 }
    : { fontSize: 10, lineHeight: 1 };
  const iconSize = large ? 'w-4 h-4 mr-2' : 'w-2.5 h-2.5 ml-0.5 mr-0.5';

  return (
    <div className={`relative ${large ? '' : 'mb-4'}`}>
      <div
        className={wrapperClass}
        style={wrapperStyle}
        onClick={() => inputRef.current?.focus()}
      >
        <SearchIcon className={`text-muted-foreground shrink-0 ${iconSize}`} />
        {pills.map((p, i) => (
          <FilterPill key={`${tokenToString(p)}-${i}`} token={p} onRemove={() => removePill(i)} />
        ))}
        <input
          ref={inputRef}
          type="text"
          value={draft}
          onChange={(ev) => {
            setDraft(ev.target.value);
            setOpen(true);
          }}
          onKeyDown={handleKeyDown}
          onFocus={() => setOpen(true)}
          onBlur={() => {
            // Commit any in-progress complete token then close suggestions.
            commitDraft();
            setTimeout(() => {
              setOpen(false);
              setForceOpen(false);
            }, 120);
          }}
          placeholder={pills.length === 0 ? (placeholder || 'filter — try kind:component or belongs_to:syde-cli') : ''}
          spellCheck={false}
          style={inputStyle}
          className={`flex-1 min-w-[80px] bg-transparent border-0 outline-none ${
            large ? 'font-sans' : 'font-mono'
          } text-foreground placeholder:text-muted-foreground/70 ${
            large ? 'px-1 py-3' : 'px-1 py-1'
          }`}
        />
        {pills.length > 0 && (
          <button
            type="button"
            onClick={() => {
              setPills([]);
              setDraft('');
              inputRef.current?.focus();
            }}
            title="Clear all"
            className="text-muted-foreground/40 hover:text-muted-foreground px-1 shrink-0 leading-none"
            style={{ fontSize: '9px' }}
          >
            ✕
          </button>
        )}
      </div>

      {showDropdown && (
        <div
          className={
            large
              ? 'border-t border-white/10 max-h-56 overflow-y-auto bg-black'
              : 'absolute z-30 mt-1.5 left-0 right-0 max-h-72 overflow-y-auto rounded-md border border-border bg-background/98 backdrop-blur-sm shadow-2xl'
          }
        >
          <ul className={large ? 'py-1.5' : 'py-1'}>
            {suggestions.map((s, i) => (
              <li key={`${s.insert}-${i}`}>
                <button
                  type="button"
                  onMouseDown={(ev) => {
                    ev.preventDefault();
                    applySuggestion(s);
                  }}
                  onMouseEnter={() => setActiveIdx(i)}
                  className={`w-full text-left flex items-center gap-3 ${
                    large ? 'px-5 py-1.5' : 'px-2.5 py-1'
                  } ${i === activeIdx ? (large ? 'bg-white/[0.06]' : 'bg-accent/60') : (large ? 'hover:bg-white/[0.03]' : 'hover:bg-accent/30')}`}
                >
                  <span className={`font-mono text-foreground ${large ? 'text-[12px]' : 'text-[10px]'}`}>
                    {s.label}
                  </span>
                  {s.hint && (
                    <span className={`ml-auto text-muted-foreground ${large ? 'text-[10px]' : 'text-[9px]'}`}>
                      {s.hint}
                    </span>
                  )}
                </button>
              </li>
            ))}
          </ul>
          <div
            className={`border-t flex justify-between ${
              large
                ? 'border-white/10 px-5 py-1.5 text-[10px] text-muted-foreground'
                : 'border-border/60 px-2.5 py-1 text-[9px] text-muted-foreground'
            }`}
          >
            <span>↑↓ navigate · tab complete · space pill · ⏎ search · esc close</span>
            <span>⌃␣ open · prefix - to negate</span>
          </div>
        </div>
      )}
    </div>
  );
}

function FilterPill({ token, onRemove }: { token: FilterToken; onRemove: () => void }) {
  const palette = paletteForToken(token);
  return (
    <span
      className={`inline-flex items-center gap-0.5 rounded border px-1.5 py-[3px] text-[10px] font-mono leading-none ${palette}`}
    >
      {token.negate && <span className="opacity-70">−</span>}
      {token.key && (
        <>
          <span className="opacity-70">{token.key}</span>
          <span className="opacity-50">:</span>
        </>
      )}
      <span>{token.value}</span>
      <button
        type="button"
        onMouseDown={(ev) => {
          ev.preventDefault();
          onRemove();
        }}
        className="opacity-50 hover:opacity-100 ml-0.5 -mr-0.5"
        aria-label="remove"
      >
        ✕
      </button>
    </span>
  );
}
