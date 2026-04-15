---
id: REQ-0156
kind: requirement
name: Query Engine Searches Component Owned Code
slug: query-engine-searches-component-owned-code-600w
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:07Z"
statement: When SearchCode is called with a term, the query engine shall search across files owned by components using ripgrep when available and a Go fallback otherwise.
req_type: functional
priority: must
verification: unit test of SearchCode in engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Code search is scoped to component-owned paths so results stay relevant.
---
