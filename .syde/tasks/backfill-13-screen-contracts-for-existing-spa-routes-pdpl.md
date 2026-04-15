---
acceptance: syde query --kind contract --format refs | grep -i screen lists 13 new contracts; dashboard Contracts page renders their wireframes.
affected_entities:
    - web-spa
completed_at: "2026-04-14T11:35:40Z"
created_at: "2026-04-14T11:23:14Z"
details: 'Write /tmp/syde-screens.sh containing 13 syde add contract calls. Use contract_kind=screen, interaction_pattern=render, input=<route path>, wireframe=<small UIML block matching the actual layout>, --add-rel ''web-spa:references'', --add-rel ''syded-dashboard:belongs_to'', --file pointing at the tsx source. Screens: Overview, File Tree, Graph, Plan View, Task Board, Learning Feed, Concepts ERD, Systems Inbox, Components Inbox, Contracts Inbox, Concepts Inbox (list mode), Flows Inbox, Decisions Inbox. Run the script and verify all 13 appear under the Contracts page.'
id: TSK-0079
kind: task
name: Backfill 13 screen contracts for existing SPA routes
objective: Every React page has a matching screen contract with a UIML wireframe
plan_phase: phase_4
plan_ref: screen-contract-kind
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: backfill-13-screen-contracts-for-existing-spa-routes-pdpl
task_status: completed
updated_at: "2026-04-14T11:35:40Z"
---
