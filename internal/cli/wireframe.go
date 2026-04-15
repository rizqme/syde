package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/feedloop/syde/internal/model"
	"github.com/feedloop/syde/internal/uiml"
	"github.com/spf13/cobra"
)

// `syde wireframe` is the command group for working with screen
// contracts' UIML wireframes outside the dashboard. The dashboard
// renders wireframes server-side via FormatJSON for browser viewing;
// these commands let you render the same wireframe to stdout, a
// file, or a PNG image without leaving the terminal — useful for
// iterating on wireframe authoring.

var wireframeCmd = &cobra.Command{
	Use:   "wireframe",
	Short: "Render screen contract wireframes (UIML) to html / ascii / image",
}

var (
	wfRenderFormat string
	wfRenderOut    string
	wfRenderOpen   bool
)

var wireframeRenderCmd = &cobra.Command{
	Use:   "render <contract-slug>",
	Short: "Render a screen contract's UIML wireframe",
	Long: `Render the UIML wireframe stored on a screen contract.

Output formats:
  html  (default) self-contained dark-mode HTML wireframe
  ascii           plain-text ASCII rendering (80 cols)
  image           PNG screenshot via headless Google Chrome

By default the rendered output goes to stdout. Use --out <path> to
write to a file, or --open to write to a temp file and open it with
the system default app. --format image REQUIRES --out (PNGs to
stdout would be binary garbage in the terminal).

Examples:
  syde wireframe render components-inbox-screen
  syde wireframe render components-inbox-screen --format ascii
  syde wireframe render components-inbox-screen --out /tmp/x.html --open
  syde wireframe render components-inbox-screen --format image --out /tmp/x.png
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := openWriteClient()
		if err != nil {
			return err
		}
		defer store.Close()

		entity, _, err := store.GetByKind(model.KindContract, args[0])
		if err != nil {
			return fmt.Errorf("contract not found: %s", args[0])
		}
		c, ok := entity.(*model.ContractEntity)
		if !ok {
			return fmt.Errorf("entity %s is not a contract", args[0])
		}
		if c.ContractKind != "screen" {
			return fmt.Errorf("contract %q has contract_kind=%q (expected \"screen\") — only screen contracts carry wireframes", c.Name, c.ContractKind)
		}
		if c.Wireframe == "" {
			return fmt.Errorf("contract %q has no wireframe — set one via syde update %s --wireframe '<screen>...</screen>'", c.Name, args[0])
		}

		result := uiml.Parse(c.Wireframe)
		for _, e := range result.Errors {
			fmt.Fprintf(os.Stderr, "  WARN: UIML parse: %s\n", e.Error())
		}

		switch wfRenderFormat {
		case "", "html":
			html := uiml.RenderWireframeHTML(result.Nodes)
			return wfWriteOrOpen([]byte(html), args[0], "html")

		case "ascii":
			ascii := uiml.RenderASCII(result.Nodes, 80)
			return wfWriteOrOpen([]byte(ascii), args[0], "txt")

		case "image":
			if wfRenderOut == "" {
				return fmt.Errorf("--format image requires --out <path.png>")
			}
			html := uiml.RenderWireframeHTML(result.Nodes)
			tmpHTML := filepath.Join(os.TempDir(), "syde-wireframe-"+args[0]+".html")
			if err := os.WriteFile(tmpHTML, []byte(html), 0644); err != nil {
				return fmt.Errorf("write temp html: %w", err)
			}
			defer os.Remove(tmpHTML)
			if err := wfChromeScreenshot(tmpHTML, wfRenderOut); err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "→ wrote %s\n", wfRenderOut)
			if wfRenderOpen {
				return wfOpenInDefault(wfRenderOut)
			}
			return nil

		default:
			return fmt.Errorf("unknown --format %q (expected html, ascii, or image)", wfRenderFormat)
		}
	},
}

// wfWriteOrOpen sends the rendered bytes to stdout, --out, or a
// temp file (when --open is set without --out). Returns the path
// in the temp/--out cases so the caller can echo it.
func wfWriteOrOpen(data []byte, slug, ext string) error {
	if wfRenderOut == "" && !wfRenderOpen {
		_, err := os.Stdout.Write(data)
		return err
	}
	out := wfRenderOut
	if out == "" {
		out = filepath.Join(os.TempDir(), "syde-wireframe-"+slug+"."+ext)
	}
	if err := os.WriteFile(out, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", out, err)
	}
	fmt.Fprintf(os.Stderr, "→ wrote %s\n", out)
	if wfRenderOpen {
		return wfOpenInDefault(out)
	}
	return nil
}

// wfChromeScreenshot drives headless Google Chrome to capture a PNG
// of a local HTML file. Mirrors scripts/wireframe-shot.sh but
// targets a file:// URL so we don't need syded running.
func wfChromeScreenshot(htmlPath, outPath string) error {
	chrome := wfChromeBinary()
	if chrome == "" {
		return fmt.Errorf("Google Chrome not found (looked for /Applications/Google Chrome.app/.../Google Chrome on macOS, google-chrome / chromium on linux)")
	}
	args := []string{
		"--headless",
		"--disable-gpu",
		"--hide-scrollbars",
		"--window-size=1440,900",
		"--virtual-time-budget=2000",
		"--screenshot=" + outPath,
		"file://" + htmlPath,
	}
	out, err := exec.Command(chrome, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("chrome screenshot failed: %w\n%s", err, string(out))
	}
	if info, statErr := os.Stat(outPath); statErr != nil || info.Size() == 0 {
		return fmt.Errorf("chrome did not write %s", outPath)
	}
	return nil
}

// wfChromeBinary returns the first existing Chrome binary path we
// know about for the current OS, or "" if none.
func wfChromeBinary() string {
	candidates := []string{}
	switch runtime.GOOS {
	case "darwin":
		candidates = append(candidates,
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
		)
	default:
		candidates = append(candidates,
			"/usr/bin/google-chrome",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/chromium",
			"/usr/bin/chromium-browser",
		)
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	// Fall back to PATH lookup
	for _, name := range []string{"google-chrome", "chromium", "chromium-browser", "chrome"} {
		if p, err := exec.LookPath(name); err == nil {
			return p
		}
	}
	return ""
}

// wfOpenInDefault opens the file with the OS default app.
func wfOpenInDefault(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	return cmd.Start()
}

func init() {
	wireframeRenderCmd.Flags().StringVar(&wfRenderFormat, "format", "html", "output format: html | ascii | image")
	wireframeRenderCmd.Flags().StringVar(&wfRenderOut, "out", "", "write to this path instead of stdout (REQUIRED for --format image)")
	wireframeRenderCmd.Flags().BoolVar(&wfRenderOpen, "open", false, "after writing, open the file with the system default app")

	wireframeCmd.AddCommand(wireframeRenderCmd)
	rootCmd.AddCommand(wireframeCmd)
}
