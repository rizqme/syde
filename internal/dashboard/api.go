package dashboard

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/query"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
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

	// Reuse a cached Store per project. BadgerDB holds an exclusive
	// directory lock, so opening a fresh handle per request 500s under
	// concurrency. GetStore opens lazily and keeps the handle alive for
	// the daemon's lifetime — do NOT Close() it here.
	sydeDir := project.Path + "/.syde"
	store, err := GetStore(sydeDir)
	if err != nil {
		jsonError(w, "cannot open .syde/: "+err.Error(), 500)
		return
	}

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
	case endpoint == "tasks":
		handleTasks(w, store)
	case endpoint == "search":
		handleSearchAPI(w, r, store)
	case endpoint == "constraints":
		handleConstraintsAPI(w, store)
	case endpoint == "constraints-check":
		handleConstraintsCheckAPI(w, r, store)
	case endpoint == "validate" || endpoint == "sync-check":
		handleValidateAPI(w, r, store)
	case endpoint == "context":
		handleContextAPI(w, store)
	case endpoint == "query":
		handleQueryAPI(w, r, store)
	case endpoint == "reindex":
		handleReindexAPI(w, r, store)
	case endpoint == "entity" && (r.Method == http.MethodPost || r.Method == http.MethodPut):
		handleEntityWrite(w, r, store)
	case strings.HasPrefix(endpoint, "entity-raw/"):
		handleEntityRaw(w, strings.TrimPrefix(endpoint, "entity-raw/"), store)
	case strings.HasPrefix(endpoint, "entity/") && r.Method == http.MethodDelete:
		// /entity/<kind>/<slug>
		rest := strings.TrimPrefix(endpoint, "entity/")
		parts := strings.SplitN(rest, "/", 2)
		if len(parts) != 2 {
			jsonError(w, "delete path must be entity/<kind>/<slug>", 400)
			return
		}
		handleEntityDelete(w, parts[0], parts[1], store)
	case strings.HasPrefix(endpoint, "files/"):
		handleFilesAPI(w, r, strings.TrimPrefix(endpoint, "files/"), store)
	case endpoint == "tree":
		handleTree(w, store)
	case strings.HasPrefix(endpoint, "tree/"):
		handleTreeNode(w, strings.TrimPrefix(endpoint, "tree/"), store)
	default:
		jsonError(w, "unknown endpoint: "+endpoint, 404)
	}
}

func handleTree(w http.ResponseWriter, store *storage.Store) {
	t, err := tree.Load(store.FS.Root)
	if err != nil {
		jsonError(w, "cannot load tree: "+err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"scanned_at": t.ScannedAt,
		"root":       t.Root,
		"nodes":      t.Nodes,
	})
}

func handleTreeNode(w http.ResponseWriter, path string, store *storage.Store) {
	t, err := tree.Load(store.FS.Root)
	if err != nil {
		jsonError(w, "cannot load tree: "+err.Error(), 500)
		return
	}
	projectRoot := store.FS.Root + "/.."
	bundle, err := tree.BuildContext(t, path, tree.ContextOptions{
		IncludeContent: true,
		ProjectRoot:    projectRoot,
	})
	if err != nil {
		jsonError(w, err.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(bundle)
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
	results, _ := eng.Filter(kind, tagFilter)
	json.NewEncoder(w).Encode(map[string]interface{}{"entities": results, "count": len(results)})
}

func handleEntityDetail(w http.ResponseWriter, endpoint string, store *storage.Store) {
	// Supports: entity/{slug} or entity/{kind}/{slug}
	rest := strings.TrimPrefix(endpoint, "entity/")
	parts := strings.SplitN(rest, "/", 2)
	slug := parts[0]
	if len(parts) == 2 {
		slug = parts[1]
	}
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
			"status":   p.PlanStatus,
			"progress": p.Progress(),
			"phases":   p.Phases,
			"created":  p.CreatedAt,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"plans": result})
}

func handleTasks(w http.ResponseWriter, store *storage.Store) {
	tasks, _ := store.List(model.KindTask)
	var result []map[string]interface{}
	for _, ewb := range tasks {
		t := ewb.Entity.(*model.TaskEntity)
		result = append(result, map[string]interface{}{
			"name":     t.Name,
			"status":   t.TaskStatus,
			"priority": t.Priority,
			"plan_ref": t.PlanRef,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"tasks": result})
}

func handleSearchAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	q := r.URL.Query().Get("q")
	if q == "" {
		jsonError(w, "missing q parameter", 400)
		return
	}
	eng := query.NewEngine(store)
	hits, _ := eng.Search(query.SearchOptions{Query: q})
	json.NewEncoder(w).Encode(map[string]interface{}{"query": q, "results": hits, "count": len(hits)})
}

func handleConstraintsAPI(w http.ResponseWriter, store *storage.Store) {
	requirements, _ := store.List(model.KindRequirement)
	var activeReqs []map[string]string
	for _, ewb := range requirements {
		r := ewb.Entity.(*model.RequirementEntity)
		if r.RequirementStatus != "" && r.RequirementStatus != model.RequirementActive {
			continue
		}
		activeReqs = append(activeReqs, map[string]string{
			"name": r.Name, "statement": r.Statement, "rationale": r.Rationale,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"requirements": activeReqs,
	})
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
