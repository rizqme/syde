package cli

import (
	"strings"
	"time"

	"github.com/feedloop/syde/internal/model"
)

type requirementCapture struct {
	Name          string
	Statement     string
	Source        string
	SourceRef     string
	Rationale     string
	ApprovedAt    string
	Relationships []model.Relationship
	Body          string
}

func createRequirementIfMissing(store *writeClient, cap requirementCapture) (*model.RequirementEntity, bool, string, error) {
	if cap.SourceRef != "" {
		if existing, file, ok := findRequirementBySourceRef(store, cap.Source, cap.SourceRef); ok {
			return existing, false, file, nil
		}
	}

	now := time.Now().UTC().Format(time.RFC3339)
	approvedAt := cap.ApprovedAt
	if approvedAt == "" {
		approvedAt = now
	}
	source := cap.Source
	if source == "" {
		source = "manual"
	}

	req := &model.RequirementEntity{
		BaseEntity: model.BaseEntity{
			Kind:          model.KindRequirement,
			Name:          cap.Name,
			Relationships: cap.Relationships,
		},
		Statement:         cap.Statement,
		Source:            source,
		SourceRef:         cap.SourceRef,
		RequirementStatus: model.RequirementActive,
		Rationale:         cap.Rationale,
		ApprovedAt:        approvedAt,
	}
	if !hasRelationshipType(req.Relationships, model.RelBelongsTo) {
		if parent := rootSystemTarget(store); parent != "" {
			req.Relationships = append(req.Relationships, model.Relationship{
				Target: parent,
				Type:   model.RelBelongsTo,
			})
		}
	}

	file, err := store.Create(req, cap.Body)
	if err != nil {
		return nil, false, "", err
	}
	return req, true, file, nil
}

func findRequirementBySourceRef(store *writeClient, source, sourceRef string) (*model.RequirementEntity, string, bool) {
	if strings.TrimSpace(sourceRef) == "" {
		return nil, "", false
	}
	requirements, err := store.List(model.KindRequirement)
	if err != nil {
		return nil, "", false
	}
	for _, ewb := range requirements {
		req, ok := ewb.Entity.(*model.RequirementEntity)
		if !ok {
			continue
		}
		if req.SourceRef != sourceRef {
			continue
		}
		if source != "" && req.Source != source {
			continue
		}
		return req, store.FS.RelativePath(model.KindRequirement, req.CanonicalSlug()), true
	}
	return nil, "", false
}

func rootSystemTarget(store *writeClient) string {
	systems, err := store.List(model.KindSystem)
	if err != nil {
		return ""
	}
	if len(systems) == 0 {
		return ""
	}
	for _, ewb := range systems {
		b := ewb.Entity.GetBase()
		if !hasRelationshipType(b.Relationships, model.RelBelongsTo) {
			return b.CanonicalSlug()
		}
	}
	return systems[0].Entity.GetBase().CanonicalSlug()
}

func hasRelationshipType(rels []model.Relationship, relType string) bool {
	for _, rel := range rels {
		if rel.Type == relType {
			return true
		}
	}
	return false
}

func appendRelationshipOnce(rels []model.Relationship, rel model.Relationship) []model.Relationship {
	if rel.Target == "" || rel.Type == "" {
		return rels
	}
	for _, existing := range rels {
		if existing.Target == rel.Target && existing.Type == rel.Type && existing.Label == rel.Label {
			return rels
		}
	}
	return append(rels, rel)
}

func requirementName(prefix, text string) string {
	oneLine := strings.Join(strings.Fields(text), " ")
	if oneLine == "" {
		return prefix
	}
	const max = 72
	if len(oneLine) > max {
		oneLine = oneLine[:max-3] + "..."
	}
	return prefix + ": " + oneLine
}

// parseAuditedOverlaps turns repeatable --audited flag values of the form
// "slug" or "slug:distinction" into AuditedOverlap structs. The distinction
// rationale is required by the audit engine; the CLI allows legacy slug-only
// input but the downstream audit rule flags it.
func parseAuditedOverlaps(values []string) []model.AuditedOverlap {
	if len(values) == 0 {
		return nil
	}
	out := make([]model.AuditedOverlap, 0, len(values))
	for _, v := range values {
		if v == "" {
			continue
		}
		slug, distinction, _ := strings.Cut(v, ":")
		slug = strings.TrimSpace(slug)
		if slug == "" {
			continue
		}
		out = append(out, model.AuditedOverlap{
			Slug:        slug,
			Distinction: strings.TrimSpace(distinction),
		})
	}
	return out
}
