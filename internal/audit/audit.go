// Package audit centralizes the syde model-health checks that back
// `syde validate`, `syde sync check`, and the `syde files` commands.
// Each check produces Finding values with a Severity and Category so
// callers can group, filter, and render them consistently.
package audit

import (
	"path/filepath"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/tree"
)

type Severity string

// SeverityFinding is the sole audit severity. The project is strict —
// every finding blocks sync-check, validate, and plan-complete. Older
// aliases SeverityError, SeverityWarning, and SeverityHint collapse to
// this single constant so existing call-sites keep compiling, but
// semantically there is one level.
const SeverityFinding Severity = "finding"

// Legacy aliases — all equal to SeverityFinding. Kept as a single
// name each so the audit producers in this package can be migrated
// incrementally; nothing downgrades a finding.
const (
	SeverityError   = SeverityFinding
	SeverityWarning = SeverityFinding
	SeverityHint    = SeverityFinding
)

type Category string

const (
	CatMissingField     Category = "missing_field"
	CatRecommendedField Category = "recommended_field"
	CatBrokenRel        Category = "broken_rel"
	CatContractOwner    Category = "contract_owner"
	CatSystemCycle      Category = "system_cycle"
	CatComponentCycle   Category = "component_cycle"
	CatFileNotInTree    Category = "file_not_in_tree"
	CatOrphanFile       Category = "orphan_file"
	CatFileDrift        Category = "file_drift"
	CatPlanPhase        Category = "plan_phase"
	CatScreenUnclaimed  Category = "screen_unclaimed"
	CatRequirement      Category = "requirement_lifecycle"
	CatTraceability     Category = "requirement_traceability"
	CatHierarchy        Category = "hierarchy"
	CatContractFlow     Category = "contract_flow"
	CatPlanCompletion   Category = "plan_completion"
	CatPlanAuthoring    Category = "plan_authoring"
)

// Finding is a single audit result. Path is set for file-scoped findings
// (orphans, drift). EntityKind+EntitySlug are set for entity-scoped
// findings. Field is set when the finding points at a specific entity
// field (missing, recommended).
type Finding struct {
	Severity   Severity         `json:"severity"`
	Category   Category         `json:"category"`
	Message    string           `json:"message"`
	EntityKind model.EntityKind `json:"entity_kind,omitempty"`
	EntitySlug string           `json:"entity_slug,omitempty"`
	EntityName string           `json:"entity_name,omitempty"`
	Field      string           `json:"field,omitempty"`
	Path       string           `json:"path,omitempty"`
}

// Report is the full set of findings plus counts, returned from Run.
type Report struct {
	Findings []Finding `json:"findings"`
	Entities int       `json:"entities"`
}

// Counts returns (findings, 0, 0) — the second and third slots are
// vestigial from the old Error/Warning/Hint split and preserved so
// existing call sites that unpack three values keep compiling. All
// findings are at the single SeverityFinding level.
func (r *Report) Counts() (int, int, int) {
	return len(r.Findings), 0, 0
}

// HasErrors reports whether the report has any findings. There is no
// non-blocking severity, so any finding makes the gate fail.
func (r *Report) HasErrors() bool {
	return len(r.Findings) > 0
}

// Options toggles which categories Run evaluates. Zero value runs
// everything; callers asking for a narrower audit (e.g. `files orphans`)
// can disable irrelevant checks.
type Options struct {
	SkipEntityValidation bool
	SkipRelationships    bool
	SkipCycles           bool
	SkipTreeConsistency  bool
	SkipOrphans          bool
	SkipDrift            bool
}

// Run evaluates every enabled check against the given store (entities)
// and optional tree (file-level checks). A nil tree disables all tree-
// scoped findings without erroring — first-run projects may not have a
// tree yet.
func Run(store *storage.Store, t *tree.Tree, opts Options) (*Report, error) {
	all, err := store.ListAll()
	if err != nil {
		return nil, err
	}
	rep := &Report{Entities: len(all)}

	if !opts.SkipEntityValidation {
		rep.Findings = append(rep.Findings, entityFieldFindings(all)...)
	}
	if !opts.SkipRelationships {
		rep.Findings = append(rep.Findings, relationshipFindings(all)...)
		rep.Findings = append(rep.Findings, requirementTraceFindings(all)...)
		rep.Findings = append(rep.Findings, requirementFanoutFindings(all)...)
		rep.Findings = append(rep.Findings, goodRequirementFindings(all)...)
		rep.Findings = append(rep.Findings, coverageFindings(all)...)
		rep.Findings = append(rep.Findings, hierarchyFindings(all)...)
		rep.Findings = append(rep.Findings, contractFlowFindings(all)...)
	}
	if !opts.SkipCycles {
		rep.Findings = append(rep.Findings, cycleFindings(all)...)
		rep.Findings = append(rep.Findings, planPhaseFindings(all)...)
		rep.Findings = append(rep.Findings, requirementFindings(all)...)
		rep.Findings = append(rep.Findings, requirementOverlapFindings(all)...)
		rep.Findings = append(rep.Findings, requirementContractSurfaceFindings(all)...)
		rep.Findings = append(rep.Findings, bidirectionalFindings(all)...)
		rep.Findings = append(rep.Findings, requirementStaleFindings(all, filepath.Dir(store.FS.Root))...)
		rep.Findings = append(rep.Findings, planAuthoringFindings(all)...)
		rep.Findings = append(rep.Findings, planCompletionFindings(all)...)
	}
	if t != nil && !opts.SkipTreeConsistency {
		rep.Findings = append(rep.Findings, screenFindings(all, t)...)
	}
	if t != nil && !opts.SkipTreeConsistency {
		rep.Findings = append(rep.Findings, fileRefFindings(all, t)...)
	}
	if t != nil && !opts.SkipOrphans {
		rep.Findings = append(rep.Findings, orphanFindings(all, t)...)
	}
	if t != nil && !opts.SkipDrift {
		rep.Findings = append(rep.Findings, driftFindings(all, t)...)
	}

	return rep, nil
}
