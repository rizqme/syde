package tree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Size beyond which a file is tracked but not summarized (1 MiB).
const MaxSummarizableBytes = 1 << 20

// Size of the binary-detection sniff window.
const sniffBytes = 512

// WalkedFile is the output of WalkProject for a single file. Folders are
// implied by the unique parent paths.
type WalkedFile struct {
	Path   string // relative to project root, forward slashes
	Size   int64
	Hash   string // sha256 hex
	Mtime  string // RFC3339
	Binary bool
}

// WalkProject scans the project tree under root, honoring the matcher,
// and returns a WalkedFile for every file. Directories are NOT returned
// (the scan step derives folder nodes from parents).
func WalkProject(root string, m *Matcher) ([]WalkedFile, error) {
	var files []WalkedFile
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(absRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(absRoot, path)
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}

		if info.IsDir() {
			if m.Match(rel, true) {
				return filepath.SkipDir
			}
			return nil
		}

		if m.Match(rel, false) {
			return nil
		}

		// Hash the file and detect binary
		hash, binary, err := hashAndSniff(path)
		if err != nil {
			// Skip unreadable files silently
			return nil
		}

		files = append(files, WalkedFile{
			Path:   rel,
			Size:   info.Size(),
			Hash:   hash,
			Mtime:  info.ModTime().UTC().Format(time.RFC3339),
			Binary: binary,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// hashAndSniff returns (sha256hex, isBinary, err) for a file, reading it
// once. Uses the first 512 bytes for binary detection.
func hashAndSniff(path string) (string, bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", false, err
	}
	defer f.Close()

	h := sha256.New()

	// Read first sniffBytes into a buffer AND into the hash
	sniff := make([]byte, sniffBytes)
	n, err := io.ReadFull(f, sniff)
	if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
		return "", false, err
	}
	sniff = sniff[:n]
	h.Write(sniff)

	// Continue hashing the rest
	if _, err := io.Copy(h, f); err != nil {
		return "", false, err
	}

	binary := bytes.IndexByte(sniff, 0) >= 0
	return hex.EncodeToString(h.Sum(nil)), binary, nil
}
