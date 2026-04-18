---
id: REQ-0489
kind: requirement
name: Design model shall contain exactly two systems named syde and syded
slug: design-model-shall-contain-exactly-two-systems-named-syde-and-syded-f2q8
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:55:14Z"
statement: While the syde design model is in the standalone-systems configuration, the syde CLI shall persist exactly two system entities named syde and syded corresponding to the two process binaries produced by the repository build.
req_type: functional
priority: must
verification: syde list system returns exactly two entries with names syde and syded
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: One system per standalone process eliminates overlap between syde and syde-cli-2478.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:55:14Z"
---
