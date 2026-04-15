---
acceptance: A concept with no attributes fails syde sync check with a missing-attributes ERROR. A concept with relates_to label 'many' (not in enum) fails with a cardinality ERROR.
affected_entities:
    - entity-model
    - audit-engine
affected_files:
    - internal/model/validation.go
    - internal/audit/audit.go
completed_at: "2026-04-14T09:50:43Z"
created_at: "2026-04-14T09:44:26Z"
details: 'internal/model/validation.go: new case *ConceptEntity requiring Meaning and len(Attributes)>=1; each attribute needs Name and Type. internal/audit adds CatConceptIntegrity constant and conceptFindings() that walks plan entities of kind concept, validates every relates_to relationship with non-empty Label against the cardinality enum {one-to-one, one-to-many, many-to-one, many-to-many}. Integrate into Run() alongside planPhaseFindings.'
id: TSK-0059
kind: task
name: Concept validation and cardinality audit
objective: syde sync check --strict flags concepts missing meaning or attributes as ERROR and invalid cardinality labels as ERROR
plan_phase: phase_3
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: concept-validation-and-cardinality-audit-tzac
task_status: completed
updated_at: "2026-04-14T09:50:43Z"
---
