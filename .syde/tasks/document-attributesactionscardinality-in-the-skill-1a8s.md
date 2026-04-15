---
acceptance: Fresh install-skill lands the updated files. A reader can copy the Order example and it parses + validates.
affected_entities:
    - skill-installer
affected_files:
    - skill/SKILL.md
    - skill/references/entity-spec.md
    - skill/references/commands.md
completed_at: "2026-04-14T10:00:17Z"
created_at: "2026-04-14T09:44:26Z"
details: skill/SKILL.md Concept rules rewritten with the new required fields (meaning, >=1 attribute) and a worked example using Order/LineItem/Customer. Add cardinality enum table. skill/references/entity-spec.md Concept section gains attribute/action fields and cardinality documentation. skill/references/commands.md syde add flag reference gains --attribute and --action with pipe-syntax notes.
id: TSK-0061
kind: task
name: Document attributes/actions/cardinality in the skill
objective: Agents reading the skill create fully-populated ERD concepts on the first try
plan_phase: phase_5
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: document-attributesactionscardinality-in-the-skill-1a8s
task_status: completed
updated_at: "2026-04-14T10:00:17Z"
---
