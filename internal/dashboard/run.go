package dashboard

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/feedloop/syde/web"
)

const defaultPort = 5703

// Server runtime counters used by /health and the idle-timeout shutdown.
var (
	serverStarted  time.Time
	lastRequestTS  int64 // unix seconds; atomic
	buildVersion   = "dev"
)

// Run starts the dashboard daemon process.
func Run(args []string) error {
	fs := flag.NewFlagSet("syded", flag.ContinueOnError)
	daemon := fs.Bool("daemon", false, "run as background daemon")
	stop := fs.Bool("stop", false, "stop running daemon")
	port := fs.Int("port", defaultPort, "server port")
	idle := fs.Duration("idle-timeout", 30*time.Minute, "shut down after this long without /api traffic (0 to disable)")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *stop {
		return stopDaemon()
	}

	if *daemon {
		return daemonize(*port, *idle)
	}

	return startServer(*port, *idle)
}

// lastRequestMiddleware records every non-/health request for the idle
// shutdown ticker to reference. /health is excluded so the auto-launch
// poll from the CLI doesn't keep the daemon alive forever.
func lastRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			atomic.StoreInt64(&lastRequestTS, time.Now().Unix())
		}
		next.ServeHTTP(w, r)
	})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uptime := int64(0)
	if !serverStarted.IsZero() {
		uptime = int64(time.Since(serverStarted).Seconds())
	}
	last := atomic.LoadInt64(&lastRequestTS)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":              true,
		"version":         buildVersion,
		"uptime_sec":      uptime,
		"last_request_ts": last,
	})
}

func startServer(port int, idleTimeout time.Duration) error {
	serverStarted = time.Now()
	atomic.StoreInt64(&lastRequestTS, serverStarted.Unix())

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/api/projects", handleProjects)
	mux.HandleFunc("/api/", handleProjectAPI)
	mux.HandleFunc("/ws/", handleWebSocket)
	go startProjectWatchers()
	mux.HandleFunc("/", handleSPA)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("syded: dashboard at http://localhost%s\n", addr)

	srv := &http.Server{Addr: addr, Handler: lastRequestMiddleware(mux)}

	// Idle-timeout auto shutdown. CLI auto-launch fires syded for a
	// single command in CI contexts — we don't want the daemon to
	// outlive its usefulness. 0 disables (nice for dev).
	if idleTimeout > 0 {
		go func() {
			tick := time.NewTicker(60 * time.Second)
			defer tick.Stop()
			for range tick.C {
				last := atomic.LoadInt64(&lastRequestTS)
				if time.Since(time.Unix(last, 0)) > idleTimeout {
					fmt.Println("syded: idle timeout, shutting down")
					_ = srv.Close()
					return
				}
			}
		}()
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("\nsyded: shutting down")
		_ = srv.Close()
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func handleProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	projects, err := loadProjectRegistry()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"projects": []interface{}{}})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"projects": projects})
}

func handleSPA(w http.ResponseWriter, r *http.Request) {
	distFS, err := fs.Sub(web.DistFS, "dist")
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body>Dashboard not built. Run: make frontend</body></html>"))
		return
	}

	// Try to serve the exact file first (JS, CSS, assets)
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	// Strip leading slash for fs.Open
	fsPath := path[1:]
	if f, err := distFS.Open(fsPath); err == nil {
		f.Close()
		http.FileServer(http.FS(distFS)).ServeHTTP(w, r)
		return
	}

	// SPA fallback: serve index.html for all other routes
	indexData, err := fs.ReadFile(distFS, "index.html")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(indexData)
}

func startProjectWatchers() {
	projects, err := loadProjectRegistry()
	if err != nil {
		return
	}
	for _, p := range projects {
		sydeDir := filepath.Join(p.Path, ".syde")
		go watchSydeDir(sydeDir)
	}
}

func daemonize(port int, idleTimeout time.Duration) error {
	// Write PID
	home, _ := os.UserHomeDir()
	globalDir := filepath.Join(home, ".syde")
	os.MkdirAll(globalDir, 0755)

	pidFile := filepath.Join(globalDir, "syded.pid")
	os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0644)

	defer os.Remove(pidFile)
	return startServer(port, idleTimeout)
}

func stopDaemon() error {
	home, _ := os.UserHomeDir()
	pidFile := filepath.Join(home, ".syde", "syded.pid")

	data, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("syded not running (no PID file)")
	}

	var pid int
	fmt.Sscanf(string(data), "%d", &pid)

	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process not found: %d", pid)
	}

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("signal: %w", err)
	}

	os.Remove(pidFile)
	fmt.Printf("syded: stopped (PID %d)\n", pid)
	return nil
}
