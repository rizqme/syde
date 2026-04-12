package skill

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Install installs the syde skill and hooks into a project's .claude/ directory.
func Install(projectRoot string) error {
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
		"entity-spec.md":    EntitySpecRef,
		"commands.md":       CommandsRef,
		"plan-workflow.md":  PlanWorkflowRef,
		"constraints.md":    ConstraintsRef,
		"scan-workflow.md":  ScanWorkflowRef,
	}
	for name, content := range refs {
		if err := os.WriteFile(filepath.Join(refsDir, name), []byte(content), 0644); err != nil {
			return err
		}
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

const claudeMDSection = `## syde Design Model

This project uses syde for architecture management. These rules are mandatory:

1. **Architecture auto-loaded**: ` + "`syde context`" + ` runs at session start. Use ` + "`syde query <slug> --full`" + ` for targeted deep dives only.
2. **Clarify first**: Before implementing, ask ALL clarifying questions with recommendations. Wait for user confirmation.
3. **Plan before code**: Create plan with ` + "`syde plan create`" + ` + ` + "`syde plan add-step`" + `. Present to user. Do NOT implement until approved.
4. **After writing files**: Run ` + "`syde constraints check <file>`" + ` to verify new files are mapped to components.
5. **After completing work**: Run ` + "`syde validate`" + `. Capture learnings with ` + "`syde remember`" + `.
6. **Never read .syde/ files directly** — always use syde CLI commands.
`
