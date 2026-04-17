package audit

import "testing"

// TestSingleSeverityLevel guards against accidental reintroduction of
// a Warning or Hint severity. The project is strict — every finding
// blocks, and the severity enum collapses to one level.
func TestSingleSeverityLevel(t *testing.T) {
	if SeverityError != SeverityFinding {
		t.Errorf("SeverityError must alias SeverityFinding; got %q vs %q", SeverityError, SeverityFinding)
	}
	if SeverityWarning != SeverityFinding {
		t.Errorf("SeverityWarning must alias SeverityFinding; got %q vs %q", SeverityWarning, SeverityFinding)
	}
	if SeverityHint != SeverityFinding {
		t.Errorf("SeverityHint must alias SeverityFinding; got %q vs %q", SeverityHint, SeverityFinding)
	}
	if SeverityFinding != "finding" {
		t.Errorf("SeverityFinding string value should be %q, got %q", "finding", SeverityFinding)
	}
}

// TestReportHasErrorsAnyFinding asserts Report.HasErrors returns true
// for any finding, regardless of the (aliased) severity used.
func TestReportHasErrorsAnyFinding(t *testing.T) {
	r := &Report{Findings: []Finding{{Severity: SeverityFinding}}}
	if !r.HasErrors() {
		t.Error("HasErrors should return true for any finding")
	}
	empty := &Report{}
	if empty.HasErrors() {
		t.Error("HasErrors on empty Report should be false")
	}
}
