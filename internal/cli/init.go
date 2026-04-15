package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/dashboard"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/skill"
	"github.com/spf13/cobra"
)

var initInstallSkill bool
var initInstallCodex bool
var initInstallAll bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a .syde/ directory in the current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, _ := os.Getwd()
		dir := filepath.Join(cwd, ".syde")

		if _, err := os.Stat(dir); err == nil {
			return fmt.Errorf(".syde/ already exists")
		}

		// Directory skeleton — one dir per kind + the index dir. No
		// BadgerDB open here: syded owns the index, it'll create the
		// lock file on first open. Keeping init dependency-free means
		// `syde init` works in CI without a syded binary on the PATH.
		for _, kind := range model.AllEntityKinds() {
			subDir := filepath.Join(dir, kind.KindPlural())
			if err := os.MkdirAll(subDir, 0755); err != nil {
				return fmt.Errorf("create %s: %w", subDir, err)
			}
		}
		indexDir := filepath.Join(dir, "index")
		if err := os.MkdirAll(indexDir, 0755); err != nil {
			return fmt.Errorf("create index dir: %w", err)
		}

		// Seed syde.yaml with the project name.
		projectName := filepath.Base(cwd)
		cfg := config.DefaultConfig(projectName)
		if err := config.Save(dir, cfg); err != nil {
			return fmt.Errorf("create syde.yaml: %w", err)
		}

		gitignore := filepath.Join(dir, ".gitignore")
		if err := os.WriteFile(gitignore, []byte("index/\n"), 0644); err != nil {
			return fmt.Errorf("create .gitignore: %w", err)
		}

		// Register the project in ~/.syde/projects.json so any running
		// or future syded instance can serve this project's /api/ paths
		// immediately. Uses the same slug derivation the dashboard +
		// client use.
		absPath, _ := filepath.Abs(cwd)
		slug := dashboard.MakeProjectSlug(projectName, absPath)
		if err := dashboard.RegisterProject(slug, absPath, projectName); err != nil {
			fmt.Printf("  Warning: could not register project with syded: %v\n", err)
		}

		fmt.Printf("Initialized .syde/ in %s\n", cwd)
		fmt.Printf("  %d entity directories created\n", len(model.AllEntityKinds()))
		fmt.Println("  syde.yaml created")
		fmt.Println("  index/ created (gitignored)")
		fmt.Printf("  Project registered as %s\n", slug)

		if initInstallSkill || initInstallAll {
			fmt.Println("\nInstalling Claude Code skill...")
			if err := skill.InstallClaude(cwd); err != nil {
				fmt.Printf("  Warning: %v\n", err)
			} else {
				fmt.Println("  SKILL.md + references + CLAUDE.md + hooks installed")
			}
		}
		if initInstallCodex || initInstallAll {
			fmt.Println("\nInstalling Codex skill...")
			if err := skill.InstallCodex(cwd); err != nil {
				fmt.Printf("  Warning: %v\n", err)
			} else {
				fmt.Println("  .agents skill + AGENTS.md + .codex hooks installed")
			}
		}

		return nil
	},
}

func init() {
	initCmd.Flags().BoolVar(&initInstallSkill, "install-skill", false, "also install Claude Code skill")
	initCmd.Flags().BoolVar(&initInstallCodex, "install-codex", false, "also install Codex skill and hooks")
	initCmd.Flags().BoolVar(&initInstallAll, "install-all", false, "install Claude Code and Codex support")
	rootCmd.AddCommand(initCmd)
}
