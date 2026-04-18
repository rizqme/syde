---
id: PLN-0018
kind: plan
name: Bidirectional requirement-component coupling with content-hash recheck gate
slug: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
relationships:
- target: approved-plan-bidirectional-requirement-component-coupling-with-content-hash-rechec-sgfy
  type: references
  label: requirement
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-18T08:21:26Z'
plan_status: completed
background: 'Today requirements link to architecture via belongs_to:system (loose ownership) and optional refines:component (used inconsistently — only 192/399 active reqs have it). This breaks the propagation chain the user wants: file change → component update → requirement re-check. Without a guaranteed component link, audit can''t surface ''which requirements should I re-validate after editing this file''. Symmetrically, components have no enforcement that any requirement constrains them, so a feature can ship architecturally with zero stated requirements.'
objective: Every active requirement refines ≥1 component. Every component with mapped files has ≥1 active requirement refining it. Cross-cutting (system-level) requirements are disallowed. When any file in a refining component changes content, the requirement is automatically marked stale until re-verified by the agent via syde requirement verify <slug>. PostToolUse hook surfaces affected requirements when files are edited. syde query --refined-by <comp> lists requirements per component.
scope: 'IN: requirement schema (verified_against field), four new audit findings, syde requirement verify CLI, syde query --refined-by filter, PostToolUse hook for affected-requirements surfacing, migration triage of 207 reqs lacking refines:component + 1 component lacking refining req, skill+entity-spec doc updates. OUT: backfilling rationale/verification text on existing reqs, content-hash for files outside any refining component, cross-component requirements (still allowed via multiple refines), drift detection for non-file changes (renames, deletes — out for v1).'
source: manual
created_at: '2026-04-18T07:54:59Z'
approved_at: '2026-04-18T08:00:32Z'
completed_at: '2026-04-18T08:21:26Z'
phases:
- id: phase_1
  name: Migration triage
  status: completed
  description: Walk the 207 active requirements lacking refines:component and assign each to ≥1 component; draft a refining requirement for plan-detail-panel-nqq1.
  objective: Every active requirement carries refines:component before audit enforcement turns on; the one orphan component (plan-detail-panel-nqq1) has at least one refining requirement.
  changes: 207 requirement files get refines:component added (1+ targets each); 207 requirement files get belongs_to:system removed; 1 new requirement created for plan-detail-panel-nqq1.
  details: 'Generate triage worksheet with heuristic suggestions: (a) intersect with components in the system the req currently belongs_to, (b) keyword-match component names against the requirement statement, (c) list any component file paths mentioned. Present to user, walk through assignments, then bulk-apply via direct YAML edit + syde reindex (per the bulk-reparenting workflow). Drop belongs_to:system from each migrated requirement at the same time.'
  notes: Direct YAML edit chosen over 207 'syde update' calls because the user pre-approved the bulk-reparenting workflow for migrations >50 entities. Verify with syde reindex after edits. Plan-detail-panel-nqq1 needs a brand-new requirement drafted (e.g. 'The plan detail panel shall render plan markdown with Plan and Tasks tabs.').
  tasks:
  - generate-migration-triage-worksheet
  - sample-and-validate-triage-assignments
  - bulk-apply-migrated-relationships-and-reindex
  - draft-seed-requirement-for-plan-detail-panel-nqq1
- id: phase_2
  name: Schema + audit enforcement
  status: completed
  description: Add the verified_against field to RequirementEntity and wire four new finding rules into the audit engine. Findings only emitted after Phase 1 migration so sync check stays green.
  objective: 'RequirementEntity carries verified_against snapshot field; audit emits four new findings: requirement_no_component, component_no_requirement, requirement_targets_system, requirement_stale.'
  changes: internal/model/entity.go gets new VerifiedAgainst field on RequirementEntity; internal/audit/requirements.go adds requirementNoComponent + requirementTargetsSystem + requirementStale; internal/audit/relationships.go (or new file) adds componentNoRequirement; CatRequirement category covers all four.
  details: 'RequirementEntity.VerifiedAgainst is a map keyed by component-canonical-slug, value is {hash: sha256-hex, at: RFC3339-timestamp}. Audit rules: requirement_no_component fires when an active req has zero refines:component edges; component_no_requirement fires when a component has files mapped AND zero incoming refines from active reqs; requirement_targets_system fires when an active req has any refines:system or belongs_to:system edge; requirement_stale fires when, for any refining component, the SHA-256 of any file in component.files differs from the corresponding stored hash (or no entry exists). All four use SeverityFinding (the only severity).'
  notes: verified_against entries created lazily — first 'syde requirement verify' call snapshots; before that, requirement_stale fires for all. This is intentional — Phase 3 task does the bulk initial verify after schema lands.
  tasks:
  - add-verifiedagainst-field-to-requiremententity
  - implement-requirementnocomponent-audit-finding
  - implement-componentnorequirement-audit-finding
  - implement-requirementtargetssystem-audit-finding
  - implement-requirementstale-audit-finding-content-hash-gate
  - verify-post-phase-2-sync-check-passes
- id: phase_3
  name: Verify CLI and initial snapshot
  status: completed
  description: Implement the verify command and run it once across all active requirements so the initial verified_against state matches current code.
  objective: syde requirement verify <slug> snapshots SHA-256 of every file in each refining component into the requirement's verified_against; bulk-verify run on all 399 active reqs after migration.
  changes: internal/cli/requirement.go (new file) implements 'syde requirement verify <slug>'; rootCmd registers it; one-shot scripts/bulk-verify-requirements.sh (or inline bash) runs verify for all active reqs after Phase 1+2 land.
  details: 'syde requirement verify <slug> resolves the requirement, walks each refines:component edge, opens each file in component.files, computes SHA-256, writes verified_against[component-canonical-slug] = {hash, at: now}. Saves the requirement file via existing FileStore. Bulk task: shell loop or scripts/bulk-verify-requirements.sh (one-shot).'
  notes: 'Sub-task: write a small Go helper utils.FileSHA256 if not already present. The command must be idempotent — calling verify on an already-verified req just refreshes timestamps.'
  tasks:
  - bulk-verify-all-399-active-requirements-initial-snapshot
  - implement-syde-requirement-verify-cli-command
- id: phase_4
  name: 'Discovery CLI: --refined-by'
  status: completed
  description: Add the inverse-of-refines query so agents and humans can ask 'what requirements constrain this component?' as a single command.
  objective: syde query --refined-by <component-slug> lists active requirements that refine the named component.
  changes: internal/cli/query.go adds --refined-by flag; internal/query/engine.go adds matching resolver method; refs format already returns slug list, so no formatter change.
  details: Mirrors --depended-by/--depends-on patterns. Resolves the component slug via the same alias map used by audit, walks all requirement entities, filters those with active status AND a refines edge whose resolved target equals the queried component.
  notes: Output uses the existing rich/json/compact/refs formatters (no new format needed).
  tasks:
  - implement-resolverefinedby-in-query-engine
  - add-refined-by-flag-to-syde-query
- id: phase_5
  name: PostToolUse hook + skill docs
  status: completed
  description: Wire the skill's PostToolUse hook to surface affected requirements, and update the user-facing skill docs to teach the bidirectional invariant + recheck workflow.
  objective: When the agent edits a file mapped to any component, the syde skill emits the affected active requirements as a re-verification reminder; SKILL.md and entity-spec.md document the new model.
  changes: skill/hooks.go PostToolUse handler extended; skill/SKILL.md adds Recheck section; skill/references/entity-spec.md updates Requirement section.
  details: Hook receives the edited file path, looks up owning component via existing constraints-check API, calls /api/<proj>/query?refined-by=<comp-slug>, returns a context block listing affected requirement slugs + a hint to run syde requirement verify after re-reading them. SKILL.md gains a 'Recheck affected requirements' subsection in the workflow rules. entity-spec.md updates the Requirement section to describe verified_against, the bidirectional invariant, and the gate behaviour.
  notes: Hook lives in skill/hooks.go (PostToolUse handler). Hook is no-op when no component owns the edited path (avoids noise on .syde/, scripts/, docs).
  tasks:
  - update-skillmd-with-recheck-workflow
  - update-entity-specmd-requirement-section
  - implement-posttooluse-affected-requirements-hook
changes:
  requirements:
    new:
    - id: c7q9
      name: Active requirement shall refine at least one component
      what: Audit finding when an active requirement has zero refines:component edges
      why: Without a guaranteed component link the propagation chain (file→component→requirement) has dead ends; agents cannot answer 'which requirements should I re-verify after editing this file'.
      draft:
        priority: must
        rationale: Bidirectional invariant requires every requirement to constrain a concrete component.
        req_type: functional
        source: plan
        statement: When an active requirement has no refines edge to a component, the syde audit engine shall report a finding.
        verification: syde sync check errors when an active requirement is created or migrated without a refines:component edge
    - id: s2pq
      name: Component with mapped files shall have at least one refining requirement
      what: Audit finding when a component has files mapped and zero incoming refines from active reqs
      why: Components that have shipped code without stated requirements are architecturally unanchored — nothing constrains their behaviour and nothing triggers re-verification when they change.
      draft:
        priority: must
        rationale: Symmetric half of the bidirectional invariant; design-phase components without files are exempt.
        req_type: functional
        source: plan
        statement: When a component has at least one file in its files list and no active requirement refines it, the syde audit engine shall report a finding.
        verification: syde sync check errors when a component with files mapped has zero incoming active-requirement refines edges
    - id: ccrf
      name: Requirement shall not refine or belong to a system
      what: Audit finding when an active requirement has refines:system or belongs_to:system
      why: User decided cross-cutting requirements are disallowed (Q1). System-level reqs must be decomposed per component so the propagation chain remains symmetric and precise.
      draft:
        priority: must
        rationale: Cross-cutting reqs hide which component is actually responsible; force decomposition.
        req_type: functional
        source: plan
        statement: If a requirement carries a refines edge to a system or a belongs_to edge to a system, then the syde audit engine shall report a finding.
        verification: syde sync check errors when any active requirement targets an entity of kind=system via refines or belongs_to
    - id: ja3g
      name: Requirement shall be marked stale when refining component file content changes
      what: Audit finding when SHA-256 of any file in a refining component differs from the requirement's stored verified_against entry
      why: User asked for a deterministic recheck gate (Q3). Content hash is more robust than mtime — survives file copies, format-only edits don't trigger.
      draft:
        priority: must
        rationale: Content-hash chosen over mtime for resilience against copy/touch operations and to ignore format-only changes.
        req_type: functional
        source: plan
        statement: Where any file in a component refining a requirement has a SHA-256 content hash differing from the requirement's stored verified_against entry, the syde audit engine shall report a finding.
        verification: Editing a file in a refining component without subsequently running 'syde requirement verify' on the requirement causes syde sync check to error
    - id: 7h0k
      name: syde requirement verify shall snapshot SHA-256 hashes for every file in each refining component
      what: New CLI command that updates verified_against for a requirement
      why: Operators need a deterministic action to clear the stale-hash finding after re-reading the requirement and confirming it still holds.
      draft:
        priority: must
        rationale: Verify is idempotent — re-running refreshes timestamps; failure to verify any file aborts the snapshot to avoid partial state.
        req_type: functional
        source: plan
        statement: When the user runs syde requirement verify against a requirement slug, the syde requirement entity shall snapshot the SHA-256 content hash of every file in each refining component into verified_against.
        verification: Running syde requirement verify <slug> updates verified_against entries for each refining component to the current SHA-256 hashes and timestamp
    - id: v0o7
      name: syde query shall support --refined-by component slug
      what: New CLI flag that lists active requirements refining a named component
      why: Inverse-of-refines lookup is the discovery primitive agents use to ask 'what requirements constrain this component'.
      draft:
        priority: should
        rationale: Mirrors --depended-by/--depends-on naming; resolves component slug via the same alias map the audit uses.
        req_type: functional
        source: plan
        statement: The syde query command shall support a --refined-by component-slug flag that lists every active requirement carrying a refines edge to the resolved component.
        verification: syde query --refined-by <component-slug> returns the slugs of every active requirement with refines edges resolving to that component
    - id: tjh8
      name: PostToolUse hook shall surface affected requirements when a file mapped to a component is edited
      what: Skill hook that emits affected requirement slugs into agent context after Edit or Write
      why: 'Closes the propagation loop: file edit → component owner lookup → list of requirements to re-verify, surfaced to the agent without requiring a manual query.'
      draft:
        priority: should
        rationale: Hook is no-op for paths owned by no component (e.g. .syde/, scripts/, docs).
        req_type: functional
        source: plan
        statement: When a PostToolUse Edit or Write touches a path mapped to a component files list, the syde skill hook shall surface the affected active requirements into the agent context.
        verification: Editing a file owned by a component causes the next agent turn to receive a context block listing the active requirements refining that component
    - id: 6gro
      name: Plan detail panel shall render plan markdown with Plan and Tasks tabs
      what: First refining requirement for plan-detail-panel-nqq1 — the one component currently lacking any
      why: Component has files mapped (web/src/components/PlanDetailPanel.tsx) but zero refining requirements; Phase 1 migration must close this gap before Phase 2 audit enforcement turns on.
      draft:
        priority: must
        rationale: Existing plan-detail-panel-nqq1 component has zero incoming refines from active requirements; this is the seed requirement to satisfy the new component_no_requirement audit rule.
        req_type: functional
        source: plan
        statement: The plan detail panel component shall render the selected plan's markdown design and switch between Plan and Tasks tabs via the tab query parameter.
        verification: Selecting a plan in the plans inbox shows the plan markdown by default and switching to the Tasks tab shows the phase-grouped task list
  concepts:
    new:
    - id: w8g5
      name: Requirement
      what: 'Define the Requirement concept: an EARS-format constraint on ≥1 component, refreshed via verified_against content-hash snapshots'
      why: There is no Requirement concept today even though it is the most-populated kind (410 entities) and gains significant new semantics in this plan.
      draft:
        invariants: Active requirements have ≥1 refines:component edge. Components with files have ≥1 incoming refines from active reqs. No requirement targets a system via refines or belongs_to. verified_against is a map keyed by component canonical slug.
        meaning: A Requirement is an EARS-format constraint that one or more components must satisfy. Every active requirement carries refines edges to ≥1 component (no system-level reqs). When the SHA-256 content of any refining component file diverges from the requirement's last verified_against snapshot, the requirement is automatically marked stale and must be re-verified by an agent (syde requirement verify) after re-reading the requirement and confirming it still holds against the current code.
  components:
    extended:
    - id: cpum
      slug: entity-model-f28o
      what: Add VerifiedAgainst map[string]VerifiedSnapshot field (keyed by component canonical slug; VerifiedSnapshot{Hash, At RFC3339}) to RequirementEntity in internal/model/entity.go. This is a Go-struct-level change, not a syde entity field change.
      why: Stores per-refining-component SHA-256 snapshots so the audit can deterministically detect when refining components have drifted.
    - id: mczu
      slug: audit-engine-4ktg
      what: Add four new audit rules in a new internal/audit/bidirectional.go file — requirement_no_component (CatRequirement), component_no_requirement (CatTraceability), requirement_targets_system (CatRequirement), requirement_stale (CatRequirement). All use SeverityFinding.
      why: Implements the bidirectional invariant + content-hash recheck gate as audit findings (single severity, blocks sync check).
    - id: cqbm
      slug: cli-commands-hpjb
      what: Add 'syde requirement' parent command with 'verify <slug>' subcommand in a new internal/cli/requirement.go; add --refined-by flag to 'syde query' in internal/cli/query.go.
      why: Operators need (a) a deterministic action to clear requirement_stale findings after re-reading reqs, and (b) a discovery primitive for the inverse-of-refines lookup.
    - id: 9l5g
      slug: query-engine-9k84
      what: Add RefinedBy(componentSlug) method to the Engine type in internal/query/engine.go — uses existing four-way alias map to canonicalize the component slug, walks active requirements, filters by refines edge resolving to the component.
      why: Backs both the CLI --refined-by flag and the PostToolUse hook's affected-requirements lookup.
    - id: qztj
      slug: skill-installer-wbmu
      what: Extend skill/hooks.json with a PostToolUse Write|Edit|MultiEdit handler that surfaces active refining requirements; add a Recheck section to skill/SKILL.md; update the Requirement section in skill/references/entity-spec.md to document verified_against and the gate workflow. All bundled via go:embed.
      why: Skill files are go:embed'd at build time; any new hook logic or doc changes need to be in the embedded source tree.
---
