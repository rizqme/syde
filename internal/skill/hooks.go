package skill

// HooksJSON is the hooks template for Claude Code settings.
var HooksJSON = `{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write",
        "hooks": [
          {
            "type": "command",
            "command": "bash -c 'input=$(cat); file_path=$(echo \"$input\" | jq -r \".tool_input.file_path // empty\" 2>/dev/null); if [ -z \"$file_path\" ]; then exit 0; fi; if echo \"$file_path\" | grep -q \".claude/plans/\"; then echo \"{\\\"systemMessage\\\": \\\"Plan file updated. Run syde plan sync or syde plan add-step to track in design model.\\\"}\"; elif [ -d .syde ] && ! echo \"$file_path\" | grep -q \".syde/\\|.claude/\\|node_modules/\\|vendor/\"; then result=$(syde constraints check \"$file_path\" --json 2>/dev/null); if [ \"$result\" = \"{}\" ] || [ -z \"$result\" ]; then echo \"{\\\"systemMessage\\\": \\\"New file not mapped to any syde component. Run syde update <component> --add-rel to link it.\\\"}\"; fi; fi'",
            "timeout": 5000
          }
        ]
      },
      {
        "matcher": "Edit",
        "hooks": [
          {
            "type": "command",
            "command": "bash -c 'if [ -d .syde ]; then input=$(cat); file_path=$(echo \"$input\" | jq -r \".tool_input.file_path // empty\" 2>/dev/null); if [ -n \"$file_path\" ] && ! echo \"$file_path\" | grep -q \".syde/\\|.claude/\"; then result=$(syde constraints check \"$file_path\" --json 2>/dev/null); if [ -n \"$result\" ] && [ \"$result\" != \"{}\" ]; then echo \"{\\\"systemMessage\\\": \\\"syde constraints: $result\\\"}\"; fi; fi; fi'",
            "timeout": 5000
          }
        ]
      }
    ],
    "SessionStart": [
      {
        "matcher": "*",
        "hooks": [
          {
            "type": "command",
            "command": "bash -c 'if [ -d .syde ]; then ctx=$(syde context --json 2>/dev/null); if [ -n \"$ctx\" ]; then echo \"{\\\"systemMessage\\\": \\\"syde architecture context: $ctx\\\"}\"; fi; fi'",
            "timeout": 5000
          }
        ]
      }
    ]
  }
}`
