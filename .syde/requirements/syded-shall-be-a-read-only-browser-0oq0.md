---
id: REQ-0014
kind: requirement
name: syded shall be a read-only browser
slug: syded-shall-be-a-read-only-browser-0oq0
relationships:
    - target: syded-dashboard-e82c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syded daemon shall not expose any HTTP endpoint that mutates entity files.
req_type: constraint
priority: must
verification: audit walking dashboard write handlers
source: manual
source_ref: system:syded-dashboard-e82c:scope
requirement_status: active
rationale: 'Read-only contract simplifies concurrency: only the CLI writes. Mutations would race the file watcher and confuse clients.'
---
