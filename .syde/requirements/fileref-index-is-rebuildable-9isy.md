---
id: REQ-0039
kind: requirement
name: FileRef Index Is Rebuildable
slug: fileref-index-is-rebuildable-9isy
relationships:
    - target: fileref-7vac
      type: refines
    - target: storage-engine-ahgm
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:40Z"
statement: When the user runs syde reindex, the syde CLI shall rebuild every FileRef from the markdown files under .syde/ without loss of information.
req_type: functional
priority: must
verification: integration test deleting the BadgerDB index and re-running syde reindex
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: The index must never be authoritative over markdown files.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:40Z"
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:40Z"
---
