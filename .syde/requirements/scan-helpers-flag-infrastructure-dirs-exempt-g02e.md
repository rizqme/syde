---
id: REQ-0163
kind: requirement
name: Scan Helpers Flag Infrastructure Dirs Exempt
slug: scan-helpers-flag-infrastructure-dirs-exempt-g02e
relationships:
    - target: scan-helpers-legacy-sa6d
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:28Z"
statement: When preparing a coverage report, the scan helpers shall flag known infrastructure directories such as migrations, deploy, and k8s as exempt from component coverage requirements.
req_type: functional
priority: should
verification: inspection of exempt-dir handling in coverage.go
source: manual
source_ref: component:scan-helpers-legacy-sa6d
requirement_status: active
rationale: Infrastructure directories do not correspond to design components.
---
