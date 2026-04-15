---
boundaries: Read-only — no entity editing. Does NOT talk to BadgerDB directly (always via HTTP API).
capabilities:
    - Overview page with entity counts and relationships
    - Kind-scoped entity list and detail views
    - Plan and task board views with status tracking
    - Learning feed and file tree explorer
    - Command palette search over all entities
description: React 18 + TypeScript + Vite + Tailwind v4 dashboard UI embedded in syded.
files:
    - web/src/main.tsx
    - web/src/App.tsx
    - web/src/index.css
    - web/src/components/EntityDetail.tsx
    - web/src/components/EntityList.tsx
    - web/src/components/KindBadge.tsx
    - web/src/components/RelationshipChip.tsx
    - web/src/components/SearchPalette.tsx
    - web/src/components/Sidebar.tsx
    - web/src/components/icons.tsx
    - web/src/components/EntityEmptyState.tsx
    - web/src/components/EntityFilterBar.tsx
    - web/src/hooks/useApi.ts
    - web/src/hooks/useWebSocket.ts
    - web/src/lib/api.ts
    - web/src/pages/Overview.tsx
    - web/src/pages/FileTree.tsx
    - web/src/pages/LearningFeed.tsx
    - web/src/pages/PlanView.tsx
    - web/src/pages/TaskBoard.tsx
    - web/src/pages/Graph.tsx
    - web/src/pages/ERD.tsx
    - web/embed.go
    - web/index.html
    - web/package.json
    - web/vite.config.ts
    - web/tsconfig.json
    - web/tsconfig.app.json
    - web/tsconfig.node.json
    - web/eslint.config.js
    - scripts/wireframe-shot.sh
id: COM-0018
kind: component
name: Web SPA
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
purpose: Provide the browser-facing UI for syde projects
relationships:
    - target: syded-dashboard
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
slug: web-spa-jy9z
updated_at: "2026-04-15T03:08:45Z"
---
