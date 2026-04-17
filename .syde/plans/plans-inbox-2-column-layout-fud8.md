---
id: PLN-0002
kind: plan
name: Plans Inbox 2-Column Layout
slug: plans-inbox-2-column-layout-fud8
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-17T11:08:10Z"
plan_status: completed
background: The Plans dashboard currently uses a custom PlansInboxScreen + PlanDetailScreen at /__plan__/<slug>, with the floating EntityDetail panel as a fallback. Every other entity kind uses the canonical 2-column inbox (420px left list + flex-1 right detail) handled by App.tsx. The result is inconsistent UX and the right-side floating panel is dead weight. The user wants plans to follow the same inbox pattern as components/contracts/concepts/systems/flows.
objective: Plans render via the default 2-column inbox at /<proj>/plan with the entity list on the left and a PlanDetailPanel on the right (in-page, not floating). Selecting a plan from the list shows Design + Changes + Tasks tabs in the right column. The floating panel is gone.
scope: 'Phase 1: switch routing — drop plan from SPECIAL_VIEWS, wire EntityList + PlanDetailPanel via the default 2-column path. Phase 2: build PlanDetailPanel by refactoring PlanDetailScreen for inline use. Phase 3: empty-state for PlanChangesView. Phase 4: delete dead PlansInboxScreen + PlanDetailScreen files. Phase 5: verify (build, browser smoke). EXCLUDED: any backend changes, the API endpoint, the audit rules, the CLI commands — all stay as-is. Only the React routing layer and the plan detail component shape change.'
design: |-
    Today, App.tsx switches between two render paths based on whether activeKind is in SPECIAL_VIEWS = ['plan', 'task', 'design']: special kinds get their own custom page component (PlansInboxScreen → PlanDetailScreen), other kinds get the canonical EntityList + EntityDetail 2-column layout. Plan currently lives in SPECIAL_VIEWS because the plan detail UI is structurally different from a generic entity (tabs, phases, structured changes view).

    The fix: drop plan from SPECIAL_VIEWS, route it through the same 2-column path everything else uses, and refactor PlanDetailScreen into a PlanDetailPanel component that App.tsx renders in the right column instead of EntityDetail when activeKind === 'plan'. The left column is the existing EntityList (which already lists plans correctly via api.entities('plan')).

    PlanDetailPanel preserves everything PlanDetailScreen had — Design / Background / Objective / Scope / Changes / Phases / Tasks tabs — but its props match the inline pattern (slug, onClose, onNavigate, no router). Tab state still persists via ?tab=plan|tasks query param.

    PlansInboxScreen and PlanDetailScreen become dead code and get deleted. The /__plan__/<slug> nested route also goes — selection flows through setSelectedSlug() like every other inbox kind.

    Bonus fix: PlanChangesView currently hides empty lanes and renders nothing when all six lanes are empty. Add an explicit empty state ('No changes declared yet — use syde plan add-change ...') so reviewers know it's deliberate rather than broken.
source: manual
created_at: "2026-04-15T13:02:32Z"
approved_at: "2026-04-15T14:56:31Z"
completed_at: "2026-04-16T05:08:20Z"
phases:
    - id: phase_1
      name: Switch plan routing to default 2-column inbox
      status: completed
      description: Drop plan from SPECIAL_VIEWS in App.tsx and route it through the same EntityList + right-column-detail path components/contracts/concepts/systems/flows already use.
      objective: Opening /<proj>/plan renders the canonical 2-column inbox with the entity list on the left.
      changes: web/src/App.tsx (remove plan from SPECIAL_VIEWS, drop the PlansInboxScreen branch, add the plan-specific right-column rendering).
      details: 'App.tsx already has a 2-column branch for normal kinds. Add a conditional inside the right column: if activeKind === ''plan'' and selectedSlug, render <PlanDetailPanel slug={selectedSlug} ... /> instead of <EntityDetail .../>. Remove the SPECIAL_VIEWS reference for ''plan'' so the floating-panel branch never triggers for plans.'
      tasks:
        - drop-plan-from-specialviews-in-apptsx
        - wire-plandetailpanel-into-the-right-column-for-plan-kind
        - add-tasks-string-field-to-change-types
    - id: phase_2
      name: Build PlanDetailPanel from PlanDetailScreen
      status: completed
      description: Refactor PlanDetailScreen into a PlanDetailPanel component that renders inline in the right column. Same content, no routing, prop-driven.
      objective: PlanDetailPanel takes (slug, onClose, onNavigate, onOpenFile) props and renders the full plan content (header + Plan/Tasks tabs).
      changes: web/src/components/PlanDetailPanel.tsx (new), web/src/lib/api.ts (no changes — api.planDetail already exists).
      details: 'Copy the body of PlanDetailScreen.tsx into a new component file under components/. Replace useParams/useNavigate with props. Tab state still persists via window.location query param ?tab=plan|tasks (or hoist to App.tsx if simpler). Header layout matches the existing inline EntityDetail style: name, status pill, progress bar, approved_at timestamp.'
      tasks:
        - create-plandetailpanel-component
        - pre-render-wireframe-html-for-screen-contracts-in-plan-detail-api
        - render-wireframe-html-in-planchangesview-extended-and-new-cards
        - tweak-phasetasklist-phase-header-rendering
        - add-navigate-websocket-message-type-to-syded
        - spa-listens-for-navigate-websocket-messages
        - implement-syde-plan-open-cli-command
        - group-extended-changes-by-target-slug-in-planchangesview
        - pre-render-proposed-wireframe-html-for-screen-contract-fieldchanges
        - render-wireframe-fieldchanges-as-tabbed-old-or-new-view
        - add-task-flag-to-syde-plan-add-change-subcommands
    - id: phase_3
      name: Empty-state for PlanChangesView
      status: completed
      description: PlanChangesView currently hides empty lanes and renders nothing when all six are empty. Show an explicit empty state instead.
      objective: Opening a plan with zero declared changes shows a muted 'No changes declared yet' message with a hint about syde plan add-change.
      changes: web/src/components/PlanChangesView.tsx (add empty-state branch when every lane has zero entries).
      details: 'After computing the visible lanes list, if the total entry count is zero render a centered card: ''No changes declared yet. Use syde plan add-change to record deletions, extensions, or new entities this plan will make.'' Use the same muted-foreground / border-dashed styling as the existing EntityEmptyState.'
      tasks:
        - add-empty-state-branch-to-planchangesview
    - id: phase_4
      name: Delete dead PlansInboxScreen and PlanDetailScreen
      status: completed
      description: After plans route through the default path, the dedicated screen files are dead code.
      objective: Both files are deleted; nothing in web/src/ imports them; the build still passes.
      changes: Delete web/src/pages/PlansInboxScreen.tsx and web/src/pages/PlanDetailScreen.tsx; remove their imports from App.tsx; verify rg PlansInboxScreen|PlanDetailScreen returns no matches in web/src.
      details: Delete the two files; rg-grep for any lingering imports; clean up.
      tasks:
        - delete-dead-plansinboxscreen-and-plandetailscreen-files
        - update-plan-view-screen-contract-wireframe-and-files-list
    - id: phase_5
      name: Verify
      status: completed
      description: Build, install, browser smoke.
      objective: bun run build clean, dashboard renders the new layout, syde plan complete passes the validator.
      changes: No source changes; test, build, install, manual browser inspection.
      details: cd web && bun run build. make install. Open the dashboard, navigate to /<proj>/plan, click the Plans Inbox 2-Column Layout plan, verify the Plan tab shows the populated Changes block and the Tasks tab shows the phases. Run syde plan complete after marking all tasks done — validator should pass.
      tasks:
        - frontend-build-and-browser-smoke
        - run-plan-complete-validator
        - update-plan-lifecycle-flow-narrative-and-happypath
    - id: phase_6
      name: Plan authoring completeness harness
      status: completed
      description: Add a planAuthoringFindings audit rule that flags gaps in draft plans before approval, plus a syde plan check CLI command, plus a mandatory skill rule that agents must run the check after drafting and revise until clean.
      objective: Drafted plans surface their own gaps via syde plan check; the skill workflow requires running that check after every drafting session.
      changes: internal/audit/plan_authoring.go (new), internal/audit/audit.go (wire into Run + add CatPlanAuthoring), internal/cli/plan.go (new check subcommand), skill/SKILL.md + skill/codex/SKILL.md (new mandatory step in the plan workflow), installed skill copies.
      details: 'planAuthoringFindings runs on draft/approved plans (skip completed) and emits:\n- WARN: Requirements lane is empty (plan declares no durable property).\n- WARN: ExtendedChange has no field_changes (programmatic verification disabled).\n- ERROR: NewChange draft missing kind-required fields (component.responsibility, contract.contract_kind/input/output, requirement.statement matching EARS, concept.meaning/invariants).\n- WARN: Orphan change — declared change whose target slug appears in no task''s affected_entities/affected_files.\n- WARN: Orphan task — task whose affected_entities/affected_files don''t match any declared change target.\n- WARN: Extended change targets a screen contract but field_changes doesn''t include wireframe.\n\nsyde plan check <plan> runs syde sync check, filters findings to plan_authoring + plan_completion categories scoped to the target plan, and prints them grouped by severity. Used as the post-drafting gate in the skill workflow.\n\nSkill docs additions: in SKILL.md Phase 2 CREATE PLAN and codex SKILL.md plan workflow, after the ''add tasks'' step, add: ''Run syde plan check <plan>. Revise the plan to address every WARN and ERROR. Do NOT present the plan for approval until syde plan check exits 0 (errors only — warnings may be acknowledged in the presentation).'''
      tasks:
        - implement-planauthoringfindings-audit-rule
        - wire-planauthoringfindings-into-auditrun
        - implement-syde-plan-check-cli-command
        - update-skill-docs-to-require-syde-plan-check-after-drafting
        - refresh-installed-skill-copies-after-phase-6
        - run-syde-plan-check-on-this-plan-and-address-findings
changes:
    requirements:
        new:
            - id: 1iz1
              name: Plans shall render via the canonical 2-column inbox
              what: New EARS requirement declaring that the syded dashboard renders plans via the same 2-column inbox layout used by all other entity kinds, with no floating detail panel for plan kind.
              why: The user explicitly called this out — plans should follow the same UX pattern as components/contracts/concepts/systems/flows. Capturing it as a requirement makes the rule durable and audit-checkable, not just a one-off refactor.
              draft:
                priority: must
                rationale: Plans need the same inbox UX as every other entity kind so reviewers do not learn a separate navigation pattern, and the floating EntityDetail panel becomes dead weight.
                req_type: usability
                statement: The syded dashboard shall render the Plans inbox as a 2-column layout with the entity list on the left and the selected plan's detail rendered inline in the right column.
                verification: 'Manual inspection of /<project>/plan in the dashboard: a plan list is on the left, selecting a plan renders the detail inline on the right, no floating panel appears.'
            - id: po2x
              name: Plan changes view shall render screen wireframes inline
              what: New EARS requirement declaring that screen-kind contracts referenced by a plan's Extended or New changes have their wireframe rendered as HTML inline in the plan changes view, not as a UIML source dump.
              why: Screen contracts already render as wireframes in EntityDetail. The plan changes view should match — reviewers comparing plan diffs need to see what the screen will look like, not parse UIML source by eye.
              draft:
                priority: must
                rationale: UIML source is unreadable to humans during review; rendered wireframes are the whole point of having a wireframe field.
                req_type: usability
                statement: When the plan changes view renders an Extended or New change targeting a screen-kind contract, the syded dashboard shall render the contract's wireframe as HTML using the canonical UIML renderer.
                verification: 'Manual inspection: open a plan whose changes block extends a screen contract, confirm the wireframe renders visually inside the Extended card.'
            - id: finm
              name: Plans shall pass syde plan check before approval
              what: New EARS requirement declaring that any plan presented for approval must first pass syde plan check (zero ERROR findings).
              why: Captures the new gate as a durable system property the audit can later enforce.
              draft:
                priority: must
                rationale: Plan authoring gaps are the dominant failure mode (forgotten requirements, missing field_changes, thin NewChange drafts). A programmatic gate catches them before approval instead of after execution drift.
                req_type: constraint
                statement: When the syde planning workflow drafts a plan, the agent shall run syde plan check on the plan and address every ERROR before presenting the plan for approval.
                verification: 'Manual workflow inspection: every approved plan in the repo must have been preceded by a syde plan check that exited 0.'
            - id: ie37
              name: Plan approvals shall be preceded by a dashboard open
              what: New EARS requirement declaring that whenever the syde planning workflow asks the user to approve a plan, the agent first runs syde plan open which reuses the existing dashboard tab via WebSocket broadcast or spawns a new tab if none is connected.
              why: Reviewers should never have to manually navigate to the plan they're being asked to approve. The browser-open step is part of the approval handoff.
              draft:
                priority: must
                rationale: Manual URL navigation breaks reviewer flow. The dashboard is already open in a tab during a planning session; the agent should drive that tab directly.
                req_type: usability
                statement: When the syde planning workflow drafts a plan and is ready to ask for user approval, the agent shall run syde plan open <plan> before presenting the approval prompt.
                verification: 'Manual workflow inspection: every approval request in a session is preceded by a syde plan open invocation.'
            - id: gcj8
              name: Plan changes view shall group Extended by target
              what: 'New EARS requirement: when a plan declares multiple Extended changes targeting the same entity, the dashboard shall render them as a single card grouped by target slug.'
              why: Per-entry rendering creates duplicate cards for the same target, fragmenting the reviewer's mental model of what's changing on each entity.
              draft:
                priority: must
                rationale: Reviewers think in terms of 'what is changing on entity X', not 'what change entries exist'. The card boundary should match the entity boundary.
                req_type: usability
                statement: When the plan changes view renders multiple Extended changes targeting the same entity slug, the syded dashboard shall group them into a single card with all what/why/field_changes stacked inside.
                verification: 'Manual inspection: open a plan with two Extended changes on the same target and confirm only one card renders for that target.'
            - id: mcm1
              name: User requests shall cascade requirement-first across all lanes
              what: New EARS requirement declaring that any user request the agent acts on must first be captured as a New or Extended Requirement (with conflicting existing requirements marked superseded or obsolete) before being expressed as Component/Contract/Concept/Flow changes, and that any flow whose behavior changes must also be Extended with field_changes.
              why: Requirements are the why; components/contracts/concepts/flows are the how. Inverting the order means the plan tells you what is happening but not why it is allowed to happen, conflicting requirements rot in place, and behavior flows silently drift away from reality.
              draft:
                priority: must
                rationale: Requirements are the why, the other lanes are the how. Flows in particular rot the fastest when authors forget to update them — narrative drifts, happy_path no longer matches the code, edge_cases get stale. The cascade rule forces explicit acknowledgement.
                req_type: constraint
                statement: When the user requests a change to the syde system, the agent shall first capture the underlying requirement as a New or Extended Requirement entry in the plan's Requirements lane (marking any conflicting existing requirement superseded or obsolete) before declaring any Component, Contract, Concept, or Flow changes that implement it; and any flow whose narrative, happy_path, or edge_cases changes shall be Extended with corresponding field_changes.
                verification: 'Manual workflow inspection: every plan in the repo must show that its Requirements lane was populated before its Components/Contracts/Concepts/Flows lanes for any given user-driven change cluster, and any flow whose behavior changes must appear in the Flows lane Extended list.'
            - id: a0da
              name: Plan changes shall list their implementing tasks
              what: New EARS requirement declaring that every Deleted/Extended/NewChange entry in the plan's Changes block carries an explicit tasks []string field listing the task slugs that implement it. The planAuthoringFindings audit shall ERROR when any change has an empty tasks list or references a task slug that does not exist.
              why: Heuristic matching (affected_entities, name fuzzy match) is brittle and surprises authors. An explicit one-way mapping FROM each change TO its implementing task slugs is unambiguous and trivially auditable.
              draft:
                priority: must
                rationale: A change without a claiming task is dead intent. Heuristic matching forces guesswork; explicit mapping makes the linkage canonical and trivially auditable.
                req_type: constraint
                statement: When the syde plan model declares a Deleted, Extended, or NewChange entry, the entry shall carry a non-empty tasks []string field listing the slugs of one or more tasks in the same plan that implement the change, and the syde plan authoring audit shall ERROR on any change with an empty or unresolvable tasks list.
                verification: syde plan check on a plan whose changes contain at least one entry with empty tasks reports the orphan change as ERROR and exits non-zero.
    components:
        extended:
            - id: b5zn
              slug: websocket-server-hdup
              what: Add project-scoped dashboard navigation broadcast support.
              why: Tab reuse needs server-side push so the SPA can switch URLs without spawning a fresh tab.
              field_changes:
                responsibility: Broadcast entity/tree change events over WebSocket
            - id: roxj
              slug: web-spa-jy9z
              what: Refactor plan detail rendering into the inline PlanDetailPanel and remove dead standalone plan pages.
              why: Plans should follow the same inbox UX as components, contracts, concepts, systems, and flows.
              field_changes:
                responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
            - id: ltc1
              slug: web-spa-jy9z
              what: Switch App.tsx to render plan kind through the default 2-column path.
              why: Plan should be a first-class inbox kind instead of a special floating-panel view.
              field_changes:
                responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
            - id: pre1
              slug: web-spa-jy9z
              what: Tweak PhaseTaskList phase headers with a distinct phase icon and Phase N prefix.
              why: Phases and tasks need visually distinct hierarchy in the Tasks tab.
              field_changes:
                responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
            - id: 68ov
              slug: web-spa-jy9z
              what: Group Extended changes by target and render wireframe changes as rendered Current/Proposed tabs.
              why: Reviewers need one card per changed entity and visual wireframe diffs instead of UIML source dumps.
              field_changes:
                responsibility: React 18 + TypeScript + Vite + Tailwind v4 single-page app rendered by syded
            - id: vsa3
              slug: http-api-afos
              what: Pre-render wireframe HTML in the plan detail API for screen contract changes.
              why: The dashboard plan changes view needs rendered wireframe previews in one round trip.
              field_changes:
                responsibility: Handle HTTP routes for entities, tree, search, constraints, context
            - id: nlbj
              slug: http-api-afos
              what: Add a project-scoped navigate endpoint for dashboard tab reuse.
              why: The CLI needs an HTTP entry point to trigger WebSocket navigation broadcasts.
              field_changes:
                responsibility: Handle HTTP routes for entities, tree, search, constraints, context
            - id: 2m7g
              slug: http-api-afos
              what: Pre-render proposed wireframe HTML for screen contract field changes.
              why: The frontend tab switcher needs current and proposed rendered wireframes in the plan detail response.
              field_changes:
                responsibility: Handle HTTP routes for entities, tree, search, constraints, context
        new:
            - id: lkzy
              name: Plan Detail Panel
              what: Inline React component rendering the Plan + Tasks tabs for a selected plan in the right column of the Plans inbox.
              why: Replaces PlanDetailScreen which was a routed page; needs to be a component that App.tsx can render inline.
              draft:
                boundaries: Does NOT own routing (App.tsx does). Does NOT fetch plan list (EntityList does).
                capabilities:
                    - Render plan header (name, status, progress, approved_at)
                    - Switch between Plan and Tasks tabs via ?tab= query param
                    - Embed PlanChangesView for the structured diff
                    - Embed PhaseTaskList for the Tasks tab
                purpose: Match the canonical 2-column inbox UX for plans.
                responsibility: Render a selected plan's design, structured changes, phases, and nested tasks inside the right column of the Plans inbox.
    contracts:
        extended:
            - id: "0e07"
              slug: plan-view-screen-gb2y
              what: 'Update the Plan View Screen wireframe to depict the 2-column inbox layout (left column: plans list with name/status/progress, right column: selected plan''s Design + Changes + Tasks tabs) instead of the old single-column PlanView mockup.'
              why: The contract's wireframe is the canonical visual spec for the screen; if the layout changes the wireframe must change too. Otherwise reviewers see the old mockup when looking at the contract.
              field_changes:
                wireframe: '<screen name="Plans Inbox" direction="horizontal"><sidebar><list name="plans"><item><label>Plans Inbox 2-Column Layout</label><label>draft</label></item><item><label>Revamp Planning</label><label>completed</label></item></list></sidebar><main name="Plan Detail"><navbar><heading>Plan: Plans Inbox 2-Column Layout</heading><button>Approve</button></navbar><tabs><tab name="Plan"/><tab name="Tasks"/></tabs><section><heading>Design</heading><paragraph>Detailed implementation prose...</paragraph></section><section><heading>Changes</heading><tabs><tab name="Requirements"/><tab name="Components"/><tab name="Contracts"/></tabs><card><label>Extended audit-engine</label></card></section></main></screen>'
            - id: pbtm
              slug: plan-view-screen-gb2y
              what: Switch the screen contract's files list to PlanDetailPanel.tsx (the inline component) instead of the deleted PlanDetailScreen.tsx page.
              why: The screen contract should point at the file that actually renders the plan detail UI.
              field_changes:
                files: web/src/components/PlanDetailPanel.tsx
    flows:
        extended:
            - id: yt4b
              slug: plan-lifecycle-pwb1
              what: 'Update Plan Lifecycle flow to reflect the new authoring loop: Requirements-first cascade, structured Changes block authoring (delete/extend/new), syde plan check gate after drafting, syde plan open before approval, syde plan complete with validator gate at finish.'
              why: The Plan Lifecycle flow's happy_path and narrative still describe the pre-revamp workflow with no Changes block, no plan check, no plan open, no plan complete validator. The cascade rule itself REQUIRES this flow to be updated in this plan.
              field_changes:
                happy_path: clarify → cascade requirement-first → draft plan with Design + Changes (Requirements/Components/Contracts/Concepts/Flows lanes) + phases + tasks → syde plan check (must exit 0) → syde plan open → user approves → execute task-by-task → syde plan complete (validator-gated) → finish with clean tree
                narrative: 'Phase 1: agent clarifies requirements critically and waits for user confirmation. Phase 2: agent identifies the underlying requirement (search existing requirements first; mark conflicting ones superseded or obsolete), then drafts a plan with background/objective/scope/design, populates the structured Changes block lane-by-lane (Requirements first, then Components/Contracts/Concepts/Flows), adds phases with objective/changes/details, adds tasks with objective/details/acceptance/affected-entities/affected-files. Phase 2.5: agent runs syde plan check, addresses every ERROR and reviews every WARN, then runs syde plan open to surface the plan in the user''s existing dashboard tab. Phase 3: agent presents plan with summary of caught gaps, user approves via syde plan approve. Phase 4: agent executes tasks one phase at a time without pausing for permission. Phase 5: agent runs syde plan complete which invokes planCompletionFindings to verify every declared change against actual entity state.'
---

-
