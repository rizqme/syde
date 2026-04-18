---
id: REQ-0405
kind: requirement
name: Sync check validate and plan complete shall have no strict toggle
slug: sync-check-validate-and-plan-complete-shall-have-no-strict-toggle-zku7
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:38:00Z"
statement: The syde sync-check, validate, and plan-complete commands shall not accept a --strict flag and shall always block on any audit finding.
req_type: constraint
priority: must
verification: syde sync check --strict exits non-zero with unknown-flag error; syde sync check without flag exits non-zero on any finding
source: plan
requirement_status: active
rationale: Removing the toggle removes the option to run non-strict.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:00Z"
---
