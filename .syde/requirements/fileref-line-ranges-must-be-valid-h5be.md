---
id: REQ-0037
kind: requirement
name: FileRef Line Ranges Must Be Valid
slug: fileref-line-ranges-must-be-valid-h5be
relationships:
    - target: fileref-7vac
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:17Z"
statement: The storage layer shall ensure that every FileRef line range lies within the current byte span of its markdown file.
req_type: constraint
priority: must
verification: integration test asserting each FileRef line span is in bounds after reindex
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Out-of-range line spans would surface invalid slices to consumers of syde get.
---
