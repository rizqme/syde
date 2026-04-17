---
id: PLN-0001
kind: plan
name: Revamp Planning To Structured Design And Diff
slug: revamp-planning-to-structured-design-and-diff-m8p5
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
plan_status: completed
background: 'The current plan model captures background/objective/scope/phases/tasks but has no structured design section and no way to declare the specific entity changes a plan will make. Reviewers can''t see the diff before approval and the system can''t verify what actually landed afterwards. The old decision REQ-0007 ''Plans shall not carry draft or embedded entity definitions'' is being reversed: structured embedded drafts are exactly what makes plans reviewable and verifiable.'
objective: Extend the plan model with a Design prose field and a structured Changes block (per-kind Deleted/Extended/New lanes with what/why plus optional field-level diffs), add CLI commands to build the diff, render the whole plan in the dashboard with kind-specific contract viewers (screen/CLI/REST/storage/event), and add a completion validator that errors when the claimed diff does not match the actual entity state.
scope: 'Phase 1 model (Design field + Changes struct + per-lane/entry types + YAML round-trip + validation). Phase 2 CLI (--design flag, syde plan add-change / remove-change / show-changes with kind-specific subcommands). Phase 3 completion validator (planCompletionFindings rule, block complete on ERROR). Phase 4 dashboard plan viewer (route, sections per kind, specialized New-contract viewers per contract_kind, Extended side-by-side field diff). Phase 5 reverse REQ-0007 and refresh skill docs. Phase 6 verify. EXCLUDED: task↔change linkage (validator runs on all changes at plan completion regardless of task ownership).'
source: manual
created_at: "2026-04-15T11:39:14Z"
approved_at: "2026-04-15T11:45:29Z"
completed_at: "2026-04-15T12:46:06Z"
phases:
    - id: phase_1
      name: Plan model with design and changes
      status: completed
      description: Extend PlanEntity with a free-form Design prose field and a structured Changes block carrying per-kind lanes (Requirements, Systems, Concepts, Components, Contracts, Flows). Each lane has Deleted, Extended, and New lists. Extended entries may declare field-level diffs that the completion validator will check.
      objective: PlanEntity carries Design and Changes; YAML round-trips cleanly; validator rejects Changes entries that lack what/why.
      changes: internal/model/plan.go (new Design field + PlanChanges + ChangeLane + DeletedChange + ExtendedChange + NewChange types with YAML tags), internal/model/validation.go (require what and why on every change, validate kind-consistent NewChange drafts).
      details: 'PlanChanges groups per kind: Requirements ChangeLane, Systems ChangeLane, Concepts ChangeLane, Components ChangeLane, Contracts ChangeLane, Flows ChangeLane. ChangeLane contains Deleted []DeletedChange, Extended []ExtendedChange, New []NewChange. DeletedChange: Slug + Why. ExtendedChange: Slug + What + Why + optional FieldChanges map[string]string (value ''DELETE'' means drop field). NewChange: Name + What + Why + kind-specific Draft fields embedded inline (e.g. Responsibility/Capabilities/Boundaries for components, ContractKind/Input/Output/Wireframe for contracts, Statement/ReqType/Priority/Verification for requirements). Validation error if what or why is empty on any entry.'
      tasks:
        - add-planentity-design-field-and-changes-struct
        - validate-plan-changes-entries
        - plan-yaml-round-trip-smoke-test
    - id: phase_2
      name: CLI plan-change surface
      status: completed
      description: 'Add CLI commands to build the changes block: syde plan create --design, syde plan add-change, syde plan remove-change, syde plan show-changes.'
      objective: Users can fully author a plan with design and structured diff from the command line without hand-editing markdown.
      changes: internal/cli/plan.go (new subcommands add-change delete/extend/new, remove-change, show-changes; --design flag on plan create), internal/cli/add.go helpers reused for NewChange kind-specific fields.
      details: 'syde plan add-change <plan> requirement|system|concept|component|contract|flow delete <slug> --why. syde plan add-change ... extend <slug> --what --why [--field key=value, repeatable]. syde plan add-change ... new --name --what --why [kind-specific flags: --responsibility/--capability/--boundaries for component, --statement/--type/--priority/--verification for requirement, etc.]. syde plan remove-change <plan> <change-id>. syde plan show-changes <plan> --format rich|json. Each change gets a short ULID-style change-id on insertion so remove-change can address it.'
      tasks:
        - add-design-flag-to-syde-plan-create
        - syde-plan-add-change-delete-subcommand
        - syde-plan-add-change-extend-subcommand
        - syde-plan-add-change-new-subcommand
        - syde-plan-remove-change-and-show-changes
    - id: phase_3
      name: Plan completion validator
      status: completed
      description: Add an audit rule that compares a plan's Changes block against actual entity state and reports drift. Block plan completion on ERROR.
      objective: syde plan complete fails if any Deleted entity still exists, any New entity is missing, or any Extended entity has declared field_changes that do not match current values.
      changes: internal/audit/plan_completion.go (new planCompletionFindings function), internal/audit/audit.go (wire the rule into Run when plans are in scope), internal/cli/plan.go (plan complete blocks on ERROR from the validator).
      details: 'For each plan with PlanStatus=approved or completed: for each DeletedChange, ERROR if the target slug still resolves; for each NewChange, ERROR if no entity with the declared name/kind exists; for each ExtendedChange with FieldChanges, resolve the target entity and ERROR for any field whose current value does not match the declared new value (normalize string, slice, map comparisons). Soft case: ExtendedChange without FieldChanges is WARN only (hand-review only — there is nothing to verify mechanically). Completion validator is included in syde sync check --strict automatically.'
      tasks:
        - implement-plancompletionfindings-audit-rule
        - wire-plancompletionfindings-into-auditrun
        - block-syde-plan-complete-on-validator-errors
    - id: phase_4
      name: Dashboard plan viewer with kind-specific contract views
      status: completed
      description: 'Render the revamped plan in the browser: design prose, phases/tasks (existing), and a new Changes section grouped by kind with specialized views for new contract drafts (screen/CLI/REST/storage/event) and a side-by-side diff for Extended entries.'
      objective: A user opening /plans/<slug> in the dashboard sees the full plan design, diff, and specialized contract previews without leaving the page.
      changes: 'internal/dashboard/api.go + api_readall.go (new /api/plan/<slug> endpoint returning full plan with changes and pre-resolved Extended old-values), web/src/pages/PlanView.tsx (new Design + Changes section + kind tabs), web/src/components/PlanChanges.tsx (new component — groups Deleted/Extended/New per kind), web/src/components/NewContractDraftView.tsx (switches on contract_kind: screen -> UIML render, cli -> command+flags table, rest/rpc -> METHOD+params, storage -> key+fields, event/websocket -> name+payload), web/src/components/ExtendedFieldDiff.tsx (side-by-side before/after per declared field).'
      details: The API returns for each ExtendedChange both the current entity snapshot and the declared FieldChanges so the side-by-side diff can render without an extra round-trip. NewContractDraftView reuses the existing wireframe renderer for screen contracts and renders UIML via the shared uiml package. CLI command preview renders the Input line bolded with each Input Parameter as a flag row. REST/rpc contracts show METHOD + path + input_parameters as a request body/query table and output_parameters as a response body table. Storage contracts render the key pattern + fields. Event contracts render event name + payload schema. Component New drafts render responsibility/capabilities/boundaries bullets. Requirement New drafts render statement + req_type/priority/verification pills. Concept New drafts render meaning/invariants/attributes/actions.
      tasks:
        - expose-plan-with-changes-via-http-api
        - planview-page-renders-design-and-changes
        - newcontractdraftview-specialized-per-contractkind
        - extendedfielddiff-side-by-side-view
        - kind-specific-new-entity-draft-previews
        - frontend-build-and-browser-smoke
        - remove-standalone-tasks-sidebar-nav-and-taskboard-page
        - plans-inbox-column-page
        - plan-detail-page-shell-with-plan-and-tasks-tabs
        - tasks-tab-grouped-by-phases
    - id: phase_5
      name: Reverse REQ-0007 and refresh skill docs
      status: completed
      description: Mark the old 'Plans shall reference entities, never draft them' requirement as superseded and document the new plan workflow in the skill.
      objective: REQ-0007 is superseded_by a new requirement that declares plans MUST carry structured change diffs. Skill docs teach the new plan workflow.
      changes: .syde/requirements (new requirement authored + REQ-0007 status set to superseded + superseded_by link added), skill/SKILL.md, skill/codex/SKILL.md, skill/references/commands.md, skill/references/entity-spec.md, skill/references/sync-workflow.md, skill/references/requirement-derivation.md (add a Plan kind section if missing).
      details: 'New requirement statement: ''The syde plan model shall record every entity change a plan will make (deleted, extended, or new) with a structured what+why diff, and the completion validator shall verify the declared diff against actual entity state before a plan can be marked completed.'' Type: constraint. Priority: must. Verification: integration test building a plan with a New component draft, executing it, and asserting syde plan complete rejects mismatches. Skill docs get a new Plan workflow section: design prose, diff entries, phases, tasks, completion. After editing, make install + syde install-skill --all.'
      tasks:
        - author-new-plans-shall-carry-diff-requirement-and-supersede-req-0007
        - update-skill-docs-for-new-plan-workflow
        - refresh-installed-skill-copies-after-phase-5
    - id: phase_6
      name: Verify
      status: completed
      description: Run the full health gate and refresh the summary tree after the revamp lands.
      objective: go test, make install, syde sync check --strict, and syde tree status --strict all pass.
      changes: No source changes; test, build, install, and summary tree refresh.
      details: 'go test ./..., make install, syde install-skill --all, syde reindex, syde sync check --strict (expect clean modulo flow-traceability WARNs), syde tree scan and leaves-first summarize loop until syde tree status --strict exits 0. Plus an end-to-end smoke: create a plan with --design and --add-change new component/extend contract/delete decision, approve, author and run the phase tasks, then verify syde plan complete blocks on a deliberately mismatched field.'
      tasks:
        - end-to-end-plan-smoke-test
        - refresh-summary-tree-after-the-revamp
---
