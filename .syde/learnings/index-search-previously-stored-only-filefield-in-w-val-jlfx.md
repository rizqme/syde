---
category: gotcha
confidence: medium
description: 'Index Search() previously stored only {file,field} in w: values, so every hit had empty entity identity. WordRef now carries full FileRef+field+word and a scanWordPrefix iterator handles the new value shape separately from scanPrefix (which still unmarshals FileRef for e:/s:/t: keys).'
discovered_at: "2026-04-14T08:29:48Z"
entity_refs:
    - storage-engine
id: LRN-0005
kind: learning
name: 'Index Search() previously stored only {file,field} in w: val'
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: index-search-previously-stored-only-filefield-in-w-val-jlfx
source: session-observation
updated_at: "2026-04-14T08:29:48Z"
---
