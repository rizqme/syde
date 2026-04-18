---
id: REQ-0163
kind: requirement
name: Scan Helpers Flag Infrastructure Dirs Exempt
slug: scan-helpers-flag-infrastructure-dirs-exempt-g02e
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
updated_at: "2026-04-18T09:38:03Z"
statement: When preparing a coverage report, the scan helpers shall flag known infrastructure directories such as migrations, deploy, and k8s as exempt from component coverage requirements.
req_type: functional
priority: should
verification: inspection of exempt-dir handling in coverage.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Infrastructure directories do not correspond to design components.
verified_against:
    scan-helpers-legacy-sa6d:
        hash: c9b19fd18480f13b89908a38caa95ff8ea66c15c73dc6cf6e551e414e71ce3ba
        at: "2026-04-18T09:38:03Z"
---
