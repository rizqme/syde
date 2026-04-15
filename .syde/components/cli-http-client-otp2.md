---
id: COM-0020
kind: component
name: CLI HTTP Client
slug: cli-http-client-otp2
description: Thin HTTP client the syde CLI uses to talk to a local syded — the sole BadgerDB writer
purpose: Give the CLI a zero-BadgerDB path to read + write the syde model by routing every operation through syded's /api/<project>/* HTTP surface
notes:
    - 'Codex compatibility follow-up: client slug derivation now honors syde.yaml project names so CLI routes match syded dashboard registry slugs.'
files:
    - internal/client/client.go
relationships:
    - target: syde-cli
      type: belongs_to
    - target: daemon-launcher
      type: depends_on
    - target: http-api
      type: depends_on
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
updated_at: "2026-04-15T09:27:08Z"
responsibility: Wrap every syded read + write endpoint as a typed Go method and transparently auto-launch syded via daemon.EnsureRunning on first call
capabilities:
    - Derive the project slug from a .syde/ dir so CLI + syded agree on the project key
    - Wrap read endpoints (Status, List, Get, Query, Validate, SyncCheck, Context, Constraints, Search, FilesOrphans, FilesCoverage)
    - Wrap write endpoints (CreateEntity, UpdateEntity, DeleteEntity, Reindex) with YAML-frontmatter payloads
    - Auto-launch syded on first request via internal/daemon.EnsureRunning
boundaries: Does NOT talk to BadgerDB directly. Does NOT cache results. Does NOT render output — callers do their own formatting.
---
