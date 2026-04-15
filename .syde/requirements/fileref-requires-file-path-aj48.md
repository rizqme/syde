---
id: REQ-0041
kind: requirement
name: FileRef Requires File Path
slug: fileref-requires-file-path-aj48
relationships:
    - target: fileref-7vac
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:21Z"
statement: The storage layer shall require a non-empty relative file path on every FileRef instance.
req_type: constraint
priority: must
verification: unit test rejecting FileRef with empty file field
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Without a file path the FileRef cannot point back to any source.
---
