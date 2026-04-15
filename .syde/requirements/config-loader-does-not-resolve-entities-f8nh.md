---
id: REQ-0124
kind: requirement
name: Config Loader Does Not Resolve Entities
slug: config-loader-does-not-resolve-entities-f8nh
relationships:
    - target: config-loader-bx7x
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:46Z"
statement: The config loader shall not resolve entities and shall delegate entity lookup to the storage and query engines.
req_type: constraint
priority: must
verification: code review of internal/config imports
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Keeping config loading decoupled from the entity store avoids cyclic initialization.
---
