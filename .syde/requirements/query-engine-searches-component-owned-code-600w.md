---
id: REQ-0156
kind: requirement
name: Query Engine Searches Component Owned Code
slug: query-engine-searches-component-owned-code-600w
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:37:17Z"
statement: When SearchCode is called with a term, the query engine shall search across files owned by components using ripgrep when available and a Go fallback otherwise.
req_type: functional
priority: must
verification: unit test of SearchCode in engine.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Code search is scoped to component-owned paths so results stay relevant.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:37:16Z"
---
