---
approved_at: "2026-04-14T10:40:37Z"
background: 'User feedback on the just-landed ERD in-page toggle: (1) the segmented List/ERD tab occupies a dedicated row above the content which wastes vertical space and looks disconnected from the detail panel; (2) the ERD node cards still cannot be dragged despite explicit nodesDraggable on the ReactFlow component. The toggle should move to the top-right of the detail area and replace the X close button there. The drag issue needs a defensive fix — most likely the custom-node buttons or handle elements are intercepting pointer events that React Flow v12 relies on to start dragging.'
completed_at: "2026-04-14T10:51:43Z"
created_at: "2026-04-14T10:38:38Z"
id: PLN-0009
kind: plan
name: erd-polish
objective: The concept page has no dedicated toggle row. The List/ERD toggle sits at the top-right of the detail/canvas area and replaces the X close button in EntityDetail when kind=concept. Dragging ERD cards works — clicking and holding any non-button area of a node card moves it on the canvas.
phases:
    - changes: web/src/App.tsx removes the border-b dedicated tab row and instead renders the toggle as an absolute-positioned div (top-right of the concept main area) with z-10. web/src/components/EntityDetail.tsx hides the X close button when kind='concept' so the toggle overlay sits in its place without a conflicting button.
      description: Move List/ERD segmented toggle into an absolute top-right overlay; hide the X close button in EntityDetail when kind=concept
      details: Wrap the concept main-area in a relative container and absolute-position the toggle at top-4 right-4. EntityDetail's close button is rendered via DetailShell — add an optional prop hideClose that apps pass when kind=concept. Default stays unchanged for all other kinds.
      id: phase_1
      name: Toggle to top-right, replace X
      notes: 'The overlay needs pointer-events: auto (default) while the underlying area has pointer-events: none only on the space the toggle occupies — simplest is just to let the toggle sit on top.'
      objective: The concept page no longer has a dedicated tab row; the toggle sits at the top-right of the main area, in the same visual spot as the old X button
      status: completed
      tasks:
        - move-toggle-to-top-right-overlay-hide-x-in-concept-detail
    - changes: 'web/src/pages/ERD.tsx ConceptNode: replace the <button> wrapping the name with a <div> that has onClick + cursor-pointer; the root <div> no longer has onDoubleClick (click on the name div triggers selection). Per-attribute <Handle> elements gain style={{ pointerEvents: ''none'' }} so they do not capture pointer-down events from React Flow''s drag handler. Keep the top/bottom belonging handles interactive since they are still used for edge routing.'
      description: Cards must be grabbable — clicking and holding any non-interactive area of a node card should drag it on the canvas
      details: 'React Flow v12 treats <button> elements inside custom nodes as ''nodrag'' regions by default. When the entire interactive surface of a card is a button, drag never starts. Replacing the button with a <div role=''button'' onClick=...> keeps the click handler while leaving the surface draggable. Handles with pointer-events: none remain visible and still anchor edges; they just no longer swallow the pointer-down event.'
      id: phase_2
      name: ERD drag fix
      notes: If drag still fails after this change, add an explicit dragHandle class selector on the node to force React Flow to only start drag from a specific area.
      objective: ERD nodes drag smoothly under React Flow's default behaviour with the buttons removed and handles made non-interactive
      status: completed
      tasks:
        - harden-conceptnode-drag-behaviour
    - changes: 'internal/model/entity.go: remove Type field from ConceptAttribute. ParseConceptAttribute: parts[0]=name, parts[1]=description, parts[2]=refs. internal/model/validation.go: drop the attribute type requirement, keep name requirement. web/src/lib/api.ts: drop type from attribute TS type. web/src/components/EntityDetail.tsx: concept ParamTable no longer shows type column (the generic ParamTable accepts optional type so mapping stays clean). web/src/pages/ERD.tsx: already doesn''t render type. skill/SKILL.md, entity-spec.md, commands.md: rewrite 3-part example, drop all ''type'' references in the concept section. Back-filled YAML still loads; orphan ''type:'' keys drop on next save because the struct has no matching field.'
      description: ConceptAttribute drops the Type field — at the conceptual level attributes are just name + description + optional refs
      details: Breaking change to the CLI flag semantics — users will need to adapt. Re-backfill the existing internal concepts (Entity, Plan, Task, Plan Phase, Relationship, Decision, Learning, FileRef, Skill, Summary Tree, Tree Node) with the new 3-part specs so the ERD renders clean data from the start.
      id: phase_3
      name: Remove Type from concept attribute
      notes: The old 'id|uuid|primary key' reads naturally as 'id|primary key — uuid is implied'. Where the type carried information (e.g. 'status|OrderStatus|draft→placed'), merge it into the description prose.
      objective: Attribute spec becomes 'name|description|refs' (3 parts). Type is gone from the model, parser, validator, dashboard, ERD, and skill docs.
      status: completed
      tasks:
        - drop-type-from-conceptattribute-model-and-parser
        - update-entitydetail-erd-and-skill-docs-for-3-part-attribute-spec
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) move the segmented toggle out of App.tsx''s dedicated row into an absolute top-right overlay that sits above both the detail panel and the ERD canvas; (2) hide or replace the X close button in EntityDetail when kind=concept; (3) harden ConceptNode drag behaviour — remove any button elements that wrap the whole drag surface, add pointer-events: none on non-interactive handles, ensure no onMouseDown interferes with React Flow''s pointer-down handler. Out-of-scope: redesigning the concept detail layout, keyboard shortcuts, persisting node positions across reloads.'
slug: erd-polish-qgfe
source: manual
updated_at: "2026-04-14T10:51:43Z"
---
