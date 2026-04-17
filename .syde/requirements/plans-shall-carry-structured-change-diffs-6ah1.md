---
id: REQ-0331
kind: requirement
name: Plans shall carry structured change diffs
slug: plans-shall-carry-structured-change-diffs-6ah1
relationships:
    - target: plans-shall-reference-entities-never-draft-them-ibdv
      type: supersedes
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-16T01:33:47Z"
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
---
