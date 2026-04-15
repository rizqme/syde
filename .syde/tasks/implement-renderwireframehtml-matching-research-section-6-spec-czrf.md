---
acceptance: Calling uiml.RenderWireframeHTML on the example 2 UIML from research section 7 produces HTML that, when screenshotted via scripts/wireframe-shot.sh, matches /tmp/wf-sandbox.png within obvious tolerances (region badges visible, ✕ thumbnails, line placeholders, active row fill, charcoal-on-white).
affected_entities:
    - uiml-parser
completed_at: "2026-04-15T03:04:48Z"
created_at: "2026-04-15T02:53:51Z"
details: 'New file internal/uiml/render_wireframe.go. Function: func RenderWireframeHTML(nodes []*Node) string. Self-contained <html><head><style>...wf classes...</style></head><body>...rendered nodes...</body></html>. Per-NodeKind switch handling screen, layout, grid, stack, sidebar, main, panel, card, section, navbar, footer, heading, text, paragraph, label, button, button-group, input, search, list, item, menu, image, placeholder, icon, divider, spacer, metric, badge, tabs, tab, stepper, step, breadcrumb, progress, table. Empty text nodes get the 3-bar placeholder. Items with an <image> child get the small ✕-thumb pattern. Active items get the fill. See sections 6 and 7 of .syde/research/uiml-survey.md for exact CSS values and per-tag DOM template.'
id: TSK-0087
kind: task
name: Implement RenderWireframeHTML matching research section 6 spec
objective: internal/uiml/render_wireframe.go exists and produces wireframe HTML matching the sandbox screenshot at /tmp/wf-sandbox.png
plan_phase: phase_2
plan_ref: uiml-wireframe-render
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: implement-renderwireframehtml-matching-research-section-6-spec-czrf
task_status: completed
updated_at: "2026-04-15T03:04:48Z"
---
