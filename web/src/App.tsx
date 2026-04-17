import { useState, useEffect, useCallback } from 'react';
import { useNavigate, useParams, useLocation } from 'react-router-dom';
import { Sidebar } from './components/Sidebar';
import { EntityList } from './components/EntityList';
import { EntityDetail } from './components/EntityDetail';
import { PlanDetailPanel } from './components/PlanDetailPanel';
import { SearchPalette } from './components/SearchPalette';
import { Overview } from './pages/Overview';
import { FileTree } from './pages/FileTree';
import { Graph } from './pages/Graph';
import { EntityEmptyState } from './components/EntityEmptyState';
import { useApi } from './hooks/useApi';
import { useWebSocket } from './hooks/useWebSocket';
import { api, fetchProjects, setProject, getProject, StatusResponse, EntitiesResponse } from './lib/api';

const SPECIAL_VIEWS = ['task', '__tree__', '__graph__'];

export default function App() {
  const [projectName, setProjectName] = useState('Loading...');
  const [projectSlug, setProjectSlug] = useState<string | null>(null);
  const [ready, setReady] = useState(false);
  const [selectedSlug, setSelectedSlug] = useState<string | null>(null);
  const [searchOpen, setSearchOpen] = useState(false);
  const [reloadKey, setReloadKey] = useState(0);

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

  // The standalone Tasks page was removed — tasks now live inside the
  // plan detail page. If someone lands on a legacy /<proj>/task URL
  // (bookmark, external link) redirect them to the Plans inbox.
  useEffect(() => {
    if (activeKind === 'task' && !urlEntitySlug && projectSlug) {
      navigate(`/${projectSlug}/plan`, { replace: true });
    }
  }, [activeKind, urlEntitySlug, projectSlug, navigate]);

  const handleSocketNavigate = useCallback((path: string) => {
    const current = `${location.pathname}${location.search}`;
    if (path === current) return;
    navigate(path);
  }, [location.pathname, location.search, navigate]);

  // WebSocket live reload and dashboard-driven navigation
  useWebSocket(ready ? getProject() : null, useCallback(() => {
    setReloadKey((k) => k + 1);
  }, []), handleSocketNavigate);

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
                activeKind === 'plan' ? (
                  <PlanDetailPanel
                    slug={selectedSlug}
                    onNavigate={handleNavigateEntity}
                    onOpenFile={handleOpenFile}
                    onClose={() => {
                      setSelectedSlug(null);
                      if (projectSlug) navigate(`/${projectSlug}/plan`);
                    }}
                  />
                ) : (
                  <EntityDetail
                    slug={selectedSlug}
                    onNavigate={handleNavigateEntity}
                    onOpenFile={handleOpenFile}
                    onClose={() => {
                      setSelectedSlug(null);
                      if (projectSlug && activeKind) navigate(`/${projectSlug}/${activeKind}`);
                    }}
                    inline
                  />
                )
              ) : (
                <EntityEmptyState kind={activeKind} />
              )}
            </div>
          </div>
        )}
      </main>

      {/* Floating detail panel only used by special views such as task detail. */}
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
