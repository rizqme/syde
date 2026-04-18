---
id: SYS-0003
kind: system
name: syde
slug: syde-5tdt
description: Text-first software design CLI + Claude Code skill + dashboard. Two standalone binaries (syde, syded) sharing one markdown-backed entity store.
purpose: Let AI agents and humans co-own a living architecture model in version control — plans, tasks, components, contracts, concepts — without leaving the terminal or the editor.
updated_at: "2026-04-18T10:04:41Z"
context: Claude Code and similar agents have no persistent architectural memory between sessions. syde fills that gap with a file-native model that lives in the repo, stays human-readable, and is enforced by a mandatory skill workflow.
scope: 'In: CLI for entity CRUD, plans/tasks, validation, file summary tree, dashboard. Out: runtime observability, code generation, LLM calls from syde itself.'
design_principles: Markdown files are source of truth. BadgerDB is a rebuildable cache, never authoritative. Every agent workflow is enforced by SKILL.md hooks. No LLM calls from syde binary — agents drive all summarization.
quality_goals: Fast session startup (architecture auto-loaded in <1s). Zero-drift guarantee via tree+validator. Idempotent CLI (every command safe to re-run).
assumptions: Single developer or small team per project. Projects are git repos. Agents have access to a syde binary in PATH.
---
