---
category: pattern
confidence: medium
description: Screen contracts (contract_kind=screen, interaction_pattern=render) carry a Wireframe field with UIML source. Validator runs uiml.Parse and surfaces every error as a ValidationError; FormatJSON pre-renders wireframe_html and wireframe_ascii server-side so the React detail panel can drop the HTML straight into a sandboxed div via dangerouslySetInnerHTML. Use this for every UI page — the audit screen-unclaimed check warns on web/src/pages/*.tsx files lacking a screen contract.
discovered_at: "2026-04-14T11:40:50Z"
entity_refs:
    - entity-model
id: LRN-0025
kind: learning
name: Screen contracts (contract_kind=screen, interaction_pattern=
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: screen-contracts-contractkindscreen-interactionpattern-g06d
source: session-observation
updated_at: "2026-04-14T11:40:50Z"
---
