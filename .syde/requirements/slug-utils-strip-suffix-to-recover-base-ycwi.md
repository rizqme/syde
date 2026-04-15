---
id: REQ-0173
kind: requirement
name: Slug Utils Strip Suffix To Recover Base
slug: slug-utils-strip-suffix-to-recover-base-ycwi
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:00Z"
statement: When BaseSlug is called on a suffixed slug, the slug utils shall return the bare name form without the four-character suffix.
req_type: functional
priority: must
verification: unit test of BaseSlug in slug.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Stripping the suffix enables ambiguity resolution in commands that accept bare names.
---
