---
acceptance: 'All three formats work: ''syde wireframe render components-inbox-screen-c5jh > /tmp/x.html && open /tmp/x.html'' shows the dark wireframe in a browser; ''--format ascii'' prints to stdout; ''--format image --out /tmp/x.png'' writes a PNG that the Read tool can display. Errors cleanly when the slug isn''t a screen contract or when --format image without --out.'
affected_entities:
    - cli-commands
    - uiml-parser
completed_at: "2026-04-15T03:06:45Z"
created_at: "2026-04-15T02:58:11Z"
details: 'New file internal/cli/wireframe.go. Cobra parent ''wireframe'' command + ''render'' subcommand. Args: <contract-slug>. Flags: --format html|ascii|image (default html), --out <path> (default stdout for html/ascii, REQUIRED for image), --open (write to a temp file and ''open'' it via the system default app). Implementation: load contract via write client, type-assert to ContractEntity, error if contract_kind != ''screen'' or wireframe is empty, parse via uiml.Parse, surface parse errors. html mode: call uiml.RenderWireframeHTML and write the HTML to the sink. ascii mode: call uiml.RenderASCII(nodes, 80) and write to the sink. image mode: render HTML to a temp file, then exec /Applications/Google Chrome.app/Contents/MacOS/Google Chrome --headless --disable-gpu --window-size=1440,900 --virtual-time-budget=2000 --screenshot=<out> file://<temp> to produce a PNG. Reuse the auto-launch and project-resolution logic conceptually from scripts/wireframe-shot.sh but operate on the wireframe HTML directly (no need for syded since the renderer runs in-process). Register under root cmd. Mirror the existing ''syde design'' command structure for consistency.'
id: TSK-0091
kind: task
name: Add 'syde wireframe render' CLI command
objective: syde wireframe render <contract-slug> outputs the rendered wireframe HTML (or ASCII) for any screen contract, pipeable to a file or directly to a browser
plan_phase: phase_2
plan_ref: uiml-wireframe-render
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-syde-wireframe-render-cli-command-cs0x
task_status: completed
updated_at: "2026-04-15T03:06:45Z"
---
