---
id: SYS-0005
kind: system
name: syded Dashboard
slug: syded-dashboard-e82c
description: HTTP daemon binary 'syded' that serves the React SPA + REST/WebSocket API for browsing syde projects in a browser.
purpose: Provide a visual, clickable interface for humans reviewing architectures that are primarily edited via CLI.
notes:
    - 'Dashboard caches Store handles per project (registry.GetStore) to fix concurrent-request 500s from BadgerDB directory-lock contention. KNOWN CAVEAT: while syded is running'
    - ' the syde CLI cannot open .syde/ in the same project — stop syded before running CLI commands. (2026-04-13)'
relationships:
    - target: syde
      type: belongs_to
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
updated_at: "2026-04-15T11:05:37Z"
context: Launched via 'syde open' or run standalone. Opens a project registry, serves the embedded SPA, and watches markdown files for live updates via WebSocket.
scope: 'In: HTTP API, WebSocket live-refresh, embedded React SPA, multi-project registry. Out: Entity editing (read-only browser for now).'
---
