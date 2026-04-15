package uiml

import (
	"fmt"
	"strings"
)

// RenderWireframeHTML renders a UIML AST as a self-contained HTML
// document styled like a classic mid-fidelity wireframe — bordered
// regions, ✕-rect placeholders, line-bar text, bordered button
// labels, region badges. The colour scheme is dark-mode (light
// strokes on a dark card background) so the output sits cleanly
// inside the syded dashboard's contract detail panel.
//
// Used exclusively by screen-kind contracts (contract_kind=="screen").
// Design entities continue to use RenderHTML which produces a more
// realistic Tailwind preview.
//
// Visual spec (from .syde/research/uiml-survey.md section 6, adapted
// for dark mode):
//
//	stroke         #a1a1aa   muted foreground — borders, labels, glyphs
//	stroke-light   #71717a   secondary text + thin dividers
//	bg             #18181b   card background — wireframe canvas
//	bg-page        #09090b   page background (one shade darker)
//	fill-active    #27272a   selected item row background
//	w-region       2px       outer border + region splitters
//	w-thin         1.5px     placeholders, button outlines, X-cross
//	w-hair         1px       item dividers
//	radius         12px      outer container
func RenderWireframeHTML(nodes []*Node) string {
	var body strings.Builder
	for _, node := range nodes {
		renderNodeWireframe(&body, node, 0)
	}

	return wireframeDocPrefix + body.String() + wireframeDocSuffix
}

const wireframeDocPrefix = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>UIML Wireframe</title>
<style>
/*
 * All rules are scoped under .wf-root so the CSS is safe to inject
 * into the dashboard's contract detail panel via dangerouslySetInnerHTML
 * without leaking into the parent page. The .wf-frame wrapper enforces
 * a fixed 16:9 aspect ratio and uses a CSS container query to scale
 * the 1440×810 .wf-screen down to whatever width the parent gives it.
 */
.wf-root {
  --wf-stroke: #a1a1aa;
  --wf-stroke-light: #71717a;
  --wf-bg: #18181b;
  --wf-bg-page: #09090b;
  --wf-fill-active: #27272a;
  --wf-w-region: 2px;
  --wf-w-thin: 1.5px;
  --wf-w-hair: 1px;
  --wf-radius: 12px;
  --wf-radius-btn: 8px;
  --wf-radius-thumb: 4px;
  --wf-font: -apple-system, system-ui, "Inter", sans-serif;
  font-family: var(--wf-font);
  color: var(--wf-stroke);
  display: block;
  width: 100%;
}
.wf-root, .wf-root *, .wf-root *::before, .wf-root *::after { box-sizing: border-box; }
/*
 * Pure-HTML wireframe with CSS transform scaling. The .wf-frame
 * sets up a container query context with a fixed 16:9 aspect
 * ratio. The .wf-screen inside is laid out at a fixed 1440×810
 * design size and scaled down via transform: scale(100cqi / 1440px)
 * — a length-by-length division in calc() that yields the unitless
 * number transform: scale() needs. When the parent is 1440px wide,
 * scale is 1.0; at 720px it's 0.5; etc. Always preserves aspect.
 * Chrome 2022+ and Firefox 2023+ support this.
 */
.wf-root .wf-outer {
  padding: 6px;
}
.wf-root .wf-frame {
  container-type: inline-size;
  position: relative;
  width: 100%;
  aspect-ratio: 1440 / 810;
  background: var(--wf-bg);
  border: var(--wf-w-thin) solid var(--wf-stroke-light);
  border-radius: 8px;
  overflow: hidden;
}
.wf-root .wf-screen {
  position: absolute;
  top: 0;
  left: 0;
  width: 1440px;
  height: 810px;
  transform-origin: top left;
  transform: scale(calc(100cqi / 1440px));
  display: flex;
  flex-direction: column;
}
.wf-root .wf-screen.wf-row { flex-direction: row; }
.wf-root .wf-region {
  position: relative;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  overflow: hidden;
  flex: 1;
  min-width: 0;
  transition: background-color 150ms ease;
}
.wf-root .wf-region:hover { background-color: rgba(255, 255, 255, 0.03); }
.wf-root .wf-region:has(.wf-region:hover) { background-color: transparent; }
.wf-root .wf-screen.wf-row > .wf-region + .wf-region { border-left: var(--wf-w-region) solid var(--wf-stroke); }
.wf-root .wf-screen.wf-col > .wf-region + .wf-region { border-top: var(--wf-w-region) solid var(--wf-stroke); }
.wf-root .wf-screen > .wf-region + .wf-region { border-left: var(--wf-w-region) solid var(--wf-stroke); }
.wf-root .wf-region.wf-card { flex: 0 0 auto; border: var(--wf-w-region) solid var(--wf-stroke); border-radius: var(--wf-radius-thumb); }
.wf-root .wf-region.wf-panel { flex: 0 0 auto; }
.wf-root .wf-region-label {
  position: absolute;
  top: 10px;
  left: 10px;
  font: 700 10px/1.2 var(--wf-font);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: var(--wf-stroke);
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: 4px;
  padding: 3px 6px;
  background: rgba(9, 9, 11, 0.55);
  backdrop-filter: blur(2px);
  -webkit-backdrop-filter: blur(2px);
  z-index: 2;
  pointer-events: none;
  opacity: 0.4;
  transition: opacity 150ms ease, background-color 150ms ease, border-color 150ms ease;
}
.wf-root .wf-region:hover > .wf-region-label {
  opacity: 1;
  background-color: var(--wf-bg);
  border-color: #e4e4e7;
  color: #fafafa;
}
.wf-root .wf-region:has(.wf-region:hover) > .wf-region-label { opacity: 0.4; background-color: rgba(9, 9, 11, 0.55); border-color: var(--wf-stroke); color: var(--wf-stroke); }
.wf-root .wf-heading {
  font: 700 16px/1.2 var(--wf-font);
  color: var(--wf-stroke);
  margin: 0;
}
.wf-root .wf-section-title {
  font: 700 11px/1.2 var(--wf-font);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: var(--wf-stroke-light);
  margin: 0;
}
.wf-root .wf-text-lines {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.wf-root .wf-text-lines > span {
  height: 2px;
  background: var(--wf-stroke-light);
  border-radius: 2px;
}
.wf-root .wf-text-lines > span:nth-child(1) { width: 80%; }
.wf-root .wf-text-lines > span:nth-child(2) { width: 95%; }
.wf-root .wf-text-lines > span:nth-child(3) { width: 60%; }
.wf-root .wf-text {
  font: 400 12px/1.4 var(--wf-font);
  color: var(--wf-stroke-light);
  margin: 0;
}
.wf-root .wf-button {
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: var(--wf-radius-btn);
  padding: 8px 14px;
  font: 600 12px/1 var(--wf-font);
  color: var(--wf-stroke);
  background: transparent;
  display: inline-block;
}
.wf-root .wf-row-flex {
  display: flex;
  gap: 8px;
  align-items: center;
  flex-wrap: wrap;
}
.wf-root .wf-stack {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.wf-root .wf-grid {
  display: grid;
  gap: 16px;
}
.wf-root .wf-list {
  display: flex;
  flex-direction: column;
}
.wf-root .wf-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 0;
  border-bottom: var(--wf-w-hair) solid var(--wf-stroke-light);
  font: 500 13px/1.2 var(--wf-font);
  color: var(--wf-stroke);
}
.wf-root .wf-item.wf-active {
  background: var(--wf-fill-active);
  font-weight: 700;
  margin: 0 -16px;
  padding-left: 16px;
  padding-right: 16px;
  border-bottom-color: var(--wf-stroke);
}
.wf-root .wf-item-thumb {
  width: 32px;
  height: 32px;
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: var(--wf-radius-thumb);
  flex-shrink: 0;
  position: relative;
  overflow: hidden;
}
.wf-root .wf-item-thumb::before,
.wf-root .wf-item-thumb::after {
  content: "";
  position: absolute;
  top: 50%;
  left: -10%;
  width: 120%;
  height: var(--wf-w-thin);
  background: var(--wf-stroke);
}
.wf-root .wf-item-thumb::before { transform: translateY(-50%) rotate(35deg); }
.wf-root .wf-item-thumb::after  { transform: translateY(-50%) rotate(-35deg); }
.wf-root .wf-item-label { flex: 1; min-width: 0; }
.wf-root .wf-item-meta {
  font: 400 11px/1.2 var(--wf-font);
  color: var(--wf-stroke-light);
  flex-shrink: 0;
}
.wf-root .wf-placeholder {
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: 6px;
  aspect-ratio: 16 / 9;
  position: relative;
  background: transparent;
}
.wf-root .wf-placeholder::before,
.wf-root .wf-placeholder::after {
  content: "";
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: var(--wf-w-thin);
  background: var(--wf-stroke);
}
.wf-root .wf-placeholder::before { transform: translateY(-50%) rotate(28deg); }
.wf-root .wf-placeholder::after  { transform: translateY(-50%) rotate(-28deg); }
.wf-root .wf-divider {
  height: 0;
  border-top: var(--wf-w-region) solid var(--wf-stroke);
  margin: 4px 0;
}
.wf-root .wf-spacer { flex: 1; }
.wf-root .wf-icon {
  width: 24px;
  height: 24px;
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: 4px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font: 500 12px/1 var(--wf-font);
  color: var(--wf-stroke);
  flex-shrink: 0;
}
.wf-root .wf-metric { display: flex; flex-direction: column; gap: 4px; }
.wf-root .wf-metric-label {
  font: 500 10px/1.2 var(--wf-font);
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: var(--wf-stroke-light);
}
.wf-root .wf-metric-value {
  font: 700 22px/1.1 var(--wf-font);
  color: var(--wf-stroke);
}
.wf-root .wf-badge {
  display: inline-block;
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: 4px;
  padding: 2px 6px;
  font: 600 10px/1.2 var(--wf-font);
  color: var(--wf-stroke);
}
.wf-root .wf-input {
  display: flex;
  flex-direction: column;
  gap: 4px;
}
.wf-root .wf-input-label {
  font: 500 11px/1.2 var(--wf-font);
  color: var(--wf-stroke);
}
.wf-root .wf-input-underline {
  height: 22px;
  border-bottom: var(--wf-w-thin) solid var(--wf-stroke);
}
.wf-root .wf-search {
  border: var(--wf-w-thin) solid var(--wf-stroke);
  border-radius: 6px;
  padding: 8px 12px;
  font: 400 12px/1 var(--wf-font);
  color: var(--wf-stroke-light);
}
.wf-root .wf-navbar {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-bottom: var(--wf-w-region) solid var(--wf-stroke);
}
.wf-root .wf-footer {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-top: var(--wf-w-region) solid var(--wf-stroke);
  font: 400 11px/1.2 var(--wf-font);
  color: var(--wf-stroke-light);
}
.wf-root .wf-tabs {
  display: flex;
  gap: 0;
  border-bottom: var(--wf-w-thin) solid var(--wf-stroke-light);
}
.wf-root .wf-tab {
  padding: 8px 14px;
  font: 600 12px/1 var(--wf-font);
  color: var(--wf-stroke-light);
  border-bottom: 2px solid transparent;
}
.wf-root .wf-tab.wf-active {
  color: var(--wf-stroke);
  border-bottom-color: var(--wf-stroke);
}
.wf-root .wf-step {
  display: flex;
  align-items: center;
  gap: 10px;
}
.wf-root .wf-step-circle {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  border: var(--wf-w-thin) solid var(--wf-stroke);
  flex-shrink: 0;
}
.wf-root .wf-step.wf-active .wf-step-circle { background: var(--wf-stroke); }
.wf-root .wf-progress {
  height: 4px;
  background: var(--wf-stroke-light);
  border-radius: 2px;
  overflow: hidden;
  position: relative;
}
.wf-root .wf-progress-fill {
  position: absolute;
  top: 0; left: 0; bottom: 0;
  background: var(--wf-stroke);
}
</style>
</head>
<body>
<div class="wf-root"><div class="wf-outer"><div class="wf-frame">
`

const wireframeDocSuffix = `</div></div></div>
</body>
</html>`

func renderNodeWireframe(sb *strings.Builder, node *Node, depth int) {
	if node == nil {
		return
	}
	ind := strings.Repeat("  ", depth)

	switch node.Kind {
	case NodeScreen:
		// The screen itself doesn't get a label chip — the outer
		// frame IS the screen. A chip here would take flex space
		// in row direction (pushing sidebar right) and its own
		// border would visually interrupt the outer border at the
		// corner where they meet.
		dir := node.Attr("direction")
		cls := "wf-screen wf-col"
		if dir == "horizontal" || dir == "row" {
			cls = "wf-screen wf-row"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"%s\">\n", ind, cls))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeSidebar:
		w := node.Attr("width")
		style := ""
		if w != "" {
			style = fmt.Sprintf(" style=\"flex:0 0 %spx;width:%spx\"", w, w)
		} else {
			style = " style=\"flex:0 0 220px;width:220px\""
		}
		sb.WriteString(fmt.Sprintf("%s<aside class=\"wf-region wf-sidebar\"%s>\n", ind, style))
		writeRegionLabel(sb, node, "sidebar")
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</aside>\n", ind))

	case NodeMain:
		sb.WriteString(fmt.Sprintf("%s<main class=\"wf-region wf-main\">\n", ind))
		writeRegionLabel(sb, node, "main")
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</main>\n", ind))

	case NodePanel:
		w := node.Attr("width")
		style := ""
		if w != "" {
			style = fmt.Sprintf(" style=\"flex:0 0 %spx;width:%spx\"", w, w)
		}
		sb.WriteString(fmt.Sprintf("%s<section class=\"wf-region wf-panel\"%s>\n", ind, style))
		writeRegionLabel(sb, node, "panel")
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</section>\n", ind))

	case NodeCard:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-region wf-card\">\n", ind))
		writeRegionLabel(sb, node, "card")
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeSection:
		title := node.Attr("title")
		sb.WriteString(fmt.Sprintf("%s<section class=\"wf-stack\">\n", ind))
		if title != "" {
			sb.WriteString(fmt.Sprintf("%s  <h3 class=\"wf-section-title\">%s</h3>\n", ind, escapeHTML(title)))
		}
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</section>\n", ind))

	case NodeNavbar:
		sb.WriteString(fmt.Sprintf("%s<nav class=\"wf-navbar\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</nav>\n", ind))

	case NodeFooter:
		sb.WriteString(fmt.Sprintf("%s<footer class=\"wf-footer\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</footer>\n", ind))

	case NodeLayout:
		dir := node.Attr("direction")
		cls := "wf-stack"
		if dir == "horizontal" || dir == "row" {
			cls = "wf-row-flex"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"%s\">\n", ind, cls))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeStack:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-stack\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeGrid:
		cols := node.Attr("cols")
		if cols == "" {
			cols = "2"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-grid\" style=\"grid-template-columns:repeat(%s,1fr)\">\n", ind, cols))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeHeading:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s<h2 class=\"wf-heading\">%s</h2>\n", ind, escapeHTML(text)))

	case NodeText, NodeParagraph:
		text := collectText(node)
		if strings.TrimSpace(text) == "" {
			sb.WriteString(fmt.Sprintf("%s<div class=\"wf-text-lines\"><span></span><span></span><span></span></div>\n", ind))
		} else {
			sb.WriteString(fmt.Sprintf("%s<p class=\"wf-text\">%s</p>\n", ind, escapeHTML(text)))
		}

	case NodeLabel:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s<span class=\"wf-text\">%s</span>\n", ind, escapeHTML(text)))

	case NodeButton:
		text := collectText(node)
		if text == "" {
			text = "Button"
		}
		sb.WriteString(fmt.Sprintf("%s<span class=\"wf-button\">%s</span>\n", ind, escapeHTML(text)))

	case NodeButtonGroup:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-row-flex\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeInput:
		label := node.Attr("label")
		if label == "" {
			label = node.Attr("placeholder")
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-input\">\n", ind))
		if label != "" {
			sb.WriteString(fmt.Sprintf("%s  <span class=\"wf-input-label\">%s</span>\n", ind, escapeHTML(label)))
		}
		sb.WriteString(fmt.Sprintf("%s  <div class=\"wf-input-underline\"></div>\n", ind))
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeSearch:
		placeholder := node.Attr("placeholder")
		if placeholder == "" {
			placeholder = "Search"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-search\">%s</div>\n", ind, escapeHTML(placeholder)))

	case NodeList:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-list\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeMenu:
		sb.WriteString(fmt.Sprintf("%s<nav class=\"wf-stack\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</nav>\n", ind))

	case NodeItem:
		// An item is a row with optional <image> first child as a
		// thumbnail, then label(s) — the second label slides to the
		// right edge as metadata.
		cls := "wf-item"
		if node.HasAttr("active") {
			cls = "wf-item wf-active"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"%s\">\n", ind, cls))
		var thumbDone bool
		var labels []string
		for _, child := range node.Children {
			if child == nil {
				continue
			}
			if !thumbDone && child.Kind == NodeImage {
				sb.WriteString(fmt.Sprintf("%s  <span class=\"wf-item-thumb\"></span>\n", ind))
				thumbDone = true
				continue
			}
			if child.Kind == NodeLabel || child.Kind == NodeText || child.Kind == NodeTextContent {
				labels = append(labels, collectText(child))
				continue
			}
			// Anything else recurses normally so nested structure still works.
			renderNodeWireframe(sb, child, depth+1)
		}
		// If the item has plain text (no <label> children), use it.
		if len(labels) == 0 {
			text := collectText(node)
			if text != "" {
				labels = append(labels, text)
			}
		}
		for i, lbl := range labels {
			if i == len(labels)-1 && len(labels) > 1 {
				sb.WriteString(fmt.Sprintf("%s  <span class=\"wf-item-meta\">%s</span>\n", ind, escapeHTML(lbl)))
			} else {
				cls := "wf-item-label"
				if i > 0 {
					cls = "wf-item-meta"
				}
				sb.WriteString(fmt.Sprintf("%s  <span class=\"%s\">%s</span>\n", ind, cls, escapeHTML(lbl)))
			}
		}
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeImage, NodePlaceholder:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-placeholder\"></div>\n", ind))

	case NodeIcon:
		glyph := node.Attr("glyph")
		if glyph == "" {
			glyph = collectText(node)
		}
		if glyph == "" {
			glyph = "•"
		}
		sb.WriteString(fmt.Sprintf("%s<span class=\"wf-icon\">%s</span>\n", ind, escapeHTML(glyph)))

	case NodeDivider:
		sb.WriteString(fmt.Sprintf("%s<hr class=\"wf-divider\"/>\n", ind))

	case NodeSpacer:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-spacer\"></div>\n", ind))

	case NodeMetric:
		label := node.Attr("label")
		value := node.Attr("value")
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-metric\">\n", ind))
		if label != "" {
			sb.WriteString(fmt.Sprintf("%s  <span class=\"wf-metric-label\">%s</span>\n", ind, escapeHTML(label)))
		}
		if value != "" {
			sb.WriteString(fmt.Sprintf("%s  <span class=\"wf-metric-value\">%s</span>\n", ind, escapeHTML(value)))
		}
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeBadge:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s<span class=\"wf-badge\">%s</span>\n", ind, escapeHTML(text)))

	case NodeTabs:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-tabs\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeTab:
		text := collectText(node)
		cls := "wf-tab"
		if node.HasAttr("active") {
			cls = "wf-tab wf-active"
		}
		sb.WriteString(fmt.Sprintf("%s<span class=\"%s\">%s</span>\n", ind, cls, escapeHTML(text)))

	case NodeStepper:
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-stack\">\n", ind))
		renderChildrenWireframe(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("%s</div>\n", ind))

	case NodeStep:
		text := collectText(node)
		cls := "wf-step"
		if node.HasAttr("active") {
			cls = "wf-step wf-active"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"%s\"><span class=\"wf-step-circle\"></span><span class=\"wf-text\">%s</span></div>\n", ind, cls, escapeHTML(text)))

	case NodeProgress:
		val := node.Attr("value")
		if val == "" {
			val = "50"
		}
		sb.WriteString(fmt.Sprintf("%s<div class=\"wf-progress\"><div class=\"wf-progress-fill\" style=\"width:%s%%\"></div></div>\n", ind, val))

	case NodeTextContent:
		// Bare text inside a parent — emit verbatim. The parent
		// usually consumes this through collectText.
		if strings.TrimSpace(node.Text) != "" {
			sb.WriteString(escapeHTML(node.Text))
		}

	default:
		// Unknown / unhandled node — render children only so the
		// wireframe degrades gracefully instead of dropping content.
		renderChildrenWireframe(sb, node, depth)
	}
}

func renderChildrenWireframe(sb *strings.Builder, node *Node, depth int) {
	for _, child := range node.Children {
		renderNodeWireframe(sb, child, depth)
	}
}

// writeRegionLabel emits the small UPPERCASE chip in the top-left
// corner of every structural region. Uses the entity's `name` attr
// if present, otherwise the tag's default label.
func writeRegionLabel(sb *strings.Builder, node *Node, defaultLabel string) {
	label := node.Attr("name")
	if label == "" {
		label = defaultLabel
	}
	sb.WriteString(fmt.Sprintf("    <span class=\"wf-region-label\">%s</span>\n", escapeHTML(label)))
}

func escapeHTML(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
	)
	return r.Replace(s)
}
