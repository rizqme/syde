---
id: REQ-0395
kind: requirement
name: Audited overlap acknowledgements shall carry a distinction rationale
slug: audited-overlap-acknowledgements-shall-carry-a-distinction-rationale-yruo
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T11:07:56Z"
statement: The syde requirement entity shall persist, for every audited overlap entry, a distinction rationale of at least 20 characters explaining why the two requirements are semantically distinct.
req_type: functional
priority: must
verification: sync check errors on any audited_overlaps entry with empty or below-threshold distinction
source: plan
requirement_status: superseded
rationale: TF-IDF text overlap is only a candidate signal; semantic distinction must be authored explicitly.
superseded_by:
    - acknowledged-requirement-overlaps-shall-carry-non-trivial-distinction-text
    - acknowledged-requirement-overlaps-shall-carry-non-trivial-distinction-text-5wqs
---
