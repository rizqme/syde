package uiml

import (
	"fmt"
	"strings"
)

// RenderHTML renders a UIML AST to a self-contained HTML document.
func RenderHTML(nodes []*Node) string {
	var body strings.Builder
	for _, node := range nodes {
		renderNodeHTML(&body, node, 0)
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>UIML Preview</title>
<script src="https://cdn.tailwindcss.com"></script>
<script>
tailwind.config = {
  darkMode: 'class',
  theme: { extend: { colors: {
    background: '#09090b', foreground: '#fafafa', card: '#18181b',
    border: '#27272a', muted: '#a1a1aa',
  }}}
}
</script>
</head>
<body class="dark bg-background text-foreground font-sans">
%s
</body>
</html>`, body.String())
}

func renderNodeHTML(sb *strings.Builder, node *Node, depth int) {
	ind := strings.Repeat("  ", depth)

	switch node.Kind {
	case NodeScreen:
		name := node.Attr("name")
		w := node.Attr("width")
		style := ""
		if w != "" {
			style = fmt.Sprintf(` style="max-width:%spx"`, w)
		}
		sb.WriteString(fmt.Sprintf(`%s<div class="mx-auto border border-border rounded-lg overflow-hidden"%s>`, ind, style))
		if name != "" {
			sb.WriteString(fmt.Sprintf(`<div class="px-4 py-2 border-b border-border text-sm font-medium text-muted">%s</div>`, name))
		}
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</div>\n", ind))

	case NodeNavbar:
		sb.WriteString(fmt.Sprintf(`%s<nav class="flex items-center gap-4 px-4 py-3 border-b border-border">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</nav>\n", ind))

	case NodeSidebar:
		w := node.Attr("width")
		style := ""
		if w != "" {
			style = fmt.Sprintf(` style="width:%spx;min-width:%spx"`, w, w)
		}
		sb.WriteString(fmt.Sprintf(`%s<aside class="border-r border-border p-3"%s>`, ind, style))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</aside>\n", ind))

	case NodeLayout:
		dir := node.Attr("direction")
		cls := "flex flex-col"
		if dir == "horizontal" {
			cls = "flex flex-row"
		}
		sb.WriteString(fmt.Sprintf(`%s<div class="%s min-h-0 flex-1">`, ind, cls))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</div>\n", ind))

	case NodeMain:
		sb.WriteString(fmt.Sprintf(`%s<main class="flex-1 p-6 overflow-y-auto">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</main>\n", ind))

	case NodeGrid:
		cols := node.Attr("cols")
		gap := node.Attr("gap")
		if gap == "" {
			gap = "16"
		}
		sb.WriteString(fmt.Sprintf(`%s<div class="grid grid-cols-%s" style="gap:%spx">`, ind, cols, gap))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</div>\n", ind))

	case NodeCard:
		sb.WriteString(fmt.Sprintf(`%s<div class="bg-card border border-border rounded-lg p-4">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</div>\n", ind))

	case NodeSection:
		title := node.Attr("title")
		sb.WriteString(fmt.Sprintf(`%s<section class="mt-6">`, ind))
		if title != "" {
			sb.WriteString(fmt.Sprintf(`<h3 class="text-sm font-medium text-muted mb-3">%s</h3>`, title))
		}
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</section>\n", ind))

	case NodeHeading:
		text := collectText(node)
		level := node.Attr("level")
		cls := "text-xl font-semibold"
		tag := "h2"
		if level == "2" || level == "3" {
			cls = "text-lg font-medium"
			tag = "h3"
		}
		sb.WriteString(fmt.Sprintf(`%s<%s class="%s">%s</%s>`, ind, tag, cls, text, tag))

	case NodeText, NodeParagraph:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf(`%s<p class="text-sm text-muted">%s</p>`, ind, text))

	case NodeButton:
		text := collectText(node)
		variant := node.Attr("variant")
		cls := "px-3 py-1.5 text-sm rounded-md border border-border bg-card hover:bg-[#27272a] transition"
		if variant == "primary" {
			cls = "px-3 py-1.5 text-sm rounded-md bg-foreground text-background font-medium hover:bg-[#d4d4d8] transition"
		} else if variant == "outline" {
			cls = "px-3 py-1.5 text-sm rounded-md border border-border bg-transparent hover:bg-card transition"
		} else if variant == "danger" {
			cls = "px-3 py-1.5 text-sm rounded-md bg-red-600 text-white hover:bg-red-700 transition"
		}
		sb.WriteString(fmt.Sprintf(`%s<button class="%s">%s</button>`, ind, cls, text))

	case NodeButtonGroup:
		sb.WriteString(fmt.Sprintf(`%s<div class="flex gap-2">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</div>", ind))

	case NodeInput:
		placeholder := node.Attr("placeholder")
		inputType := node.Attr("type")
		if inputType == "" {
			inputType = "text"
		}
		sb.WriteString(fmt.Sprintf(`%s<input type="%s" placeholder="%s" class="bg-background border border-border rounded-md px-3 py-1.5 text-sm focus:outline-none focus:border-[#52525b]" />`, ind, inputType, placeholder))

	case NodeTable:
		sb.WriteString(fmt.Sprintf(`%s<table class="w-full text-sm">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</table>", ind))

	case NodeColumns:
		text := collectText(node)
		cols := strings.Split(text, ",")
		sb.WriteString(fmt.Sprintf(`%s<thead><tr class="border-b border-border text-muted">`, ind))
		for _, c := range cols {
			sb.WriteString(fmt.Sprintf(`<th class="p-2 text-left font-medium">%s</th>`, strings.TrimSpace(c)))
		}
		sb.WriteString("</tr></thead>")

	case NodeRow:
		sb.WriteString(fmt.Sprintf(`%s<tr class="border-b border-border hover:bg-[#1a1a1e]">`, ind))
		text := collectText(node)
		cells := strings.Split(text, ",")
		for _, c := range cells {
			sb.WriteString(fmt.Sprintf(`<td class="p-2">%s</td>`, strings.TrimSpace(c)))
		}
		sb.WriteString("</tr>")

	case NodeMetric:
		label := node.Attr("label")
		value := node.Attr("value")
		sb.WriteString(fmt.Sprintf(`%s<div><div class="text-xs text-muted">%s</div><div class="text-2xl font-semibold">%s</div></div>`, ind, label, value))

	case NodeTrend:
		text := collectText(node)
		dir := node.Attr("direction")
		cls := "text-muted"
		arrow := ""
		switch dir {
		case "up":
			cls = "text-green-400"
			arrow = "↑ "
		case "down":
			cls = "text-red-400"
			arrow = "↓ "
		case "flat":
			arrow = "→ "
		}
		sb.WriteString(fmt.Sprintf(`%s<div class="text-xs %s">%s%s</div>`, ind, cls, arrow, text))

	case NodeBadge:
		text := collectText(node)
		color := node.Attr("color")
		cls := "inline-flex items-center px-2 py-0.5 rounded-full text-xs border border-border"
		switch color {
		case "green":
			cls += " text-green-400 border-green-800"
		case "red":
			cls += " text-red-400 border-red-800"
		case "yellow":
			cls += " text-yellow-400 border-yellow-800"
		case "gray":
			cls += " text-[#71717a] border-[#3f3f46]"
		}
		sb.WriteString(fmt.Sprintf(`<span class="%s">%s</span>`, cls, text))

	case NodeLogo:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf(`%s<span class="font-semibold text-sm">%s</span>`, ind, text))

	case NodeNav:
		text := collectText(node)
		cls := "text-sm text-muted hover:text-foreground"
		if node.HasAttr("active") {
			cls = "text-sm text-foreground font-medium"
		}
		sb.WriteString(fmt.Sprintf(`%s<a class="%s cursor-pointer">%s</a>`, ind, cls, text))

	case NodeAvatar:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf(`%s<div class="w-8 h-8 rounded-full bg-[#27272a] flex items-center justify-center text-xs font-medium">%s</div>`, ind, text))

	case NodeSpacer:
		sb.WriteString(fmt.Sprintf(`%s<div class="flex-1"></div>`, ind))

	case NodeDivider:
		sb.WriteString(fmt.Sprintf(`%s<hr class="border-border" />`, ind))

	case NodeMenu:
		sb.WriteString(fmt.Sprintf(`%s<nav class="flex flex-col gap-1">`, ind))
		renderChildren(sb, node, depth+1)
		sb.WriteString(fmt.Sprintf("\n%s</nav>", ind))

	case NodeItem:
		text := collectText(node)
		cls := "px-3 py-1.5 text-sm text-muted rounded-md hover:bg-[#27272a] hover:text-foreground cursor-pointer"
		if node.HasAttr("active") {
			cls = "px-3 py-1.5 text-sm text-foreground rounded-md bg-[#27272a] font-medium"
		}
		sb.WriteString(fmt.Sprintf(`%s<a class="%s">%s</a>`, ind, cls, text))

	case NodeAlert:
		text := collectText(node)
		alertType := node.Attr("type")
		cls := "p-3 rounded-md border text-sm"
		switch alertType {
		case "error":
			cls += " border-red-800 text-red-400 bg-red-950"
		case "warning":
			cls += " border-yellow-800 text-yellow-400 bg-yellow-950"
		case "success":
			cls += " border-green-800 text-green-400 bg-green-950"
		default:
			cls += " border-border text-muted bg-card"
		}
		sb.WriteString(fmt.Sprintf(`%s<div class="%s">%s</div>`, ind, cls, text))

	case NodeProgress:
		value := node.Attr("value")
		sb.WriteString(fmt.Sprintf(`%s<div class="h-1.5 bg-[#27272a] rounded-full overflow-hidden"><div class="h-full bg-foreground rounded-full" style="width:%s%%"></div></div>`, ind, value))

	case NodeTextContent:
		if node.Text != "" {
			sb.WriteString(node.Text)
		}

	default:
		renderChildren(sb, node, depth)
	}

	sb.WriteString("\n")
}

func renderChildren(sb *strings.Builder, node *Node, depth int) {
	for _, child := range node.Children {
		renderNodeHTML(sb, child, depth)
	}
}
