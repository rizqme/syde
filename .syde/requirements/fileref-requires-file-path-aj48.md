---
id: REQ-0041
kind: requirement
name: FileRef Requires File Path
slug: fileref-requires-file-path-aj48
relationships:
    - target: fileref-7vac
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:38:00Z"
statement: The storage layer shall require a non-empty relative file path on every FileRef instance.
req_type: constraint
priority: must
verification: unit test rejecting FileRef with empty file field
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Without a file path the FileRef cannot point back to any source.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:38:00Z"
---
