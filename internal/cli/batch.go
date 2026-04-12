package cli

import (
	"fmt"
	"os"

	"github.com/feedloop/syde/internal/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type batchFile struct {
	Entities []batchEntity `yaml:"entities"`
}

type batchEntity struct {
	Kind               string   `yaml:"kind"`
	Name               string   `yaml:"name"`
	Description        string   `yaml:"description,omitempty"`
	Purpose            string   `yaml:"purpose,omitempty"`
	Status             string   `yaml:"status,omitempty"`
	Tags               []string `yaml:"tags,omitempty"`
	Body               string   `yaml:"body,omitempty"`
	// Component
	Responsibility     string   `yaml:"responsibility,omitempty"`
	Boundaries         string   `yaml:"boundaries,omitempty"`
	BehaviorSummary    string   `yaml:"behavior_summary,omitempty"`
	InteractionSummary string   `yaml:"interaction_summary,omitempty"`
	DataHandling       string   `yaml:"data_handling,omitempty"`
	ScalingNotes       string   `yaml:"scaling_notes,omitempty"`
	// Contract
	ContractKind       string   `yaml:"contract_kind,omitempty"`
	InteractionPattern string   `yaml:"interaction_pattern,omitempty"`
	ProtocolNotes      string   `yaml:"protocol_notes,omitempty"`
	InputDesc          string   `yaml:"input_description,omitempty"`
	OutputDesc         string   `yaml:"output_description,omitempty"`
	Constraints        string   `yaml:"constraints,omitempty"`
	// Concept
	Meaning            string   `yaml:"meaning,omitempty"`
	StructureNotes     string   `yaml:"structure_notes,omitempty"`
	Lifecycle          string   `yaml:"lifecycle,omitempty"`
	Invariants         string   `yaml:"invariants,omitempty"`
	// Flow
	Trigger            string   `yaml:"trigger,omitempty"`
	Goal               string   `yaml:"goal,omitempty"`
	Narrative          string   `yaml:"narrative,omitempty"`
	// Decision
	Category           string   `yaml:"category,omitempty"`
	Statement          string   `yaml:"statement,omitempty"`
	Rationale          string   `yaml:"rationale,omitempty"`
	Alternatives       string   `yaml:"alternatives_considered,omitempty"`
	Tradeoffs          string   `yaml:"tradeoffs,omitempty"`
	// System
	Context            string   `yaml:"context,omitempty"`
	Scope              string   `yaml:"scope,omitempty"`
	DesignPrinciples   string   `yaml:"design_principles,omitempty"`
	QualityGoals       string   `yaml:"quality_goals,omitempty"`
}

var batchCmd = &cobra.Command{
	Use:   "batch <file.yaml>",
	Short: "Bulk create entities from a YAML file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("read batch file: %w", err)
		}

		var bf batchFile
		if err := yaml.Unmarshal(data, &bf); err != nil {
			return fmt.Errorf("parse batch file: %w", err)
		}

		if len(bf.Entities) == 0 {
			return fmt.Errorf("no entities in batch file")
		}

		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		created := 0
		for _, be := range bf.Entities {
			kind, ok := model.ParseEntityKind(be.Kind)
			if !ok {
				fmt.Printf("  SKIP: unknown kind '%s' for '%s'\n", be.Kind, be.Name)
				continue
			}

			entity := model.NewEntityForKind(kind)
			b := entity.GetBase()
			b.Name = be.Name
			b.Description = be.Description
			b.Purpose = be.Purpose
			b.Tags = be.Tags
			if be.Status != "" {
				b.Status = model.Status(be.Status)
			}

			switch v := entity.(type) {
			case *model.ComponentEntity:
				v.Responsibility = be.Responsibility
				v.Boundaries = be.Boundaries
				v.BehaviorSummary = be.BehaviorSummary
				v.InteractionSummary = be.InteractionSummary
				v.DataHandling = be.DataHandling
				v.ScalingNotes = be.ScalingNotes
			case *model.ContractEntity:
				v.ContractKind = be.ContractKind
				v.InteractionPattern = be.InteractionPattern
				v.ProtocolNotes = be.ProtocolNotes
				v.InputDescription = be.InputDesc
				v.OutputDescription = be.OutputDesc
				v.Constraints = be.Constraints
			case *model.ConceptEntity:
				v.Meaning = be.Meaning
				v.StructureNotes = be.StructureNotes
				v.Lifecycle = be.Lifecycle
				v.Invariants = be.Invariants
			case *model.FlowEntity:
				v.Trigger = be.Trigger
				v.Goal = be.Goal
				v.Narrative = be.Narrative
			case *model.DecisionEntity:
				v.Category = be.Category
				v.Statement = be.Statement
				v.Rationale = be.Rationale
				v.AlternativesConsidered = be.Alternatives
				v.Tradeoffs = be.Tradeoffs
			case *model.SystemEntity:
				v.Context = be.Context
				v.Scope = be.Scope
				v.DesignPrinciples = be.DesignPrinciples
				v.QualityGoals = be.QualityGoals
			}

			filePath, err := store.Create(entity, be.Body)
			if err != nil {
				fmt.Printf("  ERROR: %s: %v\n", be.Name, err)
				continue
			}

			fmt.Printf("  ✓ %s: %s → %s\n", kind, be.Name, filePath)
			created++
		}

		fmt.Printf("\nCreated %d/%d entities.\n", created, len(bf.Entities))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)
}
