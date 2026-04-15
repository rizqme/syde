---
contract_kind: cli
description: Create a new entity of any kind with all its kind-specific fields.
id: CON-0008
input: syde add <kind> <name> [flags]
input_parameters:
    - description: positional, required. One of system|component|contract|concept|flow|decision|plan|task|design|requirement
      path: kind
      type: string
    - description: positional, required. Human-readable entity name
      path: name
      type: string
    - description: short explanation
      path: --description
      type: string
    - description: why it exists
      path: --purpose
      type: string
    - description: repeatable tag label
      path: --tag
      type: '[]string'
    - description: repeatable concrete source file path (no wildcards)
      path: --file
      type: '[]string'
    - description: markdown body content
      path: --body
      type: string
    - description: repeatable informal note
      path: --note
      type: '[]string'
    - description: repeatable relationship in 'target-slug:type' form
      path: --add-rel
      type: '[]string'
    - description: 'component: what it does (required for component)'
      path: --responsibility
      type: string
    - description: 'component: repeatable capability (≥1 required for component)'
      path: --capability
      type: '[]string'
    - description: 'component: what it does NOT do'
      path: --boundaries
      type: string
    - description: 'contract: rest|cli|event|rpc|graphql|websocket|pubsub'
      path: --contract-kind
      type: string
    - description: 'contract: sync|async|request-response|pub-sub|streaming'
      path: --interaction-pattern
      type: string
    - description: 'contract: invocation signature (required)'
      path: --input
      type: string
    - description: 'contract: ''path|type|description'' (≥1 required)'
      path: --input-parameter
      type: '[]string'
    - description: 'contract: output signature (required)'
      path: --output
      type: string
    - description: 'contract: ''path|type|description'' (≥1 required)'
      path: --output-parameter
      type: '[]string'
    - description: 'concept: domain meaning'
      path: --meaning
      type: string
    - description: 'concept: rules that must always hold'
      path: --invariants
      type: string
interaction_pattern: request-response
kind: contract
name: Add Entity
output: exit 0 on success; prints created ID and file path
output_parameters:
    - description: allocated counter-based ID (PFX-NNNN)
      path: id
      type: string
    - description: absolute path to the new markdown file
      path: file_path
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: storage-engine
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: add-entity-jbmc
updated_at: "2026-04-14T03:27:03Z"
---
