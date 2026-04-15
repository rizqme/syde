---
id: REQ-0030
kind: requirement
name: Entity Bodies Must Not Contain Secrets
slug: entity-bodies-must-not-contain-secrets-3g96
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:02Z"
statement: The syde CLI shall not accept entity bodies that contain secrets or credentials and shall treat entity content as non-sensitive git-committable data.
req_type: security
priority: must
verification: manual review of contribution guidelines and secret scanning in CI
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Entities are intended to be committed to git so leaking secrets would expose them publicly.
---
