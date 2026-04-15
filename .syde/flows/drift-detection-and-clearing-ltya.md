---
description: How file mtime vs entity updated_at surfaces drift and how task-done clears it.
edge_cases: Entity files field references a deleted file → validator errors. File was rewritten during a task without listing it as affected → drift WARN surfaces next run.
failure_modes: 'If agents routinely ignore drift warnings, the model decays. Mitigation: session-end hook runs ''syde tree status --strict'' and surfaces drift.'
goal: Ensure entities and their mapped source files are kept in sync, surfacing unreviewed drift to the agent
happy_path: task done → touch affected entities → next validate run shows no WARN
id: FLW-0004
kind: flow
name: Drift Detection And Clearing
narrative: 'Validator iterates every entity and compares the mtime of each file in its ''files'' list against the entity''s ''updated_at''. If file mtime > updated_at, a WARN is raised: ''file changed but these entities were not updated since''. Agent must update the entity (re-read the file, adjust description/responsibility/etc.) OR complete a task whose affected_entities or affected_files covers this pair, which auto-bumps updated_at. Validator also raises an ERROR for any non-ignored file that has zero owning entities — agent must either add the file to an entity''s files list or ''syde tree ignore'' it.'
relationships:
    - target: syde
      type: belongs_to
    - target: validate-model
      type: involves
    - target: complete-task
      type: involves
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: drift-detection-and-clearing-ltya
trigger: syde validate runs, or the agent is reviewing an entity
updated_at: "2026-04-14T03:27:02Z"
---
