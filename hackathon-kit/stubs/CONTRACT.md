# API contract (зафиксировать в час 0)

## `GET /api/activity?pod=<name>&namespace=<ns>`

Response = shape of `fixtures/sample-activity.json`.

### Fields

- `processes[]`: id, pid, ppid, binary, role?, cmdline, firstSeen, lastSeen
- `network[]`: id, fromPid, dst, port, proto, label?, t
- `events[]`: id, t, type, pid?, summary, raw?

### Tree rule

Edge `parent → child` iff `child.ppid === parent.pid`. Root: `ppid === 0` or parent missing.

### Fallback

If Runtime Radar unreachable → serve `fixtures/sample-activity.json` and set header `X-Data-Source: fixture`.
