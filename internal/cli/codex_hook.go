package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type codexHookInput struct {
	SessionID        string `json:"session_id"`
	TranscriptPath   string `json:"transcript_path"`
	Cwd              string `json:"cwd"`
	HookEventName    string `json:"hook_event_name"`
	Model            string `json:"model"`
	TurnID           string `json:"turn_id"`
	Source           string `json:"source"`
	Prompt           string `json:"prompt"`
	StopHookActive   bool   `json:"stop_hook_active"`
	ToolName         string `json:"tool_name"`
	ToolUseID        string `json:"tool_use_id"`
	LastAssistantMsg string `json:"last_assistant_message"`
	ToolInput        struct {
		Command string `json:"command"`
	} `json:"tool_input"`
	ToolResponse interface{} `json:"tool_response"`
}

type codexHookOutput struct {
	Continue           *bool                  `json:"continue,omitempty"`
	StopReason         string                 `json:"stopReason,omitempty"`
	SystemMessage      string                 `json:"systemMessage,omitempty"`
	Decision           string                 `json:"decision,omitempty"`
	Reason             string                 `json:"reason,omitempty"`
	HookSpecificOutput map[string]interface{} `json:"hookSpecificOutput,omitempty"`
}

var codexHookCmd = &cobra.Command{
	Use:    "codex-hook",
	Short:  "Run a Codex lifecycle hook handler",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var input codexHookInput
		if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
			return fmt.Errorf("decode hook input: %w", err)
		}
		if input.Cwd != "" {
			if err := os.Chdir(input.Cwd); err != nil {
				return fmt.Errorf("chdir %s: %w", input.Cwd, err)
			}
		}
		if _, err := findSydeDirHelper(); err != nil {
			return nil
		}

		switch input.HookEventName {
		case "SessionStart":
			return codexSessionStart(input)
		case "UserPromptSubmit":
			return codexUserPromptSubmit(input)
		case "PreToolUse":
			return codexPreToolUse(input)
		case "PostToolUse":
			return codexPostToolUse(input)
		case "Stop":
			return codexStop(input)
		default:
			return nil
		}
	},
}

func codexSessionStart(input codexHookInput) error {
	ctx, err := runSyde("context", "--json")
	if err != nil {
		return writeCodexJSON(codexHookOutput{
			SystemMessage: "syde: could not load architecture context: " + shortErr(err),
		})
	}

	var b strings.Builder
	b.WriteString("syde architecture context: ")
	b.WriteString(strings.TrimSpace(string(ctx)))
	if _, err := runSyde("tree", "status", "--strict"); err != nil {
		b.WriteString("\n\nsyde tree summary is stale. Run `syde tree scan`, then summarize stale leaves until `syde tree status --strict` exits 0.")
	}

	return writeCodexJSON(codexHookOutput{
		HookSpecificOutput: map[string]interface{}{
			"hookEventName":     "SessionStart",
			"additionalContext": b.String(),
		},
	})
}

func codexUserPromptSubmit(input codexHookInput) error {
	msg := "syde is active in this repository. Before code changes: refresh the summary tree, use `syde query` for context, create/approve a syde plan, start a syde task, and finish with `syde sync check --strict`."
	if _, err := runSyde("tree", "status", "--strict"); err != nil {
		msg += " Current tree status is stale, so Phase 0 should run first."
	}
	if ref, ok := captureCodexPromptRequirement(input); ok {
		msg += " Captured this user prompt as requirement `" + ref + "`."
	}
	return writeCodexJSON(codexHookOutput{
		HookSpecificOutput: map[string]interface{}{
			"hookEventName":     "UserPromptSubmit",
			"additionalContext": msg,
		},
	})
}

func captureCodexPromptRequirement(input codexHookInput) (string, bool) {
	prompt := strings.TrimSpace(input.Prompt)
	if prompt == "" {
		return "", false
	}
	store, err := openWriteClient()
	if err != nil {
		return "", false
	}
	defer store.Close()

	req, _, _, err := createRequirementIfMissing(store, requirementCapture{
		Name:      requirementName("User request", prompt),
		Statement: prompt,
		Source:    "user",
		SourceRef: codexPromptSourceRef(input),
		Rationale: "Captured automatically from Codex UserPromptSubmit.",
	})
	if err != nil {
		return "", false
	}
	return req.CanonicalSlug(), true
}

func codexPromptSourceRef(input codexHookInput) string {
	seed := strings.Join([]string{
		input.SessionID,
		input.TurnID,
		input.TranscriptPath,
		input.Prompt,
	}, "\x00")
	sum := sha256.Sum256([]byte(seed))
	session := input.SessionID
	if session == "" {
		session = "unknown-session"
	}
	turn := input.TurnID
	if turn == "" {
		turn = fmt.Sprintf("%x", sum[:6])
	}
	return "codex:" + session + ":" + turn
}

func codexPreToolUse(input codexHookInput) error {
	if input.ToolName != "" && input.ToolName != "Bash" {
		return nil
	}
	command := input.ToolInput.Command
	if strings.TrimSpace(command) == "" {
		return nil
	}

	if isGitPublishCommand(command) {
		if reason := codexPublishBlockReason(); reason != "" {
			return codexDenyPreTool(reason)
		}
		return nil
	}

	if isLikelyMutatingShellCommand(command) {
		if !hasApprovedPlan() {
			return codexDenyPreTool("syde BLOCK: this looks like a mutating command, but no approved or in-progress syde plan exists. Create a plan, present it, and run `syde plan approve <slug>` before changing files.")
		}
		if !hasActiveTask() {
			return codexDenyPreTool("syde BLOCK: this looks like a mutating command, but no syde task is active. Run `syde task start <slug>` before changing files.")
		}
	}

	return nil
}

func codexPostToolUse(input codexHookInput) error {
	command := input.ToolInput.Command
	if strings.TrimSpace(command) == "" {
		return nil
	}

	var messages []string
	if isGitPublishCommand(command) {
		if reason := codexPublishBlockReason(); reason != "" {
			messages = append(messages, reason)
		}
	}
	if isLikelyMutatingShellCommand(command) {
		if orphans := changedOrUnmappedFiles(); len(orphans) > 0 {
			messages = append(messages, "syde: changed files may be unmapped to design entities: "+strings.Join(orphans, ", ")+". Run `syde update <component> --file <path>` or ignore intentional non-design files.")
		}
	}
	if len(messages) == 0 {
		return nil
	}

	return writeCodexJSON(codexHookOutput{
		SystemMessage: strings.Join(messages, " "),
		HookSpecificOutput: map[string]interface{}{
			"hookEventName":     "PostToolUse",
			"additionalContext": strings.Join(messages, "\n"),
		},
	})
}

func codexStop(input codexHookInput) error {
	_, _ = runSyde("tree", "scan")
	out, err := runSyde("sync", "check", "--strict")
	if err == nil {
		return nil
	}
	reason := "syde FINISH GATE: `syde sync check --strict` failed. Fix every error and warning before ending the session. Highlights: " + firstInterestingLines(string(out), 5)
	if input.StopHookActive {
		cont := false
		return writeCodexJSON(codexHookOutput{
			Continue:      &cont,
			StopReason:    reason,
			SystemMessage: reason,
		})
	}
	return writeCodexJSON(codexHookOutput{
		Decision: "block",
		Reason:   reason,
	})
}

func codexDenyPreTool(reason string) error {
	return writeCodexJSON(codexHookOutput{
		SystemMessage: reason,
		HookSpecificOutput: map[string]interface{}{
			"hookEventName":            "PreToolUse",
			"permissionDecision":       "deny",
			"permissionDecisionReason": reason,
		},
	})
}

func codexPublishBlockReason() string {
	var messages []string
	if out, err := runSyde("task", "list"); err == nil && hasOpenTaskOutput(string(out)) {
		messages = append(messages, "syde tasks are still pending or in progress. Run `syde task done` or `syde task block` before publishing.")
	}
	if out, err := runSyde("sync", "check", "--strict"); err != nil {
		messages = append(messages, "`syde sync check --strict` failed: "+firstInterestingLines(string(out), 5))
	}
	return strings.Join(messages, " ")
}

func hasApprovedPlan() bool {
	out, err := runSyde("plan", "list")
	if err != nil {
		return false
	}
	text := string(out)
	return strings.Contains(text, " approved") || strings.Contains(text, " in-progress") || strings.Contains(text, " in_progress")
}

func hasActiveTask() bool {
	out, err := runSyde("task", "list")
	if err != nil {
		return false
	}
	text := string(out)
	return strings.Contains(text, " in_progress") || strings.Contains(text, " in-progress")
}

func hasOpenTaskOutput(text string) bool {
	return strings.Contains(text, " pending") || strings.Contains(text, " in_progress") || strings.Contains(text, " in-progress")
}

func changedOrUnmappedFiles() []string {
	changed := map[string]bool{}
	for _, args := range [][]string{
		{"diff", "--name-only"},
		{"ls-files", "--others", "--exclude-standard"},
	} {
		out, err := exec.Command("git", args...).Output()
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(out), "\n") {
			p := strings.TrimSpace(line)
			if p == "" || shouldIgnoreHookPath(p) {
				continue
			}
			changed[p] = true
		}
	}

	var unmapped []string
	for p := range changed {
		out, err := runSyde("constraints", "check", p, "--json")
		if err != nil {
			continue
		}
		if strings.TrimSpace(string(out)) == "{}" {
			unmapped = append(unmapped, p)
			if len(unmapped) >= 10 {
				break
			}
		}
	}
	return unmapped
}

func shouldIgnoreHookPath(path string) bool {
	ignore := []string{
		".syde/", ".codex/", ".agents/", ".claude/", ".git/",
		"node_modules/", "vendor/", "web/dist/", "web/node_modules/",
	}
	for _, prefix := range ignore {
		if strings.HasPrefix(path, prefix) || strings.Contains(path, "/"+prefix) {
			return true
		}
	}
	return false
}

var mutatingShellPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(^|[;&|]\s*)(cat|tee)\b.*>{1,2}`),
	regexp.MustCompile(`(^|[[:space:];&|])>{1,2}\s*[^&[:space:]]`),
	regexp.MustCompile(`\b(sed|perl)\s+-i\b`),
	regexp.MustCompile(`\bgofmt\s+-w\b`),
	regexp.MustCompile(`\b(go|npm|pnpm|yarn|bun|cargo)\s+(get|mod|install|add|update)\b`),
	regexp.MustCompile(`(^|[;&|]\s*)(rm|mv|cp|touch|mkdir|install)\b`),
	regexp.MustCompile(`\bapply_patch\b`),
}

func isLikelyMutatingShellCommand(command string) bool {
	lower := strings.ToLower(command)
	for _, re := range mutatingShellPatterns {
		if re.MatchString(lower) {
			return true
		}
	}
	return false
}

func isGitPublishCommand(command string) bool {
	lower := strings.ToLower(command)
	return regexp.MustCompile(`\bgit\s+(commit|push)\b`).MatchString(lower)
}

func runSyde(args ...string) ([]byte, error) {
	exe, err := os.Executable()
	if err != nil || exe == "" {
		exe = "syde"
	}
	cmd := exec.Command(exe, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("%s: %w", strings.Join(append([]string{filepath.Base(exe)}, args...), " "), err)
	}
	return out, nil
}

func writeCodexJSON(out codexHookOutput) error {
	enc := json.NewEncoder(os.Stdout)
	return enc.Encode(out)
}

func shortErr(err error) string {
	if err == nil {
		return ""
	}
	return strings.TrimSpace(err.Error())
}

func firstInterestingLines(text string, max int) string {
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
		if len(lines) >= max {
			break
		}
	}
	if len(lines) == 0 {
		return "no output"
	}
	return strings.Join(lines, " ")
}

func init() {
	rootCmd.AddCommand(codexHookCmd)
}
