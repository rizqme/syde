---
id: REQ-0275
kind: requirement
name: Ignore Tree Node Invocation
slug: ignore-tree-node-invocation-0kg2
relationships:
    - target: ignore-tree-node-m460
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde tree ignore <path>, the syde CLI shall flag the named node as ignored so the orphan validator exempts it.
req_type: interface
priority: must
verification: integration test invoking syde tree ignore and running the orphan validator
source: manual
source_ref: contract:ignore-tree-node-m460
requirement_status: active
rationale: Ignore lists let operators silence intentional gaps without breaking validation.
---
