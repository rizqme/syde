package audit

import (
	"regexp"
	"sort"
	"strings"
)

// SurfaceKind categorises contract-governed surfaces that can be
// extracted from a requirement statement's prose.
type SurfaceKind string

const (
	SurfaceCLI    SurfaceKind = "cli"
	SurfaceREST   SurfaceKind = "rest"
	SurfaceScreen SurfaceKind = "screen"
	SurfaceEvent  SurfaceKind = "event"
)

// Surface identifies one extracted contract-governed surface mention.
// Raw is the literal match from the statement; Normalised is a
// canonical form used to match against contract.Input so that minor
// phrasing differences do not miss a covering contract.
type Surface struct {
	Kind       SurfaceKind
	Raw        string
	Normalised string
}

var (
	// CLI pattern: `syde <sub>` or `syde <sub> <subsub>` anchored at a
	// word boundary. Matches both Ubiquitous ("The syde add ... shall")
	// and Event-driven ("When syde add ... is invoked") phrasings.
	surfaceCLIRegex = regexp.MustCompile(`(?i)\bsyde\s+([a-z][a-z0-9-]+(?:\s+[a-z0-9-]+)?)`)

	// REST pattern: HTTP verb + meaningful path. Path must start with
	// a slash + alphanumeric segment and end on an alphanumeric or
	// closing-brace character so bare roots ("GET /", "GET /api/")
	// are not matched. The regex still handles params ({id}, :slug).
	surfaceRESTRegex = regexp.MustCompile(`\b(GET|POST|PUT|DELETE|PATCH)\s+(/[a-zA-Z0-9][a-zA-Z0-9/_\-:{}]*[a-zA-Z0-9}])`)

	// Screen pattern: "the dashboard shall render X" / "display X" / "show X"
	// where X is a capitalised proper noun phrase (up to 3 words), e.g.
	// "Plans inbox", "Plan Detail". Lower-case fragments are not
	// considered screen names and are intentionally not matched.
	surfaceScreenRegex = regexp.MustCompile(`\b(?:the\s+)?(?:dashboard|SPA|UI)\s+shall\s+(?:render|display|show|open|navigate\s+to)\s+(?:the\s+|a\s+)?([A-Z][A-Za-z]+(?:\s+[A-Za-z]+){0,2}(?:\s+(?:inbox|page|panel|view|screen|detail|list|canvas))?)`)

	// Event pattern: only a dot-delimited or underscore-delimited topic
	// identifier (e.g. `users.registered`, `plan_completed`). Prose
	// "when the plan is approved" conditions are intentionally not
	// considered events — that is behaviour, not a pub-sub topic.
	surfaceEventRegex = regexp.MustCompile(`\b([a-z][a-z0-9]+(?:[._][a-z0-9]+)+)\b`)
)

// cliSecondWordAllowlist names every real CLI sub-subcommand or
// entity-kind argument that is valid as the 2nd word after "syde".
// Anything else gets truncated to the single-word surface.
var cliSecondWordAllowlist = map[string]bool{
	// Plan subcommands
	"create": true, "update": true, "approve": true, "show": true,
	"list": true, "complete": true, "remove": true, "open": true,
	"check": true, "execute": true, "estimate": true,
	"add-change": true, "add-phase": true, "remove-change": true,
	"show-changes": true,
	// Task subcommands
	"start": true, "done": true, "block": true, "sub": true,
	// Tree subcommands
	"scan": true, "status": true, "changes": true, "summarize": true,
	"ignore": true, "context": true,
	// Files subcommands
	"coverage": true, "orphans": true,
	// Install subcommand
	"skill": true,
	// Entity kinds (for syde add / remove / get / update / list)
	"component": true, "contract": true, "concept": true, "flow": true,
	"system": true, "task": true, "requirement": true, "decision": true,
	"plan": true,
}

// cliSecondWordBlocklist strips English filler that the regex
// accidentally captures as a second CLI word. Without it,
// "When syde add is invoked" would extract "syde add is" as
// a surface.
var cliSecondWordBlocklist = map[string]bool{
	"is":        true,
	"shall":     true,
	"succeeds":  true,
	"fails":     true,
	"invokes":   true,
	"returns":   true,
	"completes": true,
	"does":      true,
	"has":       true,
	"will":      true,
	"must":      true,
	"cannot":    true,
	"when":      true,
	"then":      true,
	"while":     true,
	"where":     true,
	"emits":     true,
	"prints":    true,
	"accepts":   true,
	"refuses":   true,
	"blocks":    true,
	"warns":     true,
	"errors":    true,
	"creates":   true,
	"detects":   true,
	"rejects":   true,
	"parses":    true,
	"runs":      true,
	"checks":    true,
	"gets":      true,
	"reads":     true,
	"writes":    true,
}

// cliSubcommandAllowlist filters out CLI matches that are actually
// English words or tool names (not syde subcommands). Without this,
// phrases like "syde add requirement shall" would surface matches
// for non-commands.
var cliSubcommandAllowlist = map[string]bool{
	"add":         true,
	"get":         true,
	"update":      true,
	"remove":      true,
	"list":        true,
	"query":       true,
	"search":      true,
	"graph":       true,
	"status":      true,
	"context":     true,
	"sync":        true,
	"scan":        true,
	"tree":        true,
	"files":       true,
	"plan":        true,
	"task":        true,
	"constraints": true,
	"init":        true,
	"install":     true,
	"open":        true,
	"reindex":     true,
	"server":      true,
	"validate":    true,
	"wireframe":   true,
}

// ExtractSurfaces returns every contract-governed surface mention in
// a requirement statement, in a deterministic order (CLI → REST →
// screen → event). Duplicate surfaces are de-duplicated by
// Normalised identity.
func ExtractSurfaces(statement string) []Surface {
	if strings.TrimSpace(statement) == "" {
		return nil
	}
	seen := map[string]bool{}
	var out []Surface

	for _, m := range surfaceCLIRegex.FindAllStringSubmatch(statement, -1) {
		words := strings.Fields(strings.ToLower(m[1]))
		if len(words) == 0 {
			continue
		}
		if !cliSubcommandAllowlist[words[0]] {
			continue
		}
		// Drop a second word that is really English filler captured
		// by the regex (e.g. "syde add is invoked" → "syde add").
		if len(words) > 1 && cliSecondWordBlocklist[words[1]] {
			words = words[:1]
		}
		// Require the second word to be an actual CLI sub-subcommand
		// or entity kind. Otherwise fall back to the single-word
		// surface so phrases like "syde plan authoring" become
		// "syde plan" instead of a bogus two-word surface.
		if len(words) > 1 && !cliSecondWordAllowlist[words[1]] {
			words = words[:1]
		}
		raw := "syde " + strings.Join(words, " ")
		norm := raw
		key := string(SurfaceCLI) + "|" + norm
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, Surface{Kind: SurfaceCLI, Raw: raw, Normalised: norm})
	}

	for _, m := range surfaceRESTRegex.FindAllStringSubmatch(statement, -1) {
		verb := strings.ToUpper(m[1])
		path := m[2]
		raw := verb + " " + path
		norm := strings.ToLower(verb + " " + path)
		key := string(SurfaceREST) + "|" + norm
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, Surface{Kind: SurfaceREST, Raw: raw, Normalised: norm})
	}

	for _, m := range surfaceScreenRegex.FindAllStringSubmatch(statement, -1) {
		raw := strings.TrimSpace(m[0])
		norm := strings.ToLower(strings.TrimSpace(m[1]))
		// Truncate trailing English stopwords ("as", "a", "with") so
		// the normalised form matches a screen contract's name even
		// when the sentence continues with follow-on prose.
		words := strings.Fields(norm)
		trimmed := make([]string, 0, len(words))
		for _, w := range words {
			if w == "as" || w == "a" || w == "with" || w == "to" || w == "for" || w == "of" {
				break
			}
			trimmed = append(trimmed, w)
		}
		if len(trimmed) == 0 {
			continue
		}
		norm = strings.Join(trimmed, " ")
		key := string(SurfaceScreen) + "|" + norm
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, Surface{Kind: SurfaceScreen, Raw: raw, Normalised: norm})
	}

	for _, m := range surfaceEventRegex.FindAllStringSubmatch(statement, -1) {
		raw := strings.TrimSpace(m[1])
		norm := strings.ToLower(strings.TrimSpace(m[1]))
		// Skip common file-path-like tokens that aren't events.
		if strings.Contains(norm, "/") {
			continue
		}
		// A genuine pub-sub event identifier looks like
		// `users.registered` (dotted). Underscore-only tokens like
		// `belongs_to` are YAML field names / relationship types, not
		// events — skip them.
		if !strings.Contains(norm, ".") {
			continue
		}
		// Skip filename-like tokens (hooks.json, syde.yaml, tree.yaml)
		// whose extension is a common data format.
		parts := strings.Split(norm, ".")
		if ext := parts[len(parts)-1]; len(parts) >= 2 && (ext == "json" || ext == "yaml" || ext == "yml" || ext == "md" || ext == "go" || ext == "ts" || ext == "tsx" || ext == "js") {
			continue
		}
		key := string(SurfaceEvent) + "|" + norm
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, Surface{Kind: SurfaceEvent, Raw: raw, Normalised: norm})
	}

	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Kind != out[j].Kind {
			return out[i].Kind < out[j].Kind
		}
		return out[i].Normalised < out[j].Normalised
	})
	return out
}

// ContractCoversSurface reports whether the given contract's input
// signature covers the requirement-statement surface. The check is a
// simple case-insensitive substring match of the normalised surface
// against the contract input — good enough to catch the obvious
// cases (contract `input: "syde add requirement"` covers a
// requirement naming that invocation) without over-matching.
func ContractCoversSurface(contractInput string, surface Surface) bool {
	if contractInput == "" || surface.Normalised == "" {
		return false
	}
	hay := strings.ToLower(contractInput)
	return strings.Contains(hay, surface.Normalised)
}
