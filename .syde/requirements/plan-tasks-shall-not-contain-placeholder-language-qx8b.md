---
id: REQ-0493
kind: requirement
name: Plan tasks shall not contain placeholder language
slug: plan-tasks-shall-not-contain-placeholder-language-qx8b
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:48Z"
statement: If a plan task's details or acceptance contains placeholder phrases including TBD, TODO, implement later, add appropriate error handling, similar to Task N, or write tests for the above, then the syde skill shall require the author to rewrite the task with explicit content before approval.
req_type: functional
priority: should
verification: skill/references/plan-authoring.md lists the blacklist verbatim
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Superpowers No Placeholders rule.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:48Z"
---
