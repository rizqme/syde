---
id: REQ-0221
kind: requirement
name: Contracts Inbox Kind And Pattern Filters
slug: contracts-inbox-kind-and-pattern-filters-fjzp
relationships:
    - target: contracts-inbox-screen-x2tr
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:09Z"
statement: Where a filter query parameter with contract_kind or pattern terms is provided on the /contract route, the dashboard shall restrict the contracts list to entities matching those terms.
req_type: interface
priority: should
verification: manual inspection of /contract?filter=contract_kind:screen in the dashboard
source: manual
source_ref: contract:contracts-inbox-screen-x2tr
requirement_status: active
rationale: Kind and pattern filters let users slice the contract inventory by surface type.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:09Z"
---
