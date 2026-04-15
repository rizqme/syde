---
id: REQ-0323
kind: requirement
name: Validate Model Reports Errors
slug: validate-model-reports-errors-jhp2
relationships:
    - target: validate-model-tjzs
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When syde validate succeeds, the syde CLI shall return errors as a list of fatal issues including missing required fields, unknown targets, cycles, orphan files, and contract schema gaps.
req_type: interface
priority: must
verification: integration test invoking syde validate with a model containing each error class
source: manual
source_ref: contract:validate-model-tjzs
requirement_status: active
rationale: Clear error categorization lets operators fix issues without reading Go source.
---
