---
id: REQ-0385
kind: requirement
name: Requirement entity shall carry audited overlaps list
slug: requirement-entity-shall-carry-audited-overlaps-list-adu1
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: The syde requirement entity shall persist a list of acknowledged overlapping requirement slugs.
req_type: functional
priority: must
verification: RequirementEntity has an AuditedOverlaps field serialised to YAML.
source: plan
requirement_status: active
rationale: Audit state must round-trip through YAML.
---
