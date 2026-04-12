package dashboard

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/query"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/utils"
)

// handleProjectAPI handles all /api/{project-slug}/... requests.
func handleProjectAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse URL: /api/{project-slug}/{endpoint}...
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) == 0 {
		jsonError(w, "missing project slug", 400)
		return
	}

	projectSlug := parts[0]
	endpoint := ""
	if len(parts) > 1 {
		endpoint = parts[1]
	}

	// Find project
	project, err := FindProjectBySlug(projectSlug)
	if err != nil {
		jsonError(w, "project not found: "+projectSlug, 404)
		return
	}

	// Open store for this project
	sydeDir := project.Path + "/.syde"
	store, err := storage.NewStore(sydeDir)
	if err != nil {
		jsonError(w, "cannot open .syde/: "+err.Error(), 500)
		return
	}
	defer store.Close()

	switch {
	case endpoint == "" || endpoint == "status":
		handleStatus(w, store)
	case endpoint == "entities":
		handleEntities(w, r, store)
	case strings.HasPrefix(endpoint, "entity/"):
		handleEntityDetail(w, endpoint, store)
	case endpoint == "graph":
		handleGraph(w, r, store)
	case endpoint == "plans":
		handlePlans(w, store)
	case endpoint == "learnings":
		handleLearnings(w, store)
	case endpoint == "tasks":
		handleTasks(w, store)
	case endpoint == "designs":
		handleDesigns(w, store)
	case endpoint == "search":
		handleSearchAPI(w, r, store)
	case endpoint == "constraints":
		handleConstraintsAPI(w, store)
	default:
		jsonError(w, "unknown endpoint: "+endpoint, 404)
	}
}

func handleStatus(w http.ResponseWriter, store *storage.Store) {
	counts := make(map[string]int)
	total := 0
	for _, kind := range model.AllEntityKinds() {
		entities, _ := store.List(kind)
		if len(entities) > 0 {
			counts[string(kind)] = len(entities)
			total += len(entities)
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"counts": counts,
		"total":  total,
	})
}

func handleEntities(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	kindFilter := r.URL.Query().Get("kind")
	tagFilter := r.URL.Query().Get("tag")
	statusFilter := r.URL.Query().Get("status")

	eng := query.NewEngine(store)
	var kind model.EntityKind
	if kindFilter != "" {
		k, ok := model.ParseEntityKind(kindFilter)
		if !ok {
			jsonError(w, "unknown kind: "+kindFilter, 400)
			return
		}
		kind = k
	}
	results, _ := eng.Filter(kind, tagFilter, statusFilter)
	json.NewEncoder(w).Encode(map[string]interface{}{"entities": results, "count": len(results)})
}

func handleEntityDetail(w http.ResponseWriter, endpoint string, store *storage.Store) {
	// endpoint format: entity/{kind}/{slug}
	rest := strings.TrimPrefix(endpoint, "entity/")
	parts := strings.SplitN(rest, "/", 2)
	if len(parts) < 2 {
		jsonError(w, "expected entity/{kind}/{slug}", 400)
		return
	}
	slug := parts[1]
	eng := query.NewEngine(store)
	resolved, err := eng.Lookup(slug)
	if err != nil {
		jsonError(w, "entity not found: "+slug, 404)
		return
	}
	w.Write([]byte(query.FormatJSON(resolved)))
}

func handleGraph(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	all, _ := store.ListAll()

	type GraphNode struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Kind string `json:"kind"`
	}
	type GraphEdge struct {
		Source string `json:"source"`
		Target string `json:"target"`
		Type   string `json:"type"`
		Label  string `json:"label,omitempty"`
	}

	var nodes []GraphNode
	var edges []GraphEdge
	seen := make(map[string]bool)

	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		if !seen[b.ID] {
			nodes = append(nodes, GraphNode{ID: b.ID, Name: b.Name, Kind: string(b.Kind)})
			seen[b.ID] = true
		}
		for _, rel := range b.Relationships {
			edges = append(edges, GraphEdge{
				Source: b.ID,
				Target: rel.Target,
				Type:   rel.Type,
				Label:  rel.Label,
			})
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"nodes": nodes, "edges": edges})
}

func handlePlans(w http.ResponseWriter, store *storage.Store) {
	plans, _ := store.List(model.KindPlan)
	var result []map[string]interface{}
	for _, ewb := range plans {
		p := ewb.Entity.(*model.PlanEntity)
		result = append(result, map[string]interface{}{
			"name":     p.Name,
			"status":   p.Status,
			"progress": p.Progress(),
			"steps":    p.Steps,
			"created":  p.CreatedAt,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"plans": result})
}

func handleLearnings(w http.ResponseWriter, store *storage.Store) {
	learnings, _ := store.List(model.KindLearning)
	var result []map[string]interface{}
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		result = append(result, map[string]interface{}{
			"name":        l.Name,
			"category":    l.Category,
			"confidence":  l.ConfLevel,
			"description": l.Description,
			"entity_refs": l.EntityRefs,
			"file":        store.FS.RelativePath(model.KindLearning, utils.Slugify(l.Name)),
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"learnings": result})
}

func handleTasks(w http.ResponseWriter, store *storage.Store) {
	tasks, _ := store.List(model.KindTask)
	var result []map[string]interface{}
	for _, ewb := range tasks {
		t := ewb.Entity.(*model.TaskEntity)
		result = append(result, map[string]interface{}{
			"name":     t.Name,
			"status":   t.Status,
			"priority": t.Priority,
			"plan_ref": t.PlanRef,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"tasks": result})
}

func handleDesigns(w http.ResponseWriter, store *storage.Store) {
	designs, _ := store.List(model.KindDesign)
	var result []map[string]interface{}
	for _, ewb := range designs {
		d := ewb.Entity.(*model.DesignEntity)
		result = append(result, map[string]interface{}{
			"name":        d.Name,
			"design_type": d.DesignType,
			"status":      d.Status,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"designs": result})
}

func handleSearchAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	q := r.URL.Query().Get("q")
	if q == "" {
		jsonError(w, "missing q parameter", 400)
		return
	}
	eng := query.NewEngine(store)
	hits, _ := eng.Search(q)
	json.NewEncoder(w).Encode(map[string]interface{}{"query": q, "results": hits, "count": len(hits)})
}

func handleConstraintsAPI(w http.ResponseWriter, store *storage.Store) {
	decisions, _ := store.List(model.KindDecision)
	var activeDecisions []map[string]string
	for _, ewb := range decisions {
		d := ewb.Entity.(*model.DecisionEntity)
		if d.Status == model.StatusActive || d.Status == "" || d.Status == model.StatusDraft {
			activeDecisions = append(activeDecisions, map[string]string{
				"name": d.Name, "statement": d.Statement, "category": d.Category,
			})
		}
	}

	learnings, _ := store.List(model.KindLearning)
	var criticalLearnings []map[string]string
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		if l.ConfLevel == model.ConfidenceHigh && (l.Category == model.CatGotcha || l.Category == model.CatConstraint) {
			criticalLearnings = append(criticalLearnings, map[string]string{
				"name": l.Name, "category": string(l.Category), "description": l.Description,
			})
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"decisions": activeDecisions,
		"learnings": criticalLearnings,
	})
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
