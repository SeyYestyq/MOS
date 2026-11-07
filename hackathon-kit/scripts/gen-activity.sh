#!/usr/bin/env bash
# Damir — 4-phase activity in ns hackathon (Must for RR demo).
# Preflight: Sources process + incoming listen; Allow hackathon; All; Apply.
set -euo pipefail

NS="${NS:-hackathon}"
ATTACKER="${ATTACKER:-attacker-pod}"
TARGET="${TARGET:-target-service}"
CTX="${KUBE_CONTEXT:-}"

k() {
  if [[ -n "$CTX" ]]; then
    kubectl --context "$CTX" "$@"
  else
    kubectl "$@"
  fi
}

echo "[*] 0. Preparing namespace and pods..."
k create namespace "$NS" --dry-run=client -o yaml | k apply -f -
k run "$TARGET" -n "$NS" --image=nginx:alpine --port=80 --expose --dry-run=client -o yaml | k apply -f -
if ! k get pod -n "$NS" "$ATTACKER" >/dev/null 2>&1; then
  k run "$ATTACKER" -n "$NS" --image=alpine:latest --restart=Always -- /bin/sh -c "sleep 3600"
fi
k wait -n "$NS" --for=condition=Ready "pod/$ATTACKER" --timeout=120s

echo "[*] Installing forensic & network tools..."
k exec -n "$NS" "$ATTACKER" -- apk add --no-cache curl wget netcat-openbsd >/dev/null

echo "[*] ФАЗА 1: Reconnaissance..."
k exec -n "$NS" "$ATTACKER" -- /bin/sh -c "ls -la /; sleep 2; cat /etc/passwd; sleep 1; id"

echo "[*] ФАЗА 2: Deep process tree..."
k exec -n "$NS" "$ATTACKER" -- /bin/sh -c "sh -c 'sleep 1 && ls /tmp && sh -c \"whoami && sleep 2\"'"

echo "[*] ФАЗА 3: Network (curl/wget → tcp_connect)..."
k exec -n "$NS" "$ATTACKER" -- /bin/sh -c "
  curl -s http://${TARGET}.${NS}.svc.cluster.local > /dev/null || true
  sleep 2
  wget -qO- http://${TARGET}.${NS}.svc.cluster.local > /dev/null || true
"

echo "[*] ФАЗА 4: Threat nc:8888 (TA0011/T1571) + flush ls..."
k exec -n "$NS" "$ATTACKER" -- /bin/sh -c "nc -vv -l -p 8888 &"
sleep 2
# Separate ls flushes event buffer to DB (PT detector guide)
k exec -n "$NS" "$ATTACKER" -- ls -al >/dev/null

echo "[+] Done. Check Runtime → Events (threat filter). Then: scripts/dump-rr-events.sh"
