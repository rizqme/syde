package audit

import (
	"fmt"

	"github.com/feedloop/syde/internal/model"
)

// planPhaseFindings walks every plan entity and reports structural
// defects in its Phases slice. These are ERROR severity because any
// of them can crash the renderer (via ChildPhases / CollectTasks
// recursion) or leave the plan unreviewable, and because the write
// path guards against producing them — a finding here means someone
// bypassed the CLI or is running against an old data file.
//
// Defects detected (one finding per occurrence):
//
//   - empty phase ID
//   - duplicate phase ID within the same plan
//   - self-parent (ph.ParentPhase == ph.ID)
//   - ParentPhase pointing at an ID that doesn't exist in the plan
//   - ParentPhase cycles across two or more phases (A->B->A)
func planPhaseFindings(all []model.EntityWithBody) []Finding {
	var out []Finding
	for _, ewb := range all {
		p, ok := ewb.Entity.(*model.PlanEntity)
		if !ok {
			continue
		}
		b := p.GetBase()

		// Build per-plan ID set for dangling-parent detection and
		// duplicate tracking.
		idSet := make(map[string]bool, len(p.Phases))
		seen := make(map[string]bool, len(p.Phases))

		for i, ph := range p.Phases {
			if ph.ID == "" {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatPlanPhase,
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Message:    fmt.Sprintf("plan %q phase index %d has empty id — delete and recreate the plan", b.Name, i),
				})
				continue
			}
			if seen[ph.ID] {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatPlanPhase,
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Message:    fmt.Sprintf("plan %q has duplicate phase id %q", b.Name, ph.ID),
				})
			}
			seen[ph.ID] = true
			idSet[ph.ID] = true
		}

		// Self-parent and dangling-parent in a single pass over the
		// phase list. Skip phases we already flagged above as
		// empty-ID — they have no identity to compare against.
		for _, ph := range p.Phases {
			if ph.ID == "" {
				continue
			}
			if ph.ParentPhase == "" {
				continue
			}
			if ph.ParentPhase == ph.ID {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatPlanPhase,
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Message:    fmt.Sprintf("plan %q phase %q lists itself as its own parent", b.Name, ph.ID),
				})
				continue
			}
			if !idSet[ph.ParentPhase] {
				out = append(out, Finding{
					Severity:   SeverityError,
					Category:   CatPlanPhase,
					EntityKind: b.Kind,
					EntitySlug: b.CanonicalSlug(),
					EntityName: b.Name,
					Message:    fmt.Sprintf("plan %q phase %q parent_phase %q is not a known phase id", b.Name, ph.ID, ph.ParentPhase),
				})
			}
		}

		// ParentPhase cycles (A->B->A and longer). DFS from every
		// phase, walking ParentPhase upward. Local visited + on-stack
		// set so a cycle is detected the moment we see a node twice
		// on the current walk. Cycles involving a node already
		// flagged (empty/self/dangling) are harmless to re-report —
		// the user needs to know the structure is broken regardless.
		parentOf := make(map[string]string, len(p.Phases))
		for _, ph := range p.Phases {
			if ph.ID != "" {
				parentOf[ph.ID] = ph.ParentPhase
			}
		}
		reportedCycle := make(map[string]bool)
		for startID := range parentOf {
			if reportedCycle[startID] {
				continue
			}
			onStack := map[string]bool{}
			node := startID
			for node != "" {
				if onStack[node] {
					// Mark every node on the current walk as reported
					// so we don't re-emit the same cycle from a
					// different entry point.
					for n := range onStack {
						reportedCycle[n] = true
					}
					out = append(out, Finding{
						Severity:   SeverityError,
						Category:   CatPlanPhase,
						EntityKind: b.Kind,
						EntitySlug: b.CanonicalSlug(),
						EntityName: b.Name,
						Message:    fmt.Sprintf("plan %q has a parent_phase cycle involving %q", b.Name, node),
					})
					break
				}
				onStack[node] = true
				next, ok := parentOf[node]
				if !ok {
					break
				}
				node = next
			}
		}
	}
	return out
}
