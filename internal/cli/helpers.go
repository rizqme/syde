package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/feedloop/syde/internal/client"
	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/dashboard"
	"github.com/feedloop/syde/internal/storage"
)

// openStore opens the Store, auto-detecting the .syde/ directory.
// DEPRECATED: CLI read commands should use openClient() and talk to
// syded over HTTP. Kept for write commands that still need FileStore
// for direct markdown writes (the Idx field is unused by those paths).
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

// openClient auto-launches syded and returns an HTTP client scoped to
// the current project. Every CLI read command goes through this —
// syded is the single owner of the BadgerDB index, so this avoids any
// directory-lock contention.
func openClient() (*client.Client, error) {
	dir := sydeDir
	if dir == "" {
		var err error
		dir, err = config.FindSydeDir()
		if err != nil {
			return nil, fmt.Errorf("no .syde/ directory found (run 'syde init' first)")
		}
	}

	projectRoot := filepath.Dir(dir)
	absPath, err := filepath.Abs(projectRoot)
	if err != nil {
		return nil, err
	}

	projectName := filepath.Base(absPath)
	if cfg, err := config.Load(dir); err == nil && cfg != nil && cfg.Project != "" {
		projectName = cfg.Project
	}

	slug := dashboard.MakeProjectSlug(projectName, absPath)
	if err := dashboard.RegisterProject(slug, absPath, projectName); err != nil {
		return nil, fmt.Errorf("register project with syded: %w", err)
	}

	return client.New(client.Options{SydeDir: dir, ProjectSlug: slug})
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
