---
id: REQ-0070
kind: requirement
name: CLI Commands Support Entity CRUD
slug: cli-commands-support-entity-crud-y7au
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:43Z"
statement: The syde CLI shall support entity create, read, update, list, and remove operations via add, get, list, update, and remove subcommands.
req_type: functional
priority: must
verification: integration test invoking syde add/get/list/update/remove
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: CRUD coverage is the baseline capability for managing the design model.
---
