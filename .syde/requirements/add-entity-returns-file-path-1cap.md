---
id: REQ-0245
kind: requirement
name: Add Entity Returns File Path
slug: add-entity-returns-file-path-1cap
relationships:
    - target: add-entity-jbmc
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T10:59:59Z"
statement: When syde add succeeds, the syde CLI shall return the absolute file path of the newly created markdown entity.
req_type: interface
priority: must
verification: integration test invoking syde add and asserting the printed path exists
source: manual
source_ref: contract:add-entity-jbmc
requirement_status: active
rationale: Downstream tooling and the operator need the exact location of the created entity.
---
