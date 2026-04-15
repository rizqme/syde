---
id: REQ-0183
kind: requirement
name: Storage Engine Supports Three Slug Forms
slug: storage-engine-supports-three-slug-forms-hlwy
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:30Z"
statement: When Get is called with a slug, the storage engine shall resolve full suffixed, bare-with-ambiguity-error, and parent-over-child path forms.
req_type: functional
priority: must
verification: unit test of Get in store.go covering all three slug forms
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Flexible slug addressing improves command ergonomics.
---
