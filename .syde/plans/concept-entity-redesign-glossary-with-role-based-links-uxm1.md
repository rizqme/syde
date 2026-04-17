---
id: PLN-0006
kind: plan
name: 'Concept entity redesign: glossary with role-based links'
slug: concept-entity-redesign-glossary-with-role-based-links-uxm1
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      type: references
    - target: approved-plan-concept-entity-redesign-glossary-with-role-based-links-1ojq
      type: references
      label: requirement
updated_at: "2026-04-17T08:40:46Z"
plan_status: completed
background: Concepts currently model ERD-style entities with typed attributes, actions, and cardinality-labelled relates_to edges. The ERD canvas renders draggable cards with attribute rows and FK arrows. This is over-engineered for a design model — concepts should be high-level glossary entries that explain domain terms, not database schemas. The ERD view adds complexity without proportional value. Relationships to other entity kinds (components, contracts, flows) use generic 'references' which doesn't convey the concept's architectural role.
objective: 'Concepts are glossary entries with meaning, invariants, lifecycle. No attributes, no actions, no ERD canvas. Relationships use explicit role-based types: implemented_by (-> component), exposed_via (-> contract), used_in (-> flow), relates_to (-> concept, no cardinality). Dashboard shows linked glossary cards instead of ERD.'
scope: 'In scope: ConceptEntity field removal (attributes, actions), new relationship types, ERD canvas removal, dashboard glossary view, audit rule updates, skill doc updates. Out of scope: other entity kinds, no data migration of existing concepts (just strip fields on next save).'
design: |-
    Concepts become a **domain glossary** — each entry answers "what is this thing, why does it matter, and what does it connect to?" The ERD lens (attributes, types, FK arrows, cardinality) is removed entirely. Concepts are not data schemas; they are design-level explanations.

    **Structural changes to ConceptEntity:**
    - Remove: Attributes, Actions, ConceptRelationships, DataSensitivity, StructureNotes fields
    - Keep: Meaning (required), Invariants (required), Lifecycle (optional)
    - The struct becomes much simpler — just BaseEntity + Meaning + Invariants + Lifecycle

    **Relationship types gain three new values:**
    - `implemented_by` — concept -> component (which module owns this domain object)
    - `exposed_via` — concept -> contract (which boundaries expose this concept)
    - `used_in` — concept -> flow (which user journeys involve this concept)
    - `relates_to` — concept -> concept (no cardinality, optional prose label via rel.Label)

    These are more informative than generic `references` because they answer: where does this concept live in the architecture? The dashboard renders these as grouped relationship chips.

    **Dashboard changes:**
    - Remove the ERD canvas (concepts-erd-screen contract, the List/ERD toggle, and the ERD rendering code)
    - Keep the standard 2-column inbox (list + detail)
    - Detail panel shows: meaning, lifecycle, invariants, then grouped relationship chips (implemented_by, exposed_via, used_in, relates_to)

    **Audit changes:**
    - Remove concept attribute validation (required attributes check)
    - Remove cardinality enum validation on relates_to
    - Keep meaning as required (ERROR)
    - Keep invariants as recommended (WARN)
source: manual
created_at: "2026-04-16T11:17:10Z"
approved_at: "2026-04-17T02:45:17Z"
completed_at: "2026-04-17T08:40:46Z"
phases:
    - id: phase_1
      name: Simplify ConceptEntity and CLI
      status: completed
      description: Remove ERD fields from the entity model and CLI flags
      objective: ConceptEntity has only meaning, invariants, lifecycle; CLI flags match
      changes: internal/model/entity.go, internal/cli/add.go, internal/cli/update.go
      details: Remove Attributes, Actions, ConceptRelationships, DataSensitivity, StructureNotes from ConceptEntity struct. Remove --attribute, --action, --data-sensitivity, --structure-notes CLI flags. Keep --meaning, --invariants, --lifecycle.
      tasks:
        - remove-erd-fields-from-conceptentity
        - remove-erd-cli-flags
    - id: phase_2
      name: Update audit and API
      status: completed
      description: Remove concept attribute validation, cardinality check; update API serialization
      objective: Audit validates meaning required, invariants recommended; API returns simplified concept
      changes: internal/audit/concepts.go, internal/query/engine.go, internal/query/formatter.go
      details: Remove conceptFindings attribute checks, cardinality enum validation. Remove attributes/actions from EntitySummary and FormatJSON.
      tasks:
        - simplify-concept-audit-rules
        - remove-concept-erd-fields-from-api
    - id: phase_3
      name: 'Dashboard: remove ERD, update detail'
      status: completed
      description: Remove ERD canvas and toggle, update concept detail panel
      objective: Concepts page is standard inbox; detail shows glossary with grouped relationship chips
      changes: web/src/ (remove ERD components, update EntityDetail concept case)
      details: Remove List/ERD toggle and ERD canvas rendering code. Update the concept case in EntityDetail to show meaning, lifecycle, invariants as prose sections, then grouped relationship chips (implemented_by, exposed_via, used_in, relates_to).
      tasks:
        - remove-erd-canvas-from-concepts-page
        - update-concept-detail-to-glossary-style
    - id: phase_4
      name: Migrate concepts and update docs
      status: completed
      description: Update existing concepts to glossary style, add role-based links, update skill
      objective: All 10 existing concepts use the new model; skill docs teach glossary pattern
      changes: Existing concept entities, skill/SKILL.md, skill/references/entity-spec.md
      details: Batch update existing concepts to remove attributes/actions (they'll be stripped on save). Add implemented_by/exposed_via/used_in relationships where appropriate. Update skill docs.
      tasks:
        - migrate-existing-concepts-to-glossary-style
        - update-skill-docs-for-glossary-concepts
    - id: phase_5
      name: Verify
      status: pending
      description: Build, browser smoke, sync check
      objective: Everything builds; dashboard renders glossary concepts; sync check passes
      changes: No source changes
      details: go build, bun run build, browser check, syde sync check --strict
      tasks:
        - build-and-verify
changes:
    requirements:
        extended:
            - id: ozl5
              slug: concepts-erd-node-click-returns-to-list
              what: 'Supersede: ERD canvas is being removed entirely'
              why: The ERD view no longer exists; this requirement has no target
              field_changes:
                requirement_status: superseded
              tasks:
                - remove-erd-canvas-from-concepts-page
            - id: yz29
              slug: concepts-erd-route-renders-canvas
              what: 'Supersede: ERD canvas is being removed entirely'
              why: The ERD view no longer exists; this requirement has no target
              field_changes:
                requirement_status: superseded
              tasks:
                - remove-erd-canvas-from-concepts-page
        new:
            - id: h4pl
              name: Concepts shall be glossary entries not data schemas
              what: ConceptEntity is glossary
              why: ERD is over-engineered
              draft:
                priority: must
                rationale: Concepts explain terms, not schemas
                req_type: functional
                source: plan
                statement: The syde entity model shall define concept entities as domain glossary entries with meaning, invariants, and lifecycle fields.
                verification: ConceptEntity has meaning, invariants, lifecycle
              tasks:
                - remove-erd-fields-from-conceptentity
            - id: wwph
              name: Concept attributes field shall be removed
              what: Remove attributes
              why: ERD detail
              draft:
                priority: must
                rationale: Properties in code
                req_type: constraint
                source: plan
                statement: The syde entity model shall not include an attributes field on concept entities.
                verification: No attributes field
              tasks:
                - remove-erd-fields-from-conceptentity
            - id: v2kl
              name: Concept actions field shall be removed
              what: Remove actions
              why: ERD detail
              draft:
                priority: must
                rationale: Verbs in lifecycle
                req_type: constraint
                source: plan
                statement: The syde entity model shall not include an actions field on concept entities.
                verification: No actions field
              tasks:
                - remove-erd-fields-from-conceptentity
            - id: uvkj
              name: Concept meaning shall be required
              what: ERROR if missing
              why: Core of glossary
              draft:
                priority: must
                rationale: Meaning is core
                req_type: functional
                source: plan
                statement: The syde audit engine shall report an error for any concept entity with an empty meaning field.
                verification: sync check errors on empty meaning
              tasks:
                - simplify-concept-audit-rules
            - id: 2wpv
              name: Concept invariants shall be recommended
              what: WARN if missing
              why: Important but optional
              draft:
                priority: should
                rationale: Invariants document rules
                req_type: functional
                source: plan
                statement: The syde audit engine shall warn when a concept entity has an empty invariants field.
                verification: sync check warns on empty invariants
              tasks:
                - simplify-concept-audit-rules
            - id: trl6
              name: Concepts shall use implemented-by for components
              what: New rel type
              why: Role-based
              draft:
                priority: must
                rationale: Explicit roles
                req_type: functional
                source: plan
                statement: The syde entity model shall support an implemented_by relationship type from concept entities to component entities.
                verification: --add-rel works
              tasks:
                - migrate-existing-concepts-to-glossary-style
            - id: czsy
              name: Concepts shall use exposed-via for contracts
              what: New rel type
              why: Role-based
              draft:
                priority: must
                rationale: Explicit roles
                req_type: functional
                source: plan
                statement: The syde entity model shall support an exposed_via relationship type from concept entities to contract entities.
                verification: --add-rel works
              tasks:
                - migrate-existing-concepts-to-glossary-style
            - id: 7mmk
              name: Concepts shall use used-in for flows
              what: New rel type
              why: Role-based
              draft:
                priority: must
                rationale: Explicit roles
                req_type: functional
                source: plan
                statement: The syde entity model shall support a used_in relationship type from concept entities to flow entities.
                verification: --add-rel works
              tasks:
                - migrate-existing-concepts-to-glossary-style
            - id: xok9
              name: Relates-to shall drop cardinality labels
              what: No cardinality enum
              why: Glossary style
              draft:
                priority: must
                rationale: Prose labels more natural
                req_type: constraint
                source: plan
                statement: The syde audit engine shall not validate cardinality labels on relates_to relationships between concept entities.
                verification: relates_to without cardinality succeeds
              tasks:
                - simplify-concept-audit-rules
            - id: xpeu
              name: ERD canvas shall be removed
              what: Remove ERD toggle and canvas
              why: No data to render
              draft:
                priority: must
                rationale: Without attributes there is nothing to render
                req_type: constraint
                source: plan
                statement: The dashboard shall not render an ERD canvas for concept entities.
                verification: No ERD toggle on concepts page
              tasks:
                - remove-erd-canvas-from-concepts-page
            - id: 6dyq
              name: Concept detail shall show grouped relationship chips
              what: Grouped by type
              why: Scannable
              draft:
                priority: should
                rationale: Flat lists hard to scan
                req_type: usability
                source: plan
                statement: When displaying a concept entity, the dashboard shall group relationship chips by type with labeled sections.
                verification: Grouped sections visible in browser
              tasks:
                - update-concept-detail-to-glossary-style
    components:
        extended:
            - id: 8nw1
              slug: entity-model
              what: Remove ERD fields from ConceptEntity
              why: Glossary model
              tasks:
                - remove-erd-fields-from-conceptentity
            - id: zkc0
              slug: cli-commands
              what: Remove ERD CLI flags
              why: Match model
              tasks:
                - remove-erd-cli-flags
            - id: j0mx
              slug: audit-engine
              what: Simplify concept audit
              why: Match model
              tasks:
                - simplify-concept-audit-rules
            - id: hjjv
              slug: web-spa
              what: Remove ERD, glossary detail
              why: Match model
              tasks:
                - remove-erd-canvas-from-concepts-page
                - update-concept-detail-to-glossary-style
            - id: duaj
              slug: query-engine
              what: Remove ERD fields from API
              why: Match model
              tasks:
                - remove-concept-erd-fields-from-api
            - id: lknl
              slug: skill-installer
              what: Update concept docs
              why: Match model
              tasks:
                - update-skill-docs-for-glossary-concepts
    contracts:
        deleted:
            - id: sm85
              slug: concepts-erd-screen
              why: ERD canvas removed
              tasks:
                - remove-erd-canvas-from-concepts-page
---
