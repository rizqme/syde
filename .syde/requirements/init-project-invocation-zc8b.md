---
id: REQ-0276
kind: requirement
name: Init Project Invocation
slug: init-project-invocation-zc8b
relationships:
    - target: init-project-xm6c
      type: refines
    - target: config-loader-bx7x
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:03Z"
statement: When the user runs syde init, the syde CLI shall create the .syde directory structure and write a default syde.yaml configuration file.
req_type: interface
priority: must
verification: integration test invoking syde init in an empty directory
source: manual
source_ref: contract:init-project-xm6c
requirement_status: active
rationale: Project initialization is a prerequisite for every other syde workflow.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:03Z"
    config-loader-bx7x:
        hash: 981e7e6e050ed123ce3540f4c43142510c90cbd61fa77ad42f51505e6eda8cea
        at: "2026-04-18T09:38:03Z"
---
