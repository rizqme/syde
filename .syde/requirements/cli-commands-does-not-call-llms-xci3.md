---
id: REQ-0093
kind: requirement
name: CLI Commands Does Not Call LLMs
slug: cli-commands-does-not-call-llms-xci3
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:31Z"
statement: The syde CLI shall not invoke any LLM API during command execution.
req_type: constraint
priority: must
verification: code review of internal/cli for LLM client imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: syde stays deterministic and offline-capable so agents remain the sole LLM consumers.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:31Z"
---
