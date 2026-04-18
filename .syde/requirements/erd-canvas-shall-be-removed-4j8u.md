---
id: REQ-0377
kind: requirement
name: ERD canvas shall be removed
slug: erd-canvas-shall-be-removed-4j8u
relationships:
    - target: web-spa
      type: refines
updated_at: "2026-04-18T09:37:16Z"
statement: The dashboard shall not render an ERD canvas for concept entities.
req_type: constraint
priority: must
verification: No ERD toggle on concepts page
source: plan
requirement_status: active
rationale: Without attributes there is nothing to render
supersedes:
    - concepts-erd-node-click-returns-to-list-cbne
    - concepts-erd-route-renders-canvas-up23
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:16Z"
---
