package skill

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Install installs the syde skill and hooks into a project's .claude/ directory.
func Install(projectRoot string) error {
	return InstallClaude(projectRoot)
}

// InstallClaude installs the syde skill and hooks into a project's .claude/ directory.
func InstallClaude(projectRoot string) error {
	claudeDir := filepath.Join(projectRoot, ".claude")
	skillDir := filepath.Join(claudeDir, "skills", "syde")
	refsDir := filepath.Join(skillDir, "references")

	// Create directories
	if err := os.MkdirAll(refsDir, 0755); err != nil {
		return fmt.Errorf("create skill dir: %w", err)
	}

	// Write SKILL.md
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte(SkillMD), 0644); err != nil {
		return err
	}

	// Write reference files
	refs := map[string]string{
		"entity-spec.md":             EntitySpecRef,
		"commands.md":                CommandsRef,
		"clarify-guide.md":           ClarifyGuideRef,
		"sync-workflow.md":           SyncWorkflowRef,
		"requirement-derivation.md":  RequirementDerivationRef,
	}
	for name, content := range refs {
		if err := os.WriteFile(filepath.Join(refsDir, name), []byte(content), 0644); err != nil {
			return err
		}
	}

	// Clean up old reference files that were merged into SKILL.md
	for _, old := range []string{"plan-workflow.md", "constraints.md", "scan-workflow.md"} {
		os.Remove(filepath.Join(refsDir, old))
	}

	// Write hooks template
	hooksDir := filepath.Join(claudeDir, "hooks")
	os.MkdirAll(hooksDir, 0755)
	if err := os.WriteFile(filepath.Join(hooksDir, "syde-hooks.json"), []byte(HooksJSON), 0644); err != nil {
		return err
	}

	// Update CLAUDE.md
	if err := updateClaudeMD(projectRoot); err != nil {
		return err
	}

	return nil
}

// InstallCodex installs the syde skill and hooks into Codex's repo-local
// discovery locations.
func InstallCodex(projectRoot string) error {
	agentsDir := filepath.Join(projectRoot, ".agents", "skills", "syde")
	refsDir := filepath.Join(agentsDir, "references")
	if err := os.MkdirAll(refsDir, 0755); err != nil {
		return fmt.Errorf("create Codex skill dir: %w", err)
	}

	if err := os.WriteFile(filepath.Join(agentsDir, "SKILL.md"), []byte(CodexSkillMD), 0644); err != nil {
		return err
	}

	refs := map[string]string{
		"entity-spec.md":             EntitySpecRef,
		"commands.md":                CommandsRef,
		"clarify-guide.md":           ClarifyGuideRef,
		"sync-workflow.md":           SyncWorkflowRef,
		"requirement-derivation.md":  RequirementDerivationRef,
	}
	for name, content := range refs {
		if err := os.WriteFile(filepath.Join(refsDir, name), []byte(content), 0644); err != nil {
			return err
		}
	}

	codexDir := filepath.Join(projectRoot, ".codex")
	if err := os.MkdirAll(codexDir, 0755); err != nil {
		return fmt.Errorf("create .codex dir: %w", err)
	}
	if err := os.WriteFile(filepath.Join(codexDir, "hooks.json"), []byte(CodexHooksJSON), 0644); err != nil {
		return err
	}
	if err := enableCodexHooksFeature(filepath.Join(codexDir, "config.toml")); err != nil {
		return err
	}

	if err := updateAgentsMD(projectRoot); err != nil {
		return err
	}

	return nil
}

func updateClaudeMD(projectRoot string) error {
	claudeMDPath := filepath.Join(projectRoot, "CLAUDE.md")

	content := ""
	if data, err := os.ReadFile(claudeMDPath); err == nil {
		content = string(data)
	}

	// Check if syde section already exists
	marker := "## syde Design Model"
	if strings.Contains(content, marker) {
		// Replace existing section
		startIdx := strings.Index(content, marker)
		endIdx := len(content)
		// Find next ## heading
		if nextH2 := strings.Index(content[startIdx+len(marker):], "\n## "); nextH2 >= 0 {
			endIdx = startIdx + len(marker) + nextH2
		}
		content = content[:startIdx] + claudeMDSection + content[endIdx:]
	} else {
		if content != "" && !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		content += "\n" + claudeMDSection
	}

	return os.WriteFile(claudeMDPath, []byte(content), 0644)
}

func updateAgentsMD(projectRoot string) error {
	agentsMDPath := filepath.Join(projectRoot, "AGENTS.md")

	content := ""
	if data, err := os.ReadFile(agentsMDPath); err == nil {
		content = string(data)
	}

	marker := "## syde Design Model"
	if strings.Contains(content, marker) {
		startIdx := strings.Index(content, marker)
		endIdx := len(content)
		if nextH2 := strings.Index(content[startIdx+len(marker):], "\n## "); nextH2 >= 0 {
			endIdx = startIdx + len(marker) + nextH2
		}
		content = content[:startIdx] + agentsMDSection + content[endIdx:]
	} else {
		if content != "" && !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		content += "\n" + agentsMDSection
	}

	return os.WriteFile(agentsMDPath, []byte(content), 0644)
}

func enableCodexHooksFeature(path string) error {
	content := ""
	if data, err := os.ReadFile(path); err == nil {
		content = string(data)
	}
	if strings.Contains(content, "codex_hooks") {
		return nil
	}
	if strings.Contains(content, "[features]") {
		idx := strings.Index(content, "[features]") + len("[features]")
		content = content[:idx] + "\ncodex_hooks = true" + content[idx:]
	} else {
		if content != "" && !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		content += "\n[features]\ncodex_hooks = true\n"
	}
	return os.WriteFile(path, []byte(content), 0644)
}

const claudeMDSection = `## syde Design Model

This project uses syde for architecture management. These rules are mandatory:

1. **Phase 0 — always run first**: ` + "`syde tree scan`" + `, then iterate ` + "`syde tree changes --leaves-only`" + ` + ` + "`syde tree summarize <path> --summary \"...\"`" + ` until ` + "`syde tree status --strict`" + ` exits 0. The summary tree is the cheap way to understand the project without re-reading every file. ` + "`.gitignore`" + ` is honored automatically.
2. **Architecture auto-loaded**: ` + "`syde context`" + ` runs at session start. Do NOT re-run it. Use ` + "`syde query <slug> --full`" + ` for targeted deep dives only.
3. **Use ` + "`syde tree context <path>`" + `, NOT naive ` + "`Read`" + `, when creating entities on an existing codebase.** It returns the ancestor breadcrumb + file summary + content in one call — the right framing for ` + "`--purpose`" + ` / ` + "`--responsibility`" + ` / ` + "`--boundaries`" + ` and for picking ` + "`belongs_to`" + `.
4. **Clarify first**: Be critical — challenge assumptions, identify missing requirements, propose constraints. Use the project-type checklists in the syde skill. Wait for user confirmation before proceeding.
5. **Design before code**: Create a plan with ` + "`syde plan create`" + `, add entity drafts with ` + "`syde plan add-entity`" + `, add phases with ` + "`syde plan add-phase`" + `. Present to user. Do NOT implement until approved.
6. **Track implementation**: Use ` + "`syde task create`" + ` / ` + "`syde task start`" + ` / ` + "`syde task done`" + ` for each unit of work.
7. **Verify after writing files**: Run ` + "`syde constraints check <file>`" + ` to verify new files are mapped to components.
8. **Finish**: Run ` + "`syde validate`" + `, then **refresh the summary tree**: ` + "`syde tree scan`" + ` + leaves-first summarize loop until ` + "`syde tree status --strict`" + ` exits 0. The Stop hook will block the session from ending cleanly if the tree is dirty.
9. **Never read .syde/ files directly** — always use syde CLI commands.
`

const agentsMDSection = `## syde Design Model

This project uses syde for architecture management. These Codex rules are mandatory:

1. **Start with syde context**: Run ` + "`syde tree scan`" + ` and ensure ` + "`syde tree status --strict`" + ` passes before planning or editing.
2. **Use syde query before raw reads**: Prefer ` + "`syde query --file <path> --content`" + `, ` + "`syde query --code <symbol>`" + `, and ` + "`syde query --search \"<term>\"`" + ` for tracked files.
3. **Design before code**: Create a syde plan, add phases/entities, present it, wait for approval, then run ` + "`syde plan approve <slug>`" + `.
4. **Track implementation**: Start a task with ` + "`syde task start <slug>`" + ` before changing files and finish with ` + "`syde task done <slug> --affected-entity <entity> --affected-file <path>`" + `.
5. **Verify mappings**: After source edits, run ` + "`syde constraints check <file>`" + ` and map new files with ` + "`syde update <component> --file <path>`" + `.
6. **Finish cleanly**: Run tests, ` + "`syde sync check --strict`" + `, and refresh stale summary-tree nodes before final response.
7. **Hook limitation**: Codex hooks currently intercept Bash only. They are guardrails, not a complete enforcement boundary for ` + "`apply_patch`" + ` or other non-Bash tools.
8. **Never read ` + "`.syde/`" + ` files directly** — always use syde CLI commands.
`
