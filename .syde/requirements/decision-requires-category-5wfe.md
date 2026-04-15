---
id: REQ-0019
kind: requirement
name: Decision Requires Category
slug: decision-requires-category-5wfe
relationships:
    - target: decision-m2um
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:47Z"
statement: The syde CLI shall require a category on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --category
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: Category enables constraint checks to filter decisions by domain.
---
