package audit

import (
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
)

// cycleFindings reports cyclic system nesting (belongs_to) and cyclic
// component dependencies (depends_on). Both are blocking errors.
func cycleFindings(all []model.EntityWithBody) []Finding {
	var out []Finding
	for _, cycle := range detectCycles(all, model.KindSystem, model.RelBelongsTo) {
		out = append(out, Finding{
			Severity: SeverityError,
			Category: CatSystemCycle,
			Message:  "system cycle: " + cycle,
		})
	}
	for _, cycle := range detectCycles(all, model.KindComponent, model.RelDependsOn) {
		out = append(out, Finding{
			Severity: SeverityError,
			Category: CatComponentCycle,
			Message:  "component cycle: " + cycle,
		})
	}
	return out
}

func detectCycles(all []model.EntityWithBody, kind model.EntityKind, relType string) []string {
	graph := make(map[string][]string)
	nameBySlug := make(map[string]string)
	inKind := make(map[string]bool)

	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if b.Kind != kind {
			continue
		}
		slug := utils.Slugify(b.Name)
		inKind[slug] = true
		nameBySlug[slug] = b.Name
		for _, rel := range b.Relationships {
			if rel.Type == relType {
				graph[slug] = append(graph[slug], rel.Target)
			}
		}
	}

	var cycles []string
	visited := make(map[string]bool)
	stack := make(map[string]bool)

	var dfs func(node string, path []string)
	dfs = func(node string, path []string) {
		if stack[node] {
			start := -1
			for i, n := range path {
				if n == node {
					start = i
					break
				}
			}
			if start >= 0 {
				cycle := append(path[start:], node)
				names := make([]string, len(cycle))
				for i, s := range cycle {
					if n, ok := nameBySlug[s]; ok {
						names[i] = n
					} else {
						names[i] = s
					}
				}
				cycles = append(cycles, joinArrow(names))
			}
			return
		}
		if visited[node] {
			return
		}
		visited[node] = true
		stack[node] = true
		for _, next := range graph[node] {
			if inKind[next] {
				dfs(next, append(path, node))
			}
		}
		stack[node] = false
	}

	for slug := range inKind {
		if !visited[slug] {
			dfs(slug, []string{})
		}
	}
	return cycles
}

func joinArrow(names []string) string {
	result := ""
	for i, n := range names {
		if i > 0 {
			result += " → "
		}
		result += n
	}
	return result
}
