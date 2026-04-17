---
id: PLN-0007
kind: plan
name: Requirement overlap audit with mandatory acknowledgement
slug: requirement-overlap-audit-with-mandatory-acknowledgement-u5lj
relationships:
    - target: approved-plan-requirement-overlap-audit-with-mandatory-acknowledgement-xgvu
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T09:16:36Z"
plan_status: completed
background: 'The current overlap detection runs only at plan check time and emits a WARN. Authors can ignore it. Requirements with unacknowledged overlaps slip through, causing confusion when two requirements say nearly the same thing. The overlap check needs to be mandatory: detect at creation time, ERROR at sync check if not audited.'
objective: When syde add requirement creates a requirement, it automatically finds similar existing requirements and prints them. The requirement entity carries an audited_overlaps []string field. syde sync check ERRORs on any requirement whose detected overlaps are not listed in audited_overlaps. Authors must pass --audited <slug> per overlap to acknowledge.
scope: 'In scope: RequirementEntity.AuditedOverlaps field, syde add requirement overlap detection + --audited flag, syde sync check ERROR rule, syde update --audited flag. Out of scope: changing the plan-check WARN (it stays as an early signal).'
design: |-
    The core idea is **detect-then-acknowledge**: the system finds overlaps automatically, but a human must explicitly say "I reviewed this and it is not a duplicate."

    **RequirementEntity gains AuditedOverlaps []string** — slugs of existing requirements the author reviewed and confirmed are distinct. Stored in YAML as `audited_overlaps`.

    **syde add requirement** runs the same significantTerms/termOverlap logic that plan_authoring uses today. If any existing active requirement has >50% term overlap, the CLI prints each match with its slug, name, and overlap percentage. The command still succeeds (the requirement is created) but without --audited flags the sync check will ERROR later. If the author passes --audited <slug> for each overlap, those slugs are stored in audited_overlaps on creation.

    **syde sync check** gains a new audit rule: for each active requirement, compute overlaps against all other active requirements. If an overlap is detected and the target slug is NOT in audited_overlaps, emit ERROR. This catches both newly created requirements (from syde add) and existing requirements that gained new overlaps when a nearby requirement was added.

    **syde update <req> --audited <slug>** appends to audited_overlaps, allowing post-hoc acknowledgement.

    The plan-check WARN stays as-is — it is an early signal during plan drafting before the requirement entity exists. The sync-check ERROR is the enforcement gate.
source: manual
created_at: "2026-04-16T11:29:02Z"
approved_at: "2026-04-16T11:33:51Z"
completed_at: "2026-04-17T09:16:36Z"
phases:
    - id: phase_1
      name: Model and CLI
      status: completed
      description: Add AuditedOverlaps field, --audited flag, overlap detection on create
      objective: syde add requirement detects overlaps and accepts --audited; syde update supports --audited
      changes: internal/model/entity.go, internal/cli/add.go, internal/cli/update.go
      details: Add field to RequirementEntity. Add --audited StringArrayVar to add and update commands. In add.go, after creating the requirement, run overlap detection against all active requirements and print matches. Store --audited values in AuditedOverlaps.
      tasks:
        - add-auditedoverlaps-field-to-requiremententity
        - add-overlap-detection-to-syde-add-requirement
        - add-audited-flag-to-syde-update
    - id: phase_2
      name: Audit rule
      status: completed
      description: Add requirementOverlapFindings to sync check
      objective: syde sync check ERRORs on requirements with unaudited overlaps
      changes: internal/audit/plan_authoring.go or new file
      details: 'New finding generator: for each active requirement, compute overlaps against all other active requirements. If overlap >50% and target slug not in audited_overlaps, emit ERROR on both sides.'
      tasks:
        - add-sync-check-overlap-error-rule
    - id: phase_3
      name: Verify
      status: completed
      description: Build and sync check
      objective: Everything builds; sync check passes with properly audited requirements
      changes: No source changes
      details: go build, syde sync check
      tasks:
        - build-and-verify
changes:
    requirements:
        new:
            - id: 6gai
              name: Requirement creation shall detect similar requirements
              what: syde add requirement prints similar existing requirements on creation
              why: Authors need to see overlaps at the moment they create the requirement
              draft:
                priority: must
                rationale: Detection at creation time is the earliest possible feedback
                req_type: functional
                source: plan
                statement: When a requirement entity is created, the syde CLI shall print any existing active requirements with greater than 50 percent term similarity.
                verification: syde add requirement with a near-duplicate statement prints the similar requirement
              tasks:
                - add-overlap-detection-to-syde-add-requirement
            - id: 8uza
              name: Requirement creation shall accept audited flag
              what: syde add requirement --audited <slug> stores acknowledged overlaps
              why: Authors need a way to acknowledge overlaps at creation time
              draft:
                priority: must
                rationale: Acknowledgement at creation avoids a separate update step
                req_type: functional
                source: plan
                statement: When creating a requirement, the syde CLI shall accept a repeatable --audited flag storing acknowledged overlap slugs in audited_overlaps.
                verification: syde add requirement --audited <slug> populates audited_overlaps on the entity
              tasks:
                - add-overlap-detection-to-syde-add-requirement
            - id: hrxh
              name: Sync check shall error on unaudited requirement overlaps
              what: ERROR when a requirement has detected overlaps not in audited_overlaps
              why: Unacknowledged overlaps must be flagged as errors, not ignored
              draft:
                priority: must
                rationale: WARNs are ignorable; ERRORs block the session-end gate
                req_type: functional
                source: plan
                statement: If a requirement entity has term overlap greater than 50 percent with another active requirement whose slug is not in audited_overlaps, then the syde audit engine shall report an error.
                verification: syde sync check errors on requirement with unaudited overlap
              tasks:
                - add-sync-check-overlap-error-rule
            - id: egtx
              name: Requirement update shall support audited flag
              what: syde update <req> --audited <slug> appends to audited_overlaps
              why: Post-hoc acknowledgement for requirements created without --audited
              draft:
                priority: must
                rationale: Not all overlaps are known at creation time
                req_type: functional
                source: plan
                statement: When updating a requirement, the syde CLI shall accept a repeatable --audited flag appending acknowledged overlap slugs to audited_overlaps.
                verification: syde update <req> --audited <slug> adds to audited_overlaps list
              tasks:
                - add-audited-flag-to-syde-update
            - id: 71w9
              name: Overlap detection shall be bidirectional
              what: Both the new and the existing requirement must audit each other
              why: A new requirement overlapping an old one means both need acknowledgement
              draft:
                priority: must
                rationale: Unidirectional auditing leaves the older requirement unaware of the overlap
                req_type: functional
                source: plan
                statement: When the syde audit engine detects term overlap between two requirements, the engine shall report an error on each requirement that does not list the other in its audited_overlaps.
                verification: Both requirements show errors if neither has audited the other
              tasks:
                - add-sync-check-overlap-error-rule
            - id: zg6m
              name: Requirement entity shall carry audited overlaps list
              what: New AuditedOverlaps []string field
              why: Authors must explicitly acknowledge overlaps
              draft:
                priority: must
                rationale: Unacknowledged overlaps cause confusion
                req_type: functional
                source: plan
                statement: The syde entity model shall provide an audited_overlaps field on requirement entities listing slugs of reviewed overlapping requirements.
                verification: RequirementEntity has AuditedOverlaps with yaml tag
              tasks:
                - add-auditedoverlaps-field-to-requiremententity
    components:
        extended:
            - id: 8r3e
              slug: cli-commands
              what: Add --audited flag to syde add requirement and syde update; overlap detection on create
              why: CLI must support the new workflow
              tasks:
                - add-overlap-detection-to-syde-add-requirement
                - add-audited-flag-to-syde-update
            - id: 3qs8
              slug: audit-engine
              what: 'New requirementOverlapFindings rule: ERROR on unaudited overlaps'
              why: Sync check enforces mandatory acknowledgement
              tasks:
                - add-sync-check-overlap-error-rule
            - id: slk7
              slug: entity-model
              what: Add AuditedOverlaps []string to RequirementEntity
              why: New field for overlap acknowledgement
              tasks:
                - add-auditedoverlaps-field-to-requiremententity
---
