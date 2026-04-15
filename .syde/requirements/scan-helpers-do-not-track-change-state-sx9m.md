---
id: REQ-0164
kind: requirement
name: Scan Helpers Do Not Track Change State
slug: scan-helpers-do-not-track-change-state-sx9m
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:30Z"
statement: The scan helpers shall not track file change state and shall leave drift detection to the summary tree.
req_type: constraint
priority: must
verification: code review of internal/scan for change state storage
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Change tracking is the summary tree's responsibility; scan helpers are stateless.
---
