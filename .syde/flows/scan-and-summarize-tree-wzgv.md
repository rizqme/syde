---
id: FLW-0034
kind: flow
name: Scan and Summarize Tree
slug: scan-and-summarize-tree-wzgv
description: Agent scans the source tree and summarizes stale files
tags:
    - tree
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: Session start or source files changed
goal: Every file and folder in tree.yaml has a current summary
steps:
    - id: s1
      action: Agent runs syde tree scan
      contract: scan-tree
      description: Walks fs, marks changed files stale
      on_success: s2
    - id: s2
      action: Agent lists stale leaves
      contract: list-tree-changes
      description: Returns stale file paths
      on_success: s3
    - id: s3
      action: Agent reads file context
      contract: tree-context-bundle
      description: Returns breadcrumb + content
      on_success: s4
    - id: s4
      action: Agent writes summary
      contract: set-tree-summary
      description: Stores summary on node
      on_success: s5
    - id: s5
      action: Agent checks tree status
      contract: tree-status
      description: Reports stale count
      on_success: s2
      on_failure: done
    - id: s6
      action: Agent views tree
      contract: show-tree
      description: Shows tree hierarchy
      on_success: done
    - id: s7
      action: Agent reads node summary
      contract: get-tree-summary
      description: Returns summary for path
      on_success: done
    - id: s8
      action: Dashboard reads tree node
      contract: get-tree-node-http
      description: HTTP API returns node detail
      on_success: done
    - id: s9
      action: Dashboard reads full tree
      contract: get-full-tree-http
      description: HTTP API returns full tree
      on_success: done
---
