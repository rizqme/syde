package query

import (
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/utils"
)

// ResolvedRelationship is a relationship with the target entity's summary.
type ResolvedRelationship struct {
	Type       string `json:"type"`
	Direction  string `json:"direction"` // "outbound" or "inbound"
	TargetID   string `json:"target_id"`
	TargetName string `json:"target_name"`
	TargetKind string `json:"target_kind"`
	TargetFile string `json:"target_file"`
	TargetDesc string `json:"target_description"`
	Label      string `json:"label,omitempty"`
}

// ResolvedEntity is an entity with all its context resolved.
type ResolvedEntity struct {
	Entity        model.Entity               `json:"entity"`
	Body          string                     `json:"body,omitempty"`
	File          string                     `json:"file"`
	Relationships []ResolvedRelationship     `json:"relationships,omitempty"`
	Learnings     []LearningSummary          `json:"learnings,omitempty"`
	Tasks         []TaskSummary              `json:"tasks,omitempty"`
	Decisions     []DecisionSummary          `json:"decisions,omitempty"`
	Suggested     []string                   `json:"suggested_queries,omitempty"`
}

// LearningSummary is a compact learning reference.
type LearningSummary struct {
	Name       string `json:"name"`
	Category   string `json:"category"`
	Confidence string `json:"confidence"`
	Desc       string `json:"description"`
	File       string `json:"file"`
}

// TaskSummary is a compact task reference.
type TaskSummary struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	File     string `json:"file"`
}

// DecisionSummary is a compact decision reference.
type DecisionSummary struct {
	Name      string `json:"name"`
	Statement string `json:"statement"`
	Category  string `json:"category"`
	File      string `json:"file"`
}

// Resolve builds a full ResolvedEntity for the given slug.
func Resolve(store *storage.Store, slug string) (*ResolvedEntity, error) {
	entity, body, err := store.Get(slug)
	if err != nil {
		return nil, err
	}
	b := entity.GetBase()
	entitySlug := utils.Slugify(b.Name)
	file := store.FS.RelativePath(b.Kind, entitySlug)

	result := &ResolvedEntity{
		Entity: entity,
		Body:   body,
		File:   file,
	}

	// Resolve outbound relationships
	for _, rel := range b.Relationships {
		rr := ResolvedRelationship{
			Type:      rel.Type,
			Direction: "outbound",
			TargetID:  rel.Target,
			Label:     rel.Label,
		}
		if targetKind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := store.Idx.LookupByID(targetKind, rel.Target); err == nil {
				rr.TargetName = ref.Name
				rr.TargetKind = string(ref.Kind)
				rr.TargetFile = ref.File
			}
		}
		result.Relationships = append(result.Relationships, rr)
	}

	// Resolve inbound relationships
	inbound, _ := store.Idx.GetInbound(b.ID)
	for _, in := range inbound {
		rr := ResolvedRelationship{
			Type:      in.Type,
			Direction: "inbound",
			TargetID:  in.Source,
			Label:     in.Rel.Label,
		}
		if sourceKind, ok := utils.ParseIDKind(in.Source); ok {
			if ref, err := store.Idx.LookupByID(sourceKind, in.Source); err == nil {
				rr.TargetName = ref.Name
				rr.TargetKind = string(ref.Kind)
				rr.TargetFile = ref.File
				rr.TargetDesc = "" // would need file read for desc
			}
		}
		result.Relationships = append(result.Relationships, rr)
	}

	// Resolve learnings referencing this entity
	learnings, _ := store.List(model.KindLearning)
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		for _, ref := range l.EntityRefs {
			if ref == b.ID || ref == entitySlug {
				result.Learnings = append(result.Learnings, LearningSummary{
					Name:       l.Name,
					Category:   string(l.Category),
					Confidence: string(l.ConfLevel),
					Desc:       l.Description,
					File:       store.FS.RelativePath(model.KindLearning, utils.Slugify(l.Name)),
				})
				break
			}
		}
	}

	// Resolve tasks referencing this entity
	tasks, _ := store.List(model.KindTask)
	for _, ewb := range tasks {
		t := ewb.Entity.(*model.TaskEntity)
		for _, ref := range t.EntityRefs {
			if ref == b.ID || ref == entitySlug {
				result.Tasks = append(result.Tasks, TaskSummary{
					Name:     t.Name,
					Status:   string(t.Status),
					Priority: string(t.Priority),
					File:     store.FS.RelativePath(model.KindTask, utils.Slugify(t.Name)),
				})
				break
			}
		}
	}

	// Resolve applicable decisions
	decisions, _ := store.List(model.KindDecision)
	for _, ewb := range decisions {
		d := ewb.Entity.(*model.DecisionEntity)
		for _, rel := range d.Relationships {
			if rel.Target == b.ID {
				result.Decisions = append(result.Decisions, DecisionSummary{
					Name:      d.Name,
					Statement: d.Statement,
					Category:  d.Category,
					File:      store.FS.RelativePath(model.KindDecision, utils.Slugify(d.Name)),
				})
				break
			}
		}
	}

	// Suggested queries
	result.Suggested = []string{
		"syde query --impacts " + entitySlug,
		"syde graph " + entitySlug + " --depth 2",
		"syde learn about " + entitySlug,
	}

	return result, nil
}
