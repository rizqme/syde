package dashboard

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/feedloop/syde/internal/audit"
	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/query"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
)

// This file holds the new endpoints that back the syde CLI's "everything
// reads through syded" refactor. Each handler mirrors what the CLI
// command used to compute directly against storage.Store — the CLI now
// fetches the already-computed JSON via HTTP and just renders it.

// handleValidateAPI runs audit.Run and returns the findings as a health
// envelope {errors, warnings, hints, entities}. Used by both `syde
// validate` (errors-only mode) and `syde sync check` (full).
func handleValidateAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	strict := r.URL.Query().Get("strict") == "true"

	var t *tree.Tree
	if td, err := tree.Load(store.FS.Root); err == nil {
		t = td
	}

	rep, err := audit.Run(store, t, audit.Options{})
	if err != nil {
		jsonError(w, err.Error(), 500)
		return
	}

	// When called as sync-check we also fold in tree staleness so the
	// endpoint mirrors the CLI's current behavior.
	if r.URL.Path != "" && strings.HasSuffix(r.URL.Path, "/sync-check") && t != nil {
		staleFiles := 0
		staleDirs := 0
		for _, n := range t.Nodes {
			if !n.SummaryStale {
				continue
			}
			if n.Type == tree.TypeFile {
				staleFiles++
			} else {
				staleDirs++
			}
		}
		if staleFiles+staleDirs > 0 {
			rep.Findings = append(rep.Findings, audit.Finding{
				Severity: audit.SeverityWarning,
				Category: "tree_stale",
				Message: formatStaleMsg(staleFiles, staleDirs),
			})
		}
	}

	errs, _, _ := rep.Counts()
	_ = strict // strict affects CLI exit code, not the payload
	_ = errs

	json.NewEncoder(w).Encode(buildHealthPayload(rep))
}

func formatStaleMsg(files, dirs int) string {
	return "summary tree has " + itoa(files) + " stale files + " + itoa(dirs) + " stale dirs — run 'syde tree changes --leaves-only' and summarize"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}

// buildHealthPayload splits findings by severity for the JSON response.
// Mirrors reportPayload() in internal/cli/validate.go.
func buildHealthPayload(rep *audit.Report) map[string]interface{} {
	var errs, warns, hints []audit.Finding
	for _, f := range rep.Findings {
		switch f.Severity {
		case audit.SeverityError:
			errs = append(errs, f)
		case audit.SeverityWarning:
			warns = append(warns, f)
		case audit.SeverityHint:
			hints = append(hints, f)
		}
	}
	return map[string]interface{}{
		"errors":   errs,
		"warnings": warns,
		"hints":    hints,
		"entities": rep.Entities,
	}
}

// handleContextAPI returns the full project context — replaces the
// computation currently in internal/cli/context.go so the CLI doesn't
// need a Store.
func handleContextAPI(w http.ResponseWriter, store *storage.Store) {
	type entry struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Desc string `json:"description,omitempty"`
	}
	entitiesByKind := make(map[string][]entry)
	var decisions []map[string]string
	var learnings []map[string]string
	var plans []map[string]interface{}
	taskStats := map[string]int{"completed": 0, "in_progress": 0, "pending": 0, "blocked": 0}
	total := 0

	for _, kind := range model.AllEntityKinds() {
		entities, _ := store.List(kind)
		if len(entities) == 0 {
			continue
		}
		total += len(entities)
		for _, ewb := range entities {
			b := ewb.Entity.GetBase()
			switch v := ewb.Entity.(type) {
			case *model.DecisionEntity:
				decisions = append(decisions, map[string]string{
					"name": b.Name, "statement": v.Statement, "category": v.Category,
				})
			case *model.LearningEntity:
				if v.ConfLevel == model.ConfidenceHigh || v.Category == model.CatGotcha || v.Category == model.CatConstraint {
					learnings = append(learnings, map[string]string{
						"name":        b.Name,
						"category":    string(v.Category),
						"confidence":  string(v.ConfLevel),
						"description": v.Description,
					})
				}
			case *model.PlanEntity:
				done := 0
				for _, ph := range v.Phases {
					if ph.Status == model.PhaseCompleted || ph.Status == model.PhaseSkipped {
						done++
					}
				}
				plans = append(plans, map[string]interface{}{
					"name": b.Name, "status": string(v.PlanStatus),
					"progress": v.Progress(), "total_phases": len(v.Phases), "done_phases": done,
				})
			case *model.TaskEntity:
				switch v.TaskStatus {
				case model.TaskCompleted:
					taskStats["completed"]++
				case model.TaskInProgress:
					taskStats["in_progress"]++
				case model.TaskBlocked:
					taskStats["blocked"]++
				default:
					taskStats["pending"]++
				}
			default:
				entitiesByKind[string(kind)] = append(entitiesByKind[string(kind)], entry{
					ID: b.ID, Name: b.Name, Desc: b.Description,
				})
			}
		}
	}

	projectName := ""
	if cfg, err := config.Load(store.FS.Root); err == nil && cfg != nil {
		projectName = cfg.Project
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"project":   projectName,
		"entities":  entitiesByKind,
		"decisions": decisions,
		"learnings": learnings,
		"plans":     plans,
		"tasks":     taskStats,
		"total":     total,
	})
}

// handleConstraintsCheckAPI maps a source file path to its owning
// component via syde.yaml's component_paths, then attaches the
// component's boundaries + any learnings referencing it. Mirrors the
// old `syde constraints check <file>` behavior.
func handleConstraintsCheckAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	relFile := r.URL.Query().Get("path")
	if relFile == "" {
		jsonError(w, "missing path parameter", 400)
		return
	}

	cfg, err := config.Load(store.FS.Root)
	if err != nil {
		jsonError(w, "load config: "+err.Error(), 500)
		return
	}

	componentSlug := ""
	if cfg.ComponentPaths != nil {
		for slug, patterns := range cfg.ComponentPaths {
			for _, pattern := range patterns {
				if matched, _ := filepath.Match(pattern, relFile); matched {
					componentSlug = slug
					break
				}
				if strings.HasSuffix(pattern, "/**") {
					prefix := strings.TrimSuffix(pattern, "/**")
					if strings.HasPrefix(relFile, prefix+"/") {
						componentSlug = slug
						break
					}
				}
			}
			if componentSlug != "" {
				break
			}
		}
	}

	if componentSlug == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{"file": relFile})
		return
	}

	entity, _, err := store.GetByKind(model.KindComponent, componentSlug)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"file": relFile, "component_slug": componentSlug})
		return
	}
	comp := entity.(*model.ComponentEntity)

	learnings, _ := store.List(model.KindLearning)
	var compLearnings []map[string]string
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		for _, ref := range l.EntityRefs {
			if ref == comp.ID || ref == componentSlug {
				compLearnings = append(compLearnings, map[string]string{
					"category":    string(l.Category),
					"description": l.Description,
					"confidence":  string(l.ConfLevel),
				})
				break
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"file":           relFile,
		"component":      comp.Name,
		"component_slug": componentSlug,
		"responsibility": comp.Responsibility,
		"boundaries":     comp.Boundaries,
		"learnings":      compLearnings,
	})
}

// handleFilesAPI implements /files/orphans and /files/coverage.
func handleFilesAPI(w http.ResponseWriter, r *http.Request, sub string, store *storage.Store) {
	t, err := tree.Load(store.FS.Root)
	if err != nil {
		jsonError(w, "load tree: "+err.Error(), 500)
		return
	}
	all, err := store.ListAll()
	if err != nil {
		jsonError(w, err.Error(), 500)
		return
	}

	switch sub {
	case "orphans":
		orphans := audit.Orphans(all, t)
		json.NewEncoder(w).Encode(map[string]interface{}{"orphans": orphans, "count": len(orphans)})
	case "coverage":
		cov := audit.FileCoverage(all, t)
		path := r.URL.Query().Get("path")
		if path != "" {
			owners := cov[path]
			ownerLabels := make([]string, len(owners))
			for i, o := range owners {
				ownerLabels[i] = o.Label()
			}
			json.NewEncoder(w).Encode(map[string]interface{}{"path": path, "owners": ownerLabels})
			return
		}
		// Full dump: path → owners
		full := make(map[string][]string, len(cov))
		for p, owners := range cov {
			labels := make([]string, len(owners))
			for i, o := range owners {
				labels[i] = o.Label()
			}
			full[p] = labels
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"coverage": full})
	default:
		jsonError(w, "unknown files subcommand: "+sub, 404)
	}
}

// handleQueryAPI routes query modes that the existing entity/{slug}
// endpoint doesn't cover: impacts, related-to, depends-on, depended-by,
// full, search. Supports format=rich|json|compact|refs so CLI rendering
// stays server-side and the CLI stays a thin dumb client.
func handleQueryAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	mode := r.URL.Query().Get("mode")
	slug := r.URL.Query().Get("slug")
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json"
	}
	eng := query.NewEngine(store)

	writeResolved := func(res *query.ResolvedEntity) {
		if format == "rich" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatRich(res)))
			return
		}
		w.Write([]byte(query.FormatJSON(res)))
	}
	writeSummaries := func(list []query.EntitySummary) {
		switch format {
		case "rich", "compact":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatCompact(list)))
		case "refs":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatRefs(list)))
		default:
			json.NewEncoder(w).Encode(map[string]interface{}{"entities": list, "count": len(list)})
		}
	}

	switch mode {
	case "", "lookup":
		res, err := eng.Lookup(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		writeResolved(res)
	case "full":
		res, err := eng.FullContext(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		writeResolved(res)
	case "impacts":
		res, err := eng.Impacts(slug, 3)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		if format == "rich" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatImpact(res)))
			return
		}
		json.NewEncoder(w).Encode(res)
	case "related-to":
		res, err := eng.RelatedTo(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		writeSummaries(res)
	case "depends-on":
		res, err := eng.DependsOn(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		writeSummaries(res)
	case "depended-by":
		res, err := eng.DependedBy(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		writeSummaries(res)
	case "flow-components":
		res, err := eng.FlowComponents(slug)
		if err != nil {
			jsonError(w, err.Error(), 404)
			return
		}
		if format == "rich" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatFlowDecomposition(res)))
			return
		}
		json.NewEncoder(w).Encode(res)
	case "filter":
		kind := model.EntityKind(r.URL.Query().Get("kind"))
		tag := r.URL.Query().Get("tag")
		res, err := eng.Filter(kind, tag)
		if err != nil {
			jsonError(w, err.Error(), 500)
			return
		}
		writeSummaries(res)
	case "search":
		opts := query.SearchOptions{
			Query: r.URL.Query().Get("q"),
			Kind:  model.EntityKind(r.URL.Query().Get("kind")),
			Tag:   r.URL.Query().Get("tag"),
			Any:   r.URL.Query().Get("any") == "true" || r.URL.Query().Get("any") == "1",
		}
		if ls := r.URL.Query().Get("limit"); ls != "" {
			if n, err := strconv.Atoi(ls); err == nil {
				opts.Limit = n
			}
		}
		hits, err := eng.Search(opts)
		if err != nil {
			jsonError(w, err.Error(), 500)
			return
		}
		switch format {
		case "rich", "compact":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatSearchHits(hits, opts.Query)))
		case "refs":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatSearchHitsRefs(hits)))
		default:
			json.NewEncoder(w).Encode(map[string]interface{}{"hits": hits, "count": len(hits), "query": opts.Query})
		}
	case "by-file":
		path := r.URL.Query().Get("path")
		withRelated := r.URL.Query().Get("with_related") != "false"
		withContent := r.URL.Query().Get("content") == "true" || r.URL.Query().Get("content") == "1"
		res, err := eng.ByFileWith(path, withRelated, withContent)
		if err != nil {
			jsonError(w, err.Error(), 500)
			return
		}
		if format == "rich" || format == "compact" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatByFile(res)))
			return
		}
		json.NewEncoder(w).Encode(res)
	case "code":
		opts := query.SearchCodeOptions{
			Pattern: r.URL.Query().Get("q"),
		}
		if ls := r.URL.Query().Get("limit"); ls != "" {
			if n, err := strconv.Atoi(ls); err == nil {
				opts.Limit = n
			}
		}
		res, err := eng.SearchCode(opts)
		if err != nil {
			jsonError(w, err.Error(), 500)
			return
		}
		switch format {
		case "rich", "compact":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatCodeHits(res)))
		case "refs":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(query.FormatCodeHitsRefs(res)))
		default:
			json.NewEncoder(w).Encode(res)
		}
	default:
		jsonError(w, "unknown query mode: "+mode, 400)
	}
}

// handleReindexAPI accepts POST {paths:[]string, full:bool} and either
// rebuilds the whole index (full:true) or reloads the listed markdown
// paths into the BadgerDB index. Single-writer (syded) so no lock
// contention — this replaces the CLI's direct IndexEntity calls.
func handleReindexAPI(w http.ResponseWriter, r *http.Request, store *storage.Store) {
	if r.Method != http.MethodPost {
		jsonError(w, "POST required", 405)
		return
	}
	var body struct {
		Paths []string `json:"paths"`
		Full  bool     `json:"full"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err.Error() != "EOF" {
		jsonError(w, "bad body: "+err.Error(), 400)
		return
	}

	if body.Full || len(body.Paths) == 0 {
		stats, err := store.Reindex()
		if err != nil {
			jsonError(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"mode":           "full",
			"entities":       stats.Entities,
			"tags":           stats.Tags,
			"relationships":  stats.Relationships,
			"words":          stats.Words,
		})
		return
	}

	// Incremental: each path maps to a kind by its parent directory in
	// .syde/<kind>/ and a slug from the filename stem. For every listed
	// path we load the markdown, remove the stale index entries, and
	// reindex. The caller is responsible for passing the path relative
	// to .syde/ (e.g. "components/foo-abcd.md").
	indexed := 0
	var failed []string
	for _, p := range body.Paths {
		kindName, slug, ok := splitSydeRelPath(p)
		if !ok {
			failed = append(failed, p)
			continue
		}
		kind, ok := model.ParseEntityKind(kindName)
		if !ok {
			failed = append(failed, p)
			continue
		}
		if err := store.ReindexOne(kind, slug); err != nil {
			// Entity was removed — purge stale index entries using the
			// slug as a best-effort identifier and keep going.
			store.Idx.RemoveEntity(kind, slug)
			continue
		}
		indexed++
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mode":    "incremental",
		"indexed": indexed,
		"failed":  failed,
	})
}

// splitSydeRelPath turns e.g. "components/foo-abcd.md" into ("component",
// "foo-abcd", true). Accepts the kind directory name (plural) OR the
// singular kind as prefix.
func splitSydeRelPath(p string) (kind, slug string, ok bool) {
	p = strings.TrimPrefix(p, ".syde/")
	parts := strings.SplitN(p, "/", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	dir := parts[0]
	base := strings.TrimSuffix(parts[1], ".md")
	// Directories are plural; map back to singular kind names.
	kindName := strings.TrimSuffix(dir, "s")
	return kindName, base, true
}
