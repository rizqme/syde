---
id: COM-0002
kind: component
name: CLI Commands
slug: cli-commands-hpjb
description: Cobra command tree exposing every syde feature as a CLI subcommand.
purpose: Expose every syde feature as a cobra subcommand
notes:
    - 'Verification cleanup: scan completeness audit output now avoids fmt.Println with an embedded trailing newline so go vet passes under go test ./....'
files:
    - cmd/syde/main.go
    - internal/cli/add.go
    - internal/cli/codex_hook.go
    - internal/cli/constraints.go
    - internal/cli/context.go
    - internal/cli/design.go
    - internal/cli/files.go
    - internal/cli/get.go
    - internal/cli/graph.go
    - internal/cli/helpers.go
    - internal/cli/init.go
    - internal/cli/install_skill.go
    - internal/cli/list.go
    - internal/cli/memory.go
    - internal/cli/open.go
    - internal/cli/output.go
    - internal/cli/plan.go
    - internal/cli/query.go
    - internal/cli/reindex.go
    - internal/cli/remember.go
    - internal/cli/remove.go
    - internal/cli/root.go
    - internal/cli/scan.go
    - internal/cli/search.go
    - internal/cli/status.go
    - internal/cli/sync_check.go
    - internal/cli/task.go
    - internal/cli/tree.go
    - internal/cli/update.go
    - internal/cli/validate.go
    - internal/cli/wireframe.go
    - internal/cli/writes.go
    - internal/cli/requirements.go
relationships:
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T07:05:30Z"
responsibility: Define and register all syde CLI commands under the root cobra tree
capabilities:
    - Entity CRUD (add, get, list, update, remove)
    - Query and search (query, search, graph)
    - Plan and task lifecycle (plan create/approve/execute, task create/start/done)
    - Summary tree management (tree scan/summarize/show/context/status)
    - Validation + constraints enforcement
    - Init + skill install + project bootstrap
boundaries: Does NOT own entity persistence (delegates to storage). Does NOT render HTML or serve HTTP (syded owns that). Does NOT call LLMs.
---
