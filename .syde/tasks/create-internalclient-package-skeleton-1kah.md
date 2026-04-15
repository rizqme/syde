---
acceptance: Client.Status() against a live syded returns parsed StatusResponse
completed_at: "2026-04-14T06:46:49Z"
created_at: "2026-04-14T06:38:16Z"
details: 'NEW FILE internal/client/client.go: Client{base url, http.Client}. Methods: Get, List, Query, Status, Validate, SyncCheck, Context, Constraints, Reindex. Each marshals query params, calls json.Decode on response. Client.New(sydeDir) derives project slug from MakeProjectSlug.'
id: TSK-0020
kind: task
name: Create internal/client package skeleton
objective: A Client struct with every syded endpoint wrapped as a typed method
plan_phase: phase_3
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: create-internalclient-package-skeleton-1kah
task_status: completed
updated_at: "2026-04-14T06:46:49Z"
---
