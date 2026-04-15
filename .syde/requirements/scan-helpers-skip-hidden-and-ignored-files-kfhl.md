---
id: REQ-0162
kind: requirement
name: Scan Helpers Skip Hidden And Ignored Files
slug: scan-helpers-skip-hidden-and-ignored-files-kfhl
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:25Z"
statement: While walking the project tree, the scan helpers shall skip hidden files and known skip-list directories.
req_type: functional
priority: must
verification: inspection of walk filters in internal/scan/guide.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Hidden and vendor files distort the language histogram if not filtered.
---
