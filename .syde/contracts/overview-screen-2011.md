---
id: CON-0070
kind: contract
name: Overview Screen
slug: overview-screen-2011
description: 'Project home: entity counts grid plus recent file activity feed'
files:
    - web/src/pages/Overview.tsx
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: screen
interaction_pattern: render
input: /
input_parameters:
    - path: project
      type: string
      description: active project slug from URL
output: rendered overview page
output_parameters:
    - path: kind-card-click
      type: event
      description: navigate to that kind inbox
wireframe: <screen name="Overview" direction="vertical"><navbar><heading>syde Overview</heading></navbar><main name="Main"><grid cols="4"><card name="Systems"><metric label="systems" value="3"/></card><card name="Components"><metric label="components" value="22"/></card><card name="Contracts"><metric label="contracts" value="82"/></card><card name="Concepts"><metric label="concepts" value="11"/></card></grid><section title="Recent activity"><list><item><label>storage-engine updated</label></item><item><label>web-spa file added</label></item><item><label>concept Order created</label></item></list></section></main></screen>
---
