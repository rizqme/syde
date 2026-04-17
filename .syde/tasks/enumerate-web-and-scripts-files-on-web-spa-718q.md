---
id: TSK-0099
kind: task
name: Enumerate web/ and scripts/ files on web-spa
slug: enumerate-web-and-scripts-files-on-web-spa-718q
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-clear-all-remaining-sync-check-drift
      type: references
updated_at: "2026-04-17T10:46:19Z"
task_status: completed
objective: web-spa-jy9z's file list covers every tracked web/** source and config plus scripts/wireframe-shot.sh
details: Gather the current file list via syde query web-spa --full; append the 29 orphan paths from syde sync check output; re-invoke syde update web-spa --file <...> with the full union (flag replaces the list)
acceptance: syde sync check reports 0 orphan errors under web/ or scripts/
affected_entities:
    - web-spa-jy9z
plan_ref: clear-all-remaining-sync-check-drift-aokb
plan_phase: phase_1
created_at: "2026-04-17T08:48:21Z"
completed_at: "2026-04-17T09:05:21Z"
---
