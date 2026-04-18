---
id: REQ-0121
kind: requirement
name: Config Loader Maps Component Paths To Slugs
slug: config-loader-maps-component-paths-to-slugs-w2qd
relationships:
    - target: config-loader-bx7x
      type: refines
updated_at: "2026-04-18T09:38:02Z"
statement: The config loader shall expose the component_paths glob to component slug mapping used by the constraints check command.
req_type: functional
priority: must
verification: inspection of Config.ComponentPaths handling
source: manual
source_ref: component:config-loader-bx7x
requirement_status: active
rationale: File-to-component mapping is the basis of constraint enforcement.
verified_against:
    config-loader-bx7x:
        hash: 981e7e6e050ed123ce3540f4c43142510c90cbd61fa77ad42f51505e6eda8cea
        at: "2026-04-18T09:38:02Z"
---
