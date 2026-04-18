---
id: REQ-0037
kind: requirement
name: FileRef Line Ranges Must Be Valid
slug: fileref-line-ranges-must-be-valid-h5be
relationships:
    - target: fileref-7vac
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:12Z"
statement: The storage layer shall ensure that every FileRef line range lies within the current byte span of its markdown file.
req_type: constraint
priority: must
verification: integration test asserting each FileRef line span is in bounds after reindex
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Out-of-range line spans would surface invalid slices to consumers of syde get.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:12Z"
---
