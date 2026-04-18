---
id: REQ-0176
kind: requirement
name: Storage Engine Serializes Entities As Markdown
slug: storage-engine-serializes-entities-as-markdown-4try
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:29Z"
statement: When saving an entity, the storage engine shall serialize it as a markdown file with YAML frontmatter followed by the body.
req_type: functional
priority: must
verification: unit test of serializer in internal/storage/serializer.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Markdown on disk is the source of truth and must be git-friendly.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:29Z"
---
