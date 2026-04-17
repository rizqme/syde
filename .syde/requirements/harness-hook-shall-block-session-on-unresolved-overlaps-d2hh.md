---
id: REQ-0397
kind: requirement
name: Harness hook shall block session on unresolved overlaps
slug: harness-hook-shall-block-session-on-unresolved-overlaps-d2hh
relationships:
    - target: skill-installer
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When syde add requirement prints an overlap banner, the installed Claude Code PostToolUse hook shall emit a system reminder naming the merge, rename, and distinct resolution paths before the session continues.
req_type: functional
priority: must
verification: after syde install-skill --all, a syde add requirement call with a high-overlap statement triggers a visible system reminder in the session transcript
source: plan
requirement_status: active
rationale: The skill infrastructure must enforce the semantic review, not just document it.
---
