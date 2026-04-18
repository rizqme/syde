---
id: REQ-0407
kind: requirement
name: Smoke test single severity gate
slug: smoke-test-single-severity-gate-t5n0
relationships: []
updated_at: '2026-04-17T10:58:20Z'
statement: The syde audit engine shall emit findings at a single blocking severity level without any warning tier.
req_type: constraint
priority: must
verification: manual
source: manual
requirement_status: obsolete
rationale: Smoke test of overlap gate
obsolete_reason: Smoke test artefact from end-to-end gate verification; not a real requirement
audited_overlaps:
- slug: audit-shall-emit-a-single-severity-level-without-any-non-blocking-tier-baeq
  distinction: smoke test asserts the gate accepts --audited slug:reason and proceeds — distinct from the production rule because this is a manual one-off probe
---
