---
id: REQ-0169
kind: requirement
name: Skill Installer Does Not Own Skill Content
slug: skill-installer-does-not-own-skill-content-ctmq
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:46Z"
statement: The skill installer shall not own the skill content and shall read all templates from the embedded skill FS.
req_type: constraint
priority: must
verification: code review of internal/skill for embedded FS usage
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Skill content lives in skill/ so it can be edited and versioned without touching installer logic.
---
