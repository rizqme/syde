---
id: REQ-0014
kind: requirement
name: syded shall be a read-only browser
slug: syded-shall-be-a-read-only-browser-0oq0
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:30Z"
statement: The syded daemon shall not expose any HTTP endpoint that mutates entity files.
req_type: constraint
priority: must
verification: audit walking dashboard write handlers
source: manual
source_ref: system:syded-dashboard-e82c:scope
requirement_status: active
rationale: 'Read-only contract simplifies concurrency: only the CLI writes. Mutations would race the file watcher and confuse clients.'
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:30Z"
---
