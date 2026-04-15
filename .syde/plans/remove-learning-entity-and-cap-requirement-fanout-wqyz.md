---
id: PLN-0021
kind: plan
name: Remove Learning Entity And Cap Requirement Fanout
slug: remove-learning-entity-and-cap-requirement-fanout-wqyz
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: approved-plan-remove-learning-entity-and-cap-requirement-fanout-wf5u
      type: references
      label: requirement
updated_at: "2026-04-15T09:38:13Z"
plan_status: completed
background: A migration baseline requirement currently links to hundreds of entities, and learning entities are considered flawed and should be removed rather than exempted.
objective: Remove the learning entity kind completely and make strict sync prevent catch-all requirements by capping linked entities per kind.
scope: Remove learning model/CLI/API/dashboard/docs/data surfaces; add an audit fanout cap of 10 entities per kind per requirement; do not redesign requirements beyond the cap.
source: manual
created_at: "2026-04-15T08:22:19Z"
approved_at: "2026-04-15T08:25:15Z"
completed_at: "2026-04-15T09:38:13Z"
phases:
    - id: phase_1
      name: Learning model removal
      status: completed
      description: Remove learning as a first-class entity kind from the Go model and storage surfaces.
      objective: The codebase no longer defines KindLearning or LearningEntity.
      changes: Delete learning model definitions and remove learning kind dispatch/validation/query hooks.
      details: Update model entity kind lists, constructors, validation, resolver/query learning counters, memory sync, and audit assumptions so learning is no longer loaded as an entity.
      tasks:
        - remove-learning-entity-model-paths
    - id: phase_2
      name: Learning UI and command removal
      status: completed
      description: Remove user-facing learning commands and dashboard views.
      objective: No syde CLI or dashboard surface offers learning entities.
      changes: Remove remember/learn command registration, dashboard LearningFeed route/sidebar/API, relationship icon styling, and docs references.
      details: Delete or disconnect learning-specific CLI file code, HTTP handlers, React page and special-view routing, sidebar item, API client methods/types, and skill command docs.
      tasks:
        - remove-learning-cli-and-http-apis
        - remove-learning-dashboard-ui
    - id: phase_3
      name: Requirement fanout cap
      status: completed
      description: Prevent catch-all requirements from satisfying strict traceability.
      objective: Strict sync reports requirements with more than 10 linked entities of the same kind.
      changes: Add audit fanout validation and update docs.
      details: Count outbound non-requirement relationships to requirement targets by requirement and source kind; emit traceability errors over the cap.
      tasks:
        - enforce-requirement-fanout-cap
    - id: phase_4
      name: Data cleanup and verification
      status: completed
      description: Delete existing learning files, refresh generated skill copies, and verify expected strict behavior.
      objective: The build passes and strict sync reports the intended baseline fanout violations until requirements are backfilled.
      changes: Remove .syde/learnings data files, run tests/build/install, tree summaries, and strict sync.
      details: Delete learning markdown files because the entity kind is gone; use expected sync output to prove baseline fanout is now caught.
      tasks:
        - update-docs-and-installed-skills-after-learning-removal
        - delete-learning-data-and-verify-behavior
---
