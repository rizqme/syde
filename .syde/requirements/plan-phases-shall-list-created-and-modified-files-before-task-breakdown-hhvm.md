---
id: REQ-0491
kind: requirement
name: Plan phases shall list created and modified files before task breakdown
slug: plan-phases-shall-list-created-and-modified-files-before-task-breakdown-hhvm
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:47Z"
statement: When a plan phase is authored, the syde skill shall require the phase to enumerate every file to be created or modified in that phase with a one-line responsibility per file before any task in the phase is added.
req_type: functional
priority: should
verification: skill/references/plan-authoring.md documents the Files-section requirement and skill/SKILL.md Phase 2 section references it
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Derived from obra/superpowers writing-plans.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:47Z"
---
