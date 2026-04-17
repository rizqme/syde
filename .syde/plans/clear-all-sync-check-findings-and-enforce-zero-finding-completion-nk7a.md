---
id: PLN-0009
kind: plan
name: Clear all sync check findings and enforce zero-finding completion
slug: clear-all-sync-check-findings-and-enforce-zero-finding-completion-nk7a
relationships:
    - target: approved-plan-clear-all-sync-check-findings-and-enforce-zero-finding-completion-peda
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T09:16:36Z"
plan_status: completed
background: 'syde sync check --strict shows 138 errors and 98 warnings. Most are batch-fixable: TF-IDF overlaps need --audited, plans/tasks need belongs_to, auto-generated requirements need EARS fields. Also, plans can currently be completed while findings exist — the completion gate should block on any sync check error.'
objective: syde sync check --strict exits 0. syde plan complete refuses when sync check has errors. Skill docs enforce this.
scope: Batch entity fixes, plan completion gate enhancement, skill doc update.
design: 'Phase 1: Fix auto-generated Approved plan requirements (EARS, fields, belongs_to). Phase 2: Fix all belongs_to on plans/tasks. Phase 3: Batch --audited on all 92 TF-IDF overlaps. Phase 4: Fix missing requirement relationships on entities. Phase 5: Add sync check gate to plan complete and skill docs.'
source: manual
created_at: "2026-04-17T01:36:05Z"
approved_at: "2026-04-17T01:36:05Z"
completed_at: "2026-04-17T09:16:36Z"
phases:
    - id: phase_1
      name: Fix auto-generated requirements
      status: completed
      description: Fix EARS, fields, belongs_to on auto-generated Approved plan requirements
      objective: All Approved plan requirements pass validation
      changes: 3 requirement entities
      details: syde update each with --statement, --type, --priority, --verification, --add-rel syde:belongs_to
      tasks:
        - fix-approved-plan-requirements
    - id: phase_2
      name: Fix belongs_to on plans and tasks
      status: completed
      description: Add belongs_to:syde on all plans and tasks missing it
      objective: 0 belongs_to errors
      changes: ~18 plan/task entities
      details: 'Batch script: syde update <slug> --add-rel syde:belongs_to'
      tasks:
        - batch-fix-belongsto
    - id: phase_3
      name: Acknowledge TF-IDF overlaps
      status: completed
      description: Batch --audited on all 92 requirement overlap pairs
      objective: 0 TF-IDF overlap errors
      changes: ~46 requirement entities get audited_overlaps populated
      details: 'Script: parse sync check output, extract slug pairs, run syde update --audited for each'
      tasks:
        - batch-acknowledge-overlaps
    - id: phase_4
      name: Fix missing requirement relationships
      status: completed
      description: Add outgoing requirement references on entities missing them
      objective: 0 missing-requirement-relationship warnings
      changes: ~46 entities get --add-rel <req>:references
      details: For each entity without a requirement relationship, find the nearest matching requirement and link it
      tasks:
        - batch-fix-requirement-relationships
    - id: phase_5
      name: Enforce sync check in plan complete
      status: completed
      description: Add sync check gate to syde plan complete and update skill docs
      objective: syde plan complete refuses with sync check errors; skill teaches this
      changes: internal/cli/plan.go, skill/SKILL.md
      details: In planCompleteCmd, run sync check before marking complete. If errors exist (excluding the plan-completion findings themselves), refuse. Update SKILL.md Phase 5 to document this.
      tasks:
        - add-sync-check-gate-to-plan-complete
        - update-skill-docs
changes:
    requirements:
        new:
            - id: quu5
              name: Plan complete shall require clean sync check
              what: syde plan complete blocks when syde sync check has errors
              why: Plans should not complete while the model has integrity issues
              draft:
                priority: must
                rationale: Completion without clean sync is a false gate
                req_type: functional
                source: plan
                statement: When completing a plan, the syde CLI shall refuse if syde sync check reports any errors.
                verification: syde plan complete fails when sync check has errors
    components:
        extended:
            - id: okhq
              slug: cli-commands
              what: Add sync check gate to plan complete
              why: Enforce clean model at completion time
            - id: 9zmt
              slug: skill-installer
              what: 'Add rule: no plan completion with sync check errors'
              why: Skill must teach the enforcement
---
