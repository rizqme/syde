---
id: REQ-0277
kind: requirement
name: Init Project Install Skill Flag
slug: init-project-install-skill-flag-j272
relationships:
    - target: init-project-xm6c
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: Where --install-skill is passed to syde init, the syde CLI shall also install the Claude Code skill files after initializing the project.
req_type: interface
priority: must
verification: integration test invoking syde init --install-skill
source: manual
source_ref: contract:init-project-xm6c
requirement_status: active
rationale: Combining init and skill install reduces setup friction for new projects.
---
