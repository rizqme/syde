---
id: REQ-0119
kind: requirement
name: Config Loader Provides Default Config
slug: config-loader-provides-default-config-l7vi
relationships:
    - target: config-loader-bx7x
      type: refines
updated_at: "2026-04-18T09:37:27Z"
statement: When a project has no syde.yaml, the config loader shall return a default Config with project name, version, and default tree ignore patterns.
req_type: functional
priority: must
verification: inspection of default config in config.go
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: Fresh projects must be usable without manual config editing.
verified_against:
    config-loader-bx7x:
        hash: 981e7e6e050ed123ce3540f4c43142510c90cbd61fa77ad42f51505e6eda8cea
        at: "2026-04-18T09:37:27Z"
---
