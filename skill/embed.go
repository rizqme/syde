package skill

import "embed"

// FS contains all embedded skill files (SKILL.md, references/*, hooks.json).
//
//go:embed SKILL.md hooks.json references/* codex/*
var FS embed.FS
