---
id: REQ-0064
kind: requirement
name: Registry Caches Open Badger Handles
slug: registry-caches-open-badger-handles-82op
relationships:
    - target: project-registry-q1zx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: The project registry shall cache open BadgerDB handles keyed by project .syde directory for the daemon lifetime.
req_type: functional
priority: must
verification: integration test asserting handle reuse across requests
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Reopening BadgerDB on every request is both slow and conflicts with single-writer semantics.
---
