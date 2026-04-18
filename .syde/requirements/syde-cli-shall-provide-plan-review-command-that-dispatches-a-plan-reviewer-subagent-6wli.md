---
id: REQ-0494
kind: requirement
name: syde CLI shall provide plan review command that dispatches a plan-reviewer subagent
slug: syde-cli-shall-provide-plan-review-command-that-dispatches-a-plan-reviewer-subagent-6wli
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:55:16Z"
statement: When the user runs syde plan review against a plan slug, the syde CLI shall dispatch a plan-reviewer subagent whose prompt is bundled with the skill, and shall print the reviewer's verdict as either Approved or Issues Found along with any blocking issues listed by the reviewer.
req_type: functional
priority: should
verification: Running syde plan review against a plan with a deliberate placeholder step surfaces the placeholder as a blocking issue
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Structural syde plan check does not catch spec-alignment or task-decomposition issues.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:55:16Z"
---
