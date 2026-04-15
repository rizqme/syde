---
acceptance: Final section is a copy-pasteable bash block that runs without edits; covers all 5 implementation tasks.
affected_entities:
    - uiml-parser
completed_at: "2026-04-15T02:50:16Z"
created_at: "2026-04-15T02:37:51Z"
details: 'Append ''## Implementation plan (next session)'' to uiml-survey.md containing a single bash code block. Tasks to include: (1) lexer fix with inTag flag + parser unaffected, (2) new uiml.RenderWireframeHTML function emitting the visual rules from phase 3, (3) FormatJSON rewire to call RenderWireframeHTML for contract_kind=screen, (4) re-backfill the 13 existing screen contracts using the patterns from phase 4, (5) skill docs cleanup dropping the attribute-free caveat. Each syde task create includes objective/details/acceptance/affected-entity. The block must run as-is in a future session.'
id: TSK-0085
kind: task
name: Draft the follow-up implementation plan inside uiml-survey.md
objective: uiml-survey.md ends with a runnable shell block that creates the implementation plan + phases + tasks via syde plan create
plan_phase: phase_5
plan_ref: uiml-wireframe-research
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: draft-the-follow-up-implementation-plan-inside-uiml-surveymd-jd8w
task_status: completed
updated_at: "2026-04-15T02:50:16Z"
---
