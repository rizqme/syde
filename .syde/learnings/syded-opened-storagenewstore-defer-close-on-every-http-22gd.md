---
category: gotcha
confidence: high
description: 'syded opened storage.NewStore() + defer Close on every HTTP request. BadgerDB holds an exclusive directory lock, so parallel requests (SPA fetches status + tree on page load) race on the lock and return 500 ''Cannot acquire directory lock''. Fix: package-level sync.Mutex-protected Store cache in internal/dashboard/registry.go GetStore, opens lazy on first request, never closed. TRADEOFF: while syded runs, ''syde'' CLI in the same project blocks on the DB lock — stop syded before CLI work, or teach the daemon to release on idle / share via a socket.'
discovered_at: "2026-04-13T14:11:44Z"
entity_refs:
    - http-api
id: LRN-0004
kind: learning
name: 'syded opened storage.NewStore() + defer Close on every HTTP '
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: syded-opened-storagenewstore-defer-close-on-every-http-22gd
source: session-observation
updated_at: "2026-04-13T14:11:44Z"
---
