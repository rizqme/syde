---
id: REQ-0183
kind: requirement
name: Storage Engine Supports Three Slug Forms
slug: storage-engine-supports-three-slug-forms-hlwy
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:39Z"
statement: When Get is called with a slug, the storage engine shall resolve full suffixed, bare-with-ambiguity-error, and parent-over-child path forms.
req_type: functional
priority: must
verification: unit test of Get in store.go covering all three slug forms
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Flexible slug addressing improves command ergonomics.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:39Z"
---
