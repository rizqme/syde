---
approved_at: "2026-04-14T09:45:49Z"
background: 'ConceptEntity today only holds meaning/lifecycle/invariants/data_sensitivity prose. It cannot represent the ERD-level detail agents need to reason about domain models — no attributes, no actions, no typed cardinality on relates_to. The dashboard renders concepts as plain text fields with no visual connection between them. Result: the concept kind is currently the weakest part of the design model; domain modelling bypasses syde and lives in prose docs or diagrams somewhere else.'
completed_at: "2026-04-14T10:00:17Z"
created_at: "2026-04-14T09:42:53Z"
id: PLN-0007
kind: plan
name: concept-as-erd
objective: 'Concept entities behave like first-class ERD nodes: structured attributes (name/type/description), structured actions (name/description), and relates_to relationships that carry a typed cardinality label. syde sync check enforces the new required fields and the cardinality enum. The dashboard renders concepts in two places — the existing detail panel (with attribute / action tables) and a new ERD page using React Flow (nodes = concepts, edges = relates_to with cardinality). Skill docs teach agents to populate attributes and actions on every concept.'
phases:
    - changes: internal/model/entity.go adds ConceptAttribute {Name, Type, Description}, ConceptAction {Name, Description}, ParseConceptAttribute, ParseConceptAction; ConceptEntity gains Attributes []ConceptAttribute and Actions []ConceptAction yaml-tagged slices. Relationship.Label already exists — no struct change needed, only parser wiring in a later phase.
      description: Typed schema for attributes, actions, and cardinality-labelled relates_to
      details: 'Mirror ContractParam/ParseContractParam exactly: pipe-separated specs with SplitN on |, path/name required, rest optional, whitespace trimmed. Attributes = ''name|type|description''; Actions = ''name|description''. YAML tags omitempty so empty concepts still round-trip cleanly.'
      id: phase_1
      name: Model — ConceptAttribute, ConceptAction, and Label-aware rels
      notes: No DB / storage layer change — the YAML serializer handles slice fields transparently via reflection.
      objective: ConceptEntity carries attributes + actions as structured YAML; rel labels persist through read/write
      status: completed
      tasks:
        - conceptattribute-conceptaction-structs-and-parsers
    - changes: 'internal/cli/add.go adds addConceptAttributes [] and addConceptActions [] flag slices plus the concept case wiring. internal/cli/update.go mirrors the same. --add-rel parsing upgraded from SplitN('':'', 2) to SplitN('':'', 3); when three parts are present, the third becomes Relationship.Label. Back-compat: two-part spec still accepted and produces an empty label.'
      description: Expose the new fields through syde add concept / syde update concept and extend --add-rel to carry cardinality labels
      details: 'parseConceptAttributes() and parseConceptActions() helpers in add.go mirror parseContractParams(). Flag help text: ''--attribute "name|type|description" (repeatable) — structured concept attribute''. Relationship parser: split on colon, len==2 → no label, len==3 → third is label, len!=2 and !=3 → error. Both add.go and update.go share the new rel parser; factor out to helpers.go if duplication is ugly.'
      id: phase_2
      name: CLI — --attribute, --action, and three-part --add-rel
      notes: update.go should only replace the Attributes/Actions slices when the user actually passed the flag (cmd.Flags().Changed) — prevents a no-op update from wiping data.
      objective: Agents can create a fully-populated ERD concept in one shell call
      status: completed
      tasks:
        - cli-flags-for-attribute-action-and-three-part-add-rel
    - changes: internal/model/validation.go adds a case *ConceptEntity block requiring Meaning and >=1 Attribute (ERROR), recommending Invariants (WARN). Each attribute must have non-empty Name and Type (ERROR on missing). internal/audit adds CatConceptIntegrity category and a conceptFindings() function that validates cardinality labels on relates_to targets — must be one of {one-to-one, one-to-many, many-to-one, many-to-many} when non-empty. Integrate into audit.Run() alongside planPhaseFindings.
      description: Validator rejects malformed concepts and invalid cardinality labels
      details: 'Validator entry: ErrorsOnMissing {meaning, attributes}, warns on invariants. Audit conceptFindings walks every ConceptEntity, scans b.Relationships for type==relates_to with non-empty Label, checks against the closed cardinality set. Message format: ''concept %q relates_to %q has invalid cardinality label %q — expected one of one-to-one/one-to-many/many-to-one/many-to-many''.'
      id: phase_3
      name: Validation — required concept fields and cardinality enum
      notes: Existing concepts without attributes will become ERRORs on next sync check — that is intentional. Each must be back-filled or the concept deleted. Flag this in the commit description.
      objective: syde sync check --strict flags every concept gap as ERROR or WARN with actionable messages
      status: completed
      tasks:
        - concept-validation-and-cardinality-audit
    - changes: 'web/package.json adds @xyflow/react. web/src/components/EntityDetail.tsx concept case renders two new sections: AttributeTable (name, type, description columns) and ActionList (name, description). web/src/pages/ERD.tsx is a new page using React Flow: nodes are concepts with attribute/action previews, edges are relates_to with cardinality label. web/src/components/Sidebar.tsx adds ''ERD'' item to the Architecture group. web/src/App.tsx wires the new route under ''__erd__''.'
      description: Render attributes + actions on the concept detail panel; add /web/src/pages/ERD.tsx using React Flow for the full domain view
      details: bun add @xyflow/react in web/; vite picks it up automatically. ERD.tsx fetches /api/<proj>/list?kind=concept, maps entities to nodes with auto-layout (fit-view on mount, manual drag allowed but not persisted), maps relates_to to edges labelled with r.label. Node component is a small card showing name + top-3 attributes. AttributeTable is a simple 3-column div grid. Unknown cardinality labels render in gray with a warning pill.
      id: phase_4
      name: Dashboard — EntityDetail panel + new ERD page
      notes: 'Layout: use the default React Flow auto-placement (no dagre) and let the user drag. Good enough for <30 concepts; spatial persistence can come later.'
      objective: Clicking into a concept shows its ERD shape; the ERD page renders every concept + relates_to edge with cardinality labels
      status: completed
      tasks:
        - dashboard-entitydetail-panel-react-flow-erd-page
        - erd-page-implementation-react-flow
    - changes: skill/SKILL.md Concept rules section rewritten with --attribute / --action / cardinality label guidance. skill/references/entity-spec.md Concept section gains attribute/action fields, cardinality enum documentation, and worked examples. skill/references/commands.md syde add reference gains the new flags.
      description: Teach agents to populate attributes, actions, and cardinality on every concept they create
      details: 'Concept example block: syde add concept ''Order'' --meaning ''...'' --invariants ''...'' --attribute ''id|uuid|primary key'' --attribute ''status|OrderStatus|draft|placed|paid|shipped|delivered'' --attribute ''total|decimal|must be > 0'' --action ''place|transitions from draft to placed'' --action ''cancel|reverts to draft if not yet shipped'' --add-rel ''customer:relates_to:many-to-one'' --add-rel ''line-item:relates_to:one-to-many'' --add-rel ''ecommerce:belongs_to''. Note: attribute description may contain pipes if wrapped in quotes (the parser keeps everything after the second pipe as description). Explicitly state that every concept must have at least one attribute or sync check will fail.'
      id: phase_5
      name: Skill — document the new concept shape
      notes: Skill changes are embedded via go:embed — just edit the .md and rebuild.
      objective: SKILL.md and entity-spec.md reflect the new rules with examples so agents get it right on the first try
      status: completed
      tasks:
        - document-attributesactionscardinality-in-the-skill
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) ConceptAttribute + ConceptAction structs with pipe-separated parsers mirroring ContractParam; (2) ConceptEntity gains Attributes and Actions slices; (3) syde add/update concept gain --attribute and --action repeatable flags; (4) --add-rel parser accepts optional third part ''target:type:label'' for cardinality, two-part form still valid; (5) validator enforces meaning + >=1 attribute on concepts, cardinality label must be one of {one-to-one, one-to-many, many-to-one, many-to-many} when present on relates_to; (6) new audit category CatConceptIntegrity; (7) dashboard EntityDetail concept panel renders attributes + actions; (8) new /web/src/pages/ERD.tsx using React Flow to visualize concepts + relates_to edges; (9) Sidebar gets an ERD item under Architecture; (10) SKILL.md and entity-spec.md document the new shape. Out-of-scope: back-filling existing concepts (validator will warn on zero-attribute concepts so agents fill them over time); attribute required/default fields (start with name|type|description, extend later); concept actions as full contract-style APIs (name + description only); drag-to-save layouts (ERD is read-only for this pass).'
slug: concept-as-erd-mitx
source: manual
updated_at: "2026-04-14T10:00:17Z"
---
