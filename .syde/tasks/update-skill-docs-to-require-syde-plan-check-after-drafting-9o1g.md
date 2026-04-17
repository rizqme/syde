---
id: TSK-0039
kind: task
name: Update skill docs to require syde plan check after drafting
slug: update-skill-docs-to-require-syde-plan-check-after-drafting-9o1g
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: skill/SKILL.md and skill/codex/SKILL.md require agents to run syde plan check after drafting any plan and revise until clean before presenting for approval.
details: |-
    skill/SKILL.md and skill/codex/SKILL.md require agents to perform three new mandatory rules in the planning workflow:

    1. **Requirement first, then cascade.** When the user requests anything (a feature, a fix, a UX change, a behavior tweak), the agent's first step is to identify or author the underlying requirement, NOT to jump to the implementation. Search existing requirements for semantic overlap via 'syde search' / 'syde query --kind requirement'. If a relevant requirement exists and the new request supersedes or conflicts with it, mark the old one superseded (status=superseded, superseded_by=<new>) or obsolete (status=obsolete, obsolete_reason=...). Add the new requirement (or the extended existing one) to the plan's Requirements lane FIRST. THEN cascade the implementation as Component/Contract/Concept/Flow changes that refine the requirement. The cascade order is: Requirements → Components → Contracts → Concepts → Flows. Requirements are the why; the other lanes are the how. Never invert that order. Always check existing flows: if the request changes a behavior captured by a flow entity, that flow MUST be Extended in the plan with field_changes declaring the new happy_path/narrative/edge_cases — flows rot silently when authors forget them.

    2. Run 'syde plan check <plan>' after drafting. Address every ERROR and review every WARN. Do NOT present the plan for approval until 'syde plan check' exits 0 (errors only — warnings may be acknowledged in the presentation).

    3. Run 'syde plan open <plan>' before asking for approval. This reuses the dashboard tab if one is connected (via WebSocket navigate broadcast) or spawns a new tab if not. The plan must be visible in the browser before the agent asks the user to approve.

    Add all three rules to SKILL.md Phase 2 CREATE PLAN. The cascade rule (#1) goes at the TOP of the phase since it shapes the whole authoring loop. Rules 2 and 3 are post-drafting gates and go immediately before the 'Show the plan for approval' step. Add the same three rules to codex SKILL.md plan workflow in the compact format.
acceptance: rg 'syde plan check' skill/ returns matches in SKILL.md and codex/SKILL.md.
affected_entities:
    - skill-installer-wbmu
affected_files:
    - skill/SKILL.md
    - skill/codex/SKILL.md
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_6
created_at: "2026-04-15T13:15:56Z"
completed_at: "2026-04-15T21:43:25Z"
---
