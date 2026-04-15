package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/skill"
	"github.com/spf13/cobra"
)

var (
	installSkillClaude bool
	installSkillCodex  bool
	installSkillAll    bool
)

var installSkillCmd = &cobra.Command{
	Use:   "install-skill",
	Short: "Install syde agent skills and hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := sydeDir
		if dir == "" {
			var err error
			dir, err = config.FindSydeDir()
			if err != nil {
				return fmt.Errorf("no .syde/ directory found")
			}
		}

		projectRoot := filepath.Dir(dir)

		installClaude := installSkillClaude || installSkillAll
		installCodex := installSkillCodex || installSkillAll
		if !installClaude && !installCodex {
			installClaude = true
		}

		var installed []string
		if installClaude {
			if err := skill.InstallClaude(projectRoot); err != nil {
				return fmt.Errorf("install Claude skill: %w", err)
			}
			installed = append(installed, "Claude Code")
		}
		if installCodex {
			if err := skill.InstallCodex(projectRoot); err != nil {
				return fmt.Errorf("install Codex skill: %w", err)
			}
			installed = append(installed, "Codex")
		}

		fmt.Printf("Installed syde support for %s:\n", strings.Join(installed, " + "))
		if installClaude {
			fmt.Println("  .claude/skills/syde/SKILL.md")
			fmt.Println("  .claude/skills/syde/references/*.md")
			fmt.Println("  .claude/hooks/syde-hooks.json")
			fmt.Println("  CLAUDE.md")
		}
		if installCodex {
			fmt.Println("  .agents/skills/syde/SKILL.md")
			fmt.Println("  .agents/skills/syde/references/*.md")
			fmt.Println("  .codex/hooks.json")
			fmt.Println("  .codex/config.toml")
			fmt.Println("  AGENTS.md")
		}
		return nil
	},
}

func init() {
	installSkillCmd.Flags().BoolVar(&installSkillClaude, "claude", false, "install Claude Code skill and hooks")
	installSkillCmd.Flags().BoolVar(&installSkillCodex, "codex", false, "install Codex skill and hooks")
	installSkillCmd.Flags().BoolVar(&installSkillAll, "all", false, "install support for Claude Code and Codex")
	rootCmd.AddCommand(installSkillCmd)
}
