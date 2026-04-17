---
id: PLN-0004
kind: plan
name: Flow steps with contract references and flowchart rendering
slug: flow-steps-with-contract-references-and-flowchart-rendering-m0b5
description: Structured flow steps with contract refs, step-based audit, flowchart dashboard, real user-journey flows
relationships:
    - target: approved-plan-flow-steps-with-contract-references-and-flowchart-rendering-7n8n
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T09:16:35Z"
plan_status: completed
background: Flows currently store behavior as free-text prose (narrative, happy_path, edge_cases). The contract-flow audit rule exists but is defeated by a catch-all 'Design Model Operations Coverage' flow that every contract points to. Flows need structured steps that explicitly reference contracts, a flowchart visualization in the dashboard, and audit rules that enforce meaningful coverage. The catch-all flow must be replaced with real user-journey flows.
objective: FlowEntity has structured steps referencing contracts. Each flow represents one user goal (create plan, browse components, search entities, etc.) with steps describing what user and system do. Dashboard renders flows as flowcharts. Tags group flows by category. Audit rules enforce contract coverage via steps. The catch-all flow is replaced with ~20-30 user-goal flows covering all 71 contracts.
scope: 'In scope: FlowEntity.Steps model, CLI flags, audit rules, dashboard flowchart view, real flow authoring, catch-all deletion. Out of scope: sequence diagram DSL, flow simulation/execution, contract-side changes.'
design: |-
    Flows are **user journeys** — one flow per user goal, describing what the user and system do at each step. Examples: "Create Plan", "Approve Plan", "Browse Components", "Search Entities", "Reindex Storage". Each flow answers: what does the user want to achieve, and what sequence of contract interactions does the system perform to get there?

    **FlowStep** is a value type on FlowEntity: id (short intra-flow identifier), action (describes what user or system does — the actor is embedded in the action text, e.g. "User runs syde plan create", "System allocates plan ID"), contract (slug of the contract this step exercises — empty for internal steps), description (optional elaboration), on_success (step ID or "done"), on_failure (step ID, "done", or "abort").

    One contract per step. One flow per user goal. If a goal touches many contracts, the flow has many steps — that is correct and desirable. The step sequence IS the user journey.

    **Tags group flows** by category (planning, authoring, querying, dashboard, storage, init). The dashboard filters by tag. No structural nesting.

    **Audit replaces the edge-only check** with step-aware rules: ERROR when a contract has no flow step across all flows, ERROR on unresolvable step contract/on_success/on_failure refs, ERROR on duplicate step IDs, WARN on steps with empty contract field.

    **Dashboard flowchart** renders each step as a card node with action text and a clickable contract chip. on_success edges are solid green, on_failure are dashed red. Implicit next for steps without on_success. Terminal nodes for done/abort.

    **Skill docs** teach the flow-per-user-goal pattern, the --step flag, and the audit rules.
source: manual
created_at: "2026-04-16T09:20:27Z"
approved_at: "2026-04-16T10:30:53Z"
completed_at: "2026-04-17T09:16:35Z"
phases:
    - id: phase_1
      name: Flow step model and CLI
      status: completed
      description: Add FlowStep struct to entity model and CLI flags for authoring steps
      objective: FlowEntity has Steps []FlowStep with Action, Contract, Description, OnSuccess, OnFailure; CLI supports --step flag
      changes: internal/model/entity.go (FlowStep struct, Steps field), internal/cli/add.go and update.go (--step flag parsing)
      details: Add FlowStep struct with yaml/json tags. Add Steps []FlowStep to FlowEntity. Add --step repeatable flag with pipe-separated format action|contract|description|on_success|on_failure. Parse in add.go and update.go flow cases. Update entity-spec.md reference doc.
      tasks:
        - add-flowstep-struct-and-steps-field-to-flowentity
        - add-step-cli-flag-for-flow-add-and-update
        - update-entity-spec-reference-doc
    - id: phase_2
      name: Step-based audit rules
      status: completed
      description: Replace edge-only contract-flow check with step-based validation
      objective: Audit WARNs on steps missing contract refs and ERRORs on contracts not in any flow step
      changes: internal/audit/graph_rules.go (rewrite contractFlowFindings), new step validation
      details: Rewrite contractFlowFindings to scan FlowEntity.Steps for contract refs instead of relationship edges. Build inFlow set from step contract slugs. WARN for steps with empty contract field. ERROR for contracts not in any step. Remove the catch-all detection heuristic since we are deleting the flow anyway.
      tasks:
        - rewrite-contractflowfindings-to-use-steps
        - create-requirements-for-new-audit-rules
    - id: phase_3
      name: Dashboard flowchart rendering
      status: completed
      description: Render flow steps as a connected flowchart in the flow detail panel
      objective: Flow detail panel shows a visual flowchart with step nodes, contract chips, and directed edges for success/failure paths
      changes: web/src/components/ (new FlowChart component), dashboard API (expose steps in flow detail)
      details: 'Backend: ensure GET /api/<proj>/entity/<slug> returns steps array for flow entities. Frontend: create a FlowChart React component that renders steps as connected nodes. Each node shows action text and a clickable contract chip. Edges connect on_success (green) and on_failure (red) branches. Use CSS grid or absolute positioning for layout. Integrate into the flow detail panel (EntityDetail or a flow-specific panel).'
      tasks:
        - expose-flow-steps-in-entity-detail-api
        - integrate-flowchart-into-flow-detail-panel
        - build-flowchart-react-component
        - render-plan-prose-fields-as-markdown-in-plandetailpanel
    - id: phase_4
      name: Author user-goal flows
      status: completed
      description: Delete catch-all flow, author ~30 user-goal flows with structured steps covering all 71 contracts
      objective: Every contract is referenced by at least one flow step in a user-goal flow; flows are tagged by category
      changes: Delete design-model-operations-coverage flow. Create ~25 new flows with steps. Add steps to 4 existing flows. Tag all flows by category (planning, authoring, querying, tree, dashboard, init).
      details: Use bulk shell scripts for efficiency. Each flow gets trigger, goal, narrative, tags, and structured steps. Steps describe what user and system do at each point. After authoring, run syde sync check to verify all 71 contracts are covered.
      tasks:
        - delete-catch-all-flow-and-remove-contract-relationships
        - author-entity-crud-lifecycle-flow-with-steps
        - author-query-and-search-flow-with-steps
        - author-dashboard-browsing-flow-with-steps
        - author-index-and-storage-operations-flow-with-steps
        - author-project-initialization-flow-with-steps
        - add-steps-to-existing-flows
    - id: phase_5
      name: Verify and finish
      status: completed
      description: Build, browser smoke, sync check, update skill docs
      objective: Everything builds, dashboard renders flowcharts, all contracts have flow step coverage, sync check passes
      changes: No source changes; build, test, verify
      details: go build, bun run build, browser check flowchart rendering, syde sync check --strict, update skill SKILL.md if flow step documentation needed
      tasks:
        - build-and-browser-smoke-test
        - run-sync-check-and-plan-complete
        - update-skill-docs-for-flow-steps
changes:
    requirements:
        new:
            - id: hzt9
              name: Flow entity shall carry structured steps
              what: FlowEntity gains a Steps []FlowStep field
              why: Free-text prose makes contract-flow traceability unverifiable
              draft:
                priority: must
                rationale: Structured steps enable programmatic audit and visual rendering
                req_type: functional
                source: plan
                statement: The syde entity model shall provide a structured steps field on flow entities.
                verification: FlowEntity has Steps []FlowStep with yaml/json tags; round-trip preserves data
              tasks:
                - add-flowstep-struct-and-steps-field-to-flowentity
            - id: 20ko
              name: Each flow step shall reference at most one contract
              what: contract field is a single slug, not a list
              why: One-to-one step-contract mapping keeps coverage audit simple and deterministic
              draft:
                priority: must
                rationale: If a behavior touches two contracts, model as two steps — the step sequence IS the journey
                req_type: constraint
                source: plan
                statement: The syde entity model shall restrict each flow step to reference at most one contract by slug.
                verification: FlowStep.Contract is string not []string; multi-contract behaviors split into multiple steps
              tasks:
                - add-flowstep-struct-and-steps-field-to-flowentity
            - id: dmz1
              name: Audit shall error on contracts not in any flow step
              what: ERROR when a contract entity is not referenced by any flow step across all flows
              why: Replaces the edge-only check defeated by the catch-all flow
              draft:
                priority: must
                rationale: Every contract boundary must participate in a documented user journey
                req_type: functional
                source: plan
                statement: The syde audit engine shall report an error for any contract entity whose slug does not appear in the contract field of at least one flow step across all flows.
                verification: syde sync check errors on contracts missing from all flow steps
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: dhzo
              name: Audit shall error on unresolvable step contract refs
              what: ERROR when a step contract slug does not resolve to an existing contract entity
              why: Broken refs mean the flow documents a contract that does not exist
              draft:
                priority: must
                rationale: Dangling refs rot silently; catch them at audit time
                req_type: functional
                source: plan
                statement: If a flow step references a contract slug that does not resolve to an existing contract entity, then the syde audit engine shall report an error.
                verification: Adding a step with a nonexistent contract slug causes syde sync check to error
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: n1sq
              name: Audit shall error on unresolvable step branching refs
              what: ERROR when on_success or on_failure references a step ID that does not exist in the same flow
              why: Broken branching refs produce disconnected flowcharts
              draft:
                priority: must
                rationale: Graph integrity is a precondition for correct flowchart rendering
                req_type: functional
                source: plan
                statement: If a flow step on_success or on_failure field references a step ID that does not exist in the same flow, then the syde audit engine shall report an error.
                verification: Adding a step with on_success pointing to nonexistent ID causes syde sync check to error
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: 3eyk
              name: Flowchart success edges shall be solid green
              what: on_success edges render as solid green arrows
              why: Visual distinction between success and failure paths
              draft:
                priority: should
                rationale: Color-coded edges make branching immediately scannable
                req_type: usability
                source: plan
                statement: The dashboard flowchart shall render on_success edges as solid green arrows.
                verification: Flowchart shows green solid arrows for success paths
              tasks:
                - build-flowchart-react-component
            - id: cmk3
              name: Flowchart failure edges shall be dashed red
              what: on_failure edges render as dashed red arrows
              why: Visual distinction between success and failure paths
              draft:
                priority: should
                rationale: Color-coded edges make branching immediately scannable
                req_type: usability
                source: plan
                statement: The dashboard flowchart shall render on_failure edges as dashed red arrows.
                verification: Flowchart shows red dashed arrows for failure paths
              tasks:
                - build-flowchart-react-component
            - id: jsuz
              name: Catch-all flow shall be replaced with user-goal flows
              what: Design Model Operations Coverage flow deleted; ~30 user-goal flows created
              why: The catch-all defeats contract-flow traceability
              draft:
                priority: must
                rationale: A single flow covering 71 contracts documents nothing; per-goal flows document everything
                req_type: constraint
                source: plan
                statement: The syde design model shall not contain a catch-all flow covering all contracts and shall instead use per-user-goal flows with structured steps.
                verification: syde query design-model-operations-coverage returns not found; all contracts covered by user-goal flow steps
              tasks:
                - delete-catch-all-flow-and-remove-contract-relationships
            - id: uvo9
              name: Flow steps shall describe user and system actions
              what: Step action text names the actor performing the action
              why: User journeys must show the handoff between user and system
              draft:
                priority: should
                rationale: Steps like 'do thing' are useless; 'User runs syde plan create' is actionable
                req_type: usability
                source: plan
                statement: When a flow has structured steps, the syde skill shall require each step action to describe what the user or system does at that point in the journey.
                verification: Skill docs teach action text conventions; review during flow authoring
              tasks:
                - update-skill-docs-for-flow-steps
            - id: fzop
              name: Flow step shall have six fields
              what: FlowStep struct has six fields
              why: Minimum fields for a step that can be audited and rendered as a flowchart node
              draft:
                priority: must
                rationale: 'Each field serves a distinct purpose: id for linking, action for display, contract for traceability, on_success/on_failure for graph edges'
                req_type: functional
                source: plan
                statement: The syde entity model shall define each flow step with an id, action, contract, description, on_success, and on_failure field.
                verification: FlowStep struct has all six fields with correct yaml/json tags
              tasks:
                - add-flowstep-struct-and-steps-field-to-flowentity
            - id: t5dh
              name: Each flow shall represent one user goal
              what: Flows are user journeys — one flow per user goal
              why: Broad catch-all flows defeat the purpose of behavioral documentation
              draft:
                priority: must
                rationale: Granular flows are reviewable and auditable; catch-all flows are not
                req_type: constraint
                source: plan
                statement: When creating a flow entity, the syde CLI shall enforce that each flow represents a single user goal with a clear trigger and outcome.
                verification: Skill docs teach the one-flow-per-goal pattern; plan check warns on flows with excessive contract coverage
              tasks:
                - update-skill-docs-for-flow-steps
            - id: mtlb
              name: Steps without on-success shall connect to next
              what: Empty on_success means implicit next in array order
              why: Most flows are linear; requiring explicit next on every step is noise
              draft:
                priority: must
                rationale: Convention over configuration for linear flows
                req_type: functional
                source: plan
                statement: When a flow step has an empty on_success field, the syde dashboard shall render it as connected to the next step in array order.
                verification: Flowchart renders linear steps with implicit edges; last step connects to done
              tasks:
                - build-flowchart-react-component
            - id: hs6p
              name: Audit shall error on duplicate step IDs
              what: ERROR when two steps in the same flow share the same ID
              why: Duplicate IDs make on_success/on_failure refs ambiguous
              draft:
                priority: must
                rationale: Step IDs are the addressing mechanism for the intra-flow graph
                req_type: functional
                source: plan
                statement: If two or more steps within the same flow entity share the same id value, then the syde audit engine shall report an error.
                verification: Adding two steps with id 's1' in one flow causes syde sync check to error
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: ne2q
              name: Audit shall warn on steps with empty contract
              what: WARN when a step has no contract ref (internal/human step)
              why: Internal steps are allowed but should be visible for review
              draft:
                priority: should
                rationale: Internal steps are legitimate but catch-all abuse hides behind them
                req_type: functional
                source: plan
                statement: When a flow step has an empty contract field, the syde audit engine shall report a warning.
                verification: A step without contract slug causes syde sync check to warn but not error
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: bdnp
              name: Dashboard shall render flow steps as flowchart
              what: Flow detail panel shows steps as card nodes with directed edges
              why: Visual flowcharts make user journeys reviewable at a glance
              draft:
                priority: must
                rationale: Prose flows are unreadable at scale; visual rendering is the payoff of structured steps
                req_type: functional
                source: plan
                statement: When a flow entity has structured steps, the dashboard shall render them as a connected flowchart with card nodes and directed edges.
                verification: Opening a flow with steps in the dashboard shows a visual flowchart
              tasks:
                - build-flowchart-react-component
            - id: hy40
              name: Flows shall be tagged by category
              what: Each flow carries tags like planning, authoring, querying, dashboard, init, tree
              why: 30 flows need grouping for dashboard navigation
              draft:
                priority: should
                rationale: Tags are the lightweight grouping mechanism; no structural nesting needed
                req_type: functional
                source: plan
                statement: When creating a flow entity, the syde CLI shall support tagging flows by category for dashboard filtering.
                verification: syde query --kind flow --tag planning returns only planning-tagged flows
              tasks:
                - add-steps-to-existing-flows
            - id: xom3
              name: Plans shall decompose into granular requirements
              what: Skill docs enforce granular requirement decomposition before implementation
              why: Plans with few broad requirements miss design intent that should be traceable
              draft:
                priority: must
                rationale: Every design decision must trace to a requirement; a plan with many phases and few requirements is incomplete
                req_type: constraint
                source: plan
                statement: When creating a plan, the syde skill shall require decomposing the request into granular requirements before declaring implementation changes.
                verification: Skill SKILL.md documents the granular requirement decomposition pattern with categories checklist
              tasks:
                - update-skill-docs-for-flow-steps
            - id: glnh
              name: Dashboard shall render plan prose as markdown
              what: Background, objective, scope, and design fields render as markdown in the plan detail panel
              why: These fields contain structured prose with headers, lists, and code blocks that are unreadable as plain text
              draft:
                priority: must
                rationale: Plan design fields are written as markdown but displayed as plain text, losing all formatting
                req_type: usability
                source: plan
                statement: When displaying a plan entity, the dashboard shall render the background, objective, scope, and design fields as markdown.
                verification: Opening a plan with markdown in its design field renders headers, lists, and code blocks correctly
              tasks:
                - build-and-browser-smoke-test
    components:
        extended:
            - id: ouvf
              slug: entity-model
              what: Add Steps []FlowStep to FlowEntity with structured step fields
              why: Flows need structured steps instead of free-text prose
              tasks:
                - add-flowstep-struct-and-steps-field-to-flowentity
            - id: age8
              slug: cli-commands
              what: Add --step repeatable flag for flow add/update
              why: CLI must support authoring structured steps
              tasks:
                - add-step-cli-flag-for-flow-add-and-update
            - id: 0sjd
              slug: audit-engine
              what: New step-based contract-flow audit replacing edge-only check
              why: Current check is defeated by catch-all flow pattern
              tasks:
                - rewrite-contractflowfindings-to-use-steps
            - id: euf9
              slug: skill-installer
              what: Update SKILL.md flow section with structured steps authoring guide and --step flag documentation
              why: Agents need to know how to author flow steps from the start
              tasks:
                - update-skill-docs-for-flow-steps
            - id: bm9d
              slug: web-spa
              what: Render flow steps as flowchart, render plan prose as markdown
              why: Flowchart visualization for flows; markdown rendering for plan design/background/objective/scope
              tasks:
                - build-flowchart-react-component
                - integrate-flowchart-into-flow-detail-panel
                - render-plan-prose-fields-as-markdown-in-plandetailpanel
    flows:
        deleted:
            - id: dp7p
              slug: design-model-operations-coverage
              why: Catch-all backfill flow defeats the contract-flow audit
              tasks:
                - delete-catch-all-flow-and-remove-contract-relationships
        extended:
            - id: oaem
              slug: plan-lifecycle
              what: Add structured steps referencing plan/task contracts
              why: Existing flow needs steps added
              tasks:
                - add-steps-to-existing-flows
            - id: hdqx
              slug: drift-detection-and-clearing
              what: Add structured steps referencing validate/task-done contracts
              why: Existing flow needs steps added
              tasks:
                - add-steps-to-existing-flows
            - id: sd1d
              slug: session-start-bootstrap
              what: Add structured steps referencing session-context contract
              why: Existing flow needs steps added
              tasks:
                - add-steps-to-existing-flows
            - id: had5
              slug: source-file-summary-tree-maintenance
              what: Add structured steps referencing tree contracts
              why: Existing flow needs steps added
              tasks:
                - add-steps-to-existing-flows
        new:
            - id: q3bq
              name: Create Plan
              what: 'User-goal flow: agent creates a plan with phases and tasks'
              why: One flow per user goal, describing user+system interactions
              draft:
                goal: A draft plan with background, objective, scope, design, phases, and tasks exists
                trigger: User asks agent to implement a non-trivial change
              tasks:
                - author-entity-crud-lifecycle-flow-with-steps
            - id: 6e2p
              name: Add Entity
              what: 'User-goal flow: user adds a new entity to the design model'
              why: CRUD contracts need per-goal user journey flows
              draft:
                goal: A new entity file exists with allocated ID and index entries
                trigger: User runs syde add <kind> <name>
              tasks:
                - author-query-and-search-flow-with-steps
            - id: djm2
              name: Browse Components
              what: 'User-goal flow: user navigates to components inbox in dashboard'
              why: Each dashboard screen is a distinct user goal
              draft:
                goal: User sees component list and can select one for detail view
                trigger: User navigates to /<project>/component in the dashboard
              tasks:
                - author-dashboard-browsing-flow-with-steps
---
