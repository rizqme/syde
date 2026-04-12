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
	addStatus      string
	addTags        []string
	addBody        string
	// Component
	addResponsibility  string
	addBoundaries      string
	addBehaviorSummary string
	addInteractionSum  string
	addDataHandling    string
	addScalingNotes    string
	// Contract
	addContractKind       string
	addInteractionPattern string
	addProtocolNotes      string
	addInputDesc          string
	addOutputDesc         string
	addConstraintsText    string
	addVersioningNotes    string
	// Concept
	addMeaning       string
	addStructureNotes string
	addLifecycle     string
	addInvariants    string
	addDataSensitivity string
	// Flow
	addTrigger     string
	addGoal        string
	addNarrative   string
	addHappyPath   string
	addEdgeCases   string
	addFailureModes string
	addPerfNotes   string
	// Decision
	addCategory     string
	addStatement    string
	addRationale    string
	addAlternatives string
	addTradeoffs    string
	addConsequences string
	// System
	addSysContext    string
	addSysScope     string
	addSysPrinciples string
	addSysQuality   string
	addSysAssumptions string
)

var addCmd = &cobra.Command{
	Use:   "add <kind> <name>",
	Short: "Add a new entity to the design model",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		kindStr := strings.ToLower(args[0])
		name := args[1]

		kind, ok := model.ParseEntityKind(kindStr)
		if !ok {
			return fmt.Errorf("unknown entity kind: %s\nValid: system, component, contract, concept, flow, decision, plan, task, design, learning", kindStr)
		}

		store, err := openStore()
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
		if addStatus != "" {
			b.Status = model.Status(addStatus)
		}

		switch v := entity.(type) {
		case *model.ComponentEntity:
			v.Responsibility = addResponsibility
			v.Boundaries = addBoundaries
			v.BehaviorSummary = addBehaviorSummary
			v.InteractionSummary = addInteractionSum
			v.DataHandling = addDataHandling
			v.ScalingNotes = addScalingNotes
		case *model.ContractEntity:
			v.ContractKind = addContractKind
			v.InteractionPattern = addInteractionPattern
			v.ProtocolNotes = addProtocolNotes
			v.InputDescription = addInputDesc
			v.OutputDescription = addOutputDesc
			v.Constraints = addConstraintsText
			v.VersioningNotes = addVersioningNotes
		case *model.ConceptEntity:
			v.Meaning = addMeaning
			v.StructureNotes = addStructureNotes
			v.Lifecycle = addLifecycle
			v.Invariants = addInvariants
			v.DataSensitivity = addDataSensitivity
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
	f.StringVar(&addStatus, "status", "", "entity status")
	f.StringSliceVar(&addTags, "tag", nil, "entity tags")
	f.StringVar(&addBody, "body", "", "markdown body content")
	// Component
	f.StringVar(&addResponsibility, "responsibility", "", "component responsibility")
	f.StringVar(&addBoundaries, "boundaries", "", "component boundaries")
	f.StringVar(&addBehaviorSummary, "behavior-summary", "", "behavior summary")
	f.StringVar(&addInteractionSum, "interaction-summary", "", "interaction summary")
	f.StringVar(&addDataHandling, "data-handling", "", "data handling")
	f.StringVar(&addScalingNotes, "scaling-notes", "", "scaling notes")
	// Contract
	f.StringVar(&addContractKind, "contract-kind", "", "contract kind (api/event/command/query)")
	f.StringVar(&addInteractionPattern, "interaction-pattern", "", "interaction pattern")
	f.StringVar(&addProtocolNotes, "protocol-notes", "", "protocol notes")
	f.StringVar(&addInputDesc, "input-desc", "", "input description")
	f.StringVar(&addOutputDesc, "output-desc", "", "output description")
	f.StringVar(&addConstraintsText, "constraints-text", "", "constraints")
	f.StringVar(&addVersioningNotes, "versioning-notes", "", "versioning notes")
	// Concept
	f.StringVar(&addMeaning, "meaning", "", "concept meaning")
	f.StringVar(&addStructureNotes, "structure-notes", "", "structure notes")
	f.StringVar(&addLifecycle, "lifecycle", "", "lifecycle")
	f.StringVar(&addInvariants, "invariants", "", "invariants")
	f.StringVar(&addDataSensitivity, "data-sensitivity", "", "data sensitivity")
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
	// System
	f.StringVar(&addSysContext, "context-text", "", "system context")
	f.StringVar(&addSysScope, "scope", "", "system scope")
	f.StringVar(&addSysPrinciples, "design-principles", "", "design principles")
	f.StringVar(&addSysQuality, "quality-goals", "", "quality goals")
	f.StringVar(&addSysAssumptions, "assumptions", "", "assumptions")

	rootCmd.AddCommand(addCmd)
}
