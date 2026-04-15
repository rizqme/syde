package utils

import (
	"crypto/rand"
	"regexp"
	"strings"
)

var (
	nonAlphaNum    = regexp.MustCompile(`[^a-z0-9-]+`)
	multiDash      = regexp.MustCompile(`-{2,}`)
	suffixPattern  = regexp.MustCompile(`-[a-z0-9]{4}$`)
	suffixAlphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
)

// Slugify converts a name to a URL-friendly slug (base form, no suffix).
// Example: "syde CLI" → "syde-cli". Use this when you need to compare
// or match the "name stem" of an entity across versions, or to build
// a parent/child path.
func Slugify(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = strings.ReplaceAll(s, " ", "-")
	s = nonAlphaNum.ReplaceAllString(s, "")
	s = multiDash.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// SlugifyWithSuffix returns base slug + 4-char random alphanumeric
// suffix (e.g. "cli-a3f2"). This is the file-level addressable slug
// stored in BaseEntity.Slug and used as the markdown filename.
func SlugifyWithSuffix(name string) string {
	return Slugify(name) + "-" + randSuffix(4)
}

// BaseSlug strips a trailing 4-char suffix from a full slug if present.
// Returns the slug unchanged when no suffix is detected. Used to match
// a full slug against its bare "name stem".
func BaseSlug(slug string) string {
	if suffixPattern.MatchString(slug) {
		return slug[:len(slug)-5]
	}
	return slug
}

// HasSuffix reports whether a slug carries a -XXXX disambiguation
// suffix. Useful to distinguish "cli" (bare) from "cli-a3f2" (full).
func HasSuffix(slug string) bool {
	return suffixPattern.MatchString(slug)
}

func randSuffix(n int) string {
	b := make([]byte, n)
	random := make([]byte, n)
	if _, err := rand.Read(random); err != nil {
		// Extremely unlikely; fall back to a deterministic seed.
		for i := range b {
			b[i] = suffixAlphabet[i%len(suffixAlphabet)]
		}
		return string(b)
	}
	for i := range b {
		b[i] = suffixAlphabet[int(random[i])%len(suffixAlphabet)]
	}
	return string(b)
}
