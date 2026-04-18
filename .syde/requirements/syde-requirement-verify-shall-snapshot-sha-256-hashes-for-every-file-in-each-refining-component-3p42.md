---
id: REQ-0482
kind: requirement
name: syde requirement verify shall snapshot SHA-256 hashes for every file in each refining component
slug: syde-requirement-verify-shall-snapshot-sha-256-hashes-for-every-file-in-each-refining-component-3p42
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:24Z"
statement: When the user runs syde requirement verify against a requirement slug, the syde requirement entity shall snapshot the SHA-256 content hash of every file in each refining component into verified_against.
req_type: functional
priority: must
verification: Running syde requirement verify <slug> updates verified_against entries for each refining component to the current SHA-256 hashes and timestamp
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Verify is idempotent; re-running refreshes timestamps; failure to verify any file aborts the snapshot to avoid partial state.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:24Z"
---
