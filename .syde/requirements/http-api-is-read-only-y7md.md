---
id: REQ-0048
kind: requirement
name: HTTP API Is Read Only
slug: http-api-is-read-only-y7md
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: The syded HTTP API shall not mutate entities in response to dashboard requests.
req_type: constraint
priority: must
verification: inspection of internal/dashboard/api.go handler table for mutating verbs
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The dashboard is a read-only reviewer; writes flow through the CLI.
---
