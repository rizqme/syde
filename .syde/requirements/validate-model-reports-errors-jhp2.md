---
id: REQ-0323
kind: requirement
name: Validate Model Reports Errors
slug: validate-model-reports-errors-jhp2
relationships:
    - target: validate-model-tjzs
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: entity-model-f28o
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:31Z"
statement: When syde validate succeeds, the syde CLI shall return errors as a list of fatal issues including missing required fields, unknown targets, cycles, orphan files, and contract schema gaps.
req_type: interface
priority: must
verification: integration test invoking syde validate with a model containing each error class
source: manual
source_ref: contract:validate-model-tjzs
requirement_status: active
rationale: Clear error categorization lets operators fix issues without reading Go source.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:31Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:31Z"
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:31Z"
---
