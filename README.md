# syde

Text-first software design model for source code repositories. Works as a standalone CLI and as a Claude Code skill that enforces architectural constraints during development.

syde stores your system's architecture as markdown files in `.syde/` — requirements, components, contracts, flows, decisions, and more. A BadgerDB index enables fast queries. When integrated with Claude Code or Codex, syde auto-loads your architecture at session start, enforces design constraints before code changes, and tracks plans and tasks throughout implementation.

## Install

```bash
# Build from source
go build -o syde ./cmd/syde/
go build -o syded ./cmd/syded/

# Move to PATH
mv syde syded /usr/local/bin/
```

## Quick Start

```bash
# Initialize in your project
cd your-project
syde init --install-skill

# Create your design model
syde add system "My App" --description "What this system does"
syde add requirement "Initial architecture" \
  --statement "Track the app architecture with requirements, entities, and flows." \
  --source manual \
  --add-rel my-app:belongs_to
syde add component "API Server" \
  --description "HTTP server" \
  --purpose "Serve product API requests" \
  --responsibility "Request routing" \
  --capability "Routes HTTP requests" \
  --add-rel my-app:belongs_to \
  --add-rel initial-architecture:references
syde add decision "Use PostgreSQL" \
  --statement "All data stored in PostgreSQL" \
  --rationale "Team expertise" \
  --add-rel my-app:belongs_to \
  --add-rel initial-architecture:references

# Check your model
syde status
syde validate
```

## Using syde with Claude Code

This is the primary use case. syde acts as a design guardrail for Claude — it loads architecture context automatically, enforces constraints, and tracks work.

### Setup

```bash
cd your-project
syde init --install-skill
```

This creates:
- `.syde/` — design model directory with entity subdirectories and BadgerDB index
- `.claude/skills/syde/SKILL.md` — skill that controls Claude's behavior
- `.claude/skills/syde/references/` — entity specs, command reference, workflow docs
- `.claude/hooks/syde-hooks.json` — hooks for session start and post-write verification
- `CLAUDE.md` — mandatory rules Claude must follow

### What happens in a Claude Code session

**1. Session starts — architecture auto-loads**

The SessionStart hook runs `syde context --json` and injects your full architecture into Claude's context: all requirements, components, contracts, flows, decisions, and learnings. Claude starts every session already knowing your system's design.

**2. Claude asks clarifying questions first**

The skill instructs Claude to list all questions with recommended answers before doing anything. Claude will present something like:

```
Before I start, I have these questions:
1. Should the API use REST or GraphQL? (Recommended: REST)
2. Do we need authentication? (Recommended: No for MVP)
3. Should I use an ORM? (Recommended: No, raw SQL)
Please confirm or adjust.
```

Claude waits for your confirmation before proceeding.

**3. Claude creates a plan and waits for approval**

Claude creates design entities first, then builds a structured plan:

```bash
# Claude runs these automatically:
syde plan create "Add User Auth"
syde plan add-entity add-user-auth component "Auth Middleware" \
  --description "JWT validation" --responsibility "Token verification"
syde plan add-entity add-user-auth contract "Auth API" \
  --description "Login/logout endpoints" --contract-kind api
syde plan add-phase add-user-auth --description "Phase 1: Implement middleware"
syde plan add-phase add-user-auth --description "Phase 2: Add JWT token handling"
syde plan estimate add-user-auth  # Checks if plan needs splitting
syde plan show add-user-auth      # Presents plan to you
```

Claude then tells you: *"Here is the plan. Run `syde plan approve add-user-auth` to proceed."*

**Claude does NOT write code until you approve.**

**4. Claude implements with task tracking**

After approval, Claude works through each task:

```bash
syde task start create-auth-middleware    # Marks task active
# ... writes code ...
syde constraints check src/middleware/auth.go  # Verifies file maps to component
syde task done create-auth-middleware     # Marks complete, updates plan progress
syde plan phase add-user-auth phase_1 --status completed
```

**5. Claude captures learnings**

When Claude discovers undocumented behavior during implementation:

```bash
syde remember "JWT tokens must be refreshed before each API call" \
  --category constraint --entity auth-middleware --confidence high
```

These learnings persist across sessions — next time Claude (or anyone) works on the auth component, the constraint surfaces automatically.

**6. Post-write verification**

The PostToolUse hook fires after every file write. If a new source file isn't mapped to any syde component, Claude gets a reminder to update the design model.

### Workflow summary

```
Session start
  │ (syde context auto-loaded)
  ▼
Clarify requirements
  │ (questions + recommendations → user confirms)
  ▼
Create plan with entity drafts
  │ (syde plan create + add-entity + add-phase)
  ▼
Present plan
  │ (syde plan create + add-phase + estimate + show)
  ▼
Wait for approval ← USER APPROVES HERE
  │ (syde plan approve)
  ▼
Implement with task tracking
  │ (syde task start → write code → constraints check → task done)
  ▼
Validate and capture learnings
  │ (syde validate + syde remember)
  ▼
Done
```

Plan approval creates a plan-sourced requirement and links the plan to it.
Codex hooks additionally capture each user prompt as a user-sourced
requirement. Requirements are append-only: mark conflicts as superseded or
obsolete instead of deleting history.

### Large plans — sub-plan splitting

When `syde plan estimate` detects a plan with >10 phases, it recommends splitting into sub-plans. Each sub-plan fits in one Claude session turn. Claude implements one sub-plan per turn, and you can resume across sessions.

## Entity Types

syde models your architecture with 11 entity types:

| Kind | What it represents | Key fields |
|------|-------------------|------------|
| **system** | The overall product/service | scope, design_principles, quality_goals |
| **component** | A service, module, or major part | responsibility, boundaries, failure_modes |
| **contract** | An API, event, or protocol boundary | contract_kind, interaction_pattern, input/output |
| **concept** | A domain object or business entity | meaning, lifecycle, invariants |
| **flow** | An end-to-end workflow or user journey | trigger, goal, narrative (step-by-step) |
| **decision** | An architectural decision (ADR) | statement, rationale, tradeoffs |
| **requirement** | Append-only user/plan/migration intent | statement, source, status |
| **plan** | A tracked implementation plan | phases with action/status |
| **task** | A work item linked to plans/entities | priority, status, entity_refs |
| **design** | A UI mockup in UIML format | design_type, UIML body |
| **learning** | Captured design knowledge | category, confidence, entity_refs |

Each entity is a markdown file with YAML frontmatter in `.syde/<kind-plural>/<slug>.md`. Human-readable, git-friendly, editable by hand or via CLI.

## CLI Reference

### Architecture overview

```bash
syde context                    # Full snapshot: entities, decisions, learnings, plans, tasks
syde context --json             # Machine-readable (used by session start hook)
syde status                     # Entity counts
syde validate                   # Check integrity (broken refs, missing fields)
```

### Entity CRUD

```bash
# Create
syde add component "Auth Service" \
  --description "Handles authentication" \
  --responsibility "JWT tokens, OAuth2" \
  --boundaries "Does not own user profiles" \
  --tag security

# Read
syde get auth-service           # Full entity details
syde list components            # List all components
syde list --status active       # Filter by status
syde search "auth"              # Full-text search

# Update (supports ALL kind-specific fields)
syde update auth-service \
  --responsibility "JWT, OAuth2, session management" \
  --add-tag sso \
  --body "## Details\n\nThe auth service handles..."

# Delete
syde remove auth-service        # With confirmation
syde remove auth-service --force
```

### Relationships

```bash
syde update auth-service --add-rel "comp_xyz:depends_on"
syde update auth-service --remove-rel "comp_xyz"
syde graph auth-service         # Show connections
syde graph auth-service --format dot  # Graphviz output
```

Relationship types: `belongs_to`, `depends_on`, `exposes`, `consumes`, `uses`, `involves`, `references`, `applies_to`, `modifies`, `visualizes`

### Rich queries

```bash
syde query auth-service                    # Entity + relationships + learnings + tasks + decisions
syde query auth-service --full             # Everything including body
syde query auth-service --full --format json  # Machine-readable with file:line refs
syde query --kind component --tag security # Filter
syde query --impacts auth-service          # What breaks if this changes
syde query --flow login --components       # Flow decomposition with component details
syde query --related-to auth-service       # All direct connections
syde query --diff auth-service --since 7d  # Git change history
```

### Plans

```bash
syde plan create "Add Payment Processing"
syde plan add-phase add-payment-processing \
  --description "Create payment service" \
  --action create \
  --entity-kind component \
  --entity-name "Payment Service"
syde plan estimate add-payment-processing  # Size check + split recommendation
syde plan show add-payment-processing      # View phases with progress
syde plan approve add-payment-processing   # Approve before implementation
syde plan phase add-payment-processing phase_1 --status completed  # Mark phase done
syde plan execute add-payment-processing   # Auto-scaffold entity files for pending phases
```

### Tasks

```bash
syde task create "Build payment webhook" --plan add-payment-processing --priority high
syde task start build-payment-webhook
syde task done build-payment-webhook       # Auto-updates linked plan phase
syde task block build-payment-webhook --reason "Waiting for Stripe API key"
syde task sub build-payment-webhook "Add signature verification"  # Subtask
syde task link build-payment-webhook payment-service  # Link to entity
syde task list
```

### Learnings

```bash
# Capture
syde remember "Payment webhooks retry up to 5 times with exponential backoff" \
  --category constraint \
  --entity payment-service \
  --confidence high

# Query
syde learn list                         # All learnings
syde learn about payment-service        # Learnings for a specific entity
syde learn search "webhook"             # Search text
syde learn stale                        # Learnings referencing changed entities
syde learn promote webhook-retry --to decision  # Promote to formal decision
```

Learning categories: `gotcha`, `constraint`, `convention`, `context`, `dependency`, `performance`, `workaround`

### Constraints

```bash
syde constraints                        # Active decisions + critical learnings
syde constraints --json                 # For hook injection
syde constraints check src/auth/handler.go  # Map file → component → constraints
```

File-to-component mapping uses `component_paths` in `syde.yaml`:

```yaml
component_paths:
  auth-service: ["src/auth/**", "pkg/auth/**"]
  api-gateway: ["src/gateway/**"]
  web-app: ["web/**"]
```

### UI Designs (UIML)

```bash
syde design create "Dashboard" --type screen
syde design show dashboard              # ASCII art render
syde design preview dashboard           # Open HTML in browser
syde design export dashboard --format html
syde design validate dashboard          # Check UIML syntax
syde design link dashboard web-app      # Link to component
```

UIML uses HTML-like syntax:

```html
<screen name="Dashboard">
  <navbar direction="horizontal">
    <logo>My App</logo>
    <nav active>Home</nav>
    <nav>Settings</nav>
  </navbar>
  <layout direction="horizontal">
    <sidebar width="240">
      <menu>
        <item icon="dashboard" active>Overview</item>
        <item icon="users">Users</item>
      </menu>
    </sidebar>
    <main>
      <heading>Dashboard</heading>
      <grid cols="3" gap="16">
        <card>
          <metric label="Users" value="1,234" />
          <trend direction="up">+12%</trend>
        </card>
      </grid>
    </main>
  </layout>
</screen>
```

### Dashboard

```bash
syde server start                       # Start on port 5703
syde server stop                        # Stop
syde open                               # Start + register project + open browser
```

Dashboard at `http://localhost:5703/<project-slug>` shows: entity overview, plans with progress, learnings, tasks, and design previews.

### Sync with existing codebase

```bash
syde sync                               # Analyze codebase + check model alignment
syde sync --dry-run                     # Preview only
syde sync --coverage                    # Check which directories map to components
```

Sync generates a `scan-guide.json` with directory structure and language detection. For new projects, the syde skill in Claude Code uses this guide to drive 5 rounds of agent-powered extraction. For existing models, it verifies entities against the implementation and detects drift.

### Memory sync

```bash
syde memory sync                        # Generate Claude Code memory files from learnings
syde memory list                        # Show memory files
syde memory clean                       # Remove all syde memories
```

Syncs learnings to `.claude/projects/<hash>/memory/` so they persist across Claude Code sessions.

## `.syde/` directory structure

```
.syde/
├── syde.yaml              # Project config
├── index/                 # BadgerDB (gitignored, rebuildable)
├── systems/               # System entities
├── components/            # Component entities
├── contracts/             # Contract entities
├── concepts/              # Concept entities
├── flows/                 # Flow entities
├── decisions/             # Decision entities
├── plans/                 # Plan entities
├── tasks/                 # Task entities
├── designs/               # Design entities (UIML)
└── learnings/             # Captured learnings
```

Add to `.gitignore`:
```
.syde/index/
```

Everything else in `.syde/` should be committed — it's your architecture documentation.

## License

MIT
