---
id: REQ-0176
kind: requirement
name: Storage Engine Serializes Entities As Markdown
slug: storage-engine-serializes-entities-as-markdown-4try
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:10Z"
statement: When saving an entity, the storage engine shall serialize it as a markdown file with YAML frontmatter followed by the body.
req_type: functional
priority: must
verification: unit test of serializer in internal/storage/serializer.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Markdown on disk is the source of truth and must be git-friendly.
---
