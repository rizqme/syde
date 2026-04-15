---
boundaries: Does NOT persist registry state across daemon restarts. Does NOT scan the filesystem for projects.
capabilities:
    - Register new project paths from CLI invocations
    - Cache open BadgerDB handles to avoid repeated opens
    - Provide thread-safe lookups for API handlers
description: Daemon-side project registry and per-project Store handle cache.
files:
    - internal/dashboard/registry.go
id: COM-0015
kind: component
name: Project Registry
notes:
    - Added GetStore(sydeDir) cache keyed by .syde dir with sync.Mutex. Opens lazy
    - ' holds for daemon lifetime'
    - ' callers must not Close. Stops CLI from opening the same DB — that''s the known tradeoff.'
purpose: Track which syde projects the daemon currently knows about
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
responsibility: Open, cache, and serve Store handles per project path
slug: project-registry-q1zx
updated_at: "2026-04-14T07:04:05Z"
---
