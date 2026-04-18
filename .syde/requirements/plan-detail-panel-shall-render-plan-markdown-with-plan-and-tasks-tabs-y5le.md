---
id: REQ-0477
kind: requirement
name: Plan detail panel shall render plan markdown with Plan and Tasks tabs
slug: plan-detail-panel-shall-render-plan-markdown-with-plan-and-tasks-tabs-y5le
relationships:
    - target: plan-detail-panel-nqq1
      type: refines
updated_at: "2026-04-18T09:36:48Z"
statement: The plan detail panel component shall render the selected plan's markdown design and switch between Plan and Tasks tabs via the tab query parameter.
req_type: functional
priority: must
verification: Selecting a plan in the plans inbox shows the plan markdown by default and switching to the Tasks tab shows the phase-grouped task list
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Plan-detail-panel-nqq1 component had zero incoming refines from active requirements before this migration; this is the seed requirement satisfying the new component_no_requirement audit rule.
verified_against:
    plan-detail-panel-nqq1:
        hash: 39bbb2e6626805136a51b9e96a97e85736e2eff9aad23ad7cc67ef65d45f1543
        at: "2026-04-18T09:36:48Z"
---
