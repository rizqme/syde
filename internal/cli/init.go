package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/skill"
	"github.com/feedloop/syde/internal/storage"
	"github.com/spf13/cobra"
)

var initInstallSkill bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a .syde/ directory in the current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, _ := os.Getwd()
		dir := filepath.Join(cwd, ".syde")

		if _, err := os.Stat(dir); err == nil {
			return fmt.Errorf(".syde/ already exists")
		}

		// Create directory structure
		for _, kind := range model.AllEntityKinds() {
			subDir := filepath.Join(dir, kind.KindPlural())
			if err := os.MkdirAll(subDir, 0755); err != nil {
				return fmt.Errorf("create %s: %w", subDir, err)
			}
		}

		// Create index directory
		indexDir := filepath.Join(dir, "index")
		if err := os.MkdirAll(indexDir, 0755); err != nil {
			return fmt.Errorf("create index dir: %w", err)
		}

		// Create syde.yaml
		projectName := filepath.Base(cwd)
		cfg := config.DefaultConfig(projectName)
		if err := config.Save(dir, cfg); err != nil {
			return fmt.Errorf("create syde.yaml: %w", err)
		}

		// Create .gitignore
		gitignore := filepath.Join(dir, ".gitignore")
		if err := os.WriteFile(gitignore, []byte("index/\n"), 0644); err != nil {
			return fmt.Errorf("create .gitignore: %w", err)
		}

		// Initialize empty index
		idx, err := storage.OpenIndex(indexDir)
		if err != nil {
			return fmt.Errorf("init index: %w", err)
		}
		idx.Close()

		fmt.Printf("Initialized .syde/ in %s\n", cwd)
		fmt.Printf("  %d entity directories created\n", len(model.AllEntityKinds()))
		fmt.Println("  syde.yaml created")
		fmt.Println("  index/ created (gitignored)")

		if initInstallSkill {
			fmt.Println("\nInstalling Claude Code skill...")
			if err := skill.Install(cwd); err != nil {
				fmt.Printf("  Warning: %v\n", err)
			} else {
				fmt.Println("  SKILL.md + references + CLAUDE.md + hooks installed")
			}
		}

		return nil
	},
}

func init() {
	initCmd.Flags().BoolVar(&initInstallSkill, "install-skill", false, "also install Claude Code skill")
	rootCmd.AddCommand(initCmd)
}
