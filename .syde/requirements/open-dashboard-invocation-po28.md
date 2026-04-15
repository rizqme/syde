---
id: REQ-0287
kind: requirement
name: Open Dashboard Invocation
slug: open-dashboard-invocation-po28
relationships:
    - target: open-dashboard-cupm
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde open, the syde CLI shall start the syded daemon in the background and open the browser to the project's dashboard URL.
req_type: interface
priority: must
verification: integration test invoking syde open and inspecting syded process plus opened URL
source: manual
source_ref: contract:open-dashboard-cupm
requirement_status: active
rationale: The open command is the one-shot entry point into the visual dashboard.
---
