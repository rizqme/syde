---
boundaries: Does NOT allocate counter-based IDs (storage/counters.go owns that).
capabilities:
    - Normalize names into URL-safe slugs
    - Append a 4-char random alnum suffix for uniqueness
    - Strip suffix to recover the bare-name form for ambiguity resolution
description: Slug normalization, suffix generation, and legacy ID helpers.
files:
    - internal/utils/slug.go
    - internal/utils/id.go
id: COM-0013
kind: component
name: Slug and ID Utils
purpose: Shared slug generation and legacy ID helpers
relationships:
    - target: syde-cli
      type: belongs_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
responsibility: Provide Slugify, SlugifyWithSuffix, BaseSlug, HasSuffix, and a deprecated ID stub
slug: slug-and-id-utils-8kr7
updated_at: "2026-04-14T03:35:54Z"
---
