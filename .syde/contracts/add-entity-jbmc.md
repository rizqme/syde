---
id: CON-0008
kind: contract
name: Add Entity
slug: add-entity-jbmc
description: Create a new entity of any kind with all its kind-specific fields.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: storage-engine
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde add <kind> <name> [flags]
input_parameters:
    - path: kind
      type: string
      description: positional, required. One of system|component|contract|concept|flow|decision|plan|task|design|requirement
    - path: name
      type: string
      description: positional, required. Human-readable entity name
    - path: --description
      type: string
      description: short explanation
    - path: --purpose
      type: string
      description: why it exists
    - path: --tag
      type: '[]string'
      description: repeatable tag label
    - path: --file
      type: '[]string'
      description: repeatable concrete source file path (no wildcards)
    - path: --body
      type: string
      description: markdown body content
    - path: --note
      type: '[]string'
      description: repeatable informal note
    - path: --add-rel
      type: '[]string'
      description: repeatable relationship in 'target-slug:type' form
    - path: --responsibility
      type: string
      description: 'component: what it does (required for component)'
    - path: --capability
      type: '[]string'
      description: 'component: repeatable capability (≥1 required for component)'
    - path: --boundaries
      type: string
      description: 'component: what it does NOT do'
    - path: --contract-kind
      type: string
      description: 'contract: rest|cli|event|rpc|graphql|websocket|pubsub'
    - path: --interaction-pattern
      type: string
      description: 'contract: sync|async|request-response|pub-sub|streaming'
    - path: --input
      type: string
      description: 'contract: invocation signature (required)'
    - path: --input-parameter
      type: '[]string'
      description: 'contract: ''path|type|description'' (≥1 required)'
    - path: --output
      type: string
      description: 'contract: output signature (required)'
    - path: --output-parameter
      type: '[]string'
      description: 'contract: ''path|type|description'' (≥1 required)'
    - path: --meaning
      type: string
      description: 'concept: domain meaning'
    - path: --invariants
      type: string
      description: 'concept: rules that must always hold'
output: exit 0 on success; prints created ID and file path
output_parameters:
    - path: id
      type: string
      description: allocated counter-based ID (PFX-NNNN)
    - path: file_path
      type: string
      description: absolute path to the new markdown file
---
