---
id: REQ-0126
kind: requirement
name: Config Loader Does Not Touch BadgerDB
slug: config-loader-does-not-touch-badgerdb-a7qh
relationships:
    - target: config-loader-bx7x
      type: refines
updated_at: "2026-04-18T09:37:01Z"
statement: The config loader shall not open or query BadgerDB and shall operate only on YAML files on disk.
req_type: constraint
priority: must
verification: code review of internal/config for BadgerDB imports
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Config IO must be possible even when the index is absent or corrupt.
verified_against:
    config-loader-bx7x:
        hash: 981e7e6e050ed123ce3540f4c43142510c90cbd61fa77ad42f51505e6eda8cea
        at: "2026-04-18T09:37:01Z"
---
