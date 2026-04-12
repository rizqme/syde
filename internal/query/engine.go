package query

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/graph"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/utils"
)

// Engine executes rich queries against the store.
type Engine struct {
	Store *storage.Store
}

// NewEngine creates a query engine.
func NewEngine(store *storage.Store) *Engine {
	return &Engine{Store: store}
}

// Lookup returns a fully resolved entity by slug.
func (e *Engine) Lookup(slug string) (*ResolvedEntity, error) {
	return Resolve(e.Store, slug)
}

// Filter returns entities matching kind/tag/status filters.
func (e *Engine) Filter(kind model.EntityKind, tag, status string) ([]EntitySummary, error) {
	var kinds []model.EntityKind
	if kind != "" {
		kinds = []model.EntityKind{kind}
	} else {
		kinds = model.AllEntityKinds()
	}

	var results []EntitySummary
	for _, k := range kinds {
		entities, err := e.Store.List(k)
		if err != nil {
			continue
		}
		for _, ewb := range entities {
			b := ewb.Entity.GetBase()
			if status != "" && string(b.Status) != status {
				continue
			}
			if tag != "" && !hasTag(b.Tags, tag) {
				continue
			}

			relCount := len(b.Relationships)
			learnCount := e.countLearningsFor(b.ID, utils.Slugify(b.Name))

			results = append(results, EntitySummary{
				ID:            b.ID,
				Kind:          string(b.Kind),
				Name:          b.Name,
				Status:        string(b.Status),
				Description:   b.Description,
				File:          e.Store.FS.RelativePath(b.Kind, utils.Slugify(b.Name)),
				RelCount:      relCount,
				LearningCount: learnCount,
			})
		}
	}
	return results, nil
}

// EntitySummary is a compact entity listing.
type EntitySummary struct {
	ID            string `json:"id"`
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	Description   string `json:"description"`
	File          string `json:"file"`
	RelCount      int    `json:"relationship_count"`
	LearningCount int    `json:"learning_count"`
}

// Impacts returns transitive impact analysis grouped by hop distance.
type ImpactResult struct {
	EntityID   string              `json:"entity_id"`
	EntityName string              `json:"entity_name"`
	Hops       map[int][]EntitySummary `json:"hops"`
	Total      int                 `json:"total"`
}

func (e *Engine) Impacts(slug string, maxDepth int) (*ImpactResult, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	hops, err := graph.ImpactAnalysis(e.Store.Idx, b.ID, maxDepth)
	if err != nil {
		return nil, err
	}

	result := &ImpactResult{
		EntityID:   b.ID,
		EntityName: b.Name,
		Hops:       make(map[int][]EntitySummary),
	}

	for depth, nodes := range hops {
		for _, node := range nodes {
			summary := EntitySummary{ID: node.ID}
			if kind, ok := utils.ParseIDKind(node.ID); ok {
				if ref, err := e.Store.Idx.LookupByID(kind, node.ID); err == nil {
					summary.Kind = string(ref.Kind)
					summary.Name = ref.Name
					summary.File = ref.File
				}
			}
			result.Hops[depth] = append(result.Hops[depth], summary)
			result.Total++
		}
	}
	return result, nil
}

// FlowDecomposition shows a flow with its participating components.
type FlowDecomposition struct {
	FlowName   string          `json:"flow_name"`
	Trigger    string          `json:"trigger"`
	Goal       string          `json:"goal"`
	Components []EntitySummary `json:"components"`
	Contracts  []EntitySummary `json:"contracts"`
	Concepts   []EntitySummary `json:"concepts"`
}

func (e *Engine) FlowComponents(slug string) (*FlowDecomposition, error) {
	entity, _, err := e.Store.GetByKind(model.KindFlow, slug)
	if err != nil {
		return nil, err
	}
	f := entity.(*model.FlowEntity)

	result := &FlowDecomposition{
		FlowName: f.Name,
		Trigger:  f.Trigger,
		Goal:     f.Goal,
	}

	for _, rel := range f.Relationships {
		kind, ok := utils.ParseIDKind(rel.Target)
		if !ok {
			continue
		}
		ref, err := e.Store.Idx.LookupByID(kind, rel.Target)
		if err != nil {
			continue
		}
		summary := EntitySummary{
			ID:            ref.ID,
			Kind:          string(ref.Kind),
			Name:          ref.Name,
			File:          ref.File,
			LearningCount: e.countLearningsFor(ref.ID, ""),
		}
		switch kind {
		case model.KindComponent:
			result.Components = append(result.Components, summary)
		case model.KindContract:
			result.Contracts = append(result.Contracts, summary)
		case model.KindConcept:
			result.Concepts = append(result.Concepts, summary)
		}
	}
	return result, nil
}

// Search performs full-text search with context.
type SearchHit struct {
	ID    string `json:"id"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
	File  string `json:"file"`
	Field string `json:"field,omitempty"`
}

func (e *Engine) Search(query string) ([]SearchHit, error) {
	refs, err := e.Store.Idx.Search(query)
	if err != nil {
		return nil, err
	}

	var hits []SearchHit
	for _, ref := range refs {
		hits = append(hits, SearchHit{
			ID:   ref.ID,
			Kind: string(ref.Kind),
			Name: ref.Name,
			File: ref.File,
		})
	}
	return hits, nil
}

func (e *Engine) countLearningsFor(entityID, slug string) int {
	learnings, _ := e.Store.List(model.KindLearning)
	count := 0
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		for _, ref := range l.EntityRefs {
			if ref == entityID || ref == slug {
				count++
				break
			}
		}
	}
	return count
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

// FullContext is alias for Lookup — returns everything.
func (e *Engine) FullContext(slug string) (*ResolvedEntity, error) {
	return e.Lookup(slug)
}

// RelatedTo returns all directly connected entities.
func (e *Engine) RelatedTo(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	seen := make(map[string]bool)
	var results []EntitySummary

	// Outbound
	for _, rel := range b.Relationships {
		if seen[rel.Target] {
			continue
		}
		seen[rel.Target] = true
		if kind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, rel.Target); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, File: ref.File})
			}
		}
	}

	// Inbound
	inbound, _ := e.Store.Idx.GetInbound(b.ID)
	for _, in := range inbound {
		if seen[in.Source] {
			continue
		}
		seen[in.Source] = true
		if kind, ok := utils.ParseIDKind(in.Source); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, in.Source); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, File: ref.File})
			}
		}
	}

	return results, nil
}

// DependsOn returns entities this slug depends on.
func (e *Engine) DependsOn(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	var results []EntitySummary
	for _, rel := range b.Relationships {
		if rel.Type != model.RelDependsOn {
			continue
		}
		if kind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, rel.Target); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, File: ref.File})
			}
		}
	}
	return results, nil
}

// DependedBy returns entities that depend on this slug.
func (e *Engine) DependedBy(slug string) ([]EntitySummary, error) {
	entity, _, err := e.Store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()

	var results []EntitySummary
	inbound, _ := e.Store.Idx.GetInbound(b.ID)
	for _, in := range inbound {
		if in.Type != model.RelDependsOn {
			continue
		}
		if kind, ok := utils.ParseIDKind(in.Source); ok {
			if ref, err := e.Store.Idx.LookupByID(kind, in.Source); err == nil {
				results = append(results, EntitySummary{ID: ref.ID, Kind: string(ref.Kind), Name: ref.Name, File: ref.File})
			}
		}
	}
	return results, nil
}

// unused import guard
var _ = fmt.Sprintf
