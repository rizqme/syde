---
approved_at: "2026-04-14T11:06:07Z"
background: Seven BadgerDB key contracts (Tag/Entity/Slug/Word/Counter/Outgoing Rel/Incoming Rel Index Keys) are currently tagged contract_kind=event, interaction_pattern=pub-sub. That's wrong — they aren't events, they're storage schemas (key prefix templates the Storage Engine writes and reads). The screenshot showing 'event · pub-sub' on Tag Index Key exposed the mis-classification. The SKILL.md already talks about 'Schema contracts' as a first-class concept but the documented contract_kind enum has no 'storage' value so these contracts were shoehorned into 'event'.
completed_at: "2026-04-14T11:10:55Z"
created_at: "2026-04-14T11:05:35Z"
id: PLN-0011
kind: plan
name: storage-schema-contract-kind
objective: Every BadgerDB key contract reads as contract_kind=storage, interaction_pattern=schema. The SKILL.md contract-kind enum is extended to include 'storage' for data schemas (KV prefixes, SQL tables, proto messages, cache keys) with 'schema' as the matching interaction pattern. Future schema contracts get the right kind on creation.
phases:
    - changes: syde update on each of Tag/Entity/Slug/Word/Counter/Outgoing Rel/Incoming Rel Index Keys with --contract-kind storage --interaction-pattern schema. skill/SKILL.md and skill/references/entity-spec.md add storage to the contract_kind enum with a note that schema-style contracts (KV key prefix, SQL table, proto message, queue topic, cache key) use it.
      description: Flip 7 BadgerDB contracts to storage/schema and document the new enum values in the skill
      details: Seven contracts, one shell script. Then edit SKILL.md and entity-spec.md to add 'storage' to the contract_kind list and 'schema' to the interaction_pattern list. commands.md already says 'free-form' so the flag reference stays valid.
      id: phase_1
      name: Retag + document storage/schema
      notes: contract_kind is a free-form string on ContractEntity today — the validator only requires non-empty. No Go code changes needed.
      objective: BadgerDB schema contracts render as storage · schema in the dashboard; skill explicitly lists storage as a contract_kind option
      status: completed
      tasks:
        - retag-7-badgerdb-contracts-and-document-storage-kind
    - changes: web/src/components/EntityFilterBar.tsx adds two new dropdowns (Kind, Pattern) that only render when activeKind === 'contract'. The dropdowns populate their options dynamically from the loaded entities so new values (like storage/schema) show up automatically. Filter state wiring in App.tsx / EntityList passes the new filters through and prunes the list accordingly.
      description: EntityFilterBar grows contract_kind and interaction_pattern filters for the contract kind page
      details: Inspect the current filter bar + how filter state flows into EntityList. Add kind + pattern multi-select or single-select dropdowns. Prune the list in-place. No server-side filter needed — all entities already come down in the list response.
      id: phase_2
      name: 'Filter bar: kind and pattern'
      notes: Dashboard-only change — no Go/CLI code involved.
      objective: Users can filter contracts by contract_kind and interaction_pattern via the filter bar on the Contracts page
      status: completed
      tasks:
        - add-kind-pattern-filters-to-entityfilterbar-for-contracts
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) update the 7 BadgerDB key contracts via syde update; (2) document ''storage'' as a contract_kind value and ''schema'' as an interaction_pattern value in skill/SKILL.md and skill/references/entity-spec.md. Out-of-scope: validator-enforced enum for contract_kind (still free-form string), retroactive re-tagging of other schema contracts outside the BadgerDB index keys, ERD changes.'
slug: storage-schema-contract-kind-2ind
source: manual
updated_at: "2026-04-14T11:10:55Z"
---
