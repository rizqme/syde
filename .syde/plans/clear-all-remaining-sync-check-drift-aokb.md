---
id: PLN-0010
kind: plan
name: Clear all remaining sync check drift
slug: clear-all-remaining-sync-check-drift-aokb
relationships:
    - target: approved-plan-clear-all-remaining-sync-check-drift-phkm
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T09:16:52Z"
plan_status: completed
background: 'After completing the Concept entity redesign plan with --force, 49 sync check errors remain in the model. They fall into four groups: (1) 29 orphan files under web/ and scripts/ that no component claims, (2) 16 requirements declared by three older plans (Flow steps with contract references and flowchart rendering, Requirement overlap audit with mandatory acknowledgement, Clear all sync check findings and enforce zero-finding completion) but never authored, (3) one flow the Flow-steps plan declared for deletion (design-model-operations-coverage) but left on disk, and (4) three approved-plan requirements over the 10-per-kind link cap — 20 tasks and 30 flows and 19 tasks respectively. Each error is genuine drift between plan intent and entity state; none are audit bugs.'
objective: syde sync check --strict exits 0 from a clean session. syde plan complete no longer needs --force on the three affected older plans. Every tracked source file is either claimed by a component or explicitly tree-ignored. Every plan-declared new requirement exists as an EARS-validated entity. Every plan-declared deletion is executed. No approved-plan requirement exceeds the 10-per-kind link cap.
scope: 'In scope: map every orphan web/ and scripts/ file onto the existing Web SPA component''s file list; author the 16 missing requirements from each older plan''s draft fields (statement, type, priority, verification, rationale); delete the design-model-operations-coverage flow; split three over-linked approved-plan requirements into kind-scoped children with derives_from links and retarget every linking entity. Out of scope: changes to audit rules or the 10-per-kind cap; new features; modifications to already-completed plans'' designs; any refactor of web SPA structure.'
design: 'The cleanup proceeds by error category rather than by plan so each phase closes a whole class of findings in one pass. Phase 1 extends web-spa''s --file list to enumerate every web/** source and config file plus scripts/wireframe-shot.sh; we pick one canonical order (folders alphabetical, files alphabetical within each folder) so diffs stay readable. Phase 2 authors the 16 missing requirements in three batch scripts (one per source plan); each requirement uses the exact EARS statement/type/priority/rationale/verification from the plan''s draft map, with --add-rel <target>:refines pointing at the implementing component or contract called out in the plan''s change entry. Phase 3 deletes the design-model-operations-coverage flow via syde remove --force. Phase 4 splits the three over-linked approved-plan requirements: for each, author N kind-scoped child requirements (one per ≤10-entity batch), each deriving from the parent; repoint every linking task/flow from the parent onto its assigned child via --remove-rel parent-slug followed by --add-rel child-slug:references; leave the parent in place with requirement_status=active as the umbrella node. Phase 5 runs reindex + syde sync check --strict and asserts exit 0. Throughout: every syde update uses the two-step remove-then-add pattern for relationship edits so the known clobbering bug (add-rel then remove-rel removes both) does not bite us.'
source: manual
created_at: "2026-04-17T08:47:14Z"
approved_at: "2026-04-17T09:04:08Z"
completed_at: "2026-04-17T09:16:52Z"
phases:
    - id: phase_1
      name: Map orphan files to web-spa
      status: completed
      description: Add every web/ and scripts/wireframe-shot.sh path to the Web SPA component's --file list
      objective: Zero tree/* orphan errors under sync check
      changes: components/web-spa-jy9z.md gains a complete --file enumeration of web/** + scripts/wireframe-shot.sh
      details: 'Run a single syde update web-spa --file <path1> --file <path2> ... passing every orphan path reported by syde sync check plus the existing files already on the list. The --file flag replaces the list so we must pass the full union. Path order: existing web/ files first, then new additions alphabetically grouped by folder.'
      tasks:
        - enumerate-web-and-scripts-files-on-web-spa
    - id: phase_2
      name: Author missing requirements from older plans
      status: completed
      description: Create the 16 new-change requirements that the Flow steps, Requirement overlap audit, and Clear all sync check plans declared but never authored
      objective: Zero 'plan claims to create new requirement X but no entity with that name exists' errors
      changes: 16 new requirement entities under .syde/requirements/
      details: Three /tmp/syde-reqs-phase2-<plan>.sh batch scripts, one per source plan. Each syde add requirement call sources statement/type/priority/rationale/verification verbatim from the plan's --draft map; --add-rel targets the implementing component or contract the plan change entry called out. Requirements that would trigger TF-IDF overlap warnings get --audited <slug> on subsequent calls.
      tasks:
        - author-9-flow-steps-plan-requirements
        - author-6-requirement-overlap-plan-requirements
        - author-1-clear-all-sync-check-plan-requirement
    - id: phase_3
      name: Execute declared flow deletion
      status: completed
      description: Delete the design-model-operations-coverage flow the Flow-steps plan claimed to delete
      objective: Zero 'plan claims to delete flow X but it still exists' errors
      changes: flows/design-model-operations-coverage-*.md removed from disk and index
      details: syde remove <slug> --force. If the flow has inbound references from steps on other flows, break those first by editing the referring flows; expect this to be a clean delete since the declared intent was straightforward removal.
      tasks:
        - delete-design-model-operations-coverage-flow
    - id: phase_4
      name: Split over-linked approved-plan requirements
      status: completed
      description: For each requirement exceeding the 10-per-kind link cap, author kind-scoped children and repoint linking entities
      objective: Zero 'requirement is linked by N tasks/flows (cap is 10 per kind) — split into kind-specific requirements' errors
      changes: ~7 new child requirements (2 task-scope for Clear-all, 3 flow-scope + 2 task-scope for Flow-steps), each deriving from its parent. Dozens of tasks/flows --remove-rel parent --add-rel child:references
      details: 'Per parent: (a) enumerate linking entities from syde query <parent> --full inbound list; (b) group into batches of <=10 by kind; (c) for each batch, author a child requirement with a scoped name (''Clear all sync check: concept redesign tasks'' / ''Flow steps: plan lifecycle flows'' etc.), --add-rel <parent>:derives_from, --add-rel <implementing-component>:refines; (d) repoint each entity in the batch via syde update <slug> --remove-rel <parent> then --add-rel <child>:references. The two-step pattern is mandatory because of the known add-then-remove clobber bug.'
      tasks:
        - split-clear-all-sync-check-over-linked-requirement
        - split-flow-steps-plan-over-linked-requirement
    - id: phase_5
      name: Verify clean gate
      status: completed
      description: Reindex, refresh summary tree, confirm syde sync check --strict exits 0
      objective: syde sync check --strict exits 0 and the three older plans can syde plan complete without --force
      changes: No entity changes; reindex run and tree scan run
      details: Run syde reindex, syde tree scan, iterate summarize loop for any stale nodes, then syde sync check --strict; assert exit 0. Finally attempt syde plan complete on each of the three older plans without --force and confirm success.
      tasks:
        - reindex-and-assert-clean-sync-check
changes:
    requirements:
        new:
            - id: w2l1
              name: Requirement creation shall detect similar requirements
              what: Authored from source plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Overlap detection prevents redundant intent statements.
                req_type: functional
                source: plan
                statement: When syde add requirement is invoked, the CLI shall compute TF-IDF similarity against existing active requirements and warn the user for any match above 50%.
                verification: syde add requirement prints overlap warnings for similar names
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: hb77
              name: Requirement creation shall accept audited flag
              what: Authored from source plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Authors need a way to proceed when overlap is intentional.
                req_type: functional
                source: plan
                statement: The syde add requirement command shall accept a repeatable --audited flag that acknowledges a specific overlap by slug.
                verification: --audited <slug> consumed by syde add requirement; acknowledged overlaps not re-warned
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: swtd
              name: Requirement update shall support audited flag
              what: Authored from source plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Audit acknowledgement must be possible post-creation.
                req_type: functional
                source: plan
                statement: The syde update command shall accept a repeatable --audited flag to add acknowledged overlaps to an existing requirement.
                verification: syde update --audited <slug> persists on the requirement
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: p2ik
              name: Overlap detection shall be bidirectional
              what: Authored from source plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Mutual acknowledgement is stronger than unilateral ignore.
                req_type: functional
                source: plan
                statement: The syde sync check overlap rule shall require both requirements in an overlap pair to acknowledge each other before clearing the finding.
                verification: sync check still errors if only one side of the pair acknowledges the other
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: zdxx
              name: Requirement entity shall carry audited overlaps list
              what: Authored from source plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Audit state must round-trip through YAML.
                req_type: functional
                source: plan
                statement: The syde requirement entity shall persist a list of acknowledged overlapping requirement slugs.
                verification: RequirementEntity has AuditedOverlaps field serialised to YAML
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: qige
              name: 'Clear all sync check: concept redesign tasks'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: tcua
              name: 'Clear all sync check: audit and overlap tasks'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: ryhb
              name: 'Flow steps: plan lifecycle flows'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: 17j3
              name: 'Flow steps: entity operation flows'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: ji5j
              name: 'Flow steps: dashboard browsing flows'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: qiif
              name: 'Flow steps: flow authoring tasks'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: p386
              name: 'Flow steps: chart and doc tasks'
              what: Kind-scoped split child under its parent approved-plan requirement
              why: Parent exceeds 10-per-kind link cap; this scopes one batch of <=10 entities
              draft:
                priority: must
                rationale: Cap compliance plus preserved traceability
                req_type: constraint
                source: plan
                statement: The syde design model shall trace the scoped batch of entities covered by this requirement back to its parent approved-plan requirement via derives_from.
                verification: sync check reports no cap violations and each entity in the batch references this requirement instead of its parent
              tasks:
                - split-clear-all-sync-check-over-linked-requirement
                - split-flow-steps-plan-over-linked-requirement
            - id: n79b
              name: Sync check shall error on unaudited requirement overlaps
              what: Authored from Requirement-overlap plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Mandatory acknowledgement forces the author to confirm or resolve overlaps.
                req_type: constraint
                source: plan
                statement: If two active requirements overlap above 50% similarity without mutual acknowledgement, then the syde sync check engine shall report an error.
                verification: sync check errors when two reqs overlap and neither acknowledges the other
              tasks:
                - author-6-requirement-overlap-plan-requirements
            - id: vqil
              name: Plan complete shall require clean sync check
              what: Authored from Clear-all plan draft
              why: Plan declared; never created
              draft:
                priority: must
                rationale: Plan completion is the canonical gate; forcing should be rare and intentional.
                req_type: constraint
                source: plan
                statement: If the syde sync check reports any errors, then the syde plan complete command shall refuse to mark the plan completed without --force.
                verification: syde plan complete exits non-zero when sync check has errors unless --force is passed
              tasks:
                - author-1-clear-all-sync-check-plan-requirement
    components:
        extended:
            - id: ud1o
              slug: web-spa
              what: Enumerate every web/** source/config plus scripts/wireframe-shot.sh on the --file list
              why: 29 orphan tree/ errors — every tracked web/ file must be claimed by a component
              tasks:
                - enumerate-web-and-scripts-files-on-web-spa
    flows:
        deleted:
            - id: f4b4
              slug: design-model-operations-coverage
              why: Declared for deletion by the Flow-steps plan but never removed
              tasks:
                - delete-design-model-operations-coverage-flow
---
