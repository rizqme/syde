---
id: REQ-0093
kind: requirement
name: CLI Commands Does Not Call LLMs
slug: cli-commands-does-not-call-llms-xci3
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:04Z"
statement: The syde CLI shall not invoke any LLM API during command execution.
req_type: constraint
priority: must
verification: code review of internal/cli for LLM client imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: syde stays deterministic and offline-capable so agents remain the sole LLM consumers.
---
