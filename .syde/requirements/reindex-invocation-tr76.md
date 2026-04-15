---
id: REQ-0292
kind: requirement
name: Reindex Invocation
slug: reindex-invocation-tr76
relationships:
    - target: reindex-from-files-jblp
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde reindex, the syde CLI shall rebuild the BadgerDB index from the markdown files and print counts of entities, relationships, tags, and words.
req_type: interface
priority: must
verification: integration test invoking syde reindex and inspecting printed counts
source: manual
source_ref: contract:reindex-from-files-jblp
requirement_status: active
rationale: Reindexing is the escape hatch when the cache diverges from the source-of-truth markdown.
---
