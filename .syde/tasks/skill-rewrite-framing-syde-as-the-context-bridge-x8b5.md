---
acceptance: skill/SKILL.md and skill/references/commands.md updated. syde install-skill in a scratch dir lands the new files. The Long help for syde query mentions the three-question checklist and the Grep/Read prohibition for tracked files.
affected_entities:
    - skill-installer
    - cli-commands
affected_files:
    - skill/SKILL.md
    - skill/references/commands.md
    - internal/cli/query.go
completed_at: "2026-04-14T09:32:26Z"
created_at: "2026-04-14T09:20:18Z"
details: 'skill/SKILL.md: rename ''Finding Files to Read'' to ''Getting Context — syde first, always''. Lead with the three-question checklist. Document the architecture↔code sync feedback loop: every query that returns ''no owners'' or hits in an untracked file is the model telling you drift exists — fix it now. Explicit rule: do NOT use Grep or Read for any file tracked by the summary tree — use syde query --code, --file --content, or --search instead. Reserve Grep/Read for vendor/, node_modules/, generated assets, and .git/. skill/references/commands.md: rewrite cookbook section with 10+ recipes covering --code, --content, --kind listing, orphan triage, drift detection, broadened-search example, file→owner reverse lookup, recent-activity scoping. Update queryCmd.Long in internal/cli/query.go to match. Mention the three-question checklist.'
id: TSK-0056
kind: task
name: Skill rewrite framing syde as the context bridge
objective: Agents reading the skill default to syde for every read/explore step, understand the architecture/code sync feedback loop, and stop using Grep/Read for tracked files
plan_phase: phase_3
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: skill-rewrite-framing-syde-as-the-context-bridge-x8b5
task_status: completed
updated_at: "2026-04-14T09:32:26Z"
---
