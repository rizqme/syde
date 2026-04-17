---
id: REQ-0377
kind: requirement
name: ERD canvas shall be removed
slug: erd-canvas-shall-be-removed-4j8u
relationships:
    - target: web-spa
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T08:40:04Z"
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
---
