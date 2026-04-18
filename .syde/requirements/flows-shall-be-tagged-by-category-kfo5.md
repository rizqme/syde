---
id: REQ-0360
kind: requirement
name: Flows shall be tagged by category
slug: flows-shall-be-tagged-by-category-kfo5
description: Tags group flows for dashboard filtering
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:22Z"
statement: When creating a flow entity, the syde CLI shall support tagging flows by category for dashboard filtering.
req_type: functional
priority: should
verification: syde query --kind flow --tag planning returns only planning-tagged flows
source: plan
requirement_status: active
rationale: Tags are the lightweight grouping mechanism; no structural nesting needed
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:22Z"
---
