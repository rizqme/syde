---
id: FLW-0023
kind: flow
name: Browse Concepts
slug: browse-concepts-dpck
description: User navigates to the concepts page with list or ERD view
tags:
    - dashboard
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-dashboard-browsing-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User clicks Concepts in the sidebar
goal: User sees concept list or ERD canvas
steps:
    - id: s1
      action: User clicks Concepts in sidebar
      contract: concepts-inbox-screen
      description: Dashboard renders list view and opens a concept in detail
      on_success: done
---
