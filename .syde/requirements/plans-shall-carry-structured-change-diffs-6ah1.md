---
id: REQ-0331
kind: requirement
name: Plans shall carry structured change diffs
slug: plans-shall-carry-structured-change-diffs-6ah1
relationships:
    - target: plans-shall-reference-entities-never-draft-them-ibdv
      type: supersedes
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:21Z"
statement: The syde plan model shall record every entity change as a structured diff with deleted, extended, and new entries, and the completion validator shall verify the declared diff against actual entity state before a plan can be marked completed.
req_type: constraint
priority: must
verification: Integration test creating a plan with an extended FieldChanges entry, executing the tasks, then asserting syde plan complete blocks on a deliberately mismatched value.
source: manual
source_ref: plan:revamp-planning-to-structured-design-and-diff-m8p5
requirement_status: active
rationale: Without an embedded diff reviewers cannot approve plans with full context and the system cannot verify what actually landed.
supersedes:
    - plans-shall-reference-entities-never-draft-them-ibdv
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:21Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:21Z"
---
