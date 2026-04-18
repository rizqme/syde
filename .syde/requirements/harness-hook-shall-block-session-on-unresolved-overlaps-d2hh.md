---
id: REQ-0397
kind: requirement
name: Harness hook shall block session on unresolved overlaps
slug: harness-hook-shall-block-session-on-unresolved-overlaps-d2hh
relationships:
    - target: skill-installer
      type: refines
updated_at: "2026-04-18T10:04:46Z"
statement: When syde add requirement prints an overlap banner, the installed Claude Code PostToolUse hook shall emit a system reminder naming the merge, rename, and distinct resolution paths before the session continues.
req_type: functional
priority: must
verification: after syde install-skill --all, a syde add requirement call with a high-overlap statement triggers a visible system reminder in the session transcript
source: plan
requirement_status: active
rationale: The skill infrastructure must enforce the semantic review, not just document it.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:46Z"
---
