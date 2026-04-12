package cli

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/utils"
	"github.com/spf13/cobra"
)

var (
	// Base
	updateName        string
	updateDescription string
	updatePurpose     string
	updateStatus      string
	updateAddTags     []string
	updateRemoveTags  []string
	updateBody        string
	updateAddRel      string
	updateRemoveRel   string
	// Component
	updCompResponsibility  string
	updCompBoundaries      string
	updCompBehaviorSummary string
	updCompInteractionSum  string
	updCompDataHandling    string
	updCompScalingNotes    string
	// Contract
	updContKind       string
	updContPattern    string
	updContProtocol   string
	updContInputDesc  string
	updContOutputDesc string
	updContConstraints string
	updContVersioning string
	// Concept
	updConcMeaning       string
	updConcStructure     string
	updConcLifecycle     string
	updConcInvariants    string
	updConcSensitivity   string
	// Flow
	updFlowTrigger     string
	updFlowGoal        string
	updFlowNarrative   string
	updFlowHappyPath   string
	updFlowEdgeCases   string
	updFlowFailures    string
	updFlowPerformance string
	// Decision
	updDecCategory     string
	updDecStatement    string
	updDecRationale    string
	updDecAlternatives string
	updDecTradeoffs    string
	updDecConsequences string
	// System
	updSysContext    string
	updSysScope     string
	updSysPrinciples string
	updSysQuality   string
	updSysAssumptions string
)

var updateCmd = &cobra.Command{
	Use:   "update <id-or-slug>",
	Short: "Update an existing entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, body, err := store.Get(slug)
		if err != nil {
			return fmt.Errorf("entity not found: %s", slug)
		}

		b := entity.GetBase()
		changed := false

		// Base fields
		if cmd.Flags().Changed("name") { b.Name = updateName; changed = true }
		if cmd.Flags().Changed("description") { b.Description = updateDescription; changed = true }
		if cmd.Flags().Changed("purpose") { b.Purpose = updatePurpose; changed = true }
		if cmd.Flags().Changed("status") { b.Status = model.Status(updateStatus); changed = true }
		if cmd.Flags().Changed("body") { body = updateBody; changed = true }

		for _, tag := range updateAddTags {
			if !hasTagSlice(b.Tags, tag) { b.Tags = append(b.Tags, tag); changed = true }
		}
		for _, tag := range updateRemoveTags {
			b.Tags = removeTagSlice(b.Tags, tag); changed = true
		}

		// Relationship management
		if cmd.Flags().Changed("add-rel") && updateAddRel != "" {
			parts := strings.SplitN(updateAddRel, ":", 2)
			if len(parts) == 2 {
				b.Relationships = append(b.Relationships, model.Relationship{Target: parts[0], Type: parts[1]})
				changed = true
			}
		}
		if cmd.Flags().Changed("remove-rel") && updateRemoveRel != "" {
			var kept []model.Relationship
			for _, r := range b.Relationships {
				if r.Target != updateRemoveRel {
					kept = append(kept, r)
				}
			}
			b.Relationships = kept
			changed = true
		}

		// Kind-specific fields
		switch v := entity.(type) {
		case *model.ComponentEntity:
			if cmd.Flags().Changed("responsibility") { v.Responsibility = updCompResponsibility; changed = true }
			if cmd.Flags().Changed("boundaries") { v.Boundaries = updCompBoundaries; changed = true }
			if cmd.Flags().Changed("behavior-summary") { v.BehaviorSummary = updCompBehaviorSummary; changed = true }
			if cmd.Flags().Changed("interaction-summary") { v.InteractionSummary = updCompInteractionSum; changed = true }
			if cmd.Flags().Changed("data-handling") { v.DataHandling = updCompDataHandling; changed = true }
			if cmd.Flags().Changed("scaling-notes") { v.ScalingNotes = updCompScalingNotes; changed = true }
		case *model.ContractEntity:
			if cmd.Flags().Changed("contract-kind") { v.ContractKind = updContKind; changed = true }
			if cmd.Flags().Changed("interaction-pattern") { v.InteractionPattern = updContPattern; changed = true }
			if cmd.Flags().Changed("protocol-notes") { v.ProtocolNotes = updContProtocol; changed = true }
			if cmd.Flags().Changed("input-desc") { v.InputDescription = updContInputDesc; changed = true }
			if cmd.Flags().Changed("output-desc") { v.OutputDescription = updContOutputDesc; changed = true }
			if cmd.Flags().Changed("constraints-text") { v.Constraints = updContConstraints; changed = true }
			if cmd.Flags().Changed("versioning-notes") { v.VersioningNotes = updContVersioning; changed = true }
		case *model.ConceptEntity:
			if cmd.Flags().Changed("meaning") { v.Meaning = updConcMeaning; changed = true }
			if cmd.Flags().Changed("structure-notes") { v.StructureNotes = updConcStructure; changed = true }
			if cmd.Flags().Changed("lifecycle") { v.Lifecycle = updConcLifecycle; changed = true }
			if cmd.Flags().Changed("invariants") { v.Invariants = updConcInvariants; changed = true }
			if cmd.Flags().Changed("data-sensitivity") { v.DataSensitivity = updConcSensitivity; changed = true }
		case *model.FlowEntity:
			if cmd.Flags().Changed("trigger") { v.Trigger = updFlowTrigger; changed = true }
			if cmd.Flags().Changed("goal") { v.Goal = updFlowGoal; changed = true }
			if cmd.Flags().Changed("narrative") { v.Narrative = updFlowNarrative; changed = true }
			if cmd.Flags().Changed("happy-path") { v.HappyPath = updFlowHappyPath; changed = true }
			if cmd.Flags().Changed("edge-cases") { v.EdgeCases = updFlowEdgeCases; changed = true }
			if cmd.Flags().Changed("failure-modes") { v.FlowFailureModes = updFlowFailures; changed = true }
			if cmd.Flags().Changed("performance-notes") { v.PerformanceNotes = updFlowPerformance; changed = true }
		case *model.DecisionEntity:
			if cmd.Flags().Changed("category") { v.Category = updDecCategory; changed = true }
			if cmd.Flags().Changed("statement") { v.Statement = updDecStatement; changed = true }
			if cmd.Flags().Changed("rationale") { v.Rationale = updDecRationale; changed = true }
			if cmd.Flags().Changed("alternatives") { v.AlternativesConsidered = updDecAlternatives; changed = true }
			if cmd.Flags().Changed("tradeoffs") { v.Tradeoffs = updDecTradeoffs; changed = true }
			if cmd.Flags().Changed("consequences") { v.Consequences = updDecConsequences; changed = true }
		case *model.SystemEntity:
			if cmd.Flags().Changed("context-text") { v.Context = updSysContext; changed = true }
			if cmd.Flags().Changed("scope") { v.Scope = updSysScope; changed = true }
			if cmd.Flags().Changed("design-principles") { v.DesignPrinciples = updSysPrinciples; changed = true }
			if cmd.Flags().Changed("quality-goals") { v.QualityGoals = updSysQuality; changed = true }
			if cmd.Flags().Changed("assumptions") { v.Assumptions = updSysAssumptions; changed = true }
		}

		if !changed {
			return fmt.Errorf("no changes specified")
		}

		filePath, err := store.Update(entity, body)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}

		fmt.Printf("Updated %s: %s\n", b.Kind, b.Name)
		fmt.Printf("  File: %s\n", filePath)
		return nil
	},
}

func hasTagSlice(tags []string, tag string) bool {
	for _, t := range tags { if strings.EqualFold(t, tag) { return true } }
	return false
}

func removeTagSlice(tags []string, tag string) []string {
	var r []string
	for _, t := range tags { if !strings.EqualFold(t, tag) { r = append(r, t) } }
	return r
}

func init() {
	_ = utils.Slugify
	f := updateCmd.Flags()
	// Base
	f.StringVar(&updateName, "name", "", "new name")
	f.StringVar(&updateDescription, "description", "", "new description")
	f.StringVar(&updatePurpose, "purpose", "", "new purpose")
	f.StringVar(&updateStatus, "status", "", "new status")
	f.StringSliceVar(&updateAddTags, "add-tag", nil, "add tag")
	f.StringSliceVar(&updateRemoveTags, "remove-tag", nil, "remove tag")
	f.StringVar(&updateBody, "body", "", "set markdown body")
	f.StringVar(&updateAddRel, "add-rel", "", "add relationship (target:type)")
	f.StringVar(&updateRemoveRel, "remove-rel", "", "remove relationship by target")
	// Component
	f.StringVar(&updCompResponsibility, "responsibility", "", "component responsibility")
	f.StringVar(&updCompBoundaries, "boundaries", "", "component boundaries")
	f.StringVar(&updCompBehaviorSummary, "behavior-summary", "", "component behavior summary")
	f.StringVar(&updCompInteractionSum, "interaction-summary", "", "component interaction summary")
	f.StringVar(&updCompDataHandling, "data-handling", "", "component data handling")
	f.StringVar(&updCompScalingNotes, "scaling-notes", "", "component scaling notes")
	// Contract
	f.StringVar(&updContKind, "contract-kind", "", "contract kind (api/event/command/query)")
	f.StringVar(&updContPattern, "interaction-pattern", "", "interaction pattern")
	f.StringVar(&updContProtocol, "protocol-notes", "", "protocol notes")
	f.StringVar(&updContInputDesc, "input-desc", "", "input description")
	f.StringVar(&updContOutputDesc, "output-desc", "", "output description")
	f.StringVar(&updContConstraints, "constraints-text", "", "constraints")
	f.StringVar(&updContVersioning, "versioning-notes", "", "versioning notes")
	// Concept
	f.StringVar(&updConcMeaning, "meaning", "", "concept meaning")
	f.StringVar(&updConcStructure, "structure-notes", "", "structure notes")
	f.StringVar(&updConcLifecycle, "lifecycle", "", "lifecycle")
	f.StringVar(&updConcInvariants, "invariants", "", "invariants")
	f.StringVar(&updConcSensitivity, "data-sensitivity", "", "data sensitivity")
	// Flow
	f.StringVar(&updFlowTrigger, "trigger", "", "flow trigger")
	f.StringVar(&updFlowGoal, "goal", "", "flow goal")
	f.StringVar(&updFlowNarrative, "narrative", "", "flow narrative")
	f.StringVar(&updFlowHappyPath, "happy-path", "", "happy path")
	f.StringVar(&updFlowEdgeCases, "edge-cases", "", "edge cases")
	f.StringVar(&updFlowFailures, "failure-modes", "", "failure modes")
	f.StringVar(&updFlowPerformance, "performance-notes", "", "performance notes")
	// Decision
	f.StringVar(&updDecCategory, "category", "", "decision category")
	f.StringVar(&updDecStatement, "statement", "", "decision statement")
	f.StringVar(&updDecRationale, "rationale", "", "rationale")
	f.StringVar(&updDecAlternatives, "alternatives", "", "alternatives considered")
	f.StringVar(&updDecTradeoffs, "tradeoffs", "", "tradeoffs")
	f.StringVar(&updDecConsequences, "consequences", "", "consequences")
	// System
	f.StringVar(&updSysContext, "context-text", "", "system context")
	f.StringVar(&updSysScope, "scope", "", "system scope")
	f.StringVar(&updSysPrinciples, "design-principles", "", "design principles")
	f.StringVar(&updSysQuality, "quality-goals", "", "quality goals")
	f.StringVar(&updSysAssumptions, "assumptions", "", "assumptions")

	rootCmd.AddCommand(updateCmd)
}
