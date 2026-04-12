package cli

import (
	"fmt"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/memory"
	"github.com/spf13/cobra"
)

var memoryForce bool

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Manage Claude Code memory files",
}

var memorySyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Generate/update Claude memory files from learnings",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		dir := sydeDir
		if dir == "" {
			dir, _ = config.FindSydeDir()
		}
		projectRoot := filepath.Dir(dir)

		mgr := memory.NewManager(store, projectRoot)
		if err := mgr.SyncAll(memoryForce); err != nil {
			return fmt.Errorf("sync: %w", err)
		}

		fmt.Println("Memory files synced to Claude Code.")
		return nil
	},
}

var memoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show syde memory files",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		dir := sydeDir
		if dir == "" {
			dir, _ = config.FindSydeDir()
		}
		projectRoot := filepath.Dir(dir)

		mgr := memory.NewManager(store, projectRoot)
		infos := mgr.ListMemories()

		if len(infos) == 0 {
			fmt.Println("No syde memory files found. Run 'syde memory sync' to generate.")
			return nil
		}

		for _, info := range infos {
			fmt.Printf("  %-40s %d bytes\n", info.File, info.Size)
		}
		return nil
	},
}

var memoryCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove all syde memory files",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openStore()
		if err != nil {
			return err
		}
		defer store.Close()

		dir := sydeDir
		if dir == "" {
			dir, _ = config.FindSydeDir()
		}
		projectRoot := filepath.Dir(dir)

		mgr := memory.NewManager(store, projectRoot)
		if err := mgr.Clean(); err != nil {
			return err
		}

		fmt.Println("Removed all syde memory files.")
		return nil
	},
}

func init() {
	memorySyncCmd.Flags().BoolVar(&memoryForce, "force", false, "force regeneration")
	memoryCmd.AddCommand(memorySyncCmd)
	memoryCmd.AddCommand(memoryListCmd)
	memoryCmd.AddCommand(memoryCleanCmd)
	rootCmd.AddCommand(memoryCmd)
}
