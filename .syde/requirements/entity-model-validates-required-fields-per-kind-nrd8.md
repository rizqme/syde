---
id: REQ-0141
kind: requirement
name: Entity Model Validates Required Fields Per Kind
slug: entity-model-validates-required-fields-per-kind-nrd8
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:38:03Z"
statement: When ValidateEntity is invoked, the entity model shall return validation errors for any missing required or recommended fields defined for the entity kind.
req_type: functional
priority: must
verification: unit test of ValidateEntity in internal/model/validation.go
source: manual
source_ref: component:entity-model-f28o
requirement_status: active
rationale: Per-kind validation guarantees entities meet minimum schema requirements before persistence.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:38:03Z"
---
