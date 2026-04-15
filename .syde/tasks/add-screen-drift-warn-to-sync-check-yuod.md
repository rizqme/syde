---
acceptance: Creating a fake page web/src/pages/Stub.tsx and running syde sync check produces a WARN for the unclaimed page. Removing the file clears the warning.
affected_entities:
    - audit-engine
affected_files:
    - internal/audit/audit.go
completed_at: "2026-04-14T11:37:11Z"
created_at: "2026-04-14T11:23:14Z"
details: 'internal/audit/audit.go: add CatScreenUnclaimed category. internal/audit/screens.go: new screenFindings(all, tree) function. For each tree node under web/src/pages matching *.tsx and not .test.tsx / .stories.tsx / ignored, check if any ContractEntity with ContractKind==''screen'' has the path in Files. Unowned files emit WARN Findings. Integrated into audit.Run alongside planPhaseFindings and conceptFindings.'
id: TSK-0080
kind: task
name: Add screen drift WARN to sync check
objective: syde sync check warns on web/src/pages/*.tsx files with no owning screen contract
plan_phase: phase_5
plan_ref: screen-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-screen-drift-warn-to-sync-check-yuod
task_status: completed
updated_at: "2026-04-14T11:37:11Z"
---
