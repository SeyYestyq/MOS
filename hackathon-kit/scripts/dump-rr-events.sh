#!/usr/bin/env bash
# Damir — dump full events via REST slice (has args/sock_arg).
# IMPORTANT: history uses /slice/right (registered_at < cursor), NOT left.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="${OUT:-$ROOT/fixtures/raw/events.json}"
RR_BASE_URL="${RR_BASE_URL:-https://localhost:9000}"
SLICE_SIZE="${SLICE_SIZE:-100}"
PAGES="${PAGES:-10}"
# right = past from cursor; left = future
DIRECTION="${RR_SLICE_DIR:-right}"

if [[ -z "${RR_TOKEN:-}" ]]; then
  echo "Set RR_TOKEN to personal Access-Token from RR UI" >&2
  exit 1
fi

mkdir -p "$(dirname "$OUT")"
export OUT RR_BASE_URL SLICE_SIZE PAGES RR_TOKEN DIRECTION

python3 <<'PY'
import json, os, ssl, urllib.parse, urllib.request
from datetime import datetime, timezone

base = os.environ["RR_BASE_URL"].rstrip("/")
token = os.environ["RR_TOKEN"]
out = os.environ["OUT"]
slice_size = int(os.environ["SLICE_SIZE"])
pages = int(os.environ["PAGES"])
direction = os.environ.get("DIRECTION", "right")

ctx = ssl._create_unverified_context()
cursor = datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%S.%fZ")
acc = []

for i in range(1, pages + 1):
    q = urllib.parse.urlencode({"cursor": cursor, "slice_size": slice_size})
    url = f"{base}/api/v1/public-api/runtime-event/slice/{direction}?{q}"
    print(f"[*] page {i} dir={direction} cursor={cursor}")
    req = urllib.request.Request(url, headers={"Access-Token": token, "Accept": "application/json"})
    with urllib.request.urlopen(req, context=ctx, timeout=60) as resp:
        page = json.load(resp)
    evs = page.get("runtime_events") or page.get("RuntimeEvents") or []
    if not evs:
        break
    acc.extend(evs)
    if direction == "left":
        nxt = page.get("left_cursor") or page.get("LeftCursor") or ""
    else:
        nxt = page.get("right_cursor") or page.get("RightCursor") or ""
    if not nxt or nxt == cursor:
        break
    cursor = nxt

with open(out, "w", encoding="utf-8") as f:
    json.dump(acc, f, ensure_ascii=False)
print(f"[+] wrote {len(acc)} RR envelopes → {out}")
print(f"    BFF: DATA_MODE=raw RAW_DUMP_PATH={out} go run .")
print("    Adapter unwraps process_* envelopes (same as live).")
PY
