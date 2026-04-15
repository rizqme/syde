---
approved_at: "2026-04-15T02:56:47Z"
background: 'Research plan uiml-wireframe-research produced .syde/research/uiml-survey.md with the lexer fix, visual language spec, per-tag mapping, and 3 worked UIML examples. This plan executes the implementation: fix the lexer, add a wireframe HTML renderer matching the user''s reference style, rewire FormatJSON for screen contracts, re-backfill the 13 existing screen contracts, and clean up the skill docs that warn about attribute-free UIML.'
completed_at: "2026-04-15T03:11:22Z"
created_at: "2026-04-15T02:52:25Z"
id: PLN-0014
kind: plan
name: uiml-wireframe-render
objective: Screen contract detail panels in the dashboard render UIML wireframes that look like classic mid-fidelity wireframes (charcoal-on-white, X-rect placeholders, bordered button labels, region badges) instead of the current stacked headings. The UIML lexer accepts attributes correctly so syde design create skeletons stop producing parse warnings.
phases:
    - changes: 'internal/uiml/lexer.go: add inTag field on Lexer struct; in Next(), set true after reading TokTagOpen, clear on TokGT/TokSelfClose; when inTag is true and the current char isn''t a structural token, read an identifier (reusing readTagName) and emit it as TokText; never call readTextContent in the inTag branch.'
      description: Add inTag bool flag to internal/uiml/lexer.go so attributes parse correctly
      details: See section 1 of .syde/research/uiml-survey.md for the exact diff. ~10 lines. Parser is unchanged — the existing parseElement attribute loop already handles distinct identifier/equals/value tokens correctly.
      id: phase_1
      name: Lexer inTag fix
      notes: 'Smoke test: run syde design create Smoke and verify syde design show smoke produces zero parse warnings.'
      objective: <layout direction="horizontal"> and <grid cols="3"> parse without errors and the existing syde design create skeleton stops producing warnings
      status: completed
      tasks:
        - add-intag-flag-and-identifier-branch-to-lexernext
    - changes: internal/uiml/render_wireframe.go (NEW). Mirrors render_html.go's structure but emits the CSS classes from research section 6 and the per-NodeKind rules from section 7. Uses an inline <style> block with the wf- class palette; no Tailwind dependency.
      description: internal/uiml/render_wireframe.go implementing the visual language spec from research section 6
      details: Self-contained HTML doc. <style> block declares all wf- classes from section 6. Switch on NodeKind in renderNodeWireframe (mirrors renderNodeHTML but maps to wf- classes). Region tags get a top-left UPPERCASE label chip. Empty <text>/<paragraph> render as 3 line bars; non-empty render the text content. <image>/<placeholder> as X-rect. <button> as bordered rounded label. <list>/<item> as inbox row pattern with first-child <image> as thumbnail. Active items get the wf-active fill.
      id: phase_2
      name: RenderWireframeHTML new function
      notes: RenderHTML stays untouched — design entities continue to use the realistic Tailwind preview. Only screen contracts use the new wireframe renderer.
      objective: uiml.RenderWireframeHTML(nodes []*Node) returns a self-contained HTML document matching the user-reference wireframe style
      status: completed
      tasks:
        - implement-renderwireframehtml-matching-research-section-6-spec
        - add-syde-wireframe-render-cli-command
    - changes: 'internal/query/formatter.go FormatJSON case *model.ContractEntity: when ContractKind==''screen'', set entityMap[''wireframe_html''] = uiml.RenderWireframeHTML(parsedNodes) instead of uiml.RenderHTML(parsedNodes). wireframe_ascii unchanged. wireframe (raw source) unchanged.'
      description: FormatJSON contract case calls RenderWireframeHTML for screen contracts
      details: One-line change. The existing dashboard EntityDetail.tsx contract case already mounts wireframe_html via dangerouslySetInnerHTML — it doesn't care which renderer produced it.
      id: phase_3
      name: FormatJSON rewire
      notes: Verify by screenshotting any screen contract via scripts/wireframe-shot.sh after the change.
      objective: Dashboard contract detail panel renders screen wireframes via the new wireframe renderer instead of RenderHTML
      status: completed
      tasks:
        - switch-screen-contracts-to-renderwireframehtml-in-formatjson
    - changes: 'Bash script /tmp/syde-rebackfill-screens.sh: for each of the 13 screen contracts, syde update <slug> --wireframe ''<new UIML>'' using the patterns from research section 7.'
      description: Rewrite every screen contract --wireframe using attribute-rich UIML matching the 3 worked examples
      details: 'Pattern reuse: Overview uses example 1 (grid of metric cards). Inbox screens (Systems/Components/Contracts/Concepts/Flows/Decisions) use example 2 (sidebar + list + detail). Plan View / Task Board / Learning Feed / File Tree / Graph each get a tailored UIML based on their actual layout. ERD screen uses example 3 (placeholder canvas with toggle nav). Each --wireframe should be inline, single-line shell-quoted.'
      id: phase_4
      name: Re-backfill 13 screen contracts
      notes: This is the rare bulk-update case where a shell script is appropriate per the SKILL.md guidance.
      objective: Every screen contract in the project carries a wireframe that renders correctly under RenderWireframeHTML and looks like the scene it represents
      status: completed
      tasks:
        - re-backfill-13-screen-contracts-with-attribute-rich-uiml
    - changes: 'skill/SKILL.md screen contract section: remove the ''stick to attribute-free structural tags'' paragraph; replace with positive guidance on the UIML tag vocabulary. skill/references/entity-spec.md and skill/references/commands.md: same treatment.'
      description: Drop the attribute-free UIML caveat from SKILL.md, entity-spec.md, commands.md
      details: Search for 'attribute-free' in skill/ — three occurrences. Remove them and replace with 'Use <layout direction="horizontal|vertical">, <grid cols="N">, and the structural region tags. The wireframe renderer adds region badges automatically.' Add a brief table of the most common tags.
      id: phase_5
      name: Skill docs cleanup
      notes: After this lands, future agents authoring wireframes get the full UIML vocabulary instead of being told to avoid attributes.
      objective: Skill docs no longer warn agents to avoid UIML attributes
      status: completed
      tasks:
        - drop-attribute-free-uiml-caveat-from-skill-docs
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) lexer inTag flag fix in internal/uiml/lexer.go; (2) new uiml.RenderWireframeHTML function in internal/uiml/render_wireframe.go matching the CSS spec in .syde/research/uiml-survey.md section 6; (3) FormatJSON contract case calls RenderWireframeHTML for contract_kind=screen instead of RenderHTML; (4) re-backfill all 13 screen contracts using the worked examples in section 7; (5) drop the attribute-free caveat from skill/SKILL.md, skill/references/entity-spec.md, skill/references/commands.md. Out-of-scope: changing the existing RenderHTML used by design entities, adding new UIML tags from the missing primitives list, syde wireframe shot CLI command, lexer test suite (smoke tests via syde sync check are enough).'
slug: uiml-wireframe-render-iaer
source: manual
updated_at: "2026-04-15T03:11:22Z"
---
