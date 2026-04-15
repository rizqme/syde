
## syde Design Model

This project uses syde for architecture management. These Codex rules are mandatory:

1. **Start with syde context**: Run `syde tree scan` and ensure `syde tree status --strict` passes before planning or editing.
2. **Use syde query before raw reads**: Prefer `syde query --file <path> --content`, `syde query --code <symbol>`, and `syde query --search "<term>"` for tracked files.
3. **Design before code**: Create a syde plan, add phases/entities, present it, wait for approval, then run `syde plan approve <slug>`.
4. **Track implementation**: Start a task with `syde task start <slug>` before changing files and finish with `syde task done <slug> --affected-entity <entity> --affected-file <path>`.
5. **Verify mappings**: After source edits, run `syde constraints check <file>` and map new files with `syde update <component> --file <path>`.
6. **Finish cleanly**: Run tests, `syde sync check --strict`, and refresh stale summary-tree nodes before final response.
7. **Hook limitation**: Codex hooks currently intercept Bash only. They are guardrails, not a complete enforcement boundary for `apply_patch` or other non-Bash tools.
8. **Never read `.syde/` files directly** — always use syde CLI commands.
