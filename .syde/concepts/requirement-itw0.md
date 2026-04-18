---
id: CPT-0012
kind: concept
name: Requirement
slug: requirement-itw0
description: An EARS-format constraint that one or more components must satisfy, anchored bidirectionally via refines:component.
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-18T08:45:16Z"
meaning: A Requirement is an EARS-format constraint that one or more components must satisfy. Every active requirement carries refines edges to ≥1 component (no system-level reqs). When the SHA-256 content of any refining component file diverges from the requirement's last verified_against snapshot, the requirement is automatically marked stale and must be re-verified by an agent (syde requirement verify) after re-reading the requirement and confirming it still holds against the current code.
invariants: Active requirements have ≥1 refines:component edge. Components with files have ≥1 incoming refines from active reqs. No requirement targets a system via refines or belongs_to. verified_against is a map keyed by component canonical slug.
---
