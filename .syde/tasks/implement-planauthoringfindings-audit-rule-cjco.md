---
id: TSK-0036
kind: task
name: Implement planAuthoringFindings audit rule
slug: implement-planauthoringfindings-audit-rule-cjco
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: internal/audit/plan_authoring.go defines planAuthoringFindings(all []EntityWithBody) []Finding emitting the six checks listed in the phase details.
details: |-
    internal/audit/plan_authoring.go: new file. New CatPlanAuthoring constant in audit.go. The rule walks every plan whose status is draft/approved/in_progress and emits findings for:

    - WARN: Requirements lane is empty (plan declares no durable property).
    - WARN: ExtendedChange has no field_changes (programmatic verification disabled).
    - ERROR: NewChange draft missing kind-required fields. Component drafts must declare responsibility. Contract drafts must declare contract_kind, input, output. Requirement drafts must declare a statement that matches an EARS pattern (use model.MatchEARS). Concept drafts must declare meaning and invariants.
    - ERROR: **Orphan change** — every Deleted/Extended/NewChange must list at least one task slug in its tasks []string field, and every listed slug must resolve to an actual task entity in the plan's phase tree. Empty tasks list -> ERROR. Unresolvable task slug -> ERROR.
    - WARN: Orphan task — task whose slug does not appear in any change.tasks list. Verification-only tasks (build/test/smoke/refresh-tree) are exempt by name pattern.
    - WARN: Extended targets a screen contract but field_changes doesn't include wireframe.

    Reuse the existing auditGraph helper for slug resolution. Skip plans whose status is completed.
acceptance: Running syde sync check --strict on the in-flight plan surfaces at least one finding from each rule the plan deliberately violates.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_6
created_at: "2026-04-15T13:15:56Z"
completed_at: "2026-04-15T21:38:25Z"
---
