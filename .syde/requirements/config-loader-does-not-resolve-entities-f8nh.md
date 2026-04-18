---
id: REQ-0124
kind: requirement
name: Config Loader Does Not Resolve Entities
slug: config-loader-does-not-resolve-entities-f8nh
relationships:
    - target: config-loader-bx7x
      type: refines
updated_at: "2026-04-18T09:37:24Z"
statement: The config loader shall not resolve entities and shall delegate entity lookup to the storage and query engines.
req_type: constraint
priority: must
verification: code review of internal/config imports
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Keeping config loading decoupled from the entity store avoids cyclic initialization.
verified_against:
    config-loader-bx7x:
        hash: 981e7e6e050ed123ce3540f4c43142510c90cbd61fa77ad42f51505e6eda8cea
        at: "2026-04-18T09:37:24Z"
---
