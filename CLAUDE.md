# syde — Development Guide

This is the syde project itself (the CLI tool + dashboard + skill). Go module at `github.com/feedloop/syde`.

## Project structure

```
cmd/syde/main.go         CLI entry point
cmd/syded/main.go        Dashboard daemon entry point
internal/
  cli/                   All cobra commands (add, get, list, update, query, plan, task, etc.)
  model/                 Entity types: BaseEntity + SystemEntity, ComponentEntity, etc.
  storage/               FileStore (markdown read/write) + Index (BadgerDB) + Store (unified)
  config/                syde.yaml config loading
  query/                 Rich query engine (resolver, formatter, diff)
  graph/                 Relationship traversal (BFS, impact analysis) + DOT/ASCII rendering
  memory/                Claude Code memory file generation
  scan/                  Codebase sync guide + coverage
  skill/                 SKILL.md, hooks.json, references/*.md (embedded via go:embed)
  uiml/                  UIML parser (lexer, parser, AST, ASCII/HTML renderers, validator)
  dashboard/             HTTP server, API handlers, project registry, embedded SPA
  utils/                 ID generation (ULID-style), slugify
```

## Build

```bash
go build -o syde ./cmd/syde/
go build -o syded ./cmd/syded/
```

## Key design decisions

- **Dual storage**: Markdown files are source of truth (git-friendly). BadgerDB index is a rebuildable cache for fast queries. Never trust the index over the files — `syde reindex` rebuilds it.
- **Entity kind dispatch**: `model.NewEntityForKind()` creates typed entities. CLI commands type-switch on the entity to access kind-specific fields. Keep these in sync when adding fields.
- **SKILL.md is the behavior lever**: The skill files in `skill/` (SKILL.md, references/*.md, hooks.json) are embedded via `go:embed` in `skill/embed.go`. Edit the actual files, not Go strings. Changes here directly affect agent behavior.
- **Hooks drive automation**: `skill/hooks.go` defines PostToolUse and SessionStart hooks. These inject syde context into Claude's system messages.
- **FileRef pointers**: BadgerDB values point to file paths + line numbers in the markdown files. No entity data is duplicated in the index.

## Adding a new entity field

1. Add field to the struct in `internal/model/entity.go` (or plan.go, task.go, etc.) with `yaml` tag
2. Add CLI flag in `internal/cli/add.go` and `internal/cli/update.go`
3. Add to `applyEntityData()` in `internal/cli/plan.go` for plan execute support
4. Update `skill/references/entity-spec.md` to document it
5. Build and test: `go build ./cmd/syde/ && syde add <kind> <name> --new-flag "value"`

## Adding a new CLI command

1. Create `internal/cli/newcommand.go`
2. Define cobra command with `Use`, `Short`, `Args`, `RunE`
3. Register in `init()`: `rootCmd.AddCommand(newCmd)` or `parentCmd.AddCommand(newCmd)`
4. Update `skill/references/commands.md`

## Testing

```bash
# Quick build + smoke test
go build -o syde ./cmd/syde/ && \
  rm -rf /tmp/syde-smoke && mkdir /tmp/syde-smoke && cd /tmp/syde-smoke && \
  syde init && syde add component "Test" --description "test" && syde status

# Full integration test with Claude Code
cd /tmp/syde-smoke && syde init --install-skill && \
  claude --print --output-format stream-json --verbose \
    --max-turns 100 --max-budget-usd 10 \
    -p "Build a todo app following the syde workflow" \
    < /dev/null > output.jsonl 2>&1
```

## What syde install-skill writes

- `.claude/skills/syde/SKILL.md` — skill definition (from `skill/SKILL.md`)
- `.claude/skills/syde/references/*.md` — entity spec, commands, clarify guide, sync workflow
- `.claude/hooks/syde-hooks.json` — PostToolUse + SessionStart hooks
- `CLAUDE.md` — appends mandatory syde rules section (idempotent)

## syde Design Model

This project uses syde for architecture management. These rules are mandatory:

1. **Phase 0 — always run first**: `syde tree scan`, then iterate `syde tree changes --leaves-only` + `syde tree summarize <path> --summary "..."` until `syde tree status --strict` exits 0. The summary tree is the cheap way to understand the project without re-reading every file. `.gitignore` is honored automatically.
2. **Architecture auto-loaded**: `syde context` runs at session start. Do NOT re-run it. Use `syde query <slug> --full` for targeted deep dives only.
3. **Use `syde tree context <path>`, NOT naive `Read`, when creating entities on an existing codebase.** It returns the ancestor breadcrumb + file summary + content in one call — the right framing for `--purpose` / `--responsibility` / `--boundaries` and for picking `belongs_to`.
4. **Clarify first**: Be critical — challenge assumptions, identify missing requirements, propose constraints. Use the project-type checklists in the syde skill. Wait for user confirmation before proceeding.
5. **Design before code**: Create a plan with `syde plan create`, add entity drafts with `syde plan add-entity`, add phases with `syde plan add-phase`. Present to user. Do NOT implement until approved.
6. **Track implementation**: Use `syde task create` / `syde task start` / `syde task done` for each unit of work.
7. **Verify after writing files**: Run `syde constraints check <file>` to verify new files are mapped to components.
8. **Finish**: Run `syde validate`, then **refresh the summary tree**: `syde tree scan` + leaves-first summarize loop until `syde tree status --strict` exits 0. The Stop hook will block the session from ending cleanly if the tree is dirty.
9. **Never read .syde/ files directly** — always use syde CLI commands.
