---
id: PLN-0019
kind: plan
name: Remove root system and allow components to belong to multiple standalone systems
slug: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: design-model-shall-contain-exactly-two-systems-named-syde-and-syded-f2q8
      type: implements
    - target: approved-plan-remove-root-system-and-allow-components-to-belong-to-multiple-standal-5cvq
      type: references
      label: requirement
updated_at: "2026-04-18T10:00:54Z"
plan_status: completed
background: The current model requires exactly one root system (syde-5tdt) with every other system and entity carrying belongs_to to anchor into that hierarchy. This has produced three systems (syde, syde CLI, syded Dashboard) with confused ownership — HTTP API belongs_to root syde but is daemon-exclusive, Audit Engine belongs_to syde CLI but runs in both processes, etc. The root-system abstraction doesn't match the two actual process boundaries (syde binary + syded binary) and forces components that live in both into an arbitrary single-owner.
objective: 'Model has exactly two systems (syde + syded) — both standalone top-level with no belongs_to. Components reparent to belongs_to: syde and/or syded per their actual process reachability from cmd/syde/main.go and cmd/syded/main.go. Audit rules drop the root-system requirement and accept multi-system belongs_to on components. Frontend renders all systems at a single visual tier.'
scope: 'IN: audit rule changes (hierarchy root rule removed, belongs_to forbidden on systems, multi belongs_to allowed on components), frontend Graph.tsx tier collapse, 12 component + 57 contract reparenting, syde-cli-2478 deletion, syded-dashboard-e82c rename to ''syded'', skill+entity-spec+CLAUDE.md doc updates. OUT: removing belongs_to entirely for non-system entities (they still need ≥1), system→system relationships other than belongs_to (depends_on/exposes remain fine), refactoring components that span both systems into separate shared libraries.'
source: manual
created_at: "2026-04-18T09:07:08Z"
approved_at: "2026-04-18T10:00:13Z"
completed_at: "2026-04-18T10:00:54Z"
phases:
    - id: phase_1
      name: Audit rules + frontend tier collapse
      status: completed
      description: Modify internal/audit/graph_rules.go so systems are forbidden from carrying belongs_to (they are top-level by model rule), drop the exactly-one-root-system check, and keep the non-system entities needing ≥1 belongs_to. Simplify Graph.tsx to treat every system kind identically instead of splitting system-root vs system-sub.
      objective: hierarchyFindings drops 'one root system' rule and forbids belongs_to on systems; Graph.tsx renders every system at a single visual tier.
      changes: 'internal/audit/graph_rules.go: remove root-selection + ''must have belongs_to parent'' for systems; add ''system must not have belongs_to'' finding. web/src/pages/Graph.tsx: drop subSystemIds + system-root/system-sub sizeKeys.'
      details: 'In hierarchyFindings: replace the root-selection logic with a flat check — every system entity with any belongs_to edge is a finding; every non-system non-requirement entity without belongs_to is a finding. In Graph.tsx: remove subSystemIds calculation and set sizeKey=''system'' uniformly; drop the LEGEND entries that distinguish root vs sub.'
      notes: 'These checks activate immediately — before Phase 2 runs, the syded-dashboard-e82c → syde belongs_to edge will fail the new systems-cannot-have-belongs_to rule. That is intentional: Phase 2 removes that edge as part of the rename.'
      tasks:
        - rewrite-hierarchyfindings-to-drop-root-system-rule-and-forbid-belongsto-on-systems
        - simplify-graphtsx-to-render-every-system-at-one-tier
    - id: phase_2
      name: System entity consolidation
      status: completed
      description: Collapse the redundant syde-cli-2478 system into syde-5tdt. Rename syded-dashboard-e82c name field from 'syded Dashboard' to 'syded' while preserving the slug (to avoid churn in the 57 refining reqs).
      objective: 'Exactly two systems remain: syde (CLI) and syded (daemon). syde-cli-2478 is deleted after its 12 components + 57 contracts are reparented. syded-dashboard-e82c is renamed to ''syded'' and loses its belongs_to:syde edge.'
      changes: 'systems/syded-dashboard-e82c.md: name ''syded Dashboard'' → ''syded''; remove belongs_to:syde relationship. systems/syde-cli-2478.md: deleted after Phase 3 completes.'
      details: 'Order matters: Phase 3 reparents first so syde-cli-2478 ends up with 0 owned entities, then Phase 2 deletes it. Renaming syded-dashboard-e82c is a YAML-field edit. The belongs_to:syde removal makes syded standalone.'
      notes: Keep syded-dashboard-e82c slug — renaming would force updating ~25 refining requirements and every relationship target. Only the name field changes.
      tasks:
        - rename-syded-dashboard-e82c-name-field-to-syded-and-remove-belongstosyde
        - delete-syde-cli-2478-system-entity
    - id: phase_3
      name: Reparent components and contracts
      status: completed
      description: 'Generate assignment worksheet using import-graph heuristic: for each component, resolve its files and trace which cmd/ entrypoint imports them; for each contract, inspect its input signature (syde-prefix → syde, GET/POST /api/ → syded, ws:// → syded).'
      objective: All 12 components and 57 contracts currently belongs_to syde-cli-2478 get reparented to syde-5tdt and/or syded-dashboard-e82c per import-graph reachability.
      changes: 'components/*.md (12 files) and contracts/*.md (57 files): belongs_to edges repointed from syde-cli-2478 to syde-5tdt and/or syded-dashboard-e82c. Several components gain a second belongs_to edge (multi-system).'
      details: 'Shared components (Audit Engine, Entity Model, Storage Engine, Query Engine, Graph Engine, Slug and ID Utils, Summary Tree, UIML Parser, Config Loader, Scan Helpers, Skill Installer) gain belongs_to: syde-5tdt + belongs_to: syded-dashboard-e82c. CLI-only (CLI Commands, CLI HTTP Client, Daemon Launcher) become belongs_to: syde-5tdt. Daemon-only (Project Registry, HTTP API, WebSocket Server, Dashboard Daemon Entry Point, Plan Detail Panel, Web SPA) stay belongs_to: syded-dashboard-e82c. Contracts: syde CLI invocation contracts → syde-5tdt; HTTP/screen/websocket contracts → syded-dashboard-e82c. Generate triage worksheet, sample 15 rows for user validation, bulk-apply via YAML edit + reindex (per the bulk-reparenting workflow).'
      notes: User pre-approved the direct YAML edit + reindex workflow for migrations >50 entities. Sample-and-validate gate before bulk apply.
      tasks:
        - generate-reparenting-worksheet-for-12-components-57-contracts
        - sample-and-validate-reparenting-assignments
        - bulk-apply-component-and-contract-reparenting-and-reindex
    - id: phase_4
      name: Documentation updates
      status: completed
      description: 'Update the user-facing skill docs and the project''s own CLAUDE.md to reflect the new rules: no root system, no system→system belongs_to, components may carry belongs_to to multiple systems.'
      objective: SKILL.md, entity-spec.md, and CLAUDE.md describe systems as standalone top-level entities with no belongs_to, and document multi-system belongs_to on components.
      changes: skill/SKILL.md system rules section; skill/references/entity-spec.md System section; CLAUDE.md syde rules section.
      details: Touch skill/SKILL.md (system rules section — remove root/subsystem language, state belongs_to forbidden on systems, add multi-belongs_to note for components). Touch skill/references/entity-spec.md (System section — describe standalone-process semantics). Touch CLAUDE.md (any 'root system' references).
      notes: Keep the language tight — the behaviour change is small, the docs should just reflect it accurately without bloat.
      tasks:
        - rewrite-system-rules-in-skillmd
        - rewrite-system-section-in-entity-specmd
        - purge-root-system-language-from-claudemd
        - draft-referencesplan-authoringmd
        - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
        - implement-syde-plan-review-slug-cli-command
        - draft-referencesplan-review-promptmd
        - build-reinstall-to-bundle-the-two-new-reference-docs
    - id: phase_5
      name: Verify + refresh tree
      status: completed
      description: Run sync check end-to-end. Re-verify any requirements whose refining component files changed during the audit-rule edits. Refresh summary tree summaries for touched files and propagate stale up to root.
      objective: syde sync check exits 0; summary tree fully clean.
      changes: No entity changes in this phase — verification only. Summary tree nodes for audit/graph_rules.go, web/src/pages/Graph.tsx, skill/SKILL.md, skill/references/entity-spec.md, CLAUDE.md re-summarized.
      details: 'Expected touched files: internal/audit/graph_rules.go, web/src/pages/Graph.tsx, skill/SKILL.md, skill/references/entity-spec.md, CLAUDE.md. Run bulk-verify across all active requirements to re-snapshot content hashes. Rebuild syde+syded, reinstall skill, restart daemon.'
      notes: This phase is the finish gate — if anything fails here, loop back to the phase whose task missed something.
      tasks:
        - build-reinstall-restart-daemon
        - bulk-re-verify-all-active-requirements-after-audit-rule-change
        - final-sync-check-and-tree-refresh
changes:
    requirements:
        extended:
            - id: nmq3
              slug: system-entities-shall-not-carry-belongs-to-6hg2
              what: 'New audit finding: system entity carrying belongs_to is a blocking finding'
              why: Systems are standalone top-level process entities; belongs_to creates false hierarchy
              tasks:
                - rewrite-hierarchyfindings-to-drop-root-system-rule-and-forbid-belongsto-on-systems
            - id: 803y
              slug: design-model-shall-not-designate-a-root-system-8joh
              what: Audit drops exactly-one-root-system rule; all systems are standalone
              why: Flat system set reflects actual process boundaries; root was an artifact
              tasks:
                - rewrite-hierarchyfindings-to-drop-root-system-rule-and-forbid-belongsto-on-systems
            - id: y4yy
              slug: component-shall-be-allowed-to-belong-to-multiple-systems-qd6u
              what: Audit accepts multiple belongs_to edges on a single component
              why: Shared libraries run in both binaries; single owner misrepresents their scope
              tasks:
                - generate-reparenting-worksheet-for-12-components-57-contracts
                - sample-and-validate-reparenting-assignments
                - bulk-apply-component-and-contract-reparenting-and-reindex
                - bulk-re-verify-all-active-requirements-after-audit-rule-change
            - id: hmex
              slug: dashboard-graph-shall-render-every-system-at-the-same-visual-tier-pb2g
              what: Graph.tsx drops system-root vs system-sub distinction; single visual tier
              why: With no root, the tier split has no semantic meaning; all systems are peers
              tasks:
                - simplify-graphtsx-to-render-every-system-at-one-tier
            - id: qexi
              slug: design-model-shall-contain-exactly-two-systems-named-syde-and-syded-f2q8
              what: After consolidation, system set is exactly syde and syded
              why: Three-system state had semantic overlap; collapsing to one system per binary removes ambiguity
              tasks:
                - rename-syded-dashboard-e82c-name-field-to-syded-and-remove-belongstosyde
                - delete-syde-cli-2478-system-entity
                - build-reinstall-restart-daemon
                - final-sync-check-and-tree-refresh
            - id: ad9z
              slug: skill-documentation-shall-describe-systems-as-standalone-processes-rsf8
              what: SKILL.md and entity-spec.md describe systems as top-level standalone processes with no belongs_to
              why: Docs are the behaviour lever for future agents; without updating them sessions keep authoring root-system model
              tasks:
                - rewrite-system-rules-in-skillmd
                - rewrite-system-section-in-entity-specmd
                - purge-root-system-language-from-claudemd
                - build-reinstall-to-bundle-the-two-new-reference-docs
            - id: n94c
              slug: plan-phases-shall-list-created-and-modified-files-before-task-breakdown-hhvm
              what: SKILL.md plan-authoring rules require a Files section per phase before task decomposition
              why: Locking decomposition in a Files section is the superpowers-derived practice for subagent-executable plans
              tasks:
                - draft-referencesplan-authoringmd
                - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
            - id: rgdm
              slug: plan-tasks-shall-decompose-work-into-bite-sized-checkbox-steps-jv3l
              what: SKILL.md requires bite-sized checkbox steps with exact commands and expected output
              why: Bite-sized steps make plans executable by fresh subagents without ambiguity
              tasks:
                - draft-referencesplan-authoringmd
                - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
            - id: fyvf
              slug: plan-tasks-shall-not-contain-placeholder-language-qx8b
              what: SKILL.md defines a no-placeholder blacklist enforced via plan-authoring reference doc
              why: Placeholder phrases produce plans that pass structural checks but fail in implementation
              tasks:
                - draft-referencesplan-authoringmd
            - id: x1db
              slug: plan-author-shall-run-self-review-checklist-ka3p
              what: SKILL.md Phase 2 mandates self-review pass before plan is presented for approval
              why: 'User feedback: drafted plans routinely ship with authoring gaps — mandatory self-review catches them'
              tasks:
                - draft-referencesplan-authoringmd
                - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
            - id: lpmk
              slug: syde-cli-shall-provide-plan-review-command-that-dispatches-a-plan-reviewer-subagent-6wli
              what: New syde plan review <slug> command returns Approved or Issues Found verdict
              why: Structural syde plan check does not catch spec-alignment issues; semantic reviewer closes that gap
              tasks:
                - implement-syde-plan-review-slug-cli-command
                - draft-referencesplan-review-promptmd
            - id: 84eb
              slug: plan-approval-step-shall-offer-execution-handoff-choice-mez3
              what: After plan approval, SKILL.md requires surfacing subagent-driven vs inline execution choice
              why: 'Superpowers pattern: explicit alignment checkpoint before burning context on implementation'
              tasks:
                - draft-referencesplan-authoringmd
                - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
    systems:
        deleted:
            - id: 4jzu
              slug: syde-cli-2478
              why: Semantic duplicate of syde-5tdt (both represent the syde CLI process). Collapses 3-system state into 2. All 12 components + 57 contracts get reparented to syde-5tdt and/or syded-dashboard-e82c in Phase 3 before deletion.
              tasks:
                - delete-syde-cli-2478-system-entity
        extended:
            - id: 90ib
              slug: syded-dashboard-e82c
              what: Rename name field from 'syded Dashboard' to 'syded' and remove the belongs_to:syde relationship.
              why: Matches the syded binary name; belongs_to removal makes it a standalone top-level system.
              tasks:
                - rename-syded-dashboard-e82c-name-field-to-syded-and-remove-belongstosyde
    components:
        extended:
            - id: xt4a
              slug: audit-engine-4ktg
              what: 'Replace root-selection logic in hierarchyFindings with flat rules: systems must not carry belongs_to; non-system non-requirement entities must carry ≥1 belongs_to (can be multiple).'
              why: Implements requirements for no-root-system and multi-system belongs_to on components.
              tasks:
                - rewrite-hierarchyfindings-to-drop-root-system-rule-and-forbid-belongsto-on-systems
            - id: l2f9
              slug: web-spa-jy9z
              what: Graph.tsx drops subSystemIds calculation; all system entities render with sizeKey='system' and a single legend entry.
              why: Flat visual tier matches the flat model topology.
              tasks:
                - simplify-graphtsx-to-render-every-system-at-one-tier
            - id: ykiu
              slug: skill-installer-wbmu
              what: Bundle updated skill/SKILL.md (system rules rewrite), skill/references/entity-spec.md (System section rewrite), and CLAUDE.md notes referencing the new standalone-systems model.
              why: Skill embeds drive agent behaviour in new projects; docs must reflect the new rules.
              tasks:
                - rewrite-system-rules-in-skillmd
                - rewrite-system-section-in-entity-specmd
                - purge-root-system-language-from-claudemd
                - draft-referencesplan-authoringmd
                - rewrite-skillmd-phase-2-create-plan-to-link-to-plan-authoringmd
                - implement-syde-plan-review-slug-cli-command
                - draft-referencesplan-review-promptmd
                - build-reinstall-to-bundle-the-two-new-reference-docs
            - id: spjk
              slug: cli-commands-hpjb
              what: Add 'syde plan review <slug>' subcommand that loads the plan-reviewer prompt from the embedded skill bundle, dispatches it to a subagent (or prints the prompt if no subagent backend is available), and renders Approved or Issues Found verdict.
              why: Implements the new plan-review requirement; complements the structural syde plan check with a semantic reviewer.
              tasks:
                - implement-syde-plan-review-slug-cli-command
                - draft-referencesplan-review-promptmd
---
