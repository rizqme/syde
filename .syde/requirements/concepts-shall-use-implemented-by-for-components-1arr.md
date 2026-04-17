---
id: REQ-0373
kind: requirement
name: Concepts shall use implemented-by for components
slug: concepts-shall-use-implemented-by-for-components-1arr
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:50:15Z"
statement: The syde entity model shall support an implemented_by relationship type from concept entities to component entities.
req_type: functional
priority: must
verification: --add-rel works
source: plan
requirement_status: active
rationale: Explicit roles
audited_overlaps:
    - slug: concepts-shall-use-exposed-via-for-contracts
      distinction: implemented-by points concept→component for code realisation; exposed-via points concept→contract for boundary exposure — different role types
    - slug: concepts-shall-use-used-in-for-flows
      distinction: implemented-by points concept→component for code realisation; used-in points concept→flow for journey participation — different role types
    - slug: concepts-shall-use-exposed-via-for-contracts-r0la
      distinction: Defines the implemented_by edge from concept to component entities; the exposed_via requirement defines a different edge name pointing to contract entities.
    - slug: concepts-shall-use-used-in-for-flows-cyol
      distinction: Defines the implemented_by edge targeting component entities; the used_in requirement is a distinct relationship name pointing to flow entities.
---
