---
id: REQ-0383
kind: requirement
name: Requirement update shall support audited flag
slug: requirement-update-shall-support-audited-flag-qfs2
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:21Z"
statement: The syde update command shall accept a repeatable --audited flag to add acknowledged overlaps to an existing requirement.
req_type: functional
priority: must
verification: syde update --audited <slug> persists the overlap on the requirement.
source: plan
requirement_status: active
rationale: Audit acknowledgement must be possible post-creation.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:21Z"
---
