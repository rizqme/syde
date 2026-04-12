package uiml

import (
	"fmt"
	"strings"
)

// RenderASCII renders a UIML AST to ASCII art.
func RenderASCII(nodes []*Node, width int) string {
	if width <= 0 {
		width = 80
	}
	var sb strings.Builder
	for _, node := range nodes {
		renderNodeASCII(&sb, node, width, 0)
	}
	return sb.String()
}

func renderNodeASCII(sb *strings.Builder, node *Node, width, depth int) {
	indent := strings.Repeat("  ", depth)

	switch node.Kind {
	case NodeScreen:
		name := node.Attr("name")
		border := strings.Repeat("─", width-2)
		sb.WriteString(fmt.Sprintf("┌%s┐\n", border))
		if name != "" {
			sb.WriteString(fmt.Sprintf("│ %-*s │\n", width-4, name))
			sb.WriteString(fmt.Sprintf("├%s┤\n", border))
		}
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width-2, depth)
		}
		sb.WriteString(fmt.Sprintf("└%s┘\n", border))

	case NodeNavbar:
		items := collectTextChildren(node)
		line := strings.Join(items, "  ")
		sb.WriteString(fmt.Sprintf("%s│ %s\n", indent, line))
		sb.WriteString(fmt.Sprintf("%s├%s\n", indent, strings.Repeat("─", width-2)))

	case NodeSidebar:
		sb.WriteString(fmt.Sprintf("%s┌─ Sidebar ─┐\n", indent))
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth+1)
		}
		sb.WriteString(fmt.Sprintf("%s└───────────┘\n", indent))

	case NodeLayout:
		dir := node.Attr("direction")
		if dir == "horizontal" {
			// Render children side-by-side (simplified)
			for _, child := range node.Children {
				renderNodeASCII(sb, child, width, depth)
			}
		} else {
			for _, child := range node.Children {
				renderNodeASCII(sb, child, width, depth)
			}
		}

	case NodeMain:
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth)
		}

	case NodeGrid:
		sb.WriteString(fmt.Sprintf("%s┌─ Grid (%s cols) ─┐\n", indent, node.Attr("cols")))
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth+1)
		}
		sb.WriteString(fmt.Sprintf("%s└─────────────────┘\n", indent))

	case NodeCard:
		sb.WriteString(fmt.Sprintf("%s┌──────────────┐\n", indent))
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth+1)
		}
		sb.WriteString(fmt.Sprintf("%s└──────────────┘\n", indent))

	case NodeSection:
		title := node.Attr("title")
		sb.WriteString(fmt.Sprintf("%s── %s ──\n", indent, title))
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth+1)
		}

	case NodeHeading:
		text := collectText(node)
		level := node.Attr("level")
		prefix := "##"
		if level == "2" {
			prefix = "###"
		}
		sb.WriteString(fmt.Sprintf("%s%s %s\n", indent, prefix, text))

	case NodeText, NodeParagraph, NodeLabel:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s%s\n", indent, text))

	case NodeButton:
		text := collectText(node)
		variant := node.Attr("variant")
		if variant == "primary" {
			sb.WriteString(fmt.Sprintf("%s[*%s*]\n", indent, text))
		} else {
			sb.WriteString(fmt.Sprintf("%s[%s]\n", indent, text))
		}

	case NodeButtonGroup:
		var buttons []string
		for _, child := range node.Children {
			if child.Kind == NodeButton {
				text := collectText(child)
				buttons = append(buttons, fmt.Sprintf("[%s]", text))
			}
		}
		sb.WriteString(fmt.Sprintf("%s%s\n", indent, strings.Join(buttons, " ")))

	case NodeInput:
		placeholder := node.Attr("placeholder")
		inputType := node.Attr("type")
		if inputType == "" {
			inputType = "text"
		}
		sb.WriteString(fmt.Sprintf("%s[_________ %s (%s)]\n", indent, placeholder, inputType))

	case NodeTable:
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth)
		}

	case NodeColumns:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s│ %s │\n", indent, text))
		sb.WriteString(fmt.Sprintf("%s├%s┤\n", indent, strings.Repeat("─", len(text)+2)))

	case NodeRow:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s│ %s │\n", indent, text))

	case NodeMetric:
		label := node.Attr("label")
		value := node.Attr("value")
		sb.WriteString(fmt.Sprintf("%s%s: %s\n", indent, label, value))

	case NodeTrend:
		text := collectText(node)
		dir := node.Attr("direction")
		arrow := "─"
		switch dir {
		case "up":
			arrow = "▲"
		case "down":
			arrow = "▼"
		}
		sb.WriteString(fmt.Sprintf("%s%s %s\n", indent, arrow, text))

	case NodeBadge:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("[%s]", text))

	case NodeBreadcrumb:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s%s\n", indent, text))

	case NodeDivider:
		sb.WriteString(fmt.Sprintf("%s%s\n", indent, strings.Repeat("─", width-depth*2)))

	case NodeMenu:
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth)
		}

	case NodeItem:
		text := collectText(node)
		icon := node.Attr("icon")
		active := node.HasAttr("active")
		prefix := "  "
		if active {
			prefix = "● "
		}
		if icon != "" {
			sb.WriteString(fmt.Sprintf("%s%s[%s] %s\n", indent, prefix, icon, text))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s%s\n", indent, prefix, text))
		}

	case NodeLogo:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s◆ %s\n", indent, text))

	case NodeNav:
		text := collectText(node)
		if node.HasAttr("active") {
			sb.WriteString(fmt.Sprintf("[*%s*] ", text))
		} else {
			sb.WriteString(fmt.Sprintf("[%s] ", text))
		}

	case NodeAlert:
		text := collectText(node)
		alertType := node.Attr("type")
		sb.WriteString(fmt.Sprintf("%s[%s] %s\n", indent, strings.ToUpper(alertType), text))

	case NodeEmptyState:
		sb.WriteString(fmt.Sprintf("%s┌─────────────────┐\n", indent))
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth+1)
		}
		sb.WriteString(fmt.Sprintf("%s└─────────────────┘\n", indent))

	case NodeProgress:
		value := node.Attr("value")
		sb.WriteString(fmt.Sprintf("%s[████░░░░░░] %s%%\n", indent, value))

	case NodeLoading:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s⟳ %s\n", indent, text))

	case NodeNote:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s📝 %s\n", indent, text))

	case NodeTodo:
		text := collectText(node)
		sb.WriteString(fmt.Sprintf("%s☐ %s\n", indent, text))

	case NodeTextContent:
		if node.Text != "" {
			sb.WriteString(fmt.Sprintf("%s%s\n", indent, node.Text))
		}

	default:
		// Generic container
		for _, child := range node.Children {
			renderNodeASCII(sb, child, width, depth)
		}
	}
}

func collectText(node *Node) string {
	if node.Text != "" {
		return node.Text
	}
	var parts []string
	for _, child := range node.Children {
		if child.Kind == NodeTextContent {
			parts = append(parts, child.Text)
		} else if child.Kind == NodeBadge {
			parts = append(parts, fmt.Sprintf("[%s]", collectText(child)))
		} else {
			parts = append(parts, collectText(child))
		}
	}
	return strings.Join(parts, " ")
}

func collectTextChildren(node *Node) []string {
	var items []string
	for _, child := range node.Children {
		text := collectText(child)
		if text != "" {
			items = append(items, text)
		}
	}
	return items
}
