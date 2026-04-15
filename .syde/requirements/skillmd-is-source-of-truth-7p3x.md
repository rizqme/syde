---
id: REQ-0094
kind: requirement
name: SKILL.md Is Source Of Truth
slug: skillmd-is-source-of-truth-7p3x
relationships:
    - target: skill-7fmf
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:07Z"
statement: The syde CLI shall treat the embedded skill/SKILL.md as the sole source of truth for skill behavior installed into user projects.
req_type: constraint
priority: must
verification: code review confirming install-skill reads from the embedded skill asset and not external locations
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: A single source prevents drift between the binary and installed skill files.
---
