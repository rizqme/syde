---
id: REQ-0398
kind: requirement
name: Approved-plan requirements shall carry plan-specific statements
slug: approved-plan-requirements-shall-carry-plan-specific-statements-3yih
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When the syde plan approval workflow creates the approved-plan requirement, the statement shall be a plan-specific EARS statement derived from the plan's objective field rather than the plan's name.
req_type: functional
priority: must
verification: syde query --search 'Approved plan' shows plan-specific statements with pairwise TF-IDF similarity below 0.6
source: plan
requirement_status: active
rationale: Generic approval statements flood the overlap audit with false positives.
---
