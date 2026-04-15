---
id: REQ-0175
kind: requirement
name: Slug Utils Do Not Allocate Counter IDs
slug: slug-utils-do-not-allocate-counter-ids-ykyh
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:05Z"
statement: The slug utils shall not allocate counter-based entity IDs and shall defer allocation to the storage engine counter module.
req_type: constraint
priority: must
verification: code review of internal/utils for counter writes
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Counter allocation requires BadgerDB access that is owned by storage.
---
