---
id: REQ-0371
kind: requirement
name: Concept meaning shall be required
slug: concept-meaning-shall-be-required-suc1
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T08:25:33Z"
statement: The syde audit engine shall report an error for any concept entity with an empty meaning field.
req_type: functional
priority: must
verification: sync check errors on empty meaning
source: plan
requirement_status: active
rationale: Meaning is core
---
