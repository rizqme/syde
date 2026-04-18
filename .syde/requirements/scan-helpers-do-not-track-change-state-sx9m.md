---
id: REQ-0164
kind: requirement
name: Scan Helpers Do Not Track Change State
slug: scan-helpers-do-not-track-change-state-sx9m
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
updated_at: "2026-04-18T09:36:45Z"
statement: The scan helpers shall not track file change state and shall leave drift detection to the summary tree.
req_type: constraint
priority: must
verification: code review of internal/scan for change state storage
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Change tracking is the summary tree's responsibility; scan helpers are stateless.
verified_against:
    scan-helpers-legacy-sa6d:
        hash: c9b19fd18480f13b89908a38caa95ff8ea66c15c73dc6cf6e551e414e71ce3ba
        at: "2026-04-18T09:36:45Z"
---
