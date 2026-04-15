---
id: REQ-0300
kind: requirement
name: Set Tree Summary Stdin Input
slug: set-tree-summary-stdin-input-ibav
relationships:
    - target: set-tree-summary-2vdt
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When syde tree summarize is invoked with --summary set to a dash, the syde CLI shall read the summary text from standard input.
req_type: interface
priority: must
verification: integration test piping summary text into syde tree summarize --summary -
source: manual
source_ref: contract:set-tree-summary-2vdt
requirement_status: active
rationale: Stdin input supports long summaries that would be awkward on the command line.
---
