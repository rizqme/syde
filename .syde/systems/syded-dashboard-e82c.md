---
context: Launched via 'syde open' or run standalone. Opens a project registry, serves the embedded SPA, and watches markdown files for live updates via WebSocket.
description: HTTP daemon binary 'syded' that serves the React SPA + REST/WebSocket API for browsing syde projects in a browser.
id: SYS-0005
kind: system
name: syded Dashboard
notes:
    - 'Dashboard caches Store handles per project (registry.GetStore) to fix concurrent-request 500s from BadgerDB directory-lock contention. KNOWN CAVEAT: while syded is running'
    - ' the syde CLI cannot open .syde/ in the same project — stop syded before running CLI commands. (2026-04-13)'
purpose: Provide a visual, clickable interface for humans reviewing architectures that are primarily edited via CLI.
relationships:
    - target: syde
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
scope: 'In: HTTP API, WebSocket live-refresh, embedded React SPA, multi-project registry. Out: Entity editing (read-only browser for now).'
slug: syded-dashboard-e82c
updated_at: "2026-04-15T03:08:45Z"
---
