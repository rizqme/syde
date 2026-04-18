---
id: REQ-0011
kind: requirement
name: Markdown shall be the source of truth
slug: markdown-shall-be-the-source-of-truth-urm4
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:49Z"
statement: The syde storage layer shall treat markdown files under .syde as the only authoritative entity representation.
req_type: constraint
priority: must
verification: syde reindex rebuilds the BadgerDB index from files alone with no data loss
source: manual
source_ref: system:syde-5tdt:design_principles
requirement_status: active
rationale: Git-friendly text storage lets humans diff and merge without special tools.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:49Z"
---
