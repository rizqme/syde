package graph

import (
	"fmt"
	"strings"
)

// RenderASCII renders a graph result as an ASCII table.
func RenderASCII(entityName string, result *SubgraphResult) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Graph for: %s\n\n", entityName))

	if len(result.Edges) == 0 {
		sb.WriteString("  (no relationships)\n")
		return sb.String()
	}

	sb.WriteString("Outbound:\n")
	for _, e := range result.Edges {
		if e.Source != "" {
			label := ""
			if e.Label != "" {
				label = " — " + e.Label
			}
			sb.WriteString(fmt.Sprintf("  → %s: %s%s\n", e.Type, e.Target, label))
		}
	}

	sb.WriteString("\nInbound:\n")
	for _, e := range result.Edges {
		if e.Target != "" {
			label := ""
			if e.Label != "" {
				label = " — " + e.Label
			}
			sb.WriteString(fmt.Sprintf("  ← %s: %s%s\n", e.Type, e.Source, label))
		}
	}

	return sb.String()
}

// RenderDOT renders a graph result in Graphviz DOT format.
func RenderDOT(entityName string, result *SubgraphResult) string {
	var sb strings.Builder

	sb.WriteString("digraph syde {\n")
	sb.WriteString("  rankdir=LR;\n")
	sb.WriteString("  node [shape=box, style=rounded];\n\n")

	// Add nodes
	seen := make(map[string]bool)
	for _, e := range result.Edges {
		if !seen[e.Source] {
			sb.WriteString(fmt.Sprintf("  \"%s\";\n", e.Source))
			seen[e.Source] = true
		}
		if !seen[e.Target] {
			sb.WriteString(fmt.Sprintf("  \"%s\";\n", e.Target))
			seen[e.Target] = true
		}
	}
	sb.WriteString("\n")

	// Add edges
	for _, e := range result.Edges {
		label := e.Type
		if e.Label != "" {
			label = e.Label
		}
		sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [label=\"%s\"];\n", e.Source, e.Target, label))
	}

	sb.WriteString("}\n")
	return sb.String()
}
