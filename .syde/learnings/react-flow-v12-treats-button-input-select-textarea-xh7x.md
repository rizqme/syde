---
category: gotcha
confidence: medium
description: 'React Flow v12 treats <button>, <input>, <select>, <textarea> as implicit ''nodrag'' regions — if your custom node wraps interactive surface in a button element, dragging never starts. Fix: use <div role=''button'' className=''nodrag'' onClick=...> instead. Also put pointer-events: none on any <Handle> element that isn''t actively being used for connectable edges, otherwise handles swallow the pointer-down React Flow needs to start dragging.'
discovered_at: "2026-04-14T10:51:57Z"
entity_refs:
    - web-spa
id: LRN-0019
kind: learning
name: React Flow v12 treats <button>, <input>, <select>, <textarea
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: react-flow-v12-treats-button-input-select-textarea-xh7x
source: session-observation
updated_at: "2026-04-14T10:51:57Z"
---
