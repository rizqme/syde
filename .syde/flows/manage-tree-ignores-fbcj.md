---
id: FLW-0035
kind: flow
name: Manage Tree Ignores
slug: manage-tree-ignores-fbcj
description: User marks files as ignored or un-ignores them
tags:
    - tree
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User wants to exclude a file from orphan validation
goal: File is marked ignored or un-ignored in tree.yaml
steps:
    - id: s1
      action: User runs syde tree ignore
      contract: ignore-tree-node
      description: Marks file as ignored
      on_success: done
      on_failure: s2
    - id: s2
      action: User runs syde tree unignore
      contract: unignore-tree-node
      description: Removes ignore flag
      on_success: done
---
