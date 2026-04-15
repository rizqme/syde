---
id: REQ-0096
kind: requirement
name: Skill Installer Is Idempotent
slug: skill-installer-is-idempotent-1u2y
relationships:
    - target: skill-7fmf
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:10Z"
statement: When the user runs syde install-skill repeatedly, the syde CLI shall produce the same installed skill state without duplicating CLAUDE.md rule sections.
req_type: functional
priority: must
verification: integration test running syde install-skill twice and diffing the workspace
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Idempotence lets users re-run install without fear of corrupting their project.
---
