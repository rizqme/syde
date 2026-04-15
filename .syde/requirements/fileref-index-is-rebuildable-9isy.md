---
id: REQ-0039
kind: requirement
name: FileRef Index Is Rebuildable
slug: fileref-index-is-rebuildable-9isy
relationships:
    - target: fileref-7vac
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:20Z"
statement: When the user runs syde reindex, the syde CLI shall rebuild every FileRef from the markdown files under .syde/ without loss of information.
req_type: functional
priority: must
verification: integration test deleting the BadgerDB index and re-running syde reindex
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: The index must never be authoritative over markdown files.
---
