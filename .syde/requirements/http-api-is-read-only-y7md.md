---
id: REQ-0048
kind: requirement
name: HTTP API Is Read Only
slug: http-api-is-read-only-y7md
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:36:57Z"
statement: The syded HTTP API shall not mutate entities in response to dashboard requests.
req_type: constraint
priority: must
verification: inspection of internal/dashboard/api.go handler table for mutating verbs
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The dashboard is a read-only reviewer; writes flow through the CLI.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:57Z"
---
