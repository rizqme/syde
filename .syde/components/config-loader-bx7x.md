---
boundaries: Does NOT resolve entities. Does NOT touch BadgerDB.
capabilities:
    - YAML (un)marshal of Config struct
    - Default config for new projects
    - component_paths glob → component slug mapping used by constraints check
description: 'syde.yaml loader: project metadata, component file mappings, and tree-ignore patterns.'
files:
    - internal/config/config.go
id: COM-0009
kind: component
name: Config Loader
purpose: Own the syde.yaml project config shape and IO
relationships:
    - target: syde-cli
      type: belongs_to
responsibility: Load/save/default the project config (name, version, component_paths, tree_ignore)
slug: config-loader-bx7x
updated_at: "2026-04-14T03:35:54Z"
---
