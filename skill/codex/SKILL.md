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
2. Create a syde plan with `syde plan create`.
3. Add planned entities and phases with `syde plan add-entity` and
   `syde plan add-phase`.
4. Create a comprehensive, granular task list before approval:
   `syde task create "<task>" --plan <slug> --phase <phase-id>`. Every phase
   must have at least one concrete task; split broad tasks until each one has a
   clear file/entity target or verification outcome.
5. Mirror those syde tasks into Codex `update_plan` before showing the plan.
   Keep statuses synced whenever a syde task is started, completed, split,
   renamed, or blocked.
6. Show the plan and task list, then wait for user approval.
7. Run `syde plan approve <slug>`; approval creates a plan-sourced
   requirement and links it to the plan.
8. Start tasks with `syde task start` as implementation begins.
9. Link every new or changed entity back to the relevant requirement
   with an outbound relationship such as
   `--add-rel <requirement>:references`.

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
