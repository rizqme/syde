package audit

import (
	"testing"
)

func TestExtractSurfaces_CLI(t *testing.T) {
	cases := []struct {
		stmt     string
		wantNorm string
	}{
		{"The syde add requirement command shall block.", "syde add requirement"},
		{"When syde plan approve is invoked, the CLI shall do X.", "syde plan approve"},
		{"The syde sync check command shall exit non-zero.", "syde sync check"},
	}
	for _, tc := range cases {
		got := ExtractSurfaces(tc.stmt)
		found := false
		for _, s := range got {
			if s.Kind == SurfaceCLI && s.Normalised == tc.wantNorm {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%q: expected CLI surface %q in %+v", tc.stmt, tc.wantNorm, got)
		}
	}
}

func TestExtractSurfaces_CLI_RejectsEnglishWord(t *testing.T) {
	stmt := "The syde system shall ensure X."
	got := ExtractSurfaces(stmt)
	for _, s := range got {
		if s.Kind == SurfaceCLI {
			t.Errorf("%q: unexpected CLI surface %+v — 'system' is not an allowlisted subcommand", stmt, s)
		}
	}
}

func TestExtractSurfaces_REST(t *testing.T) {
	stmt := "The syded server shall return 200 from GET /api/projects and 404 from POST /api/unknown."
	got := ExtractSurfaces(stmt)
	want := map[string]bool{"get /api/projects": true, "post /api/unknown": true}
	for _, s := range got {
		if s.Kind == SurfaceREST {
			delete(want, s.Normalised)
		}
	}
	if len(want) > 0 {
		t.Errorf("missing REST surfaces: %+v in %+v", want, got)
	}
}

func TestExtractSurfaces_Screen(t *testing.T) {
	stmt := "When a plan is opened, the dashboard shall render the Plan Detail page."
	got := ExtractSurfaces(stmt)
	found := false
	for _, s := range got {
		if s.Kind == SurfaceScreen {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a screen surface, got %+v", got)
	}
}

func TestExtractSurfaces_Event(t *testing.T) {
	stmt := "When a user registers successfully, the auth service shall emit users.registered."
	got := ExtractSurfaces(stmt)
	found := false
	for _, s := range got {
		if s.Kind == SurfaceEvent {
			found = true
		}
	}
	if !found {
		t.Errorf("expected an event surface, got %+v", got)
	}
}

func TestExtractSurfaces_Empty(t *testing.T) {
	if got := ExtractSurfaces(""); got != nil {
		t.Errorf("empty statement should return nil, got %+v", got)
	}
	if got := ExtractSurfaces("   "); got != nil {
		t.Errorf("whitespace statement should return nil, got %+v", got)
	}
}

func TestContractCoversSurface(t *testing.T) {
	s := Surface{Kind: SurfaceCLI, Normalised: "syde add requirement"}
	if !ContractCoversSurface("syde add requirement <name> [--statement ...]", s) {
		t.Error("contract input should cover matching surface")
	}
	if ContractCoversSurface("syde plan create <name>", s) {
		t.Error("contract input should not cover non-matching surface")
	}
	if ContractCoversSurface("", s) {
		t.Error("empty contract input should never match")
	}
}
