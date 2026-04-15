package scan

import (
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
	"github.com/feedloop/syde/internal/utils"
)

// CoverageEntry represents coverage of a directory by a component.
type CoverageEntry struct {
	Directory     string
	ComponentSlug string
	ComponentName string
	Covered       bool
	IsInfra       bool // migrations, deploy, etc.
}

// CheckCoverage compares scan guide directories against existing components.
func CheckCoverage(guide *ScanGuide, store *storage.Store) ([]CoverageEntry, error) {
	// Load all components
	components, err := store.List(model.KindComponent)
	if err != nil {
		return nil, err
	}

	// Build slug → name map and dir → component map from files field
	compMap := make(map[string]string)
	dirToComp := make(map[string]string) // directory path → component name
	for _, ewb := range components {
		b := ewb.Entity.GetBase()
		slug := utils.Slugify(b.Name)
		compMap[slug] = b.Name
		// Map file globs to directories
		for _, f := range b.Files {
			// Extract directory from glob (e.g., "internal/storage/*.go" → "internal/storage")
			dir := f
			if idx := strings.LastIndex(f, "/"); idx >= 0 {
				dir = f[:idx]
			}
			dirToComp[dir] = b.Name
		}
	}

	var entries []CoverageEntry
	infraDirs := map[string]bool{
		"migrations": true, "deploy": true, "scripts": true,
		"docs": true, "config": true, "terraform": true,
		"k8s": true, "helm": true, ".github": true,
	}

	for _, dir := range guide.DirectoryMap {
		entry := CoverageEntry{
			Directory: dir.Path,
		}

		// Check if this is an infrastructure directory
		parts := strings.Split(dir.Path, "/")
		lastPart := parts[len(parts)-1]
		if infraDirs[lastPart] {
			entry.IsInfra = true
			entries = append(entries, entry)
			continue
		}

		// Try to match directory to a component
		// 1. Check files field mapping (e.g., "internal/storage/*.go" covers "internal/storage")
		if name, ok := dirToComp[dir.Path]; ok {
			entry.Covered = true
			entry.ComponentName = name
			entry.ComponentSlug = utils.Slugify(name)
		} else {
			// 2. Try slug match on directory name
			slug := utils.Slugify(lastPart)
			if name, ok := compMap[slug]; ok {
				entry.Covered = true
				entry.ComponentSlug = slug
				entry.ComponentName = name
			}
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// FormatCoverage renders coverage results as a human-readable report.
func FormatCoverage(entries []CoverageEntry) string {
	var sb strings.Builder
	sb.WriteString("Coverage Report:\n\n")

	covered := 0
	source := 0
	for _, e := range entries {
		if e.IsInfra {
			sb.WriteString(fmt.Sprintf("  %-35s (infrastructure)          -\n", e.Directory))
			continue
		}
		source++
		if e.Covered {
			sb.WriteString(fmt.Sprintf("  %-35s → %s ✓\n", e.Directory, e.ComponentName))
			covered++
		} else {
			sb.WriteString(fmt.Sprintf("  %-35s → (not mapped)            ✗\n", e.Directory))
		}
	}

	sb.WriteString(fmt.Sprintf("\n  Coverage: %d/%d source directories mapped (%d%%)\n",
		covered, source, percent(covered, source)))
	return sb.String()
}

func percent(a, b int) int {
	if b == 0 {
		return 0
	}
	return a * 100 / b
}
