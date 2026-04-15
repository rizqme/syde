---
acceptance: syde sync check passes; spot-screenshotting at least Components Inbox + Overview + Concepts ERD shows wireframes that match their respective patterns.
affected_entities:
    - web-spa
completed_at: "2026-04-15T03:08:44Z"
created_at: "2026-04-15T02:53:51Z"
details: Write /tmp/syde-rebackfill-screens.sh containing 13 syde update calls. Each updates the --wireframe field to the appropriate pattern from research section 7. Run the script. Verify with syde sync check + a few spot screenshots via scripts/wireframe-shot.sh.
id: TSK-0089
kind: task
name: Re-backfill 13 screen contracts with attribute-rich UIML
objective: All 13 screen contracts have wireframes that render correctly under the new renderer
plan_phase: phase_4
plan_ref: uiml-wireframe-render
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: re-backfill-13-screen-contracts-with-attribute-rich-uiml-h2tt
task_status: completed
updated_at: "2026-04-15T03:08:44Z"
---
