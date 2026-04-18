---
id: REQ-0162
kind: requirement
name: Scan Helpers Skip Hidden And Ignored Files
slug: scan-helpers-skip-hidden-and-ignored-files-kfhl
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
updated_at: "2026-04-18T09:38:04Z"
statement: While walking the project tree, the scan helpers shall skip hidden files and known skip-list directories.
req_type: functional
priority: must
verification: inspection of walk filters in internal/scan/guide.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Hidden and vendor files distort the language histogram if not filtered.
verified_against:
    scan-helpers-legacy-sa6d:
        hash: c9b19fd18480f13b89908a38caa95ff8ea66c15c73dc6cf6e551e414e71ce3ba
        at: "2026-04-18T09:38:04Z"
---
