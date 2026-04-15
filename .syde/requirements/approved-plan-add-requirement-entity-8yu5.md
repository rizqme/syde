---
id: REQ-0003
kind: requirement
name: 'Approved plan: Add Requirement Entity'
slug: approved-plan-add-requirement-entity-8yu5
relationships:
    - target: add-requirement-entity-56dv
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T06:45:33Z"
statement: 'Add a requirement entity kind and validation rules so requirements are captured, linked, and governed: user-approved work creates requirements; conflicts are marked superseded or obsolete; every non-root entity traces to a requirement and parent; every contract participates in at least one flow.'
source: plan
source_ref: plan:add-requirement-entity-56dv
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-15T06:20:05Z"
---
