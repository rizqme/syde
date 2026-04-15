---
acceptance: syde query tag-index-key --format json shows contract_kind=storage, interaction_pattern=schema. Dashboard Tag Index Key detail panel renders storage + schema pills. skill install-skill in scratch dir lands the updated SKILL.md.
affected_entities:
    - skill-installer
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
completed_at: "2026-04-14T11:08:01Z"
created_at: "2026-04-14T11:06:07Z"
details: Run syde update on the seven index-key contracts with --contract-kind storage --interaction-pattern schema. Update skill/SKILL.md contract rules section to add 'storage' to the contract_kind enum list and mention 'schema' as the interaction pattern used with storage contracts. Mirror in skill/references/entity-spec.md.
id: TSK-0074
kind: task
name: Retag 7 BadgerDB contracts and document storage kind
objective: BadgerDB key contracts render as storage·schema; skill lists storage as a contract_kind value
plan_phase: phase_1
plan_ref: storage-schema-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: retag-7-badgerdb-contracts-and-document-storage-kind-b11r
task_status: completed
updated_at: "2026-04-14T11:08:01Z"
---
