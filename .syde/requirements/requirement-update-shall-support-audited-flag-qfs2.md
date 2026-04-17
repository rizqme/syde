---
id: REQ-0383
kind: requirement
name: Requirement update shall support audited flag
slug: requirement-update-shall-support-audited-flag-qfs2
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: The syde update command shall accept a repeatable --audited flag to add acknowledged overlaps to an existing requirement.
req_type: functional
priority: must
verification: syde update --audited <slug> persists the overlap on the requirement.
source: plan
requirement_status: active
rationale: Audit acknowledgement must be possible post-creation.
---
