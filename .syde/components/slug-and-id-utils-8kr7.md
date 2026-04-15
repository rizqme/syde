---
id: COM-0013
kind: component
name: Slug and ID Utils
slug: slug-and-id-utils-8kr7
description: Slug normalization, suffix generation, and legacy ID helpers.
purpose: Shared slug generation and legacy ID helpers
files:
    - internal/utils/slug.go
    - internal/utils/id.go
relationships:
    - target: syde-cli
      type: belongs_to
updated_at: "2026-04-15T10:15:40Z"
responsibility: Provide Slugify, SlugifyWithSuffix, BaseSlug, HasSuffix, and a deprecated ID stub
capabilities:
    - Normalize names into URL-safe slugs
    - Append a 4-char random alnum suffix for uniqueness
    - Strip suffix to recover the bare-name form for ambiguity resolution
boundaries: Does NOT allocate counter-based IDs (storage/counters.go owns that).
---
