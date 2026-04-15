---
id: REQ-0079
kind: requirement
name: CLI Commands Support Init And Skill Install
slug: cli-commands-support-init-and-skill-install-nkf6
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:56Z"
statement: The syde CLI shall provide init and install-skill subcommands that bootstrap new projects and write agent skill files.
req_type: functional
priority: must
verification: integration test invoking syde init and syde install-skill
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Bootstrap UX is the first contact point for new users.
---
