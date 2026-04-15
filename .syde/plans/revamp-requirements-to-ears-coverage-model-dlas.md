---
id: PLN-0022
kind: plan
name: Revamp Requirements To EARS Coverage Model
slug: revamp-requirements-to-ears-coverage-model-dlas
relationships:
    - target: approved-plan-revamp-requirements-to-ears-coverage-model-4yeg
      type: references
      label: requirement
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T11:09:22Z"
plan_status: completed
background: Current requirements are sparse (15), free-form prose, indistinguishable from tasks, and a single baseline requirement is linked by hundreds of entities. Decisions overlap with requirements conceptually, and the design kind is a vestigial stub with zero instances. Requirements should be dense, EARS-structured declarative system properties, and strict sync should enforce both well-formed-ness and coverage across every component, contract, concept, and system.
objective: Reset the requirement store, absorb decisions, remove the design kind, tighten the model to EARS + req_type/priority/verification/refines, programmatically enforce good requirements and coverage, then systematically backfill 150–250 narrow requirements derived from components/contracts/concepts/system via a documented deterministic derivation algorithm.
scope: 'Phase 1 reset + absorb: delete existing requirements, demote non-coverage traceability to WARN, convert decisions to requirements, delete decision and design entity kinds. Phase 2 model + algorithm: add req_type/priority/verification/refines, EARS validator, CLI flags, dashboard, and document the derivation algorithm. Phase 3 audits: good-requirement check (EARS + fields) and coverage check (component/contract/concept/system). Phase 4 authoring: subagent-dispatched backfill following the documented algorithm. Phase 5 verify. EXCLUDED: touching plan/task entities, rewriting or re-linking existing tasks/plans — user will revamp separately later.'
source: manual
created_at: "2026-04-15T09:51:59Z"
approved_at: "2026-04-15T10:07:06Z"
completed_at: "2026-04-15T11:09:22Z"
phases:
    - id: phase_1
      name: Reset and absorb
      status: completed
      description: Delete existing requirements, demote dangling traceability to WARN, convert decisions to requirements, remove decision and design entity kinds.
      objective: The requirement store is empty except for decisions rewritten as EARS requirements, and KindDecision + KindDesign are gone from Go code.
      changes: Delete .syde/requirements/*, relax requirementTraceFindings severity, convert DecisionEntity instances to RequirementEntity files, remove KindDecision + KindDesign + all dispatch/validation/query/CLI/API/dashboard references.
      details: 'Follow the same removal pattern used for learning: model dispatch, validation, query resolver/engine, dashboard API + UI, CLI flags, skill docs, installed skill copy refresh. Decision conversion rewrites each decision''s statement into EARS form with req_type inferred from category (constraint/functional/non-functional) and verification inferred from consequences.'
      tasks:
        - delete-existing-requirements-and-relax-traceability
        - convert-decisions-to-ears-requirements
        - remove-decision-entity-kind-from-go
        - remove-design-entity-kind-from-go
        - refresh-installed-skill-copies-after-phase-1
    - id: phase_2
      name: Model tightening and derivation algorithm
      status: completed
      description: Add req_type, priority, verification, refines/derives_from, EARS validator, CLI and dashboard plumbing, and write the deterministic derivation algorithm doc used by Phase 4 subagents.
      objective: RequirementEntity carries req_type + priority + verification; refines/derives_from relationship types exist; save-time validator rejects non-EARS statements; skill references include requirement-derivation.md.
      changes: internal/model/entity.go (RequirementEntity fields, enums), internal/model/relationship.go (rel constants), internal/model/validation.go (EARS regex check), internal/cli/add.go + update.go (flags), internal/dashboard/api.go + web EntityDetail, skill/references/entity-spec.md + commands.md + requirement-derivation.md, skill SKILL.md, make install + syde install-skill --all.
      details: 'Verification field replaces acceptance_criteria. Five EARS pattern regexes cover Ubiquitous, Event-driven, State-driven, Unwanted-behavior, Optional-feature. Derivation algorithm doc specifies per-kind procedures: Component (responsibility + each capability + boundaries NOT-lines + failure_modes + scaling_notes), Contract (input-output + each parameter + constraints), System (quality_goals + design_principles), Concept (invariants + lifecycle + data_sensitivity + required attributes). Quality checks: pattern match, 10-250 chars, uniqueness, no task verbs, back-link via refines.'
      tasks:
        - add-reqtype-priority-verification-fields
        - add-refines-and-derivesfrom-relationship-types
        - implement-ears-statement-validator
        - wire-requirement-cli-flags-and-dashboard-fields
        - write-requirement-derivation-algorithm-doc
        - update-skill-docs-and-refresh-installed-copies
    - id: phase_3
      name: Audit rules for good and covered requirements
      status: completed
      description: 'Implement two new audit rules: good-requirement (EARS plus complete fields) and coverage (component/contract/concept/system must be connected to a requirement).'
      objective: syde sync check --strict reports ERROR for any requirement missing EARS/type/priority/verification and ERROR for any component/contract/concept/system without a requirement link.
      changes: internal/audit/graph_rules.go (new goodRequirementFindings + coverageFindings functions), internal/audit/audit.go (wire new rules into Run).
      details: 'goodRequirementFindings iterates requirement entities and checks: statement matches one of 5 EARS patterns, req_type in allowed enum, priority in allowed enum, verification non-empty. coverageFindings builds the set of entities linked by any requirement relationship (inbound or outbound) and emits ERROR for each component/contract/concept/system not in the set. Flows/plans/tasks/designs excluded from coverage per user scope.'
      tasks:
        - good-requirement-audit-rule
        - coverage-audit-rule
    - id: phase_4
      name: Dense requirement authoring
      status: completed
      description: Systematically backfill ~150-250 narrow EARS requirements across the system, components, contracts, and concepts, dispatched per subsystem via subagents following the Phase 2 derivation algorithm.
      objective: Every component, contract, concept, and system entity is linked by at least one requirement, and the good-requirement and coverage audits pass with only the expected task/plan WARNs remaining.
      changes: .syde/requirements/*.md (many new files), relationship backlinks from source entities to their refining requirements.
      details: 'Subagents receive skill/references/requirement-derivation.md as their prompt plus a batch of entities to walk. Batches: system entity, syde-cli components, syded-dashboard components, shared core components, CLI contracts, HTTP API contracts, screen contracts, concepts. Each generated requirement gets a refines relationship back to its source entity so coverage and traceability are both satisfied.'
      tasks:
        - backfill-system-entity-requirements
        - backfill-syde-cli-component-requirements
        - backfill-syded-dashboard-component-requirements
        - backfill-cli-contract-requirements
        - backfill-http-api-contract-requirements
        - backfill-screen-contract-requirements
        - backfill-concept-requirements
    - id: phase_5
      name: Verify
      status: completed
      description: Run the full health gate and refresh the summary tree.
      objective: go test and make install are clean, syde sync check --strict passes for the new audit rules (task/plan traceability WARNs are expected), and the summary tree is clean under --strict.
      changes: No source changes; test, build, install, and summary tree refresh.
      details: Run go test ./..., make install, syde install-skill --all, syde sync check --strict, syde tree scan + leaves-first summarize loop until syde tree status --strict exits 0.
      tasks:
        - strict-sync-and-build-smoke-test
        - refresh-summary-tree
---
