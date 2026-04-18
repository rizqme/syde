---
id: TSK-0269
kind: task
name: Build + reinstall to bundle the two new reference docs
slug: build-reinstall-to-bundle-the-two-new-reference-docs-oofa
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: skill-documentation-shall-describe-systems-as-standalone-processes-rsf8
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: medium
objective: go build ./... succeeds with references/plan-authoring.md and references/plan-review-prompt.md embedded; make install copies them into .claude/skills/syde/references/ and .agents/skills/syde/references/.
acceptance: After make install, both files exist under .claude/skills/syde/references/ and are nonzero-length.
affected_entities:
    - skill-installer-wbmu
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:24:42Z"
completed_at: "2026-04-18T09:36:31Z"
---
