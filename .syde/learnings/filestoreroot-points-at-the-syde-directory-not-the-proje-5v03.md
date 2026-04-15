---
category: gotcha
confidence: medium
description: FileStore.Root points at the .syde/ directory, NOT the project root. Tree paths are stored relative to the project root. Any os.Open / exec.Command that needs to resolve a tracked file must join against filepath.Dir(FS.Root), not FS.Root itself. SearchCode and ByFile --content both got bitten by this — added projectRoot() helper on Engine to make it explicit.
discovered_at: "2026-04-14T09:36:12Z"
entity_refs:
    - query-engine
id: LRN-0010
kind: learning
name: FileStore.Root points at the .syde/ directory, NOT the proje
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: filestoreroot-points-at-the-syde-directory-not-the-proje-5v03
source: session-observation
updated_at: "2026-04-14T09:36:12Z"
---
