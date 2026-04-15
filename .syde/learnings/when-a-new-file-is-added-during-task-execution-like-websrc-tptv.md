---
category: gotcha
confidence: medium
description: 'When a new file is added during task execution (like web/src/pages/ERD.tsx), syde task create rejects --affected-file for a path that doesn''t exist in the tree yet. Workaround: omit --affected-file, mention the new file in --details, and at task-done time run syde update <owning-component> --file <new-path> to attach it. The full file list must be re-passed because --file replaces rather than appends.'
discovered_at: "2026-04-14T10:12:28Z"
entity_refs:
    - cli-commands
id: LRN-0015
kind: learning
name: When a new file is added during task execution (like web/src
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: when-a-new-file-is-added-during-task-execution-like-websrc-tptv
source: session-observation
updated_at: "2026-04-14T10:12:28Z"
---
