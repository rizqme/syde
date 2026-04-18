---
id: REQ-0079
kind: requirement
name: CLI Commands Support Init And Skill Install
slug: cli-commands-support-init-and-skill-install-nkf6
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:52Z"
statement: The syde CLI shall provide init and install-skill subcommands that bootstrap new projects and write agent skill files.
req_type: functional
priority: must
verification: integration test invoking syde init and syde install-skill
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Bootstrap UX is the first contact point for new users.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:52Z"
---
