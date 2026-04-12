package memory

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/storage"
)

// Manager handles Claude Code memory file generation.
type Manager struct {
	store      *storage.Store
	memoryPath string
}

// NewManager creates a memory manager.
func NewManager(store *storage.Store, projectRoot string) *Manager {
	return &Manager{
		store:      store,
		memoryPath: detectMemoryPath(projectRoot),
	}
}

// SyncAll generates all memory files from learnings + overview.
func (m *Manager) SyncAll(force bool) error {
	if err := os.MkdirAll(m.memoryPath, 0755); err != nil {
		return fmt.Errorf("create memory dir: %w", err)
	}

	// Generate overview memory
	overview := m.generateOverview()
	if err := m.writeMemoryFile("syde_overview.md", overview); err != nil {
		return err
	}

	// Generate per-learning memories
	learnings, _ := m.store.List(model.KindLearning)
	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		content := m.generateLearningMemory(l)
		slug := slugify(l.Name)
		filename := fmt.Sprintf("syde_learn_%s.md", slug)
		if err := m.writeMemoryFile(filename, content); err != nil {
			return err
		}
	}

	// Update MEMORY.md
	return m.updateMemoryIndex(learnings)
}

// Clean removes all syde memory files.
func (m *Manager) Clean() error {
	entries, err := os.ReadDir(m.memoryPath)
	if err != nil {
		return nil // directory doesn't exist
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "syde_") {
			os.Remove(filepath.Join(m.memoryPath, entry.Name()))
		}
	}
	// Remove MEMORY.md section
	return m.removeMemoryIndexSection()
}

// ListMemories returns info about all syde memory files.
func (m *Manager) ListMemories() []MemoryInfo {
	var infos []MemoryInfo

	entries, err := os.ReadDir(m.memoryPath)
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "syde_") {
			info, _ := entry.Info()
			infos = append(infos, MemoryInfo{
				File: entry.Name(),
				Size: info.Size(),
			})
		}
	}
	return infos
}

// MemoryInfo describes a memory file.
type MemoryInfo struct {
	File string
	Size int64
}

func (m *Manager) generateOverview() string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("name: syde-overview\n")
	sb.WriteString("description: This project has a syde design model — run syde status for overview, syde get <name> for details, syde learn about <entity> for gotchas\n")
	sb.WriteString("type: project\n")
	sb.WriteString("---\n\n")

	sb.WriteString("**Design model**: `.syde/` contains the architecture design model.\n\n")
	sb.WriteString("**Quick commands**:\n")
	sb.WriteString("- `syde status` — entity counts, plan progress, learning count\n")
	sb.WriteString("- `syde get <name>` — full entity details\n")
	sb.WriteString("- `syde search <query>` — find entities by keyword\n")
	sb.WriteString("- `syde graph <name>` — show entity relationships\n")
	sb.WriteString("- `syde learn about <entity>` — get learnings/gotchas\n")
	sb.WriteString("- `syde remember \"<text>\" --entity <name>` — save a learning\n")
	sb.WriteString("- `syde query --component <name> --full` — complete context\n")

	// Add counts
	counts := make(map[model.EntityKind]int)
	for _, kind := range model.AllEntityKinds() {
		entities, _ := m.store.List(kind)
		if len(entities) > 0 {
			counts[kind] = len(entities)
		}
	}

	if len(counts) > 0 {
		sb.WriteString("\n**Entities**: ")
		parts := []string{}
		for kind, count := range counts {
			parts = append(parts, fmt.Sprintf("%d %s", count, kind.KindPlural()))
		}
		sb.WriteString(strings.Join(parts, ", "))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m *Manager) generateLearningMemory(l *model.LearningEntity) string {
	var sb strings.Builder

	catUpper := strings.ToUpper(string(l.Category))
	desc := fmt.Sprintf("%s — %s", catUpper, l.Description)

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("name: syde-learn-%s\n", slugify(l.Name)))
	sb.WriteString(fmt.Sprintf("description: %s\n", desc))
	sb.WriteString("type: project\n")
	sb.WriteString("---\n\n")

	sb.WriteString(fmt.Sprintf("**%s**\n\n", l.Description))

	if len(l.EntityRefs) > 0 {
		sb.WriteString(fmt.Sprintf("**Entities**: %s\n\n", strings.Join(l.EntityRefs, ", ")))
	}

	sb.WriteString(fmt.Sprintf("**Category**: %s | **Confidence**: %s | **Source**: %s\n\n", l.Category, l.ConfLevel, l.Source))
	sb.WriteString(fmt.Sprintf("**How to apply:** Check `.syde/learnings/%s.md` for full details.\n", slugify(l.Name)))

	return sb.String()
}

func (m *Manager) writeMemoryFile(filename, content string) error {
	path := filepath.Join(m.memoryPath, filename)
	return os.WriteFile(path, []byte(content), 0644)
}

func (m *Manager) updateMemoryIndex(learnings []model.EntityWithBody) error {
	memoryMD := filepath.Join(m.memoryPath, "MEMORY.md")

	// Read existing content
	existing, _ := os.ReadFile(memoryMD)
	content := string(existing)

	// Build syde section
	var section strings.Builder
	section.WriteString("<!-- syde:start -->\n")
	section.WriteString(fmt.Sprintf("- [syde](syde_overview.md) — design model CLI: run `syde status` for overview\n"))

	for _, ewb := range learnings {
		l := ewb.Entity.(*model.LearningEntity)
		catUpper := strings.ToUpper(string(l.Category))
		slug := slugify(l.Name)
		desc := l.Description
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		section.WriteString(fmt.Sprintf("- [%s](syde_learn_%s.md) — %s: %s\n", l.Name, slug, catUpper, desc))
	}

	section.WriteString("<!-- syde:end -->\n")

	// Replace or append
	startMarker := "<!-- syde:start -->"
	endMarker := "<!-- syde:end -->"

	if startIdx := strings.Index(content, startMarker); startIdx >= 0 {
		if endIdx := strings.Index(content, endMarker); endIdx >= 0 {
			content = content[:startIdx] + section.String() + content[endIdx+len(endMarker)+1:]
		}
	} else {
		if content != "" && !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		content += "\n" + section.String()
	}

	return os.WriteFile(memoryMD, []byte(content), 0644)
}

func (m *Manager) removeMemoryIndexSection() error {
	memoryMD := filepath.Join(m.memoryPath, "MEMORY.md")
	data, err := os.ReadFile(memoryMD)
	if err != nil {
		return nil
	}

	content := string(data)
	startMarker := "<!-- syde:start -->"
	endMarker := "<!-- syde:end -->"

	startIdx := strings.Index(content, startMarker)
	endIdx := strings.Index(content, endMarker)

	if startIdx >= 0 && endIdx >= 0 {
		content = content[:startIdx] + content[endIdx+len(endMarker)+1:]
		return os.WriteFile(memoryMD, []byte(strings.TrimRight(content, "\n")+"\n"), 0644)
	}
	return nil
}

func detectMemoryPath(projectRoot string) string {
	absPath, _ := filepath.Abs(projectRoot)
	// Convert path to Claude Code format: /Users/user/project → -Users-user-project
	hash := strings.ReplaceAll(absPath, string(filepath.Separator), "-")
	if !strings.HasPrefix(hash, "-") {
		hash = "-" + hash
	}

	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude", "projects", hash, "memory")
}

func slugify(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = strings.ReplaceAll(s, " ", "-")
	// Simple slugification
	var result []byte
	for _, c := range []byte(s) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' {
			result = append(result, c)
		}
	}
	s = string(result)
	// Remove consecutive dashes
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}
