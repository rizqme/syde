---
id: PLN-0005
kind: plan
name: Plan requirement coverage and overlap audit
slug: plan-requirement-coverage-and-overlap-audit-y33i
relationships:
    - target: approved-plan-plan-requirement-coverage-and-overlap-audit-eikl
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-16T09:49:07Z"
plan_status: completed
background: Plans can pass syde plan check with too few requirements relative to their implementation scope. Also, new requirements may overlap existing ones without explicit acknowledgement.
objective: 'Two new plan authoring audit rules: (1) requirement coverage score ensures the requirements lane is proportional to the plan''s changes, (2) similarity check finds existing requirements that overlap new ones and requires explicit acknowledgement.'
scope: internal/audit/plan_authoring.go only. No CLI or dashboard changes.
design: 'Coverage algorithm: count the plan''s non-requirement changes (deleted + extended + new across components/contracts/concepts/flows/systems). Count the plan''s requirement changes. If requirements < changes/3, WARN that the plan may be under-specified. Overlap detection: for each new requirement in the plan, search all existing active requirements by comparing key terms in the statement. If cosine similarity of term sets > 0.5 (or >50% shared significant words), WARN that the new requirement may overlap an existing one and should be marked as refines/derives_from or the existing one superseded.'
source: manual
created_at: "2026-04-16T09:43:42Z"
approved_at: "2026-04-16T09:46:15Z"
completed_at: "2026-04-16T09:49:07Z"
phases:
    - id: phase_1
      name: Implement audit rules
      status: completed
      description: Add coverage and overlap checks to plan_authoring.go
      objective: syde plan check warns on under-specified plans and overlapping requirements
      changes: internal/audit/plan_authoring.go
      details: Add two new finding generators inside planAuthoringFindings
      tasks:
        - add-requirement-coverage-check
        - add-requirement-overlap-detection
        - build-and-verify
changes:
    requirements:
        new:
            - id: cxiy
              name: Plan check shall warn on low requirement coverage
              what: WARN when requirements lane is thin relative to implementation changes
              why: Under-specified plans miss design intent
              draft:
                priority: must
                rationale: Every design decision should trace to a requirement
                req_type: functional
                source: plan
                statement: When a plan has fewer requirement changes than one third of its non-requirement changes, the syde audit engine shall warn that the plan may be under-specified.
                verification: A plan with 9 component changes and 1 requirement triggers the warning
            - id: 7lik
              name: Plan check shall warn on overlapping requirements
              what: WARN when a new requirement in the plan overlaps an existing active requirement
              why: Overlapping requirements cause confusion and should be explicitly related
              draft:
                priority: must
                rationale: Overlapping requirements should be linked via refines/derives_from or the old one superseded
                req_type: functional
                source: plan
                statement: When a new requirement in a plan shares significant terms with an existing active requirement, the syde audit engine shall warn that the requirements may overlap.
                verification: A plan adding a requirement similar to an existing one triggers the overlap warning
    components:
        extended:
            - id: 0anl
              slug: audit-engine
              what: Add requirement coverage and overlap checks to planAuthoringFindings
              why: Plans with too few requirements pass check; overlapping requirements go unnoticed
---
