---
acceptance: bash scripts/wireframe-shot.sh components-inbox-screen writes a non-empty PNG that the Read tool can display. .syde/research/ marked ignored in the tree.
affected_entities:
    - web-spa
completed_at: "2026-04-15T02:39:58Z"
created_at: "2026-04-15T02:37:51Z"
details: 'Bash script: usage ''wireframe-shot.sh <slug> [out.png]''. Detect project slug by curling http://localhost:5703/api/projects and matching path == $(pwd). Wait-for-port loop on :5703 with curl, auto-launch syde open if dead. Call /Applications/Google Chrome.app/Contents/MacOS/Google\ Chrome --headless --disable-gpu --window-size=1440,900 --virtual-time-budget=5000 --screenshot=<out>. Default out: /tmp/wireframe-<slug>.png. After writing the script, run syde tree ignore .syde/research so the orphan gate stays clean for the survey artifact in phase 2.'
id: TSK-0081
kind: task
name: Write scripts/wireframe-shot.sh + tree ignore .syde/research
objective: scripts/wireframe-shot.sh works end to end and .syde/research/ is tree-ignored so research artifacts don't trip the orphan/sync gates
plan_phase: phase_1
plan_ref: uiml-wireframe-research
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: write-scriptswireframe-shotsh-tree-ignore-syderesearch-5kz1
task_status: completed
updated_at: "2026-04-15T02:39:58Z"
---
