---
id: REQ-0402
kind: requirement
name: Requirement statements mentioning CLI REST screen or event surfaces shall map to contracts
slug: requirement-statements-mentioning-cli-rest-screen-or-event-surfaces-shall-map-to-contracts-81c4
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When an active requirement's prose names a CLI invocation, HTTP route, dashboard screen, or pub-sub topic, the syde design model shall hold an active contract whose input definition covers the surface.
req_type: constraint
priority: must
verification: zero sync check findings about requirement prose surfaces lacking contract coverage
source: plan
requirement_status: active
rationale: Prevents requirement-contract drift by forcing the surface to exist in both places.
---
