---
id: FLW-0003
kind: flow
name: Source File Summary Tree Maintenance
slug: source-file-summary-tree-maintenance-cxmb
description: Leaves-first summarize loop with subagent delegation that keeps tree.yaml fresh.
relationships:
    - target: syde
      type: belongs_to
    - target: summary-tree
      type: involves
    - target: tree-node
      type: involves
    - target: scan-tree
      type: involves
    - target: tree-summaries-shall-cascade-stale-to-root-yqxa
      type: references
updated_at: "2026-04-16T10:59:47Z"
trigger: Session start, or any source file write during a session, or explicit 'syde tree scan'
goal: Keep .syde/tree.yaml in lock-step with the source tree — every file and folder has a current human-written summary
steps:
    - id: s1
      action: Agent runs syde tree scan
      contract: scan-tree
      description: Walks filesystem, marks stale
      on_success: s2
    - id: s2
      action: Agent lists stale leaves
      contract: list-tree-changes
      description: Returns paths needing summaries
      on_success: s3
    - id: s3
      action: Subagents read file context
      contract: tree-context-bundle
      description: Each reads breadcrumb + content
      on_success: s4
    - id: s4
      action: Subagents write summaries
      contract: set-tree-summary
      description: Each stores summary
      on_success: s5
    - id: s5
      action: Agent summarizes folders
      contract: show-tree
      description: Reads children summaries
      on_success: s6
    - id: s6
      action: Agent checks tree status
      contract: tree-status
      description: Must exit 0
      on_success: done
      on_failure: s2
narrative: 'Agent runs ''syde tree scan''. syde walks the project (respecting .gitignore + defaults + config tree_ignore), hashes files, diffs against tree.yaml. Changed files have their hash updated and summary_stale set, cascading stale up to the root. Agent lists stale leaves via ''syde tree changes --leaves-only --format json''. Agent dispatches subagents in parallel, each given a batch of stale file paths. Each subagent loops: ''syde tree context <path>'' → compose summary → ''syde tree summarize <path> --summary ...''. After subagents return, agent summarizes stale folders in the main session using ''syde tree show <folder>'' to see children. Final check: ''syde tree status --strict'' must exit 0.'
happy_path: scan → stale list → subagent batch → main session folder pass → strict status OK
edge_cases: Binary files and >1 MiB files get auto-summaries and never become stale from content. Deleted files prune from tree and mark parent stale. Ignored nodes are skipped by status --strict. Parallel tree.yaml writers race on the atomic rename — subagents must retry sequentially on collision.
failure_modes: Race condition on .syde/tree.yaml.tmp when multiple writers run concurrently → some summaries lost → caught by the next 'syde tree status' showing a straggler stale. SessionEnd hook exits non-zero if the tree isn't clean, blocking the session from ending quietly.
performance_notes: Subagent dispatch is essential — reading 100+ files into the main session blows the token budget. Folder summaries are cheap (derive from already-stored children summaries).
---
