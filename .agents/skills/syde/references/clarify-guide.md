# Clarification Guide

Be critical. The user's first description is always incomplete. Your job is to
find the gaps before they become bugs.

## Principles

1. **Challenge assumptions** — "build a REST API" hides 20 decisions. Surface them.
2. **Identify what's missing** — what the user didn't say matters more than what they did.
3. **Propose constraints** — recommend boundaries proactively. "I suggest we limit this to X because Y."
4. **Ask "what happens when this fails?"** — for every feature and external dependency.
5. **Flag skipped infrastructure** — if the user hasn't mentioned auth, error handling, logging, or testing, ask: "Is that intentional, or should we plan for it?"
6. **Push back on vague scope** — if the user says "simple" or "basic", ask: "Simple in what dimension? What are you willing to cut?"

## By Project Type

### Web / Frontend

- **Rendering**: SSR, SPA, or static? Framework choice?
- **Routing**: File-based, code-based? Nested routes? Protected routes?
- **Auth flow**: Login, signup, password reset, session expiry? OAuth/social login?
- **State management**: Local state, context, external store? What's global vs local?
- **Forms**: Validation strategy? Client-side, server-side, or both? Error display?
- **Accessibility**: WCAG level? Screen reader testing? Keyboard navigation?
- **Responsive**: Breakpoints? Mobile-first? Which devices must work?
- **UI states**: Every view needs: loading, empty, error, success states. Which exist?
- **SEO**: Meta tags? Open Graph? Sitemap? Server rendering for crawlers?
- **Performance**: Bundle size budget? Lazy loading? Image optimization?
- **Design system**: Existing component library? Tailwind, CSS modules, styled-components?

### Backend / API

- **Auth**: JWT, OAuth2, API key, session? Token refresh? Who are the API consumers?
- **Error format**: Structured errors? HTTP status code conventions? Error codes?
- **Pagination**: Cursor-based or offset? Default page size? Max page size?
- **Rate limiting**: Per-user, per-IP, per-endpoint? Limits?
- **Database**: Which DB? Migration tool? Schema versioning strategy?
- **Caching**: Which layers? TTL strategy? Cache invalidation?
- **Idempotency**: Which write endpoints need idempotency keys?
- **Async work**: Background jobs? Queue system? Retry policy?
- **Observability**: Structured logging? Metrics? Tracing? Health check endpoint?
- **Validation**: Input validation at which layer? Sanitization?
- **Versioning**: API versioning strategy (URL, header, none)?

### CLI Tool

- **Platforms**: Which OS? Cross-compilation needs?
- **Config**: File format (YAML, TOML, JSON)? Location (~/.config, ./)? Env vars?
- **Output**: Human-readable default + `--json` flag? Color output? `--quiet`/`--verbose`?
- **Exit codes**: 0 success, 1 general error, 2 usage error? Custom codes?
- **Stdin/stdout/stderr**: Which output goes where? Pipe-friendly?
- **Interactive**: Prompts? Confirmations? `--yes`/`--force` to skip?
- **Subcommands**: Command hierarchy? Global vs subcommand flags?
- **Shell completion**: Bash/zsh/fish completion support?
- **Error messages**: Format? Include fix suggestions? Link to docs?

### Full-Stack

Use both Web and Backend checklists, plus:

- **API contract**: Who defines it first — frontend or backend? Shared types?
- **Monorepo vs separate**: Shared code? Build dependencies?
- **Dev workflow**: Hot reload? Proxy setup? Docker compose?
- **Deployment**: Together or separate? CDN for frontend?
- **Type sharing**: Shared type definitions between frontend and backend?

### Library / SDK

- **Compatibility**: Target language versions? Runtime requirements?
- **Public API**: What's exposed vs internal? Stability guarantees?
- **Error handling**: Exceptions, result types, or error codes?
- **Dependencies**: Zero-dep policy? Minimal deps? Peer deps?
- **Docs**: README examples? API reference? Guides?
- **Versioning**: Semver? Breaking change policy?
- **Distribution**: Package registry (npm, PyPI, crates.io)? Build targets?
- **Testing**: Unit tests? Integration tests? Example-based tests?

### Mobile

- **Platform**: Native (Swift/Kotlin), cross-platform (React Native, Flutter)?
- **Min OS version**: iOS 15+? Android API 26+?
- **Offline**: Offline-first? Local storage strategy? Sync conflict resolution?
- **Push notifications**: Provider? Permission flow? Background handling?
- **Deep linking**: URL scheme? Universal links?
- **Permissions**: Camera, location, contacts? Permission request flow?
- **App store**: Review guidelines compliance? Content restrictions?
- **Performance**: Startup time budget? Memory constraints? Battery impact?

### Data / Pipeline

- **Volume**: Expected data size? Growth rate?
- **Latency**: Real-time, near-real-time, or batch?
- **Correctness**: Exactly-once, at-least-once, or at-most-once?
- **Schema**: Schema evolution strategy? Backward/forward compatibility?
- **Monitoring**: Data quality checks? Alerting on failures?
- **Backfill**: How to reprocess historical data?
- **Security**: PII handling? Encryption at rest/in transit?

## How to Present Questions

Group questions by concern, not by project type. Lead with the most impactful
decisions — the ones that affect everything downstream.

Format:
```
1. [CONCERN] — brief context
   Recommended: your recommendation with reasoning
   Missing: things the user hasn't addressed
   Risk: what goes wrong if we skip this

2. [CONCERN] — ...
```

After the user responds, update your understanding and create the design
entities. Do NOT ask follow-up rounds unless the user's answers reveal new
ambiguity.
