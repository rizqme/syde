---
id: REQ-0158
kind: requirement
name: Query Engine Supports Multiple Output Formats
slug: query-engine-supports-multiple-output-formats-hljy
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:36:58Z"
statement: The query engine shall provide rich, JSON, compact, and refs output formatters for resolved entities and search hits.
req_type: functional
priority: must
verification: inspection of internal/query/formatter.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: CLI and HTTP API consumers need format flexibility.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:36:58Z"
---
