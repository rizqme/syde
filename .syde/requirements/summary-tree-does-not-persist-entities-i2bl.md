---
id: REQ-0195
kind: requirement
name: Summary Tree Does Not Persist Entities
slug: summary-tree-does-not-persist-entities-i2bl
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:57Z"
statement: The summary tree shall not persist syde entities and shall only own the hashed file tree state.
req_type: constraint
priority: must
verification: code review of internal/tree for entity storage
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Keeping tree state separate from entity state prevents cross-contamination.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:57Z"
---
