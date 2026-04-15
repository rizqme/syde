#!/usr/bin/env bash
#
# wireframe-shot.sh — capture a PNG of any contract detail page in the
# syded dashboard so you can iterate on UIML wireframe rendering with
# visual feedback.
#
# Usage:
#   bash scripts/wireframe-shot.sh <contract-slug> [out.png]
#   bash scripts/wireframe-shot.sh <contract-slug> --kind contract
#
# Defaults:
#   out.png   /tmp/wireframe-<slug>.png
#   url       http://localhost:5703/<project-slug>/contract/<slug>
#
# Behaviour:
#   - Auto-starts `syde open` and waits for :5703 if the daemon is not
#     listening.
#   - Resolves the active project slug by curling /api/projects and
#     matching path == $PWD (falls back to the first project on miss).
#   - Drives /Applications/Google Chrome.app via --headless to capture
#     a 1440x900 PNG of the live dashboard render.
#
# This script lives in scripts/ because it's a useful dev tool beyond
# the wireframe research itself — a future `syde wireframe shot` CLI
# can wrap the same logic.

set -euo pipefail

SLUG="${1:-}"
OUT="${2:-}"
KIND="contract"

if [[ -z "$SLUG" ]]; then
  echo "usage: $0 <contract-slug> [out.png]" >&2
  exit 2
fi

if [[ -z "$OUT" ]]; then
  OUT="/tmp/wireframe-${SLUG}.png"
fi

CHROME="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
if [[ ! -x "$CHROME" ]]; then
  echo "error: Google Chrome not found at $CHROME" >&2
  exit 3
fi

# Wait for syded to listen on :5703, auto-launching it if necessary.
if ! curl -s --max-time 1 -o /dev/null http://localhost:5703/ ; then
  echo "syded not running — launching via 'syde open' …" >&2
  syde open >/dev/null 2>&1 &
  for _ in 1 2 3 4 5 6 7 8 9 10; do
    sleep 0.5
    if curl -s --max-time 1 -o /dev/null http://localhost:5703/ ; then
      break
    fi
  done
  if ! curl -s --max-time 1 -o /dev/null http://localhost:5703/ ; then
    echo "error: syded still not listening on :5703 after 5s" >&2
    exit 4
  fi
fi

# Resolve the project slug. Match by absolute path so we always pick
# the project the script is run from. Falls back to the first project
# in the list if nothing matches.
PROJECT_JSON=$(curl -s --max-time 2 http://localhost:5703/api/projects)
PROJECT_SLUG=$(printf '%s' "$PROJECT_JSON" | python3 -c "
import json, os, sys
try:
    data = json.loads(sys.stdin.read())
except Exception:
    sys.exit(0)
projects = data.get('projects', [])
cwd = os.getcwd()
for p in projects:
    if p.get('path') == cwd:
        print(p.get('slug', ''))
        sys.exit(0)
if projects:
    print(projects[0].get('slug', ''))
")

if [[ -z "$PROJECT_SLUG" ]]; then
  echo "error: could not resolve project slug from /api/projects" >&2
  exit 5
fi

URL="http://localhost:5703/${PROJECT_SLUG}/${KIND}/${SLUG}"
echo "→ shooting $URL → $OUT" >&2

"$CHROME" \
  --headless \
  --disable-gpu \
  --hide-scrollbars \
  --window-size=1440,900 \
  --virtual-time-budget=5000 \
  --screenshot="$OUT" \
  "$URL" >/dev/null 2>&1

if [[ ! -s "$OUT" ]]; then
  echo "error: chrome did not write $OUT" >&2
  exit 6
fi

# Print the path so callers (or pipelines) can chain on it.
echo "$OUT"
