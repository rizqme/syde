---
acceptance: syded runs; syde validate runs while syded is up, no lock error; dashboard updates reflect CLI changes within ~1s
affected_entities:
    - project-registry
    - dashboard-daemon-entry-point
affected_files:
    - internal/dashboard/registry.go
completed_at: "2026-04-14T06:24:45Z"
created_at: "2026-04-14T06:03:22Z"
details: registry.GetStore returns memstore; spawn fsnotify watcher on .syde/ dir; on any .md change, rebuild memstore (full rebuild acceptable for v1). Remove badger.Open from syded code path.
id: TSK-0006
kind: task
name: Switch syded registry to memstore + fsnotify rebuild
objective: syded never opens badger; CLI and daemon coexist
plan_phase: phase_4
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: switch-syded-registry-to-memstore-fsnotify-rebuild-sxgm
task_status: completed
updated_at: "2026-04-14T06:24:45Z"
---
