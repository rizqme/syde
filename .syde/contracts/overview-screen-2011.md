---
contract_kind: screen
description: 'Project home: entity counts grid plus recent file activity feed'
files:
    - web/src/pages/Overview.tsx
id: CON-0070
input: /
input_parameters:
    - description: active project slug from URL
      path: project
      type: string
interaction_pattern: render
kind: contract
name: Overview Screen
output: rendered overview page
output_parameters:
    - description: navigate to that kind inbox
      path: kind-card-click
      type: event
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: overview-screen-2011
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Overview" direction="vertical"><navbar><heading>syde Overview</heading></navbar><main name="Main"><grid cols="4"><card name="Systems"><metric label="systems" value="3"/></card><card name="Components"><metric label="components" value="22"/></card><card name="Contracts"><metric label="contracts" value="82"/></card><card name="Concepts"><metric label="concepts" value="11"/></card></grid><section title="Recent activity"><list><item><label>storage-engine updated</label></item><item><label>web-spa file added</label></item><item><label>concept Order created</label></item></list></section></main></screen>
---
