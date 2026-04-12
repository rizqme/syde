package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"

	"github.com/feedloop/syde/internal/config"
	"github.com/feedloop/syde/internal/dashboard"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage the syde dashboard server",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the dashboard server (port 5703)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if isDaemonRunning() {
			fmt.Println("Dashboard server already running on http://localhost:5703")
			return nil
		}

		fmt.Println("Starting dashboard server...")
		if err := startDashboard(); err != nil {
			return fmt.Errorf("start dashboard: %w", err)
		}

		fmt.Println("Dashboard server started at http://localhost:5703")
		return nil
	},
}

var serverStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the dashboard server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return stopDashboard()
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open this project in the dashboard (starts server if needed)",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := sydeDir
		if dir == "" {
			var err error
			dir, err = config.FindSydeDir()
			if err != nil {
				return fmt.Errorf("no .syde/ directory found")
			}
		}

		projectRoot := filepath.Dir(dir)
		absPath, _ := filepath.Abs(projectRoot)

		cfg, _ := config.Load(dir)
		projectName := filepath.Base(absPath)
		if cfg != nil && cfg.Project != "" {
			projectName = cfg.Project
		}

		slug := dashboard.MakeProjectSlug(projectName, absPath)

		if err := dashboard.RegisterProject(slug, absPath, projectName); err != nil {
			return fmt.Errorf("register project: %w", err)
		}

		if !isDaemonRunning() {
			fmt.Println("Starting dashboard server...")
			if err := startDashboard(); err != nil {
				return fmt.Errorf("start dashboard: %w", err)
			}
		}

		url := fmt.Sprintf("http://localhost:5703/%s", slug)
		fmt.Printf("Dashboard: %s\n", url)
		openBrowser(url)

		return nil
	},
}

func isDaemonRunning() bool {
	home, _ := os.UserHomeDir()
	pidFile := filepath.Join(home, ".syde", "syded.pid")
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return false
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return false
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return process.Signal(syscall.Signal(0)) == nil
}

func startDashboard() error {
	exe, _ := os.Executable()
	syded := filepath.Join(filepath.Dir(exe), "syded")

	if _, err := os.Stat(syded); err != nil {
		syded = "syded"
	}

	cmd := exec.Command(syded, "--daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Start()
}

func stopDashboard() error {
	home, _ := os.UserHomeDir()
	pidFile := filepath.Join(home, ".syde", "syded.pid")
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("dashboard not running")
	}
	pid, _ := strconv.Atoi(string(data))
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	if err := process.Signal(syscall.SIGTERM); err != nil {
		return err
	}
	os.Remove(pidFile)
	fmt.Printf("Dashboard stopped (PID %d)\n", pid)
	return nil
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	}
	if cmd != nil {
		cmd.Start()
	}
}

func init() {
	serverCmd.AddCommand(serverStartCmd)
	serverCmd.AddCommand(serverStopCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(openCmd)
}
