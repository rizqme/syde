---
id: REQ-0476
kind: requirement
name: 'Approved plan: Bidirectional requirement-component coupling with content-hash rechec...'
slug: approved-plan-bidirectional-requirement-component-coupling-with-content-hash-rechec-sgfy
relationships:
    - target: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
      type: references
      label: approved_plan
    - target: skill-installer-wbmu
      type: refines
    - target: cli-commands-hpjb
      type: refines
    - target: requirement-itw0
      type: relates_to
updated_at: "2026-04-18T10:04:45Z"
statement: The syde design model shall treat every active requirement as refining at least one component, with no system-level requirement targets and a content-hash recheck gate that marks a requirement stale whenever any file of its refining components diverges from the stored verified_against snapshot.
req_type: functional
priority: must
verification: syde sync check reports zero findings from requirement_no_component, component_no_requirement, requirement_targets_system, and requirement_stale rules after running syde requirement verify across all active requirements.
source: plan
source_ref: plan:bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Captured automatically when the plan was approved; restates the plan objective as a ubiquitous EARS statement enforcing the bidirectional invariant and the content-hash gate.
approved_at: "2026-04-18T08:00:32Z"
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T10:04:45Z"
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:45Z"
---
