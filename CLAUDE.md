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
  scan/                  Codebase scan guide + coverage
  skill/                 SKILL.md templates, hooks.json, installer
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
- **SKILL.md is the behavior lever**: The embedded SKILL.md in `skill/templates.go` controls how Claude Code interacts with syde. Changes here directly affect agent behavior.
- **Hooks drive automation**: `skill/hooks.go` defines PostToolUse and SessionStart hooks. These inject syde context into Claude's system messages.
- **FileRef pointers**: BadgerDB values point to file paths + line numbers in the markdown files. No entity data is duplicated in the index.

## Adding a new entity field

1. Add field to the struct in `internal/model/entity.go` (or plan.go, task.go, etc.) with `yaml` tag
2. Add CLI flag in `internal/cli/add.go` and `internal/cli/update.go`
3. Add to batch struct in `internal/cli/batch.go`
4. Update `internal/skill/templates.go` EntitySpecRef to document it
5. Build and test: `go build ./cmd/syde/ && syde add <kind> <name> --new-flag "value"`

## Adding a new CLI command

1. Create `internal/cli/newcommand.go`
2. Define cobra command with `Use`, `Short`, `Args`, `RunE`
3. Register in `init()`: `rootCmd.AddCommand(newCmd)` or `parentCmd.AddCommand(newCmd)`
4. Update CommandsRef in `internal/skill/templates.go`

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

- `.claude/skills/syde/SKILL.md` — skill definition (from `skill/templates.go` SkillMD)
- `.claude/skills/syde/references/*.md` — entity spec, commands, plan workflow, constraints, scan workflow
- `.claude/hooks/syde-hooks.json` — PostToolUse + SessionStart hooks
- `CLAUDE.md` — appends mandatory syde rules section (idempotent)
