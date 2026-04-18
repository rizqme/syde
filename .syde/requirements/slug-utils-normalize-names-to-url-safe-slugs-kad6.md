---
id: REQ-0171
kind: requirement
name: Slug Utils Normalize Names To URL Safe Slugs
slug: slug-utils-normalize-names-to-url-safe-slugs-kad6
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
updated_at: "2026-04-18T09:37:30Z"
statement: When Slugify is called with a name, the slug utils shall return a lowercase URL-safe slug that strips punctuation and whitespace.
req_type: functional
priority: must
verification: unit test of Slugify in internal/utils/slug.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Slugs are the primary addressing scheme for entity files.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:37:30Z"
---
