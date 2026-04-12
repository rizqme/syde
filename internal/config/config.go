package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents syde.yaml project configuration.
type Config struct {
	Project        string            `yaml:"project"`
	Version        string            `yaml:"version"`
	DefaultStatus  string            `yaml:"default_status,omitempty"`
	IndexPath      string            `yaml:"index_path,omitempty"`
	GitignoreIndex bool              `yaml:"gitignore_index,omitempty"`
	ComponentPaths map[string][]string `yaml:"component_paths,omitempty"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig(projectName string) *Config {
	return &Config{
		Project:        projectName,
		Version:        "0.1.0",
		DefaultStatus:  "draft",
		IndexPath:      "index",
		GitignoreIndex: true,
	}
}

// Load reads syde.yaml from the given .syde/ directory.
func Load(sydeDir string) (*Config, error) {
	data, err := os.ReadFile(filepath.Join(sydeDir, "syde.yaml"))
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Save writes syde.yaml to the given .syde/ directory.
func Save(sydeDir string, cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(sydeDir, "syde.yaml"), data, 0644)
}

// FindSydeDir walks up from cwd to find a .syde/ directory.
func FindSydeDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		sydeDir := filepath.Join(dir, ".syde")
		if info, err := os.Stat(sydeDir); err == nil && info.IsDir() {
			return sydeDir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
