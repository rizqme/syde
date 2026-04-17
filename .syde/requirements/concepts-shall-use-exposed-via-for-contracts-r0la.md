---
id: REQ-0374
kind: requirement
name: Concepts shall use exposed-via for contracts
slug: concepts-shall-use-exposed-via-for-contracts-r0la
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:50:15Z"
statement: The syde entity model shall support an exposed_via relationship type from concept entities to contract entities.
req_type: functional
priority: must
verification: --add-rel works
source: plan
requirement_status: active
rationale: Explicit roles
audited_overlaps:
    - slug: concepts-shall-use-implemented-by-for-components
      distinction: exposed-via points concept→contract for boundary exposure; implemented-by points concept→component for code realisation — different role types
    - slug: concepts-shall-use-used-in-for-flows
      distinction: exposed-via points concept→contract for boundary exposure; used-in points concept→flow for journey participation — different role types
    - slug: concepts-shall-use-implemented-by-for-components-1arr
      distinction: Defines the exposed_via edge from concept to contract entities; the implemented_by requirement defines a different edge type targeting component entities.
    - slug: concepts-shall-use-used-in-for-flows-cyol
      distinction: Defines the exposed_via edge from concept to contract; the used_in requirement defines a different relationship name pointing to flow entities.
---
