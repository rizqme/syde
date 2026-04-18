---
id: REQ-0385
kind: requirement
name: Requirement entity shall carry audited overlaps list
slug: requirement-entity-shall-carry-audited-overlaps-list-adu1
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:37:30Z"
statement: The syde requirement entity shall persist a list of acknowledged overlapping requirement slugs.
req_type: functional
priority: must
verification: RequirementEntity has an AuditedOverlaps field serialised to YAML.
source: plan
requirement_status: active
rationale: Audit state must round-trip through YAML.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:30Z"
---
