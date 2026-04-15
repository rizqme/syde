---
id: REQ-0172
kind: requirement
name: Slug Utils Append Four Char Suffix
slug: slug-utils-append-four-char-suffix-hxad
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:57Z"
statement: When SlugifyWithSuffix is called, the slug utils shall append a four-character random alphanumeric suffix to the base slug.
req_type: functional
priority: must
verification: unit test of SlugifyWithSuffix in slug.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: The suffix guarantees unique filenames even when two entities share a base name.
---
