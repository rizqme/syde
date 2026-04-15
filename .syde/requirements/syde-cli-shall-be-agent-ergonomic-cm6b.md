---
id: REQ-0013
kind: requirement
name: syde CLI shall be agent-ergonomic
slug: syde-cli-shall-be-agent-ergonomic-cm6b
relationships:
    - target: syde-cli-2478
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syde CLI shall expose every entity CRUD, query, plan, task, tree, and validation operation as a single subcommand without requiring direct markdown edits.
req_type: usability
priority: must
verification: audit walking contract entities filtered to contract_kind=cli
source: manual
source_ref: system:syde-cli-2478:scope
requirement_status: active
rationale: Agents work via CLI; any operation that requires hand-editing markdown is a workflow gap.
---
