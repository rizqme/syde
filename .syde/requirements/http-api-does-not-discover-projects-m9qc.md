---
id: REQ-0049
kind: requirement
name: HTTP API Does Not Discover Projects
slug: http-api-does-not-discover-projects-m9qc
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:18Z"
statement: The syded HTTP API shall not discover projects on the filesystem by itself.
req_type: constraint
priority: must
verification: inspection confirming project lookup is delegated to Project Registry
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: Separating discovery from routing keeps handlers thin and testable.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:18Z"
---
