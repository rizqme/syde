---
id: TSK-0268
kind: task
name: Draft references/plan-review-prompt.md
slug: draft-referencesplan-review-promptmd-lrf1
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: syde-cli-shall-provide-plan-review-command-that-dispatches-a-plan-reviewer-subagent-6wli
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: medium
objective: skill/references/plan-review-prompt.md holds the reviewer-subagent prompt template derived from superpowers plan-document-reviewer-prompt — Completeness, Spec Alignment, Task Decomposition, Buildability, calibrated to flag only implementation-blocking issues.
details: 'Adapt superpowers plan-document-reviewer-prompt.md for syde''s plan model. Include: (a) four-row check table (Completeness, Spec Alignment, Task Decomposition, Buildability), (b) calibration paragraph (''only flag issues that would cause real problems during implementation''), (c) output format (Status Approved or Issues Found, bulleted Issues with plan-task-step references, advisory Recommendations). Tailor wording to reference syde-specific structures: phases, changes lane (Requirements NEW/Extended/Deleted), affected_entities, affected_files.'
acceptance: skill/references/plan-review-prompt.md bundled via go:embed; interpolation points for plan content clearly marked; output-format section is verbatim copy-pasteable.
affected_entities:
    - skill-installer-wbmu
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:24:42Z"
completed_at: "2026-04-18T09:34:35Z"
---
