---
id: REQ-0004
kind: requirement
name: Every boundary shall be modeled as its own contract
slug: every-boundary-shall-be-modeled-as-its-own-contract-h7jq
description: Fine-grained contract entities per endpoint / command / event / key prefix / table.
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:47:01Z"
statement: The syde model shall represent every API endpoint, CLI command, event, WebSocket message, RPC call, BadgerDB key prefix, and SQL table as a separate contract entity.
req_type: constraint
priority: must
verification: audit walking contract entities
source: manual
source_ref: decision:DEC-0004
requirement_status: active
rationale: Fine-grained contracts enable per-endpoint impact analysis, parameter documentation, and drift detection. The dashboard and constraints check rely on this granularity.
---
