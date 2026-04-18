---
boundaries: Largely superseded by Summary Tree. Does NOT track change state.
capabilities:
- Directory walk with hidden/skiplist filtering
- File-to-component coverage mapping
description: Legacy shallow ScanGuide for 'syde scan' and sync coverage (superseded by Summary Tree).
files:
- internal/scan/guide.go
- internal/scan/coverage.go
id: COM-0011
kind: component
name: Scan Helpers (legacy)
purpose: Produce a shallow ScanGuide for the legacy 'syde scan' and 'sync --coverage' commands
relationships:
- type: belongs_to
  target: syde-5tdt
responsibility: Walk the source tree and emit file counts, language breakdown, key files per dir
slug: scan-helpers-legacy-sa6d
updated_at: '2026-04-14T03:35:54Z'
---
