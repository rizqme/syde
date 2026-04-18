---
id: REQ-0380
kind: requirement
name: Requirement creation shall detect similar requirements
slug: requirement-creation-shall-detect-similar-requirements-l1a6
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:58Z"
statement: When syde add requirement is invoked, the CLI shall compute TF-IDF similarity against existing active requirements and warn the user for any match above 50 percent.
req_type: functional
priority: must
verification: syde add requirement prints overlap warnings for similar names.
source: plan
requirement_status: active
rationale: Overlap detection prevents redundant intent statements.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:58Z"
---
