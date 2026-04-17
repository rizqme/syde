---
id: PLN-0011
kind: plan
name: Audit requirements for overlaps, merge duplicates, enforce semantic distinction at the harness level
slug: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
relationships:
    - target: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T11:08:17Z"
plan_status: completed
background: 'Two harness gaps motivate this plan. (1) TF-IDF surfaces candidate overlaps but the existing --audited flag takes only a slug, so acknowledgements are rubber stamps rather than semantic review. (2) When a new or extended requirement names a surface governed by a contract (a CLI command, REST endpoint, screen, event), nothing in the planning or post-plan audit forces the author to also touch the corresponding contract entity. On top of that, the audit layer is asymmetric: several checks live only in plan_authoring (planning-time) with no post-plan equivalent, so a rule evaded at planning stays evaded forever.'
objective: 'Overlap handling becomes a harness-enforced semantic review workflow with distinction rationale, merging, and gate-level blocking. Contract coverage becomes a first-class audit dimension: every requirement whose statement names a contract-governed surface must co-evolve with a contract entity, enforced at planning AND at sync check time. Every rule in plan_authoring has an equivalent finding in the post-plan audit chain so evasion at planning is caught at rest. The six Approved-plan requirements get plan-specific statements. syde sync check --strict continues to exit 0 after all changes, with the rubber-stamp and coverage rules now active.'
scope: 'In scope: (A) OVERLAP WORKFLOW — extend audited_overlaps with per-entry distinction rationale, block creation at >=0.6 unless every overlap is acknowledged with slug:reason or --force, audit errors on empty distinctions, PostToolUse hook injects the resolution reminder, Claude reviews existing pairs case-by-case; (B) CONTRACT COVERAGE — add a detector that extracts contract-surface mentions from requirement statements (CLI pattern like ''syde <subcommand>'', REST pattern like ''GET /api/...'', screen pattern like ''the dashboard shall render/display X'', event pattern like ''When <event>, the system shall ...''), emit a plan-authoring warning when a new/extended requirement''s surface is not covered by a new-or-extended contract in the same plan, emit a sync-check error when an active requirement''s surface has no corresponding active contract, walk existing 386 active requirements to add missing contracts; (C) SYMMETRY PRINCIPLE — enumerate every plan_authoring rule, for each one ensure an equivalent finding exists in the post-plan audit (i.e. if planning flags X on intent, sync check flags X on state), document the principle in SKILL.md, add a Go table-driven test that fails if a new plan_authoring rule lacks a post-plan counterpart. Out of scope: dropping TF-IDF for embeddings; contract-inference from free-form prose beyond the four named patterns; automated contract authoring (we surface the gap, Claude fills it).'
design: 'The plan unfolds in three workstreams that share infrastructure. OVERLAP WORKFLOW (phases 1-5) layers semantic-distinction capture on top of the existing TF-IDF engine: data model carries {slug, distinction}, CLI gate blocks and parses slug:reason, audit escalates empty distinctions to error, hook enforces resolution, Claude executes MERGE/RENAME/DISTINCT pair-by-pair. CONTRACT COVERAGE (phases 6-7) adds a new package internal/audit/surfaces.go with regex-based extractors for CLI/REST/screen/event patterns; plan_authoring gains a rule that, for each new or extended requirement change, extracts surfaces from the statement and checks the plan''s change diff for any contract (new or extended) whose input field matches the surface — emits warning on gaps; sync check gains the mirror rule that walks active requirements and errors when no active contract matches the surface. Claude then walks existing uncovered surfaces and adds contract entities under a separate subtask. SYMMETRY PRINCIPLE (phase 8) treats plan_authoring and sync-check audits as a pair: every rule in the former must have a semantically equivalent rule in the latter so evasion is impossible. Implementation: a registry interface lists plan_authoring findings by ID and expected post-plan finding ID; a Go test iterates the registry asserting every entry has an active post-plan counterpart. SKILL.md gains a ''Planning enforcement symmetry'' subsection. Throughout: two-step remove-then-add relationship edits; bare vs. full slug awareness on --remove-rel.'
source: manual
created_at: "2026-04-17T09:39:04Z"
approved_at: "2026-04-17T10:02:42Z"
completed_at: "2026-04-17T11:08:17Z"
phases:
    - id: phase_1
      name: Extend audited_overlaps data model with distinction rationale
      status: completed
      description: Change RequirementEntity.AuditedOverlaps from []string to []AuditedOverlap{Slug, Distinction}; keep YAML backward-readable
      objective: Every acknowledgement carries a non-empty semantic distinction rationale alongside the target slug
      changes: internal/model/requirement.go (or wherever AuditedOverlaps lives); YAML marshal/unmarshal for slug-only entries treated as distinction=''; new AuditedOverlap struct
      details: Add a struct with Slug and Distinction fields. UnmarshalYAML accepts either a plain string (legacy) — in that case Distinction stays empty, flagged by audit — or a full {slug, distinction} map. MarshalYAML always writes the full form. Run a one-time reindex so existing entries round-trip.
      tasks:
        - extend-auditedoverlap-struct-with-distinction-field
    - id: phase_2
      name: CLI accepts --audited slug:reason and blocks unacknowledged overlaps at 0.6
      status: completed
      description: syde add requirement and syde update learn to parse slug:reason; add fails non-zero when overlaps are unacknowledged unless --force
      objective: Authors cannot silently create overlapping requirements; every acknowledgement carries a documented distinction
      changes: internal/cli/add.go, internal/cli/update.go
      details: Parse --audited values as slug-or-slug-colon-reason. In add.go, after overlap detection, if any overlap is not covered by an --audited entry, print the overlap list plus a hint and return error. --force bypasses. In update.go, --audited accepts the same syntax and appends to AuditedOverlaps.
      tasks:
        - parse-slugreason-in-audited-and-block-unacknowledged-overlaps-at-add-time
    - id: phase_3
      name: Audit errors on rubber-stamped acknowledgements
      status: completed
      description: requirementOverlapFindings escalates acknowledgements with empty or trivially short distinction text to error
      objective: sync check flags rubber-stamp --audited entries so they cannot silence audit without real reasoning
      changes: internal/audit/requirements.go
      details: 'Add a check: for every audited_overlaps entry, require Distinction length >=20 characters and not identical to the other requirement''s rationale (cosine < 0.5). Emit ERROR on violations with message indicating the field to populate.'
      tasks:
        - audit-requires-non-empty-distinction-on-every-acknowledgement
    - id: phase_4
      name: PostToolUse hook blocks session on unresolved overlaps
      status: completed
      description: Add a hook that fires after syde add requirement and injects a system reminder when overlaps are printed
      objective: The harness forces Claude to resolve each overlap before continuing; no silent skip
      changes: skill/hooks.json, .claude/hooks/syde-hooks.json (re-installed by syde install-skill)
      details: Add a PostToolUse matcher on 'Bash' that greps its output for the overlap banner ('⚠ Similar to'). If present, emit a system reminder listing the overlaps and requiring the next tool call to be either syde update --audited slug:reason, syde remove (merge), or syde update statement= (rename). The hook is a small inline shell command stored in hooks.json.
      tasks:
        - posttooluse-hook-on-syde-add-requirement-injects-overlap-resolution-reminder
    - id: phase_5
      name: Claude reviews the 70 existing overlap pairs
      status: completed
      description: Pair-by-pair semantic review of every current >=60 percent overlap, executing MERGE/RENAME/DISTINCT for each
      objective: Every current overlap pair above 60 percent carries either a survivor requirement (merge), a rewritten statement (rename), or a documented distinction (distinct)
      changes: Dozens of requirement entities created, superseded, or updated; dozens of refines/derives_from/references rels repointed; audited_overlaps entries gain distinction text
      details: 'Iterate cmd/listoverlaps output starting with highest similarity. For each pair, run syde query on both, classify, then execute. MERGE path: syde add the survivor, for every inbound rel on either original run syde update remove-rel then add-rel against the survivor, syde update both originals --requirement-status superseded --superseded-by survivor. RENAME path: syde update <loser> --statement ''<new EARS statement>''. DISTINCT path: syde update <a> --audited <b-slug>:''<distinction>'' and vice-versa. Approved-plan requirements get plan-specific --statement rewrites sourced from each plan''s objective field.'
      tasks:
        - review-approved-plan-requirement-overlaps-and-rewrite-statements
        - review-high-overlap-template-pairs-80-percent-and-merge
        - review-medium-overlap-pairs-60-80-percent
    - id: phase_6
      name: Update SKILL docs and add regression test
      status: completed
      description: Document the MERGE/RENAME/DISTINCT workflow in SKILL.md and add a Go test asserting the CLI gate behavior
      objective: Future Claude sessions know the workflow; CI/test asserts the gate continues to fire
      changes: skill/SKILL.md, skill/references/entity-spec.md, internal/cli/add_test.go or internal/audit/requirements_test.go
      details: SKILL.md gains a 'Requirement overlap resolution' subsection under the Requirement rules explaining MERGE vs RENAME vs DISTINCT and when each applies; entity-spec.md mentions the slug:reason syntax on audited_overlaps. Add a table-driven Go test covering (a) high-overlap add without --audited returns error, (b) with --audited slug:reason succeeds, (c) audit errors on empty distinction.
      tasks:
        - update-skillmd-and-entity-spec-with-overlap-resolution-workflow
        - regression-test-for-cli-gate-and-audit-distinction-rule
    - id: phase_7
      name: Verify
      status: completed
      description: Build, reinstall skill, run sync check strict, confirm gate behavior end-to-end
      objective: syde sync check --strict exits 0; a crafted high-overlap syde add requirement call fails without --audited and succeeds with it
      changes: No entity changes; reindex run; skill reinstalled
      details: 'go build, syde install-skill --all, syde reindex, syde sync check --strict, then a manual smoke: syde add requirement with a statement guaranteed to overlap an existing one and verify non-zero exit; rerun with --audited slug:reason and verify success.'
      tasks:
        - build-reinstall-skill-smoke-test-end-to-end-gate
    - id: phase_8
      name: Detect contract surfaces in requirement statements
      status: completed
      description: Add a detector that extracts CLI, REST, screen, and event surface mentions from requirement.statement
      objective: Given any requirement, enumerate the contract surfaces it implies — CLI invocation patterns, REST paths, screen routes, event topics
      changes: internal/audit/surfaces.go with regex-based extractors; Go test fixtures covering each pattern
      details: 'Regex patterns: CLI = /syde ([a-z][a-z0-9-]+( [a-z0-9-]+)?)/; REST = /(GET|POST|PUT|DELETE|PATCH) (\/[a-zA-Z0-9/_:-]+)/; screen = /(dashboard|SPA|UI) shall (render|display|show) ([A-Za-z ]+)/; event = /When ([a-z][a-z0-9 _-]+), the .+ shall/. Each extractor returns a normalised surface identifier that can be matched against contract.input. Unit tests cover positive and negative cases for each pattern.'
      tasks:
        - add-surface-detector-with-clirestscreenevent-patterns
    - id: phase_9
      name: Contract coverage at planning and post-plan
      status: completed
      description: plan_authoring warns when a new or extended req names a surface uncovered by a contract change in the same plan; sync check errors when an active req's surface has no active contract
      objective: Every requirement whose statement names a contract-governed surface must co-evolve with a contract entity
      changes: internal/audit/plan_authoring.go adds surface-coverage check on requirement lane new/extended entries; internal/audit/requirements.go or graph_rules.go adds post-plan surface-coverage check; also walks existing active requirements and lists uncovered surfaces for Claude to close
      details: 'Planning rule: for each new/extended requirement change, extract surfaces; check the plan''s change diff for any contract (new or extended) whose input field OR declared name matches the surface; emit WARN if gap. Post-plan rule: walks all active requirements, extracts surfaces, checks against all active contracts'' input field; emits ERROR if gap. The detector lives in internal/audit/surfaces.go (from the previous phase). After rules are in, Claude walks the current gaps and adds contract entities.'
      tasks:
        - plan-authoring-warns-on-uncovered-contract-surfaces-in-new-or-extended-requirements
        - sync-check-errors-on-active-requirements-with-uncovered-contract-surfaces
        - walk-existing-active-requirements-and-add-missing-contracts
    - id: phase_10
      name: Planning-post-plan symmetry audit and gap-fill
      status: completed
      description: Enumerate every plan_authoring rule, verify it has an equivalent post-plan rule, add missing counterparts, document the principle, add a Go test
      objective: Every planning-time enforcement has an equivalent post-plan enforcement so evasion at planning is caught at rest
      changes: internal/audit/plan_authoring.go and internal/audit/*.go gain paired finding IDs; internal/audit/symmetry_test.go asserts parity; skill/SKILL.md documents the principle
      details: Walk plan_authoring.go and list every Finding producer; for each, identify the semantically equivalent post-plan finding in requirements.go, graph_rules.go, or plan_completion.go. If none, add it. Encode the parity registry as a map[string]string (planning finding category/key -> post-plan finding category/key) in a new symmetry.go file. Add symmetry_test.go iterating the map and asserting every planning finding in the registry has a non-empty post-plan counterpart active in audit.Run. Document in SKILL.md under a new 'Planning enforcement symmetry' subsection.
      tasks:
        - build-parity-registry-and-symmetry-test
    - id: phase_11
      name: Flow coverage at planning and post-plan
      status: completed
      description: Enforce flow co-evolution when a plan touches contracts or user-facing requirements; strengthen the post-plan flow-contract rule
      objective: Every plan that introduces or extends a contract also introduces or extends a flow that exercises that contract; every active contract participates in at least one flow step; every user-facing requirement traces to at least one flow via the contracts it names
      changes: internal/audit/plan_authoring.go gains a flow-coverage rule for contract changes; internal/audit/graph_rules.go contractFlowFindings is verified (and tightened if needed) to error on any active contract without a flow step reference; Claude walks existing contracts without flows and adds flow entries
      details: 'Planning rule: for each new/extended contract change in the plan diff, check the plan''s flow lane for any flow (new or extended) whose steps reference the contract by slug. If none, emit WARN with finding key ''contract_flow_coverage''. Post-plan rule already exists (contractFlowFindings) — verify it errors rather than warns, that it walks flow steps not legacy narrative, and that it names the missing contract clearly. After the rule is in, Claude walks the current contract-without-flow gaps and either authors a new flow or adds a step to an existing flow.'
      tasks:
        - plan-authoring-warns-when-contract-changes-lack-flow-coverage
        - verify-and-tighten-contractflowfindings-post-plan-rule
        - walk-existing-contracts-without-flows-and-add-flow-coverage
    - id: phase_12
      name: Collapse audit severity to a single Finding level and remove --strict
      status: completed
      description: Replace SeverityError/SeverityWarning/SeverityHint with a single Finding severity; every finding blocks; remove the --strict flag so there is no non-strict mode
      objective: syde sync check, syde validate, syde plan complete, and syde plan check each exit non-zero on any audit finding, with no severity tiers and no --strict toggle
      changes: internal/audit/audit.go Severity type becomes a single constant Finding (rename from SeverityError); all internal/audit/*.go producers updated; internal/cli/sync_check.go drops --strict; internal/cli/validate.go (if it exists) drops --strict; skill docs updated to drop the word 'warning' from the workflow
      details: 'Step 1: rename SeverityError→Finding or introduce a new single constant and delete the other two. Keep string value ''finding'' so output reads ''FINDING: ...''. Step 2: grep every SeverityWarning and SeverityHint producer and change to Finding. Step 3: remove the Severity field branches in audit.go Report.Counts() and formatter paths; a Report simply has len(Findings). Step 4: remove --strict from sync_check.go; the stale-tree cascade always blocks. Step 5: update skill/SKILL.md and skill/references/* to replace ''WARN'' / ''warning'' / ''strict'' terminology with ''Finding'' / ''blocking''. Step 6: add a regression test asserting the enum has exactly one constant.'
      tasks:
        - collapse-severity-enum-to-single-finding-constant
        - reclassify-every-warning-and-hint-producer-as-finding
        - remove-strict-flag-from-synccheck-and-validate
        - update-skill-docs-and-hooks-for-strict-audit-vocabulary
        - regression-test-audit-severity-enum-has-exactly-one-value
changes:
    requirements:
        extended:
            - id: 1x86
              slug: audit-shall-error-on-contracts-not-in-any-flow-step
              what: Rule already exists; phase 11 task verifies/tightens it to ERROR severity on flow.Steps walk
              why: Avoid duplicating an existing requirement with the same intent
              field_changes:
                priority: must
              tasks:
                - verify-and-tighten-contractflowfindings-post-plan-rule
            - id: lpat
              slug: list-entities-accepts-kind-filter
              what: After the review pass, this overlap-era requirement is tagged overlap-reviewed
              why: Verifiable signal that the review pass ran
              field_changes:
                tags: overlap-reviewed
              tasks:
                - review-high-overlap-template-pairs-80-percent-and-merge
                - review-medium-overlap-pairs-60-80-percent
                - update-skillmd-and-entity-spec-with-overlap-resolution-workflow
            - id: az4f
              slug: each-flow-shall-represent-one-user-goal
              what: After walk-reqs pass, this requirement is tagged contract-coverage-reviewed
              why: Verifiable signal
              field_changes:
                tags: contract-coverage-reviewed
              tasks:
                - walk-existing-active-requirements-and-add-missing-contracts
            - id: ieup
              slug: list-entities-accepts-tag-filter
              what: After walk-flows pass, this requirement is tagged flow-coverage-reviewed
              why: Verifiable signal
              field_changes:
                tags: flow-coverage-reviewed
              tasks:
                - walk-existing-contracts-without-flows-and-add-flow-coverage
        new:
            - id: f79c
              name: Requirement creation shall block unacknowledged overlaps
              what: syde add requirement returns non-zero when any surfaced TF-IDF overlap above 0.6 is not acknowledged via --audited slug:reason unless --force is passed
              why: Overlap must surface at creation, not at sync check
              draft:
                priority: must
                rationale: Shift-left the overlap gate so authors resolve before the entity exists.
                req_type: functional
                source: plan
                statement: When syde add requirement detects one or more candidate overlaps above the audit threshold, the CLI shall refuse to create the entity unless every surfaced overlap is acknowledged with an --audited slug:reason entry or the author passes --force.
                verification: syde add requirement with a high-overlap statement exits non-zero without --audited and succeeds with slug:reason
              tasks:
                - parse-slugreason-in-audited-and-block-unacknowledged-overlaps-at-add-time
            - id: bq4b
              name: Harness hook shall block session on unresolved overlaps
              what: A PostToolUse hook greps syde add requirement stdout for the overlap banner and, if present, emits a system reminder listing the resolution paths
              why: Claude must not proceed until overlaps are semantically resolved
              draft:
                priority: must
                rationale: The skill infrastructure must enforce the semantic review, not just document it.
                req_type: functional
                source: plan
                statement: When syde add requirement prints an overlap banner, the installed Claude Code PostToolUse hook shall emit a system reminder naming the merge, rename, and distinct resolution paths before the session continues.
                verification: after syde install-skill --all, a syde add requirement call with a high-overlap statement triggers a visible system reminder in the session transcript
              tasks:
                - posttooluse-hook-on-syde-add-requirement-injects-overlap-resolution-reminder
            - id: ntkr
              name: Approved-plan requirements shall carry plan-specific statements
              what: Each approved-plan requirement's statement is rewritten in plan-specific EARS form instead of echoing the plan name
              why: Boilerplate statements produce 100 percent TF-IDF overlap between every approved plan
              draft:
                priority: must
                rationale: Generic approval statements flood the overlap audit with false positives.
                req_type: functional
                source: plan
                statement: When the syde plan approval workflow creates the approved-plan requirement, the statement shall be a plan-specific EARS statement derived from the plan's objective field rather than the plan's name.
                verification: syde query --search 'Approved plan' shows plan-specific statements with pairwise TF-IDF similarity below 0.6
              tasks:
                - review-approved-plan-requirement-overlaps-and-rewrite-statements
            - id: "0xba"
              name: Plan shall touch contracts when requirements name surfaces
              what: When a plan's diff introduces or extends a requirement whose statement names a contract-governed surface, the diff must also introduce or extend a contract covering that surface
              why: Catches requirement-contract co-evolution at planning time, not just at sync check
              draft:
                priority: must
                rationale: Shift-left the requirement-contract coverage gate.
                req_type: functional
                source: plan
                statement: When a syde plan's change diff introduces or extends a requirement whose statement names a CLI command, REST endpoint, screen, or event surface, the same plan's change diff shall introduce or extend a contract whose input covers that surface.
                verification: syde plan check warns on any plan whose requirement lane mentions a surface not covered by the plan's contract lane
              tasks:
                - plan-authoring-warns-on-uncovered-contract-surfaces-in-new-or-extended-requirements
            - id: duot
              name: Every planning enforcement shall have a post-plan equivalent
              what: Each rule in plan_authoring has a counterpart rule in the post-plan audit chain so evasion at planning is caught at rest
              why: An asymmetric audit lets any rule be bypassed by skipping plan_authoring
              draft:
                priority: must
                rationale: Symmetric audits prevent one-sided gates from being evaded.
                req_type: constraint
                source: plan
                statement: The syde audit engine shall maintain, for every finding category emitted by plan_authoring, an equivalent finding category emitted by the post-plan audit chain, so that an intent missed at planning is caught against actual state.
                verification: a Go table-driven test iterates the planning-post-plan parity registry and asserts both sides fire on the same crafted input
              tasks:
                - build-parity-registry-and-symmetry-test
            - id: 8kku
              name: Acknowledged requirement overlaps shall carry non-trivial distinction text
              what: requirementOverlapFindings emits an error when any audited_overlaps entry has an empty or trivially short distinction field
              why: Silence without reasoning is not acceptable resolution
              draft:
                priority: must
                rationale: Audited acknowledgements must document semantic distinction or be treated as unresolved.
                req_type: constraint
                source: plan
                statement: If an audited overlap entry on a requirement carries a distinction rationale shorter than 20 characters, then the syde audit engine shall report an error.
                verification: sync check reports errors for every acknowledgement whose distinction is empty or below 20 chars
              tasks:
                - audit-requires-non-empty-distinction-on-every-acknowledgement
            - id: z7lj
              name: Requirement statements mentioning CLI REST screen or event surfaces shall map to contracts
              what: Post-plan audit treats requirement prose that names an invocation surface as evidence a matching contract must exist
              why: Requirements describe intent; contracts describe the boundary; both must exist together
              draft:
                priority: must
                rationale: Prevents requirement-contract drift by forcing the surface to exist in both places.
                req_type: constraint
                source: plan
                statement: When an active requirement's prose names a CLI invocation, HTTP route, dashboard screen, or pub-sub topic, the syde design model shall hold an active contract whose input definition covers the surface.
                verification: zero sync check findings about requirement prose surfaces lacking contract coverage
              tasks:
                - sync-check-errors-on-active-requirements-with-uncovered-contract-surfaces
            - id: 6vdf
              name: Plan diffs introducing contracts shall also introduce matching flow steps
              what: plan_authoring emits a warning when a plan's contract lane gains entries not paired with a flow lane entry
              why: Contracts without flow coverage cannot be traced end-to-end; planning must catch the omission
              draft:
                priority: must
                rationale: Enforces co-evolution of boundary definitions and their traversal documents at planning time.
                req_type: functional
                source: plan
                statement: When a syde plan change diff adds or extends a contract entry, the same diff shall pair that entry with a flow lane item whose trajectory names the contract slug.
                verification: plan check on a diff that adds a contract without a paired flow lane item emits a warning; pairing clears it
              tasks:
                - plan-authoring-warns-when-contract-changes-lack-flow-coverage
            - id: f1nv
              name: Audit shall emit a single severity level without any non-blocking tier
              what: internal/audit defines one Finding severity; every producer blocks; there is no warning or hint category
              why: Project policy is strict; any severity hierarchy invites evasion via 'acceptable warnings'
              draft:
                priority: must
                rationale: A strict project cannot afford a non-blocking audit tier; findings either matter or they should not be emitted.
                req_type: constraint
                source: plan
                statement: The syde audit engine shall emit findings at a single blocking severity level, without any warning or hint tier, so that every finding causes the sync-check, validate, and plan-complete gates to exit non-zero.
                verification: grep 'SeverityWarning\|SeverityHint' in internal/audit returns zero results and a Go test asserts the audit severity enum has exactly one exported value
              tasks:
                - collapse-severity-enum-to-single-finding-constant
                - reclassify-every-warning-and-hint-producer-as-finding
                - regression-test-audit-severity-enum-has-exactly-one-value
            - id: hfuv
              name: Sync check validate and plan complete shall have no strict toggle
              what: --strict flag removed from syde sync check and syde validate; every invocation is strict by definition
              why: Strict is the only mode; a toggle invites non-strict runs which leave findings unresolved
              draft:
                priority: must
                rationale: Removing the toggle removes the option to run non-strict.
                req_type: constraint
                source: plan
                statement: The syde sync-check, validate, and plan-complete commands shall not accept a --strict flag and shall always block on any audit finding.
                verification: syde sync check --strict exits non-zero with unknown-flag error; syde sync check without flag exits non-zero on any finding
              tasks:
                - remove-strict-flag-from-synccheck-and-validate
    components:
        extended:
            - id: 0nbc
              slug: skill-installer
              what: skill/SKILL.md and references drop --strict mentions and use Finding vocabulary; hooks.json still renders overlap reminders
              why: Skill docs must match the new strict-only model
              field_changes:
                responsibility: Render and write skill files including the Finding-aware sync check instructions, hooks.json overlap reminder, and CLAUDE.md append rules
              tasks:
                - update-skill-docs-and-hooks-for-strict-audit-vocabulary
            - id: x0pa
              slug: cli-commands
              what: syde sync check and syde validate drop the --strict flag; stale-tree cascade is always blocking
              why: Strict is the only mode; no reason to keep a toggle
              field_changes:
                responsibility: Define and register all syde CLI commands under the root cobra tree, including the requirement overlap gate that blocks unacknowledged creates and the always-strict sync-check gate
              tasks:
                - remove-strict-flag-from-synccheck-and-validate
            - id: lvpn
              slug: audit-engine
              what: Severity enum collapsed to single Finding; all SeverityWarning/SeverityHint producers reclassified; surface detector + contract-coverage rules + parity registry + symmetry test
              why: Concentrates the new enforcement layers under the audit engine
              field_changes:
                responsibility: Produce Findings (a single strict severity level) covering entity field validation, relationship integrity, cycles, tree file references, orphan detection, file drift, contract surface coverage, flow coverage, and planning-post-plan symmetry
              tasks:
                - collapse-severity-enum-to-single-finding-constant
                - reclassify-every-warning-and-hint-producer-as-finding
                - regression-test-audit-severity-enum-has-exactly-one-value
                - add-surface-detector-with-clirestscreenevent-patterns
                - plan-authoring-warns-on-uncovered-contract-surfaces-in-new-or-extended-requirements
                - sync-check-errors-on-active-requirements-with-uncovered-contract-surfaces
                - build-parity-registry-and-symmetry-test
                - plan-authoring-warns-when-contract-changes-lack-flow-coverage
                - verify-and-tighten-contractflowfindings-post-plan-rule
                - audit-requires-non-empty-distinction-on-every-acknowledgement
                - regression-test-for-cli-gate-and-audit-distinction-rule
            - id: u9bj
              slug: entity-model
              what: AuditedOverlaps becomes []AuditedOverlap{Slug, Distinction}; YAML round-trip for legacy string entries
              why: Semantic distinction must be captured per acknowledgement, not just the slug
              field_changes:
                responsibility: BaseEntity + per-kind structs + validation rules + plan/task/relationship types + AuditedOverlap with distinction rationale
              tasks:
                - extend-auditedoverlap-struct-with-distinction-field
---
