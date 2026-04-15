package query

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/uiml"
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
	if len(b.Tags) > 0 {
		sb.WriteString(fmt.Sprintf("Tags: %s\n", strings.Join(b.Tags, ", ")))
	}
	sb.WriteString(fmt.Sprintf("File: %s\n", r.File))

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
	entityMap := map[string]interface{}{
		"id":          b.ID,
		"kind":        b.Kind,
		"name":        b.Name,
		"slug":        b.CanonicalSlug(),
		"description": b.Description,
		"purpose":     b.Purpose,
		"tags":        b.Tags,
		"notes":       b.Notes,
		"files":       b.Files,
		"file_refs":   r.FileRefs, // file path + tree summary + stale flag for each entry
		"file":        r.File,
	}

	// Kind-specific fields
	switch e := r.Entity.(type) {
	case *model.ComponentEntity:
		entityMap["responsibility"] = e.Responsibility
		entityMap["capabilities"] = e.Capabilities
		entityMap["boundaries"] = e.Boundaries
		entityMap["behavior_summary"] = e.BehaviorSummary
		entityMap["interaction_summary"] = e.InteractionSummary
		entityMap["data_handling"] = e.DataHandling
		entityMap["scaling_notes"] = e.ScalingNotes
	case *model.ContractEntity:
		entityMap["contract_kind"] = e.ContractKind
		entityMap["interaction_pattern"] = e.InteractionPattern
		entityMap["protocol_notes"] = e.ProtocolNotes
		entityMap["input"] = e.Input
		entityMap["input_parameters"] = e.InputParameters
		entityMap["output"] = e.Output
		entityMap["output_parameters"] = e.OutputParameters
		entityMap["constraints"] = e.Constraints
		entityMap["versioning_notes"] = e.VersioningNotes
		// Screen contracts ship the raw UIML wireframe plus
		// pre-rendered HTML and ASCII so the dashboard detail panel
		// can drop the HTML straight into a contained box without
		// needing a JS UIML parser. Parse errors are tolerated —
		// the audit engine reports them at WARN severity.
		if e.ContractKind == "screen" && e.Wireframe != "" {
			entityMap["wireframe"] = e.Wireframe
			res := uiml.Parse(e.Wireframe)
			// Use the dedicated wireframe renderer (dark palette,
			// region badges, ✕-rect placeholders) instead of the
			// realistic Tailwind RenderHTML — design entities still
			// use RenderHTML for their preview.
			entityMap["wireframe_html"] = uiml.RenderWireframeHTML(res.Nodes)
			entityMap["wireframe_ascii"] = uiml.RenderASCII(res.Nodes, 80)
		}
	case *model.ConceptEntity:
		entityMap["meaning"] = e.Meaning
		entityMap["structure_notes"] = e.StructureNotes
		entityMap["lifecycle"] = e.Lifecycle
		entityMap["invariants"] = e.Invariants
		entityMap["data_sensitivity"] = e.DataSensitivity
		entityMap["concept_relationships"] = e.ConceptRelationships
		entityMap["attributes"] = e.Attributes
		entityMap["actions"] = e.Actions
	case *model.FlowEntity:
		entityMap["trigger"] = e.Trigger
		entityMap["goal"] = e.Goal
		entityMap["narrative"] = e.Narrative
		entityMap["happy_path"] = e.HappyPath
		entityMap["edge_cases"] = e.EdgeCases
		entityMap["failure_modes"] = e.FlowFailureModes
		entityMap["performance_notes"] = e.PerformanceNotes
	case *model.DecisionEntity:
		entityMap["category"] = e.Category
		entityMap["statement"] = e.Statement
		entityMap["rationale"] = e.Rationale
		entityMap["alternatives_considered"] = e.AlternativesConsidered
		entityMap["tradeoffs"] = e.Tradeoffs
		entityMap["consequences"] = e.Consequences
		entityMap["supersedes"] = e.Supersedes
	case *model.RequirementEntity:
		entityMap["statement"] = e.Statement
		entityMap["source"] = e.Source
		entityMap["source_ref"] = e.SourceRef
		entityMap["requirement_status"] = e.RequirementStatus
		entityMap["rationale"] = e.Rationale
		entityMap["acceptance_criteria"] = e.AcceptanceCriteria
		entityMap["supersedes"] = e.Supersedes
		entityMap["superseded_by"] = e.SupersededBy
		entityMap["obsolete_reason"] = e.ObsoleteReason
		entityMap["approved_at"] = e.ApprovedAt
	case *model.SystemEntity:
		entityMap["context"] = e.Context
		entityMap["scope"] = e.Scope
		entityMap["design_principles"] = e.DesignPrinciples
		entityMap["quality_goals"] = e.QualityGoals
		entityMap["assumptions"] = e.Assumptions
	case *model.PlanEntity:
		entityMap["plan_status"] = e.PlanStatus
		entityMap["background"] = e.Background
		entityMap["objective"] = e.Objective
		entityMap["scope"] = e.PlanScope
		entityMap["phases"] = e.Phases
		entityMap["progress"] = e.Progress()
		entityMap["created_at"] = e.CreatedAt
		entityMap["approved_at"] = e.ApprovedAt
		entityMap["completed_at"] = e.CompletedAt
	case *model.TaskEntity:
		entityMap["task_status"] = e.TaskStatus
		entityMap["priority"] = e.Priority
		entityMap["objective"] = e.Objective
		entityMap["details"] = e.Details
		entityMap["acceptance"] = e.Acceptance
		entityMap["plan_ref"] = e.PlanRef
		entityMap["plan_phase"] = e.PlanPhase
		entityMap["entity_refs"] = e.EntityRefs
		entityMap["block_reason"] = e.BlockReason
	}

	out := map[string]interface{}{
		"entity":            entityMap,
		"file_refs":         r.FileRefs, // also at top-level so the SPA can read it without digging into entity
		"relationships":     r.Relationships,
		"learnings":         r.Learnings,
		"tasks":             r.Tasks,
		"decisions":         r.Decisions,
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
		sb.WriteString(fmt.Sprintf("  %-12s %-25s %s\n", s.Kind, s.Name, desc))
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

// FormatSearchHits renders SearchHit results for terminal output.
// Shows kind/name, matched tokens, field, and a snippet. Surfaces
// the AND→OR broadened banner when the strict pass returned zero
// and the engine relaxed to OR — agents need to know the query was
// loose so they can review the matched_tokens list.
func FormatSearchHits(hits []SearchHit, query string) string {
	var sb strings.Builder
	if len(hits) == 0 {
		sb.WriteString(fmt.Sprintf("No results for %q\n", query))
		return sb.String()
	}
	broadened := false
	for _, h := range hits {
		if h.Broadened {
			broadened = true
			break
		}
	}
	sb.WriteString(fmt.Sprintf("Search %q — %d hit(s)", query, len(hits)))
	if broadened {
		sb.WriteString("  ⚠ broadened (no entity matched every token; results match at least one)")
	}
	sb.WriteString("\n\n")
	for _, h := range hits {
		matched := strings.Join(h.Matched, ", ")
		sb.WriteString(fmt.Sprintf("  %-10s %s\n", h.Kind, h.Name))
		sb.WriteString(fmt.Sprintf("             %s\n", h.File))
		if matched != "" {
			sb.WriteString(fmt.Sprintf("             match: %s  (field: %s)\n", matched, h.Field))
		}
		if h.Snippet != "" {
			sb.WriteString(fmt.Sprintf("             %s\n", h.Snippet))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// FormatSearchHitsRefs renders hits as "file # kind/name" lines for
// shell piping.
func FormatSearchHitsRefs(hits []SearchHit) string {
	var sb strings.Builder
	for _, h := range hits {
		sb.WriteString(fmt.Sprintf("%s  # %s/%s\n", h.File, h.Kind, h.Name))
	}
	return sb.String()
}

// FormatByFile renders a ByFileResult: owners, then one-hop related,
// then file content when --content was requested. The orphan banner
// is the architecture↔code drift signal — when this triggers on a
// tracked file the agent should map it to an entity immediately.
func FormatByFile(r *ByFileResult) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("File: %s  (match: %s)\n\n", r.Path, r.Mode))
	if len(r.Owners) == 0 {
		sb.WriteString("  ⚠ DRIFT: no owning entities — this tracked file is an orphan.\n")
		sb.WriteString("    Action: add it to a component's --file list, or run 'syde tree ignore <path>' if it is not part of the design model.\n")
	} else {
		sb.WriteString(fmt.Sprintf("Owners (%d):\n", len(r.Owners)))
		for _, o := range r.Owners {
			sb.WriteString(fmt.Sprintf("  %-10s %s\n", o.Kind, o.Name))
			if o.Slug != "" {
				sb.WriteString(fmt.Sprintf("             slug: %s\n", o.Slug))
			}
		}
	}
	if len(r.Related) > 0 {
		sb.WriteString(fmt.Sprintf("\nRelated (one-hop, %d):\n", len(r.Related)))
		for _, e := range r.Related {
			sb.WriteString(fmt.Sprintf("  %-10s %s\n", e.Kind, e.Name))
			if e.File != "" {
				sb.WriteString(fmt.Sprintf("             %s\n", e.File))
			}
		}
	}
	if r.Content != "" {
		sb.WriteString(fmt.Sprintf("\nContent (%d bytes", r.ContentBytes))
		if r.ContentTruncated {
			sb.WriteString(", truncated")
		}
		sb.WriteString("):\n")
		sb.WriteString("─────────────────────────────────────────────\n")
		sb.WriteString(r.Content)
		if !strings.HasSuffix(r.Content, "\n") {
			sb.WriteString("\n")
		}
		sb.WriteString("─────────────────────────────────────────────\n")
	}
	return sb.String()
}

// FormatCodeHits renders a CodeSearchResult for terminal output:
// path:line per hit with a snippet and the owning entity (or a
// drift warning if the file is orphaned). The owner annotation is
// the whole point — every code lookup also surfaces its
// architectural framing.
func FormatCodeHits(r *CodeSearchResult) string {
	var sb strings.Builder
	if r == nil || len(r.Hits) == 0 {
		sb.WriteString(fmt.Sprintf("No code hits for %q (engine: %s, files scanned: %d)\n", r.Pattern, r.Engine, r.FilesScanned))
		return sb.String()
	}
	sb.WriteString(fmt.Sprintf("Code search %q — %d hit(s) (engine: %s, files scanned: %d)\n\n", r.Pattern, len(r.Hits), r.Engine, r.FilesScanned))
	for _, h := range r.Hits {
		sb.WriteString(fmt.Sprintf("  %s:%d\n", h.Path, h.Line))
		sb.WriteString(fmt.Sprintf("    %s\n", h.Snippet))
		if h.OwnerName != "" {
			sb.WriteString(fmt.Sprintf("    owner: %s/%s\n", h.OwnerKind, h.OwnerName))
		} else {
			sb.WriteString("    ⚠ orphan: this file has no owning entity (architecture↔code drift)\n")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// FormatCodeHitsRefs renders code hits as "path:line  # owner" for
// shell piping and grep-style consumption.
func FormatCodeHitsRefs(r *CodeSearchResult) string {
	var sb strings.Builder
	for _, h := range r.Hits {
		owner := h.OwnerName
		if owner == "" {
			owner = "(orphan)"
		}
		sb.WriteString(fmt.Sprintf("%s:%d  # %s\n", h.Path, h.Line, owner))
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
