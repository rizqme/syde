---
id: REQ-0081
kind: requirement
name: CLI Commands Does Not Persist Entities Directly
slug: cli-commands-does-not-persist-entities-directly-kn77
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:59Z"
statement: The syde CLI shall not persist entities directly and shall delegate all storage operations to the storage engine via the HTTP client.
req_type: constraint
priority: must
verification: code review of internal/cli for storage imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Centralizing writes in syded ensures there is a single BadgerDB writer.
---
