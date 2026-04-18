---
id: REQ-0088
kind: requirement
name: SPA Command Palette Search
slug: spa-command-palette-search-tzc3
relationships:
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:35Z"
statement: When the user opens the command palette, the web SPA shall search across all entities in the active project and render matching results.
req_type: functional
priority: must
verification: manual inspection of the command palette against a known query
source: manual
source_ref: component:web-spa-jy9z
requirement_status: active
rationale: Keyboard-driven search is the fastest way to jump between entities.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:35Z"
---
