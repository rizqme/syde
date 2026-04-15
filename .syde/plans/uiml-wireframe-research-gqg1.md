---
approved_at: "2026-04-15T02:38:10Z"
background: Screen contract wireframes in the dashboard look nothing like wireframes — they render as stacked headings because (a) the UIML lexer mishandles tag attributes (readTextContent slurps everything until the next '<', so 'direction="vertical">' becomes one giant text token the parser cannot decompose), and (b) RenderHTML uses solid Tailwind UI classes, not wireframe styles. Even the official 'syde design create' skeleton triggers parse warnings on its own output. We need to research what 'proper wireframe' means visually and how to get the existing UIML stack there. Earlier sessions worked around this by using attribute-free structural tags only, which makes everything stack vertically and look identical.
completed_at: "2026-04-15T02:50:17Z"
created_at: "2026-04-15T02:36:20Z"
id: PLN-0013
kind: plan
name: uiml-wireframe-research
objective: 'Produce a written research report covering: (1) the lexer bug''s exact root cause and minimal fix, (2) a wireframe visual language spec chosen from 4-5 candidates, (3) a tag-to-wireframe-DOM mapping table, (4) a copy-paste-ready follow-up implementation plan. Build a screenshot inspection loop using headless Google Chrome so every visual decision is grounded in a real PNG, not guesswork. The plan delivers a research artifact + an inspection helper script — NOT production changes to the renderer or backfilled wireframes.'
phases:
    - changes: scripts/wireframe-shot.sh (new). Resolves project slug from .syde/syde.yaml or syde status output, calls headless Google Chrome with --screenshot, defaults output to /tmp/wireframe-<slug>.png. Auto-starts syde open if :5703 isn't listening. Window size 1440x900.
      description: scripts/wireframe-shot.sh — bash helper that drives headless Chrome to screenshot any contract page
      details: 'Use /Applications/Google Chrome.app/Contents/MacOS/Google Chrome with --headless --disable-gpu --virtual-time-budget=5000. Project slug discovery: parse ''syde status'' or curl /api/projects then match by path. Wait for :5703 with curl loop. Output path defaults to /tmp/wireframe-<slug>.png so it stays in tmp.'
      id: phase_1
      name: Screenshot inspection helper
      notes: Committed under scripts/ because this is a useful dev tool beyond research. Future syde wireframe shot CLI can wrap the same logic.
      objective: I can run 'bash scripts/wireframe-shot.sh <slug>' and read the PNG via the Read tool to visually inspect the live render
      status: completed
      tasks:
        - write-scriptswireframe-shotsh-tree-ignore-syderesearch
    - changes: '.syde/research/uiml-survey.md (new, tree-ignored). Section 1: lexer state machine trace through Lexer.Next() showing why <layout direction="vertical"> fails and the minimal inTag-flag patch. Section 2: parser confirmation that parseElement attribute loop is already correct. Section 3: AST tag catalog filtered to wireframe-relevant kinds. Section 4: per-NodeKind RenderHTML output baseline.'
      description: Read internal/uiml end to end and produce .syde/research/uiml-survey.md with lexer trace, parser confirmation, AST list, render baseline
      details: Use syde tree context internal/uiml/lexer.go etc. — never naive Read on tracked files. Walk Lexer.Next() with a concrete example string. Document the inTag fix as a code diff. The survey is the artifact agents reference next session, so be specific about line numbers and expected behaviour.
      id: phase_2
      name: UIML source survey
      notes: First time using .syde/research/ as a research artifact directory. Run syde tree ignore .syde/research immediately so the orphan check stays clean.
      objective: Have a written reference of the entire UIML stack so the implementation plan can target the minimal patches
      status: completed
      tasks:
        - survey-internaluiml-and-write-uiml-surveymd-sections-1-4
    - changes: '/tmp/wireframe-sandbox.html (throwaway). Hand-written HTML with 4-5 columns, each containing the SAME wireframe scene rendered in a different visual language: Balsamiq (sketchy dashed greys), blueprint (white-on-blue), monochrome (grey solid blocks with labels), zoned (coloured zone overlays), labelled-only (every box shows its tag name in a top-left chip). Use the screenshot helper to capture each. Update .syde/research/uiml-survey.md with inlined PNG paths and a recommendation block.'
      description: Build /tmp/wireframe-sandbox.html with 4-5 candidate visual languages, screenshot each, pick a winner
      details: 'The scene is the components-inbox-screen layout: sidebar + main(list+detail). Each candidate uses pure CSS (no images, no fonts) so the visual rules can be distilled into Go code later. Recommendation criteria: (a) instantly reads as wireframe not real UI, (b) clearly shows region boundaries, (c) labels region names, (d) doesn''t pretend to be clickable.'
      id: phase_3
      name: Wireframe visual language exploration
      notes: Sandbox is throwaway — winning rules get distilled into Go code in the implementation plan, not committed as HTML.
      objective: Have an explicit, screenshot-grounded recommendation for the wireframe visual style with reasoning
      status: completed
      tasks:
        - hand-write-4-5-wireframe-candidates-and-pick-a-winner
    - changes: 'Append to .syde/research/uiml-survey.md a mapping table: tag | wireframe DOM | honoured attributes | example. Plus 3 worked UIML examples ready for the backfill: Overview grid, Inbox layout, ERD canvas.'
      description: Per-NodeKind table mapping every structural UIML tag to its wireframe DOM, honoured attributes, and example UIML
      details: Cover screen, layout, grid, sidebar, main, panel, card, navbar, footer, stack, columns, row, heading, text, button, list, item — at minimum. Note any missing primitives (e.g. <placeholder>, <stub>, <region>) as a 'future' subsection.
      id: phase_4
      name: Tag → wireframe mapping
      notes: The mapping table feeds straight into the implementation plan's render task. No code is written here.
      objective: Implementation plan can refer to a concrete table when wiring RenderWireframeHTML cases
      status: completed
      tasks:
        - build-the-tag-wireframe-mapping-table
    - changes: '.syde/research/uiml-survey.md gains an ''## Implementation plan (next session)'' section with the syde plan create commands for: (1) lexer fix, (2) RenderWireframeHTML new function, (3) FormatJSON rewire to call the new renderer for screen contracts, (4) re-backfill the 13 screen contracts, (5) drop attribute-free caveat from skill docs.'
      description: Append a final section to the research note containing the full syde plan create + add-phase + task create commands ready to copy-paste
      details: Each task entry includes objective + details + acceptance + affected-entity refs. The shell block must be runnable as-is — no placeholder slugs, no missing flags.
      id: phase_5
      name: Synthesize follow-up implementation plan
      notes: Closes the research loop. After this plan completes, the only thing left is to run the shell block.
      objective: Next session can run one shell block to spin up the implementation plan from this research
      status: completed
      tasks:
        - draft-the-follow-up-implementation-plan-inside-uiml-surveymd
plan_status: completed
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
scope: 'In-scope: (1) scripts/wireframe-shot.sh helper that screenshots http://localhost:5703/<project>/contract/<slug> via headless Chrome and writes PNG to /tmp; (2) source survey of internal/uiml/ producing .syde/research/uiml-survey.md (tree-ignored, NOT a syde entity, free-form research note); (3) wireframe sandbox in /tmp/ with 4-5 candidate visual languages, screenshotted and compared; (4) tag→wireframe mapping table; (5) draft follow-up implementation plan as a copy-paste shell block in the research note. Out-of-scope: actually fixing the UIML lexer in production, modifying RenderHTML, re-backfilling the 13 screen contracts, updating skill docs, adding ''syde wireframe shot'' as a CLI command, committing the research notes (they stay tree-ignored).'
slug: uiml-wireframe-research-gqg1
source: manual
updated_at: "2026-04-15T02:50:17Z"
---
