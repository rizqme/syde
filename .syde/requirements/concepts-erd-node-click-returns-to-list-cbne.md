---
id: REQ-0226
kind: requirement
name: Concepts ERD Node Click Returns To List
slug: concepts-erd-node-click-returns-to-list-cbne
relationships:
- target: erd-canvas-shall-be-removed
  type: refines
updated_at: '2026-04-17T11:03:29Z'
statement: When the user clicks a node on the concepts ERD canvas, the dashboard shall switch to list mode and select the corresponding concept entity.
req_type: interface
priority: should
verification: manual inspection of /concept ERD mode in the dashboard
source: manual
source_ref: contract:concepts-erd-screen-spxe
requirement_status: superseded
rationale: Node click bridges visual exploration and list-based detail inspection.
superseded_by:
- erd-canvas-shall-be-removed-4j8u
obsolete_reason: ERD canvas was removed from the concepts page in favor of glossary layout
---
