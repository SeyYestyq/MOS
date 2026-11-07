#!/usr/bin/env bash
# Damir B5 — one working historical curl for Denis. Default: /slice/right
set -euo pipefail

RR_BASE_URL="${RR_BASE_URL:-https://localhost:9000}"
SLICE_SIZE="${SLICE_SIZE:-5}"
DIRECTION="${RR_SLICE_DIR:-right}"

if [[ -z "${RR_TOKEN:-}" ]]; then
  echo "export RR_TOKEN=<Access-Token from RR UI>" >&2
  exit 1
fi

CURSOR="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
URL="${RR_BASE_URL%/}/api/v1/public-api/runtime-event/slice/${DIRECTION}?cursor=${CURSOR}&slice_size=${SLICE_SIZE}"

echo "[*] GET $URL"
HTTP_CODE="$(curl -sk -o /tmp/rr-smoke.json -w '%{http_code}' \
  -H "Access-Token: ${RR_TOKEN}" \
  -H "Accept: application/json" \
  "$URL")"

echo "HTTP $HTTP_CODE"
python3 - <<'PY'
import json
d=json.load(open("/tmp/rr-smoke.json"))
evs=d.get("runtime_events") or d.get("RuntimeEvents") or []
print("events_in_page:", len(evs))
print("left_cursor:", d.get("left_cursor") or d.get("LeftCursor"))
print("right_cursor:", d.get("right_cursor") or d.get("RightCursor"))
if evs and isinstance(evs[0], dict):
    print("sample_keys:", list(evs[0].keys())[:12])
PY

if [[ "$HTTP_CODE" != "200" ]]; then
  echo "[!] smoke failed" >&2
  exit 1
fi
echo "[+] smoke OK (historical dir=$DIRECTION)"
