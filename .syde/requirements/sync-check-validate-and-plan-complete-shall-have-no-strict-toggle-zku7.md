---
id: REQ-0405
kind: requirement
name: Sync check validate and plan complete shall have no strict toggle
slug: sync-check-validate-and-plan-complete-shall-have-no-strict-toggle-zku7
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:33Z"
statement: The syde sync-check, validate, and plan-complete commands shall not accept a --strict flag and shall always block on any audit finding.
req_type: constraint
priority: must
verification: syde sync check --strict exits non-zero with unknown-flag error; syde sync check without flag exits non-zero on any finding
source: plan
requirement_status: active
rationale: Removing the toggle removes the option to run non-strict.
---
