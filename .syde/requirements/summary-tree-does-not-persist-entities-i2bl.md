---
id: REQ-0195
kind: requirement
name: Summary Tree Does Not Persist Entities
slug: summary-tree-does-not-persist-entities-i2bl
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:06Z"
statement: The summary tree shall not persist syde entities and shall only own the hashed file tree state.
req_type: constraint
priority: must
verification: code review of internal/tree for entity storage
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Keeping tree state separate from entity state prevents cross-contamination.
---
