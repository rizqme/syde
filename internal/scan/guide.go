package scan

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// ScanGuide is the output of programmatic directory analysis.
type ScanGuide struct {
	ProjectRoot  string            `json:"project_root"`
	FileCount    int               `json:"file_count"`
	Languages    map[string]int    `json:"languages"`
	DirectoryMap []DirEntry        `json:"directory_map"`
	KeyFiles     []string          `json:"key_files"`
	ScanStatus   map[string]string `json:"scan_status"`
}

// DirEntry represents a directory in the scan guide.
type DirEntry struct {
	Path           string `json:"path"`
	Files          int    `json:"files"`
	Language       string `json:"language"`
	HasTests       bool   `json:"has_tests,omitempty"`
	HasDockerfile  bool   `json:"has_dockerfile,omitempty"`
	HasPackageJSON bool   `json:"has_package_json,omitempty"`
	HasGoMod       bool   `json:"has_go_mod,omitempty"`
}

// GenerateGuide scans a project directory and produces a scan guide.
func GenerateGuide(root string) (*ScanGuide, error) {
	root, _ = filepath.Abs(root)
	guide := &ScanGuide{
		ProjectRoot: root,
		Languages:   make(map[string]int),
		ScanStatus: map[string]string{
			"round_1": "pending",
			"round_2": "pending",
			"round_3": "pending",
			"round_4": "pending",
			"round_5": "pending",
		},
	}

	dirFiles := make(map[string]int)
	dirLangs := make(map[string]map[string]int)

	keyFileNames := map[string]bool{
		"README.md": true, "readme.md": true,
		"docker-compose.yml": true, "docker-compose.yaml": true,
		"Dockerfile": true, "Makefile": true,
		"go.mod": true, "package.json": true,
		"Cargo.toml": true, "pyproject.toml": true,
		".github": true,
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(root, path)

		// Skip hidden dirs, node_modules, vendor, .syde
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" || name == "__pycache__" {
				return filepath.SkipDir
			}
			return nil
		}

		guide.FileCount++

		// Track languages
		ext := strings.ToLower(filepath.Ext(info.Name()))
		lang := extToLang(ext)
		if lang != "" {
			guide.Languages[lang]++
		}

		// Track directory files
		dir := filepath.Dir(rel)
		dirFiles[dir]++
		if dirLangs[dir] == nil {
			dirLangs[dir] = make(map[string]int)
		}
		if lang != "" {
			dirLangs[dir][lang]++
		}

		// Track key files
		if keyFileNames[info.Name()] {
			guide.KeyFiles = append(guide.KeyFiles, rel)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Build directory map for significant directories
	for dir, count := range dirFiles {
		if count < 2 || dir == "." {
			continue
		}
		// Only include source-like directories
		if isSourceDir(dir) {
			entry := DirEntry{
				Path:  dir,
				Files: count,
			}

			// Determine dominant language
			if langs, ok := dirLangs[dir]; ok {
				maxLang := ""
				maxCount := 0
				for l, c := range langs {
					if c > maxCount {
						maxLang = l
						maxCount = c
					}
				}
				entry.Language = maxLang
			}

			// Check for markers
			fullDir := filepath.Join(root, dir)
			if fileExists(filepath.Join(fullDir, "Dockerfile")) {
				entry.HasDockerfile = true
			}
			if fileExists(filepath.Join(fullDir, "package.json")) {
				entry.HasPackageJSON = true
			}
			if fileExists(filepath.Join(fullDir, "go.mod")) {
				entry.HasGoMod = true
			}
			// Check for test files
			entries, _ := os.ReadDir(fullDir)
			for _, e := range entries {
				name := e.Name()
				if strings.Contains(name, "_test.") || strings.Contains(name, ".test.") || strings.Contains(name, ".spec.") || strings.HasSuffix(name, "_test.go") {
					entry.HasTests = true
					break
				}
			}

			guide.DirectoryMap = append(guide.DirectoryMap, entry)
		}
	}

	return guide, nil
}

// SaveGuide writes the scan guide to .syde/scan-guide.json.
func SaveGuide(sydeDir string, guide *ScanGuide) error {
	data, err := json.MarshalIndent(guide, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(sydeDir, "scan-guide.json"), data, 0644)
}

// LoadGuide reads the scan guide from .syde/scan-guide.json.
func LoadGuide(sydeDir string) (*ScanGuide, error) {
	data, err := os.ReadFile(filepath.Join(sydeDir, "scan-guide.json"))
	if err != nil {
		return nil, err
	}
	var guide ScanGuide
	return &guide, json.Unmarshal(data, &guide)
}

func extToLang(ext string) string {
	switch ext {
	case ".go":
		return "go"
	case ".ts", ".tsx":
		return "typescript"
	case ".js", ".jsx":
		return "javascript"
	case ".py":
		return "python"
	case ".java":
		return "java"
	case ".rs":
		return "rust"
	case ".rb":
		return "ruby"
	case ".proto":
		return "proto"
	case ".sql":
		return "sql"
	case ".yaml", ".yml":
		return "yaml"
	case ".json":
		return "json"
	case ".md":
		return "markdown"
	case ".css", ".scss":
		return "css"
	case ".html":
		return "html"
	default:
		return ""
	}
}

func isSourceDir(dir string) bool {
	parts := strings.Split(dir, string(filepath.Separator))
	for _, p := range parts {
		switch p {
		case "src", "services", "apps", "pkg", "lib", "internal", "cmd",
			"web", "frontend", "client", "api", "gateway", "server",
			"handlers", "controllers", "routes", "models", "domain",
			"integrations", "adapters", "proto", "migrations":
			return true
		}
	}
	return len(parts) <= 2 // Top-level dirs with files
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
