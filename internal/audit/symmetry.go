package audit

// SymmetryEntry pairs a planning-time audit rule with its post-plan
// counterpart. Every planning-time enforcement must have an equivalent
// rule that fires against actual entity state, so an intent flagged at
// planning is still caught at rest if the plan is evaded or abandoned
// mid-flight. The paired registry below is consumed by symmetry_test.go.
type SymmetryEntry struct {
	// PlanningRule is a short human identifier for the rule emitted by
	// planAuthoringFindings (or its helpers).
	PlanningRule string
	// PostPlanRule is the identifier of the equivalent post-plan rule
	// that lives in requirements.go, graph_rules.go, orphans.go, etc.
	PostPlanRule string
	// Description explains in one sentence what the two rules enforce.
	Description string
}

// SymmetryRegistry is the canonical planning↔post-plan pairing. Adding
// a rule to either side without its counterpart is a design bug; the
// symmetry test iterates this registry asserting both sides are
// wired into audit.Run.
var SymmetryRegistry = []SymmetryEntry{
	{
		PlanningRule: "plan_authoring.requirement_overlap_warn",
		PostPlanRule: "requirements.requirementOverlapFindings",
		Description:  "Overlap between a new/extended requirement and an existing active requirement must surface at plan time and at sync check time.",
	},
	{
		PlanningRule: "plan_authoring.rubber_stamp_ack_warn",
		PostPlanRule: "requirements.rubberStampOverlapFindings",
		Description:  "An audited overlap acknowledgement without semantic distinction text must be flagged both when the plan draft is authored and against the persisted requirement.",
	},
	{
		PlanningRule: "plan_authoring.contract_surface_coverage_warn",
		PostPlanRule: "requirements.requirementContractSurfaceFindings",
		Description:  "A requirement whose statement names a CLI/REST/screen/event surface must co-evolve with a contract — caught in the plan diff and against the active corpus.",
	},
	{
		PlanningRule: "plan_authoring.flow_coverage_warn",
		PostPlanRule: "graph_rules.contractFlowFindings",
		Description:  "A contract without a flow step referencing it must be flagged at plan time (flow lane omission) and at sync check time (active contract uncovered).",
	},
	{
		PlanningRule: "plan_authoring.requirement_lane_coverage_warn",
		PostPlanRule: "graph_rules.requirementTraceFindings",
		Description:  "Plan must carry enough requirements relative to its implementation changes; every non-requirement entity must itself link to a requirement at rest.",
	},
	{
		PlanningRule: "plan_authoring.task_claimed_by_change_warn",
		PostPlanRule: "plan_completion.planCompletionFindings",
		Description:  "Every task must be claimed by a change entry (planning) and every declared change must resolve to a real entity (post-plan).",
	},
}

// SymmetryByPlanningRule returns the registry entry for the given
// planning-rule identifier, or nil if none exists.
func SymmetryByPlanningRule(id string) *SymmetryEntry {
	for i := range SymmetryRegistry {
		if SymmetryRegistry[i].PlanningRule == id {
			return &SymmetryRegistry[i]
		}
	}
	return nil
}
