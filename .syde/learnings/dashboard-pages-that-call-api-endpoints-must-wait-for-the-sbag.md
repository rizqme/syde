---
category: gotcha
confidence: high
description: 'Dashboard pages that call api.* endpoints must wait for the async fetchProjects() effect to set the module-level projectSlug via setProject(). Gate the render on App''s ''ready'' state (like status does), otherwise a fresh URL load races against setProject and the fetch hits /api/null/<endpoint> → 404 → error state. Symptom: empty/error screen on reload that disappears after navigating away and back.'
discovered_at: "2026-04-13T14:05:01Z"
entity_refs:
    - web-spa
id: LRN-0002
kind: learning
name: 'Dashboard pages that call api.* endpoints must wait for the '
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dashboard-pages-that-call-api-endpoints-must-wait-for-the-sbag
source: session-observation
updated_at: "2026-04-13T14:05:01Z"
---
