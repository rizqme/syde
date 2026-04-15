---
id: REQ-0123
kind: requirement
name: Config Loader Finds syde Dir From Subfolder
slug: config-loader-finds-syde-dir-from-subfolder-m0y6
relationships:
    - target: config-loader-bx7x
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:45Z"
statement: When invoked from a subdirectory of a syde project, the config loader shall walk upward to locate the nearest .syde directory.
req_type: functional
priority: must
verification: inspection of FindSydeDir in config.go
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Users run syde commands from any folder in their repo.
---
