package dashboard

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const defaultPort = 5703

// Run starts the dashboard daemon process.
func Run(args []string) error {
	fs := flag.NewFlagSet("syded", flag.ContinueOnError)
	daemon := fs.Bool("daemon", false, "run as background daemon")
	stop := fs.Bool("stop", false, "stop running daemon")
	port := fs.Int("port", defaultPort, "server port")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *stop {
		return stopDaemon()
	}

	if *daemon {
		return daemonize(*port)
	}

	return startServer(*port)
}

func startServer(port int) error {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/projects", handleProjects)
	mux.HandleFunc("/api/", handleProjectAPI)

	// Static files + SPA fallback
	mux.HandleFunc("/", handleSPA)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("syded: dashboard at http://localhost%s\n", addr)

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Println("\nsyded: shutting down")
		os.Exit(0)
	}()

	return http.ListenAndServe(addr, mux)
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
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(indexHTML))
}

func daemonize(port int) error {
	// Write PID
	home, _ := os.UserHomeDir()
	globalDir := filepath.Join(home, ".syde")
	os.MkdirAll(globalDir, 0755)

	pidFile := filepath.Join(globalDir, "syded.pid")
	os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0644)

	defer os.Remove(pidFile)
	return startServer(port)
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
