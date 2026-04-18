# Plan Authoring Discipline

This reference codifies the plan-quality rules for syde. It is derived from
[obra/superpowers `writing-plans`](https://github.com/obra/superpowers/blob/main/skills/writing-plans/SKILL.md)
and adapted to syde's phase / task / changes-lane model.

The rules here overlay the structural `syde plan create` → `add-phase` →
`add-change` → `task create` workflow. Follow both.

---

## 1. Audience

**Assume the executor has zero context for this codebase and questionable
taste.** The plan must be complete enough to hand to a fresh subagent or an
enthusiastic junior engineer with no prior session memory. Every path,
command, and acceptance signal must be spelled out.

If the plan reads like "you know what to do here", it fails this test.

---

## 2. Files section per phase (before tasks)

**Before declaring tasks in a phase, map out the files the phase will touch.**
This is where decomposition decisions get locked in.

Author the files list as part of the phase `--details` or as `--affected-file`
flags on the first task. The list is:

- **Create**: `exact/path/to/new/file.ext` — one-line responsibility
- **Modify**: `exact/path/to/existing/file.ext:LINE-RANGE` — what changes
- **Test**: `exact/path/to/test.ext` — what is covered

Rules:

- **Exact paths only.** `internal/cli/foo.go`, not "the CLI file".
- **One responsibility per file.** If a new file does two unrelated things,
  split it.
- **Files that change together live together.** Split by responsibility, not
  by technical layer.
- **Follow existing patterns.** Don't unilaterally restructure a file that
  the rest of the codebase treats as a single unit.

---

## 3. Bite-sized checkbox steps

Inside each task's `--details`, decompose the work into checkbox-prefixed
steps. Each step is **one 2-5 minute action**, not a multi-hour chunk.

**Canonical sequence for a code change with tests:**

- `[ ] Write the failing test` (with the actual test code in a fenced block)
- `[ ] Run the test to confirm it fails` (exact command + expected failure message)
- `[ ] Write the minimal implementation` (actual code in a fenced block)
- `[ ] Run the test to confirm it passes` (exact command + expected output)
- `[ ] Commit` (`git add ... && git commit -m "..."`)

**For non-code-change tasks** (docs, data migration, YAML edits), still use
checkbox steps; just adapt: `[ ] Open file X`, `[ ] Replace section Y with Z
(show the exact text)`, `[ ] Verify with command W (expected output: ...)`.

**Rules for steps:**

- **Exact commands.** `pytest tests/foo_test.py::test_bar -v`, not "run the
  test".
- **Expected output.** Every verification step states what the command should
  print. "Expected: PASS", "Expected: FAIL with 'X not defined'", etc.
- **Complete code blocks.** If a step changes code, show the code. Don't
  write "add handler for X" — write the handler.
- **Exact file paths** in every step that touches a file.

---

## 4. No Placeholders

These phrases are **plan failures**. The plan check and reviewer should both
flag them:

| Placeholder pattern | Fix |
|---|---|
| `TBD`, `TODO`, "fill in later", "implement later" | Write the actual content; if you can't, the plan isn't ready |
| "Add appropriate error handling" / "add validation" / "handle edge cases" | Specify which errors to handle and how |
| "Write tests for the above" (no code shown) | Include the actual test code |
| "Similar to Task N" (no content repeated) | Repeat the code — tasks may be read out of order or by different executors |
| "Refactor X as needed" | State the exact refactoring plan |
| Steps that describe instead of show | Include the code block / command |
| References to types / functions not defined in any task | Define them (or reference the file where they live) |

If a step legitimately doesn't fit these rules (e.g. a discovery task whose
outcome depends on a prior step), explicitly note that and make the
discovery itself the step's exact action.

---

## 5. Self-Review Checklist

**Run this yourself after drafting the complete plan. Not a subagent task —
five minutes with fresh eyes.**

1. **Spec / requirements coverage.** Walk the Requirements lane. Does every
   new requirement map to a task? Does every modified component/contract
   have an Extended change with both `what` and `why`? Any gaps?

2. **Placeholder scan.** Search the plan for the phrases in section 4. Every
   hit is a fix before approval.

3. **Cross-task consistency.** Do function / method / file names mentioned
   in later tasks match what earlier tasks create? A function `ClearLayers()`
   in Task 3 and `ClearFullLayers()` in Task 7 is a bug.

4. **Phase ordering.** If a later phase depends on an earlier phase's output,
   verify the dependency is explicit and correct. Catch "Phase 2 deletes X
   but Phase 3 reparents off X" — the reparent must happen first.

5. **Execution handoff prepared.** Plan ends with an explicit choice: fresh
   subagent per task vs inline execution (see section 7).

Fix issues inline. No need to re-review — just fix and move on.

---

## 6. Plan header discipline

Every plan should open with these four lines in `background` / `objective`
narrative (syde's `syde plan create` flags already map to this):

- **Goal** — one sentence on what the plan builds.
- **Architecture** — 2-3 sentences on the approach.
- **Tech stack / constraints** — key libraries, existing patterns to follow.
- **Executor** — "Task-by-task, with a plan-reviewer subagent between phases"
  or similar — tells the executor which flow to follow.

---

## 7. Execution Handoff

After the plan is saved and approved, **announce the execution choice
explicitly**:

> Plan approved and saved to
> `.syde/plans/<slug>.md`. Two execution options:
>
> 1. **Subagent-Driven** (recommended for multi-phase plans) — dispatch a
>    fresh subagent per task, review between tasks, keep the main context
>    clean.
> 2. **Inline Execution** — execute every task in the current session, mark
>    done as you go, surface findings at the end.
>
> Which approach?

Don't silently proceed inline — ask. The user's preference is a real
design decision (cost, context preservation, review frequency).

If the user has pre-approved inline execution in project memory (e.g. "do
not pause mid-plan"), that overrides this prompt — but only if the
preference was captured explicitly. Default is ask.

---

## 8. End-to-end example

Phase header + task, following all six rules:

### Phase: Add `syde foo` command

**Files:**
- Create: `internal/cli/foo.go` — cobra subcommand + flag wiring
- Modify: `internal/cli/root.go:42-48` — register fooCmd
- Test: `internal/cli/foo_test.go` — happy path + flag-missing error

### Task 1: Wire the cobra command

**Files:**
- Create: `internal/cli/foo.go`
- Modify: `internal/cli/root.go:42-48`

- [ ] **Step 1: Write the failing test**

```go
// internal/cli/foo_test.go
package cli

import "testing"

func TestFoo_RequiresArg(t *testing.T) {
    err := runFoo([]string{})
    if err == nil || err.Error() != "foo: missing required <name>" {
        t.Fatalf("expected missing-arg error, got %v", err)
    }
}
```

- [ ] **Step 2: Run to confirm fail**

Run: `go test ./internal/cli/ -run TestFoo_RequiresArg -v`
Expected: FAIL with `undefined: runFoo`

- [ ] **Step 3: Write minimal implementation**

```go
// internal/cli/foo.go
package cli

import (
    "fmt"
    "github.com/spf13/cobra"
)

var fooCmd = &cobra.Command{
    Use:   "foo <name>",
    Short: "Foo the bar",
    Args:  cobra.ExactArgs(1),
    RunE: func(cmd *cobra.Command, args []string) error {
        return runFoo(args)
    },
}

func runFoo(args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("foo: missing required <name>")
    }
    fmt.Printf("foo: %s\n", args[0])
    return nil
}
```

- [ ] **Step 4: Register command** (modify `internal/cli/root.go:42-48`)

Add `rootCmd.AddCommand(fooCmd)` in `init()`.

- [ ] **Step 5: Run to confirm pass**

Run: `go test ./internal/cli/ -run TestFoo -v`
Expected: PASS (1 test, 0 failures)

- [ ] **Step 6: Commit**

```bash
git add internal/cli/foo.go internal/cli/foo_test.go internal/cli/root.go
git commit -m "feat(cli): add syde foo command"
```

---

## Remember

- Exact paths always
- Complete code in every step where code changes
- Exact commands with expected output
- No placeholders — every phrase from section 4 is a fail
- Self-review before presenting
- Announce the execution handoff choice explicitly
