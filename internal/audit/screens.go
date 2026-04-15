package audit

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/tree"
)

// screenFindings reports drift between the SPA page files on disk
// and the screen contracts that should claim them. For every
// non-ignored .tsx file under web/src/pages, we check whether any
// contract entity with contract_kind=="screen" lists that path in
// its Files. Unowned files become WARN findings — devs can land a
// new page mid-feature without immediately blocking the session
// gate, but the warning surfaces as a reminder that a screen
// contract is missing.
//
// Test/story suffixes are skipped so future *.test.tsx and
// *.stories.tsx don't trip the gate. Files outside web/src/pages
// are out of scope — the orphan-file check (different category)
// handles those.
func screenFindings(all []model.EntityWithBody, t *tree.Tree) []Finding {
	if t == nil {
		return nil
	}

	// Build a set of every file path claimed by any screen contract.
	claimed := make(map[string]bool)
	for _, ewb := range all {
		c, ok := ewb.Entity.(*model.ContractEntity)
		if !ok {
			continue
		}
		if c.ContractKind != "screen" {
			continue
		}
		for _, fp := range c.GetBase().Files {
			claimed[fp] = true
		}
	}

	var out []Finding
	for path, n := range t.Nodes {
		if n == nil || n.Type != tree.TypeFile || n.Ignored {
			continue
		}
		if !strings.HasPrefix(path, "web/src/pages/") {
			continue
		}
		if !strings.HasSuffix(path, ".tsx") {
			continue
		}
		// Skip test and story files — they aren't first-class screens.
		if strings.HasSuffix(path, ".test.tsx") || strings.HasSuffix(path, ".stories.tsx") {
			continue
		}
		if claimed[path] {
			continue
		}
		out = append(out, Finding{
			Severity: SeverityWarning,
			Category: CatScreenUnclaimed,
			Path:     path,
			Message: fmt.Sprintf(
				"page %q has no owning screen contract — create one with 'syde add contract <name> --contract-kind screen --interaction-pattern render --wireframe \"<screen>...</screen>\" --file %s' or 'syde tree ignore %s' if it is not a screen",
				path, path, path,
			),
		})
	}
	return out
}
