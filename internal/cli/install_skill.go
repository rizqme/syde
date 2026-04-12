package cli

import (
	"fmt"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/skill"
	"github.com/spf13/cobra"
)

var installSkillCmd = &cobra.Command{
	Use:   "install-skill",
	Short: "Install Claude Code skill and hooks",
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
		if err := skill.Install(projectRoot); err != nil {
			return fmt.Errorf("install skill: %w", err)
		}

		fmt.Println("Installed syde skill:")
		fmt.Println("  .claude/skills/syde/SKILL.md")
		fmt.Println("  .claude/skills/syde/references/*.md")
		fmt.Println("  CLAUDE.md updated with syde constraints")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installSkillCmd)
}
