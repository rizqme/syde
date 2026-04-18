---
id: REQ-0198
kind: requirement
name: UIML Parser Reports Parse Errors With Line Numbers
slug: uiml-parser-reports-parse-errors-with-line-numbers-7nbz
relationships:
    - target: uiml-parser-sjdk
      type: refines
updated_at: "2026-04-18T09:37:40Z"
statement: If the UIML parser encounters a mismatched close tag or unknown tag, then the UIML parser shall record a ParseError with the line number of the offending token.
req_type: functional
priority: must
verification: unit test asserting line-numbered ParseError output
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Authors need precise line feedback when editing UIML bodies.
verified_against:
    uiml-parser-sjdk:
        hash: 4f1d204aa9053a09c0ad957ac1e2f841a9b380e4a2e80811c18e737a3736d44c
        at: "2026-04-18T09:37:40Z"
---
