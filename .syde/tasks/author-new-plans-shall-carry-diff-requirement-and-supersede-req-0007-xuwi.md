---
id: TSK-0018
kind: task
name: Author new plans-shall-carry-diff requirement and supersede REQ-0007
slug: author-new-plans-shall-carry-diff-requirement-and-supersede-req-0007-xuwi
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: A new EARS requirement declares the new plan-diff rule; REQ-0007 (Plans shall reference entities never draft them) is marked superseded with superseded_by pointing at the new one.
details: syde add requirement 'Plans shall carry structured change diffs' --statement 'The syde plan model shall record every entity change a plan will make as a structured diff with deleted, extended, and new entries, and the completion validator shall verify the declared diff against actual entity state before a plan can be marked completed.' --type constraint --priority must --verification 'Integration test creating a plan with a New component draft, executing the tasks, and asserting syde plan complete rejects a mismatched FieldChanges value.' --rationale 'Without an embedded diff reviewers cannot approve plans with full context and the system cannot verify what actually landed.' --add-rel 'plans-shall-reference-entities-never-draft-them-ibdv:supersedes' --add-rel 'syde-5tdt:belongs_to'. Then syde update plans-shall-reference-entities-never-draft-them-ibdv --requirement-status superseded --superseded-by <new-slug>.
acceptance: syde query plans-shall-reference-entities-never-draft-them-ibdv shows status=superseded and superseded_by link.
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_5
created_at: "2026-04-15T11:42:14Z"
completed_at: "2026-04-15T12:20:26Z"
---
