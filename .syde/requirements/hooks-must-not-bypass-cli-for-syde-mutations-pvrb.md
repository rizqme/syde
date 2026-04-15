---
id: REQ-0098
kind: requirement
name: Hooks Must Not Bypass CLI For Syde Mutations
slug: hooks-must-not-bypass-cli-for-syde-mutations-pvrb
relationships:
    - target: skill-7fmf
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:13Z"
statement: The skill hooks shall not modify any file under .syde/ except through the syde CLI.
req_type: constraint
priority: must
verification: code review of hooks.json and hook scripts
source: manual
source_ref: concept:skill-7fmf
requirement_status: active
rationale: Going through the CLI keeps the index, tree, and file-store consistent.
---
