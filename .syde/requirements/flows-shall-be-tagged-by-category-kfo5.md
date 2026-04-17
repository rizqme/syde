---
id: REQ-0360
kind: requirement
name: Flows shall be tagged by category
slug: flows-shall-be-tagged-by-category-kfo5
description: Tags group flows for dashboard filtering
relationships:
    - target: syde
      type: belongs_to
updated_at: "2026-04-16T10:40:59Z"
statement: When creating a flow entity, the syde CLI shall support tagging flows by category for dashboard filtering.
req_type: functional
priority: should
verification: syde query --kind flow --tag planning returns only planning-tagged flows
source: plan
requirement_status: active
rationale: Tags are the lightweight grouping mechanism; no structural nesting needed
---
