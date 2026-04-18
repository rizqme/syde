---
id: REQ-0081
kind: requirement
name: CLI Commands Does Not Persist Entities Directly
slug: cli-commands-does-not-persist-entities-directly-kn77
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:22Z"
statement: The syde CLI shall not persist entities directly and shall delegate all storage operations to the storage engine via the HTTP client.
req_type: constraint
priority: must
verification: code review of internal/cli for storage imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Centralizing writes in syded ensures there is a single BadgerDB writer.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:22Z"
---
