---
id: REQ-0113
kind: requirement
name: Summary Tree Requires Scanned At
slug: summary-tree-requires-scanned-at-oivu
relationships:
    - target: summary-tree-u2fo
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: storage-engine-ahgm
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:15Z"
statement: The syde CLI shall require a scanned_at timestamp on every persisted summary tree instance.
req_type: constraint
priority: must
verification: unit test loading a tree.yaml without scanned_at
source: manual
source_ref: concept:summary-tree-u2fo
requirement_status: active
rationale: scanned_at anchors diff detection and status reporting.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:15Z"
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:15Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:15Z"
---
