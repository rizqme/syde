---
id: REQ-0340
kind: requirement
name: Plan changes shall list their implementing tasks
slug: plan-changes-shall-list-their-implementing-tasks-gb2a
description: Plan Changes traceability requirement linking each change entry to implementing tasks.
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:42Z"
statement: When the syde plan model declares a Deleted, Extended, or NewChange entry, the entry shall carry a non-empty tasks list naming one or more tasks in the same plan that implement the change.
req_type: constraint
priority: must
verification: syde plan check on a plan whose changes contain an entry with an empty tasks list reports the orphan change as ERROR and exits non-zero.
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: A change without a claiming task is dead intent. Explicit mapping makes the linkage canonical and auditable.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:42Z"
---
