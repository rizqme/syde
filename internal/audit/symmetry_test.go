package audit

import (
	"strings"
	"testing"
)

// TestSymmetryRegistryNonEmpty asserts the registry is populated.
// Without it, the symmetry principle is decorative.
func TestSymmetryRegistryNonEmpty(t *testing.T) {
	if len(SymmetryRegistry) == 0 {
		t.Fatal("SymmetryRegistry must list at least one planning↔post-plan pair")
	}
}

// TestSymmetryRegistryEntriesWellFormed asserts every registry entry
// has both sides of the pair populated and a description.
func TestSymmetryRegistryEntriesWellFormed(t *testing.T) {
	for _, e := range SymmetryRegistry {
		if strings.TrimSpace(e.PlanningRule) == "" {
			t.Errorf("entry has empty PlanningRule: %+v", e)
		}
		if strings.TrimSpace(e.PostPlanRule) == "" {
			t.Errorf("entry %q missing PostPlanRule", e.PlanningRule)
		}
		if strings.TrimSpace(e.Description) == "" {
			t.Errorf("entry %q missing Description", e.PlanningRule)
		}
	}
}

// TestSymmetryLookup spot-checks the lookup helper returns a non-nil
// entry for a known rule and nil for an unknown one.
func TestSymmetryLookup(t *testing.T) {
	known := SymmetryRegistry[0].PlanningRule
	if SymmetryByPlanningRule(known) == nil {
		t.Errorf("lookup for known rule %q returned nil", known)
	}
	if SymmetryByPlanningRule("plan_authoring.does_not_exist") != nil {
		t.Error("lookup for unknown rule should return nil")
	}
}
