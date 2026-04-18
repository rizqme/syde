---
id: TSK-0250
kind: task
name: Implement syde requirement verify CLI command
slug: implement-syde-requirement-verify-cli-command-41la
updated_at: '2026-04-18T08:14:28Z'
task_status: completed
priority: high
objective: syde requirement verify <slug> snapshots SHA-256 of every file in each refining component into req.verified_against and saves
details: 'Create new file internal/cli/requirement.go with cobra subcommand. Resolve req via store.GetEntity. Walk Relationships of type RelRefines. For each, resolve target component, compute sha256 of each file, set verified_against[component-canonical-slug] = {hash: hex, at: time.Now().UTC().Format(RFC3339)}. Save via FileStore. Exit 1 with error if any file unreadable. Register under requirementCmd parent (new) under rootCmd.'
acceptance: syde requirement verify <existing-active-req-slug> succeeds; the req YAML now contains verified_against with one entry per refining component
affected_entities:
- cli-commands-hpjb
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_3
created_at: '2026-04-18T08:00:26Z'
completed_at: '2026-04-18T08:14:28Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: syde-requirement-verify-shall-snapshot-sha-256-hashes-for-every-file-in-each-refining-component-3p42
---
