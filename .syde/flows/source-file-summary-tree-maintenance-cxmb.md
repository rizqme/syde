---
description: Leaves-first summarize loop with subagent delegation that keeps tree.yaml fresh.
edge_cases: Binary files and >1 MiB files get auto-summaries and never become stale from content. Deleted files prune from tree and mark parent stale. Ignored nodes are skipped by status --strict. Parallel tree.yaml writers race on the atomic rename — subagents must retry sequentially on collision.
failure_modes: Race condition on .syde/tree.yaml.tmp when multiple writers run concurrently → some summaries lost → caught by the next 'syde tree status' showing a straggler stale. SessionEnd hook exits non-zero if the tree isn't clean, blocking the session from ending quietly.
goal: Keep .syde/tree.yaml in lock-step with the source tree — every file and folder has a current human-written summary
happy_path: scan → stale list → subagent batch → main session folder pass → strict status OK
id: FLW-0003
kind: flow
name: Source File Summary Tree Maintenance
narrative: 'Agent runs ''syde tree scan''. syde walks the project (respecting .gitignore + defaults + config tree_ignore), hashes files, diffs against tree.yaml. Changed files have their hash updated and summary_stale set, cascading stale up to the root. Agent lists stale leaves via ''syde tree changes --leaves-only --format json''. Agent dispatches subagents in parallel, each given a batch of stale file paths. Each subagent loops: ''syde tree context <path>'' → compose summary → ''syde tree summarize <path> --summary ...''. After subagents return, agent summarizes stale folders in the main session using ''syde tree show <folder>'' to see children. Final check: ''syde tree status --strict'' must exit 0.'
performance_notes: Subagent dispatch is essential — reading 100+ files into the main session blows the token budget. Folder summaries are cheap (derive from already-stored children summaries).
relationships:
    - target: syde
      type: belongs_to
    - target: summary-tree
      type: involves
    - target: tree-node
      type: involves
    - target: scan-tree
      type: involves
slug: source-file-summary-tree-maintenance-cxmb
trigger: Session start, or any source file write during a session, or explicit 'syde tree scan'
updated_at: "2026-04-14T03:27:02Z"
---
