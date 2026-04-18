---
id: REQ-0002
kind: requirement
name: syde binary shall not call LLMs
slug: syde-binary-shall-not-call-llms-6tvj
description: syde must never make LLM API calls; agents drive all summarization.
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:36:58Z"
statement: The syde and syded binaries shall not invoke any LLM API or require API-key configuration; all summarization, planning, and description authoring shall be produced by the invoking agent via CLI commands.
req_type: constraint
priority: must
verification: rg over internal/ and cmd/ for any anthropic/openai/api-key reference returns zero matches
source: manual
source_ref: decision:DEC-0002
requirement_status: active
rationale: Keeps syde deterministic, offline-capable, and free of API-key handling. Decouples the storage tool from any specific LLM provider. Makes the CLI safe to use in CI and automation.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:58Z"
---
