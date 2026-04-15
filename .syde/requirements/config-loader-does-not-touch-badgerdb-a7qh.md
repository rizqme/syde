---
id: REQ-0126
kind: requirement
name: Config Loader Does Not Touch BadgerDB
slug: config-loader-does-not-touch-badgerdb-a7qh
relationships:
    - target: config-loader-bx7x
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:47Z"
statement: The config loader shall not open or query BadgerDB and shall operate only on YAML files on disk.
req_type: constraint
priority: must
verification: code review of internal/config for BadgerDB imports
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Config IO must be possible even when the index is absent or corrupt.
---
