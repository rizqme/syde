---
id: REQ-0484
kind: requirement
name: PostToolUse hook shall surface affected requirements when a file mapped to a component is edited
slug: posttooluse-hook-shall-surface-affected-requirements-when-a-file-mapped-to-a-component-is-edited-ofeu
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:48Z"
statement: When a PostToolUse Edit or Write touches a path mapped to a component files list, the syde skill hook shall surface the affected active requirements into the agent context.
req_type: functional
priority: should
verification: Editing a file owned by a component causes the next agent turn to receive a context block listing the active requirements refining that component
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Hook is no-op for paths owned by no component (e.g. .syde/, scripts/, docs).
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:48Z"
---
