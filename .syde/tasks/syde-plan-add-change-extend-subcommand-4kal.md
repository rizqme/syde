---
id: TSK-0006
kind: task
name: syde plan add-change extend subcommand
slug: syde-plan-add-change-extend-subcommand-4kal
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan add-change <plan> <kind> extend <slug> --what --why [--field k=v] appends an ExtendedChange.
details: Support repeatable --field key=value flags parsed into map[string]string. Field names are validated against the target kind's allowed frontmatter keys so typos are caught at authoring time (use a per-kind allowlist built from struct tags).
acceptance: Running extend with two --field entries stores them under field_changes in the plan file.
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_2
created_at: "2026-04-15T11:40:57Z"
completed_at: "2026-04-15T11:56:02Z"
---
