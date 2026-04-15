package query

import (
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
	"github.com/feedloop/syde/internal/utils"
)

// ResolvedRelationship is a relationship with the target entity's summary.
type ResolvedRelationship struct {
	Type       string `json:"type"`
	Direction  string `json:"direction"` // "outbound" or "inbound"
	TargetID   string `json:"target_id"`
	TargetSlug string `json:"target_slug"`
	TargetName string `json:"target_name"`
	TargetKind string `json:"target_kind"`
	TargetFile string `json:"target_file"`
	TargetDesc string `json:"target_description"`
	Label      string `json:"label,omitempty"`
}

// ResolvedFileRef pairs an entity's source file reference with its
// summary tree data so clients can render "path — summary" inline.
type ResolvedFileRef struct {
	Path    string `json:"path"`
	Summary string `json:"summary"`
	Stale   bool   `json:"stale"`
	InTree  bool   `json:"in_tree"`
}

// ResolvedEntity is an entity with all its context resolved.
type ResolvedEntity struct {
	Entity        model.Entity           `json:"entity"`
	Body          string                 `json:"body,omitempty"`
	File          string                 `json:"file"`
	FileRefs      []ResolvedFileRef      `json:"file_refs,omitempty"`
	Relationships []ResolvedRelationship `json:"relationships,omitempty"`
	Learnings     []LearningSummary      `json:"learnings,omitempty"`
	Tasks         []TaskSummary          `json:"tasks,omitempty"`
	Decisions     []DecisionSummary      `json:"decisions,omitempty"`
	Suggested     []string               `json:"suggested_queries,omitempty"`
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
	entitySlug := b.CanonicalSlug()
	file := store.FS.RelativePath(b.Kind, entitySlug)

	result := &ResolvedEntity{
		Entity: entity,
		Body:   body,
		File:   file,
	}

	// Resolve file references against the summary tree (best effort;
	// missing tree just leaves file_refs empty).
	if len(b.Files) > 0 {
		if treeData, err := tree.Load(store.FS.Root); err == nil && treeData != nil {
			for _, fp := range b.Files {
				ref := ResolvedFileRef{Path: fp}
				if n := treeData.Get(fp); n != nil {
					ref.InTree = true
					ref.Summary = n.Summary
					ref.Stale = n.SummaryStale
				}
				result.FileRefs = append(result.FileRefs, ref)
			}
		}
	}

	// Resolve outbound relationships
	for _, rel := range b.Relationships {
		rr := ResolvedRelationship{
			Type:      rel.Type,
			Direction: "outbound",
			TargetID:  rel.Target,
			Label:     rel.Label,
		}
		// Target may be an ID (prefix_XXX) or a slug.
		if targetKind, ok := utils.ParseIDKind(rel.Target); ok {
			if ref, err := store.Idx.LookupByID(targetKind, rel.Target); err == nil {
				rr.TargetName = ref.Name
				rr.TargetKind = string(ref.Kind)
				rr.TargetFile = ref.File
				rr.TargetSlug = utils.Slugify(ref.Name) // FileRef doesn't carry the stored slug; best effort
			}
		} else {
			// Slug target — the string itself is the slug. Try to resolve name/kind.
			rr.TargetSlug = rel.Target
			if e, _, err := store.Get(rel.Target); err == nil {
				tb := e.GetBase()
				rr.TargetName = tb.Name
				rr.TargetKind = string(tb.Kind)
				rr.TargetFile = store.FS.RelativePath(tb.Kind, rel.Target)
			}
		}
		result.Relationships = append(result.Relationships, rr)
	}

	// Resolve inbound relationships. Index entries can key the target by
	// any of: entity ID, full slug (with -XXXX suffix), or bare name slug,
	// because callers store whichever form they typed. Query all three
	// and de-dupe by source.
	aliases := []string{b.ID}
	if b.Slug != "" {
		aliases = append(aliases, b.Slug)
		if base := utils.BaseSlug(b.Slug); base != "" && base != b.Slug {
			aliases = append(aliases, base)
		}
	}
	seenSource := make(map[string]bool)
	var inbound []struct {
		Type   string
		Source string
		Rel    storage.RelRef
	}
	for _, alias := range aliases {
		batch, _ := store.Idx.GetInbound(alias)
		for _, in := range batch {
			if seenSource[in.Source+"|"+in.Type] {
				continue
			}
			seenSource[in.Source+"|"+in.Type] = true
			inbound = append(inbound, in)
		}
	}
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
				rr.TargetSlug = utils.Slugify(ref.Name) // FileRef doesn't carry the stored slug; best effort
			}
		} else {
			rr.TargetSlug = in.Source
			if e, _, err := store.Get(in.Source); err == nil {
				tb := e.GetBase()
				rr.TargetName = tb.Name
				rr.TargetKind = string(tb.Kind)
				rr.TargetFile = store.FS.RelativePath(tb.Kind, in.Source)
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
					Status:   string(t.TaskStatus),
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
