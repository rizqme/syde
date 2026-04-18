---
id: REQ-0490
kind: requirement
name: Skill documentation shall describe systems as standalone processes
slug: skill-documentation-shall-describe-systems-as-standalone-processes-rsf8
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:48Z"
statement: The syde skill documentation shall describe system entities as standalone top-level process entities that must not carry belongs_to relationships and that may be referenced by components via one or more belongs_to edges.
req_type: functional
priority: should
verification: skill/SKILL.md system rules section and skill/references/entity-spec.md system section both describe the no-root standalone-process model
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Agents learn the model from these files.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:48Z"
---
