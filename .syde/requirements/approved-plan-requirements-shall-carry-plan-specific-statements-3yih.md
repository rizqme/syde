---
id: REQ-0398
kind: requirement
name: Approved-plan requirements shall carry plan-specific statements
slug: approved-plan-requirements-shall-carry-plan-specific-statements-3yih
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:38Z"
statement: When the syde plan approval workflow creates the approved-plan requirement, the statement shall be a plan-specific EARS statement derived from the plan's objective field rather than the plan's name.
req_type: functional
priority: must
verification: syde query --search 'Approved plan' shows plan-specific statements with pairwise TF-IDF similarity below 0.6
source: plan
requirement_status: active
rationale: Generic approval statements flood the overlap audit with false positives.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:38Z"
---
