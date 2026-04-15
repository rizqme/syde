---
alternatives_considered: Optional LLM integration for auto-summaries; built-in OpenAI/Anthropic client
category: api
consequences: 'All tree summarization must be driven by the calling agent. Skill and hooks are the contract: syde provides the storage surface, agents provide the content.'
description: syde never calls an LLM — agents drive all summarization via CLI.
id: DEC-0002
kind: decision
name: No LLM Calls From syde Binary
rationale: Keeps syde deterministic, offline-capable, and free of API-key handling. Decouples the storage tool from any specific LLM provider. Makes the CLI safe to use in CI and automation.
relationships:
    - target: syde
      type: applies_to
    - target: summary-tree
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: no-llm-calls-from-syde-binary-fha5
statement: The syde binary must never call an LLM directly. All summarization, planning, and description writing is done by the invoking agent (Claude Code, Cursor, or a human) via CLI commands.
tradeoffs: Agents must do more work per session. Summaries may be inconsistent across agents. syde cannot auto-fill missing fields — that is agent responsibility.
updated_at: "2026-04-14T03:27:02Z"
---
