---
acceptance: go build clean; a concept spec 'id|primary key uuid' round-trips with Name=id, Description='primary key uuid', Refs=nil. YAML files on disk with stale 'type:' keys still load (yaml.v3 is lenient) and drop the key on next save.
affected_entities:
    - entity-model
affected_files:
    - internal/model/entity.go
    - internal/model/validation.go
completed_at: "2026-04-14T10:44:17Z"
created_at: "2026-04-14T10:40:01Z"
details: 'internal/model/entity.go: delete Type field from ConceptAttribute struct. ParseConceptAttribute reworked: parts[0]=name, parts[1]=description, parts[2]=refs (comma-separated). internal/model/validation.go: drop the ''attribute type required'' ERROR — keep name required. Rebuild confirms any lingering references fail to compile.'
id: TSK-0071
kind: task
name: Drop Type from ConceptAttribute model and parser
objective: ConceptAttribute model and parser carry only name + description + refs; Type is removed end-to-end
plan_phase: phase_3
plan_ref: erd-polish
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: drop-type-from-conceptattribute-model-and-parser-6c6x
task_status: completed
updated_at: "2026-04-14T10:44:17Z"
---
