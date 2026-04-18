---
id: REQ-0172
kind: requirement
name: Slug Utils Append Four Char Suffix
slug: slug-utils-append-four-char-suffix-hxad
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
updated_at: "2026-04-18T09:38:01Z"
statement: When SlugifyWithSuffix is called, the slug utils shall append a four-character random alphanumeric suffix to the base slug.
req_type: functional
priority: must
verification: unit test of SlugifyWithSuffix in slug.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: The suffix guarantees unique filenames even when two entities share a base name.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:38:01Z"
---
