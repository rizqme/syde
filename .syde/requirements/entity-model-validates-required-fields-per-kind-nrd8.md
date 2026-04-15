---
id: REQ-0141
kind: requirement
name: Entity Model Validates Required Fields Per Kind
slug: entity-model-validates-required-fields-per-kind-nrd8
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:26Z"
statement: When ValidateEntity is invoked, the entity model shall return validation errors for any missing required or recommended fields defined for the entity kind.
req_type: functional
priority: must
verification: unit test of ValidateEntity in internal/model/validation.go
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Per-kind validation guarantees entities meet minimum schema requirements before persistence.
---
