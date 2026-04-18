---
id: REQ-0375
kind: requirement
name: Concepts shall use used-in for flows
slug: concepts-shall-use-used-in-for-flows-cyol
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:37:18Z"
statement: The syde entity model shall support a used_in relationship type from concept entities to flow entities.
req_type: functional
priority: must
verification: --add-rel works
source: plan
requirement_status: active
rationale: Explicit roles
audited_overlaps:
    - slug: concepts-shall-use-exposed-via-for-contracts
      distinction: used-in points concept→flow for journey participation; exposed-via points concept→contract for boundary exposure — different role types
    - slug: concepts-shall-use-implemented-by-for-components
      distinction: used-in points concept→flow for journey participation; implemented-by points concept→component for code realisation — different role types
    - slug: concepts-shall-use-exposed-via-for-contracts-r0la
      distinction: Defines the used_in edge from concept to flow entities; the exposed_via requirement is a different relationship name pointing to contract entities.
    - slug: concepts-shall-use-implemented-by-for-components-1arr
      distinction: Defines the used_in edge pointing to flow entities; the implemented_by requirement is a distinct edge name pointing to component entities.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:18Z"
---
