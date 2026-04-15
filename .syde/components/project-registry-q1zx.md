---
id: COM-0015
kind: component
name: Project Registry
slug: project-registry-q1zx
description: Daemon-side project registry and per-project Store handle cache.
purpose: Track which syde projects the daemon currently knows about
notes:
    - Added GetStore(sydeDir) cache keyed by .syde dir with sync.Mutex. Opens lazy
    - ' holds for daemon lifetime'
    - ' callers must not Close. Stops CLI from opening the same DB — that''s the known tradeoff.'
files:
    - internal/dashboard/registry.go
relationships:
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:58:06Z"
responsibility: Open, cache, and serve Store handles per project path
capabilities:
    - Register new project paths from CLI invocations
    - Cache open BadgerDB handles to avoid repeated opens
    - Provide thread-safe lookups for API handlers
boundaries: Does NOT persist registry state across daemon restarts. Does NOT scan the filesystem for projects.
---
