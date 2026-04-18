---
id: REQ-0300
kind: requirement
name: Set Tree Summary Stdin Input
slug: set-tree-summary-stdin-input-ibav
relationships:
    - target: set-tree-summary-2vdt
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:46Z"
statement: When syde tree summarize is invoked with --summary set to a dash, the syde CLI shall read the summary text from standard input.
req_type: interface
priority: must
verification: integration test piping summary text into syde tree summarize --summary -
source: manual
source_ref: contract:set-tree-summary-2vdt
requirement_status: active
rationale: Stdin input supports long summaries that would be awkward on the command line.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:46Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:46Z"
---
