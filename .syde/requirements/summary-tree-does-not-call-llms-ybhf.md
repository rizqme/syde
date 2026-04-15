---
id: REQ-0194
kind: requirement
name: Summary Tree Does Not Call LLMs
slug: summary-tree-does-not-call-llms-ybhf
relationships:
    - target: summary-tree-fq6u
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:04Z"
statement: The summary tree shall not call any LLM to generate summaries and shall require summaries to be written by a human or agent via CLI.
req_type: constraint
priority: must
verification: code review of internal/tree for LLM client imports
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Summaries are intentional design knowledge and must not be auto-hallucinated.
---
