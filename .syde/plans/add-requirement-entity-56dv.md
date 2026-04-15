---
id: PLN-0015
kind: plan
name: Add Requirement Entity
slug: add-requirement-entity-56dv
relationships:
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
    - target: syde-5tdt
      type: belongs_to
    - target: approved-plan-add-requirement-entity-8yu5
      type: references
      label: requirement
updated_at: "2026-04-15T06:47:18Z"
plan_status: completed
background: Users need syde to preserve requested product/design requirements as first-class, append-only design records. Requirements should be traceable from user statements and approved plans into every design entity, and conflicting or outdated requirements must remain historically visible instead of being deleted.
objective: 'Add a requirement entity kind and validation rules so requirements are captured, linked, and governed: user-approved work creates requirements; conflicts are marked superseded or obsolete; every non-root entity traces to a requirement and parent; every contract participates in at least one flow.'
scope: 'In scope: model/storage/index support for requirement entities, CLI/API/dashboard visibility, requirement capture workflow for user statements and approved plans, validation/audit rules for requirement traceability, contract-flow coverage, and belongs_to hierarchy. Out of scope: deleting requirements, LLM interpretation inside the syde binary, or changing existing relationship semantics beyond adding requirement-oriented relationship types/statuses.'
source: manual
created_at: "2026-04-15T06:17:37Z"
approved_at: "2026-04-15T06:20:05Z"
completed_at: "2026-04-15T06:47:18Z"
phases:
    - id: phase_1
      name: Requirement model
      status: completed
      description: Add first-class requirement entities to the design model.
      objective: Represent user-approved requirements as durable, queryable, append-only entities.
      changes: Add KindRequirement, requirements/ directory support, REQ ID prefix, RequirementEntity struct, serializer/index dispatch, relationship support, and query/status/list coverage.
      details: Define fields such as statement, source, source_ref, requirement_status active|superseded|obsolete, rationale, acceptance_criteria, supersedes, superseded_by, obsolete_reason, approved_at, and entity_refs/relationships. Requirements must never be physically deleted through normal CLI flows; remove should block or deprecate only.
      notes: 'New entity kind: requirement. Relationships expected: references/applies_to/from entities to requirements plus supersedes between requirements.'
      tasks:
        - implement-requirement-entity-model
    - id: phase_2
      name: CLI and API surfaces
      status: completed
      description: Expose requirement CRUD/read operations consistently.
      objective: Let humans and agents create, inspect, update status, and query requirements without editing .syde files directly.
      changes: Update add/update/get/list/query/status/context/reindex/write-client/dashboard API paths for requirement entities.
      details: Add syde add requirement and update flags for requirement-specific fields. Add safe status transitions for superseded/obsolete. Ensure requirement entities appear in status counts, context, full query output, dashboard lists/details, and search results.
      notes: Remove/delete must not delete requirements; it should reject or require a non-deleting obsolete/superseded transition.
      tasks:
        - expose-requirement-cli-and-api
    - id: phase_3
      name: Capture workflow
      status: completed
      description: Capture user statements and approved plans as requirements.
      objective: Make requirements a natural output of agent planning and approval rather than a manual afterthought.
      changes: Update Claude and Codex skills/hooks/docs so user requirements and approved syde plans create requirement records.
      details: User-stated requirements should become requirement entities during planning or explicit requirement capture. Plan approval should create or link requirements for the approved objective/scope/phases. The syde binary must not call an LLM; agents interpret user text and call syde commands.
      notes: Codex hooks can add context/reminders but cannot fully enforce non-Bash edits, so the skill and AGENTS.md must carry the behavioral contract.
      tasks:
        - capture-prompts-and-approved-plans-as-requirements
    - id: phase_4
      name: Conflict lifecycle
      status: completed
      description: Handle conflicting requirements without deleting history.
      objective: Ensure newer decisions explicitly supersede or obsolete conflicting requirements.
      changes: Add requirement validation and command flows for supersedes/obsolete links and statuses.
      details: When a requirement conflicts with an active prior requirement, the plan/agent must either link the new requirement as superseding the old one, mark the old one obsolete with a reason, or block validation. Superseded/obsolete requirements remain queryable and must retain provenance.
      notes: Conflict detection can begin as explicit metadata/validation rather than semantic LLM conflict detection inside syde.
      tasks:
        - validate-requirement-conflict-lifecycle
    - id: phase_5
      name: Traceability validation
      status: completed
      description: Require all design entities to trace to requirements.
      objective: Every design artifact should answer which requirement justified it.
      changes: Extend audit/sync check validation to require requirement links for all entities except allowed bootstrap/root requirement exceptions.
      details: Validation should require every non-requirement entity to have at least one relationship to a requirement or inherit requirement coverage through its parent only if explicitly designed. Requirement entities themselves should have provenance and may link to parent/root system as applicable.
      notes: Need a careful bootstrap path so existing projects can backfill requirements without making every command unusable immediately.
      tasks:
        - validate-requirement-traceability
    - id: phase_6
      name: Hierarchy validation
      status: completed
      description: Require belongs_to parents for all entities except the root system.
      objective: Keep the design graph navigable and prevent orphan architecture records.
      changes: Extend validation so every entity has a belongs_to parent except the single root system entity.
      details: Define root system selection rules. Flag duplicate root systems and entities without belongs_to. Requirement entities should belong to the root system or relevant parent unless explicitly allowed as root-level requirements.
      notes: Current repo has multiple system entities, so implementation must define root vs subsystem behavior and include migration/backfill guidance.
      tasks:
        - validate-entity-hierarchy
    - id: phase_7
      name: Contract-flow validation
      status: completed
      description: Require every contract to participate in at least one flow.
      objective: Ensure contracts are tied to end-to-end behavior, not isolated interface inventory.
      changes: Extend audit/sync check validation so every contract has at least one relationship to or from a flow.
      details: Accept direct flow relationships such as involves/references/uses depending on existing semantics. Report actionable errors with contract slug/name and suggested syde update commands. Add migration guidance for existing contracts.
      notes: The rule should include screen, command, query, API, and storage/schema contracts unless an explicit deprecated/obsolete escape hatch is defined.
      tasks:
        - validate-contract-flow-coverage
    - id: phase_8
      name: Backfill, docs, tests
      status: completed
      description: Backfill existing model data and document the workflow.
      objective: Ship the requirement feature without breaking existing syde projects unexpectedly.
      changes: Add tests/smoke checks, update README/skill references, add migration/backfill docs, and backfill this repository's existing entities to satisfy new validation rules.
      details: Create focused tests for model parsing, add/update/query, no-delete behavior, supersede/obsolete validation, requirement traceability, belongs_to root exceptions, and contract-flow checks. Update command docs and agent skills. Run go test ./..., syde sync check --strict, and summary-tree refresh.
      notes: May need a transitional flag or warning mode for existing projects before strict errors become default.
      tasks:
        - backfill-requirements-and-update-docs
---
