---
id: TSK-0059
kind: task
name: Rewrite contractFlowFindings to use steps
slug: rewrite-contractflowfindings-to-use-steps-a57b
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: Audit checks flow steps for contract refs instead of relationship edges
details: Rewrite contractFlowFindings in graph_rules.go. Iterate all FlowEntity Steps, collect contract slugs into inFlow set. WARN on steps with empty Contract field. ERROR on contracts not in any flow step's Contract field. Resolve contract slugs via auditGraph lookup.
acceptance: syde sync check surfaces ERRORs for contracts without step refs after catch-all is removed
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_2
created_at: "2026-04-16T09:22:29Z"
completed_at: "2026-04-16T10:40:04Z"
---
