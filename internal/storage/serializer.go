package storage

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/feedloop/syde/internal/model"
	"gopkg.in/yaml.v3"
)

const frontmatterDelimiter = "---"

// Marshal serializes an entity with body into YAML frontmatter + markdown body.
func Marshal(e model.Entity, body string) ([]byte, error) {
	fm, err := yaml.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("marshal frontmatter: %w", err)
	}

	var buf bytes.Buffer
	buf.WriteString(frontmatterDelimiter)
	buf.WriteByte('\n')
	buf.Write(fm)
	buf.WriteString(frontmatterDelimiter)
	buf.WriteByte('\n')
	if body != "" {
		buf.WriteByte('\n')
		buf.WriteString(body)
		if !strings.HasSuffix(body, "\n") {
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes(), nil
}

// Unmarshal parses YAML frontmatter + markdown body into an entity.
func Unmarshal(data []byte, kind model.EntityKind) (model.Entity, string, error) {
	fm, body, err := splitFrontmatter(data)
	if err != nil {
		return nil, "", err
	}

	entity := model.NewEntityForKind(kind)
	if err := yaml.Unmarshal(fm, entity); err != nil {
		return nil, "", fmt.Errorf("parse frontmatter: %w", err)
	}

	return entity, body, nil
}

// UnmarshalAuto detects the entity kind from frontmatter and unmarshals.
func UnmarshalAuto(data []byte) (model.Entity, string, error) {
	fm, body, err := splitFrontmatter(data)
	if err != nil {
		return nil, "", err
	}

	// First pass: extract kind
	var base model.BaseEntity
	if err := yaml.Unmarshal(fm, &base); err != nil {
		return nil, "", fmt.Errorf("parse base: %w", err)
	}

	entity := model.NewEntityForKind(base.Kind)
	if err := yaml.Unmarshal(fm, entity); err != nil {
		return nil, "", fmt.Errorf("parse entity: %w", err)
	}

	return entity, body, nil
}

// splitFrontmatter splits "---\n<yaml>\n---\n<body>" into parts.
func splitFrontmatter(data []byte) ([]byte, string, error) {
	s := string(data)

	// Must start with ---
	if !strings.HasPrefix(s, frontmatterDelimiter) {
		return nil, "", fmt.Errorf("missing frontmatter delimiter")
	}

	// Find the closing ---
	rest := s[len(frontmatterDelimiter)+1:]
	idx := strings.Index(rest, "\n"+frontmatterDelimiter)
	if idx < 0 {
		return nil, "", fmt.Errorf("missing closing frontmatter delimiter")
	}

	fm := rest[:idx]
	body := rest[idx+len("\n"+frontmatterDelimiter):]
	body = strings.TrimPrefix(body, "\n")
	body = strings.TrimPrefix(body, "\n")

	return []byte(fm), body, nil
}

// ComputeLineMap calculates line numbers for each YAML field in the frontmatter.
func ComputeLineMap(data []byte) map[string][2]int {
	lines := strings.Split(string(data), "\n")
	result := make(map[string][2]int)

	fmStart := -1
	fmEnd := -1
	bodyStart := -1
	inFrontmatter := false

	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)

		if trimmed == frontmatterDelimiter {
			if !inFrontmatter {
				fmStart = lineNum
				inFrontmatter = true
			} else {
				fmEnd = lineNum
				inFrontmatter = false
				bodyStart = lineNum + 1
			}
			continue
		}

		if inFrontmatter && len(trimmed) > 0 && !strings.HasPrefix(trimmed, "-") && !strings.HasPrefix(trimmed, " ") {
			// Top-level YAML field
			colonIdx := strings.Index(trimmed, ":")
			if colonIdx > 0 {
				fieldName := trimmed[:colonIdx]
				// Find the end of this field (next top-level field or end of frontmatter)
				endLine := lineNum
				for j := i + 1; j < len(lines); j++ {
					nextTrimmed := strings.TrimSpace(lines[j])
					if nextTrimmed == frontmatterDelimiter {
						break
					}
					if len(nextTrimmed) > 0 && !strings.HasPrefix(nextTrimmed, " ") && !strings.HasPrefix(nextTrimmed, "-") {
						break
					}
					endLine = j + 1
				}
				result[fieldName] = [2]int{lineNum, endLine}
			}
		}
	}

	if fmStart > 0 {
		result["frontmatter"] = [2]int{fmStart, fmEnd}
	}
	if bodyStart > 0 {
		result["body"] = [2]int{bodyStart, len(lines)}
	}

	return result
}
