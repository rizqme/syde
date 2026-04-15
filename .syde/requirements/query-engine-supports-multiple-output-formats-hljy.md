---
id: REQ-0158
kind: requirement
name: Query Engine Supports Multiple Output Formats
slug: query-engine-supports-multiple-output-formats-hljy
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:12Z"
statement: The query engine shall provide rich, JSON, compact, and refs output formatters for resolved entities and search hits.
req_type: functional
priority: must
verification: inspection of internal/query/formatter.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: CLI and HTTP API consumers need format flexibility.
---
