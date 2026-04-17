---
id: FLW-0004
kind: flow
name: Drift Detection And Clearing
slug: drift-detection-and-clearing-ltya
description: How file mtime vs entity updated_at surfaces drift and how task-done clears it.
relationships:
    - target: syde
      type: belongs_to
    - target: validate-model
      type: involves
    - target: complete-task
      type: involves
    - target: tree-validator-shall-guarantee-zero-drift-5vxy
      type: references
updated_at: "2026-04-16T10:48:22Z"
trigger: syde validate runs, or the agent is reviewing an entity
goal: Ensure entities and their mapped source files are kept in sync, surfacing unreviewed drift to the agent
steps:
    - id: s1
      action: Agent runs syde validate
      contract: validate-model
      description: Checks all entities for drift
      on_success: s2
    - id: s2
      action: Validator compares file mtime vs entity updated_at
      description: Internal comparison logic
      on_success: s3
      on_failure: s4
    - id: s3
      action: No drift found — clean
      description: All files match
      on_success: done
    - id: s4
      action: 'WARN: file newer than entity'
      contract: complete-task
      description: Agent runs syde task done with --affected-entity
      on_success: done
narrative: 'Validator iterates every entity and compares the mtime of each file in its ''files'' list against the entity''s ''updated_at''. If file mtime > updated_at, a WARN is raised: ''file changed but these entities were not updated since''. Agent must update the entity (re-read the file, adjust description/responsibility/etc.) OR complete a task whose affected_entities or affected_files covers this pair, which auto-bumps updated_at. Validator also raises an ERROR for any non-ignored file that has zero owning entities — agent must either add the file to an entity''s files list or ''syde tree ignore'' it.'
happy_path: task done → touch affected entities → next validate run shows no WARN
edge_cases: Entity files field references a deleted file → validator errors. File was rewritten during a task without listing it as affected → drift WARN surfaces next run.
failure_modes: 'If agents routinely ignore drift warnings, the model decays. Mitigation: session-end hook runs ''syde tree status --strict'' and surfaces drift.'
---
