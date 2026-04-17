---
id: TSK-0014
kind: task
name: NewContractDraftView specialized per contract_kind
slug: newcontractdraftview-specialized-per-contractkind-hynw
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: 'NewChange previews for contract entities render specialized views per contract_kind: screen (UIML wireframe), cli (command + flags table), rest/rpc (METHOD + params tables), storage (key + fields), event/websocket (event + payload).'
details: web/src/components/NewContractDraftView.tsx switches on draft.contract_kind. Screen reuses the existing wireframe renderer that reads ContractEntity.Wireframe and outputs sanitized HTML. CLI renders draft.input as a monospace command header and draft.input_parameters as a flag/type/description table. REST/rpc does the same but styles input as METHOD + path and splits input_parameters into query/path/body based on a parameter location hint. Storage renders draft.input as a key pattern and draft.output_parameters as a fields table. Event shows draft.input as event name and parameters as payload.
acceptance: A plan with one NewChange per contract_kind renders five different previews without errors.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:19:14Z"
---
