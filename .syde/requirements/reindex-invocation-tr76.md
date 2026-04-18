---
id: REQ-0292
kind: requirement
name: Reindex Invocation
slug: reindex-invocation-tr76
relationships:
    - target: reindex-from-files-jblp
      type: refines
    - target: storage-engine-ahgm
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:12Z"
statement: When the user runs syde reindex, the syde CLI shall rebuild the BadgerDB index from the markdown files and print counts of entities, relationships, tags, and words.
req_type: interface
priority: must
verification: integration test invoking syde reindex and inspecting printed counts
source: manual
source_ref: contract:reindex-from-files-jblp
requirement_status: active
rationale: Reindexing is the escape hatch when the cache diverges from the source-of-truth markdown.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:12Z"
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:12Z"
---
