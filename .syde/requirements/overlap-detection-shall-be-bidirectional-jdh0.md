---
id: REQ-0384
kind: requirement
name: Overlap detection shall be bidirectional
slug: overlap-detection-shall-be-bidirectional-jdh0
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:53Z"
statement: The syde sync check overlap rule shall require both requirements in an overlap pair to acknowledge each other before clearing the finding.
req_type: functional
priority: must
verification: sync check still errors if only one side of the pair acknowledges the other.
source: plan
requirement_status: active
rationale: Mutual acknowledgement is stronger than unilateral ignore.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:53Z"
---
