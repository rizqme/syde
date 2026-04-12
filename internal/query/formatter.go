package query

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// FormatRich renders a ResolvedEntity as human-readable terminal output.
func FormatRich(r *ResolvedEntity) string {
	var sb strings.Builder
	b := r.Entity.GetBase()

	kindStr := string(b.Kind)
	if len(kindStr) > 0 {
		kindStr = strings.ToUpper(kindStr[:1]) + kindStr[1:]
	}

	sb.WriteString(fmt.Sprintf("═══ %s: %s ═══\n", kindStr, b.Name))
	sb.WriteString(fmt.Sprintf("Status: %s", b.Status))
	if len(b.Tags) > 0 {
		sb.WriteString(fmt.Sprintf(" | Tags: %s", strings.Join(b.Tags, ", ")))
	}
	sb.WriteString(fmt.Sprintf("\nFile: %s\n", r.File))

	if b.Description != "" {
		sb.WriteString(fmt.Sprintf("\n  %s\n", b.Description))
	}
	if b.Purpose != "" {
		sb.WriteString(fmt.Sprintf("  Purpose: %s\n", b.Purpose))
	}

	// Entity-specific fields
	fmBytes, _ := yaml.Marshal(r.Entity)
	sb.WriteString(fmt.Sprintf("\n%s", string(fmBytes)))

	// Relationships
	if len(r.Relationships) > 0 {
		sb.WriteString("\n── Relationships ──\n")
		for _, rel := range r.Relationships {
			arrow := "→"
			name := rel.TargetName
			if name == "" {
				name = rel.TargetID
			}
			if rel.Direction == "inbound" {
				arrow = "←"
			}
			line := fmt.Sprintf("  %s %s: %s", arrow, rel.Type, name)
			if rel.TargetKind != "" {
				line += fmt.Sprintf(" (%s)", rel.TargetKind)
			}
			if rel.Label != "" {
				line += " — " + rel.Label
			}
			sb.WriteString(line + "\n")
			if rel.TargetFile != "" {
				sb.WriteString(fmt.Sprintf("       %s\n", rel.TargetFile))
			}
		}
	}

	// Learnings
	if len(r.Learnings) > 0 {
		sb.WriteString(fmt.Sprintf("\n── Learnings (%d) ──\n", len(r.Learnings)))
		for _, l := range r.Learnings {
			icon := "ℹ"
			if l.Category == "gotcha" || l.Category == "constraint" {
				icon = "⚠"
			}
			sb.WriteString(fmt.Sprintf("  %s %s: %s [%s]\n", icon, strings.ToUpper(l.Category), l.Desc, l.Confidence))
			sb.WriteString(fmt.Sprintf("    %s\n", l.File))
		}
	}

	// Tasks
	if len(r.Tasks) > 0 {
		sb.WriteString(fmt.Sprintf("\n── Tasks (%d) ──\n", len(r.Tasks)))
		for _, t := range r.Tasks {
			icon := "○"
			switch t.Status {
			case "completed":
				icon = "✓"
			case "in_progress":
				icon = "●"
			case "blocked":
				icon = "✗"
			}
			sb.WriteString(fmt.Sprintf("  %s %s [%s, %s]\n", icon, t.Name, t.Status, t.Priority))
		}
	}

	// Decisions
	if len(r.Decisions) > 0 {
		sb.WriteString(fmt.Sprintf("\n── Decisions ──\n"))
		for _, d := range r.Decisions {
			sb.WriteString(fmt.Sprintf("  ✓ %s", d.Name))
			if d.Statement != "" {
				sb.WriteString(fmt.Sprintf(": %s", d.Statement))
			}
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("    %s\n", d.File))
		}
	}

	// Body
	if r.Body != "" {
		sb.WriteString(fmt.Sprintf("\n── Details ──\n%s\n", r.Body))
	}

	// Suggested
	if len(r.Suggested) > 0 {
		sb.WriteString("\n── Drill deeper ──\n")
		for _, s := range r.Suggested {
			sb.WriteString(fmt.Sprintf("  %s\n", s))
		}
	}

	return sb.String()
}

// FormatJSON renders a ResolvedEntity as JSON.
func FormatJSON(r *ResolvedEntity) string {
	b := r.Entity.GetBase()
	out := map[string]interface{}{
		"entity": map[string]interface{}{
			"id":          b.ID,
			"kind":        b.Kind,
			"name":        b.Name,
			"slug":        strings.ToLower(strings.ReplaceAll(b.Name, " ", "-")),
			"description": b.Description,
			"purpose":     b.Purpose,
			"status":      b.Status,
			"tags":        b.Tags,
			"file":        r.File,
		},
		"relationships":    r.Relationships,
		"learnings":        r.Learnings,
		"tasks":            r.Tasks,
		"decisions":        r.Decisions,
		"suggested_queries": r.Suggested,
	}
	if r.Body != "" {
		out["body"] = r.Body
	}
	data, _ := json.MarshalIndent(out, "", "  ")
	return string(data)
}

// FormatCompact renders a list of EntitySummary as one-line-per-entity.
func FormatCompact(summaries []EntitySummary) string {
	var sb strings.Builder
	for _, s := range summaries {
		desc := s.Description
		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}
		sb.WriteString(fmt.Sprintf("  %-12s %-25s %-8s %s\n", s.Kind, s.Name, s.Status, desc))
		sb.WriteString(fmt.Sprintf("               %s\n", s.File))
	}
	return sb.String()
}

// FormatRefs renders file references only.
func FormatRefs(summaries []EntitySummary) string {
	var sb strings.Builder
	for _, s := range summaries {
		sb.WriteString(fmt.Sprintf("%s  # %s/%s\n", s.File, s.Kind, s.Name))
	}
	return sb.String()
}

// FormatImpact renders an ImpactResult.
func FormatImpact(r *ImpactResult) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Impact analysis for %s:\n\n", r.EntityName))

	for depth := 1; depth <= 5; depth++ {
		nodes, ok := r.Hops[depth]
		if !ok || len(nodes) == 0 {
			continue
		}
		label := "Direct"
		if depth > 1 {
			label = fmt.Sprintf("Transitive (%d hops)", depth)
		}
		sb.WriteString(fmt.Sprintf("%s:\n", label))
		for _, n := range nodes {
			name := n.Name
			if name == "" {
				name = n.ID
			}
			kind := n.Kind
			if kind == "" {
				kind = "?"
			}
			sb.WriteString(fmt.Sprintf("  → %s (%s)\n", name, kind))
			if n.File != "" {
				sb.WriteString(fmt.Sprintf("    %s\n", n.File))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("Total: %d entities potentially affected\n", r.Total))
	return sb.String()
}

// FormatFlowDecomposition renders a flow decomposition.
func FormatFlowDecomposition(fd *FlowDecomposition) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Flow: %s\n", fd.FlowName))
	sb.WriteString(fmt.Sprintf("Trigger: %s\n", fd.Trigger))
	sb.WriteString(fmt.Sprintf("Goal: %s\n\n", fd.Goal))

	if len(fd.Components) > 0 {
		sb.WriteString("Components in this flow:\n")
		for i, c := range fd.Components {
			warn := ""
			if c.LearningCount > 0 {
				warn = fmt.Sprintf(" | %d learnings ⚠", c.LearningCount)
			}
			sb.WriteString(fmt.Sprintf("  %d. %s%s\n", i+1, c.Name, warn))
			sb.WriteString(fmt.Sprintf("     %s\n", c.File))
		}
	}
	if len(fd.Contracts) > 0 {
		names := make([]string, len(fd.Contracts))
		for i, c := range fd.Contracts {
			names[i] = c.Name
		}
		sb.WriteString(fmt.Sprintf("\nContracts used: %s\n", strings.Join(names, ", ")))
	}
	if len(fd.Concepts) > 0 {
		names := make([]string, len(fd.Concepts))
		for i, c := range fd.Concepts {
			names[i] = c.Name
		}
		sb.WriteString(fmt.Sprintf("Concepts involved: %s\n", strings.Join(names, ", ")))
	}
	return sb.String()
}
