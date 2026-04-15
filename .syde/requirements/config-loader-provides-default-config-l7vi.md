---
id: REQ-0119
kind: requirement
name: Config Loader Provides Default Config
slug: config-loader-provides-default-config-l7vi
relationships:
    - target: config-loader-bx7x
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:38Z"
statement: When a project has no syde.yaml, the config loader shall return a default Config with project name, version, and default tree ignore patterns.
req_type: functional
priority: must
verification: inspection of default config in config.go
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Fresh projects must be usable without manual config editing.
---
