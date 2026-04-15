---
id: REQ-0298
kind: requirement
name: Session Context Invocation
slug: session-context-invocation-wbh3
relationships:
    - target: session-context-74s7
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde context, the syde CLI shall print the full architecture snapshot including project metadata, entity counts, key decisions, top-level components, and active plans.
req_type: interface
priority: must
verification: integration test invoking syde context and inspecting each section
source: manual
source_ref: contract:session-context-74s7
requirement_status: active
rationale: Session context is auto-loaded at session start and must be comprehensive enough to orient an agent.
---
