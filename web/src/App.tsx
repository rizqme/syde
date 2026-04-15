import { useState, useEffect, useCallback } from 'react';
import { useNavigate, useParams, useLocation } from 'react-router-dom';
import { Sidebar } from './components/Sidebar';
import { EntityList } from './components/EntityList';
import { EntityDetail } from './components/EntityDetail';
import { SearchPalette } from './components/SearchPalette';
import { Overview } from './pages/Overview';
import { PlanView } from './pages/PlanView';
import { TaskBoard } from './pages/TaskBoard';
import { FileTree } from './pages/FileTree';
import { Graph } from './pages/Graph';
import { ERD } from './pages/ERD';
import { EntityEmptyState } from './components/EntityEmptyState';
import { useApi } from './hooks/useApi';
import { useWebSocket } from './hooks/useWebSocket';
import { api, fetchProjects, setProject, getProject, StatusResponse, EntitiesResponse } from './lib/api';

const SPECIAL_VIEWS = ['plan', 'task', '__tree__', '__graph__'];

export default function App() {
  const [projectName, setProjectName] = useState('Loading...');
  const [projectSlug, setProjectSlug] = useState<string | null>(null);
  const [ready, setReady] = useState(false);
  const [selectedSlug, setSelectedSlug] = useState<string | null>(null);
  const [searchOpen, setSearchOpen] = useState(false);
  const [reloadKey, setReloadKey] = useState(0);
  // Concept page has two view modes — the default 2-column inbox and
  // an ERD canvas. Toggle lives on App so it survives entity-select
  // state changes; reset to 'list' whenever we navigate away from the
  // concept kind so returning to it starts fresh.
  const [conceptView, setConceptView] = useState<'list' | 'erd'>('list');

  const location = useLocation();
  const navigate = useNavigate();

  // Parse active kind and selected entity from URL
  // URL format: /<project>/<kind>/<entity-slug>
  const pathParts = location.pathname.split('/').filter(Boolean);
  const activeKind = pathParts.length >= 2 ? pathParts[1] : null;
  const urlEntitySlug = pathParts.length >= 3 ? pathParts[2] : null;

  // Sync URL entity slug to state
  useEffect(() => {
    if (urlEntitySlug && urlEntitySlug !== selectedSlug) {
      setSelectedSlug(urlEntitySlug);
    }
  }, [urlEntitySlug]);

  // Reset the concept view-mode toggle whenever the active kind leaves
  // 'concept' — otherwise returning to concepts would still be in ERD
  // mode from a previous session, which breaks the default-to-list UX.
  useEffect(() => {
    if (activeKind !== 'concept') {
      setConceptView('list');
    }
  }, [activeKind]);

  // WebSocket live reload
  useWebSocket(ready ? getProject() : null, useCallback(() => {
    setReloadKey((k) => k + 1);
  }, []));

  // Init: find project
  useEffect(() => {
    fetchProjects().then((data) => {
      if (!data.projects?.length) return;
      const firstPart = pathParts[0] || '';
      const proj = data.projects.find((p) => p.slug === firstPart) || data.projects[0];
      setProject(proj.slug);
      setProjectSlug(proj.slug);
      setProjectName(proj.name || proj.slug);
      setReady(true);
      // If URL doesn't start with project slug, redirect
      if (firstPart !== proj.slug) {
        navigate(`/${proj.slug}`, { replace: true });
      }
    });
  }, []);

  // Keyboard shortcuts
  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
        e.preventDefault();
        setSearchOpen(true);
      }
      if (e.key === 'Escape') {
        if (searchOpen) setSearchOpen(false);
        else if (selectedSlug) {
          setSelectedSlug(null);
          if (projectSlug && activeKind) navigate(`/${projectSlug}/${activeKind}`);
        }
      }
    };
    window.addEventListener('keydown', handler);
    return () => window.removeEventListener('keydown', handler);
  }, [searchOpen, selectedSlug, projectSlug, activeKind, navigate]);

  const { data: status } = useApi<StatusResponse>(
    () => (ready ? api.status() : Promise.resolve({ counts: {}, total: 0 })),
    [ready, reloadKey]
  );

  const { data: entitiesData } = useApi<EntitiesResponse>(
    () => (ready && activeKind && !SPECIAL_VIEWS.includes(activeKind) ? api.entities(activeKind) : Promise.resolve({ entities: [], count: 0 })),
    [ready, activeKind, reloadKey]
  );

  const handleNavigate = useCallback((kind: string | null) => {
    setSelectedSlug(null);
    if (kind === null) {
      navigate(`/${projectSlug}`);
    } else {
      navigate(`/${projectSlug}/${kind}`);
    }
  }, [projectSlug, navigate]);

  const handleSelectEntity = useCallback((slug: string) => {
    setSelectedSlug(slug);
    if (projectSlug && activeKind) {
      navigate(`/${projectSlug}/${activeKind}/${slug}`);
    }
  }, [projectSlug, activeKind, navigate]);

  // Navigate to an entity in a possibly-different kind. Used by relationship
  // chips so clicking e.g. a 'system' link from a component page actually
  // switches to the systems list with that entity selected. The optional
  // filter param pre-seeds the destination's filter bar via ?filter=… so a
  // child clicked from a parent system lands on a list scoped to siblings.
  const handleNavigateEntity = useCallback((slug: string, kind?: string, filter?: string) => {
    if (!projectSlug) return;
    const targetKind = kind || activeKind;
    if (!targetKind) return;
    setSelectedSlug(slug);
    const qs = filter ? `?filter=${encodeURIComponent(filter)}` : '';
    navigate(`/${projectSlug}/${targetKind}/${slug}${qs}`);
  }, [projectSlug, activeKind, navigate]);

  // Open a source file in the FileTree page. Path goes via ?path= query
  // param so any '/' inside survives without escaping the URL grammar.
  const handleOpenFile = useCallback((path: string) => {
    if (!projectSlug) return;
    navigate(`/${projectSlug}/__tree__?path=${encodeURIComponent(path)}`);
  }, [projectSlug, navigate]);

  if (!ready) {
    return (
      <div className="flex items-center justify-center h-screen text-muted-foreground text-sm">
        Loading...
      </div>
    );
  }

  return (
    <>
      <Sidebar
        projectName={projectName}
        status={status}
        activeKind={activeKind}
        onNavigate={handleNavigate}
        onSearch={() => setSearchOpen(true)}
      />

      <main className="flex-1 min-w-0 overflow-hidden flex">
        {activeKind === null && (
          <div className="flex-1 overflow-y-auto px-8 py-8 max-w-4xl">
            <Overview onNavigate={handleNavigate} />
          </div>
        )}
        {activeKind === 'plan' && (
          <div className="flex-1 overflow-y-auto px-8 py-8 max-w-4xl">
            <PlanView onSelectEntity={handleSelectEntity} />
          </div>
        )}
        {activeKind === 'task' && (
          <div className="flex-1 overflow-y-auto px-8 py-8 max-w-4xl">
            <TaskBoard onSelectEntity={handleSelectEntity} />
          </div>
        )}
        {activeKind === '__tree__' && ready && (
          <div className="flex-1 min-w-0 min-h-0">
            <FileTree key={reloadKey} />
          </div>
        )}
        {activeKind === '__graph__' && ready && (
          <div className="flex-1 min-w-0 min-h-0">
            <Graph onSelectEntity={handleNavigateEntity} />
          </div>
        )}
        {activeKind && !SPECIAL_VIEWS.includes(activeKind) && entitiesData && (
          <div className="flex-1 flex min-w-0 min-h-0 relative">
            {/* Concept page renders a floating List/ERD toggle at
                the top-right of the main area, replacing the X close
                button in the detail panel. Other kinds render nothing
                here. */}
            {activeKind === 'concept' && (
              <div className="absolute top-3 right-4 z-10">
                <div className="inline-flex rounded-md border border-border bg-card overflow-hidden text-xs shadow-sm">
                  <button
                    onClick={() => setConceptView('list')}
                    className={`px-3 py-1 ${conceptView === 'list' ? 'bg-muted text-foreground' : 'text-muted-foreground hover:text-foreground'}`}
                  >
                    List
                  </button>
                  <button
                    onClick={() => setConceptView('erd')}
                    className={`px-3 py-1 border-l border-border ${conceptView === 'erd' ? 'bg-muted text-foreground' : 'text-muted-foreground hover:text-foreground'}`}
                  >
                    ERD
                  </button>
                </div>
              </div>
            )}
            {activeKind === 'concept' && conceptView === 'erd' ? (
              <div className="flex-1 min-w-0 min-h-0">
                <ERD
                  onSelectEntity={(slug, kind) => {
                    // Clicking a concept node returns to list mode so
                    // the user immediately sees the detail panel for
                    // the selected entity.
                    setConceptView('list');
                    handleNavigateEntity(slug, kind);
                  }}
                />
              </div>
            ) : (
              <>
                {/* Inbox-style 2-column: list on the left, detail on the right. */}
                <div className="w-[420px] shrink-0 border-r border-border overflow-y-auto px-6 py-6">
                  <EntityList
                    key={activeKind}
                    entities={entitiesData.entities || []}
                    kind={activeKind}
                    onSelect={handleSelectEntity}
                    selectedSlug={selectedSlug || undefined}
                  />
                </div>
                <div className="flex-1 min-w-0 overflow-y-auto">
                  {selectedSlug ? (
                    <EntityDetail
                      slug={selectedSlug}
                      onNavigate={handleNavigateEntity}
                      onOpenFile={handleOpenFile}
                      onClose={() => {
                        setSelectedSlug(null);
                        if (projectSlug && activeKind) navigate(`/${projectSlug}/${activeKind}`);
                      }}
                      inline
                      hideClose={activeKind === 'concept'}
                    />
                  ) : (
                    <EntityEmptyState kind={activeKind} />
                  )}
                </div>
              </>
            )}
          </div>
        )}
      </main>

      {/* Floating detail panel only used by special views (plan / task) */}
      {selectedSlug && activeKind && SPECIAL_VIEWS.includes(activeKind) && (
        <EntityDetail
          slug={selectedSlug}
          onNavigate={handleNavigateEntity}
          onOpenFile={handleOpenFile}
          onClose={() => {
            setSelectedSlug(null);
            if (projectSlug && activeKind) navigate(`/${projectSlug}/${activeKind}`);
          }}
        />
      )}

      <SearchPalette
        open={searchOpen}
        onClose={() => setSearchOpen(false)}
        onSelect={(slug, kind) => {
          handleNavigateEntity(slug, kind);
          setSearchOpen(false);
        }}
      />
    </>
  );
}
