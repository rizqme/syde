---
id: REQ-0495
kind: requirement
name: Plan approval step shall offer execution handoff choice
slug: plan-approval-step-shall-offer-execution-handoff-choice-mez3
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:47Z"
statement: When a plan transitions to approved, the syde skill shall prompt the user to choose between subagent-driven execution and inline execution before any task starts.
req_type: functional
priority: could
verification: skill/SKILL.md documents the handoff prompt verbatim
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Superpowers-derived alignment checkpoint.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:47Z"
---
