package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// FileSHA256 returns the hex-encoded SHA-256 of a file's bytes. Used by
// the requirement verify command and the requirement_stale audit rule
// to detect when code under a refining component has drifted from the
// state the requirement was last verified against.
func FileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// CombinedFilesSHA256 hashes the concatenation of len-prefixed file
// contents for a stable ordering-agnostic representation of a
// component's files on disk. Missing files abort with the underlying
// error so the caller can surface it.
func CombinedFilesSHA256(paths []string) (string, error) {
	h := sha256.New()
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err != nil {
			return "", err
		}
		// prefix with length so "ab"+"c" can't collide with "a"+"bc"
		lenBytes := []byte{byte(len(b) >> 24), byte(len(b) >> 16), byte(len(b) >> 8), byte(len(b))}
		h.Write(lenBytes)
		h.Write(b)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
