package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/storage"
)

// openStore opens the Store, auto-detecting the .syde/ directory.
func openStore() (*storage.Store, error) {
	dir := sydeDir
	if dir == "" {
		var err error
		dir, err = config.FindSydeDir()
		if err != nil {
			return nil, fmt.Errorf("no .syde/ directory found (run 'syde init' first)")
		}
	}
	return storage.NewStore(dir)
}

// resolveSydeDir returns the .syde/ directory path.
func resolveSydeDir() string {
	if sydeDir != "" {
		return sydeDir
	}
	dir, _ := config.FindSydeDir()
	return dir
}

// loadProjectConfig loads the syde.yaml config safely.
func loadProjectConfig() (string, string) {
	dir := resolveSydeDir()
	if dir == "" {
		return "unknown", ""
	}
	cfg, err := config.Load(dir)
	if err != nil {
		return filepath.Base(filepath.Dir(dir)), ""
	}
	return cfg.Project, cfg.Version
}

func findSydeDirHelper() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		sd := filepath.Join(dir, ".syde")
		if info, err := os.Stat(sd); err == nil && info.IsDir() {
			return sd, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
