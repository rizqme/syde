---
id: TSK-0051
kind: task
name: Add Tasks []string field to change types
slug: add-tasks-string-field-to-change-types-nb31
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: 'internal/model/plan.go: DeletedChange, ExtendedChange, and NewChange each carry a Tasks []string YAML+JSON-tagged ''tasks,omitempty'' field that lists the slugs of tasks implementing the change.'
details: Add Tasks []string with yaml:tasks,omitempty + json:tasks,omitempty to all three change struct types. Update validation to require at least one task (deferred until Phase 6 audit rule lands; structural validation can stay lenient at this layer).
acceptance: 'go build clean; round-tripping a plan YAML with tasks: [foo,bar] under a change preserves the list.'
affected_entities:
    - entity-model-f28o
affected_files:
    - internal/model/plan.go
    - internal/model/plan_test.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_1
created_at: "2026-04-15T13:38:32Z"
completed_at: "2026-04-15T15:03:57Z"
---
