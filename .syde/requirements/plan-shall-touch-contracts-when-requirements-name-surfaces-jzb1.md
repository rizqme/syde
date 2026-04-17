---
id: REQ-0399
kind: requirement
name: Plan shall touch contracts when requirements name surfaces
slug: plan-shall-touch-contracts-when-requirements-name-surfaces-jzb1
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When a syde plan's change diff introduces or extends a requirement whose statement names a CLI command, REST endpoint, screen, or event surface, the same plan's change diff shall introduce or extend a contract whose input covers that surface.
req_type: functional
priority: must
verification: syde plan check warns on any plan whose requirement lane mentions a surface not covered by the plan's contract lane
source: plan
requirement_status: active
rationale: Shift-left the requirement-contract coverage gate.
---
