---
id: REQ-0198
kind: requirement
name: UIML Parser Reports Parse Errors With Line Numbers
slug: uiml-parser-reports-parse-errors-with-line-numbers-7nbz
relationships:
    - target: uiml-parser-sjdk
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:57:17Z"
statement: If the UIML parser encounters a mismatched close tag or unknown tag, then the UIML parser shall record a ParseError with the line number of the offending token.
req_type: functional
priority: must
verification: unit test asserting line-numbered ParseError output
source: manual
source_ref: component:uiml-parser-sjdk
requirement_status: active
rationale: Authors need precise line feedback when editing UIML bodies.
---
