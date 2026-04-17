---
id: REQ-0384
kind: requirement
name: Overlap detection shall be bidirectional
slug: overlap-detection-shall-be-bidirectional-jdh0
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: The syde sync check overlap rule shall require both requirements in an overlap pair to acknowledge each other before clearing the finding.
req_type: functional
priority: must
verification: sync check still errors if only one side of the pair acknowledges the other.
source: plan
requirement_status: active
rationale: Mutual acknowledgement is stronger than unilateral ignore.
---
