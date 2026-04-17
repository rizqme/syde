---
id: FLW-0028
kind: flow
name: View File Tree
slug: view-file-tree-ulqa
description: User browses the source file tree in the dashboard
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks File Tree in the sidebar
goal: User sees the project file tree with summaries and stale indicators
steps:
    - id: s1
      action: User clicks File Tree in sidebar
      contract: file-tree-screen
      description: Dashboard renders tree view
      on_success: done
---
