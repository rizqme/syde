---
id: REQ-0276
kind: requirement
name: Init Project Invocation
slug: init-project-invocation-zc8b
relationships:
    - target: init-project-xm6c
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde init, the syde CLI shall create the .syde directory structure and write a default syde.yaml configuration file.
req_type: interface
priority: must
verification: integration test invoking syde init in an empty directory
source: manual
source_ref: contract:init-project-xm6c
requirement_status: active
rationale: Project initialization is a prerequisite for every other syde workflow.
---
