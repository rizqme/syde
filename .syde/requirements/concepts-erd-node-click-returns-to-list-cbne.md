---
id: REQ-0226
kind: requirement
name: Concepts ERD Node Click Returns To List
slug: concepts-erd-node-click-returns-to-list-cbne
relationships:
    - target: concepts-erd-screen-spxe
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:39Z"
statement: When the user clicks a node on the concepts ERD canvas, the dashboard shall switch to list mode and select the corresponding concept entity.
req_type: interface
priority: should
verification: manual inspection of /concept ERD mode in the dashboard
source: manual
source_ref: contract:concepts-erd-screen-spxe
requirement_status: active
rationale: Node click bridges visual exploration and list-based detail inspection.
---
