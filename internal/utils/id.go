package utils

import (
	"strings"

	"github.com/feedloop/syde/internal/model"
)

// GenerateID is kept only as a compile-time placeholder for legacy
// callers. Entity IDs are now allocated via storage.NextID (counter
// per kind, persisted in .syde/counters.yaml). Calling this returns an
// empty string — callers must use storage.NextID instead.
//
// Deprecated: use storage.NextID(sydeDir, kind).
func GenerateID(kind model.EntityKind) string { return "" }

// ParseIDKind extracts the entity kind from an ID in the form
// "PFX-0001" (new counter format) or "pfx_xxxxxxxx" (legacy ULID).
// Returns (kind, true) on success.
func ParseIDKind(id string) (model.EntityKind, bool) {
	// New format: PFX-0001
	if i := strings.Index(id, "-"); i > 0 {
		prefix := id[:i]
		for _, k := range model.AllEntityKinds() {
			if k.IDPrefix() == prefix {
				return k, true
			}
		}
	}
	// Legacy format: pfx_xxxxxxxx
	if i := strings.Index(id, "_"); i > 0 {
		prefix := strings.ToLower(id[:i])
		for _, k := range model.AllEntityKinds() {
			if strings.ToLower(k.IDPrefix()) == prefix {
				return k, true
			}
		}
	}
	return "", false
}
