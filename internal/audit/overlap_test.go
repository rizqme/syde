package audit

import (
	"strings"
	"testing"

	"github.com/feedloop/syde/internal/model"
)

// buildReqEntry is a tiny helper for the rubber-stamp test.
func buildReqEntry(name, slug, statement string, acks []model.AuditedOverlap) reqEntry {
	req := &model.RequirementEntity{
		BaseEntity: model.BaseEntity{
			Kind: model.KindRequirement,
			Name: name,
			Slug: slug,
		},
		Statement:         statement,
		RequirementStatus: model.RequirementActive,
		AuditedOverlaps:   acks,
	}
	return reqEntry{req: req, terms: SignificantTerms(statement)}
}

// TestRubberStampOverlapFindings_EmptyDistinctionErrors asserts the
// rule fires on an acknowledgement with no distinction text.
func TestRubberStampOverlapFindings_EmptyDistinctionErrors(t *testing.T) {
	r := buildReqEntry("A", "a", "The syde engine shall X.", []model.AuditedOverlap{
		{Slug: "other-slug", Distinction: ""},
	})
	findings := rubberStampOverlapFindings([]reqEntry{r})
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d: %+v", len(findings), findings)
	}
	if findings[0].Severity != SeverityFinding {
		t.Errorf("severity should be Finding, got %q", findings[0].Severity)
	}
	if !strings.Contains(findings[0].Message, "no distinction rationale") {
		t.Errorf("message should mention missing rationale, got %q", findings[0].Message)
	}
}

// TestRubberStampOverlapFindings_ShortDistinctionErrors asserts the
// rule fires on a distinction shorter than the 20-char minimum.
func TestRubberStampOverlapFindings_ShortDistinctionErrors(t *testing.T) {
	r := buildReqEntry("A", "a", "The syde engine shall X.", []model.AuditedOverlap{
		{Slug: "other-slug", Distinction: "too short"},
	})
	findings := rubberStampOverlapFindings([]reqEntry{r})
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d: %+v", len(findings), findings)
	}
	if !strings.Contains(findings[0].Message, "character") {
		t.Errorf("short-distinction message should mention character count, got %q", findings[0].Message)
	}
}

// TestRubberStampOverlapFindings_SubstantiveDistinctionPasses asserts
// the rule is silent when the distinction carries real reasoning.
func TestRubberStampOverlapFindings_SubstantiveDistinctionPasses(t *testing.T) {
	r := buildReqEntry("A", "a", "The syde engine shall X.", []model.AuditedOverlap{
		{Slug: "other-slug", Distinction: "This rule applies at sync check time while the other applies at plan approval time — different triggers, different cadence."},
	})
	findings := rubberStampOverlapFindings([]reqEntry{r})
	if len(findings) != 0 {
		t.Errorf("substantive distinction should not emit findings, got %+v", findings)
	}
}
