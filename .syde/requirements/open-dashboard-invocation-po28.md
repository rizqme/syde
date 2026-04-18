---
id: REQ-0287
kind: requirement
name: Open Dashboard Invocation
slug: open-dashboard-invocation-po28
relationships:
    - target: open-dashboard-cupm
      type: refines
    - target: http-api-afos
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:56Z"
statement: When the user runs syde open, the syde CLI shall start the syded daemon in the background and open the browser to the project's dashboard URL.
req_type: interface
priority: must
verification: integration test invoking syde open and inspecting syded process plus opened URL
source: manual
source_ref: contract:open-dashboard-cupm
requirement_status: active
rationale: The open command is the one-shot entry point into the visual dashboard.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:56Z"
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:56Z"
---
