---
id: REQ-0173
kind: requirement
name: Slug Utils Strip Suffix To Recover Base
slug: slug-utils-strip-suffix-to-recover-base-ycwi
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
updated_at: "2026-04-18T09:36:55Z"
statement: When BaseSlug is called on a suffixed slug, the slug utils shall return the bare name form without the four-character suffix.
req_type: functional
priority: must
verification: unit test of BaseSlug in slug.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Stripping the suffix enables ambiguity resolution in commands that accept bare names.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:36:55Z"
---
