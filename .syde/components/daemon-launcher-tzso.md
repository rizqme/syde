---
boundaries: Does NOT own syded's lifecycle beyond the initial spawn. Does NOT restart syded if it crashes mid-session.
capabilities:
- GET /health with a short timeout to detect running syded
- Locate the syded binary via SYDED_BIN env var, sibling of current exe, or PATH
- Fork syded with -daemon -port N -idle-timeout 30m and release the child
- Poll /health every 50ms up to 3s for readiness
description: CLI-side helper that auto-launches syded if it isn't running before the client makes its first request
files:
- internal/daemon/daemon.go
- internal/daemon/detach_unix.go
id: COM-0021
kind: component
name: Daemon Launcher
notes:
- 'Codex compatibility follow-up: daemon auto-start now writes syded startup output to ~/.syde/syded.log and includes the log tail when readiness polling fails.'
purpose: Make syded transparent — users never have to start the daemon manually. First syde command in a CI run or fresh clone forks a detached syded and waits until /health responds.
relationships:
- type: belongs_to
  target: syde-5tdt
responsibility: Probe syded at /health and spawn a detached syded process with sensible defaults when the probe fails
slug: daemon-launcher-tzso
updated_at: '2026-04-15T06:02:08Z'
---
