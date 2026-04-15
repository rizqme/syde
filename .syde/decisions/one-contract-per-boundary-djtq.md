---
alternatives_considered: One 'API' contract per service with free-form body (opaque to tooling)
category: api
consequences: Validator enforces input/input_parameters/output/output_parameters as required. CLI contracts must enumerate every flag as an input parameter. Schema contracts split per-prefix / per-table.
description: Every endpoint, command, event, key prefix, or table is its own contract entity.
id: DEC-0004
kind: decision
name: One Contract Per Boundary
rationale: Fine-grained contracts allow per-endpoint impact analysis, parameter documentation, and drift detection. The dashboard and constraints check rely on this granularity.
relationships:
    - target: syde
      type: applies_to
    - target: entity-model
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: one-contract-per-boundary-djtq
statement: Every API endpoint, CLI command, event, WebSocket message, RPC call, BadgerDB key prefix, and SQL table is modeled as its own contract entity. Umbrella contracts covering many endpoints are rejected.
tradeoffs: More entities to manage. Bulk create via shell scripts mitigates manual overhead.
updated_at: "2026-04-14T03:27:03Z"
---
