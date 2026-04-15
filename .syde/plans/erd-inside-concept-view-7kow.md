---
approved_at: "2026-04-14T10:21:37Z"
background: 'The previous session shipped ERD as a separate sidebar entry (__erd__) alongside File Tree and Graph. That is wrong: ERD is a lens over concepts, not a sibling of them. It belongs inside the Concept page as a view-mode toggle so users who click ''Concepts'' can flip between the list/detail inbox and the ERD canvas without leaving the page.'
completed_at: "2026-04-14T10:30:17Z"
created_at: "2026-04-14T10:16:11Z"
id: PLN-0008
kind: plan
name: erd-inside-concept-view
objective: The dashboard Sidebar has no standalone ERD entry. Clicking 'Concepts' reveals a view-mode toggle (List / ERD). List mode is the default 2-column inbox. ERD mode renders the full React Flow canvas in the main area. The ERD page component is reused as-is, just mounted differently. Skill docs + Web SPA summary reflect the new UX.
phases:
    - changes: internal/query/formatter.go FormatJSON case *model.ConceptEntity adds attributes, actions, concept_relationships to the entity map
      description: FormatJSON omits attributes + actions + concept_relationships on the concept case
      details: Minimal surgical edit — three map assignments in the existing ConceptEntity case. The manual-map pattern is brittle (every new kind field must be added here by hand); document that in a follow-up gotcha learning but don't refactor now.
      id: phase_1
      name: Fix FormatJSON concept serialization
      notes: Discovered when the user screenshot showed a Decision concept detail panel with Meaning + Invariants + Relationships but NO attributes/actions sections, even though the backfill had written them to disk.
      objective: GET /api/<proj>/entity/<slug> returns attributes and actions for concept entities so the dashboard detail panel renders them
      status: completed
      tasks:
        - emit-attributes-and-actions-from-formatjson
    - changes: web/src/components/Sidebar.tsx drops the ERD button and the ErdIcon import. web/src/components/icons.tsx deletes the ErdIcon export. web/src/App.tsx drops __erd__ from SPECIAL_VIEWS and drops the __erd__ render block. The ERD page file itself stays — it is reused by phase 3.
      description: Drop the standalone ERD route in favour of an in-concept toggle
      details: Search for __erd__ references and remove every one. ErdIcon export is unused after the sidebar change; delete the definition too.
      id: phase_2
      name: Remove __erd__ sidebar entry
      notes: Keep web/src/pages/ERD.tsx — phase 3 mounts it differently.
      objective: Sidebar has no standalone ERD item; __erd__ routing is gone; ErdIcon is deleted
      status: completed
      tasks:
        - remove-erd-sidebar-entry-and-erdicon
    - changes: web/src/App.tsx adds a conceptView state ('list' | 'erd'). When activeKind === 'concept' and conceptView === 'erd', render <ERD /> in the main area. Otherwise render the existing 2-column. A small ConceptViewTabs component at the top of the concept area switches between the two. Selecting a node in ERD navigates to the entity detail (reuses handleNavigateEntity with auto-switch back to list).
      description: Render a List ↔ ERD segmented toggle inside the Concept page that swaps the main area between the inbox view and the React Flow canvas
      details: State lives on App.tsx so it survives selection changes. Reset conceptView to 'list' when activeKind changes away from concept. The segmented toggle renders above both list and ERD — Sticky at the top of the main area. Click-to-select in ERD should call handleNavigateEntity AND set conceptView back to 'list' so the user sees the detail pane.
      id: phase_3
      name: Concept view-mode toggle (List / ERD)
      notes: Keep the ERD canvas lightweight — if it becomes slow we can memoize the nodes/edges. For now the existing ERD.tsx works unchanged.
      objective: Clicking 'Concepts' shows the normal list+detail inbox by default. A segmented control at the top of the concept view toggles into ERD mode which fills the main area with the React Flow canvas. Toggle state is local component state (not persisted).
      status: completed
      tasks:
        - add-listerd-toggle-inside-the-concept-page
    - changes: skill/SKILL.md Concept rules mention 'open the dashboard, click Concepts, toggle to ERD view' instead of 'click ERD in the sidebar'. skill/references/entity-spec.md same. Web SPA component notes updated.
      description: Skill SKILL.md references to the ERD page updated; Web SPA notes refreshed
      details: Small text edits — one paragraph per file. Keep the Order/LineItem/Customer example unchanged since the ERD canvas still renders the same way.
      id: phase_4
      name: Docs and component summary refresh
      objective: SKILL.md describes the in-page toggle, not a sidebar item; Web SPA component summary reflects the UX change
      status: completed
      tasks:
        - refresh-skill-docs-for-in-page-erd-toggle
    - changes: internal/model/entity.go ConceptAttribute gains Refs []string with yaml tag. ParseConceptAttribute splits on |, 4th part (if present) is comma-split into Refs with each element trimmed. skill/SKILL.md and entity-spec.md document the 4-part syntax with an example (customer_id|uuid|FK|customer).
      description: ConceptAttribute gains a Refs []string field so an attribute like 'customer_id|uuid|foreign key|customer' can carry its FK-style link to another concept
      details: 'Parser: SplitN limit goes from 3 to 4. When parts[3] exists, strings.Split(parts[3], '','') with TrimSpace each. Empty refs filtered out. Back-compat: 1/2/3-part specs still parse. JSON tag ''refs,omitempty'' so existing empty attributes stay clean. No validator rule on ref target existence yet — the audit already catches unknown relationship targets and future extension can flag unknown attribute refs.'
      id: phase_5
      name: Attribute refs — attribute-level concept references
      notes: Breaking nothing — the new field is optional. Existing back-filled concepts will continue to work; their attributes just won't draw FK edges until edited.
      objective: The attribute pipe spec accepts an optional fourth part — a comma-separated list of concept slugs the attribute references — so the ERD can draw attribute-level edges
      status: completed
      tasks:
        - conceptattributerefs-and-4-part-parser
    - changes: 'web/src/pages/ERD.tsx ConceptNode rewritten: header shows concept name + meaning (short); attribute rows render name (bold) + description (muted, not type); each attribute with refs gets a source Handle keyed by attribute name; edges iterate concept.attributes and emit one edge per ref to the target concept using sourceHandle=attr-<name>. Relates_to edges still render as before but are secondary. Verify nodesDraggable/panOnDrag are not disabled.'
      description: ERD node renders name + description (not type), nodes are actually draggable, and attributes with refs draw edges from the attribute handle to the target concept
      details: 'ReactFlow props: explicitly set nodesDraggable, nodesConnectable={false}, elementsSelectable, panOnDrag, zoomOnScroll true. Custom node has Handle elements positioned per attribute row (Position.Right) with id=''attr-''+attr.name. Edges of kind ''attribute-ref'' use sourceHandle=''attr-''+attr.name and target=target concept slug, labelled with attribute name. Use a distinct color for attribute-ref edges vs relates_to edges. Position the handle absolutely within the attribute row so it sits flush to the card edge.'
      id: phase_6
      name: ERD node UX — name+description, drag, attribute edges
      notes: 'The ''cannot be moved'' complaint probably means pan/drag works but is not obvious — the default React Flow Controls component should help. Make sure Controls is rendered. If nodes still aren''t draggable verify no CSS pointer-events: none on the parent.'
      objective: 'The ERD canvas is a high-level conceptual view: each concept card shows name + description (terse), attributes listed as name with a one-line description, pannable/draggable, and FK-style attribute edges route from the source attribute directly to the target concept'
      status: completed
      tasks:
        - rework-erd-node-namedescription-drag-attribute-edges
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) remove __erd__ route, Sidebar button, and ErdIcon import; (2) add a local viewMode state when activeKind === ''concept'' and render List vs ERD conditionally with a segmented toggle; (3) keep ERD page component file and logic unchanged; (4) update skill/SKILL.md concept rules to describe the in-page toggle instead of a sidebar item; (5) refresh Web SPA component notes. Out-of-scope: redesigning ERD layout, persisting the toggle selection across reloads, keeping ErdIcon (delete it — unused now).'
slug: erd-inside-concept-view-7kow
source: manual
updated_at: "2026-04-14T10:30:17Z"
---
