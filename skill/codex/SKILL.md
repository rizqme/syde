---
name: syde
description: >
  Use whenever a repository has a .syde/ design model, or when the user asks
  about architecture, design constraints, implementation planning, tasks,
  codebase synchronization, or syde commands. Keeps Codex aligned with
  syde's text-first architecture model.
---

# syde for Codex

This repository uses syde as the source of truth for architecture and work
tracking. syde stores design entities in `.syde/` and exposes them through the
`syde` CLI and the `syded` daemon.

## First Moves

1. Run `syde tree scan`.
2. If `syde tree status --strict` fails, clean the summary tree before planning:
   use `syde tree changes --leaves-only`, `syde tree context <path>`, and
   `syde tree summarize <path> --summary "..."` until strict status passes.
3. Use `syde query` for architectural context before raw file reads:
   `syde query --file <path> --content`, `syde query --code <symbol>`, or
   `syde query --search "<term>"`.

## Design Before Code

The Codex `UserPromptSubmit` hook records each user prompt as a
`requirement` entity. Treat those records as append-only intent: if a later
prompt conflicts, create a new requirement and mark the older one superseded
or obsolete instead of deleting it.

## Requirements (EARS)

Requirement statements MUST match one of the five EARS patterns. The save
validator rejects any non-conforming statement:

- Ubiquitous: `The <subject> shall <action>.`
- Event-driven: `When <trigger>, the <subject> shall <action>.`
- State-driven: `While <state>, the <subject> shall <action>.`
- Unwanted-behavior: `If <unwanted condition>, then the <subject> shall <action>.`
- Optional-feature: `Where <feature is included>, the <subject> shall <action>.`

`syde add requirement` requires `--statement`, `--type`
(`functional|non-functional|constraint|interface|performance|security|usability`),
`--priority` (MoSCoW: `must|should|could|wont`), `--verification` (short
free-form note on how fulfillment is verified), `--source`, and `--rationale`.
The legacy `--acceptance` flag is gone for requirements — use `--verification`.

Requirements never carry a `files` list. They link only via `refines`
(requirement → component/contract/concept/system, or req → req) and
`derives_from` (requirement → parent requirement), plus the usual
`belongs_to`.

Backfilling requirements from an existing codebase follows the deterministic
algorithm in `references/requirement-derivation.md`.

Before changing source files:

1. Clarify requirements and assumptions with the user through the available
   ask-user-question tool. In Codex Plan mode use `request_user_input`; in
   Default mode or runtimes without that tool, ask plainly in chat and wait.
2. Decompose the request into **granular requirements** — one EARS statement
   per behavioral property, constraint, audit rule, or UI expectation. Search
   existing requirements with `syde query --kind requirement` /
   `syde query --search`; supersede or obsolete conflicting intent. A plan
   with many phases but few requirements is incomplete — each design decision
   must trace to a requirement. Add all requirements to the plan's
   Requirements lane before any implementation lane. Cascade:
   Requirements → Components → Contracts → Concepts → Flows. If behavior
   changes, extend the relevant flow with `field_changes`.
3. Create a syde plan with `syde plan create <name> --design "..."`.
   The `--design` flag captures the detailed implementation prose that
   reviewers read before approving.
4. Declare every entity change the plan will make using
   `syde plan add-change`. Three subcommands, one per diff type:
   - `syde plan add-change delete <plan> <kind> <slug> --why "..."`
   - `syde plan add-change extend <plan> <kind> <slug> --what "..." --why "..." [--field key=value]`
     (the optional repeatable `--field key=value` declares a programmatically
     verified field-level diff; the sentinel value `DELETE` means "clear this
     field")
   - `syde plan add-change new <plan> <kind> --name "..." --what "..." --why "..." --draft key=value`
     (repeatable `--draft` flags carry kind-specific fields; JSON literal
     values like `[{"path":"x","type":"string"}]` auto-decode)
   Inspect the diff with `syde plan show-changes <plan>` and remove an entry
   with `syde plan remove-change <plan> <change-id>`. Once tasks exist, every
   change must list its implementing task slug(s) with repeatable
   `--task <task-slug>` flags.
5. Add phases with `syde plan add-phase`.
6. Create a comprehensive, granular task list before approval:
   `syde task create "<task>" --plan <slug> --phase <phase-id>`. Every phase
   must have at least one concrete task; split broad tasks until each one has a
   clear file/entity target or verification outcome.
7. Run `syde plan check <plan>` after drafting. Address every ERROR and review
   every WARN. Do not present the plan for approval until `syde plan check`
   exits 0; warnings may be acknowledged in the presentation.
8. Run `syde plan open <plan>` before asking for approval so the dashboard plan
   detail is visible in an existing or newly opened browser tab.
9. Mirror those syde tasks into Codex `update_plan` before showing the plan.
   Keep statuses synced whenever a syde task is started, completed, split,
   renamed, or blocked.
10. Show the plan and task list (the dashboard plan detail page at
   `/<project>/plan/<slug>` renders Design + structured Changes for review),
   then wait for user approval.
11. Run `syde plan approve <slug>`.
12. Start tasks with `syde task start` as implementation begins.
13. Link every new or changed entity back to the relevant requirement
    with an outbound relationship such as
    `--add-rel <requirement>:references`.

Plans now embed structured drafts. The old "plans shall reference entities,
never draft them" rule (REQ-0007) was reversed by REQ-0331: plans MUST
record every entity change as a structured diff so reviewers can see the
delta before approval and the system can verify it after execution.

Do not rely on Codex hooks as a complete enforcement boundary. Codex hooks
currently intercept Bash, not every file-editing tool. You still need to follow
this workflow when using `apply_patch` or other non-Bash tools.

## During Implementation

- Prefer `apply_patch` for source edits.
- After writing a new or changed source file, verify design ownership with
  `syde constraints check <path>`.
- If a file is unmapped, attach it to the owning component with
  `syde update <component> --file <path>` or ignore intentional non-design
  files through the summary tree.
- **Do not stop until every task in every phase of the approved plan is
  done.** Plan approval is a commitment to finish, not to check in. Do not
  pause for permission between tasks or phases. Stop only on a real blocker,
  on a discovered plan flaw that needs revision, or on user interrupt — never
  to ask "should I continue?".

## Plan Completion Gate

When all tasks in all phases are done, mark the plan completed with
`syde plan complete <plan-slug>`. This invokes `planCompletionFindings`
which compares every declared change against actual entity state:

- Deleted change still present → ERROR
- New change missing from the model → ERROR
- Extended change with `field_changes` whose value doesn't match the
  current entity → ERROR (per field)
- Extended change without `field_changes` → WARN (hand review only)

`syde plan complete` blocks on any ERROR. Use `--force` only after
deliberate manual review. The same findings surface inside
`syde sync check --strict` so the strict gate covers in-flight plans
too.

## Finish Gate

Before reporting done:

1. Run tests/builds relevant to the change.
2. Run `syde sync check --strict`.
3. Refresh the summary tree with `syde tree scan`, then summarize stale nodes
   until `syde tree status --strict` passes.
4. Mark tasks done with affected entities/files:
   `syde task done <slug> --affected-entity <entity> --affected-file <path>`.

## References

- `references/entity-spec.md` documents entity kinds and fields.
- `references/commands.md` documents syde CLI commands.
- `references/clarify-guide.md` helps with requirement questions.
- `references/sync-workflow.md` describes bootstrapping an existing codebase.
- `references/requirement-derivation.md` — deterministic EARS backfill algorithm.
