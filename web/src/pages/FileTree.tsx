import { useEffect, useMemo, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { Highlight, themes, type Language } from 'prism-react-renderer';
import { useApi } from '../hooks/useApi';
import { api, TreeResponse, TreeContextBundle } from '../lib/api';
import {
  ChevronDownIcon,
  ChevronRightIcon,
  FileIcon,
  FolderClosedIcon,
  FolderOpenIcon,
} from '../components/icons';

export function FileTree() {
  const location = useLocation();
  // Initial selection comes from ?path=… so opening a file from an entity's
  // Files list deep-links into the tree at that node.
  const initialPath = useMemo(() => {
    const params = new URLSearchParams(location.search);
    return params.get('path') || '.';
  }, [location.search]);
  const { data: tree, loading, error } = useApi<TreeResponse>(() => api.tree(), []);
  const [selected, setSelected] = useState<string>(initialPath);
  const [expanded, setExpanded] = useState<Record<string, boolean>>(() =>
    expandAncestors(initialPath, { '.': true }),
  );

  // If the URL changes (e.g. the user opens a different file from another
  // entity card without leaving the page), re-sync.
  useEffect(() => {
    if (initialPath && initialPath !== selected) {
      setSelected(initialPath);
      setExpanded((prev) => expandAncestors(initialPath, prev));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [initialPath]);

  const { data: bundle, loading: bundleLoading } = useApi<TreeContextBundle>(
    () => (selected ? api.treeNode(selected) : Promise.resolve(null as any)),
    [selected]
  );

  const rootNode = tree?.nodes['.'];

  const counts = useMemo(() => {
    let files = 0, dirs = 0, stale = 0, ignored = 0;
    if (tree) {
      for (const n of Object.values(tree.nodes)) {
        if (n.type === 'file') files++;
        else dirs++;
        if (n.summary_stale) stale++;
        if (n.ignored) ignored++;
      }
    }
    return { files, dirs, stale, ignored };
  }, [tree]);

  if (loading) {
    return <div className="h-full flex items-center justify-center text-muted-foreground text-sm">Loading tree…</div>;
  }
  if (error || !tree || !rootNode) {
    return (
      <div className="h-full flex items-center justify-center text-sm text-muted-foreground">
        No tree available. Run <code className="font-mono text-foreground mx-1">syde tree scan</code> to generate it.
      </div>
    );
  }

  return (
    <div className="h-full flex flex-col min-h-0">
      {/* Header */}
      <div className="px-6 py-4 border-b border-border shrink-0">
        <div className="flex items-baseline gap-3">
          <h1 className="text-lg font-semibold">File Tree</h1>
          <div className="flex items-center gap-3 text-[11px] text-muted-foreground">
            <span><span className="text-foreground font-medium">{counts.files}</span> files</span>
            <span className="text-border">·</span>
            <span><span className="text-foreground font-medium">{counts.dirs}</span> dirs</span>
            {counts.stale > 0 && (
              <>
                <span className="text-border">·</span>
                <span className="text-orange-400">{counts.stale} stale</span>
              </>
            )}
            {counts.ignored > 0 && (
              <>
                <span className="text-border">·</span>
                <span>{counts.ignored} ignored</span>
              </>
            )}
            {tree.scanned_at && (
              <>
                <span className="text-border">·</span>
                <span>scanned {formatTime(tree.scanned_at)}</span>
              </>
            )}
          </div>
        </div>
      </div>

      {/* Body: 2 columns, full height */}
      <div className="flex-1 min-h-0 grid grid-cols-[minmax(280px,360px)_minmax(0,1fr)]">
        {/* Tree pane */}
        <div className="border-r border-border overflow-y-auto py-2 bg-background/40">
          <TreeRow
            tree={tree}
            path="."
            depth={0}
            ancestorIsLast={[]}
            isLast
            expanded={expanded}
            setExpanded={setExpanded}
            selected={selected}
            onSelect={setSelected}
          />
        </div>

        {/* Detail pane */}
        <div className="overflow-y-auto">
          {bundleLoading && !bundle ? (
            <div className="h-full flex items-center justify-center text-muted-foreground text-sm">Loading…</div>
          ) : bundle ? (
            <ContextView bundle={bundle} />
          ) : (
            <div className="h-full flex items-center justify-center text-muted-foreground text-sm">
              Select a file or folder…
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

interface TreeRowProps {
  tree: TreeResponse;
  path: string;
  depth: number;
  // For each ancestor depth, true when that ancestor was the last child of
  // its parent. Drives whether we draw a continuing vertical line at that
  // column or leave it blank.
  ancestorIsLast: boolean[];
  isLast: boolean;
  expanded: Record<string, boolean>;
  setExpanded: (e: Record<string, boolean>) => void;
  selected: string;
  onSelect: (path: string) => void;
}

function TreeRow({
  tree,
  path,
  depth,
  ancestorIsLast,
  isLast,
  expanded,
  setExpanded,
  selected,
  onSelect,
}: TreeRowProps) {
  const node = tree.nodes[path];
  if (!node) return null;
  const name = path === '.' ? 'project root' : path.split('/').pop() || path;
  const isDir = node.type === 'dir';
  const isOpen = !!expanded[path];
  const isSelected = selected === path;
  const isRoot = depth === 0;

  // Sort children: dirs first, then files, alphabetical within group.
  const children = (node.children || []).slice().sort((a, b) => {
    const na = tree.nodes[a];
    const nb = tree.nodes[b];
    if (!na || !nb) return 0;
    if (na.type !== nb.type) return na.type === 'dir' ? -1 : 1;
    return a.localeCompare(b);
  });

  return (
    <div>
      <button
        type="button"
        onClick={() => {
          onSelect(path);
          if (isDir) setExpanded({ ...expanded, [path]: !isOpen });
        }}
        className={`group w-full flex items-center text-left text-[12px] h-7 pr-3 transition-colors ${
          isSelected ? 'bg-accent/60' : 'hover:bg-accent/30'
        } ${node.ignored ? 'opacity-40 hover:opacity-100' : ''}`}
      >
        {/* Tree connector columns. One <span> per ancestor depth. */}
        {!isRoot && (
          <>
            {ancestorIsLast.map((wasLast, i) => (
              <span
                key={i}
                className="relative shrink-0 self-stretch"
                style={{ width: 16 }}
              >
                {!wasLast && (
                  <span className="absolute inset-y-0 left-1/2 -translate-x-px w-px bg-border" />
                )}
              </span>
            ))}
            {/* The connector for THIS row: a vertical line that either
                stops at the middle (last child) or continues, plus a
                horizontal nub into the icon. */}
            <span className="relative shrink-0 self-stretch" style={{ width: 16 }}>
              <span
                className={`absolute left-1/2 -translate-x-px w-px bg-border ${
                  isLast ? 'top-0 h-1/2' : 'inset-y-0'
                }`}
              />
              <span className="absolute top-1/2 left-1/2 h-px w-1/2 bg-border" />
            </span>
          </>
        )}

        {/* Chevron (dirs only) */}
        <span className="shrink-0 w-4 h-4 flex items-center justify-center text-muted-foreground">
          {isDir ? (
            isOpen ? <ChevronDownIcon className="w-3 h-3" /> : <ChevronRightIcon className="w-3 h-3" />
          ) : null}
        </span>

        {/* Icon */}
        <span className="shrink-0 w-4 h-4 mr-1.5 flex items-center justify-center">
          {isDir ? (
            isOpen ? (
              <FolderOpenIcon className="w-4 h-4 text-kind-system" />
            ) : (
              <FolderClosedIcon className="w-4 h-4 text-kind-system" />
            )
          ) : (
            <FileIcon className="w-3.5 h-3.5 text-muted-foreground" />
          )}
        </span>

        {/* Name */}
        <span className={`truncate ${isDir ? 'text-foreground font-medium' : 'text-foreground/90'}`}>
          {name}
        </span>

        {/* Badges. Ignored entries fade the whole row instead of carrying
            a loud badge — they're rarely worth eye contact. */}
        {node.summary_stale && !node.ignored && (
          <span className="ml-2 text-[9px] uppercase tracking-wider px-1.5 py-px rounded bg-orange-500/15 text-orange-400">
            stale
          </span>
        )}
      </button>
      {isDir && isOpen && children.map((c, i) => (
        <TreeRow
          key={c}
          tree={tree}
          path={c}
          depth={depth + 1}
          ancestorIsLast={isRoot ? [] : [...ancestorIsLast, isLast]}
          isLast={i === children.length - 1}
          expanded={expanded}
          setExpanded={setExpanded}
          selected={selected}
          onSelect={onSelect}
        />
      ))}
    </div>
  );
}

function ContextView({ bundle }: { bundle: TreeContextBundle }) {
  const isDir = bundle.type === 'dir';
  const segments = bundle.path === '.' ? ['project root'] : bundle.path.split('/');
  const name = segments[segments.length - 1];

  return (
    <div className="px-6 py-6 space-y-6 max-w-3xl">
      {/* Title block */}
      <div>
        <div className="flex items-center gap-1.5 text-[11px] text-muted-foreground mb-3 flex-wrap">
          {bundle.breadcrumb?.map((b) => (
            <span key={b.path} className="flex items-center gap-1.5">
              <span>{b.path === '.' ? 'project root' : b.path.split('/').pop()}</span>
              <ChevronRightIcon className="w-3 h-3 opacity-60" />
            </span>
          ))}
          <span className="text-foreground font-medium">{name}</span>
        </div>
        <div className="flex items-center gap-3">
          <span className="flex items-center justify-center w-9 h-9 rounded-lg border border-border bg-card">
            {isDir ? (
              <FolderOpenIcon className="w-5 h-5 text-kind-system" />
            ) : (
              <FileIcon className="w-5 h-5 text-muted-foreground" />
            )}
          </span>
          <div className="min-w-0">
            <h2 className="text-lg font-semibold truncate">{name}</h2>
            <div className="text-[11px] font-mono text-muted-foreground truncate">
              {bundle.path}
            </div>
          </div>
          {bundle.stale && (
            <span className="ml-auto text-[10px] uppercase tracking-wider px-2 py-0.5 rounded bg-orange-500/15 text-orange-400">
              stale
            </span>
          )}
        </div>
      </div>

      {/* Summary */}
      <Card title="Summary">
        {bundle.summary ? (
          <p className="text-sm text-foreground leading-relaxed">{bundle.summary}</p>
        ) : (
          <p className="text-sm text-muted-foreground italic">No summary yet — run syde tree summarize on this node.</p>
        )}
      </Card>

      {/* Breadcrumb summaries */}
      {bundle.breadcrumb && bundle.breadcrumb.length > 0 && (
        <Card title="Tree summary">
          <ol className="space-y-2.5">
            {bundle.breadcrumb.map((b, i) => (
              <li key={b.path} className="flex items-start gap-2 text-xs">
                <span
                  className="mt-1.5 inline-block h-1.5 w-1.5 rounded-full bg-muted-foreground/60 shrink-0"
                  style={{ marginLeft: i * 10 }}
                />
                <div className="min-w-0">
                  <div className="font-mono text-[11px] text-muted-foreground">
                    {b.path === '.' ? 'project root' : b.path}
                  </div>
                  {b.summary && (
                    <div className="text-foreground/90 leading-relaxed">{b.summary}</div>
                  )}
                </div>
              </li>
            ))}
          </ol>
        </Card>
      )}

      {/* Folder children list */}
      {isDir && bundle.children && bundle.children.length > 0 && (
        <Card title={`Children · ${bundle.children.length}`}>
          <ul className="divide-y divide-border/60">
            {bundle.children.map((c) => (
              <li key={c.path} className="flex items-start gap-2.5 py-2 first:pt-0 last:pb-0">
                <span className="mt-px shrink-0">
                  {c.type === 'dir' ? (
                    <FolderClosedIcon className="w-4 h-4 text-kind-system" />
                  ) : (
                    <FileIcon className="w-3.5 h-3.5 text-muted-foreground" />
                  )}
                </span>
                <div className="min-w-0">
                  <div className="font-mono text-[11px] text-foreground">
                    {c.path.split('/').pop()}
                  </div>
                  {c.summary && (
                    <div className="text-xs text-muted-foreground leading-relaxed mt-0.5">
                      {c.summary}
                    </div>
                  )}
                </div>
                {c.stale && (
                  <span className="ml-auto text-[9px] uppercase tracking-wider px-1.5 py-px rounded bg-orange-500/15 text-orange-400 shrink-0">
                    stale
                  </span>
                )}
              </li>
            ))}
          </ul>
        </Card>
      )}

      {/* File content — CodeViewer brings its own panel chrome, no Card needed. */}
      {bundle.type === 'file' && bundle.content && (
        <section>
          <h3 className="text-[10px] font-medium uppercase tracking-widest text-muted-foreground mb-2">
            Content{bundle.truncated ? ' · truncated' : ''}
          </h3>
          <CodeViewer code={bundle.content} path={bundle.path} />
        </section>
      )}
    </div>
  );
}

function Card({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <section>
      <h3 className="text-[10px] font-medium uppercase tracking-widest text-muted-foreground mb-2">
        {title}
      </h3>
      <div className="rounded-lg border border-border bg-card/40 p-4">
        {children}
      </div>
    </section>
  );
}

// Map filename extension / basename to a Prism language id. Anything
// unrecognised falls back to plain text, which still renders fine — just
// without highlighting.
const EXT_TO_LANG: Record<string, Language> = {
  ts: 'tsx',
  tsx: 'tsx',
  js: 'jsx',
  jsx: 'jsx',
  mjs: 'jsx',
  cjs: 'jsx',
  go: 'go',
  py: 'python',
  rs: 'rust',
  java: 'java',
  rb: 'ruby',
  php: 'php',
  swift: 'swift',
  kt: 'kotlin',
  sh: 'bash',
  bash: 'bash',
  zsh: 'bash',
  json: 'json',
  yaml: 'yaml',
  yml: 'yaml',
  toml: 'toml',
  md: 'markdown',
  mdx: 'markdown',
  html: 'markup',
  htm: 'markup',
  xml: 'markup',
  svg: 'markup',
  css: 'css',
  scss: 'scss',
  sql: 'sql',
  graphql: 'graphql',
  gql: 'graphql',
  dockerfile: 'docker',
  makefile: 'makefile',
};

const BASENAME_TO_LANG: Record<string, Language> = {
  Dockerfile: 'docker',
  Makefile: 'makefile',
  'go.mod': 'go',
  'go.sum': 'go',
  '.gitignore': 'bash',
};

function detectLanguage(path: string): Language {
  const base = path.split('/').pop() || path;
  if (BASENAME_TO_LANG[base]) return BASENAME_TO_LANG[base];
  const ext = base.includes('.') ? base.split('.').pop()!.toLowerCase() : '';
  return EXT_TO_LANG[ext] || ('text' as Language);
}

function CodeViewer({ code, path }: { code: string; path: string }) {
  const language = detectLanguage(path);
  // Trim trailing newlines so the gutter doesn't gain a phantom row.
  const source = code.replace(/\n+$/, '');

  return (
    <div className="rounded-md border border-border bg-[#0d0f14] overflow-hidden">
      {/* Tiny header strip showing the detected language */}
      <div className="flex items-center justify-between px-3 py-1.5 border-b border-border/80 bg-background/40">
        <span className="text-[10px] uppercase tracking-widest text-muted-foreground">
          {language === ('text' as Language) ? 'plain text' : language}
        </span>
      </div>
      <div className="overflow-auto max-h-[70vh]">
        <Highlight code={source} language={language} theme={themes.vsDark}>
          {({ className, style, tokens, getLineProps, getTokenProps }) => (
            <pre
              className={`${className} text-[12px] leading-[1.55] font-mono m-0 py-3`}
              style={{ ...style, background: 'transparent' }}
            >
              {tokens.map((line, i) => {
                const lineProps = getLineProps({ line });
                return (
                  <div key={i} {...lineProps} className={`${lineProps.className || ''} flex`}>
                    <span
                      className="select-none text-right pr-4 pl-3 text-muted-foreground/50 tabular-nums shrink-0"
                      style={{ minWidth: '3.25rem' }}
                    >
                      {i + 1}
                    </span>
                    <span className="flex-1 whitespace-pre">
                      {line.map((token, key) => (
                        <span key={key} {...getTokenProps({ token })} />
                      ))}
                    </span>
                  </div>
                );
              })}
            </pre>
          )}
        </Highlight>
      </div>
    </div>
  );
}

// Build an `expanded` map that opens every ancestor directory of `path`
// so the tree row for that path is visible without manual clicking.
function expandAncestors(path: string, base: Record<string, boolean>): Record<string, boolean> {
  const next: Record<string, boolean> = { ...base, '.': true };
  if (!path || path === '.') return next;
  const segs = path.split('/');
  let acc = '';
  for (let i = 0; i < segs.length - 1; i++) {
    acc = acc ? `${acc}/${segs[i]}` : segs[i];
    next[acc] = true;
  }
  return next;
}

function formatTime(iso: string): string {
  try {
    const d = new Date(iso);
    return d.toLocaleString(undefined, {
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  } catch {
    return iso;
  }
}
