package skill

import skillembed "github.com/feedloop/syde/skill"

// SkillMD is the SKILL.md content for the Claude Code skill.
var SkillMD = mustRead("SKILL.md")

// CodexSkillMD is the SKILL.md content for the Codex skill.
var CodexSkillMD = mustRead("codex/SKILL.md")

// CodexHooksJSON is the hooks template for Codex.
var CodexHooksJSON = mustRead("codex/hooks.json")

// EntitySpecRef is the entity specification reference content.
var EntitySpecRef = mustRead("references/entity-spec.md")

// CommandsRef is the CLI command reference.
var CommandsRef = mustRead("references/commands.md")

// ClarifyGuideRef is the critical requirement gathering guide.
var ClarifyGuideRef = mustRead("references/clarify-guide.md")

// SyncWorkflowRef is the sync workflow reference.
var SyncWorkflowRef = mustRead("references/sync-workflow.md")

// RequirementDerivationRef is the deterministic procedure for
// backfilling EARS requirements from existing entities. Subagents
// in phase-4 backfill use this as their algorithm spec.
var RequirementDerivationRef = mustRead("references/requirement-derivation.md")

func mustRead(path string) string {
	data, err := skillembed.FS.ReadFile(path)
	if err != nil {
		panic("skill: embedded file missing: " + path)
	}
	return string(data)
}
