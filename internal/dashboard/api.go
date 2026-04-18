package dashboard

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/query"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
	"github.com/feedloop/syde/internal/uiml"
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
	case strings.HasPrefix(endpoint, "plan/"):
		handlePlanDetail(w, strings.TrimPrefix(endpoint, "plan/"), store)
	case endpoint == "navigate":
		handleNavigate(w, r, projectSlug)
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

	// Alias-to-ID map mirrors internal/audit/relationships.go: a YAML
	// `target:` may be stored as ID, canonical slug, base slug, or
	// slugified name. Resolve every form to the canonical entity ID so
	// the frontend can match edges with a single nodeById lookup.
	targetToID := make(map[string]string)
	add := func(key, id string) {
		if key == "" {
			return
		}
		targetToID[key] = id
	}
	for _, ewb := range all {
		b := ewb.Entity.GetBase()
		add(b.ID, b.ID)
		add(b.CanonicalSlug(), b.ID)
		add(utils.BaseSlug(b.CanonicalSlug()), b.ID)
		add(utils.Slugify(b.Name), b.ID)
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
			targetID, ok := targetToID[rel.Target]
			if !ok {
				continue
			}
			edges = append(edges, GraphEdge{
				Source: b.ID,
				Target: targetID,
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

// handlePlanDetail returns the full plan including design, phases,
// resolved tasks (with status/priority/objective inline), and the
// structured Changes block. Extended changes with FieldChanges get
// pre-resolved current_values so the dashboard can render a
// side-by-side diff in one round-trip.
func handlePlanDetail(w http.ResponseWriter, slug string, store *storage.Store) {
	entity, body, err := store.Get(slug)
	if err != nil {
		jsonError(w, "plan not found: "+err.Error(), 404)
		return
	}
	plan, ok := entity.(*model.PlanEntity)
	if !ok {
		jsonError(w, slug+" is not a plan", 400)
		return
	}

	// Resolve every task referenced by every phase so the Tasks tab
	// can render status/priority/objective without N round-trips.
	// Phases store bare slugs (utils.Slugify(name)) but task entities
	// have suffixed slugs (utils.SlugifyWithSuffix). Try exact match
	// first, then fall back to BaseSlug matching scoped to the plan.
	taskBySlug := map[string]map[string]interface{}{}

	// Pre-load all tasks belonging to this plan for BaseSlug fallback.
	allTasks, _ := store.List(model.KindTask)
	planSlug := plan.GetBase().CanonicalSlug()
	planBase := utils.BaseSlug(planSlug)
	tasksByBase := map[string]*model.TaskEntity{}
	for _, ewb := range allTasks {
		te, ok := ewb.Entity.(*model.TaskEntity)
		if !ok {
			continue
		}
		// Match task to plan by either full or bare slug.
		taskPlanRef := te.PlanRef
		if taskPlanRef != planSlug && taskPlanRef != planBase && utils.BaseSlug(taskPlanRef) != planBase {
			continue
		}
		baseSlug := utils.BaseSlug(te.GetBase().CanonicalSlug())
		tasksByBase[baseSlug] = te
	}

	for _, ph := range plan.Phases {
		for _, taskSlug := range ph.Tasks {
			if _, ok := taskBySlug[taskSlug]; ok {
				continue
			}
			// Try exact slug first.
			var te *model.TaskEntity
			if t, _, err := store.Get(taskSlug); err == nil {
				te, _ = t.(*model.TaskEntity)
			}
			// Fallback: match by BaseSlug within the same plan.
			if te == nil {
				te = tasksByBase[taskSlug]
			}
			if te == nil {
				continue
			}
			taskBySlug[taskSlug] = map[string]interface{}{
				"slug":      te.GetBase().CanonicalSlug(),
				"name":      te.Name,
				"status":    te.TaskStatus,
				"priority":  te.Priority,
				"objective": te.Objective,
			}
		}
	}

	// Walk every ExtendedChange that declared FieldChanges and snapshot
	// the target entity's current values for those keys.
	resolveExtended := func(lane model.ChangeLane, kind model.EntityKind) []map[string]interface{} {
		var out []map[string]interface{}
		for _, e := range lane.Extended {
			currentValues := map[string]interface{}{}
			proposedValuesHTML := map[string]interface{}{}
			if target, _, err := store.GetByKind(kind, e.Slug); err == nil {
				if len(e.FieldChanges) > 0 {
					for field := range e.FieldChanges {
						if v, ok := readEntityFieldYAML(target, field); ok {
							currentValues[field] = v
						}
					}
				}
				if kind == model.KindContract {
					if contract, ok := target.(*model.ContractEntity); ok && contract.ContractKind == "screen" {
						if strings.TrimSpace(contract.Wireframe) != "" {
							currentValues["wireframe_html"] = renderWireframeHTML(contract.Wireframe)
						}
						if proposed, ok := e.FieldChanges["wireframe"]; ok && strings.TrimSpace(proposed) != "" && proposed != "DELETE" {
							proposedValuesHTML["wireframe"] = renderWireframeHTML(proposed)
						}
					}
				}
			}
			out = append(out, map[string]interface{}{
				"id":                   e.ID,
				"slug":                 e.Slug,
				"what":                 e.What,
				"why":                  e.Why,
				"field_changes":        e.FieldChanges,
				"current_values":       currentValues,
				"proposed_values_html": proposedValuesHTML,
				"tasks":                e.Tasks,
			})
		}
		return out
	}

	resolveNew := func(lane model.ChangeLane, kind model.EntityKind) []map[string]interface{} {
		var out []map[string]interface{}
		for _, n := range lane.New {
			draft := map[string]interface{}{}
			for k, v := range n.Draft {
				draft[k] = v
			}
			if kind == model.KindContract {
				contractKind, _ := draft["contract_kind"].(string)
				wireframe, _ := draft["wireframe"].(string)
				if contractKind == "screen" && strings.TrimSpace(wireframe) != "" {
					draft["wireframe_html"] = renderWireframeHTML(wireframe)
				}
			}
			out = append(out, map[string]interface{}{
				"id":    n.ID,
				"name":  n.Name,
				"what":  n.What,
				"why":   n.Why,
				"draft": draft,
				"tasks": n.Tasks,
			})
		}
		return out
	}

	resolvedChanges := map[string]interface{}{
		"requirements": map[string]interface{}{
			"deleted":  plan.Changes.Requirements.Deleted,
			"extended": resolveExtended(plan.Changes.Requirements, model.KindRequirement),
			"new":      resolveNew(plan.Changes.Requirements, model.KindRequirement),
		},
		"systems": map[string]interface{}{
			"deleted":  plan.Changes.Systems.Deleted,
			"extended": resolveExtended(plan.Changes.Systems, model.KindSystem),
			"new":      resolveNew(plan.Changes.Systems, model.KindSystem),
		},
		"concepts": map[string]interface{}{
			"deleted":  plan.Changes.Concepts.Deleted,
			"extended": resolveExtended(plan.Changes.Concepts, model.KindConcept),
			"new":      resolveNew(plan.Changes.Concepts, model.KindConcept),
		},
		"components": map[string]interface{}{
			"deleted":  plan.Changes.Components.Deleted,
			"extended": resolveExtended(plan.Changes.Components, model.KindComponent),
			"new":      resolveNew(plan.Changes.Components, model.KindComponent),
		},
		"contracts": map[string]interface{}{
			"deleted":  plan.Changes.Contracts.Deleted,
			"extended": resolveExtended(plan.Changes.Contracts, model.KindContract),
			"new":      resolveNew(plan.Changes.Contracts, model.KindContract),
		},
		"flows": map[string]interface{}{
			"deleted":  plan.Changes.Flows.Deleted,
			"extended": resolveExtended(plan.Changes.Flows, model.KindFlow),
			"new":      resolveNew(plan.Changes.Flows, model.KindFlow),
		},
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"slug":         plan.GetBase().CanonicalSlug(),
		"id":           plan.ID,
		"name":         plan.Name,
		"description":  plan.Description,
		"status":       plan.PlanStatus,
		"progress":     plan.Progress(),
		"background":   plan.Background,
		"objective":    plan.Objective,
		"scope":        plan.PlanScope,
		"design":       plan.Design,
		"created_at":   plan.CreatedAt,
		"approved_at":  plan.ApprovedAt,
		"completed_at": plan.CompletedAt,
		"phases":       plan.Phases,
		"task_index":   taskBySlug,
		"changes":      resolvedChanges,
		"body":         body,
	})
}

func renderWireframeHTML(source string) string {
	res := uiml.Parse(source)
	return uiml.RenderWireframeHTML(res.Nodes)
}

func handleNavigate(w http.ResponseWriter, r *http.Request, projectSlug string) {
	if r.Method != http.MethodPost {
		jsonError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		jsonError(w, "decode navigate payload: "+err.Error(), 400)
		return
	}
	if strings.TrimSpace(payload.Path) == "" {
		jsonError(w, "path is required", 400)
		return
	}
	clients := NavigateAll(projectSlug, payload.Path)
	json.NewEncoder(w).Encode(map[string]interface{}{"clients": clients})
}

// readEntityFieldYAML resolves a frontmatter field on a typed entity
// by its YAML tag. Mirrors the audit package's reflection helper but
// kept local so the dashboard package doesn't depend on internal/audit.
func readEntityFieldYAML(entity model.Entity, yamlTag string) (interface{}, bool) {
	v := reflect.ValueOf(entity)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}
	return readFieldByYAMLTag(v, yamlTag)
}

func readFieldByYAMLTag(v reflect.Value, tag string) (interface{}, bool) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			if inner, ok := readFieldByYAMLTag(v.Field(i), tag); ok {
				return inner, true
			}
			continue
		}
		yt := f.Tag.Get("yaml")
		name := strings.SplitN(yt, ",", 2)[0]
		if name == "" {
			name = strings.ToLower(f.Name)
		}
		if name == tag {
			return v.Field(i).Interface(), true
		}
	}
	return nil, false
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
