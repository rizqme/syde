---
id: REQ-0194
kind: requirement
name: Summary Tree Does Not Call LLMs
slug: summary-tree-does-not-call-llms-ybhf
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:58Z"
statement: The summary tree shall not call any LLM to generate summaries and shall require summaries to be written by a human or agent via CLI.
req_type: constraint
priority: must
verification: code review of internal/tree for LLM client imports
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Summaries are intentional design knowledge and must not be auto-hallucinated.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:58Z"
---
