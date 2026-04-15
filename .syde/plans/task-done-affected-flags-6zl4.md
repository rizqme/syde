---
approved_at: "2026-04-14T07:26:57Z"
background: 'syde task done today inherits whatever --affected-entity / --affected-file the caller set at task create time. In practice the set is almost always wrong: Claude declares 2 affected entities upfront, then during implementation touches 8 more, and at done time the drift cascade (touchAffectedEntities) only bumps UpdatedAt on the original 2. Sync check then flags the other 6 as drift warnings even though the task actually covered them. The skill teaches Claude to predict affected entities at create time, but prediction is unreliable — the accurate information is only available at done time.'
completed_at: "2026-04-14T07:33:12Z"
created_at: "2026-04-14T07:23:49Z"
id: PLN-0003
kind: plan
name: Task Done Affected Flags
objective: syde task done <slug> accepts --affected-entity and --affected-file (repeatable, same grammar as task create). Additions are MERGED into the task's existing lists, validated, persisted to disk, and then fed to touchAffectedEntities so drift cascades clear the full set. SKILL.md teaches Claude to always pass the real affected set at done time. After this change, running a full task loop (create → start → code → done with correct flags) leaves sync check clean.
phases:
    - changes: New cobra flags on task done; merge-into-task before the status flip; validateTaskReferences on the merged set; persist via writeClient.Update; touchAffectedEntities runs against the merged task
      description: Accept affected entities and files at done time, merge into the task, validate, persist
      details: 'In internal/cli/task.go around the taskDoneCmd (or wherever task done is wired), register the two flags with variables scoped to the command (not reusing the package-level taskAffectedEntities which is owned by task create). On RunE: load the task via store.Get; append the flag values to t.AffectedEntities / t.AffectedFiles, dedupe; call validateTaskReferences(store, newAffectedEntities, newAffectedFiles) on just the new additions (existing ones already validated at create/previous done); if validation fails, abort before persisting. Then set TaskStatus=completed + CompletedAt=now; store.Update(t, body); THEN call touchAffectedEntities(store, t) which will see the merged set.'
      id: phase_1
      name: Wire --affected-* on task done
      notes: Dedupe is on the merged list, not the additions alone — if someone passes --affected-entity foo twice we collapse it but keep any pre-existing 'foo'. touchAffectedEntities already handles duplicates via its toTouch map so no change there.
      objective: syde task done <slug> --affected-entity X --affected-file Y updates the stored task AND clears drift for the merged set
      status: completed
      tasks:
        - add-affected-entity-affected-file-flags-to-task-done
        - merge-flag-values-into-the-loaded-task-validate
        - persist-merged-task-before-touch-cascade
    - changes: skill/SKILL.md Phase 3 'AFTER — complete and verify' step updated; Rule 4 clarified; one smoke test run
      description: Teach Claude when to use the new flags + verify end-to-end
      details: 'Edit the ''AFTER'' section of Phase 3: Implement so step 8 reads ''syde task done <slug> --affected-entity ... --affected-file ...'' and spells out that this is how drift gets cleared. Rule 4 gets a pointer to Phase 3 step 8. Install the updated skill. Run a fake task loop (create with 1 affected, code 3 files, done with all 4 affected) and confirm sync check passes.'
      id: phase_2
      name: Skill guidance + smoke test
      notes: No new entities needed — skill changes land in skill/SKILL.md and the installer copies them verbatim
      objective: SKILL.md says 'always re-declare affected set at done time'; a real task loop leaves sync check clean
      status: completed
      tasks:
        - update-skillmd-guidance
        - end-to-end-smoke-test
    - changes: New Store.UpdateCascade / CreateCascade / DeleteCascade methods in internal/storage/store.go; server write handlers (POST/PUT/DELETE /entity) call the cascade variants; cascade traversal uses a map[entityID]bool visited set keyed by ID (stable across slug aliases) so a cycle terminates after at most one visit per node.
      description: When a child entity is written (Create/Update/Delete), recursively bump its belongs_to parents' UpdatedAt — cycle-safe via a visited-ID set
      details: 'Implementation outline: add s.cascadeFromParent(childBase, visited) helper that walks childBase.Relationships, filters type==belongs_to, resolves each target via s.Get (handles all three slug forms), checks visited[parent.ID] and skips if already processed, otherwise marks visited[parent.ID]=true then calls s.Update(parent, parentBody) and recurses. Create calls cascadeFromParent(newEntity.GetBase(), seedVisited). Update calls cascadeFromParent(updatedEntity.GetBase(), seedVisited). Delete reads the entity first to capture its belongs_to, then after the filesystem delete calls cascadeFromParent(deletedEntity.GetBase(), seedVisited). seedVisited starts with the child''s own ID so the child itself can''t be re-entered if a parent somehow refers back. Broken belongs_to targets (Get fails) are silently skipped — the validator will surface them separately.'
      id: phase_3
      name: Belongs-to cascade on every entity write
      notes: Cycle safety comes from the visited-ID map not from static analysis. Even a self-loop (A belongs_to A) terminates on the first recurse because A.ID is in visited. The validator already rejects cycles but we must not crash on repos where the validator hasn't been run yet.
      objective: Any store.Update / store.Create / store.Delete cascades through the belongs_to chain so a child change marks every ancestor fresh. No infinite loops even if validator-rejected cycles sneak in.
      status: completed
      tasks:
        - add-cascadefromparent-helper-updatecascade-on-store
        - cascade-on-create-and-delete-paths
        - wire-server-entity-writedelete-handlers-to-cascade-variants
        - cascade-smoke-test
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'IN: new --affected-entity and --affected-file flags on ''syde task done''; merge-into-existing semantics (no replacement, no dedupe of the pre-existing list, union of old + new); validation via the existing validateTaskReferences helper; persistence via the existing writeClient.Update round-trip; SKILL.md Phase 3 and Rule 4 updates so Claude always declares the real affected set at completion; one smoke test. OUT: a separate ''syde task update'' command (deferred); adding flags to task start (the set isn''t known yet at start); auto-detecting affected files from git diff (nice but out of scope — Claude will type them); removing affected-entity/file from task create (create-time prediction is still useful as a plan marker).'
slug: task-done-affected-flags-6zl4
source: manual
updated_at: "2026-04-14T07:33:12Z"
---
