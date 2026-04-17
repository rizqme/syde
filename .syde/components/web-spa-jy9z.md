---
id: COM-0018
kind: component
name: Web SPA
slug: web-spa-jy9z
description: React 18 + TypeScript + Vite + Tailwind v4 dashboard UI embedded in syded.
purpose: Provide the browser-facing UI for syde projects
notes:
    - App.tsx gates FileTree render on 'ready' state so project slug is set before api.tree() fires — fixes 'No tree available' flash on page reload (2026-04-13).
    - 'FileTree.tsx: moved useMemo ahead of early returns — fix for React error #310 (hook count mismatch across renders) (2026-04-13).'
    - Entity views are now inbox-style 2-column (420px list + flex-1 detail). EntityDetail.onNavigate accepts an optional target kind so relationship chips switch pages when crossing kinds (2026-04-14).
    - 'Sidebar: File Tree promoted to a top-level item under Overview'
    - ' Codebase group removed; all sidebar icons replaced with hand-rolled flat outline SVGs (web/src/components/icons.tsx'
    - ' lucide-style'
    - ' currentColor) (2026-04-14).'
    - Component cards now show description (always present
    - ' validator-enforced) (2026-04-14).'
files:
    - scripts/wireframe-shot.sh
    - web/embed.go
    - web/eslint.config.js
    - web/index.html
    - web/package.json
    - web/src/App.tsx
    - web/src/main.tsx
    - web/src/index.css
    - web/src/components/EntityDetail.tsx
    - web/src/components/EntityEmptyState.tsx
    - web/src/components/EntityFilterBar.tsx
    - web/src/components/EntityList.tsx
    - web/src/components/ExtendedFieldDiff.tsx
    - web/src/components/FlowChart.tsx
    - web/src/components/KindBadge.tsx
    - web/src/components/NewContractDraftView.tsx
    - web/src/components/NewEntityDraftView.tsx
    - web/src/components/PhaseTaskList.tsx
    - web/src/components/PlanChangesView.tsx
    - web/src/components/PlanDetailPanel.tsx
    - web/src/components/RelationshipChip.tsx
    - web/src/components/SearchPalette.tsx
    - web/src/components/Sidebar.tsx
    - web/src/components/icons.tsx
    - web/src/hooks/useApi.ts
    - web/src/hooks/useWebSocket.ts
    - web/src/lib/api.ts
    - web/tsconfig.app.json
    - web/tsconfig.json
    - web/tsconfig.node.json
    - web/vite.config.ts
    - web/src/pages/Overview.tsx
    - web/src/pages/FileTree.tsx
    - web/src/pages/Graph.tsx
relationships:
    - target: syded-dashboard
      type: belongs_to
updated_at: "2026-04-17T09:05:21Z"
responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
capabilities:
    - Overview page with entity counts and relationships
    - Kind-scoped entity list and detail views
    - Plan and task board views with status tracking
    - File tree explorer
    - Command palette search over all entities
boundaries: Read-only — no entity editing. Does NOT talk to BadgerDB directly (always via HTTP API).
---
