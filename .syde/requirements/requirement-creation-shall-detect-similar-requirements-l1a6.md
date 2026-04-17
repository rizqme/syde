---
id: REQ-0380
kind: requirement
name: Requirement creation shall detect similar requirements
slug: requirement-creation-shall-detect-similar-requirements-l1a6
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: When syde add requirement is invoked, the CLI shall compute TF-IDF similarity against existing active requirements and warn the user for any match above 50 percent.
req_type: functional
priority: must
verification: syde add requirement prints overlap warnings for similar names.
source: plan
requirement_status: active
rationale: Overlap detection prevents redundant intent statements.
---
