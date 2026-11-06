"""
STUB FastAPI BFF — реализовать завтра.
- GET /api/activity → fixture или Runtime Radar
- Не коммитить секреты; читать RR_TOKEN из env
"""

from __future__ import annotations

import json
from pathlib import Path

# from fastapi import FastAPI, Query
# from fastapi.middleware.cors import CORSMiddleware

FIXTURE = Path(__file__).resolve().parents[2] / "fixtures" / "sample-activity.json"

# app = FastAPI(title="activity-bff")
# app.add_middleware(CORSMiddleware, allow_origins=["*"], allow_methods=["*"], allow_headers=["*"])


def load_fixture() -> dict:
    return json.loads(FIXTURE.read_text(encoding="utf-8"))


# @app.get("/api/activity")
# def activity(pod: str = Query(...), namespace: str = "default"):
#     data = load_fixture()
#     # TODO: try RR live, on failure return fixture
#     return data


if __name__ == "__main__":
    print("fixture keys:", list(load_fixture().keys()))
    print("stub only — wire FastAPI tomorrow")
