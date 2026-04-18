---
id: REQ-0298
kind: requirement
name: Session Context Invocation
slug: session-context-invocation-wbh3
relationships:
    - target: session-context-74s7
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:55Z"
statement: When the user runs syde context, the syde CLI shall print the full architecture snapshot including project metadata, entity counts, key decisions, top-level components, and active plans.
req_type: interface
priority: must
verification: integration test invoking syde context and inspecting each section
source: manual
source_ref: contract:session-context-74s7
requirement_status: active
rationale: Session context is auto-loaded at session start and must be comprehensive enough to orient an agent.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:55Z"
---
