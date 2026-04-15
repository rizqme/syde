---
id: REQ-0117
kind: requirement
name: Config Loader Loads And Saves syde yaml
slug: config-loader-loads-and-saves-syde-yaml-gz93
relationships:
    - target: config-loader-bx7x
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:35Z"
statement: The config loader shall load and save the per-project syde.yaml file as a typed Config struct.
req_type: functional
priority: must
verification: inspection of Load/Save in internal/config/config.go
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: syde.yaml is the canonical source for project metadata.
---
