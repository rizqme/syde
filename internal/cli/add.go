package cli

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
)

var (
	addDescription string
	addPurpose     string
	addTags        []string
	addFiles       []string
	addBody        string
	addRels        []string
	addNotes       []string
	// Component
	addResponsibility  string
	addCapabilities    []string
	addBoundaries      string
	addBehaviorSummary string
	addInteractionSum  string
	addDataHandling    string
	addScalingNotes    string
	// Contract
	addContractKind       string
	addInteractionPattern string
	addProtocolNotes      string
	addInput              string
	addInputParams        []string
	addOutput             string
	addOutputParams       []string
	addConstraintsText    string
	addVersioningNotes    string
	addContractWireframe  string
	// Concept
	addMeaning           string
	addStructureNotes    string
	addLifecycle         string
	addInvariants        string
	addDataSensitivity   string
	addConceptAttributes []string
	addConceptActions    []string
	// Flow
	addTrigger      string
	addGoal         string
	addNarrative    string
	addHappyPath    string
	addEdgeCases    string
	addFailureModes string
	addPerfNotes    string
	// Decision
	addCategory     string
	addStatement    string
	addRationale    string
	addAlternatives string
	addTradeoffs    string
	addConsequences string
	addSupersedes   string
	// Requirement
	addReqSource         string
	addReqSourceRef      string
	addReqStatus         string
	addReqAcceptance     string
	addReqSupersededBy   string
	addReqObsoleteReason string
	addReqApprovedAt     string
	// System
	addSysContext     string
	addSysScope       string
	addSysPrinciples  string
	addSysQuality     string
	addSysAssumptions string
)

// parseContractParams converts "path|type|description" specs into ContractParam slices.
func parseContractParams(specs []string) []model.ContractParam {
	var out []model.ContractParam
	for _, s := range specs {
		if s == "" {
			continue
		}
		if p, ok := model.ParseContractParam(s); ok {
			out = append(out, p)
		}
	}
	return out
}

// parseConceptAttributes converts "name|type|description" specs into
// ConceptAttribute slices. Empty specs are skipped; malformed specs
// (empty name) are silently dropped — the validator later flags
// concepts with zero attributes so nothing slips through unnoticed.
func parseConceptAttributes(specs []string) []model.ConceptAttribute {
	var out []model.ConceptAttribute
	for _, s := range specs {
		if s == "" {
			continue
		}
		if a, ok := model.ParseConceptAttribute(s); ok {
			out = append(out, a)
		}
	}
	return out
}

// parseConceptActions converts "name|description" specs into
// ConceptAction slices.
func parseConceptActions(specs []string) []model.ConceptAction {
	var out []model.ConceptAction
	for _, s := range specs {
		if s == "" {
			continue
		}
		if a, ok := model.ParseConceptAction(s); ok {
			out = append(out, a)
		}
	}
	return out
}

func validEntityKinds() string {
	kinds := model.AllEntityKinds()
	names := make([]string, 0, len(kinds))
	for _, kind := range kinds {
		names = append(names, string(kind))
	}
	return strings.Join(names, ", ")
}

func parseRefList(specs ...string) []string {
	var out []string
	for _, spec := range specs {
		for _, part := range strings.Split(spec, ",") {
			ref := strings.TrimSpace(part)
			if ref != "" {
				out = append(out, ref)
			}
		}
	}
	return out
}

func appendRefOnce(refs []string, ref string) []string {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return refs
	}
	for _, existing := range refs {
		if existing == ref {
			return refs
		}
	}
	return append(refs, ref)
}

func parseRequirementStatus(raw string) (model.RequirementStatus, error) {
	status := model.RequirementStatus(strings.TrimSpace(strings.ToLower(raw)))
	if status == "" {
		return model.RequirementActive, nil
	}
	switch status {
	case model.RequirementActive, model.RequirementSuperseded, model.RequirementObsolete:
		return status, nil
	default:
		return "", fmt.Errorf("invalid requirement status %q: expected active, superseded, or obsolete", raw)
	}
}

// parseRelSpec parses a --add-rel value as either "target:type" or
// "target:type:label". Label is used for cardinality on relates_to
// ("one-to-many" and friends) or any other free-form relationship
// annotation. Two-part form is preserved for back-compat with every
// existing entity on disk. Errors carry the original spec so the
// user can see exactly what was rejected.
func parseRelSpec(spec string) (model.Relationship, error) {
	parts := strings.SplitN(spec, ":", 3)
	if len(parts) < 2 {
		return model.Relationship{}, fmt.Errorf("invalid --add-rel %q: expected target:type or target:type:label", spec)
	}
	r := model.Relationship{
		Target: strings.TrimSpace(parts[0]),
		Type:   strings.TrimSpace(parts[1]),
	}
	if r.Target == "" || r.Type == "" {
		return model.Relationship{}, fmt.Errorf("invalid --add-rel %q: target and type must be non-empty", spec)
	}
	if len(parts) == 3 {
		r.Label = strings.TrimSpace(parts[2])
	}
	return r, nil
}

var addCmd = &cobra.Command{
	Use:   "add <kind> <name>",
	Short: "Add a new entity to the design model",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		kindStr := strings.ToLower(args[0])
		name := args[1]

		kind, ok := model.ParseEntityKind(kindStr)
		if !ok {
			return fmt.Errorf("unknown entity kind: %s\nValid: %s", kindStr, validEntityKinds())
		}

		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity := model.NewEntityForKind(kind)
		b := entity.GetBase()
		b.Name = name
		b.Description = addDescription
		b.Purpose = addPurpose
		b.Tags = addTags
		b.Files = addFiles
		b.Notes = addNotes
		for _, spec := range addRels {
			if spec == "" {
				continue
			}
			rel, err := parseRelSpec(spec)
			if err != nil {
				return err
			}
			b.Relationships = append(b.Relationships, rel)
		}

		switch v := entity.(type) {
		case *model.ComponentEntity:
			v.Responsibility = addResponsibility
			v.Capabilities = addCapabilities
			v.Boundaries = addBoundaries
			v.BehaviorSummary = addBehaviorSummary
			v.InteractionSummary = addInteractionSum
			v.DataHandling = addDataHandling
			v.ScalingNotes = addScalingNotes
		case *model.ContractEntity:
			v.ContractKind = addContractKind
			v.InteractionPattern = addInteractionPattern
			v.ProtocolNotes = addProtocolNotes
			v.Input = addInput
			v.InputParameters = parseContractParams(addInputParams)
			v.Output = addOutput
			v.OutputParameters = parseContractParams(addOutputParams)
			v.Constraints = addConstraintsText
			v.VersioningNotes = addVersioningNotes
			v.Wireframe = addContractWireframe
		case *model.ConceptEntity:
			v.Meaning = addMeaning
			v.StructureNotes = addStructureNotes
			v.Lifecycle = addLifecycle
			v.Invariants = addInvariants
			v.DataSensitivity = addDataSensitivity
			v.Attributes = parseConceptAttributes(addConceptAttributes)
			v.Actions = parseConceptActions(addConceptActions)
		case *model.FlowEntity:
			v.Trigger = addTrigger
			v.Goal = addGoal
			v.Narrative = addNarrative
			v.HappyPath = addHappyPath
			v.EdgeCases = addEdgeCases
			v.FlowFailureModes = addFailureModes
			v.PerformanceNotes = addPerfNotes
		case *model.DecisionEntity:
			v.Category = addCategory
			v.Statement = addStatement
			v.Rationale = addRationale
			v.AlternativesConsidered = addAlternatives
			v.Tradeoffs = addTradeoffs
			v.Consequences = addConsequences
			v.Supersedes = addSupersedes
		case *model.RequirementEntity:
			status, err := parseRequirementStatus(addReqStatus)
			if err != nil {
				return err
			}
			v.Statement = addStatement
			v.Source = addReqSource
			if v.Source == "" {
				v.Source = "manual"
			}
			v.SourceRef = addReqSourceRef
			v.RequirementStatus = status
			v.Rationale = addRationale
			v.AcceptanceCriteria = addReqAcceptance
			v.Supersedes = parseRefList(addSupersedes)
			v.SupersededBy = parseRefList(addReqSupersededBy)
			v.ObsoleteReason = addReqObsoleteReason
			v.ApprovedAt = addReqApprovedAt
			if !hasRelationshipType(b.Relationships, model.RelBelongsTo) {
				if parent := rootSystemTarget(store); parent != "" {
					b.Relationships = append(b.Relationships, model.Relationship{Target: parent, Type: model.RelBelongsTo})
				}
			}
		case *model.SystemEntity:
			v.Context = addSysContext
			v.Scope = addSysScope
			v.DesignPrinciples = addSysPrinciples
			v.QualityGoals = addSysQuality
			v.Assumptions = addSysAssumptions
		}

		filePath, err := store.Create(entity, addBody)
		if err != nil {
			return fmt.Errorf("create entity: %w", err)
		}

		fmt.Printf("Created %s: %s\n", kind, name)
		fmt.Printf("  ID: %s\n", b.ID)
		fmt.Printf("  File: %s\n", filePath)

		return nil
	},
}

func init() {
	f := addCmd.Flags()
	// Base
	f.StringVar(&addDescription, "description", "", "entity description")
	f.StringVar(&addPurpose, "purpose", "", "why the entity exists")
	f.StringSliceVar(&addTags, "tag", nil, "entity tags")
	// File / note / capability / rel use StringArrayVar (literal-valued,
	// repeatable) NOT StringSliceVar — slice splits on commas inside the
	// value, which corrupts paths, sentences with commas, and relationship
	// targets.
	f.StringArrayVar(&addFiles, "file", nil, "concrete source file path (repeatable)")
	f.StringArrayVar(&addNotes, "note", nil, "informal note (repeatable)")
	f.StringVar(&addBody, "body", "", "markdown body content")
	f.StringArrayVar(&addRels, "add-rel", nil, "add relationship (target:type) — repeatable")
	// Component
	f.StringVar(&addResponsibility, "responsibility", "", "component responsibility")
	f.StringArrayVar(&addCapabilities, "capability", nil, "component capability (repeatable)")
	f.StringVar(&addBoundaries, "boundaries", "", "component boundaries")
	f.StringVar(&addBehaviorSummary, "behavior-summary", "", "behavior summary")
	f.StringVar(&addInteractionSum, "interaction-summary", "", "interaction summary")
	f.StringVar(&addDataHandling, "data-handling", "", "data handling")
	f.StringVar(&addScalingNotes, "scaling-notes", "", "scaling notes")
	// Contract
	f.StringVar(&addContractKind, "contract-kind", "", "contract kind (api/event/command/query)")
	f.StringVar(&addInteractionPattern, "interaction-pattern", "", "interaction pattern")
	f.StringVar(&addProtocolNotes, "protocol-notes", "", "protocol notes")
	f.StringVar(&addInput, "input", "", "contract invocation signature (e.g. 'GET /api/projects', 'syde plan create <name>')")
	f.StringArrayVar(&addInputParams, "input-parameter", nil, "input parameter 'path|type|description' (repeatable)")
	f.StringVar(&addOutput, "output", "", "contract output signature / response shape")
	f.StringArrayVar(&addOutputParams, "output-parameter", nil, "output parameter 'path|type|description' (repeatable)")
	f.StringVar(&addConstraintsText, "constraints-text", "", "constraints")
	f.StringVar(&addVersioningNotes, "versioning-notes", "", "versioning notes")
	f.StringVar(&addContractWireframe, "wireframe", "", "screen contract UIML wireframe source — required when --contract-kind=screen")
	// Concept
	f.StringVar(&addMeaning, "meaning", "", "concept meaning")
	f.StringVar(&addStructureNotes, "structure-notes", "", "structure notes")
	f.StringVar(&addLifecycle, "lifecycle", "", "lifecycle")
	f.StringVar(&addInvariants, "invariants", "", "invariants")
	f.StringVar(&addDataSensitivity, "data-sensitivity", "", "data sensitivity")
	f.StringArrayVar(&addConceptAttributes, "attribute", nil, "concept attribute 'name|type|description' (repeatable) — ERD-style field")
	f.StringArrayVar(&addConceptActions, "action", nil, "concept action 'name|description' (repeatable) — domain verb on the aggregate")
	// Flow
	f.StringVar(&addTrigger, "trigger", "", "flow trigger")
	f.StringVar(&addGoal, "goal", "", "flow goal")
	f.StringVar(&addNarrative, "narrative", "", "flow narrative")
	f.StringVar(&addHappyPath, "happy-path", "", "happy path")
	f.StringVar(&addEdgeCases, "edge-cases", "", "edge cases")
	f.StringVar(&addFailureModes, "failure-modes", "", "failure modes")
	f.StringVar(&addPerfNotes, "performance-notes", "", "performance notes")
	// Decision
	f.StringVar(&addCategory, "category", "", "decision category")
	f.StringVar(&addStatement, "statement", "", "decision statement")
	f.StringVar(&addRationale, "rationale", "", "rationale")
	f.StringVar(&addAlternatives, "alternatives", "", "alternatives considered")
	f.StringVar(&addTradeoffs, "tradeoffs", "", "tradeoffs")
	f.StringVar(&addConsequences, "consequences", "", "consequences")
	f.StringVar(&addSupersedes, "supersedes", "", "slug/ref superseded by this decision or requirement")
	// Requirement
	f.StringVar(&addReqSource, "source", "", "requirement source (user/plan/migration/manual)")
	f.StringVar(&addReqSourceRef, "source-ref", "", "requirement source reference")
	f.StringVar(&addReqStatus, "requirement-status", "", "requirement status (active/superseded/obsolete)")
	f.StringVar(&addReqAcceptance, "acceptance", "", "requirement acceptance criteria")
	f.StringVar(&addReqSupersededBy, "superseded-by", "", "requirement refs that supersede this one, comma-separated")
	f.StringVar(&addReqObsoleteReason, "obsolete-reason", "", "why the requirement is obsolete")
	f.StringVar(&addReqApprovedAt, "approved-at", "", "requirement approval timestamp")
	// System
	f.StringVar(&addSysContext, "context-text", "", "system context")
	f.StringVar(&addSysScope, "scope", "", "system scope")
	f.StringVar(&addSysPrinciples, "design-principles", "", "design principles")
	f.StringVar(&addSysQuality, "quality-goals", "", "quality goals")
	f.StringVar(&addSysAssumptions, "assumptions", "", "assumptions")

	rootCmd.AddCommand(addCmd)
}
