---
acceptance: Contracts page shows two extra filter dropdowns (Kind, Pattern). Selecting 'storage' prunes to the 7 BadgerDB contracts; 'rest' prunes to REST endpoints.
affected_entities:
    - web-spa
affected_files:
    - web/src/components/EntityFilterBar.tsx
    - web/src/App.tsx
completed_at: "2026-04-14T11:10:54Z"
created_at: "2026-04-14T11:07:44Z"
details: 'web/src/components/EntityFilterBar.tsx: when activeKind === ''contract'', render two additional select controls populated from distinct contract_kind and interaction_pattern values across the loaded entities. State hoisted to App.tsx (or the existing filter state if present). EntityList / the filter pruning logic checks the new filters alongside existing ones.'
id: TSK-0075
kind: task
name: Add kind + pattern filters to EntityFilterBar for contracts
objective: Contract list filterable by contract_kind and interaction_pattern
plan_phase: phase_2
plan_ref: storage-schema-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-kind-pattern-filters-to-entityfilterbar-for-contracts-k4zw
task_status: completed
updated_at: "2026-04-14T11:10:54Z"
---
