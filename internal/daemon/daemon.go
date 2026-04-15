// Package daemon provides a tiny helper for CLI processes to make sure
// a local syded instance is reachable before sending requests. Every
// syde CLI command that needs the BadgerDB-backed index calls
// EnsureRunning before hitting the HTTP API — if no syded is up it
// forks one in the background and polls /health until it responds.
package daemon

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// DefaultPort mirrors internal/dashboard.defaultPort. Kept as a local
// copy so this package doesn't import dashboard (which would pull in
// web SPA embeds and blow up the CLI binary size).
const DefaultPort = 5703

// EnsureRunning guarantees a local syded is reachable at localhost:port
// on /health. If one is already running it returns immediately. If not
// it tries to locate a `syded` binary on PATH or next to the current
// executable, spawns it with `-daemon -port <port> -idle-timeout 30m`,
// and polls /health every 50ms for up to 3 seconds.
//
// The caller passes a logger function for progress messages (typically
// os.Stderr) so CLI output can pipe around it cleanly.
func EnsureRunning(port int, logf func(format string, args ...interface{})) error {
	if port == 0 {
		port = DefaultPort
	}
	if logf == nil {
		logf = func(string, ...interface{}) {}
	}

	if ping(port, 300*time.Millisecond) == nil {
		return nil
	}

	bin, err := findSydedBinary()
	if err != nil {
		return fmt.Errorf("syded not running and binary not found on PATH: %w", err)
	}

	logf("syde: starting syded from %s\n", bin)
	logFile, logPath := openLog()
	cmd := exec.Command(bin,
		"-daemon",
		"-port", fmt.Sprintf("%d", port),
		"-idle-timeout", "30m",
	)
	// Detach stdio so the child survives the CLI exiting.
	cmd.Stdin = nil
	if logFile != nil {
		fmt.Fprintf(logFile, "\n--- syded auto-start %s ---\n", time.Now().Format(time.RFC3339))
		cmd.Stdout = logFile
		cmd.Stderr = logFile
	} else {
		cmd.Stdout = io.Discard
		cmd.Stderr = io.Discard
	}
	cmd.SysProcAttr = detachedSysProc()
	if err := cmd.Start(); err != nil {
		if logFile != nil {
			_ = logFile.Close()
		}
		return fmt.Errorf("spawn syded: %w", err)
	}
	if logFile != nil {
		_ = logFile.Close()
	}
	// Release the child so it isn't reaped when the CLI exits.
	_ = cmd.Process.Release()

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if ping(port, 100*time.Millisecond) == nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	if tail := tailLog(logPath, 4096); tail != "" {
		return fmt.Errorf("syded failed to become ready within 3s — last %s:\n%s", logPath, tail)
	}
	return fmt.Errorf("syded failed to become ready within 3s — check %s", logPath)
}

func openLog() (*os.File, string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, "~/.syde/syded.log"
	}
	dir := filepath.Join(home, ".syde")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, filepath.Join(dir, "syded.log")
	}
	path := filepath.Join(dir, "syded.log")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, path
	}
	return f, path
}

func tailLog(path string, maxBytes int64) string {
	if path == "" || strings.HasPrefix(path, "~") {
		return ""
	}
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil || info.Size() == 0 {
		return ""
	}
	start := info.Size() - maxBytes
	if start < 0 {
		start = 0
	}
	if _, err := f.Seek(start, io.SeekStart); err != nil {
		return ""
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func ping(port int, timeout time.Duration) error {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(fmt.Sprintf("http://localhost:%d/health", port))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("health: %s", resp.Status)
	}
	return nil
}

// findSydedBinary locates syded in order of preference: $SYDED_BIN,
// alongside the current executable (installs tend to drop syde + syded
// in the same dir), then PATH. Returns an error if none found.
func findSydedBinary() (string, error) {
	if p := os.Getenv("SYDED_BIN"); p != "" {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	if self, err := os.Executable(); err == nil {
		candidate := filepath.Join(filepath.Dir(self), "syded")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}
	if p, err := exec.LookPath("syded"); err == nil {
		return p, nil
	}
	return "", fmt.Errorf("syded not found")
}
