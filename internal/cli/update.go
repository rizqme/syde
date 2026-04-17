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
	updateName             string
	updateDescription      string
	updatePurpose          string
	updateDeprecated       bool
	updateDeprecatedReason string
	updateReplacedBy       string
	updateAddTags          []string
	updateFiles            []string
	updateAddNotes         []string
	updateRemoveTags       []string
	updateBody             string
	updateAddRel           []string
	updateRemoveRel        []string
	// Component
	updCompResponsibility  string
	updCompCapabilities    []string
	updCompBoundaries      string
	updCompBehaviorSummary string
	updCompInteractionSum  string
	updCompDataHandling    string
	updCompScalingNotes    string
	// Contract
	updContKind         string
	updContPattern      string
	updContProtocol     string
	updContInput        string
	updContInputParams  []string
	updContOutput       string
	updContOutputParams []string
	updContConstraints  string
	updContVersioning   string
	updContWireframe    string
	// Concept
	updConcMeaning    string
	updConcLifecycle  string
	updConcInvariants string
	// Flow
	updFlowTrigger     string
	updFlowGoal        string
	updFlowNarrative   string
	updFlowHappyPath   string
	updFlowEdgeCases   string
	updFlowFailures    string
	updFlowPerformance string
	updFlowSteps       []string
	// Requirement (statement/rationale/supersedes)
	updDecStatement  string
	updDecRationale  string
	updDecSupersedes string
	// Requirement
	updReqType           string
	updReqPriority       string
	updReqVerification   string
	updReqSource         string
	updReqSourceRef      string
	updReqStatus         string
	updReqSupersededBy   string
	updReqObsoleteReason string
	updReqApprovedAt     string
	updReqAuditedOverlaps []string
	// System
	updSysContext     string
	updSysScope       string
	updSysPrinciples  string
	updSysQuality     string
	updSysAssumptions string
)

var updateCmd = &cobra.Command{
	Use:   "update <id-or-slug>",
	Short: "Update an existing entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]

		store, err := openWriteClient()
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
		if cmd.Flags().Changed("name") {
			b.Name = updateName
			changed = true
		}
		if cmd.Flags().Changed("description") {
			b.Description = updateDescription
			changed = true
		}
		if cmd.Flags().Changed("purpose") {
			b.Purpose = updatePurpose
			changed = true
		}
		if cmd.Flags().Changed("file") {
			b.Files = updateFiles
			changed = true
		}
		if cmd.Flags().Changed("deprecated") {
			b.Deprecated = updateDeprecated
			changed = true
		}
		if cmd.Flags().Changed("deprecated-reason") {
			b.DeprecatedReason = updateDeprecatedReason
			changed = true
		}
		if cmd.Flags().Changed("replaced-by") {
			b.ReplacedBy = updateReplacedBy
			changed = true
		}
		if cmd.Flags().Changed("body") {
			body = updateBody
			changed = true
		}
		if cmd.Flags().Changed("note") {
			b.Notes = append(b.Notes, updateAddNotes...)
			changed = true
		}

		for _, tag := range updateAddTags {
			if !hasTagSlice(b.Tags, tag) {
				b.Tags = append(b.Tags, tag)
				changed = true
			}
		}
		for _, tag := range updateRemoveTags {
			b.Tags = removeTagSlice(b.Tags, tag)
			changed = true
		}

		// Relationship management
		if cmd.Flags().Changed("add-rel") {
			for _, spec := range updateAddRel {
				if spec == "" {
					continue
				}
				rel, err := parseRelSpec(spec)
				if err != nil {
					return err
				}
				b.Relationships = append(b.Relationships, rel)
				changed = true
			}
		}
		if cmd.Flags().Changed("remove-rel") && len(updateRemoveRel) > 0 {
			removeSet := make(map[string]bool, len(updateRemoveRel))
			for _, t := range updateRemoveRel {
				if t != "" {
					removeSet[t] = true
				}
			}
			var kept []model.Relationship
			for _, r := range b.Relationships {
				if removeSet[r.Target] {
					continue
				}
				kept = append(kept, r)
			}
			b.Relationships = kept
			changed = true
		}

		// Kind-specific fields
		switch v := entity.(type) {
		case *model.ComponentEntity:
			if cmd.Flags().Changed("responsibility") {
				v.Responsibility = updCompResponsibility
				changed = true
			}
			if cmd.Flags().Changed("capability") {
				v.Capabilities = updCompCapabilities
				changed = true
			}
			if cmd.Flags().Changed("boundaries") {
				v.Boundaries = updCompBoundaries
				changed = true
			}
			if cmd.Flags().Changed("behavior-summary") {
				v.BehaviorSummary = updCompBehaviorSummary
				changed = true
			}
			if cmd.Flags().Changed("interaction-summary") {
				v.InteractionSummary = updCompInteractionSum
				changed = true
			}
			if cmd.Flags().Changed("data-handling") {
				v.DataHandling = updCompDataHandling
				changed = true
			}
			if cmd.Flags().Changed("scaling-notes") {
				v.ScalingNotes = updCompScalingNotes
				changed = true
			}
		case *model.ContractEntity:
			if cmd.Flags().Changed("contract-kind") {
				v.ContractKind = updContKind
				changed = true
			}
			if cmd.Flags().Changed("interaction-pattern") {
				v.InteractionPattern = updContPattern
				changed = true
			}
			if cmd.Flags().Changed("protocol-notes") {
				v.ProtocolNotes = updContProtocol
				changed = true
			}
			if cmd.Flags().Changed("input") {
				v.Input = updContInput
				changed = true
			}
			if cmd.Flags().Changed("input-parameter") {
				v.InputParameters = parseContractParams(updContInputParams)
				changed = true
			}
			if cmd.Flags().Changed("output") {
				v.Output = updContOutput
				changed = true
			}
			if cmd.Flags().Changed("output-parameter") {
				v.OutputParameters = parseContractParams(updContOutputParams)
				changed = true
			}
			if cmd.Flags().Changed("constraints-text") {
				v.Constraints = updContConstraints
				changed = true
			}
			if cmd.Flags().Changed("versioning-notes") {
				v.VersioningNotes = updContVersioning
				changed = true
			}
			if cmd.Flags().Changed("wireframe") {
				v.Wireframe = updContWireframe
				changed = true
			}
		case *model.ConceptEntity:
			if cmd.Flags().Changed("meaning") {
				v.Meaning = updConcMeaning
				changed = true
			}
			if cmd.Flags().Changed("lifecycle") {
				v.Lifecycle = updConcLifecycle
				changed = true
			}
			if cmd.Flags().Changed("invariants") {
				v.Invariants = updConcInvariants
				changed = true
			}
		case *model.FlowEntity:
			if cmd.Flags().Changed("trigger") {
				v.Trigger = updFlowTrigger
				changed = true
			}
			if cmd.Flags().Changed("goal") {
				v.Goal = updFlowGoal
				changed = true
			}
			if cmd.Flags().Changed("narrative") {
				v.Narrative = updFlowNarrative
				changed = true
			}
			if cmd.Flags().Changed("happy-path") {
				v.HappyPath = updFlowHappyPath
				changed = true
			}
			if cmd.Flags().Changed("edge-cases") {
				v.EdgeCases = updFlowEdgeCases
				changed = true
			}
			if cmd.Flags().Changed("failure-modes") {
				v.FlowFailureModes = updFlowFailures
				changed = true
			}
			if cmd.Flags().Changed("performance-notes") {
				v.PerformanceNotes = updFlowPerformance
				changed = true
			}
			if cmd.Flags().Changed("step") {
				v.Steps = parseFlowSteps(updFlowSteps)
				changed = true
			}
		case *model.RequirementEntity:
			if cmd.Flags().Changed("statement") {
				v.Statement = updDecStatement
				changed = true
			}
			if cmd.Flags().Changed("source") {
				v.Source = updReqSource
				changed = true
			}
			if cmd.Flags().Changed("source-ref") {
				v.SourceRef = updReqSourceRef
				changed = true
			}
			if cmd.Flags().Changed("requirement-status") {
				status, err := parseRequirementStatus(updReqStatus)
				if err != nil {
					return err
				}
				v.RequirementStatus = status
				changed = true
			}
			if cmd.Flags().Changed("rationale") {
				v.Rationale = updDecRationale
				changed = true
			}
			if cmd.Flags().Changed("type") {
				v.ReqType = model.RequirementType(updReqType)
				changed = true
			}
			if cmd.Flags().Changed("priority") {
				v.Priority = model.RequirementPriority(updReqPriority)
				changed = true
			}
			if cmd.Flags().Changed("verification") {
				v.Verification = updReqVerification
				changed = true
			}
			if cmd.Flags().Changed("supersedes") {
				v.Supersedes = parseRefList(updDecSupersedes)
				for _, oldRef := range v.Supersedes {
					oldEntity, oldBody, err := store.Get(oldRef)
					if err != nil {
						continue
					}
					oldReq, ok := oldEntity.(*model.RequirementEntity)
					if !ok {
						continue
					}
					oldReq.RequirementStatus = model.RequirementSuperseded
					oldReq.SupersededBy = appendRefOnce(oldReq.SupersededBy, b.CanonicalSlug())
					if _, err := store.Update(oldEntity, oldBody); err == nil {
						fmt.Printf("  Marked requirement superseded: %s\n", oldReq.Name)
					}
				}
				changed = true
			}
			if cmd.Flags().Changed("superseded-by") {
				v.SupersededBy = parseRefList(updReqSupersededBy)
				if len(v.SupersededBy) > 0 {
					v.RequirementStatus = model.RequirementSuperseded
				}
				changed = true
			}
			if cmd.Flags().Changed("obsolete-reason") {
				v.ObsoleteReason = updReqObsoleteReason
				if v.RequirementStatus == "" || v.RequirementStatus == model.RequirementActive {
					v.RequirementStatus = model.RequirementObsolete
				}
				changed = true
			}
			if cmd.Flags().Changed("approved-at") {
				v.ApprovedAt = updReqApprovedAt
				changed = true
			}
			if cmd.Flags().Changed("audited") {
				// Append to existing list, deduplicated by slug. Later
				// entries replace earlier ones so an author can refresh
				// the distinction text with a re-invocation.
				index := map[string]int{}
				for i, ao := range v.AuditedOverlaps {
					index[ao.Slug] = i
				}
				for _, incoming := range parseAuditedOverlaps(updReqAuditedOverlaps) {
					if existing, ok := index[incoming.Slug]; ok {
						if incoming.Distinction != "" {
							v.AuditedOverlaps[existing].Distinction = incoming.Distinction
						}
						continue
					}
					index[incoming.Slug] = len(v.AuditedOverlaps)
					v.AuditedOverlaps = append(v.AuditedOverlaps, incoming)
				}
				changed = true
			}
		case *model.SystemEntity:
			if cmd.Flags().Changed("context-text") {
				v.Context = updSysContext
				changed = true
			}
			if cmd.Flags().Changed("scope") {
				v.Scope = updSysScope
				changed = true
			}
			if cmd.Flags().Changed("design-principles") {
				v.DesignPrinciples = updSysPrinciples
				changed = true
			}
			if cmd.Flags().Changed("quality-goals") {
				v.QualityGoals = updSysQuality
				changed = true
			}
			if cmd.Flags().Changed("assumptions") {
				v.Assumptions = updSysAssumptions
				changed = true
			}
		}

		if !changed {
			return fmt.Errorf("no changes specified")
		}

		if verrs := model.ValidateEntity(entity); len(verrs) > 0 {
			var msgs []string
			for _, ve := range verrs {
				msgs = append(msgs, fmt.Sprintf("  - %s: %s", ve.Field, ve.Message))
			}
			return fmt.Errorf("validation failed:\n%s", strings.Join(msgs, "\n"))
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
	for _, t := range tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

func removeTagSlice(tags []string, tag string) []string {
	var r []string
	for _, t := range tags {
		if !strings.EqualFold(t, tag) {
			r = append(r, t)
		}
	}
	return r
}

func init() {
	_ = utils.Slugify
	f := updateCmd.Flags()
	// Base
	f.StringVar(&updateName, "name", "", "new name")
	f.StringVar(&updateDescription, "description", "", "new description")
	f.StringVar(&updatePurpose, "purpose", "", "new purpose")
	f.StringSliceVar(&updateAddTags, "add-tag", nil, "add tag")
	// Use StringArrayVar (literal-valued, repeatable) for any flag whose
	// values may legitimately contain commas — files, notes, relationships,
	// capabilities. StringSliceVar splits on commas and silently corrupts
	// the data.
	f.StringArrayVar(&updateAddNotes, "note", nil, "append informal note (repeatable)")
	f.StringArrayVar(&updateFiles, "file", nil, "concrete source file path (repeatable)")
	f.BoolVar(&updateDeprecated, "deprecated", false, "mark entity as deprecated")
	f.StringVar(&updateDeprecatedReason, "deprecated-reason", "", "reason for deprecation")
	f.StringVar(&updateReplacedBy, "replaced-by", "", "slug of replacement entity")
	f.StringSliceVar(&updateRemoveTags, "remove-tag", nil, "remove tag")
	f.StringVar(&updateBody, "body", "", "set markdown body")
	f.StringArrayVar(&updateAddRel, "add-rel", nil, "add relationship (target:type) — repeatable")
	f.StringArrayVar(&updateRemoveRel, "remove-rel", nil, "remove relationship by target — repeatable")
	// Component
	f.StringVar(&updCompResponsibility, "responsibility", "", "component responsibility")
	f.StringArrayVar(&updCompCapabilities, "capability", nil, "component capability (repeatable)")
	f.StringVar(&updCompBoundaries, "boundaries", "", "component boundaries")
	f.StringVar(&updCompBehaviorSummary, "behavior-summary", "", "component behavior summary")
	f.StringVar(&updCompInteractionSum, "interaction-summary", "", "component interaction summary")
	f.StringVar(&updCompDataHandling, "data-handling", "", "component data handling")
	f.StringVar(&updCompScalingNotes, "scaling-notes", "", "component scaling notes")
	// Contract
	f.StringVar(&updContKind, "contract-kind", "", "contract kind (api/event/command/query)")
	f.StringVar(&updContPattern, "interaction-pattern", "", "interaction pattern")
	f.StringVar(&updContProtocol, "protocol-notes", "", "protocol notes")
	f.StringVar(&updContInput, "input", "", "contract invocation signature (e.g. 'GET /api/projects')")
	f.StringArrayVar(&updContInputParams, "input-parameter", nil, "input parameter 'path|type|description' (repeatable, replaces existing)")
	f.StringVar(&updContOutput, "output", "", "contract output signature / response shape")
	f.StringArrayVar(&updContOutputParams, "output-parameter", nil, "output parameter 'path|type|description' (repeatable, replaces existing)")
	f.StringVar(&updContConstraints, "constraints-text", "", "constraints")
	f.StringVar(&updContVersioning, "versioning-notes", "", "versioning notes")
	f.StringVar(&updContWireframe, "wireframe", "", "screen contract UIML wireframe source — required when --contract-kind=screen")
	// Concept
	f.StringVar(&updConcMeaning, "meaning", "", "concept meaning (one-liner explaining what this domain term is)")
	f.StringVar(&updConcLifecycle, "lifecycle", "", "lifecycle prose for the concept")
	f.StringVar(&updConcInvariants, "invariants", "", "invariants — rules that must always hold")
	// Flow
	f.StringVar(&updFlowTrigger, "trigger", "", "flow trigger")
	f.StringVar(&updFlowGoal, "goal", "", "flow goal")
	f.StringVar(&updFlowNarrative, "narrative", "", "flow narrative")
	f.StringVar(&updFlowHappyPath, "happy-path", "", "happy path")
	f.StringVar(&updFlowEdgeCases, "edge-cases", "", "edge cases")
	f.StringVar(&updFlowFailures, "failure-modes", "", "failure modes")
	f.StringArrayVar(&updFlowSteps, "step", nil, "flow step 'id|action|contract|description|on_success|on_failure' (repeatable, replaces existing)")
	f.StringVar(&updFlowPerformance, "performance-notes", "", "performance notes")
	// Requirement
	f.StringVar(&updDecStatement, "statement", "", "requirement statement (EARS shall-form)")
	f.StringVar(&updDecRationale, "rationale", "", "requirement rationale")
	f.StringVar(&updDecSupersedes, "supersedes", "", "slug/ref superseded by this requirement")
	f.StringVar(&updReqSource, "source", "", "requirement source (user/plan/migration/manual)")
	f.StringVar(&updReqSourceRef, "source-ref", "", "requirement source reference")
	f.StringVar(&updReqStatus, "requirement-status", "", "requirement status (active/superseded/obsolete)")
	f.StringVar(&updReqType, "type", "", "requirement type (functional/non-functional/constraint/interface/performance/security/usability)")
	f.StringVar(&updReqPriority, "priority", "", "requirement priority (must/should/could/wont)")
	f.StringVar(&updReqVerification, "verification", "", "how the requirement is verified")
	f.StringVar(&updReqSupersededBy, "superseded-by", "", "requirement refs that supersede this one, comma-separated")
	f.StringVar(&updReqObsoleteReason, "obsolete-reason", "", "why the requirement is obsolete")
	f.StringVar(&updReqApprovedAt, "approved-at", "", "requirement approval timestamp")
	f.StringArrayVar(&updReqAuditedOverlaps, "audited", nil, "acknowledged overlapping requirement as slug[:distinction rationale] (repeatable, appends to existing; re-passing same slug replaces distinction)")
	// System
	f.StringVar(&updSysContext, "context-text", "", "system context")
	f.StringVar(&updSysScope, "scope", "", "system scope")
	f.StringVar(&updSysPrinciples, "design-principles", "", "design principles")
	f.StringVar(&updSysQuality, "quality-goals", "", "quality goals")
	f.StringVar(&updSysAssumptions, "assumptions", "", "assumptions")

	rootCmd.AddCommand(updateCmd)
}
